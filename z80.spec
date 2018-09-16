#
#   Z80 instructions code generation
#   Go port of the brilliant libz80 library by Gabriel Gambetta
#   https://github.com/ggambetta/libz80

#
#   Add / Sub / Adc / Sbc
#
[ADC,SBC,ADD,SUB] A,[A,B,C,D,E,H,L,IXh,IXl,IYh,IYl]
    *c.R1.A = c.doArithmetic(*c.R1.$2, F1_$1, F2_$1)

[ADC,SBC,ADD,SUB] A,(HL)
    *c.R1.A = c.doArithmetic(c.read8(*c.R1.HL), F1_$1, F2_$1)

[ADC,SBC,ADD,SUB] A,([IX,IY]+d)
    c.TStates += 5
    displacement := int8(c.read8(c.PC))
    c.PC++
    addr := uint16(int16(*c.R1.$2) + int16(displacement))
    *c.R1.A = c.doArithmetic(c.read8(addr), F1_$1, F2_$1)

[ADC,SBC,ADD,SUB] A,n
	*c.R1.A = c.doArithmetic(c.read8(c.PC), F1_$1, F2_$1)
	c.PC++

# Particular cases
[ADC,ADD,SBC] HL,[SP,BC,DE,HL]
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.$2, F1_$1, F2_$1)

ADD [IX,IY],[SP,BC,DE,IX,IY]
	c.TStates += 7
	*c.R1.$1 = c.doAddWord(*c.R1.$1, *c.R1.$2, false, false)

#
# Logic operations
#
[AND,XOR,OR] (HL)
	c.do$1(c.read8(*c.R1.HL))

[AND,XOR,OR] ([IX,IY]+d)
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int32(*c.R1.$2) + int32(displacement))
	c.do$1(c.read8(addr))

[AND,XOR,OR] [A,B,C,D,E,H,L,IXh,IXl,IYh,IYl]
	c.do$1(*c.R1.$2)

[AND,XOR,OR] n
	c.do$1(c.read8(c.PC))
	c.PC++

#
# Bit operations
#
BIT [0-7],[A,B,C,D,E,H,L]
	c.doBIT_r($1, *c.R1.$2);

BIT [0-7],(HL)
	c.TStates += 1
	c.doBIT_r($1, c.read8(*c.R1.HL))

BIT [0-7],([IX,IY]+d)
	c.TStates += 2
	displacement := int8(c.read8(c.PC))
	address := uint16(int32(*c.R1.$2) + int32(displacement))
	c.PC++
	c.doBIT_indexed($1, address)

[SET,RES] [0-7],[A,B,C,D,E,H,L]
	*c.R1.$3 = c.doSetRes(SR_$1, $2, *c.R1.$3)

[SET,RES] [0-7],(HL)
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_$1, $2, c.read8(*c.R1.HL)))

[SET,RES] [0-7],([IX,IY]+d)
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.$3, off)
	c.write8(address, c.doSetRes(SR_$1, $2, c.read8(address)))
	

#
# Jumps and calls
#
CALL [C,M,NZ,NC,P,PE,PO,Z],(nn)
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_$1) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}

CALL (nn)
	addr := c.read16(c.PC)
	c.PC += 2
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = addr

JP ([HL,IX,IY])
	c.PC = *c.R1.$1
	
JP [C,M,NZ,NC,P,PE,PO,Z],(nn)
	addr := c.read16(c.PC)
	if c.condition(C_$1) {
		c.PC = addr
	} else {
		c.PC += 2
	}

JP (nn)
	addr := c.read16(c.PC)
	c.PC = addr
	
JR [C,NZ,NC,Z],(PC+e)
	off := doComplement(c.read8(c.PC))
	c.PC++
	if c.condition(C_$1) {
		c.TStates += 5
		c.PC = doOff(c.PC, off)
	} 

JR (PC+e)
	off := doComplement(c.read8(c.PC))
	c.PC++
	c.TStates += 5
	c.PC = doOff(c.PC, off)

RETI
	c.IFF1 = c.IFF2
	c.ret()
		
RETN
	c.IFF1 = c.IFF2
	c.ret()


RET [C,M,NZ,NC,P,PE,PO,Z]
	c.TStates += 1
	if c.condition(C_$1) {
		c.ret()
	}
		
RET
	c.PC = c.doPop();

	
DJNZ (PC+e)
	c.TStates += 1
	off := doComplement(c.read8(c.PC))
	c.PC++
	*c.R1.B--
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC = doOff(c.PC, off)
	}


RST [0,8,10,18,20,28,30,38]H
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x0$1

#
# Misc
#
CCF
	c.valFlag(F_C, !c.getFlag(F_C))
	c.resFlag(F_N)
	c.adjustFlags(*c.R1.A);

SCF
	c.setFlag(F_C)
	c.resFlag(F_N|F_H)
	c.adjustFlags(*c.R1.A);

