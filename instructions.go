package z80

func (c *Context) adc_A_A() {
	*c.R1.A = c.doArithmetic(*c.R1.A, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_B() {
	*c.R1.A = c.doArithmetic(*c.R1.B, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_C() {
	*c.R1.A = c.doArithmetic(*c.R1.C, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_D() {
	*c.R1.A = c.doArithmetic(*c.R1.D, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_E() {
	*c.R1.A = c.doArithmetic(*c.R1.E, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_H() {
	*c.R1.A = c.doArithmetic(*c.R1.H, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_L() {
	*c.R1.A = c.doArithmetic(*c.R1.L, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_IXh() {
	*c.R1.A = c.doArithmetic(*c.R1.IXh, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_IXl() {
	*c.R1.A = c.doArithmetic(*c.R1.IXl, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_IYh() {
	*c.R1.A = c.doArithmetic(*c.R1.IYh, F1_ADC, F2_ADC)
}

func (c *Context) adc_A_IYl() {
	*c.R1.A = c.doArithmetic(*c.R1.IYl, F1_ADC, F2_ADC)
}

func (c *Context) sbc_A_A() {
	*c.R1.A = c.doArithmetic(*c.R1.A, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_B() {
	*c.R1.A = c.doArithmetic(*c.R1.B, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_C() {
	*c.R1.A = c.doArithmetic(*c.R1.C, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_D() {
	*c.R1.A = c.doArithmetic(*c.R1.D, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_E() {
	*c.R1.A = c.doArithmetic(*c.R1.E, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_H() {
	*c.R1.A = c.doArithmetic(*c.R1.H, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_L() {
	*c.R1.A = c.doArithmetic(*c.R1.L, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_IXh() {
	*c.R1.A = c.doArithmetic(*c.R1.IXh, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_IXl() {
	*c.R1.A = c.doArithmetic(*c.R1.IXl, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_IYh() {
	*c.R1.A = c.doArithmetic(*c.R1.IYh, F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_IYl() {
	*c.R1.A = c.doArithmetic(*c.R1.IYl, F1_SBC, F2_SBC)
}

func (c *Context) add_A_A() {
	*c.R1.A = c.doArithmetic(*c.R1.A, F1_ADD, F2_ADD)
}

func (c *Context) add_A_B() {
	*c.R1.A = c.doArithmetic(*c.R1.B, F1_ADD, F2_ADD)
}

func (c *Context) add_A_C() {
	*c.R1.A = c.doArithmetic(*c.R1.C, F1_ADD, F2_ADD)
}

func (c *Context) add_A_D() {
	*c.R1.A = c.doArithmetic(*c.R1.D, F1_ADD, F2_ADD)
}

func (c *Context) add_A_E() {
	*c.R1.A = c.doArithmetic(*c.R1.E, F1_ADD, F2_ADD)
}

func (c *Context) add_A_H() {
	*c.R1.A = c.doArithmetic(*c.R1.H, F1_ADD, F2_ADD)
}

func (c *Context) add_A_L() {
	*c.R1.A = c.doArithmetic(*c.R1.L, F1_ADD, F2_ADD)
}

func (c *Context) add_A_IXh() {
	*c.R1.A = c.doArithmetic(*c.R1.IXh, F1_ADD, F2_ADD)
}

func (c *Context) add_A_IXl() {
	*c.R1.A = c.doArithmetic(*c.R1.IXl, F1_ADD, F2_ADD)
}

func (c *Context) add_A_IYh() {
	*c.R1.A = c.doArithmetic(*c.R1.IYh, F1_ADD, F2_ADD)
}

func (c *Context) add_A_IYl() {
	*c.R1.A = c.doArithmetic(*c.R1.IYl, F1_ADD, F2_ADD)
}

func (c *Context) sub_A_A() {
	*c.R1.A = c.doArithmetic(*c.R1.A, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_B() {
	*c.R1.A = c.doArithmetic(*c.R1.B, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_C() {
	*c.R1.A = c.doArithmetic(*c.R1.C, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_D() {
	*c.R1.A = c.doArithmetic(*c.R1.D, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_E() {
	*c.R1.A = c.doArithmetic(*c.R1.E, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_H() {
	*c.R1.A = c.doArithmetic(*c.R1.H, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_L() {
	*c.R1.A = c.doArithmetic(*c.R1.L, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_IXh() {
	*c.R1.A = c.doArithmetic(*c.R1.IXh, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_IXl() {
	*c.R1.A = c.doArithmetic(*c.R1.IXl, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_IYh() {
	*c.R1.A = c.doArithmetic(*c.R1.IYh, F1_SUB, F2_SUB)
}

func (c *Context) sub_A_IYl() {
	*c.R1.A = c.doArithmetic(*c.R1.IYl, F1_SUB, F2_SUB)
}

func (c *Context) adc_A_off_HL() {
	*c.R1.A = c.doArithmetic(c.read8(*c.R1.HL), F1_ADC, F2_ADC)
}

func (c *Context) sbc_A_off_HL() {
	*c.R1.A = c.doArithmetic(c.read8(*c.R1.HL), F1_SBC, F2_SBC)
}

func (c *Context) add_A_off_HL() {
	*c.R1.A = c.doArithmetic(c.read8(*c.R1.HL), F1_ADD, F2_ADD)
}

func (c *Context) sub_A_off_HL() {
	*c.R1.A = c.doArithmetic(c.read8(*c.R1.HL), F1_SUB, F2_SUB)
}

func (c *Context) adc_A_off_IX_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IX) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_ADC, F2_ADC)
}

func (c *Context) adc_A_off_IY_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IY) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_ADC, F2_ADC)
}

func (c *Context) sbc_A_off_IX_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IX) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_SBC, F2_SBC)
}

func (c *Context) sbc_A_off_IY_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IY) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_SBC, F2_SBC)
}

func (c *Context) add_A_off_IX_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IX) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_ADD, F2_ADD)
}

func (c *Context) add_A_off_IY_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IY) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_ADD, F2_ADD)
}

func (c *Context) sub_A_off_IX_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IX) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_SUB, F2_SUB)
}

func (c *Context) sub_A_off_IY_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int16(*c.R1.IY) + int16(displacement))
	*c.R1.A = c.doArithmetic(c.read8(addr), F1_SUB, F2_SUB)
}

func (c *Context) adc_A_n() {
	*c.R1.A = c.doArithmetic(c.read8(c.PC), F1_ADC, F2_ADC)
	c.PC++
}

func (c *Context) sbc_A_n() {
	*c.R1.A = c.doArithmetic(c.read8(c.PC), F1_SBC, F2_SBC)
	c.PC++
}

func (c *Context) add_A_n() {
	*c.R1.A = c.doArithmetic(c.read8(c.PC), F1_ADD, F2_ADD)
	c.PC++
}

func (c *Context) sub_A_n() {
	*c.R1.A = c.doArithmetic(c.read8(c.PC), F1_SUB, F2_SUB)
	c.PC++
}

func (c *Context) adc_HL_SP() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.SP, F1_ADC, F2_ADC)
}

func (c *Context) adc_HL_BC() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.BC, F1_ADC, F2_ADC)
}

func (c *Context) adc_HL_DE() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.DE, F1_ADC, F2_ADC)
}

func (c *Context) adc_HL_HL() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.HL, F1_ADC, F2_ADC)
}

func (c *Context) add_HL_SP() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.SP, F1_ADD, F2_ADD)
}

func (c *Context) add_HL_BC() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.BC, F1_ADD, F2_ADD)
}

func (c *Context) add_HL_DE() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.DE, F1_ADD, F2_ADD)
}

func (c *Context) add_HL_HL() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.HL, F1_ADD, F2_ADD)
}

func (c *Context) sbc_HL_SP() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.SP, F1_SBC, F2_SBC)
}

func (c *Context) sbc_HL_BC() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.BC, F1_SBC, F2_SBC)
}

func (c *Context) sbc_HL_DE() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.DE, F1_SBC, F2_SBC)
}

func (c *Context) sbc_HL_HL() {
	c.TStates += 7
	*c.R1.HL = c.doAddWord(*c.R1.HL, *c.R1.HL, F1_SBC, F2_SBC)
}

func (c *Context) add_IX_SP() {
	c.TStates += 7
	*c.R1.IX = c.doAddWord(*c.R1.IX, *c.R1.SP, false, false)
}

func (c *Context) add_IX_BC() {
	c.TStates += 7
	*c.R1.IX = c.doAddWord(*c.R1.IX, *c.R1.BC, false, false)
}

func (c *Context) add_IX_DE() {
	c.TStates += 7
	*c.R1.IX = c.doAddWord(*c.R1.IX, *c.R1.DE, false, false)
}

func (c *Context) add_IX_IX() {
	c.TStates += 7
	*c.R1.IX = c.doAddWord(*c.R1.IX, *c.R1.IX, false, false)
}

func (c *Context) add_IX_IY() {
	c.TStates += 7
	*c.R1.IX = c.doAddWord(*c.R1.IX, *c.R1.IY, false, false)
}

func (c *Context) add_IY_SP() {
	c.TStates += 7
	*c.R1.IY = c.doAddWord(*c.R1.IY, *c.R1.SP, false, false)
}

func (c *Context) add_IY_BC() {
	c.TStates += 7
	*c.R1.IY = c.doAddWord(*c.R1.IY, *c.R1.BC, false, false)
}

func (c *Context) add_IY_DE() {
	c.TStates += 7
	*c.R1.IY = c.doAddWord(*c.R1.IY, *c.R1.DE, false, false)
}

func (c *Context) add_IY_IX() {
	c.TStates += 7
	*c.R1.IY = c.doAddWord(*c.R1.IY, *c.R1.IX, false, false)
}

func (c *Context) add_IY_IY() {
	c.TStates += 7
	*c.R1.IY = c.doAddWord(*c.R1.IY, *c.R1.IY, false, false)
}

func (c *Context) and_off_HL() {
	c.doAND(c.read8(*c.R1.HL))
}

func (c *Context) xor_off_HL() {
	c.doXOR(c.read8(*c.R1.HL))
}

func (c *Context) or_off_HL() {
	c.doOR(c.read8(*c.R1.HL))
}

func (c *Context) and_off_IX_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int32(*c.R1.IX) + int32(displacement))
	c.doAND(c.read8(addr))
}

func (c *Context) and_off_IY_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int32(*c.R1.IY) + int32(displacement))
	c.doAND(c.read8(addr))
}

func (c *Context) xor_off_IX_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int32(*c.R1.IX) + int32(displacement))
	c.doXOR(c.read8(addr))
}

func (c *Context) xor_off_IY_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int32(*c.R1.IY) + int32(displacement))
	c.doXOR(c.read8(addr))
}

func (c *Context) or_off_IX_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int32(*c.R1.IX) + int32(displacement))
	c.doOR(c.read8(addr))
}

func (c *Context) or_off_IY_d() {
	c.TStates += 5
	displacement := int8(c.read8(c.PC))
	c.PC++
	addr := uint16(int32(*c.R1.IY) + int32(displacement))
	c.doOR(c.read8(addr))
}

func (c *Context) and_A() {
	c.doAND(*c.R1.A)
}

func (c *Context) and_B() {
	c.doAND(*c.R1.B)
}

func (c *Context) and_C() {
	c.doAND(*c.R1.C)
}

func (c *Context) and_D() {
	c.doAND(*c.R1.D)
}

func (c *Context) and_E() {
	c.doAND(*c.R1.E)
}

func (c *Context) and_H() {
	c.doAND(*c.R1.H)
}

func (c *Context) and_L() {
	c.doAND(*c.R1.L)
}

func (c *Context) and_IXh() {
	c.doAND(*c.R1.IXh)
}

func (c *Context) and_IXl() {
	c.doAND(*c.R1.IXl)
}

func (c *Context) and_IYh() {
	c.doAND(*c.R1.IYh)
}

func (c *Context) and_IYl() {
	c.doAND(*c.R1.IYl)
}

func (c *Context) xor_A() {
	c.doXOR(*c.R1.A)
}

func (c *Context) xor_B() {
	c.doXOR(*c.R1.B)
}

func (c *Context) xor_C() {
	c.doXOR(*c.R1.C)
}

func (c *Context) xor_D() {
	c.doXOR(*c.R1.D)
}

func (c *Context) xor_E() {
	c.doXOR(*c.R1.E)
}

func (c *Context) xor_H() {
	c.doXOR(*c.R1.H)
}

func (c *Context) xor_L() {
	c.doXOR(*c.R1.L)
}

func (c *Context) xor_IXh() {
	c.doXOR(*c.R1.IXh)
}

func (c *Context) xor_IXl() {
	c.doXOR(*c.R1.IXl)
}

func (c *Context) xor_IYh() {
	c.doXOR(*c.R1.IYh)
}

func (c *Context) xor_IYl() {
	c.doXOR(*c.R1.IYl)
}

func (c *Context) or_A() {
	c.doOR(*c.R1.A)
}

func (c *Context) or_B() {
	c.doOR(*c.R1.B)
}

func (c *Context) or_C() {
	c.doOR(*c.R1.C)
}

func (c *Context) or_D() {
	c.doOR(*c.R1.D)
}

func (c *Context) or_E() {
	c.doOR(*c.R1.E)
}

func (c *Context) or_H() {
	c.doOR(*c.R1.H)
}

func (c *Context) or_L() {
	c.doOR(*c.R1.L)
}

func (c *Context) or_IXh() {
	c.doOR(*c.R1.IXh)
}

func (c *Context) or_IXl() {
	c.doOR(*c.R1.IXl)
}

func (c *Context) or_IYh() {
	c.doOR(*c.R1.IYh)
}

func (c *Context) or_IYl() {
	c.doOR(*c.R1.IYl)
}

func (c *Context) and_n() {
	c.doAND(c.read8(c.PC))
	c.PC++
}

func (c *Context) xor_n() {
	c.doXOR(c.read8(c.PC))
	c.PC++
}

func (c *Context) or_n() {
	c.doOR(c.read8(c.PC))
	c.PC++
}

func (c *Context) bit_0_A() {
	c.doBIT_r(0, *c.R1.A)
}

func (c *Context) bit_0_B() {
	c.doBIT_r(0, *c.R1.B)
}

func (c *Context) bit_0_C() {
	c.doBIT_r(0, *c.R1.C)
}

func (c *Context) bit_0_D() {
	c.doBIT_r(0, *c.R1.D)
}

func (c *Context) bit_0_E() {
	c.doBIT_r(0, *c.R1.E)
}

func (c *Context) bit_0_H() {
	c.doBIT_r(0, *c.R1.H)
}

func (c *Context) bit_0_L() {
	c.doBIT_r(0, *c.R1.L)
}

func (c *Context) bit_1_A() {
	c.doBIT_r(1, *c.R1.A)
}

func (c *Context) bit_1_B() {
	c.doBIT_r(1, *c.R1.B)
}

func (c *Context) bit_1_C() {
	c.doBIT_r(1, *c.R1.C)
}

func (c *Context) bit_1_D() {
	c.doBIT_r(1, *c.R1.D)
}

func (c *Context) bit_1_E() {
	c.doBIT_r(1, *c.R1.E)
}

func (c *Context) bit_1_H() {
	c.doBIT_r(1, *c.R1.H)
}

func (c *Context) bit_1_L() {
	c.doBIT_r(1, *c.R1.L)
}

func (c *Context) bit_2_A() {
	c.doBIT_r(2, *c.R1.A)
}

func (c *Context) bit_2_B() {
	c.doBIT_r(2, *c.R1.B)
}

func (c *Context) bit_2_C() {
	c.doBIT_r(2, *c.R1.C)
}

func (c *Context) bit_2_D() {
	c.doBIT_r(2, *c.R1.D)
}

func (c *Context) bit_2_E() {
	c.doBIT_r(2, *c.R1.E)
}

func (c *Context) bit_2_H() {
	c.doBIT_r(2, *c.R1.H)
}

func (c *Context) bit_2_L() {
	c.doBIT_r(2, *c.R1.L)
}

func (c *Context) bit_3_A() {
	c.doBIT_r(3, *c.R1.A)
}

func (c *Context) bit_3_B() {
	c.doBIT_r(3, *c.R1.B)
}

func (c *Context) bit_3_C() {
	c.doBIT_r(3, *c.R1.C)
}

func (c *Context) bit_3_D() {
	c.doBIT_r(3, *c.R1.D)
}

func (c *Context) bit_3_E() {
	c.doBIT_r(3, *c.R1.E)
}

func (c *Context) bit_3_H() {
	c.doBIT_r(3, *c.R1.H)
}

func (c *Context) bit_3_L() {
	c.doBIT_r(3, *c.R1.L)
}

func (c *Context) bit_4_A() {
	c.doBIT_r(4, *c.R1.A)
}

func (c *Context) bit_4_B() {
	c.doBIT_r(4, *c.R1.B)
}

func (c *Context) bit_4_C() {
	c.doBIT_r(4, *c.R1.C)
}

func (c *Context) bit_4_D() {
	c.doBIT_r(4, *c.R1.D)
}

func (c *Context) bit_4_E() {
	c.doBIT_r(4, *c.R1.E)
}

func (c *Context) bit_4_H() {
	c.doBIT_r(4, *c.R1.H)
}

func (c *Context) bit_4_L() {
	c.doBIT_r(4, *c.R1.L)
}

func (c *Context) bit_5_A() {
	c.doBIT_r(5, *c.R1.A)
}

func (c *Context) bit_5_B() {
	c.doBIT_r(5, *c.R1.B)
}

func (c *Context) bit_5_C() {
	c.doBIT_r(5, *c.R1.C)
}

func (c *Context) bit_5_D() {
	c.doBIT_r(5, *c.R1.D)
}

func (c *Context) bit_5_E() {
	c.doBIT_r(5, *c.R1.E)
}

func (c *Context) bit_5_H() {
	c.doBIT_r(5, *c.R1.H)
}

func (c *Context) bit_5_L() {
	c.doBIT_r(5, *c.R1.L)
}

func (c *Context) bit_6_A() {
	c.doBIT_r(6, *c.R1.A)
}

func (c *Context) bit_6_B() {
	c.doBIT_r(6, *c.R1.B)
}

func (c *Context) bit_6_C() {
	c.doBIT_r(6, *c.R1.C)
}

func (c *Context) bit_6_D() {
	c.doBIT_r(6, *c.R1.D)
}

func (c *Context) bit_6_E() {
	c.doBIT_r(6, *c.R1.E)
}

func (c *Context) bit_6_H() {
	c.doBIT_r(6, *c.R1.H)
}

func (c *Context) bit_6_L() {
	c.doBIT_r(6, *c.R1.L)
}

func (c *Context) bit_7_A() {
	c.doBIT_r(7, *c.R1.A)
}

func (c *Context) bit_7_B() {
	c.doBIT_r(7, *c.R1.B)
}

func (c *Context) bit_7_C() {
	c.doBIT_r(7, *c.R1.C)
}

func (c *Context) bit_7_D() {
	c.doBIT_r(7, *c.R1.D)
}

func (c *Context) bit_7_E() {
	c.doBIT_r(7, *c.R1.E)
}

func (c *Context) bit_7_H() {
	c.doBIT_r(7, *c.R1.H)
}

func (c *Context) bit_7_L() {
	c.doBIT_r(7, *c.R1.L)
}

func (c *Context) bit_0_off_HL() {
	c.TStates += 1
	c.doBIT_r(0, c.read8(*c.R1.HL))
}

func (c *Context) bit_1_off_HL() {
	c.TStates += 1
	c.doBIT_r(1, c.read8(*c.R1.HL))
}

func (c *Context) bit_2_off_HL() {
	c.TStates += 1
	c.doBIT_r(2, c.read8(*c.R1.HL))
}

func (c *Context) bit_3_off_HL() {
	c.TStates += 1
	c.doBIT_r(3, c.read8(*c.R1.HL))
}

func (c *Context) bit_4_off_HL() {
	c.TStates += 1
	c.doBIT_r(4, c.read8(*c.R1.HL))
}

func (c *Context) bit_5_off_HL() {
	c.TStates += 1
	c.doBIT_r(5, c.read8(*c.R1.HL))
}

func (c *Context) bit_6_off_HL() {
	c.TStates += 1
	c.doBIT_r(6, c.read8(*c.R1.HL))
}

func (c *Context) bit_7_off_HL() {
	c.TStates += 1
	c.doBIT_r(7, c.read8(*c.R1.HL))
}

func (c *Context) bit_0_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(0, address)
}

