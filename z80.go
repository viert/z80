package z80

/*

	Go port of Gabriel Gambetta's brilliant libz80 library.
	Original C version can be found at https://github.com/ggambetta/libz80

*/

//go:generate python generate_instructions.py -s z80.spec -o instructions.go
//go:generate go fmt instructions.go

import (
	"fmt"
)

const (

    T_LENGTH = 285.7142857142857 * 2
    T_STATES_TO_BREAK = 35000

	// Flags for doIncDec()
	ID_INC = false
	ID_DEC = true

	// Flags for enable/disable interrupts
	IE_DI = false
	IE_EI = true

	// Flags for doSetRes()
	SR_RES = false
	SR_SET = true

	// Flags for logical / arithmetic operations
	IA_L = false
	IA_A = true

	// Flags for doArithmetic() - F1 = withCarry, F2 = isSub
	F1_ADC = true
	F1_SBC = true
	F1_ADD = false
	F1_SUB = false

	F2_ADC = false
	F2_SBC = true
	F2_ADD = false
	F2_SUB = true
)

type DataIn func(address uint16) byte
type DataOut func(address uint16, data byte)

type RegisterDump struct {
	AF   uint16
	AF_  uint16
	BC   uint16
	BC_  uint16
	DE   uint16
	DE_  uint16
	HL   uint16
	HL_  uint16
	PC   uint16
	SP   uint16
	IX   uint16
	IY   uint16
	R    byte
	I    byte
	IFF1 bool
	IFF2 bool
}

type Context struct {
	R1   *RegisterSet
	R2   *RegisterSet
	PC   uint16
	R    byte
	I    byte
	IFF1 bool
	IFF2 bool
	IM   byte

	MemoryRead  DataIn
	MemoryWrite DataOut
	IoRead      DataIn
	IoWrite     DataOut

	halted  bool
	TStates uint64

	/* If true, an NMI has been requested. */
	nmiRequested bool
	/* If true, a maskable interrupt has been requested. */
	intRequested bool

	/* If true, defer checking maskable interrupts for one
	 * instruction.  This is used to keep an interrupt from happening
	 * immediately after an IE instruction. */
	deferInt bool

	// Interrupt vector
	intVector byte

	/* If true, then execute the opcode in int_vector. */
	execIntVector bool

	/* Debug assistance */
	LatestInstruction string
	LatestDump        RegisterDump

	opcodes_main Z80OpcodeTable
	opcodes_DD   Z80OpcodeTable
	opcodes_ED   Z80OpcodeTable
	opcodes_FD   Z80OpcodeTable
	opcodes_CB   Z80OpcodeTable
	opcodes_DDCB Z80OpcodeTable
	opcodes_FDCB Z80OpcodeTable

	debug  bool
	stop   bool
    state  string
    bpmode bool

	breakpoints map[uint16]bool
}

func doOff(addr uint16, off int) uint16 {
	return (uint16)(int32(addr) + int32(off))
}

func doComplement_ported(v byte) int {
	if v&0x80 == 0 {
		return int(v)
	}
	v = ^v
	v &= 0x7F
	v++
	return -(int(v))
}

func doComplement(v byte) int {
	return int(int8(v))
}

func NewContext(debug bool) *Context {
	c := new(Context)
	c.debug = debug
	c.stop = false
    c.state = "idle"
    c.bpmode = false
	c.breakpoints = make(map[uint16]bool)
	c.createTables()
	c.R1 = NewRegisterSet()
	c.R2 = NewRegisterSet()
	return c
}

func (c *Context) State() string {
    return c.state
}

func (c *Context) SetBPMode(value bool) {
    c.bpmode = value
}

func (c *Context) GetBPMode() bool {
    return c.bpmode
}

func (c *Context) AddBreakpoint(addr uint16) {
	c.breakpoints[addr] = true
}

func (c *Context) RemoveBreakpoint(addr uint16) {
	if _, found := c.breakpoints[addr]; found {
		delete(c.breakpoints, addr)
	}
}

func (c *Context) GetBreakpoints() []uint16 {
    bpList := make([]uint16, 0)
    for key, _ := range c.breakpoints {
        bpList = append(bpList, key)
    }
    return bpList
}