CPL
	*c.R1.A = ^*c.R1.A
	c.setFlag(F_N|F_H)
	c.adjustFlags(*c.R1.A);
	
DAA
	c.doDAA();
	
EX (SP),[HL,IX,IY]
	c.TStates += 3
	tmp := c.read16(*c.R1.SP)
	c.write16(*c.R1.SP, *c.R1.$1)
	*c.R1.$1 = tmp

EX AF,AF'
	tmp := *c.R1.AF
	*c.R1.AF = *c.R2.AF
	*c.R2.AF = tmp

EX DE,HL
	tmp := *c.R1.DE
	*c.R1.DE = *c.R1.HL
	*c.R1.HL = tmp

EXX
	var tmp uint16
	
	tmp = *c.R1.BC
	*c.R1.BC = *c.R2.BC
	*c.R2.BC = tmp

	tmp = *c.R1.DE
	*c.R1.DE = *c.R2.DE
	*c.R2.DE = tmp

	tmp = *c.R1.HL
	*c.R1.HL = *c.R2.HL
	*c.R2.HL = tmp

HALT
	c.halted = true
	c.PC--
	
#
# Compare
#
CP (HL)
	c.doCP_HL();

CP ([IX,IY]+d)
	c.TStates += 5;
	displacement := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$1, displacement)
	val := c.read8(addr)
	c.doArithmetic(val, false, true)
	c.adjustFlags(val)

CP [A,B,C,D,E,H,L,IXh,IXl,IYh,IYl]
	c.doArithmetic(*c.R1.$1, false, true)
	c.adjustFlags(*c.R1.$1)

CP n
	val := c.read8(c.PC)
	c.PC++
	c.doArithmetic(val, false, true)
	c.adjustFlags(val)

CPDR
	c.cpd()
	if *c.R1.BC != 0 && !c.getFlag(F_Z) {
		c.TStates += 5
		c.PC -= 2
	}

CPD
	c.TStates += 5
	carry := c.getFlag(F_C)
	value := c.doCP_HL()

	if c.getFlag(F_H) {
		value--
	}
	*c.R1.HL--
	*c.R1.BC--
	c.valFlag(F_PV, *c.R1.BC != 0)
	c.valFlag(F_C, carry)
	c.valFlag(F_5, value & (1<<1) != 0)
	c.valFlag(F_3, value & (1<<3) != 0)

CPIR
	c.cpi()
	if *c.R1.BC != 0 && !c.getFlag(F_Z) {
		c.TStates += 5
		c.PC -= 2
	}

CPI
	c.TStates += 5
	carry := c.getFlag(F_C)
	value := c.doCP_HL()
	if c.getFlag(F_H) {
		value--
	}
	*c.R1.HL++
	*c.R1.BC--
	c.valFlag(F_PV, *c.R1.BC != 0)
	c.valFlag(F_C, carry)
	c.valFlag(F_5, value & (1<<2) != 0)
	c.valFlag(F_3, value & (1<<3) != 0)

#
# INC and DEC
#
[INC,DEC] (HL)
	c.TStates += 1
	value := c.read8(*c.R1.HL)
	c.write8(*c.R1.HL, c.doIncDec(value, ID_$1))

[INC,DEC] ([IX,IY]+d)
	c.TStates += 6
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$2, off)
	value := c.read8(addr)
	c.write8(addr, c.doIncDec(value, ID_$1))
	
[INC,DEC] [A,B,C,D,E,H,L,IXh,IXl,IYh,IYl]
	*c.R1.$2 = c.doIncDec(*c.R1.$2, ID_$1)

INC [BC,DE,HL,SP,IX,IY]
	c.TStates += 2
	*c.R1.$1++

DEC [BC,DE,HL,SP,IX,IY]
	c.TStates += 2
	*c.R1.$1--

#
# Interrupts
#
[EI,DI]
	c.IFF1 = IE_$1
	c.IFF2 = IE_$1
	c.deferInt = true

IM [012]
	c.IM = $1


#
# IO ports
#
IN [A,B,C,D,E,F,H,L],(C)
	*c.R1.$1 = c.ioread(*c.R1.BC)
	c.resFlag(F_H|F_N)
	c.adjustFlagSZP(*c.R1.$1)
	c.adjustFlags(*c.R1.$1)

IN A,(n)
	port := uint16(c.read8(c.PC))
	c.PC++
	port = uint16(*c.R1.A)<<8 | port 
	*c.R1.A = c.ioread(port)

INDR
	c.ind()
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC -= 2
	}

IND
	c.TStates += 1
	val := c.ioread(*c.R1.BC)
	c.write8(*c.R1.HL, val)
	*c.R1.HL--
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)
	c.valFlag(F_N, val&0x80 != 0)
	flagval := int(val) + int((*c.R1.C - 1) & 0xFF)
	c.valFlag(F_H, flagval > 0xFF)
	c.valFlag(F_C, flagval > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flagval&7) ^ *c.R1.B])