func (c *Context) bit_0_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(0, address)
}

func (c *Context) bit_1_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(1, address)
}

func (c *Context) bit_1_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(1, address)
}

func (c *Context) bit_2_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(2, address)
}

func (c *Context) bit_2_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(2, address)
}

func (c *Context) bit_3_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(3, address)
}

func (c *Context) bit_3_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(3, address)
}

func (c *Context) bit_4_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(4, address)
}

func (c *Context) bit_4_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(4, address)
}

func (c *Context) bit_5_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(5, address)
}

func (c *Context) bit_5_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(5, address)
}

func (c *Context) bit_6_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(6, address)
}

func (c *Context) bit_6_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(6, address)
}

func (c *Context) bit_7_off_IX_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IX) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(7, address)
}

func (c *Context) bit_7_off_IY_d() {
	c.TStates += 2
	address := uint16(int32(*c.R1.IY) + int32(c.read8(c.PC)))
	c.PC++
	c.doBIT_indexed(7, address)
}

func (c *Context) set_0_A() {
	*c.R1.A = c.doSetRes(SR_SET, 0, *c.R1.A)
}

func (c *Context) set_0_B() {
	*c.R1.B = c.doSetRes(SR_SET, 0, *c.R1.B)
}

func (c *Context) set_0_C() {
	*c.R1.C = c.doSetRes(SR_SET, 0, *c.R1.C)
}

func (c *Context) set_0_D() {
	*c.R1.D = c.doSetRes(SR_SET, 0, *c.R1.D)
}

func (c *Context) set_0_E() {
	*c.R1.E = c.doSetRes(SR_SET, 0, *c.R1.E)
}

func (c *Context) set_0_H() {
	*c.R1.H = c.doSetRes(SR_SET, 0, *c.R1.H)
}

func (c *Context) set_0_L() {
	*c.R1.L = c.doSetRes(SR_SET, 0, *c.R1.L)
}

func (c *Context) set_1_A() {
	*c.R1.A = c.doSetRes(SR_SET, 1, *c.R1.A)
}

func (c *Context) set_1_B() {
	*c.R1.B = c.doSetRes(SR_SET, 1, *c.R1.B)
}

func (c *Context) set_1_C() {
	*c.R1.C = c.doSetRes(SR_SET, 1, *c.R1.C)
}

func (c *Context) set_1_D() {
	*c.R1.D = c.doSetRes(SR_SET, 1, *c.R1.D)
}

func (c *Context) set_1_E() {
	*c.R1.E = c.doSetRes(SR_SET, 1, *c.R1.E)
}

func (c *Context) set_1_H() {
	*c.R1.H = c.doSetRes(SR_SET, 1, *c.R1.H)
}

func (c *Context) set_1_L() {
	*c.R1.L = c.doSetRes(SR_SET, 1, *c.R1.L)
}

func (c *Context) set_2_A() {
	*c.R1.A = c.doSetRes(SR_SET, 2, *c.R1.A)
}

func (c *Context) set_2_B() {
	*c.R1.B = c.doSetRes(SR_SET, 2, *c.R1.B)
}

func (c *Context) set_2_C() {
	*c.R1.C = c.doSetRes(SR_SET, 2, *c.R1.C)
}

func (c *Context) set_2_D() {
	*c.R1.D = c.doSetRes(SR_SET, 2, *c.R1.D)
}

func (c *Context) set_2_E() {
	*c.R1.E = c.doSetRes(SR_SET, 2, *c.R1.E)
}

func (c *Context) set_2_H() {
	*c.R1.H = c.doSetRes(SR_SET, 2, *c.R1.H)
}

func (c *Context) set_2_L() {
	*c.R1.L = c.doSetRes(SR_SET, 2, *c.R1.L)
}

func (c *Context) set_3_A() {
	*c.R1.A = c.doSetRes(SR_SET, 3, *c.R1.A)
}

func (c *Context) set_3_B() {
	*c.R1.B = c.doSetRes(SR_SET, 3, *c.R1.B)
}

func (c *Context) set_3_C() {
	*c.R1.C = c.doSetRes(SR_SET, 3, *c.R1.C)
}

func (c *Context) set_3_D() {
	*c.R1.D = c.doSetRes(SR_SET, 3, *c.R1.D)
}

func (c *Context) set_3_E() {
	*c.R1.E = c.doSetRes(SR_SET, 3, *c.R1.E)
}

func (c *Context) set_3_H() {
	*c.R1.H = c.doSetRes(SR_SET, 3, *c.R1.H)
}

func (c *Context) set_3_L() {
	*c.R1.L = c.doSetRes(SR_SET, 3, *c.R1.L)
}

func (c *Context) set_4_A() {
	*c.R1.A = c.doSetRes(SR_SET, 4, *c.R1.A)
}

func (c *Context) set_4_B() {
	*c.R1.B = c.doSetRes(SR_SET, 4, *c.R1.B)
}

