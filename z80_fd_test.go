package z80

import "testing"


func TestFD09(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x466A
  *c.R1.BC = 0xA623
  *c.R1.DE = 0xBAB2
  *c.R1.HL = 0xD788
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC9E8
  *c.R1.IY = 0xF698
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x09
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4649) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4649, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA623) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA623, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBAB2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBAB2, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD788) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD788, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC9E8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC9E8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9CBB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9CBB, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 15) {
    t.Errorf("TStates mismatch: %d expected, got %d", 15, c.TStates)
  }
}

func TestFD19(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB3E5
  *c.R1.BC = 0x5336
  *c.R1.DE = 0x76CB
  *c.R1.HL = 0x54E2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB9CE
  *c.R1.IY = 0x8624
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x19
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB3EC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB3EC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5336) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5336, *c.R1.BC)
  }
  if (*c.R1.DE != 0x76CB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x76CB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x54E2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x54E2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB9CE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB9CE, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFCEF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFCEF, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 15) {
    t.Errorf("TStates mismatch: %d expected, got %d", 15, c.TStates)
  }
}

func TestFD21(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC924
  *c.R1.BC = 0x5C83
  *c.R1.DE = 0xE0E2
  *c.R1.HL = 0xEDDB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6E9F
  *c.R1.IY = 0xBA55
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x21
  mem[0x0002] = 0x46
  mem[0x0003] = 0x47
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC924) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC924, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5C83) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5C83, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE0E2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE0E2, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEDDB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEDDB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6E9F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6E9F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4746) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4746, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 14) {
    t.Errorf("TStates mismatch: %d expected, got %d", 14, c.TStates)
  }
}

func TestFD22(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1235
  *c.R1.BC = 0xF0B6
  *c.R1.DE = 0xB74C
  *c.R1.HL = 0xCC9F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8B00
  *c.R1.IY = 0x81E4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x22
  mem[0x0002] = 0x9A
  mem[0x0003] = 0xE2
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1235) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1235, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF0B6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF0B6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB74C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB74C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCC9F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCC9F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8B00) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8B00, *c.R1.IX)
  }
  if (*c.R1.IY != 0x81E4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x81E4, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 20) {
    t.Errorf("TStates mismatch: %d expected, got %d", 20, c.TStates)
  }
}

func TestFD23(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x69F2
  *c.R1.BC = 0xC1D3
  *c.R1.DE = 0x0F6F
  *c.R1.HL = 0x2169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE39E
  *c.R1.IY = 0x2605
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x23
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x69F2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x69F2, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC1D3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC1D3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0F6F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0F6F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE39E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE39E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2606) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2606, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestFD24(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5554
  *c.R1.BC = 0x9684
  *c.R1.DE = 0xD36A
  *c.R1.HL = 0xDAC3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7803
  *c.R1.IY = 0x6434
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x24
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5520) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5520, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9684) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9684, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD36A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD36A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDAC3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDAC3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7803) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7803, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6534) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6534, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD25(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCD0B
  *c.R1.BC = 0xB5E4
  *c.R1.DE = 0xA754
  *c.R1.HL = 0x9526
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3DCB
  *c.R1.IY = 0x03B2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x25
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCD03) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCD03, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB5E4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB5E4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA754) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA754, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9526) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9526, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3DCB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3DCB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x02B2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x02B2, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD26(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2452
  *c.R1.BC = 0x300B
  *c.R1.DE = 0xB4A1
  *c.R1.HL = 0x929D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC259
  *c.R1.IY = 0x3F30
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x26
  mem[0x0002] = 0x77
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2452) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2452, *c.R1.AF)
  }
  if (*c.R1.BC != 0x300B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x300B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB4A1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB4A1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x929D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x929D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC259) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC259, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7730) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7730, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestFD29(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5812
  *c.R1.BC = 0x49D0
  *c.R1.DE = 0xEC95
  *c.R1.HL = 0x011C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEC6C
  *c.R1.IY = 0x594C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x29
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5830) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5830, *c.R1.AF)
  }
  if (*c.R1.BC != 0x49D0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x49D0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEC95) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEC95, *c.R1.DE)
  }
  if (*c.R1.HL != 0x011C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x011C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEC6C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEC6C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB298) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB298, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 15) {
    t.Errorf("TStates mismatch: %d expected, got %d", 15, c.TStates)
  }
}

func TestFD2A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0F82
  *c.R1.BC = 0x3198
  *c.R1.DE = 0x87E3
  *c.R1.HL = 0x7C1C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1BB4
  *c.R1.IY = 0xEB1A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x2A
  mem[0x0002] = 0x91
  mem[0x0003] = 0xF9
  mem[0xF991] = 0x92
  mem[0xF992] = 0xBF
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0F82) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0F82, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3198) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3198, *c.R1.BC)
  }
  if (*c.R1.DE != 0x87E3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x87E3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7C1C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7C1C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1BB4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1BB4, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBF92) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBF92, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 20) {
    t.Errorf("TStates mismatch: %d expected, got %d", 20, c.TStates)
  }
}