INIR
	c.ini()
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC -= 2
	}

INI
	c.TStates += 1
	val := c.ioread(*c.R1.BC)
	c.write8(*c.R1.HL, val)
	*c.R1.HL++
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)
	c.valFlag(F_N, val&0x80 != 0)
	flagval := int(val) + int((*c.R1.C + 1) & 0xFF)
	c.valFlag(F_H, flagval > 0xFF)
	c.valFlag(F_C, flagval > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flagval&7) ^ *c.R1.B])

OUTI
	c.TStates += 1
	value := c.read8(*c.R1.HL)
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)

	c.iowrite(*c.R1.BC, value)
	*c.R1.HL++

	flag_value := int(value) + int(*c.R1.L)
	c.valFlag(F_N, value&0x80 != 0)
	c.valFlag(F_H, flag_value > 0xFF)
	c.valFlag(F_C, flag_value > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flag_value&7) ^ *c.R1.B])
	c.adjustFlags(*c.R1.B)

OTIR
	c.outi()
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC -= 2
	}

OUTD
	c.TStates += 1
	value := c.read8(*c.R1.HL)
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)

	c.iowrite(*c.R1.BC, value)
	*c.R1.HL--

	flag_value := int(value) + int(*c.R1.L)
	c.valFlag(F_N, value&0x80 != 0)
	c.valFlag(F_H, flag_value > 0xFF)
	c.valFlag(F_C, flag_value > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flag_value&7) ^ *c.R1.B])
	c.adjustFlags(*c.R1.B)

OTDR
	c.outd()
	if *c.R1.B != 0 {
		c.TStates += 5;
		c.PC -= 2;
	}

OUT (C),0
	c.iowrite(*c.R1.BC, 0)
	
OUT (C),[A,B,C,D,E,H,L]
	c.iowrite(*c.R1.BC, *c.R1.$1)

OUT (n),A
	port := uint16(c.read8(c.PC))
	c.PC++
	port = uint16(*c.R1.A)<<8 | port
	c.iowrite(port, *c.R1.A)

#
# Stack
#
POP [AF,BC,DE,HL,IX,IY]
	*c.R1.$1 = c.doPop()

PUSH [AF,BC,DE,HL,IX,IY]
	c.TStates += 1
	c.doPush(*c.R1.$1)

#
# Rotate & shift
#

[RLC,RRC,RL,RR] (HL)
	c.TStates += 1
	c.write8(*c.R1.HL, c.do$1(true, c.read8(*c.R1.HL)))
	
[RLC,RRC,RL,RR] [A,B,C,D,E,H,L,IXh,IXl,IYh,IYl]
	*c.R1.$2 = c.do$1(true, *c.R1.$2)

[RLC,RRC,RL,RR] ([IX,IY]+d)
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$2, off)
	c.write8(addr, c.do$1(true, c.read8(addr)))

[RL,RR,RLC,RRC]A
	*c.R1.A = c.do$1(false, *c.R1.A)
	
RLD
	c.TStates += 4
	ah := *c.R1.A & 0x0F
	hl := c.read8(*c.R1.HL)

	*c.R1.A = (*c.R1.A&0xF0) | ((hl&0xF0) >> 4)
	hl = (hl << 4) | ah

	c.write8(*c.R1.HL, hl)
	c.resFlag(F_H|F_N)
	c.adjustFlagSZP(*c.R1.A)
	c.adjustFlags(*c.R1.A)


RRD
	c.TStates += 4
	ah := *c.R1.A & 0x0F
	hl := c.read8(*c.R1.HL)

	*c.R1.A = (*c.R1.A&0xF0) | (hl & 0x0F)
	hl = (hl >> 4) | (ah << 4)

	c.write8(*c.R1.HL, hl)
	c.resFlag(F_H|F_N)
	c.adjustFlagSZP(*c.R1.A)


[SL,SR][L,A] (HL)
	c.TStates += 1
	c.write8(*c.R1.HL, c.do$1(c.read8(*c.R1.HL), IA_$2))


[SL,SR][L,A] ([IX,IY]+d)
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$3, off)
	c.write8(addr, c.do$1(c.read8(addr), IA_$2))

[SL,SR][L,A] [A,B,C,D,E,H,L,IXh,IXl,IYh,IYl]
	*c.R1.$3 = c.do$1(*c.R1.$3, IA_$2)


#
# Loads
#
LD ([BC,DE,HL]),A
	c.write8(*c.R1.$1, *c.R1.A)

LD (HL),[B,C,D,E,H,L]
	c.write8(*c.R1.HL, *c.R1.$1)