func (c *Context) Disassemble(addr uint16) (string, uint16) {
	var opcode byte
	var offset uint16 = 0
    var arglen uint16 = 0
	var result string = "OP_INVALID"
	currentTable := &c.opcodes_main
	tableEntries := currentTable.entries
	for {
		opcode = c.read8(addr + offset)
		addr++
		opfunc := tableEntries[opcode].function
		if opfunc != nil {
			entry := tableEntries[opcode]
			switch entry.operandType {
			case OP_NONE:
				result = fmt.Sprintf(entry.format)
			case OP_BYTE:
				dByte := c.read8(addr)
				result = fmt.Sprintf(entry.format, dByte)
                arglen = 1
			case OP_OFFSET:
				dInt := int8(c.read8(addr))
				result = fmt.Sprintf(entry.format, dInt)
                arglen = 1
			case OP_WORD:
				dWord := c.read16(addr)
				result = fmt.Sprintf(entry.format, dWord)
                arglen = 2
			}
			break
		} else if tableEntries[opcode].nextTable != nil {
			currentTable = tableEntries[opcode].nextTable
			tableEntries = currentTable.entries
			offset = uint16(currentTable.opcodeOffset)
		} else {
			break
		}
	}
	return result, addr + offset + arglen + 1
}

func (c *Context) doExecute() {
	currentTable := &c.opcodes_main
	tableEntries := currentTable.entries
	var opcode byte
	var offset int = 0
	for {
		if c.execIntVector {
			opcode = c.intVector
			c.TStates += 6
		} else {
			opcode = c.read8(c.PC + uint16(offset))
			c.PC++
			c.TStates += 1
		}
		c.incr()
		opfunc := tableEntries[opcode].function
		if opfunc != nil {
			if c.debug {
				entry := tableEntries[opcode]
				switch entry.operandType {
				case OP_NONE:
					c.LatestInstruction = fmt.Sprintf(entry.format)
				case OP_BYTE:
					bytedata := c.read8(c.PC)
					c.LatestInstruction = fmt.Sprintf(entry.format, bytedata)
				case OP_OFFSET:
					intdata := int8(c.read8(c.PC))
					c.LatestInstruction = fmt.Sprintf(entry.format, intdata)
				case OP_WORD:
					worddata := c.read16(c.PC)
					c.LatestInstruction = fmt.Sprintf(entry.format, worddata)
				}
			}
			c.PC -= uint16(offset)
			opfunc()
            if c.debug {
				c.copyDump()
            }
			c.PC += uint16(offset)

			break
		} else if tableEntries[opcode].nextTable != nil {
			currentTable = tableEntries[opcode].nextTable
			tableEntries = currentTable.entries
			offset = currentTable.opcodeOffset
			if offset > 0 {
				c.decr()
			}
		} else {
			break
		}
	}
}

func (c *Context) copyDump() {
	c.LatestDump.PC = c.PC
	c.LatestDump.AF = *c.R1.AF
	c.LatestDump.BC = *c.R1.BC
	c.LatestDump.DE = *c.R1.DE
	c.LatestDump.HL = *c.R1.HL
	c.LatestDump.AF_ = *c.R2.AF
	c.LatestDump.BC_ = *c.R2.BC
	c.LatestDump.DE_ = *c.R2.DE
	c.LatestDump.HL_ = *c.R2.HL
	c.LatestDump.IX = *c.R1.IX
	c.LatestDump.IY = *c.R1.IY
	c.LatestDump.SP = *c.R1.SP
	c.LatestDump.I = c.I
	c.LatestDump.R = c.R
	c.LatestDump.IFF1 = c.IFF1
	c.LatestDump.IFF2 = c.IFF2
}

func (c *Context) Stop() {
	c.stop = true
}

func (c *Context) Resume() {
	c.stop = false
    c.state = "idle"
}

func (c *Context) Execute() {
    c.state = "running"
	if c.nmiRequested {
		c.doNmi()
	} else if c.intRequested && !c.deferInt && c.IFF1 {
		c.doInt()
	} else {
		c.deferInt = false
		c.doExecute()
	}
    c.state = "stopped"
}

func (c *Context) ExecuteTStates(tstates uint64) uint64 {
	c.TStates = 0
	for c.TStates < tstates {
		if c.stop {
            c.state = "stopped"
			break
		}
		if c.bpmode && c.breakpoints[c.PC] {
            c.state = fmt.Sprintf("stopped at brkp %04X", c.PC)
			c.stop = true
			break
		}
		c.Execute()
	}
	return c.TStates
}

func (c *Context) Reset() {
	c.PC = 0
	*c.R1.F = 0
	c.IM = 0
	c.IFF1 = false
	c.IFF2 = false
	c.R = 0
	c.I = 0
	c.halted = false
	c.TStates = 0
	c.nmiRequested = false
	c.intRequested = false
	c.deferInt = false
	c.execIntVector = false
}

func (c *Context) unhalt() {
	if c.halted {
		c.halted = false
		c.PC++
	}
}

func (c *Context) read8(addr uint16) byte {
	return c.MemoryRead(addr)
}

