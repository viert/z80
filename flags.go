package z80

/*

	Go port of Gabriel Gambetta's brilliant libz80 library.
	Original C version can be found at https://github.com/ggambetta/libz80

*/

type Z80Flag byte

const (
	F_C  Z80Flag = 1
	F_N  Z80Flag = 2
	F_PV Z80Flag = 4
	F_3  Z80Flag = 8
	F_H  Z80Flag = 16
	F_5  Z80Flag = 32
	F_Z  Z80Flag = 64
	F_S  Z80Flag = 128
)

type Z80Condition byte

const (
	C_Z Z80Condition = iota
	C_NZ
	C_C
	C_NC
	C_M
	C_P
	C_PE
	C_PO
)

var (
	parityBit = [...]bool{
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true,
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true,
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		false, true, true, false, true, false, false, true, true, false, false, true, false, true, true, false,
		true, false, false, true, false, true, true, false, false, true, true, false, true, false, false, true}
)

func (c *Context) setFlag(flag Z80Flag) {
	*c.R1.F |= byte(flag)
}

func (c *Context) resFlag(flag Z80Flag) {
	*c.R1.F &= byte(^flag)
}

func (c *Context) valFlag(flag Z80Flag, val bool) {
	if val {
		c.setFlag(flag)
	} else {
		c.resFlag(flag)
	}
}

func (c *Context) getFlag(flag Z80Flag) bool {
	return *c.R1.F&byte(flag) != 0
}

func (c *Context) adjustFlags(val byte) {
	c.valFlag(F_5, (val&byte(F_5)) != 0)
	c.valFlag(F_3, (val&byte(F_3)) != 0)
}

func (c *Context) adjustFlagSZP(val byte) {
	c.valFlag(F_S, (val&0x80) != 0)
	c.valFlag(F_Z, val == 0)
	c.valFlag(F_PV, parityBit[val])
}

// Adjust flags after AND, OR, XOR
func (c *Context) adjustLogicFlag(flagH bool) {
	c.valFlag(F_S, (*c.R1.A&0x80) != 0)
	c.valFlag(F_Z, *c.R1.A == 0)
	c.valFlag(F_H, flagH)
	c.valFlag(F_N, false)
	c.valFlag(F_C, false)
	c.valFlag(F_PV, parityBit[*c.R1.A])
	c.adjustFlags(*c.R1.A)
}

func (c *Context) condition(cond Z80Condition) bool {
	switch cond {
	case C_Z:
		return c.getFlag(F_Z)
	case C_NZ:
		return !c.getFlag(F_Z)
	case C_C:
		return c.getFlag(F_C)
	case C_NC:
		return !c.getFlag(F_C)
	case C_M:
		return c.getFlag(F_S)
	case C_P:
		return !c.getFlag(F_S)
	case C_PE:
		return c.getFlag(F_PV)
	case C_PO:
		return !c.getFlag(F_PV)
	}
	return true
}