func (c *Context) set_4_C() {
	*c.R1.C = c.doSetRes(SR_SET, 4, *c.R1.C)
}

func (c *Context) set_4_D() {
	*c.R1.D = c.doSetRes(SR_SET, 4, *c.R1.D)
}

func (c *Context) set_4_E() {
	*c.R1.E = c.doSetRes(SR_SET, 4, *c.R1.E)
}

func (c *Context) set_4_H() {
	*c.R1.H = c.doSetRes(SR_SET, 4, *c.R1.H)
}

func (c *Context) set_4_L() {
	*c.R1.L = c.doSetRes(SR_SET, 4, *c.R1.L)
}

func (c *Context) set_5_A() {
	*c.R1.A = c.doSetRes(SR_SET, 5, *c.R1.A)
}

func (c *Context) set_5_B() {
	*c.R1.B = c.doSetRes(SR_SET, 5, *c.R1.B)
}

func (c *Context) set_5_C() {
	*c.R1.C = c.doSetRes(SR_SET, 5, *c.R1.C)
}

func (c *Context) set_5_D() {
	*c.R1.D = c.doSetRes(SR_SET, 5, *c.R1.D)
}

func (c *Context) set_5_E() {
	*c.R1.E = c.doSetRes(SR_SET, 5, *c.R1.E)
}

func (c *Context) set_5_H() {
	*c.R1.H = c.doSetRes(SR_SET, 5, *c.R1.H)
}

func (c *Context) set_5_L() {
	*c.R1.L = c.doSetRes(SR_SET, 5, *c.R1.L)
}

func (c *Context) set_6_A() {
	*c.R1.A = c.doSetRes(SR_SET, 6, *c.R1.A)
}

func (c *Context) set_6_B() {
	*c.R1.B = c.doSetRes(SR_SET, 6, *c.R1.B)
}

func (c *Context) set_6_C() {
	*c.R1.C = c.doSetRes(SR_SET, 6, *c.R1.C)
}

func (c *Context) set_6_D() {
	*c.R1.D = c.doSetRes(SR_SET, 6, *c.R1.D)
}

func (c *Context) set_6_E() {
	*c.R1.E = c.doSetRes(SR_SET, 6, *c.R1.E)
}

func (c *Context) set_6_H() {
	*c.R1.H = c.doSetRes(SR_SET, 6, *c.R1.H)
}

func (c *Context) set_6_L() {
	*c.R1.L = c.doSetRes(SR_SET, 6, *c.R1.L)
}

func (c *Context) set_7_A() {
	*c.R1.A = c.doSetRes(SR_SET, 7, *c.R1.A)
}

func (c *Context) set_7_B() {
	*c.R1.B = c.doSetRes(SR_SET, 7, *c.R1.B)
}

func (c *Context) set_7_C() {
	*c.R1.C = c.doSetRes(SR_SET, 7, *c.R1.C)
}

func (c *Context) set_7_D() {
	*c.R1.D = c.doSetRes(SR_SET, 7, *c.R1.D)
}

func (c *Context) set_7_E() {
	*c.R1.E = c.doSetRes(SR_SET, 7, *c.R1.E)
}

func (c *Context) set_7_H() {
	*c.R1.H = c.doSetRes(SR_SET, 7, *c.R1.H)
}

func (c *Context) set_7_L() {
	*c.R1.L = c.doSetRes(SR_SET, 7, *c.R1.L)
}

func (c *Context) res_0_A() {
	*c.R1.A = c.doSetRes(SR_RES, 0, *c.R1.A)
}

func (c *Context) res_0_B() {
	*c.R1.B = c.doSetRes(SR_RES, 0, *c.R1.B)
}

func (c *Context) res_0_C() {
	*c.R1.C = c.doSetRes(SR_RES, 0, *c.R1.C)
}

func (c *Context) res_0_D() {
	*c.R1.D = c.doSetRes(SR_RES, 0, *c.R1.D)
}

func (c *Context) res_0_E() {
	*c.R1.E = c.doSetRes(SR_RES, 0, *c.R1.E)
}

func (c *Context) res_0_H() {
	*c.R1.H = c.doSetRes(SR_RES, 0, *c.R1.H)
}

func (c *Context) res_0_L() {
	*c.R1.L = c.doSetRes(SR_RES, 0, *c.R1.L)
}

func (c *Context) res_1_A() {
	*c.R1.A = c.doSetRes(SR_RES, 1, *c.R1.A)
}

func (c *Context) res_1_B() {
	*c.R1.B = c.doSetRes(SR_RES, 1, *c.R1.B)
}

func (c *Context) res_1_C() {
	*c.R1.C = c.doSetRes(SR_RES, 1, *c.R1.C)
}

func (c *Context) res_1_D() {
	*c.R1.D = c.doSetRes(SR_RES, 1, *c.R1.D)
}

func (c *Context) res_1_E() {
	*c.R1.E = c.doSetRes(SR_RES, 1, *c.R1.E)
}

func (c *Context) res_1_H() {
	*c.R1.H = c.doSetRes(SR_RES, 1, *c.R1.H)
}

func (c *Context) res_1_L() {
	*c.R1.L = c.doSetRes(SR_RES, 1, *c.R1.L)
}

func (c *Context) res_2_A() {
	*c.R1.A = c.doSetRes(SR_RES, 2, *c.R1.A)
}

func (c *Context) res_2_B() {
	*c.R1.B = c.doSetRes(SR_RES, 2, *c.R1.B)
}

func (c *Context) res_2_C() {
	*c.R1.C = c.doSetRes(SR_RES, 2, *c.R1.C)
}

func (c *Context) res_2_D() {
	*c.R1.D = c.doSetRes(SR_RES, 2, *c.R1.D)
}

func (c *Context) res_2_E() {
	*c.R1.E = c.doSetRes(SR_RES, 2, *c.R1.E)
}

func (c *Context) res_2_H() {
	*c.R1.H = c.doSetRes(SR_RES, 2, *c.R1.H)
}

func (c *Context) res_2_L() {
	*c.R1.L = c.doSetRes(SR_RES, 2, *c.R1.L)
}

func (c *Context) res_3_A() {
	*c.R1.A = c.doSetRes(SR_RES, 3, *c.R1.A)
}

func (c *Context) res_3_B() {
	*c.R1.B = c.doSetRes(SR_RES, 3, *c.R1.B)
}

func (c *Context) res_3_C() {
	*c.R1.C = c.doSetRes(SR_RES, 3, *c.R1.C)
}

func (c *Context) res_3_D() {
	*c.R1.D = c.doSetRes(SR_RES, 3, *c.R1.D)
}

func (c *Context) res_3_E() {
	*c.R1.E = c.doSetRes(SR_RES, 3, *c.R1.E)
}

func (c *Context) res_3_H() {
	*c.R1.H = c.doSetRes(SR_RES, 3, *c.R1.H)
}

func (c *Context) res_3_L() {
	*c.R1.L = c.doSetRes(SR_RES, 3, *c.R1.L)
}

func (c *Context) res_4_A() {
	*c.R1.A = c.doSetRes(SR_RES, 4, *c.R1.A)
}

func (c *Context) res_4_B() {
	*c.R1.B = c.doSetRes(SR_RES, 4, *c.R1.B)
}

func (c *Context) res_4_C() {
	*c.R1.C = c.doSetRes(SR_RES, 4, *c.R1.C)
}

func (c *Context) res_4_D() {
	*c.R1.D = c.doSetRes(SR_RES, 4, *c.R1.D)
}

func (c *Context) res_4_E() {
	*c.R1.E = c.doSetRes(SR_RES, 4, *c.R1.E)
}

func (c *Context) res_4_H() {
	*c.R1.H = c.doSetRes(SR_RES, 4, *c.R1.H)
}

func (c *Context) res_4_L() {
	*c.R1.L = c.doSetRes(SR_RES, 4, *c.R1.L)
}

func (c *Context) res_5_A() {
	*c.R1.A = c.doSetRes(SR_RES, 5, *c.R1.A)
}

func (c *Context) res_5_B() {
	*c.R1.B = c.doSetRes(SR_RES, 5, *c.R1.B)
}

func (c *Context) res_5_C() {
	*c.R1.C = c.doSetRes(SR_RES, 5, *c.R1.C)
}

func (c *Context) res_5_D() {
	*c.R1.D = c.doSetRes(SR_RES, 5, *c.R1.D)
}

func (c *Context) res_5_E() {
	*c.R1.E = c.doSetRes(SR_RES, 5, *c.R1.E)
}

func (c *Context) res_5_H() {
	*c.R1.H = c.doSetRes(SR_RES, 5, *c.R1.H)
}

func (c *Context) res_5_L() {
	*c.R1.L = c.doSetRes(SR_RES, 5, *c.R1.L)
}

func (c *Context) res_6_A() {
	*c.R1.A = c.doSetRes(SR_RES, 6, *c.R1.A)
}

func (c *Context) res_6_B() {
	*c.R1.B = c.doSetRes(SR_RES, 6, *c.R1.B)
}

func (c *Context) res_6_C() {
	*c.R1.C = c.doSetRes(SR_RES, 6, *c.R1.C)
}

func (c *Context) res_6_D() {
	*c.R1.D = c.doSetRes(SR_RES, 6, *c.R1.D)
}

func (c *Context) res_6_E() {
	*c.R1.E = c.doSetRes(SR_RES, 6, *c.R1.E)
}

func (c *Context) res_6_H() {
	*c.R1.H = c.doSetRes(SR_RES, 6, *c.R1.H)
}

func (c *Context) res_6_L() {
	*c.R1.L = c.doSetRes(SR_RES, 6, *c.R1.L)
}

func (c *Context) res_7_A() {
	*c.R1.A = c.doSetRes(SR_RES, 7, *c.R1.A)
}

func (c *Context) res_7_B() {
	*c.R1.B = c.doSetRes(SR_RES, 7, *c.R1.B)
}

func (c *Context) res_7_C() {
	*c.R1.C = c.doSetRes(SR_RES, 7, *c.R1.C)
}

func (c *Context) res_7_D() {
	*c.R1.D = c.doSetRes(SR_RES, 7, *c.R1.D)
}

func (c *Context) res_7_E() {
	*c.R1.E = c.doSetRes(SR_RES, 7, *c.R1.E)
}

func (c *Context) res_7_H() {
	*c.R1.H = c.doSetRes(SR_RES, 7, *c.R1.H)
}

func (c *Context) res_7_L() {
	*c.R1.L = c.doSetRes(SR_RES, 7, *c.R1.L)
}

func (c *Context) set_0_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 0, c.read8(*c.R1.HL)))
}

func (c *Context) set_1_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 1, c.read8(*c.R1.HL)))
}

func (c *Context) set_2_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 2, c.read8(*c.R1.HL)))
}

func (c *Context) set_3_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 3, c.read8(*c.R1.HL)))
}

func (c *Context) set_4_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 4, c.read8(*c.R1.HL)))
}

func (c *Context) set_5_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 5, c.read8(*c.R1.HL)))
}

func (c *Context) set_6_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 6, c.read8(*c.R1.HL)))
}

func (c *Context) set_7_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_SET, 7, c.read8(*c.R1.HL)))
}

func (c *Context) res_0_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 0, c.read8(*c.R1.HL)))
}

func (c *Context) res_1_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 1, c.read8(*c.R1.HL)))
}

func (c *Context) res_2_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 2, c.read8(*c.R1.HL)))
}

func (c *Context) res_3_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 3, c.read8(*c.R1.HL)))
}

func (c *Context) res_4_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 4, c.read8(*c.R1.HL)))
}

func (c *Context) res_5_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 5, c.read8(*c.R1.HL)))
}

func (c *Context) res_6_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 6, c.read8(*c.R1.HL)))
}

func (c *Context) res_7_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSetRes(SR_RES, 7, c.read8(*c.R1.HL)))
}

func (c *Context) set_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 0, c.read8(address)))
}

func (c *Context) set_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 0, c.read8(address)))
}

func (c *Context) set_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 1, c.read8(address)))
}

func (c *Context) set_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 1, c.read8(address)))
}

func (c *Context) set_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 2, c.read8(address)))
}

func (c *Context) set_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 2, c.read8(address)))
}

func (c *Context) set_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 3, c.read8(address)))
}

func (c *Context) set_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 3, c.read8(address)))
}

func (c *Context) set_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 4, c.read8(address)))
}

func (c *Context) set_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 4, c.read8(address)))
}

func (c *Context) set_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 5, c.read8(address)))
}

func (c *Context) set_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 5, c.read8(address)))
}

func (c *Context) set_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 6, c.read8(address)))
}

func (c *Context) set_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 6, c.read8(address)))
}

func (c *Context) set_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_SET, 7, c.read8(address)))
}