LD (HL),n
	c.write8(*c.R1.HL, c.read8(c.PC))
	c.PC++
	
LD ([IX,IY]+d),[A,B,C,D,E,H,L]
	c.TStates += 5
	addr := doOff(*c.R1.$1, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.$2);
	
LD ([IX,IY]+d),n
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	n := c.read8(c.PC)	
	c.PC++
	addr := doOff(*c.R1.$1, off)
	c.write8(addr, n)
	
LD (nn),A
	c.write8(c.read16(c.PC), *c.R1.A)
	c.PC += 2
	
LD [BC,DE,HL,IX,IY,SP],[BC,DE,HL,IX,IY,SP]
	c.TStates += 2
	*c.R1.$1 = *c.R1.$2

LD (nn),[BC,DE,HL,IX,IY,SP]
	c.write16(c.read16(c.PC), *c.R1.$1)
	c.PC += 2
	
LD A,([BC,DE])
	*c.R1.A = c.read8(*c.R1.$1)

LD [A,B,C,D,E,H,L],(HL)
	*c.R1.$1 = c.read8(*c.R1.HL)

LD [A,B,C,D,E,H,L],([IX,IY]+d)
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$2, off)
	*c.R1.$1 = c.read8(addr)

LD [A,B,C,D,E,H,L],(nn)
	*c.R1.$1 = c.read8(c.read16(c.PC))
	c.PC += 2
	
LD [A,B,C,D,E,H,L,IXh,IXl,IYh,IYl],[A,B,C,D,E,H,L,IXh,IXl,IYh,IYl]
	*c.R1.$1 = *c.R1.$2

LD [A,B,C,D,E,H,L],[SL,SR]A ([IX,IY]+d)
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$3, off)
	*c.R1.$1 = c.do$2(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.$1)
	
LD [A,B,C,D,E,H,L],[SL,SR]L ([IX,IY]+d)
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$3, off)
	*c.R1.$1 = c.do$2(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.$1)
	  
LD [A,B,C,D,E,H,L],[RL,RLC,RR,RRC] ([IX,IY]+d)
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$3, off)
	*c.R1.$1 = c.do$2(true, c.read8(addr))
	c.write8(addr, *c.R1.$1)

LD [A,B,C,D,E,H,L],[SET,RES] [0-7],([IX,IY]+d)
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.$4, off)
	*c.R1.$1 = c.doSetRes(SR_$2, $3, c.read8(addr))
	c.write8(addr, *c.R1.$1)

LD A,[IR]
	c.TStates += 1
	*c.R1.A = c.$1
	c.adjustFlags(*c.R1.A)
	c.resFlag(F_H|F_N)
	c.valFlag(F_PV, c.IFF2)
	c.valFlag(F_S, (*c.R1.A&0x80) != 0)
	c.valFlag(F_Z, *c.R1.A == 0)
	
LD [IR],A
	c.TStates += 1
	c.$1 = *c.R1.A

LD [A,B,C,D,E,H,L,IXh,IXl,IYh,IYl],n
	*c.R1.$1 = c.read8(c.PC)
	c.PC++
	
LD [BC,DE,HL,SP,IX,IY],(nn)
	addr := c.read16(c.PC)
	c.PC += 2
	*c.R1.$1 = c.read16(addr)

LD [BC,DE,HL,SP,IX,IY],nn
	*c.R1.$1 = c.read16(c.PC)
	c.PC += 2
	

LDIR
	c.ldi()
	if *c.R1.BC != 0 {
		c.TStates += 5
		c.PC -= 2
	}

LDI
	c.TStates += 2
	val := c.read8(*c.R1.HL)
	c.write8(*c.R1.DE, val)
	*c.R1.DE++
	*c.R1.HL++
	*c.R1.BC--
	c.valFlag(F_5, (*c.R1.A + val) & 0x02 != 0)
	c.valFlag(F_3, Z80Flag(*c.R1.A + val) & F_3 != 0)
	c.resFlag(F_H|F_N)
	c.valFlag(F_PV, *c.R1.BC != 0)

LDDR
	c.ldd()
	if *c.R1.BC != 0 {
		c.TStates += 5
		c.PC -= 2
	}

LDD
	c.TStates += 2
	val := c.read8(*c.R1.HL)
	c.write8(*c.R1.DE, val)
	*c.R1.DE--
	*c.R1.HL--
	*c.R1.BC--
	c.valFlag(F_5, (*c.R1.A + val) & 0x02 != 0)
	c.valFlag(F_3, Z80Flag(*c.R1.A + val) & F_3 != 0)
	c.resFlag(F_H|F_N)
	c.valFlag(F_PV, *c.R1.BC != 0)

NEG
	temp := *c.R1.A
	*c.R1.A = 0
	*c.R1.A = c.doArithmetic(temp, false, true)
	c.setFlag(F_N)

NOP
	// NOP
	
	