func (c *Context) read16(addr uint16) uint16 {
	low := c.read8(addr)
	high := c.read8(addr + 1)
	return uint16(high)<<8 | uint16(low)
}

func (c *Context) write8(addr uint16, data byte) {
	c.MemoryWrite(addr, data)
}

func (c *Context) write16(addr uint16, data uint16) {
	low := byte(data & 0xFF)
	high := byte((data >> 8) & 0xFF)
	c.MemoryWrite(addr, low)
	c.MemoryWrite(addr+1, high)
}

func (c *Context) doAddWord(a1 uint16, a2 uint16, withCarry bool, isSub bool) uint16 {
	if withCarry && c.getFlag(F_C) {
		a2++
	}
	var sum int = int(a1)
	if isSub {
		sum -= int(a2)
		c.valFlag(F_H, ((a1&0x0FFF)-(a2&0x0FFF))&0x1000 != 0)
	} else {
		sum += int(a2)
		c.valFlag(F_H, ((a1&0x0FFF)+(a2&0x0FFF))&0x1000 != 0)
	}
	c.valFlag(F_C, (sum&0x10000) != 0)
	if withCarry || isSub {
		var minuedSign int = int(a1 & 0x8000)
		var subtrahendSign int = int(a2 & 0x8000)
		var resultSign int = int(sum & 0x8000)
		var overflow bool
		if isSub {
			overflow = minuedSign != subtrahendSign && resultSign != minuedSign
		} else {
			overflow = minuedSign == subtrahendSign && resultSign != minuedSign
		}
		c.valFlag(F_PV, overflow)
		c.valFlag(F_S, (sum&0x8000) != 0)
		c.valFlag(F_Z, sum == 0)
	}
	c.valFlag(F_N, isSub)
	c.adjustFlags(byte(sum >> 8))
	return uint16(sum)
}

func (c *Context) incr() {
	c.R = c.R&0x80 | ((c.R + 1) & 0x7F)
}

func (c *Context) decr() {
	c.R = c.R&0x80 | ((c.R - 1) & 0x7F)
}

func (c *Context) doArithmetic(value byte, withCarry bool, isSub bool) byte {
	var res uint16 // to detect carry

	if isSub {
		c.setFlag(F_N)
		c.valFlag(F_H, (((*c.R1.A&0x0F)-(value&0x0F))&0x10) != 0)
		res = uint16(*c.R1.A) - uint16(value)
		if withCarry && c.getFlag(F_C) {
			res--
		}
	} else {
		c.resFlag(F_N)
		c.valFlag(F_H, (((*c.R1.A&0x0F)+(value&0x0F))&0x10) != 0)
		res = uint16(*c.R1.A) + uint16(value)
		if withCarry && c.getFlag(F_C) {
			res++
		}
	}
	c.valFlag(F_S, (res&0x80) != 0)
	c.valFlag(F_C, (res&0x100) != 0)
	c.valFlag(F_Z, (res&0xFF) == 0)

	var minuedSign int = int(*c.R1.A & 0x80)
	var subtrahendSign int = int(value & 0x80)
	var resultSign int = int(res & 0x80)
	var overflow bool
	if isSub {
		overflow = minuedSign != subtrahendSign && resultSign != minuedSign
	} else {
		overflow = minuedSign == subtrahendSign && resultSign != minuedSign
	}
	c.valFlag(F_PV, overflow)
	c.adjustFlags(byte(res & 0xFF))
	return byte(res & 0xFF)
}

func (c *Context) doAND(value byte) {
	*c.R1.A &= value
	c.adjustLogicFlag(true)
}

func (c *Context) doOR(value byte) {
	*c.R1.A |= value
	c.adjustLogicFlag(false)
}

func (c *Context) doXOR(value byte) {
	*c.R1.A ^= value
	c.adjustLogicFlag(false)
}

func (c *Context) doBIT(b byte, val byte) {
	if val&(1<<b) != 0 {
		c.resFlag(F_Z | F_PV)
	} else {
		c.setFlag(F_Z | F_PV)
	}

	c.setFlag(F_H)
	c.resFlag(F_N)

	c.resFlag(F_S)
	if b == 7 && !c.getFlag(F_Z) {
		c.setFlag(F_S)
	}
}

func (c *Context) doBIT_r(b byte, val byte) {
	c.doBIT(b, val)
	c.valFlag(F_5, Z80Flag(val)&F_5 != 0)
	c.valFlag(F_3, Z80Flag(val)&F_3 != 0)
}