func (c *Context) set_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_SET, 7, c.read8(address)))
}

func (c *Context) res_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 0, c.read8(address)))
}

func (c *Context) res_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 0, c.read8(address)))
}

func (c *Context) res_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 1, c.read8(address)))
}

func (c *Context) res_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 1, c.read8(address)))
}

func (c *Context) res_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 2, c.read8(address)))
}

func (c *Context) res_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 2, c.read8(address)))
}

func (c *Context) res_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 3, c.read8(address)))
}

func (c *Context) res_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 3, c.read8(address)))
}

func (c *Context) res_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 4, c.read8(address)))
}

func (c *Context) res_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 4, c.read8(address)))
}

func (c *Context) res_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 5, c.read8(address)))
}

func (c *Context) res_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 5, c.read8(address)))
}

func (c *Context) res_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 6, c.read8(address)))
}

func (c *Context) res_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 6, c.read8(address)))
}

func (c *Context) res_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IX, off)
	c.write8(address, c.doSetRes(SR_RES, 7, c.read8(address)))
}

func (c *Context) res_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	address := doOff(*c.R1.IY, off)
	c.write8(address, c.doSetRes(SR_RES, 7, c.read8(address)))
}

func (c *Context) call_C_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_C) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_M_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_M) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_NZ_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_NZ) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_NC_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_NC) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_P_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_P) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_PE_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_PE) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_PO_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_PO) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_Z_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	if c.condition(C_Z) {
		c.TStates += 1
		c.doPush(c.PC)
		c.PC = addr
	}
}

func (c *Context) call_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = addr
}

func (c *Context) jp_off_HL() {
	c.PC = *c.R1.HL
}

func (c *Context) jp_off_IX() {
	c.PC = *c.R1.IX
}

func (c *Context) jp_off_IY() {
	c.PC = *c.R1.IY
}

func (c *Context) jp_C_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_C) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_M_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_M) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_NZ_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_NZ) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_NC_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_NC) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_P_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_P) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_PE_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_PE) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_PO_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_PO) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_Z_off_nn() {
	addr := c.read16(c.PC)
	if c.condition(C_Z) {
		c.PC = addr
	} else {
		c.PC += 2
	}
}

func (c *Context) jp_off_nn() {
	addr := c.read16(c.PC)
	c.PC = addr
}

func (c *Context) jr_C_off_PC_e() {
	off := doComplement(c.read8(c.PC))
	c.PC++
	if c.condition(C_C) {
		c.TStates += 5
		c.PC = doOff(c.PC, off)
	}
}

func (c *Context) jr_NZ_off_PC_e() {
	off := doComplement(c.read8(c.PC))
	c.PC++
	if c.condition(C_NZ) {
		c.TStates += 5
		c.PC = doOff(c.PC, off)
	}
}

func (c *Context) jr_NC_off_PC_e() {
	off := doComplement(c.read8(c.PC))
	c.PC++
	if c.condition(C_NC) {
		c.TStates += 5
		c.PC = doOff(c.PC, off)
	}
}

func (c *Context) jr_Z_off_PC_e() {
	off := doComplement(c.read8(c.PC))
	c.PC++
	if c.condition(C_Z) {
		c.TStates += 5
		c.PC = doOff(c.PC, off)
	}
}

func (c *Context) jr_off_PC_e() {
	off := doComplement(c.read8(c.PC))
	c.PC++
	c.TStates += 5
	c.PC = doOff(c.PC, off)
}

func (c *Context) reti() {
	c.IFF1 = c.IFF2
	c.ret()
}

func (c *Context) retn() {
	c.IFF1 = c.IFF2
	c.ret()
}

func (c *Context) ret_C() {
	c.TStates += 1
	if c.condition(C_C) {
		c.ret()
	}
}

func (c *Context) ret_M() {
	c.TStates += 1
	if c.condition(C_M) {
		c.ret()
	}
}

func (c *Context) ret_NZ() {
	c.TStates += 1
	if c.condition(C_NZ) {
		c.ret()
	}
}

func (c *Context) ret_NC() {
	c.TStates += 1
	if c.condition(C_NC) {
		c.ret()
	}
}

func (c *Context) ret_P() {
	c.TStates += 1
	if c.condition(C_P) {
		c.ret()
	}
}

func (c *Context) ret_PE() {
	c.TStates += 1
	if c.condition(C_PE) {
		c.ret()
	}
}

func (c *Context) ret_PO() {
	c.TStates += 1
	if c.condition(C_PO) {
		c.ret()
	}
}

func (c *Context) ret_Z() {
	c.TStates += 1
	if c.condition(C_Z) {
		c.ret()
	}
}

func (c *Context) ret() {
	c.PC = c.doPop()
}

func (c *Context) djnz_off_PC_e() {
	c.TStates += 1
	off := doComplement(c.read8(c.PC))
	c.PC++
	*c.R1.B--
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC = doOff(c.PC, off)
	}
}

func (c *Context) rst_0H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x00
}

func (c *Context) rst_8H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x08
}

func (c *Context) rst_10H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x010
}

func (c *Context) rst_18H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x018
}

func (c *Context) rst_20H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x020
}

func (c *Context) rst_28H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x028
}

func (c *Context) rst_30H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x030
}

func (c *Context) rst_38H() {
	c.TStates += 1
	c.doPush(c.PC)
	c.PC = 0x038
}

func (c *Context) ccf() {
	c.valFlag(F_C, !c.getFlag(F_C))
	c.resFlag(F_N)
	c.adjustFlags(*c.R1.A)
}

func (c *Context) scf() {
	c.setFlag(F_C)
	c.resFlag(F_N | F_H)
	c.adjustFlags(*c.R1.A)
}

func (c *Context) cpl() {
	*c.R1.A = ^*c.R1.A
	c.setFlag(F_N | F_H)
	c.adjustFlags(*c.R1.A)
}

func (c *Context) daa() {
	c.doDAA()
}

func (c *Context) ex_off_SP_HL() {
	c.TStates += 3
	tmp := c.read16(*c.R1.SP)
	c.write16(*c.R1.SP, *c.R1.HL)
	*c.R1.HL = tmp
}

func (c *Context) ex_off_SP_IX() {
	c.TStates += 3
	tmp := c.read16(*c.R1.SP)
	c.write16(*c.R1.SP, *c.R1.IX)
	*c.R1.IX = tmp
}

func (c *Context) ex_off_SP_IY() {
	c.TStates += 3
	tmp := c.read16(*c.R1.SP)
	c.write16(*c.R1.SP, *c.R1.IY)
	*c.R1.IY = tmp
}

func (c *Context) ex_AF_AF_() {
	tmp := *c.R1.AF
	*c.R1.AF = *c.R2.AF
	*c.R2.AF = tmp
}

func (c *Context) ex_DE_HL() {
	tmp := *c.R1.DE
	*c.R1.DE = *c.R1.HL
	*c.R1.HL = tmp
}

func (c *Context) exx() {
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
}

func (c *Context) halt() {
	c.halted = true
	c.PC--
}

func (c *Context) cp_off_HL() {
	c.doCP_HL()
}

func (c *Context) cp_off_IX_d() {
	c.TStates += 5
	displacement := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, displacement)
	val := c.read8(addr)
	c.doArithmetic(val, false, true)
	c.adjustFlags(val)
}

func (c *Context) cp_off_IY_d() {
	c.TStates += 5
	displacement := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, displacement)
	val := c.read8(addr)
	c.doArithmetic(val, false, true)
	c.adjustFlags(val)
}

func (c *Context) cp_A() {
	c.doArithmetic(*c.R1.A, false, true)
	c.adjustFlags(*c.R1.A)
}

func (c *Context) cp_B() {
	c.doArithmetic(*c.R1.B, false, true)
	c.adjustFlags(*c.R1.B)
}

func (c *Context) cp_C() {
	c.doArithmetic(*c.R1.C, false, true)
	c.adjustFlags(*c.R1.C)
}

func (c *Context) cp_D() {
	c.doArithmetic(*c.R1.D, false, true)
	c.adjustFlags(*c.R1.D)
}

func (c *Context) cp_E() {
	c.doArithmetic(*c.R1.E, false, true)
	c.adjustFlags(*c.R1.E)
}

func (c *Context) cp_H() {
	c.doArithmetic(*c.R1.H, false, true)
	c.adjustFlags(*c.R1.H)
}

func (c *Context) cp_L() {
	c.doArithmetic(*c.R1.L, false, true)
	c.adjustFlags(*c.R1.L)
}

func (c *Context) cp_IXh() {
	c.doArithmetic(*c.R1.IXh, false, true)
	c.adjustFlags(*c.R1.IXh)
}

func (c *Context) cp_IXl() {
	c.doArithmetic(*c.R1.IXl, false, true)
	c.adjustFlags(*c.R1.IXl)
}

func (c *Context) cp_IYh() {
	c.doArithmetic(*c.R1.IYh, false, true)
	c.adjustFlags(*c.R1.IYh)
}

func (c *Context) cp_IYl() {
	c.doArithmetic(*c.R1.IYl, false, true)
	c.adjustFlags(*c.R1.IYl)
}

func (c *Context) cp_n() {
	val := c.read8(c.PC)
	c.PC++
	c.doArithmetic(val, false, true)
	c.adjustFlags(val)
}

func (c *Context) cpdr() {
	c.cpd()
	if *c.R1.BC != 0 && !c.getFlag(F_Z) {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) cpd() {
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
	c.valFlag(F_5, value&(1<<1) != 0)
	c.valFlag(F_3, value&(1<<3) != 0)
}

func (c *Context) cpir() {
	c.cpi()
	if *c.R1.BC != 0 && !c.getFlag(F_Z) {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) cpi() {
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
	c.valFlag(F_5, value&(1<<2) != 0)
	c.valFlag(F_3, value&(1<<3) != 0)
}

func (c *Context) inc_off_HL() {
	c.TStates += 1
	value := c.read8(*c.R1.HL)
	c.write8(*c.R1.HL, c.doIncDec(value, ID_INC))
}

func (c *Context) dec_off_HL() {
	c.TStates += 1
	value := c.read8(*c.R1.HL)
	c.write8(*c.R1.HL, c.doIncDec(value, ID_DEC))
}

func (c *Context) inc_off_IX_d() {
	c.TStates += 6
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	value := c.read8(addr)
	c.write8(addr, c.doIncDec(value, ID_INC))
}

func (c *Context) inc_off_IY_d() {
	c.TStates += 6
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	value := c.read8(addr)
	c.write8(addr, c.doIncDec(value, ID_INC))
}

func (c *Context) dec_off_IX_d() {
	c.TStates += 6
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	value := c.read8(addr)
	c.write8(addr, c.doIncDec(value, ID_DEC))
}

func (c *Context) dec_off_IY_d() {
	c.TStates += 6
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	value := c.read8(addr)
	c.write8(addr, c.doIncDec(value, ID_DEC))
}

func (c *Context) inc_A() {
	*c.R1.A = c.doIncDec(*c.R1.A, ID_INC)
}

func (c *Context) inc_B() {
	*c.R1.B = c.doIncDec(*c.R1.B, ID_INC)
}

func (c *Context) inc_C() {
	*c.R1.C = c.doIncDec(*c.R1.C, ID_INC)
}

func (c *Context) inc_D() {
	*c.R1.D = c.doIncDec(*c.R1.D, ID_INC)
}

func (c *Context) inc_E() {
	*c.R1.E = c.doIncDec(*c.R1.E, ID_INC)
}

func (c *Context) inc_H() {
	*c.R1.H = c.doIncDec(*c.R1.H, ID_INC)
}

func (c *Context) inc_L() {
	*c.R1.L = c.doIncDec(*c.R1.L, ID_INC)
}

func (c *Context) inc_IXh() {
	*c.R1.IXh = c.doIncDec(*c.R1.IXh, ID_INC)
}

func (c *Context) inc_IXl() {
	*c.R1.IXl = c.doIncDec(*c.R1.IXl, ID_INC)
}

func (c *Context) inc_IYh() {
	*c.R1.IYh = c.doIncDec(*c.R1.IYh, ID_INC)
}

func (c *Context) inc_IYl() {
	*c.R1.IYl = c.doIncDec(*c.R1.IYl, ID_INC)
}

func (c *Context) dec_A() {
	*c.R1.A = c.doIncDec(*c.R1.A, ID_DEC)
}

func (c *Context) dec_B() {
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)
}

func (c *Context) dec_C() {
	*c.R1.C = c.doIncDec(*c.R1.C, ID_DEC)
}

func (c *Context) dec_D() {
	*c.R1.D = c.doIncDec(*c.R1.D, ID_DEC)
}

func (c *Context) dec_E() {
	*c.R1.E = c.doIncDec(*c.R1.E, ID_DEC)
}

func (c *Context) dec_H() {
	*c.R1.H = c.doIncDec(*c.R1.H, ID_DEC)
}

func (c *Context) dec_L() {
	*c.R1.L = c.doIncDec(*c.R1.L, ID_DEC)
}

func (c *Context) dec_IXh() {
	*c.R1.IXh = c.doIncDec(*c.R1.IXh, ID_DEC)
}

func (c *Context) dec_IXl() {
	*c.R1.IXl = c.doIncDec(*c.R1.IXl, ID_DEC)
}

func (c *Context) dec_IYh() {
	*c.R1.IYh = c.doIncDec(*c.R1.IYh, ID_DEC)
}

func (c *Context) dec_IYl() {
	*c.R1.IYl = c.doIncDec(*c.R1.IYl, ID_DEC)
}

func (c *Context) inc_BC() {
	c.TStates += 2
	*c.R1.BC++
}

func (c *Context) inc_DE() {
	c.TStates += 2
	*c.R1.DE++
}

func (c *Context) inc_HL() {
	c.TStates += 2
	*c.R1.HL++
}

func (c *Context) inc_SP() {
	c.TStates += 2
	*c.R1.SP++
}

func (c *Context) inc_IX() {
	c.TStates += 2
	*c.R1.IX++
}

func (c *Context) inc_IY() {
	c.TStates += 2
	*c.R1.IY++
}

func (c *Context) dec_BC() {
	c.TStates += 2
	*c.R1.BC--
}

func (c *Context) dec_DE() {
	c.TStates += 2
	*c.R1.DE--
}

func (c *Context) dec_HL() {
	c.TStates += 2
	*c.R1.HL--
}

func (c *Context) dec_SP() {
	c.TStates += 2
	*c.R1.SP--
}

func (c *Context) dec_IX() {
	c.TStates += 2
	*c.R1.IX--
}

func (c *Context) dec_IY() {
	c.TStates += 2
	*c.R1.IY--
}

func (c *Context) ei() {
	c.IFF1 = IE_EI
	c.IFF2 = IE_EI
	c.deferInt = true
}

func (c *Context) di() {
	c.IFF1 = IE_DI
	c.IFF2 = IE_DI
	c.deferInt = true
}

func (c *Context) im_0() {
	c.IM = 0
}

func (c *Context) im_1() {
	c.IM = 1
}

func (c *Context) im_2() {
	c.IM = 2
}

func (c *Context) in_A_off_C() {
	*c.R1.A = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.A)
	c.adjustFlags(*c.R1.A)
}