func TestFD2B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAB27
  *c.R1.BC = 0x942F
  *c.R1.DE = 0x82FA
  *c.R1.HL = 0x6F2F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9438
  *c.R1.IY = 0xEBBC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x2B
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAB27) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAB27, *c.R1.AF)
  }
  if (*c.R1.BC != 0x942F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x942F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x82FA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x82FA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6F2F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6F2F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9438) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9438, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEBBB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEBBB, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestFD2C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x665D
  *c.R1.BC = 0x0AB1
  *c.R1.DE = 0x5656
  *c.R1.HL = 0xE5A9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5FB9
  *c.R1.IY = 0x4DF7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x2C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x66A9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x66A9, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0AB1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0AB1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5656) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5656, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE5A9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE5A9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5FB9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5FB9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4DF8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4DF8, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD2D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x32FB
  *c.R1.BC = 0xF78A
  *c.R1.DE = 0xB906
  *c.R1.HL = 0x31D0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC72A
  *c.R1.IY = 0xE91C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x2D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x320B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x320B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF78A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF78A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB906) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB906, *c.R1.DE)
  }
  if (*c.R1.HL != 0x31D0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x31D0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC72A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC72A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE91B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE91B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD2E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2114
  *c.R1.BC = 0x4923
  *c.R1.DE = 0x6E65
  *c.R1.HL = 0x006C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDA39
  *c.R1.IY = 0xC0CB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x2E
  mem[0x0002] = 0x49
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2114) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2114, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4923) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4923, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6E65) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6E65, *c.R1.DE)
  }
  if (*c.R1.HL != 0x006C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x006C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDA39) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDA39, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC049) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC049, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestFD34(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD56A
  *c.R1.BC = 0x6F24
  *c.R1.DE = 0x7DF7
  *c.R1.HL = 0x74F0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x365A
  *c.R1.IY = 0xEFC4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x34
  mem[0x0002] = 0xB8
  mem[0xEF7C] = 0xE0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD5A0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD5A0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6F24) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6F24, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7DF7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7DF7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x74F0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x74F0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x365A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x365A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEFC4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEFC4, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 23) {
    t.Errorf("TStates mismatch: %d expected, got %d", 23, c.TStates)
  }
}

func TestFD35(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8CDA
  *c.R1.BC = 0x35D8
  *c.R1.DE = 0x7C1A
  *c.R1.HL = 0x1C0A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x62BB
  *c.R1.IY = 0xAEC6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x35
  mem[0x0002] = 0xAB
  mem[0xAE71] = 0xA6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8CA2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8CA2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x35D8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x35D8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7C1A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7C1A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1C0A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1C0A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x62BB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x62BB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAEC6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAEC6, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 23) {
    t.Errorf("TStates mismatch: %d expected, got %d", 23, c.TStates)
  }
}

func TestFD36(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE0F9
  *c.R1.BC = 0xAE1F
  *c.R1.DE = 0x4AEF
  *c.R1.HL = 0xC9D5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC0DB
  *c.R1.IY = 0xBDD4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x36
  mem[0x0002] = 0x81
  mem[0x0003] = 0xC5
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE0F9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE0F9, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAE1F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAE1F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4AEF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4AEF, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC9D5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC9D5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC0DB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC0DB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBDD4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBDD4, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD39(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2603
  *c.R1.BC = 0x726F
  *c.R1.DE = 0x9C7F
  *c.R1.HL = 0xCD46
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDC45
  *c.R1.IY = 0x54D5
  *c.R1.SP = 0xDC57
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x39
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2631) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2631, *c.R1.AF)
  }
  if (*c.R1.BC != 0x726F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x726F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9C7F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9C7F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCD46) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCD46, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDC45) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDC45, *c.R1.IX)
  }
  if (*c.R1.IY != 0x312C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x312C, *c.R1.IY)
  }
  if (*c.R1.SP != 0xDC57) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xDC57, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 15) {
    t.Errorf("TStates mismatch: %d expected, got %d", 15, c.TStates)
  }
}