func (c *Context) doBIT_indexed(b byte, address uint16) {
	val := c.read8(address)
	c.doBIT(b, val)
	c.valFlag(F_5, Z80Flag(address>>8)&F_5 != 0)
	c.valFlag(F_3, Z80Flag(address>>8)&F_3 != 0)
}

func (c *Context) doSetRes(bit bool, pos byte, val byte) byte {
	if bit {
		val |= byte(1 << pos)
	} else {
		val &= ^byte(1 << pos)
	}
	return val
}

func (c *Context) doPush(val uint16) {
	*c.R1.SP -= 2
	c.write16(*c.R1.SP, val)
}

func (c *Context) doPop() uint16 {
	val := c.read16(*c.R1.SP)
	*c.R1.SP += 2
	return val
}

func (c *Context) doDAA() {
	correctionFactor := byte(0x00)
	carry := false
	if *c.R1.A > 0x99 || c.getFlag(F_C) {
		correctionFactor |= 0x60
		carry = true
	}
	if (*c.R1.A&0x0F) > 9 || c.getFlag(F_H) {
		correctionFactor |= 0x06
	}
	aBefore := *c.R1.A
	if c.getFlag(F_N) {
		*c.R1.A -= correctionFactor
	} else {
		*c.R1.A += correctionFactor
	}
	c.valFlag(F_H, (aBefore^*c.R1.A)&0x10 != 0)
	c.valFlag(F_C, carry)
	c.valFlag(F_S, (*c.R1.A&0x80) != 0)
	c.valFlag(F_Z, *c.R1.A == 0)
	c.valFlag(F_PV, parityBit[*c.R1.A])
	c.adjustFlags(*c.R1.A)
}

func (c *Context) doCP_HL() byte {
	val := c.read8(*c.R1.HL)
	result := c.doArithmetic(val, false, true)
	c.adjustFlags(val)
	return result
}

func (c *Context) doIncDec(val byte, isDec bool) byte {
	if isDec {
		c.valFlag(F_PV, val&0x80 != 0 && (val-1)&0x80 == 0)
		val--
		c.valFlag(F_H, val&0x0F == 0x0F)
	} else {
		c.valFlag(F_PV, val&0x80 == 0 && (val+1)&0x80 != 0)
		val++
		c.valFlag(F_H, val&0x0F == 0)
	}
	c.valFlag(F_S, val&0x80 != 0)
	c.valFlag(F_Z, val == 0)
	c.valFlag(F_N, isDec)
	c.adjustFlags(val)
	return val
}

func (c *Context) doRLC(adjFlags bool, val byte) byte {
	c.valFlag(F_C, (val&0x80) != 0)
	val <<= 1
	cy := byte(0)
	if c.getFlag(F_C) {
		cy = 1
	}
	val |= cy
	c.adjustFlags(val)
	c.resFlag(F_H | F_N)
	if adjFlags {
		c.adjustFlagSZP(val)
	}
	return val
}

func (c *Context) doRRC(adjFlags bool, val byte) byte {
	c.valFlag(F_C, (val&0x01) != 0)
	val >>= 1
	cy := byte(0)
	if c.getFlag(F_C) {
		cy = 1
	}
	val |= (cy << 7)
	c.adjustFlags(val)
	c.resFlag(F_H | F_N)
	if adjFlags {
		c.adjustFlagSZP(val)
	}
	return val
}

func (c *Context) doRL(adjFlags bool, val byte) byte {
	cy := byte(0)
	if c.getFlag(F_C) {
		cy = 1
	}
	c.valFlag(F_C, (val&0x80) != 0)
	val <<= 1
	val |= cy
	c.adjustFlags(val)
	c.resFlag(F_H | F_N)
	if adjFlags {
		c.adjustFlagSZP(val)
	}
	return val
}

func (c *Context) doRR(adjFlags bool, val byte) byte {
	cy := byte(0)
	if c.getFlag(F_C) {
		cy = 1
	}
	c.valFlag(F_C, (val&0x01) != 0)
	val >>= 1
	val |= (cy << 7)
	c.adjustFlags(val)
	c.resFlag(F_H | F_N)
	if adjFlags {
		c.adjustFlagSZP(val)
	}
	return val

}

func (c *Context) doSL(val byte, isArith bool) byte {
	c.valFlag(F_C, (val&0x80) != 0)
	val <<= 1
	if !isArith {
		val |= 1
	}
	c.adjustFlags(val)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(val)
	return val
}

func (c *Context) doSR(val byte, isArith bool) byte {
	b := val & 0x80
	c.valFlag(F_C, (val&0x01) != 0)
	val >>= 1
	if isArith {
		val |= b
	}
	c.adjustFlags(val)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(val)
	return val
}