func (c *Context) in_B_off_C() {
	*c.R1.B = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.B)
	c.adjustFlags(*c.R1.B)
}

func (c *Context) in_C_off_C() {
	*c.R1.C = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.C)
	c.adjustFlags(*c.R1.C)
}

func (c *Context) in_D_off_C() {
	*c.R1.D = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.D)
	c.adjustFlags(*c.R1.D)
}

func (c *Context) in_E_off_C() {
	*c.R1.E = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.E)
	c.adjustFlags(*c.R1.E)
}

func (c *Context) in_F_off_C() {
	*c.R1.F = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.F)
	c.adjustFlags(*c.R1.F)
}

func (c *Context) in_H_off_C() {
	*c.R1.H = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.H)
	c.adjustFlags(*c.R1.H)
}

func (c *Context) in_L_off_C() {
	*c.R1.L = c.ioread(*c.R1.BC)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.L)
	c.adjustFlags(*c.R1.L)
}

func (c *Context) in_A_off_n() {
	port := uint16(c.read8(c.PC))
	c.PC++
	port = uint16(*c.R1.A)<<8 | port
	*c.R1.A = c.ioread(port)
}

func (c *Context) indr() {
	c.ind()
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) ind() {
	c.TStates += 1
	val := c.ioread(*c.R1.BC)
	c.write8(*c.R1.HL, val)
	*c.R1.HL--
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)
	c.valFlag(F_N, val&0x80 != 0)
	flagval := int(val) + int((*c.R1.C-1)&0xFF)
	c.valFlag(F_H, flagval > 0xFF)
	c.valFlag(F_C, flagval > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flagval&7)^*c.R1.B])
}

func (c *Context) inir() {
	c.ini()
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) ini() {
	c.TStates += 1
	val := c.ioread(*c.R1.BC)
	c.write8(*c.R1.HL, val)
	*c.R1.HL++
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)
	c.valFlag(F_N, val&0x80 != 0)
	flagval := int(val) + int((*c.R1.C+1)&0xFF)
	c.valFlag(F_H, flagval > 0xFF)
	c.valFlag(F_C, flagval > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flagval&7)^*c.R1.B])
}

func (c *Context) outi() {
	c.TStates += 1
	value := c.read8(*c.R1.HL)
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)
	c.iowrite(*c.R1.BC, value)
	*c.R1.HL++
	flag_value := int(value) + int(*c.R1.L)
	c.valFlag(F_N, value&0x80 != 0)
	c.valFlag(F_H, flag_value > 0xFF)
	c.valFlag(F_C, flag_value > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flag_value&7)^*c.R1.B])
	c.adjustFlags(*c.R1.B)
}

func (c *Context) otir() {
	c.outi()
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) outd() {
	c.TStates += 1
	value := c.read8(*c.R1.HL)
	*c.R1.B = c.doIncDec(*c.R1.B, ID_DEC)
	c.iowrite(*c.R1.BC, value)
	*c.R1.HL--
	flag_value := int(value) + int(*c.R1.L)
	c.valFlag(F_N, value&0x80 != 0)
	c.valFlag(F_H, flag_value > 0xFF)
	c.valFlag(F_C, flag_value > 0xFF)
	c.valFlag(F_PV, parityBit[byte(flag_value&7)^*c.R1.B])
	c.adjustFlags(*c.R1.B)
}

func (c *Context) otdr() {
	c.outd()
	if *c.R1.B != 0 {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) out_off_C_0() {
	c.iowrite(*c.R1.BC, 0)
}

func (c *Context) out_off_C_A() {
	c.iowrite(*c.R1.BC, *c.R1.A)
}

func (c *Context) out_off_C_B() {
	c.iowrite(*c.R1.BC, *c.R1.B)
}

func (c *Context) out_off_C_C() {
	c.iowrite(*c.R1.BC, *c.R1.C)
}

func (c *Context) out_off_C_D() {
	c.iowrite(*c.R1.BC, *c.R1.D)
}

func (c *Context) out_off_C_E() {
	c.iowrite(*c.R1.BC, *c.R1.E)
}

func (c *Context) out_off_C_H() {
	c.iowrite(*c.R1.BC, *c.R1.H)
}

func (c *Context) out_off_C_L() {
	c.iowrite(*c.R1.BC, *c.R1.L)
}

func (c *Context) out_off_n_A() {
	port := uint16(c.read8(c.PC))
	c.PC++
	port = uint16(*c.R1.A)<<8 | port
	c.iowrite(port, *c.R1.A)
}

func (c *Context) pop_AF() {
	*c.R1.AF = c.doPop()
}

func (c *Context) pop_BC() {
	*c.R1.BC = c.doPop()
}

func (c *Context) pop_DE() {
	*c.R1.DE = c.doPop()
}

func (c *Context) pop_HL() {
	*c.R1.HL = c.doPop()
}

func (c *Context) pop_IX() {
	*c.R1.IX = c.doPop()
}

func (c *Context) pop_IY() {
	*c.R1.IY = c.doPop()
}

func (c *Context) push_AF() {
	c.TStates += 1
	c.doPush(*c.R1.AF)
}

func (c *Context) push_BC() {
	c.TStates += 1
	c.doPush(*c.R1.BC)
}

func (c *Context) push_DE() {
	c.TStates += 1
	c.doPush(*c.R1.DE)
}

func (c *Context) push_HL() {
	c.TStates += 1
	c.doPush(*c.R1.HL)
}

func (c *Context) push_IX() {
	c.TStates += 1
	c.doPush(*c.R1.IX)
}

func (c *Context) push_IY() {
	c.TStates += 1
	c.doPush(*c.R1.IY)
}

func (c *Context) rlc_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doRLC(true, c.read8(*c.R1.HL)))
}

func (c *Context) rrc_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doRRC(true, c.read8(*c.R1.HL)))
}

func (c *Context) rl_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doRL(true, c.read8(*c.R1.HL)))
}

func (c *Context) rr_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doRR(true, c.read8(*c.R1.HL)))
}

func (c *Context) rlc_A() {
	*c.R1.A = c.doRLC(true, *c.R1.A)
}

func (c *Context) rlc_B() {
	*c.R1.B = c.doRLC(true, *c.R1.B)
}

func (c *Context) rlc_C() {
	*c.R1.C = c.doRLC(true, *c.R1.C)
}

func (c *Context) rlc_D() {
	*c.R1.D = c.doRLC(true, *c.R1.D)
}

func (c *Context) rlc_E() {
	*c.R1.E = c.doRLC(true, *c.R1.E)
}

func (c *Context) rlc_H() {
	*c.R1.H = c.doRLC(true, *c.R1.H)
}

func (c *Context) rlc_L() {
	*c.R1.L = c.doRLC(true, *c.R1.L)
}

func (c *Context) rlc_IXh() {
	*c.R1.IXh = c.doRLC(true, *c.R1.IXh)
}

func (c *Context) rlc_IXl() {
	*c.R1.IXl = c.doRLC(true, *c.R1.IXl)
}

func (c *Context) rlc_IYh() {
	*c.R1.IYh = c.doRLC(true, *c.R1.IYh)
}

func (c *Context) rlc_IYl() {
	*c.R1.IYl = c.doRLC(true, *c.R1.IYl)
}

func (c *Context) rrc_A() {
	*c.R1.A = c.doRRC(true, *c.R1.A)
}

func (c *Context) rrc_B() {
	*c.R1.B = c.doRRC(true, *c.R1.B)
}

func (c *Context) rrc_C() {
	*c.R1.C = c.doRRC(true, *c.R1.C)
}

func (c *Context) rrc_D() {
	*c.R1.D = c.doRRC(true, *c.R1.D)
}

func (c *Context) rrc_E() {
	*c.R1.E = c.doRRC(true, *c.R1.E)
}

func (c *Context) rrc_H() {
	*c.R1.H = c.doRRC(true, *c.R1.H)
}

func (c *Context) rrc_L() {
	*c.R1.L = c.doRRC(true, *c.R1.L)
}

func (c *Context) rrc_IXh() {
	*c.R1.IXh = c.doRRC(true, *c.R1.IXh)
}

func (c *Context) rrc_IXl() {
	*c.R1.IXl = c.doRRC(true, *c.R1.IXl)
}

func (c *Context) rrc_IYh() {
	*c.R1.IYh = c.doRRC(true, *c.R1.IYh)
}

func (c *Context) rrc_IYl() {
	*c.R1.IYl = c.doRRC(true, *c.R1.IYl)
}

func (c *Context) rl_A() {
	*c.R1.A = c.doRL(true, *c.R1.A)
}

func (c *Context) rl_B() {
	*c.R1.B = c.doRL(true, *c.R1.B)
}

func (c *Context) rl_C() {
	*c.R1.C = c.doRL(true, *c.R1.C)
}

func (c *Context) rl_D() {
	*c.R1.D = c.doRL(true, *c.R1.D)
}

func (c *Context) rl_E() {
	*c.R1.E = c.doRL(true, *c.R1.E)
}

func (c *Context) rl_H() {
	*c.R1.H = c.doRL(true, *c.R1.H)
}

func (c *Context) rl_L() {
	*c.R1.L = c.doRL(true, *c.R1.L)
}

func (c *Context) rl_IXh() {
	*c.R1.IXh = c.doRL(true, *c.R1.IXh)
}

func (c *Context) rl_IXl() {
	*c.R1.IXl = c.doRL(true, *c.R1.IXl)
}

func (c *Context) rl_IYh() {
	*c.R1.IYh = c.doRL(true, *c.R1.IYh)
}

func (c *Context) rl_IYl() {
	*c.R1.IYl = c.doRL(true, *c.R1.IYl)
}

func (c *Context) rr_A() {
	*c.R1.A = c.doRR(true, *c.R1.A)
}

func (c *Context) rr_B() {
	*c.R1.B = c.doRR(true, *c.R1.B)
}

func (c *Context) rr_C() {
	*c.R1.C = c.doRR(true, *c.R1.C)
}

func (c *Context) rr_D() {
	*c.R1.D = c.doRR(true, *c.R1.D)
}

func (c *Context) rr_E() {
	*c.R1.E = c.doRR(true, *c.R1.E)
}

func (c *Context) rr_H() {
	*c.R1.H = c.doRR(true, *c.R1.H)
}

