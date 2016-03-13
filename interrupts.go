package z80

/*

	Go port of Gabriel Gambetta's brilliant libz80 library.
	Original C version can be found at https://github.com/ggambetta/libz80

*/

func (c *Context) doNmi() {
	c.unhalt()
	c.IFF2 = c.IFF1
	c.IFF1 = false
	c.doPush(c.PC)
	c.PC = 0x0066
	c.nmiRequested = false
	c.TStates += 5
}

func (c *Context) doInt() {
	c.unhalt()
	c.IFF1 = false
	c.IFF2 = false
	c.intRequested = false
	switch c.IM {
	case 0:
		c.execIntVector = true
		c.doExecute()
		c.execIntVector = false
	case 1:
		c.doPush(c.PC)
		c.PC = 0x0038
		c.TStates += 7
	case 2:
		c.doPush(c.PC)
		vectorAddr := (uint16(c.I) << 8) | uint16(c.intVector)
		c.PC = c.read16(vectorAddr)
		c.TStates += 7
	}
}

func (c *Context) Nmi() {
	c.nmiRequested = true
}

func (c *Context) Int(value byte) {
	c.intRequested = true
	c.intVector = value
}