func TestFD44(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0E58
  *c.R1.BC = 0x7192
  *c.R1.DE = 0x3580
  *c.R1.HL = 0x9BE4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1B79
  *c.R1.IY = 0x685E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x44
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0E58) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0E58, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6892) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6892, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3580) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3580, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9BE4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9BE4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1B79) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1B79, *c.R1.IX)
  }
  if (*c.R1.IY != 0x685E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x685E, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD45(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6555
  *c.R1.BC = 0xA488
  *c.R1.DE = 0x5AE8
  *c.R1.HL = 0xC948
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD7B8
  *c.R1.IY = 0xA177
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x45
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6555) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6555, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7788) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7788, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5AE8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5AE8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC948) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC948, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD7B8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD7B8, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA177) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA177, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD46(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x87F3
  *c.R1.BC = 0x17D5
  *c.R1.DE = 0x5EEA
  *c.R1.HL = 0x830B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDCEE
  *c.R1.IY = 0x3AFC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x46
  mem[0x0002] = 0x4D
  mem[0x3B49] = 0xC9
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x87F3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x87F3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC9D5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC9D5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5EEA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5EEA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x830B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x830B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDCEE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDCEE, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3AFC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3AFC, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD4C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7E6B
  *c.R1.BC = 0xBD4B
  *c.R1.DE = 0x24B6
  *c.R1.HL = 0xFF94
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x862D
  *c.R1.IY = 0x01D0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x4C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7E6B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7E6B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBD01) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBD01, *c.R1.BC)
  }
  if (*c.R1.DE != 0x24B6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x24B6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFF94) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFF94, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x862D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x862D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x01D0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x01D0, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD4D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x50CF
  *c.R1.BC = 0xE3FE
  *c.R1.DE = 0x998E
  *c.R1.HL = 0xDBA2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC4F5
  *c.R1.IY = 0xC7C9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x4D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x50CF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x50CF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE3C9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE3C9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x998E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x998E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDBA2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDBA2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC4F5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC4F5, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC7C9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC7C9, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD4E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2C0F
  *c.R1.BC = 0x69D7
  *c.R1.DE = 0x748A
  *c.R1.HL = 0x9290
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x904F
  *c.R1.IY = 0xBB9A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x4E
  mem[0x0002] = 0x67
  mem[0xBC01] = 0x9D
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2C0F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2C0F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x699D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x699D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x748A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x748A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9290) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9290, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x904F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x904F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBB9A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBB9A, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD54(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD7F9
  *c.R1.BC = 0xF65B
  *c.R1.DE = 0xB001
  *c.R1.HL = 0xD4C4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4B8E
  *c.R1.IY = 0xD437
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x54
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD7F9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD7F9, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF65B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF65B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD401) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD401, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD4C4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD4C4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4B8E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4B8E, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD437) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD437, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD55(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAB98
  *c.R1.BC = 0xFDAB
  *c.R1.DE = 0x254A
  *c.R1.HL = 0x010E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x126B
  *c.R1.IY = 0x13A9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x55
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAB98) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAB98, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFDAB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFDAB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA94A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA94A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x010E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x010E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x126B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x126B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x13A9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x13A9, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD56(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD3E8
  *c.R1.BC = 0xDF10
  *c.R1.DE = 0x5442
  *c.R1.HL = 0xB641
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA5A0
  *c.R1.IY = 0xFDA2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x56
  mem[0x0002] = 0xCE
  mem[0xFD70] = 0x78
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD3E8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD3E8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDF10) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDF10, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7842) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7842, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB641) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB641, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA5A0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA5A0, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFDA2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFDA2, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD5C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x11D5
  *c.R1.BC = 0xC489
  *c.R1.DE = 0xE220
  *c.R1.HL = 0x434E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3244
  *c.R1.IY = 0xD8BB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x5C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x11D5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x11D5, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC489) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC489, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE2D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE2D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x434E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x434E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3244) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3244, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD8BB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD8BB, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD5D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE945
  *c.R1.BC = 0xDBAE
  *c.R1.DE = 0x32EA
  *c.R1.HL = 0x4F7E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFA56
  *c.R1.IY = 0x074E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x5D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE945) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE945, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDBAE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDBAE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x324E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x324E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4F7E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4F7E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFA56) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFA56, *c.R1.IX)
  }
  if (*c.R1.IY != 0x074E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x074E, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD5E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6F3B
  *c.R1.BC = 0xE9DC
  *c.R1.DE = 0x7A06
  *c.R1.HL = 0x14F3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEC76
  *c.R1.IY = 0x8AAA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x5E
  mem[0x0002] = 0xC6
  mem[0x8A70] = 0x8C
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6F3B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6F3B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE9DC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE9DC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7A8C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7A8C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x14F3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x14F3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEC76) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEC76, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8AAA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8AAA, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD60(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8579
  *c.R1.BC = 0x005D
  *c.R1.DE = 0xD9EE
  *c.R1.HL = 0xFAEE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x382D
  *c.R1.IY = 0x2F95
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x60
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8579) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8579, *c.R1.AF)
  }
  if (*c.R1.BC != 0x005D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x005D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD9EE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD9EE, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFAEE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFAEE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x382D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x382D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0095) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0095, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD61(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5682
  *c.R1.BC = 0xDBC3
  *c.R1.DE = 0xB495
  *c.R1.HL = 0x9799
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x85B2
  *c.R1.IY = 0x3C1E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x61
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5682) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5682, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDBC3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDBC3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB495) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB495, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9799) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9799, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x85B2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x85B2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC31E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC31E, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD62(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x906B
  *c.R1.BC = 0xF52E
  *c.R1.DE = 0xF3D8
  *c.R1.HL = 0x1E8C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDDBA
  *c.R1.IY = 0x9A02
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x62
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x906B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x906B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF52E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF52E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF3D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF3D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1E8C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1E8C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDDBA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDDBA, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF302) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF302, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD63(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9D59
  *c.R1.BC = 0xBEB9
  *c.R1.DE = 0xD826
  *c.R1.HL = 0x0EAA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4290
  *c.R1.IY = 0xA4B9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x63
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9D59) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9D59, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBEB9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBEB9, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD826) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD826, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0EAA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0EAA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4290) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4290, *c.R1.IX)
  }
  if (*c.R1.IY != 0x26B9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x26B9, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD64(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7B0E
  *c.R1.BC = 0xE394
  *c.R1.DE = 0x8A25
  *c.R1.HL = 0xCDDF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9784
  *c.R1.IY = 0x2116
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x64
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7B0E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7B0E, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE394) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE394, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8A25) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8A25, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCDDF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCDDF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9784) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9784, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2116) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2116, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD65(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB827
  *c.R1.BC = 0xEB4F
  *c.R1.DE = 0xF666
  *c.R1.HL = 0xC52A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6206
  *c.R1.IY = 0x831F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x65
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB827) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB827, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEB4F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEB4F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF666) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF666, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC52A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC52A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6206) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6206, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1F1F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1F1F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD66(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9129
  *c.R1.BC = 0xE4EE
  *c.R1.DE = 0xE3A3
  *c.R1.HL = 0x86CA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4D93
  *c.R1.IY = 0x5B24
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x66
  mem[0x0002] = 0x80
  mem[0x5AA4] = 0x77
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9129) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9129, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE4EE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE4EE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE3A3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE3A3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x77CA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x77CA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4D93) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4D93, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5B24) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5B24, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD67(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDB7A
  *c.R1.BC = 0xB40B
  *c.R1.DE = 0x7B58
  *c.R1.HL = 0x49FD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x266F
  *c.R1.IY = 0x9E7B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x67
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDB7A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDB7A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB40B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB40B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7B58) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7B58, *c.R1.DE)
  }
  if (*c.R1.HL != 0x49FD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x49FD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x266F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x266F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDB7B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDB7B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD68(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4D1D
  *c.R1.BC = 0x4FD9
  *c.R1.DE = 0x783E
  *c.R1.HL = 0x0745
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0C3D
  *c.R1.IY = 0x82B5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x68
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4D1D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4D1D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4FD9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4FD9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x783E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x783E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0745) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0745, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0C3D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0C3D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x824F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x824F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD69(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1589
  *c.R1.BC = 0x5CEB
  *c.R1.DE = 0xB5DB
  *c.R1.HL = 0x922A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3C3A
  *c.R1.IY = 0xDC98
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x69
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1589) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1589, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5CEB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5CEB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB5DB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB5DB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x922A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x922A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3C3A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3C3A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDCEB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDCEB, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD6A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x607A
  *c.R1.BC = 0xE035
  *c.R1.DE = 0x5BB9
  *c.R1.HL = 0xDAC0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFC04
  *c.R1.IY = 0xB5B7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x6A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x607A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x607A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE035) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE035, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5BB9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5BB9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDAC0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDAC0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFC04) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFC04, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB55B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB55B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD6B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDB2A
  *c.R1.BC = 0xE244
  *c.R1.DE = 0x1182
  *c.R1.HL = 0x096F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x198E
  *c.R1.IY = 0x91A6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x6B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDB2A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDB2A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE244) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE244, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1182) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1182, *c.R1.DE)
  }
  if (*c.R1.HL != 0x096F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x096F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x198E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x198E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9182) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9182, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD6C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA0BE
  *c.R1.BC = 0x34EF
  *c.R1.DE = 0x8FCD
  *c.R1.HL = 0x40A7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4481
  *c.R1.IY = 0xC215
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x6C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA0BE) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA0BE, *c.R1.AF)
  }
  if (*c.R1.BC != 0x34EF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x34EF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8FCD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8FCD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x40A7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x40A7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4481) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4481, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC2C2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC2C2, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD6D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFDFC
  *c.R1.BC = 0x727A
  *c.R1.DE = 0xB839
  *c.R1.HL = 0x50A6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE782
  *c.R1.IY = 0x02E5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x6D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFDFC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFDFC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x727A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x727A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB839) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB839, *c.R1.DE)
  }
  if (*c.R1.HL != 0x50A6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x50A6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE782) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE782, *c.R1.IX)
  }
  if (*c.R1.IY != 0x02E5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x02E5, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD6E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCFD4
  *c.R1.BC = 0x6EF1
  *c.R1.DE = 0xC07D
  *c.R1.HL = 0xEB96
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB0F9
  *c.R1.IY = 0xB0A3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x6E
  mem[0x0002] = 0x78
  mem[0xB11B] = 0xF8
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCFD4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCFD4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6EF1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6EF1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC07D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC07D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEBF8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEBF8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB0F9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB0F9, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB0A3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB0A3, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD6F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8E1D
  *c.R1.BC = 0xA138
  *c.R1.DE = 0xF20A
  *c.R1.HL = 0x298E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB600
  *c.R1.IY = 0x0CF7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x6F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8E1D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8E1D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA138) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA138, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF20A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF20A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x298E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x298E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB600) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB600, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0C8E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0C8E, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD70(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2677
  *c.R1.BC = 0x33C5
  *c.R1.DE = 0xC0DC
  *c.R1.HL = 0x262F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD3DC
  *c.R1.IY = 0x23A1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x70
  mem[0x0002] = 0x53
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2677) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2677, *c.R1.AF)
  }
  if (*c.R1.BC != 0x33C5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x33C5, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC0DC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC0DC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x262F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x262F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD3DC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD3DC, *c.R1.IX)
  }
  if (*c.R1.IY != 0x23A1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x23A1, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD71(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x892E
  *c.R1.BC = 0x04AE
  *c.R1.DE = 0xD67F
  *c.R1.HL = 0x81EC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7757
  *c.R1.IY = 0xBFAB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x71
  mem[0x0002] = 0xB4
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x892E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x892E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x04AE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x04AE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD67F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD67F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x81EC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x81EC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7757) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7757, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBFAB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBFAB, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD72(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD2DC
  *c.R1.BC = 0xC23C
  *c.R1.DE = 0xDD54
  *c.R1.HL = 0x6559
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB32B
  *c.R1.IY = 0x7C80
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x72
  mem[0x0002] = 0xE3
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD2DC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD2DC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC23C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC23C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDD54) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDD54, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6559) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6559, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB32B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB32B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7C80) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7C80, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD73(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x49EF
  *c.R1.BC = 0xBFF2
  *c.R1.DE = 0x8409
  *c.R1.HL = 0x02DD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAF95
  *c.R1.IY = 0x8762
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x73
  mem[0x0002] = 0x17
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x49EF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x49EF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBFF2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBFF2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8409) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8409, *c.R1.DE)
  }
  if (*c.R1.HL != 0x02DD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x02DD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAF95) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAF95, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8762) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8762, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD74(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9479
  *c.R1.BC = 0x9817
  *c.R1.DE = 0xFA2E
  *c.R1.HL = 0x1FE0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA395
  *c.R1.IY = 0x92DB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x74
  mem[0x0002] = 0xF6
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9479) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9479, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9817) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9817, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFA2E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFA2E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1FE0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1FE0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA395) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA395, *c.R1.IX)
  }
  if (*c.R1.IY != 0x92DB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x92DB, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD75(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC8D6
  *c.R1.BC = 0x6AA4
  *c.R1.DE = 0x180E
  *c.R1.HL = 0xE37B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x02CF
  *c.R1.IY = 0x1724
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x75
  mem[0x0002] = 0xAB
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC8D6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC8D6, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6AA4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6AA4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x180E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x180E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE37B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE37B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x02CF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x02CF, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1724) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1724, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD77(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6F9E
  *c.R1.BC = 0x7475
  *c.R1.DE = 0x78AD
  *c.R1.HL = 0x2B8C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC6B7
  *c.R1.IY = 0x6B4D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x77
  mem[0x0002] = 0xF7
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6F9E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6F9E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7475) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7475, *c.R1.BC)
  }
  if (*c.R1.DE != 0x78AD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x78AD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2B8C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2B8C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC6B7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC6B7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6B4D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6B4D, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD7C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF228
  *c.R1.BC = 0x93FC
  *c.R1.DE = 0xA3D4
  *c.R1.HL = 0xDC9E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x21AC
  *c.R1.IY = 0xC617
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x7C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC628) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC628, *c.R1.AF)
  }
  if (*c.R1.BC != 0x93FC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x93FC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA3D4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA3D4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDC9E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDC9E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x21AC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x21AC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC617) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC617, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD7D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x93E5
  *c.R1.BC = 0x3CBE
  *c.R1.DE = 0x02C3
  *c.R1.HL = 0x26C2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCA81
  *c.R1.IY = 0x92B9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x7D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB9E5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB9E5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3CBE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3CBE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x02C3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x02C3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x26C2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x26C2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCA81) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCA81, *c.R1.IX)
  }
  if (*c.R1.IY != 0x92B9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x92B9, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD7E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1596
  *c.R1.BC = 0xDABA
  *c.R1.DE = 0x147B
  *c.R1.HL = 0xF362
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7110
  *c.R1.IY = 0xD45F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x7E
  mem[0x0002] = 0xE4
  mem[0xD443] = 0xAA
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAA96) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAA96, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDABA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDABA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x147B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x147B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF362) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF362, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7110) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7110, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD45F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD45F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD84(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBFBA
  *c.R1.BC = 0x7CAE
  *c.R1.DE = 0xC4DA
  *c.R1.HL = 0x7AEE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x43EE
  *c.R1.IY = 0xC08E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x84
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7F2D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7F2D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7CAE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7CAE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC4DA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC4DA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7AEE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7AEE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x43EE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x43EE, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC08E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC08E, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD85(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x52DD
  *c.R1.BC = 0x1DEA
  *c.R1.DE = 0x324F
  *c.R1.HL = 0x84E7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE7A8
  *c.R1.IY = 0xF799
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x85
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEBA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEBA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1DEA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1DEA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x324F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x324F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x84E7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x84E7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE7A8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE7A8, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF799) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF799, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD86(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFC9C
  *c.R1.BC = 0xB882
  *c.R1.DE = 0x43F9
  *c.R1.HL = 0x3E15
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9781
  *c.R1.IY = 0x8B33
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x86
  mem[0x0002] = 0xCE
  mem[0x8B01] = 0xE1
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDD89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDD89, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB882) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB882, *c.R1.BC)
  }
  if (*c.R1.DE != 0x43F9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x43F9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3E15) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3E15, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9781) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9781, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8B33) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8B33, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD8C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFD9C
  *c.R1.BC = 0x42B1
  *c.R1.DE = 0x5E8A
  *c.R1.HL = 0x081C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCB58
  *c.R1.IY = 0x3B4E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x8C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3839) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3839, *c.R1.AF)
  }
  if (*c.R1.BC != 0x42B1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x42B1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5E8A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5E8A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x081C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x081C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCB58) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCB58, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3B4E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3B4E, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD8D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9301
  *c.R1.BC = 0x7750
  *c.R1.DE = 0x8AD6
  *c.R1.HL = 0x295C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x695C
  *c.R1.IY = 0x99FB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x8D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8F89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8F89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7750) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7750, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8AD6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8AD6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x295C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x295C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x695C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x695C, *c.R1.IX)
  }
  if (*c.R1.IY != 0x99FB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x99FB, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD8E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x41EE
  *c.R1.BC = 0x398F
  *c.R1.DE = 0xF6DC
  *c.R1.HL = 0x06F3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF34A
  *c.R1.IY = 0x1AA2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x8E
  mem[0x0002] = 0x78
  mem[0x1B1A] = 0xC0
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0101) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0101, *c.R1.AF)
  }
  if (*c.R1.BC != 0x398F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x398F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF6DC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF6DC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x06F3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x06F3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF34A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF34A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1AA2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1AA2, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD94(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0431
  *c.R1.BC = 0xD255
  *c.R1.DE = 0xB9D6
  *c.R1.HL = 0x20BB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1E6A
  *c.R1.IY = 0xD5EF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x94
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2F3B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2F3B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD255) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD255, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB9D6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB9D6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x20BB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x20BB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1E6A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1E6A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD5EF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD5EF, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD95(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8B5D
  *c.R1.BC = 0xB455
  *c.R1.DE = 0x2388
  *c.R1.HL = 0xEC1E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7637
  *c.R1.IY = 0xCB97
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x95
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF4A3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF4A3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB455) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB455, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2388) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2388, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEC1E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEC1E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7637) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7637, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCB97) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCB97, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD96(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA0C6
  *c.R1.BC = 0x22AC
  *c.R1.DE = 0x0413
  *c.R1.HL = 0x4B13
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB44E
  *c.R1.IY = 0xC08B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x96
  mem[0x0002] = 0x55
  mem[0xC0E0] = 0x7B
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2536) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2536, *c.R1.AF)
  }
  if (*c.R1.BC != 0x22AC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x22AC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0413) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0413, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4B13) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4B13, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB44E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB44E, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC08B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC08B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFD9C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA44A
  *c.R1.BC = 0x3ECF
  *c.R1.DE = 0xCED3
  *c.R1.HL = 0x66EC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4BFF
  *c.R1.IY = 0xB133
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x9C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF3A3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF3A3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3ECF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3ECF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCED3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCED3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x66EC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x66EC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4BFF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4BFF, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB133) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB133, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD9D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x06C0
  *c.R1.BC = 0x8BD0
  *c.R1.DE = 0x131B
  *c.R1.HL = 0x3094
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAFC3
  *c.R1.IY = 0x7409
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x9D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFDBB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFDBB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8BD0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8BD0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x131B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x131B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3094) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3094, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAFC3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAFC3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7409) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7409, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFD9E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB983
  *c.R1.BC = 0x981F
  *c.R1.DE = 0xBB8E
  *c.R1.HL = 0xD6D5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5C3B
  *c.R1.IY = 0xF66C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0x9E
  mem[0x0002] = 0xF9
  mem[0xF665] = 0xF3
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC583) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC583, *c.R1.AF)
  }
  if (*c.R1.BC != 0x981F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x981F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBB8E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBB8E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD6D5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD6D5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5C3B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5C3B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF66C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF66C, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFDA4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB079
  *c.R1.BC = 0x79C0
  *c.R1.DE = 0x2C7C
  *c.R1.HL = 0x3E06
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7399
  *c.R1.IY = 0x037A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xA4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0054) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0054, *c.R1.AF)
  }
  if (*c.R1.BC != 0x79C0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x79C0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2C7C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2C7C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3E06) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3E06, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7399) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7399, *c.R1.IX)
  }
  if (*c.R1.IY != 0x037A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x037A, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDA5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x01D2
  *c.R1.BC = 0x654D
  *c.R1.DE = 0x9653
  *c.R1.HL = 0x2B33
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x61A4
  *c.R1.IY = 0x8F88
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xA5
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0054) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0054, *c.R1.AF)
  }
  if (*c.R1.BC != 0x654D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x654D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9653) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9653, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2B33) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2B33, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x61A4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x61A4, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8F88) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8F88, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDA6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDDB8
  *c.R1.BC = 0x40BB
  *c.R1.DE = 0x3742
  *c.R1.HL = 0x6FF1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAD28
  *c.R1.IY = 0x659B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xA6
  mem[0x0002] = 0x53
  mem[0x65EE] = 0x95
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9594) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9594, *c.R1.AF)
  }
  if (*c.R1.BC != 0x40BB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x40BB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3742) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3742, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6FF1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6FF1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAD28) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAD28, *c.R1.IX)
  }
  if (*c.R1.IY != 0x659B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x659B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFDAC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7A43
  *c.R1.BC = 0x72E3
  *c.R1.DE = 0xDD4D
  *c.R1.HL = 0x1B62
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4753
  *c.R1.IY = 0x5D63
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xAC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2724) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2724, *c.R1.AF)
  }
  if (*c.R1.BC != 0x72E3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x72E3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDD4D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDD4D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1B62) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1B62, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4753) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4753, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5D63) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5D63, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDAD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7D8E
  *c.R1.BC = 0x2573
  *c.R1.DE = 0x19CC
  *c.R1.HL = 0x78FB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5248
  *c.R1.IY = 0x8391
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xAD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xECA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xECA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2573) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2573, *c.R1.BC)
  }
  if (*c.R1.DE != 0x19CC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x19CC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x78FB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x78FB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5248) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5248, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8391) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8391, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDAE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA0DA
  *c.R1.BC = 0xBC27
  *c.R1.DE = 0x257B
  *c.R1.HL = 0x5489
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFA59
  *c.R1.IY = 0x81F8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xAE
  mem[0x0002] = 0x09
  mem[0x8201] = 0xCB
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6B28) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6B28, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBC27) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBC27, *c.R1.BC)
  }
  if (*c.R1.DE != 0x257B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x257B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5489) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5489, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFA59) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFA59, *c.R1.IX)
  }
  if (*c.R1.IY != 0x81F8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x81F8, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFDB4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4F95
  *c.R1.BC = 0x3461
  *c.R1.DE = 0xF173
  *c.R1.HL = 0x8AD3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC1A2
  *c.R1.IY = 0x8265
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xB4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCF8C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCF8C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3461) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3461, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF173) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF173, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8AD3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8AD3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC1A2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC1A2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8265) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8265, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDB5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x17F6
  *c.R1.BC = 0xE6EA
  *c.R1.DE = 0xF919
  *c.R1.HL = 0x327C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4299
  *c.R1.IY = 0x9733
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xB5
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3720) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3720, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE6EA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE6EA, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF919) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF919, *c.R1.DE)
  }
  if (*c.R1.HL != 0x327C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x327C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4299) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4299, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9733) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9733, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDB6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDB37
  *c.R1.BC = 0x3509
  *c.R1.DE = 0xD6CA
  *c.R1.HL = 0xB16A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA099
  *c.R1.IY = 0xDF6D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xB6
  mem[0x0002] = 0x4B
  mem[0xDFB8] = 0x64
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFFAC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFFAC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3509) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3509, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD6CA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD6CA, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB16A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB16A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA099) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA099, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDF6D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDF6D, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFDBC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB4FC
  *c.R1.BC = 0x9302
  *c.R1.DE = 0xE35D
  *c.R1.HL = 0x31BC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5C12
  *c.R1.IY = 0x1C92
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xBC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB49A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB49A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9302) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9302, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE35D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE35D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x31BC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x31BC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5C12) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5C12, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1C92) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1C92, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDBD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x391C
  *c.R1.BC = 0x7B82
  *c.R1.DE = 0xDFEB
  *c.R1.HL = 0x03EE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBE7B
  *c.R1.IY = 0xB30F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xBD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x391A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x391A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7B82) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7B82, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDFEB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDFEB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x03EE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x03EE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBE7B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBE7B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB30F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB30F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDBE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0970
  *c.R1.BC = 0x0B31
  *c.R1.DE = 0xF4AD
  *c.R1.HL = 0x9D4C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB95A
  *c.R1.IY = 0xA96B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xBE
  mem[0x0002] = 0x6B
  mem[0xA9D6] = 0xC0
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0903) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0903, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0B31) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0B31, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF4AD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF4AD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9D4C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9D4C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB95A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB95A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA96B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA96B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestFDE1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x828E
  *c.R1.BC = 0x078B
  *c.R1.DE = 0x1E35
  *c.R1.HL = 0x8F1C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4827
  *c.R1.IY = 0xB742
  *c.R1.SP = 0x716E
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xE1
  mem[0x716E] = 0xD5
  mem[0x716F] = 0x92
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x828E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x828E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x078B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x078B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E35) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E35, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8F1C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8F1C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4827) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4827, *c.R1.IX)
  }
  if (*c.R1.IY != 0x92D5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x92D5, *c.R1.IY)
  }
  if (*c.R1.SP != 0x7170) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x7170, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 14) {
    t.Errorf("TStates mismatch: %d expected, got %d", 14, c.TStates)
  }
}