func (c *Context) rr_L() {
	*c.R1.L = c.doRR(true, *c.R1.L)
}

func (c *Context) rr_IXh() {
	*c.R1.IXh = c.doRR(true, *c.R1.IXh)
}

func (c *Context) rr_IXl() {
	*c.R1.IXl = c.doRR(true, *c.R1.IXl)
}

func (c *Context) rr_IYh() {
	*c.R1.IYh = c.doRR(true, *c.R1.IYh)
}

func (c *Context) rr_IYl() {
	*c.R1.IYl = c.doRR(true, *c.R1.IYl)
}

func (c *Context) rlc_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doRLC(true, c.read8(addr)))
}

func (c *Context) rlc_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doRLC(true, c.read8(addr)))
}

func (c *Context) rrc_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doRRC(true, c.read8(addr)))
}

func (c *Context) rrc_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doRRC(true, c.read8(addr)))
}

func (c *Context) rl_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doRL(true, c.read8(addr)))
}

func (c *Context) rl_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doRL(true, c.read8(addr)))
}

func (c *Context) rr_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doRR(true, c.read8(addr)))
}

func (c *Context) rr_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doRR(true, c.read8(addr)))
}

func (c *Context) rla() {
	*c.R1.A = c.doRL(false, *c.R1.A)
}

func (c *Context) rra() {
	*c.R1.A = c.doRR(false, *c.R1.A)
}

func (c *Context) rlca() {
	*c.R1.A = c.doRLC(false, *c.R1.A)
}

func (c *Context) rrca() {
	*c.R1.A = c.doRRC(false, *c.R1.A)
}

func (c *Context) rld() {
	c.TStates += 4
	ah := *c.R1.A & 0x0F
	hl := c.read8(*c.R1.HL)
	*c.R1.A = (*c.R1.A & 0xF0) | ((hl & 0xF0) >> 4)
	hl = (hl << 4) | ah
	c.write8(*c.R1.HL, hl)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.A)
	c.adjustFlags(*c.R1.A)
}

func (c *Context) rrd() {
	c.TStates += 4
	ah := *c.R1.A & 0x0F
	hl := c.read8(*c.R1.HL)
	*c.R1.A = (*c.R1.A & 0xF0) | (hl & 0x0F)
	hl = (hl >> 4) | (ah << 4)
	c.write8(*c.R1.HL, hl)
	c.resFlag(F_H | F_N)
	c.adjustFlagSZP(*c.R1.A)
}

func (c *Context) sll_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSL(c.read8(*c.R1.HL), IA_L))
}

func (c *Context) sla_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSL(c.read8(*c.R1.HL), IA_A))
}

func (c *Context) srl_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSR(c.read8(*c.R1.HL), IA_L))
}

func (c *Context) sra_off_HL() {
	c.TStates += 1
	c.write8(*c.R1.HL, c.doSR(c.read8(*c.R1.HL), IA_A))
}

func (c *Context) sll_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doSL(c.read8(addr), IA_L))
}

func (c *Context) sll_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doSL(c.read8(addr), IA_L))
}

func (c *Context) sla_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doSL(c.read8(addr), IA_A))
}

func (c *Context) sla_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doSL(c.read8(addr), IA_A))
}

func (c *Context) srl_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doSR(c.read8(addr), IA_L))
}

func (c *Context) srl_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doSR(c.read8(addr), IA_L))
}

func (c *Context) sra_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, c.doSR(c.read8(addr), IA_A))
}

func (c *Context) sra_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, c.doSR(c.read8(addr), IA_A))
}

func (c *Context) sll_A() {
	*c.R1.A = c.doSL(*c.R1.A, IA_L)
}

func (c *Context) sll_B() {
	*c.R1.B = c.doSL(*c.R1.B, IA_L)
}

func (c *Context) sll_C() {
	*c.R1.C = c.doSL(*c.R1.C, IA_L)
}

func (c *Context) sll_D() {
	*c.R1.D = c.doSL(*c.R1.D, IA_L)
}

func (c *Context) sll_E() {
	*c.R1.E = c.doSL(*c.R1.E, IA_L)
}

func (c *Context) sll_H() {
	*c.R1.H = c.doSL(*c.R1.H, IA_L)
}

func (c *Context) sll_L() {
	*c.R1.L = c.doSL(*c.R1.L, IA_L)
}

func (c *Context) sll_IXh() {
	*c.R1.IXh = c.doSL(*c.R1.IXh, IA_L)
}

func (c *Context) sll_IXl() {
	*c.R1.IXl = c.doSL(*c.R1.IXl, IA_L)
}

func (c *Context) sll_IYh() {
	*c.R1.IYh = c.doSL(*c.R1.IYh, IA_L)
}

func (c *Context) sll_IYl() {
	*c.R1.IYl = c.doSL(*c.R1.IYl, IA_L)
}

func (c *Context) sla_A() {
	*c.R1.A = c.doSL(*c.R1.A, IA_A)
}

func (c *Context) sla_B() {
	*c.R1.B = c.doSL(*c.R1.B, IA_A)
}

func (c *Context) sla_C() {
	*c.R1.C = c.doSL(*c.R1.C, IA_A)
}

func (c *Context) sla_D() {
	*c.R1.D = c.doSL(*c.R1.D, IA_A)
}

func (c *Context) sla_E() {
	*c.R1.E = c.doSL(*c.R1.E, IA_A)
}

func (c *Context) sla_H() {
	*c.R1.H = c.doSL(*c.R1.H, IA_A)
}

func (c *Context) sla_L() {
	*c.R1.L = c.doSL(*c.R1.L, IA_A)
}

func (c *Context) sla_IXh() {
	*c.R1.IXh = c.doSL(*c.R1.IXh, IA_A)
}

func (c *Context) sla_IXl() {
	*c.R1.IXl = c.doSL(*c.R1.IXl, IA_A)
}

func (c *Context) sla_IYh() {
	*c.R1.IYh = c.doSL(*c.R1.IYh, IA_A)
}

func (c *Context) sla_IYl() {
	*c.R1.IYl = c.doSL(*c.R1.IYl, IA_A)
}

func (c *Context) srl_A() {
	*c.R1.A = c.doSR(*c.R1.A, IA_L)
}

func (c *Context) srl_B() {
	*c.R1.B = c.doSR(*c.R1.B, IA_L)
}

func (c *Context) srl_C() {
	*c.R1.C = c.doSR(*c.R1.C, IA_L)
}

func (c *Context) srl_D() {
	*c.R1.D = c.doSR(*c.R1.D, IA_L)
}

func (c *Context) srl_E() {
	*c.R1.E = c.doSR(*c.R1.E, IA_L)
}

func (c *Context) srl_H() {
	*c.R1.H = c.doSR(*c.R1.H, IA_L)
}

func (c *Context) srl_L() {
	*c.R1.L = c.doSR(*c.R1.L, IA_L)
}

func (c *Context) srl_IXh() {
	*c.R1.IXh = c.doSR(*c.R1.IXh, IA_L)
}

func (c *Context) srl_IXl() {
	*c.R1.IXl = c.doSR(*c.R1.IXl, IA_L)
}

func (c *Context) srl_IYh() {
	*c.R1.IYh = c.doSR(*c.R1.IYh, IA_L)
}

func (c *Context) srl_IYl() {
	*c.R1.IYl = c.doSR(*c.R1.IYl, IA_L)
}

func (c *Context) sra_A() {
	*c.R1.A = c.doSR(*c.R1.A, IA_A)
}

func (c *Context) sra_B() {
	*c.R1.B = c.doSR(*c.R1.B, IA_A)
}

func (c *Context) sra_C() {
	*c.R1.C = c.doSR(*c.R1.C, IA_A)
}

func (c *Context) sra_D() {
	*c.R1.D = c.doSR(*c.R1.D, IA_A)
}

func (c *Context) sra_E() {
	*c.R1.E = c.doSR(*c.R1.E, IA_A)
}

func (c *Context) sra_H() {
	*c.R1.H = c.doSR(*c.R1.H, IA_A)
}

func (c *Context) sra_L() {
	*c.R1.L = c.doSR(*c.R1.L, IA_A)
}

func (c *Context) sra_IXh() {
	*c.R1.IXh = c.doSR(*c.R1.IXh, IA_A)
}

func (c *Context) sra_IXl() {
	*c.R1.IXl = c.doSR(*c.R1.IXl, IA_A)
}

func (c *Context) sra_IYh() {
	*c.R1.IYh = c.doSR(*c.R1.IYh, IA_A)
}

func (c *Context) sra_IYl() {
	*c.R1.IYl = c.doSR(*c.R1.IYl, IA_A)
}

func (c *Context) ld_off_BC_A() {
	c.write8(*c.R1.BC, *c.R1.A)
}

func (c *Context) ld_off_DE_A() {
	c.write8(*c.R1.DE, *c.R1.A)
}

func (c *Context) ld_off_HL_A() {
	c.write8(*c.R1.HL, *c.R1.A)
}

func (c *Context) ld_off_HL_B() {
	c.write8(*c.R1.HL, *c.R1.B)
}

func (c *Context) ld_off_HL_C() {
	c.write8(*c.R1.HL, *c.R1.C)
}

func (c *Context) ld_off_HL_D() {
	c.write8(*c.R1.HL, *c.R1.D)
}

func (c *Context) ld_off_HL_E() {
	c.write8(*c.R1.HL, *c.R1.E)
}

func (c *Context) ld_off_HL_H() {
	c.write8(*c.R1.HL, *c.R1.H)
}

func (c *Context) ld_off_HL_L() {
	c.write8(*c.R1.HL, *c.R1.L)
}

func (c *Context) ld_off_HL_n() {
	c.write8(*c.R1.HL, c.read8(c.PC))
	c.PC++
}

func (c *Context) ld_off_IX_d_A() {
	c.TStates += 5
	addr := doOff(*c.R1.IX, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_off_IX_d_B() {
	c.TStates += 5
	addr := doOff(*c.R1.IX, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_off_IX_d_C() {
	c.TStates += 5
	addr := doOff(*c.R1.IX, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_off_IX_d_D() {
	c.TStates += 5
	addr := doOff(*c.R1.IX, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_off_IX_d_E() {
	c.TStates += 5
	addr := doOff(*c.R1.IX, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_off_IX_d_H() {
	c.TStates += 5
	addr := doOff(*c.R1.IX, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_off_IX_d_L() {
	c.TStates += 5
	addr := doOff(*c.R1.IX, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_off_IY_d_A() {
	c.TStates += 5
	addr := doOff(*c.R1.IY, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_off_IY_d_B() {
	c.TStates += 5
	addr := doOff(*c.R1.IY, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_off_IY_d_C() {
	c.TStates += 5
	addr := doOff(*c.R1.IY, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_off_IY_d_D() {
	c.TStates += 5
	addr := doOff(*c.R1.IY, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_off_IY_d_E() {
	c.TStates += 5
	addr := doOff(*c.R1.IY, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_off_IY_d_H() {
	c.TStates += 5
	addr := doOff(*c.R1.IY, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_off_IY_d_L() {
	c.TStates += 5
	addr := doOff(*c.R1.IY, int(c.read8(c.PC)))
	c.PC++
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_off_IX_d_n() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	n := c.read8(c.PC)
	c.PC++
	addr := doOff(*c.R1.IX, off)
	c.write8(addr, n)
}

func (c *Context) ld_off_IY_d_n() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	n := c.read8(c.PC)
	c.PC++
	addr := doOff(*c.R1.IY, off)
	c.write8(addr, n)
}

func (c *Context) ld_off_nn_A() {
	c.write8(c.read16(c.PC), *c.R1.A)
	c.PC += 2
}

func (c *Context) ld_BC_BC() {
	c.TStates += 2
	*c.R1.BC = *c.R1.BC
}

func (c *Context) ld_BC_DE() {
	c.TStates += 2
	*c.R1.BC = *c.R1.DE
}

func (c *Context) ld_BC_HL() {
	c.TStates += 2
	*c.R1.BC = *c.R1.HL
}

func (c *Context) ld_BC_IX() {
	c.TStates += 2
	*c.R1.BC = *c.R1.IX
}

func (c *Context) ld_BC_IY() {
	c.TStates += 2
	*c.R1.BC = *c.R1.IY
}

func (c *Context) ld_BC_SP() {
	c.TStates += 2
	*c.R1.BC = *c.R1.SP
}

func (c *Context) ld_DE_BC() {
	c.TStates += 2
	*c.R1.DE = *c.R1.BC
}

func (c *Context) ld_DE_DE() {
	c.TStates += 2
	*c.R1.DE = *c.R1.DE
}

func (c *Context) ld_DE_HL() {
	c.TStates += 2
	*c.R1.DE = *c.R1.HL
}

func (c *Context) ld_DE_IX() {
	c.TStates += 2
	*c.R1.DE = *c.R1.IX
}

func (c *Context) ld_DE_IY() {
	c.TStates += 2
	*c.R1.DE = *c.R1.IY
}

func (c *Context) ld_DE_SP() {
	c.TStates += 2
	*c.R1.DE = *c.R1.SP
}

func (c *Context) ld_HL_BC() {
	c.TStates += 2
	*c.R1.HL = *c.R1.BC
}

func (c *Context) ld_HL_DE() {
	c.TStates += 2
	*c.R1.HL = *c.R1.DE
}

func (c *Context) ld_HL_HL() {
	c.TStates += 2
	*c.R1.HL = *c.R1.HL
}

func (c *Context) ld_HL_IX() {
	c.TStates += 2
	*c.R1.HL = *c.R1.IX
}

func (c *Context) ld_HL_IY() {
	c.TStates += 2
	*c.R1.HL = *c.R1.IY
}

func (c *Context) ld_HL_SP() {
	c.TStates += 2
	*c.R1.HL = *c.R1.SP
}

func (c *Context) ld_IX_BC() {
	c.TStates += 2
	*c.R1.IX = *c.R1.BC
}

func (c *Context) ld_IX_DE() {
	c.TStates += 2
	*c.R1.IX = *c.R1.DE
}

func (c *Context) ld_IX_HL() {
	c.TStates += 2
	*c.R1.IX = *c.R1.HL
}

func (c *Context) ld_IX_IX() {
	c.TStates += 2
	*c.R1.IX = *c.R1.IX
}

func (c *Context) ld_IX_IY() {
	c.TStates += 2
	*c.R1.IX = *c.R1.IY
}

func (c *Context) ld_IX_SP() {
	c.TStates += 2
	*c.R1.IX = *c.R1.SP
}

func (c *Context) ld_IY_BC() {
	c.TStates += 2
	*c.R1.IY = *c.R1.BC
}

func (c *Context) ld_IY_DE() {
	c.TStates += 2
	*c.R1.IY = *c.R1.DE
}

func (c *Context) ld_IY_HL() {
	c.TStates += 2
	*c.R1.IY = *c.R1.HL
}

func (c *Context) ld_IY_IX() {
	c.TStates += 2
	*c.R1.IY = *c.R1.IX
}

func (c *Context) ld_IY_IY() {
	c.TStates += 2
	*c.R1.IY = *c.R1.IY
}

func (c *Context) ld_IY_SP() {
	c.TStates += 2
	*c.R1.IY = *c.R1.SP
}

func (c *Context) ld_SP_BC() {
	c.TStates += 2
	*c.R1.SP = *c.R1.BC
}

func (c *Context) ld_SP_DE() {
	c.TStates += 2
	*c.R1.SP = *c.R1.DE
}

func (c *Context) ld_SP_HL() {
	c.TStates += 2
	*c.R1.SP = *c.R1.HL
}

func (c *Context) ld_SP_IX() {
	c.TStates += 2
	*c.R1.SP = *c.R1.IX
}

func (c *Context) ld_SP_IY() {
	c.TStates += 2
	*c.R1.SP = *c.R1.IY
}

func (c *Context) ld_SP_SP() {
	c.TStates += 2
	*c.R1.SP = *c.R1.SP
}

func (c *Context) ld_off_nn_BC() {
	c.write16(c.read16(c.PC), *c.R1.BC)
	c.PC += 2
}

func (c *Context) ld_off_nn_DE() {
	c.write16(c.read16(c.PC), *c.R1.DE)
	c.PC += 2
}

func (c *Context) ld_off_nn_HL() {
	c.write16(c.read16(c.PC), *c.R1.HL)
	c.PC += 2
}

func (c *Context) ld_off_nn_IX() {
	c.write16(c.read16(c.PC), *c.R1.IX)
	c.PC += 2
}

func (c *Context) ld_off_nn_IY() {
	c.write16(c.read16(c.PC), *c.R1.IY)
	c.PC += 2
}

func (c *Context) ld_off_nn_SP() {
	c.write16(c.read16(c.PC), *c.R1.SP)
	c.PC += 2
}

func (c *Context) ld_A_off_BC() {
	*c.R1.A = c.read8(*c.R1.BC)
}

func (c *Context) ld_A_off_DE() {
	*c.R1.A = c.read8(*c.R1.DE)
}

func (c *Context) ld_A_off_HL() {
	*c.R1.A = c.read8(*c.R1.HL)
}

func (c *Context) ld_B_off_HL() {
	*c.R1.B = c.read8(*c.R1.HL)
}

func (c *Context) ld_C_off_HL() {
	*c.R1.C = c.read8(*c.R1.HL)
}

func (c *Context) ld_D_off_HL() {
	*c.R1.D = c.read8(*c.R1.HL)
}

func (c *Context) ld_E_off_HL() {
	*c.R1.E = c.read8(*c.R1.HL)
}

func (c *Context) ld_H_off_HL() {
	*c.R1.H = c.read8(*c.R1.HL)
}

func (c *Context) ld_L_off_HL() {
	*c.R1.L = c.read8(*c.R1.HL)
}

func (c *Context) ld_A_off_IX_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.read8(addr)
}

func (c *Context) ld_A_off_IY_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.read8(addr)
}

func (c *Context) ld_B_off_IX_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.read8(addr)
}

func (c *Context) ld_B_off_IY_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.read8(addr)
}

func (c *Context) ld_C_off_IX_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.read8(addr)
}

func (c *Context) ld_C_off_IY_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.read8(addr)
}

func (c *Context) ld_D_off_IX_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.read8(addr)
}

func (c *Context) ld_D_off_IY_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.read8(addr)
}

func (c *Context) ld_E_off_IX_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.read8(addr)
}

func (c *Context) ld_E_off_IY_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.read8(addr)
}

func (c *Context) ld_H_off_IX_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.read8(addr)
}

func (c *Context) ld_H_off_IY_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.read8(addr)
}

func (c *Context) ld_L_off_IX_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.read8(addr)
}

func (c *Context) ld_L_off_IY_d() {
	c.TStates += 5
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.read8(addr)
}

func (c *Context) ld_A_off_nn() {
	*c.R1.A = c.read8(c.read16(c.PC))
	c.PC += 2
}

func (c *Context) ld_B_off_nn() {
	*c.R1.B = c.read8(c.read16(c.PC))
	c.PC += 2
}

func (c *Context) ld_C_off_nn() {
	*c.R1.C = c.read8(c.read16(c.PC))
	c.PC += 2
}

func (c *Context) ld_D_off_nn() {
	*c.R1.D = c.read8(c.read16(c.PC))
	c.PC += 2
}

func (c *Context) ld_E_off_nn() {
	*c.R1.E = c.read8(c.read16(c.PC))
	c.PC += 2
}

func (c *Context) ld_H_off_nn() {
	*c.R1.H = c.read8(c.read16(c.PC))
	c.PC += 2
}

func (c *Context) ld_L_off_nn() {
	*c.R1.L = c.read8(c.read16(c.PC))
	c.PC += 2
}

func (c *Context) ld_A_A() {
	*c.R1.A = *c.R1.A
}

func (c *Context) ld_A_B() {
	*c.R1.A = *c.R1.B
}

func (c *Context) ld_A_C() {
	*c.R1.A = *c.R1.C
}

func (c *Context) ld_A_D() {
	*c.R1.A = *c.R1.D
}

func (c *Context) ld_A_E() {
	*c.R1.A = *c.R1.E
}

func (c *Context) ld_A_H() {
	*c.R1.A = *c.R1.H
}

func (c *Context) ld_A_L() {
	*c.R1.A = *c.R1.L
}

func (c *Context) ld_A_IXh() {
	*c.R1.A = *c.R1.IXh
}

func (c *Context) ld_A_IXl() {
	*c.R1.A = *c.R1.IXl
}

func (c *Context) ld_A_IYh() {
	*c.R1.A = *c.R1.IYh
}

func (c *Context) ld_A_IYl() {
	*c.R1.A = *c.R1.IYl
}

func (c *Context) ld_B_A() {
	*c.R1.B = *c.R1.A
}

func (c *Context) ld_B_B() {
	*c.R1.B = *c.R1.B
}

func (c *Context) ld_B_C() {
	*c.R1.B = *c.R1.C
}

func (c *Context) ld_B_D() {
	*c.R1.B = *c.R1.D
}

func (c *Context) ld_B_E() {
	*c.R1.B = *c.R1.E
}

func (c *Context) ld_B_H() {
	*c.R1.B = *c.R1.H
}

func (c *Context) ld_B_L() {
	*c.R1.B = *c.R1.L
}

func (c *Context) ld_B_IXh() {
	*c.R1.B = *c.R1.IXh
}

func (c *Context) ld_B_IXl() {
	*c.R1.B = *c.R1.IXl
}

func (c *Context) ld_B_IYh() {
	*c.R1.B = *c.R1.IYh
}

func (c *Context) ld_B_IYl() {
	*c.R1.B = *c.R1.IYl
}

func (c *Context) ld_C_A() {
	*c.R1.C = *c.R1.A
}

func (c *Context) ld_C_B() {
	*c.R1.C = *c.R1.B
}

func (c *Context) ld_C_C() {
	*c.R1.C = *c.R1.C
}

func (c *Context) ld_C_D() {
	*c.R1.C = *c.R1.D
}

func (c *Context) ld_C_E() {
	*c.R1.C = *c.R1.E
}

func (c *Context) ld_C_H() {
	*c.R1.C = *c.R1.H
}

func (c *Context) ld_C_L() {
	*c.R1.C = *c.R1.L
}

func (c *Context) ld_C_IXh() {
	*c.R1.C = *c.R1.IXh
}

func (c *Context) ld_C_IXl() {
	*c.R1.C = *c.R1.IXl
}

func (c *Context) ld_C_IYh() {
	*c.R1.C = *c.R1.IYh
}

func (c *Context) ld_C_IYl() {
	*c.R1.C = *c.R1.IYl
}

func (c *Context) ld_D_A() {
	*c.R1.D = *c.R1.A
}

func (c *Context) ld_D_B() {
	*c.R1.D = *c.R1.B
}

func (c *Context) ld_D_C() {
	*c.R1.D = *c.R1.C
}

func (c *Context) ld_D_D() {
	*c.R1.D = *c.R1.D
}

func (c *Context) ld_D_E() {
	*c.R1.D = *c.R1.E
}

func (c *Context) ld_D_H() {
	*c.R1.D = *c.R1.H
}

func (c *Context) ld_D_L() {
	*c.R1.D = *c.R1.L
}

func (c *Context) ld_D_IXh() {
	*c.R1.D = *c.R1.IXh
}

func (c *Context) ld_D_IXl() {
	*c.R1.D = *c.R1.IXl
}

func (c *Context) ld_D_IYh() {
	*c.R1.D = *c.R1.IYh
}

func (c *Context) ld_D_IYl() {
	*c.R1.D = *c.R1.IYl
}

func (c *Context) ld_E_A() {
	*c.R1.E = *c.R1.A
}

func (c *Context) ld_E_B() {
	*c.R1.E = *c.R1.B
}

func (c *Context) ld_E_C() {
	*c.R1.E = *c.R1.C
}

func (c *Context) ld_E_D() {
	*c.R1.E = *c.R1.D
}

func (c *Context) ld_E_E() {
	*c.R1.E = *c.R1.E
}

func (c *Context) ld_E_H() {
	*c.R1.E = *c.R1.H
}

func (c *Context) ld_E_L() {
	*c.R1.E = *c.R1.L
}

func (c *Context) ld_E_IXh() {
	*c.R1.E = *c.R1.IXh
}

func (c *Context) ld_E_IXl() {
	*c.R1.E = *c.R1.IXl
}

func (c *Context) ld_E_IYh() {
	*c.R1.E = *c.R1.IYh
}

func (c *Context) ld_E_IYl() {
	*c.R1.E = *c.R1.IYl
}

func (c *Context) ld_H_A() {
	*c.R1.H = *c.R1.A
}

func (c *Context) ld_H_B() {
	*c.R1.H = *c.R1.B
}

func (c *Context) ld_H_C() {
	*c.R1.H = *c.R1.C
}

func (c *Context) ld_H_D() {
	*c.R1.H = *c.R1.D
}

func (c *Context) ld_H_E() {
	*c.R1.H = *c.R1.E
}

func (c *Context) ld_H_H() {
	*c.R1.H = *c.R1.H
}

func (c *Context) ld_H_L() {
	*c.R1.H = *c.R1.L
}

func (c *Context) ld_H_IXh() {
	*c.R1.H = *c.R1.IXh
}

func (c *Context) ld_H_IXl() {
	*c.R1.H = *c.R1.IXl
}

func (c *Context) ld_H_IYh() {
	*c.R1.H = *c.R1.IYh
}

func (c *Context) ld_H_IYl() {
	*c.R1.H = *c.R1.IYl
}

func (c *Context) ld_L_A() {
	*c.R1.L = *c.R1.A
}

func (c *Context) ld_L_B() {
	*c.R1.L = *c.R1.B
}

func (c *Context) ld_L_C() {
	*c.R1.L = *c.R1.C
}

func (c *Context) ld_L_D() {
	*c.R1.L = *c.R1.D
}

func (c *Context) ld_L_E() {
	*c.R1.L = *c.R1.E
}

func (c *Context) ld_L_H() {
	*c.R1.L = *c.R1.H
}

func (c *Context) ld_L_L() {
	*c.R1.L = *c.R1.L
}

func (c *Context) ld_L_IXh() {
	*c.R1.L = *c.R1.IXh
}

func (c *Context) ld_L_IXl() {
	*c.R1.L = *c.R1.IXl
}

func (c *Context) ld_L_IYh() {
	*c.R1.L = *c.R1.IYh
}

func (c *Context) ld_L_IYl() {
	*c.R1.L = *c.R1.IYl
}

func (c *Context) ld_IXh_A() {
	*c.R1.IXh = *c.R1.A
}

func (c *Context) ld_IXh_B() {
	*c.R1.IXh = *c.R1.B
}

func (c *Context) ld_IXh_C() {
	*c.R1.IXh = *c.R1.C
}

func (c *Context) ld_IXh_D() {
	*c.R1.IXh = *c.R1.D
}

func (c *Context) ld_IXh_E() {
	*c.R1.IXh = *c.R1.E
}

func (c *Context) ld_IXh_H() {
	*c.R1.IXh = *c.R1.H
}

func (c *Context) ld_IXh_L() {
	*c.R1.IXh = *c.R1.L
}

func (c *Context) ld_IXh_IXh() {
	*c.R1.IXh = *c.R1.IXh
}

func (c *Context) ld_IXh_IXl() {
	*c.R1.IXh = *c.R1.IXl
}

func (c *Context) ld_IXh_IYh() {
	*c.R1.IXh = *c.R1.IYh
}

func (c *Context) ld_IXh_IYl() {
	*c.R1.IXh = *c.R1.IYl
}

func (c *Context) ld_IXl_A() {
	*c.R1.IXl = *c.R1.A
}

func (c *Context) ld_IXl_B() {
	*c.R1.IXl = *c.R1.B
}

func (c *Context) ld_IXl_C() {
	*c.R1.IXl = *c.R1.C
}

func (c *Context) ld_IXl_D() {
	*c.R1.IXl = *c.R1.D
}

func (c *Context) ld_IXl_E() {
	*c.R1.IXl = *c.R1.E
}

func (c *Context) ld_IXl_H() {
	*c.R1.IXl = *c.R1.H
}

func (c *Context) ld_IXl_L() {
	*c.R1.IXl = *c.R1.L
}

func (c *Context) ld_IXl_IXh() {
	*c.R1.IXl = *c.R1.IXh
}

func (c *Context) ld_IXl_IXl() {
	*c.R1.IXl = *c.R1.IXl
}

func (c *Context) ld_IXl_IYh() {
	*c.R1.IXl = *c.R1.IYh
}

func (c *Context) ld_IXl_IYl() {
	*c.R1.IXl = *c.R1.IYl
}

func (c *Context) ld_IYh_A() {
	*c.R1.IYh = *c.R1.A
}

func (c *Context) ld_IYh_B() {
	*c.R1.IYh = *c.R1.B
}

func (c *Context) ld_IYh_C() {
	*c.R1.IYh = *c.R1.C
}

func (c *Context) ld_IYh_D() {
	*c.R1.IYh = *c.R1.D
}

func (c *Context) ld_IYh_E() {
	*c.R1.IYh = *c.R1.E
}

func (c *Context) ld_IYh_H() {
	*c.R1.IYh = *c.R1.H
}

func (c *Context) ld_IYh_L() {
	*c.R1.IYh = *c.R1.L
}

func (c *Context) ld_IYh_IXh() {
	*c.R1.IYh = *c.R1.IXh
}

func (c *Context) ld_IYh_IXl() {
	*c.R1.IYh = *c.R1.IXl
}

func (c *Context) ld_IYh_IYh() {
	*c.R1.IYh = *c.R1.IYh
}

func (c *Context) ld_IYh_IYl() {
	*c.R1.IYh = *c.R1.IYl
}

func (c *Context) ld_IYl_A() {
	*c.R1.IYl = *c.R1.A
}

func (c *Context) ld_IYl_B() {
	*c.R1.IYl = *c.R1.B
}

func (c *Context) ld_IYl_C() {
	*c.R1.IYl = *c.R1.C
}

func (c *Context) ld_IYl_D() {
	*c.R1.IYl = *c.R1.D
}

func (c *Context) ld_IYl_E() {
	*c.R1.IYl = *c.R1.E
}

func (c *Context) ld_IYl_H() {
	*c.R1.IYl = *c.R1.H
}

func (c *Context) ld_IYl_L() {
	*c.R1.IYl = *c.R1.L
}

func (c *Context) ld_IYl_IXh() {
	*c.R1.IYl = *c.R1.IXh
}

func (c *Context) ld_IYl_IXl() {
	*c.R1.IYl = *c.R1.IXl
}

func (c *Context) ld_IYl_IYh() {
	*c.R1.IYl = *c.R1.IYh
}

func (c *Context) ld_IYl_IYl() {
	*c.R1.IYl = *c.R1.IYl
}

func (c *Context) ld_A_SLA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SLA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SRA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SRA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_B_SLA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SLA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SRA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SRA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_C_SLA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SLA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SRA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SRA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_D_SLA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SLA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SRA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SRA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_E_SLA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SLA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SRA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SRA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_H_SLA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SLA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SRA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SRA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_L_SLA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SLA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSL(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SRA_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SRA_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSR(c.read8(addr), IA_A)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_A_SLL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SLL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SRL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SRL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_B_SLL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SLL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SRL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SRL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_C_SLL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SLL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SRL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SRL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_D_SLL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SLL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SRL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SRL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_E_SLL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SLL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SRL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SRL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_H_SLL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SLL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SRL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SRL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_L_SLL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SLL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSL(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SRL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SRL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSR(c.read8(addr), IA_L)
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_A_RL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RLC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RLC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RR_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RR_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RRC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RRC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_B_RL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RLC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RLC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RR_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RR_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RRC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RRC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_C_RL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RLC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RLC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RR_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RR_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RRC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RRC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_D_RL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RLC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RLC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RR_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RR_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RRC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RRC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_E_RL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RLC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RLC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RR_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RR_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RRC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RRC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_H_RL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RLC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RLC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RR_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RR_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RRC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RRC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_L_RL_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RL_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doRL(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RLC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RLC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doRLC(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RR_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RR_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doRR(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RRC_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RRC_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doRRC(true, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_A_SET_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_SET_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.A = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_A_RES_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.A = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.A)
}

func (c *Context) ld_B_SET_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_SET_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.B = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_B_RES_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.B = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.B)
}

func (c *Context) ld_C_SET_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_SET_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.C = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_C_RES_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.C = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.C)
}

func (c *Context) ld_D_SET_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_SET_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.D = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_D_RES_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.D = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.D)
}

func (c *Context) ld_E_SET_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_SET_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.E = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_E_RES_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.E = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.E)
}

func (c *Context) ld_H_SET_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_SET_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.H = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_H_RES_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.H = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.H)
}

func (c *Context) ld_L_SET_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 0, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 1, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 2, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 3, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 4, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 5, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 6, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_SET_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_SET, 7, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_0_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_0_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 0, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_1_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_1_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 1, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_2_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_2_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 2, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_3_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_3_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 3, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_4_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_4_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 4, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_5_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_5_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 5, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_6_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_6_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 6, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_7_off_IX_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IX, off)
	*c.R1.L = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_L_RES_7_off_IY_d() {
	c.TStates += 2
	off := doComplement(c.read8(c.PC))
	c.PC++
	addr := doOff(*c.R1.IY, off)
	*c.R1.L = c.doSetRes(SR_RES, 7, c.read8(addr))
	c.write8(addr, *c.R1.L)
}

func (c *Context) ld_A_I() {
	c.TStates += 1
	*c.R1.A = c.I
	c.adjustFlags(*c.R1.A)
	c.resFlag(F_H | F_N)
	c.valFlag(F_PV, c.IFF2)
	c.valFlag(F_S, (*c.R1.A&0x80) != 0)
	c.valFlag(F_Z, *c.R1.A == 0)
}

func (c *Context) ld_A_R() {
	c.TStates += 1
	*c.R1.A = c.R
	c.adjustFlags(*c.R1.A)
	c.resFlag(F_H | F_N)
	c.valFlag(F_PV, c.IFF2)
	c.valFlag(F_S, (*c.R1.A&0x80) != 0)
	c.valFlag(F_Z, *c.R1.A == 0)
}

func (c *Context) ld_I_A() {
	c.TStates += 1
	c.I = *c.R1.A
}

func (c *Context) ld_R_A() {
	c.TStates += 1
	c.R = *c.R1.A
}

func (c *Context) ld_A_n() {
	*c.R1.A = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_B_n() {
	*c.R1.B = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_C_n() {
	*c.R1.C = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_D_n() {
	*c.R1.D = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_E_n() {
	*c.R1.E = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_H_n() {
	*c.R1.H = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_L_n() {
	*c.R1.L = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_IXh_n() {
	*c.R1.IXh = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_IXl_n() {
	*c.R1.IXl = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_IYh_n() {
	*c.R1.IYh = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_IYl_n() {
	*c.R1.IYl = c.read8(c.PC)
	c.PC++
}

func (c *Context) ld_BC_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	*c.R1.BC = c.read16(addr)
}

func (c *Context) ld_DE_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	*c.R1.DE = c.read16(addr)
}

func (c *Context) ld_HL_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	*c.R1.HL = c.read16(addr)
}

func (c *Context) ld_SP_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	*c.R1.SP = c.read16(addr)
}

func (c *Context) ld_IX_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	*c.R1.IX = c.read16(addr)
}

func (c *Context) ld_IY_off_nn() {
	addr := c.read16(c.PC)
	c.PC += 2
	*c.R1.IY = c.read16(addr)
}

func (c *Context) ld_BC_nn() {
	*c.R1.BC = c.read16(c.PC)
	c.PC += 2
}

func (c *Context) ld_DE_nn() {
	*c.R1.DE = c.read16(c.PC)
	c.PC += 2
}

func (c *Context) ld_HL_nn() {
	*c.R1.HL = c.read16(c.PC)
	c.PC += 2
}

func (c *Context) ld_SP_nn() {
	*c.R1.SP = c.read16(c.PC)
	c.PC += 2
}

func (c *Context) ld_IX_nn() {
	*c.R1.IX = c.read16(c.PC)
	c.PC += 2
}

func (c *Context) ld_IY_nn() {
	*c.R1.IY = c.read16(c.PC)
	c.PC += 2
}

func (c *Context) ldir() {
	c.ldi()
	if *c.R1.BC != 0 {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) ldi() {
	c.TStates += 2
	val := c.read8(*c.R1.HL)
	c.write8(*c.R1.DE, val)
	*c.R1.DE++
	*c.R1.HL++
	*c.R1.BC--
	c.valFlag(F_5, (*c.R1.A+val)&0x02 != 0)
	c.valFlag(F_3, Z80Flag(*c.R1.A+val)&F_3 != 0)
	c.resFlag(F_H | F_N)
	c.valFlag(F_PV, *c.R1.BC != 0)
}

func (c *Context) lddr() {
	c.ldd()
	if *c.R1.BC != 0 {
		c.TStates += 5
		c.PC -= 2
	}
}

func (c *Context) ldd() {
	c.TStates += 2
	val := c.read8(*c.R1.HL)
	c.write8(*c.R1.DE, val)
	*c.R1.DE--
	*c.R1.HL--
	*c.R1.BC--
	c.valFlag(F_5, (*c.R1.A+val)&0x02 != 0)
	c.valFlag(F_3, Z80Flag(*c.R1.A+val)&F_3 != 0)
	c.resFlag(F_H | F_N)
	c.valFlag(F_PV, *c.R1.BC != 0)
}

func (c *Context) neg() {
	temp := *c.R1.A
	*c.R1.A = 0
	*c.R1.A = c.doArithmetic(temp, false, true)
	c.setFlag(F_N)
}

func (c *Context) nop() {
	// NOP
}