func TestFDE3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4298
  *c.R1.BC = 0xC805
  *c.R1.DE = 0x6030
  *c.R1.HL = 0x4292
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x473B
  *c.R1.IY = 0x9510
  *c.R1.SP = 0x1A38
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xE3
  mem[0x1A38] = 0xE0
  mem[0x1A39] = 0x0F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4298) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4298, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC805) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC805, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6030) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6030, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4292) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4292, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x473B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x473B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0FE0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0FE0, *c.R1.IY)
  }
  if (*c.R1.SP != 0x1A38) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x1A38, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 23) {
    t.Errorf("TStates mismatch: %d expected, got %d", 23, c.TStates)
  }
}

func TestFDE5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD139
  *c.R1.BC = 0xAA0D
  *c.R1.DE = 0xBF2B
  *c.R1.HL = 0x2A56
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE138
  *c.R1.IY = 0xD4DA
  *c.R1.SP = 0xA8E1
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xE5
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD139) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD139, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAA0D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAA0D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBF2B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBF2B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2A56) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2A56, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE138) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE138, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD4DA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD4DA, *c.R1.IY)
  }
  if (*c.R1.SP != 0xA8DF) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xA8DF, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 15) {
    t.Errorf("TStates mismatch: %d expected, got %d", 15, c.TStates)
  }
}

func TestFDE9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC14F
  *c.R1.BC = 0x2EB6
  *c.R1.DE = 0xEDF0
  *c.R1.HL = 0x27CF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x09EE
  *c.R1.IY = 0xA2A4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xE9
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC14F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC14F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2EB6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2EB6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEDF0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEDF0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x27CF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x27CF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x09EE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x09EE, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA2A4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA2A4, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestFDF9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC260
  *c.R1.BC = 0x992E
  *c.R1.DE = 0xD544
  *c.R1.HL = 0x67FB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBA5E
  *c.R1.IY = 0x3596
  *c.R1.SP = 0x353F
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xF9
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC260) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC260, *c.R1.AF)
  }
  if (*c.R1.BC != 0x992E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x992E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD544) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD544, *c.R1.DE)
  }
  if (*c.R1.HL != 0x67FB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x67FB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBA5E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBA5E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3596) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3596, *c.R1.IY)
  }
  if (*c.R1.SP != 0x3596) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x3596, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}