package z80

import "testing"


func TestED40(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x83F9
  *c.R1.BC = 0x296B
  *c.R1.DE = 0x7034
  *c.R1.HL = 0x1F2F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x40
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8329) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8329, *c.R1.AF)
  }
  if (*c.R1.BC != 0x296B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x296B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7034) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7034, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1F2F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1F2F, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED41(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x29A2
  *c.R1.BC = 0x0881
  *c.R1.DE = 0xD7DD
  *c.R1.HL = 0xFF4E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x41
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x29A2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x29A2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0881) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0881, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD7DD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD7DD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFF4E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFF4E, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED42(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCBD3
  *c.R1.BC = 0x1C8F
  *c.R1.DE = 0xD456
  *c.R1.HL = 0x315E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x42
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCB12) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCB12, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1C8F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1C8F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD456) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD456, *c.R1.DE)
  }
  if (*c.R1.HL != 0x14CE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x14CE, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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

func TestED43(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDA36
  *c.R1.BC = 0x2732
  *c.R1.DE = 0x91CC
  *c.R1.HL = 0x9798
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5F73
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x43
  mem[0x0002] = 0xC6
  mem[0x0003] = 0x54
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDA36) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDA36, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2732) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2732, *c.R1.BC)
  }
  if (*c.R1.DE != 0x91CC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x91CC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9798) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9798, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x5F73) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5F73, *c.R1.SP)
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

func TestED44(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFE2B
  *c.R1.BC = 0x040F
  *c.R1.DE = 0xDEB6
  *c.R1.HL = 0xAFC3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5CA8
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x44
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0213) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0213, *c.R1.AF)
  }
  if (*c.R1.BC != 0x040F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x040F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDEB6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDEB6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAFC3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAFC3, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x5CA8) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5CA8, *c.R1.SP)
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

func TestED45(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x001D
  *c.R1.BC = 0x5B63
  *c.R1.DE = 0xA586
  *c.R1.HL = 0x1451
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x3100
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x45
  mem[0x3100] = 0x1F
  mem[0x3101] = 0x22
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x001D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x001D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5B63) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5B63, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA586) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA586, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1451) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1451, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x3102) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x3102, *c.R1.SP)
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

func TestED46(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB6EC
  *c.R1.BC = 0x8AFB
  *c.R1.DE = 0xCE09
  *c.R1.HL = 0x70A1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x8DEA
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x46
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB6EC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB6EC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8AFB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8AFB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCE09) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCE09, *c.R1.DE)
  }
  if (*c.R1.HL != 0x70A1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x70A1, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x8DEA) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x8DEA, *c.R1.SP)
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

func TestED47(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9A99
  *c.R1.BC = 0x9E5A
  *c.R1.DE = 0x9913
  *c.R1.HL = 0xCACC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x47
  for c.TStates < 9 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9A99) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9A99, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9E5A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9E5A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9913) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9913, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCACC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCACC, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x9A) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x9A, c.I)
  }
  if (c.TStates != 9) {
    t.Errorf("TStates mismatch: %d expected, got %d", 9, c.TStates)
  }
}

func TestED48(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDBDD
  *c.R1.BC = 0x7D1B
  *c.R1.DE = 0x141D
  *c.R1.HL = 0x5FB4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x48
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDB2D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDB2D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7D7D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7D7D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x141D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x141D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5FB4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5FB4, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED49(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x07A5
  *c.R1.BC = 0x59EC
  *c.R1.DE = 0xF459
  *c.R1.HL = 0x4316
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x49
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x07A5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x07A5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x59EC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x59EC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF459) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF459, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4316) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4316, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED4A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5741
  *c.R1.BC = 0x24B5
  *c.R1.DE = 0x83D2
  *c.R1.HL = 0x9AC8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x4A
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x57A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x57A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x24B5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x24B5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x83D2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x83D2, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBF7E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBF7E, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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

func TestED4B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x650C
  *c.R1.BC = 0xD74D
  *c.R1.DE = 0x0448
  *c.R1.HL = 0xA3B9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xB554
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x4B
  mem[0x0002] = 0x1A
  mem[0x0003] = 0xA4
  mem[0xA41A] = 0xF3
  mem[0xA41B] = 0xD4
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x650C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x650C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD4F3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD4F3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0448) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0448, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA3B9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA3B9, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xB554) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xB554, *c.R1.SP)
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

func TestED4C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5682
  *c.R1.BC = 0x7DDE
  *c.R1.DE = 0xB049
  *c.R1.HL = 0x939D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xC7BB
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x4C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAABB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAABB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7DDE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7DDE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB049) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB049, *c.R1.DE)
  }
  if (*c.R1.HL != 0x939D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x939D, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xC7BB) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xC7BB, *c.R1.SP)
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

func TestED4D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1BED
  *c.R1.BC = 0xC358
  *c.R1.DE = 0x5FD5
  *c.R1.HL = 0x6093
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x680E
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x4D
  mem[0x680E] = 0x03
  mem[0x680F] = 0x7C
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1BED) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1BED, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC358) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC358, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5FD5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5FD5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6093) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6093, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x6810) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x6810, *c.R1.SP)
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

func TestED4E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8E01
  *c.R1.BC = 0xE7C6
  *c.R1.DE = 0x880F
  *c.R1.HL = 0xD2A2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x85DA
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x4E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8E01) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8E01, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE7C6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE7C6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x880F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x880F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD2A2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD2A2, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x85DA) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x85DA, *c.R1.SP)
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

func TestED4F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2AE3
  *c.R1.BC = 0xC115
  *c.R1.DE = 0xEFF8
  *c.R1.HL = 0x9F6D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x4F
  for c.TStates < 9 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2AE3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2AE3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC115) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC115, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEFF8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEFF8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9F6D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9F6D, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x2A) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x2A, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 9) {
    t.Errorf("TStates mismatch: %d expected, got %d", 9, c.TStates)
  }
}

func TestED50(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x85AE
  *c.R1.BC = 0xBBCC
  *c.R1.DE = 0xE2A8
  *c.R1.HL = 0xF219
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x50
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x85AC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x85AC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBBCC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBBCC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBBA8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBBA8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF219) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF219, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED51(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2C4C
  *c.R1.BC = 0xC0A4
  *c.R1.DE = 0x5303
  *c.R1.HL = 0xBC25
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x51
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2C4C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2C4C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC0A4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC0A4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5303) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5303, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBC25) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBC25, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED52(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFC57
  *c.R1.BC = 0x1FC8
  *c.R1.DE = 0x47B6
  *c.R1.HL = 0xDA7C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x52
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFC82) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFC82, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1FC8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1FC8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x47B6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x47B6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x92C5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x92C5, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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

func TestED53(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1F88
  *c.R1.BC = 0x4692
  *c.R1.DE = 0x5CB2
  *c.R1.HL = 0x4915
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x7D8C
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x53
  mem[0x0002] = 0xFF
  mem[0x0003] = 0x21
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1F88) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1F88, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4692) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4692, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5CB2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5CB2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4915) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4915, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x7D8C) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x7D8C, *c.R1.SP)
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

func TestED54(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xADF9
  *c.R1.BC = 0x5661
  *c.R1.DE = 0x547C
  *c.R1.HL = 0xC322
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xD9EB
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x54
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5313) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5313, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5661) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5661, *c.R1.BC)
  }
  if (*c.R1.DE != 0x547C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x547C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC322) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC322, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xD9EB) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xD9EB, *c.R1.SP)
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

func TestED55(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB05B
  *c.R1.BC = 0x5E84
  *c.R1.DE = 0xD6E9
  *c.R1.HL = 0xCB3E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xD4B4
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x55
  mem[0xD4B4] = 0xEA
  mem[0xD4B5] = 0xC9
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB05B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB05B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5E84) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5E84, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD6E9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD6E9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCB3E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCB3E, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xD4B6) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xD4B6, *c.R1.SP)
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

func TestED56(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5CC0
  *c.R1.BC = 0x9100
  *c.R1.DE = 0x356B
  *c.R1.HL = 0x4BFD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x2C93
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x56
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5CC0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5CC0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9100) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9100, *c.R1.BC)
  }
  if (*c.R1.DE != 0x356B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x356B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4BFD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4BFD, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x2C93) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x2C93, *c.R1.SP)
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

func TestED57(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBCFE
  *c.R1.BC = 0xDFC7
  *c.R1.DE = 0xA621
  *c.R1.HL = 0x1022
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x1E
  c.R = 0x17
  mem[0x0000] = 0xED
  mem[0x0001] = 0x57
  for c.TStates < 9 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1E08) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1E08, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDFC7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDFC7, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA621) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA621, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1022) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1022, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x19) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x19, c.R)
  }
  if (c.I != 0x1E) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x1E, c.I)
  }
  if (c.TStates != 9) {
    t.Errorf("TStates mismatch: %d expected, got %d", 9, c.TStates)
  }
}

func TestED58(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC9EE
  *c.R1.BC = 0x4091
  *c.R1.DE = 0x9E46
  *c.R1.HL = 0x873A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x58
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC900, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4091) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4091, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9E40) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9E40, *c.R1.DE)
  }
  if (*c.R1.HL != 0x873A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x873A, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED59(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x388A
  *c.R1.BC = 0xD512
  *c.R1.DE = 0xECC5
  *c.R1.HL = 0x93AF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x59
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x388A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x388A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD512) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD512, *c.R1.BC)
  }
  if (*c.R1.DE != 0xECC5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xECC5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x93AF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x93AF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED5A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA41F
  *c.R1.BC = 0x751C
  *c.R1.DE = 0x19CE
  *c.R1.HL = 0x0493
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x5A
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA408) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA408, *c.R1.AF)
  }
  if (*c.R1.BC != 0x751C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x751C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x19CE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x19CE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1E62) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1E62, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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

func TestED5B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5DF1
  *c.R1.BC = 0x982E
  *c.R1.DE = 0x002F
  *c.R1.HL = 0xADB9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xF398
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x5B
  mem[0x0002] = 0x04
  mem[0x0003] = 0x9F
  mem[0x9F04] = 0x84
  mem[0x9F05] = 0x4D
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5DF1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5DF1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x982E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x982E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4D84) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4D84, *c.R1.DE)
  }
  if (*c.R1.HL != 0xADB9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xADB9, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xF398) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xF398, *c.R1.SP)
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

func TestED5C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x11C3
  *c.R1.BC = 0xB86C
  *c.R1.DE = 0x2042
  *c.R1.HL = 0xC958
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x93DC
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x5C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEFBB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEFBB, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB86C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB86C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2042) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2042, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC958) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC958, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x93DC) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x93DC, *c.R1.SP)
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

func TestED5D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1152
  *c.R1.BC = 0x1D20
  *c.R1.DE = 0x3F86
  *c.R1.HL = 0x64FC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5308
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x5D
  mem[0x5308] = 0x26
  mem[0x5309] = 0xE0
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1152) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1152, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1D20) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1D20, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3F86) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3F86, *c.R1.DE)
  }
  if (*c.R1.HL != 0x64FC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x64FC, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x530A) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x530A, *c.R1.SP)
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

func TestED5E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x611A
  *c.R1.BC = 0xC8CF
  *c.R1.DE = 0xF215
  *c.R1.HL = 0xD92B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x4D86
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x5E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x611A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x611A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC8CF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC8CF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF215) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF215, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD92B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD92B, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x4D86) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x4D86, *c.R1.SP)
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

func TestED5F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1BB5
  *c.R1.BC = 0xFC09
  *c.R1.DE = 0x2DFA
  *c.R1.HL = 0xBAB9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0xD7
  c.R = 0xF3
  mem[0x0000] = 0xED
  mem[0x0001] = 0x5F
  for c.TStates < 9 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF5A1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF5A1, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFC09) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFC09, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2DFA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2DFA, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBAB9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBAB9, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0xF5) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0xF5, c.R)
  }
  if (c.I != 0xD7) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0xD7, c.I)
  }
  if (c.TStates != 9) {
    t.Errorf("TStates mismatch: %d expected, got %d", 9, c.TStates)
  }
}

func TestED60(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2C9C
  *c.R1.BC = 0x0DAE
  *c.R1.DE = 0x621E
  *c.R1.HL = 0x2F66
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x60
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2C08) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2C08, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0DAE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0DAE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x621E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x621E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0D66) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0D66, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED61(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFFA8
  *c.R1.BC = 0x90CA
  *c.R1.DE = 0x0340
  *c.R1.HL = 0xD847
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x61
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFFA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFFA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x90CA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x90CA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0340) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0340, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD847) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD847, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED62(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA60B
  *c.R1.BC = 0xD9AA
  *c.R1.DE = 0x6623
  *c.R1.HL = 0x0B1A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x62
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA6BB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA6BB, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD9AA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD9AA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6623) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6623, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFFFF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFFFF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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

func TestED63(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5222
  *c.R1.BC = 0x88F9
  *c.R1.DE = 0x9D9A
  *c.R1.HL = 0xE4D3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xA2F0
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x63
  mem[0x0002] = 0x67
  mem[0x0003] = 0x65
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5222) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5222, *c.R1.AF)
  }
  if (*c.R1.BC != 0x88F9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x88F9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9D9A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9D9A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE4D3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE4D3, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xA2F0) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xA2F0, *c.R1.SP)
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

func TestED64(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2127
  *c.R1.BC = 0xE425
  *c.R1.DE = 0x66AC
  *c.R1.HL = 0xB2A3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F2
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x64
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDF9B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDF9B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE425) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE425, *c.R1.BC)
  }
  if (*c.R1.DE != 0x66AC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x66AC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB2A3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB2A3, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x43F2) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F2, *c.R1.SP)
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

func TestED65(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x63D2
  *c.R1.BC = 0x1FA1
  *c.R1.DE = 0x0788
  *c.R1.HL = 0x881C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xF207
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x65
  mem[0xF207] = 0xEB
  mem[0xF208] = 0x0E
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x63D2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x63D2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1FA1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1FA1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0788) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0788, *c.R1.DE)
  }
  if (*c.R1.HL != 0x881C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x881C, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xF209) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xF209, *c.R1.SP)
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

func TestED66(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4088
  *c.R1.BC = 0xA7E1
  *c.R1.DE = 0x3FFD
  *c.R1.HL = 0x919B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xD193
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x66
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4088) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4088, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA7E1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA7E1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3FFD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3FFD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x919B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x919B, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xD193) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xD193, *c.R1.SP)
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

func TestED67(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3624
  *c.R1.BC = 0xB16A
  *c.R1.DE = 0xA4DB
  *c.R1.HL = 0xB9DE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x67
  mem[0xB9DE] = 0x93
  for c.TStates < 18 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3324) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3324, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB16A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB16A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA4DB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA4DB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB9DE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB9DE, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 18) {
    t.Errorf("TStates mismatch: %d expected, got %d", 18, c.TStates)
  }
}

func TestED68(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5316
  *c.R1.BC = 0x624B
  *c.R1.DE = 0x7311
  *c.R1.HL = 0x3106
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x68
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5320) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5320, *c.R1.AF)
  }
  if (*c.R1.BC != 0x624B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x624B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7311) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7311, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3162) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3162, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED69(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xABD8
  *c.R1.BC = 0x8D2F
  *c.R1.DE = 0x89C7
  *c.R1.HL = 0xC3D6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x69
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xABD8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xABD8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8D2F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8D2F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x89C7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x89C7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC3D6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC3D6, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED6A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBB5A
  *c.R1.BC = 0x6FED
  *c.R1.DE = 0x59BB
  *c.R1.HL = 0x4E40
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x6A
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBB9C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBB9C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6FED) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6FED, *c.R1.BC)
  }
  if (*c.R1.DE != 0x59BB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x59BB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9C80) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9C80, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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

func TestED6B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9E35
  *c.R1.BC = 0xD240
  *c.R1.DE = 0x1998
  *c.R1.HL = 0xAB19
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x9275
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x6B
  mem[0x0002] = 0x98
  mem[0x0003] = 0x61
  mem[0x6198] = 0x3F
  mem[0x6199] = 0xBE
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9E35) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9E35, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD240) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD240, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1998) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1998, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBE3F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBE3F, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x9275) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x9275, *c.R1.SP)
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

func TestED6C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0FB1
  *c.R1.BC = 0x7D5B
  *c.R1.DE = 0xCADB
  *c.R1.HL = 0x0893
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xD983
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x6C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF1B3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF1B3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7D5B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7D5B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCADB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCADB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0893) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0893, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xD983) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xD983, *c.R1.SP)
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

func TestED6D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3860
  *c.R1.BC = 0x42DA
  *c.R1.DE = 0x5935
  *c.R1.HL = 0xDC10
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5CD3
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x6D
  mem[0x5CD3] = 0xA9
  mem[0x5CD4] = 0x73
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3860) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3860, *c.R1.AF)
  }
  if (*c.R1.BC != 0x42DA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x42DA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5935) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5935, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDC10) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDC10, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x5CD5) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5CD5, *c.R1.SP)
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

func TestED6E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7752
  *c.R1.BC = 0xBEC3
  *c.R1.DE = 0x0457
  *c.R1.HL = 0x8C95
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xA787
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x6E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7752) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7752, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBEC3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBEC3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0457) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0457, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8C95) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8C95, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xA787) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xA787, *c.R1.SP)
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

func TestED6F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x658B
  *c.R1.BC = 0x7A7A
  *c.R1.DE = 0xECF0
  *c.R1.HL = 0x403C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x6F
  mem[0x403C] = 0xC4
  for c.TStates < 18 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6C2D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6C2D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7A7A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7A7A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xECF0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xECF0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x403C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x403C, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 18) {
    t.Errorf("TStates mismatch: %d expected, got %d", 18, c.TStates)
  }
}

func TestED70(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC6A1
  *c.R1.BC = 0xF7D6
  *c.R1.DE = 0xA3CB
  *c.R1.HL = 0x288D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x70
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC6A1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC6A1, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF7D6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF7D6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA3CB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA3CB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x288D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x288D, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED71(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAFA0
  *c.R1.BC = 0x20B3
  *c.R1.DE = 0x7B33
  *c.R1.HL = 0x4AC1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x71
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAFA0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAFA0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x20B3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x20B3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7B33) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7B33, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4AC1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4AC1, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED72(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5FD9
  *c.R1.BC = 0x05CB
  *c.R1.DE = 0x0C6C
  *c.R1.HL = 0xD18B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x53DB
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x72
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5F3E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5F3E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x05CB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x05CB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0C6C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0C6C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7DAF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7DAF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x53DB) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x53DB, *c.R1.SP)
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

func TestED73(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x41C4
  *c.R1.BC = 0x763A
  *c.R1.DE = 0xECB0
  *c.R1.HL = 0xEE62
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xAED5
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x73
  mem[0x0002] = 0x2A
  mem[0x0003] = 0x79
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x41C4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x41C4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x763A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x763A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xECB0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xECB0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEE62) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEE62, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xAED5) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xAED5, *c.R1.SP)
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

func TestED74(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4454
  *c.R1.BC = 0xF2D2
  *c.R1.DE = 0x8340
  *c.R1.HL = 0x7E76
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0323
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x74
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBCBB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBCBB, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF2D2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF2D2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8340) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8340, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7E76) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7E76, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0323) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0323, *c.R1.SP)
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

func TestED75(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7CA4
  *c.R1.BC = 0x1615
  *c.R1.DE = 0x5D2A
  *c.R1.HL = 0xA95B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x7D00
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x75
  mem[0x7D00] = 0xFD
  mem[0x7D01] = 0x4F
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7CA4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7CA4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1615) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1615, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5D2A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5D2A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA95B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA95B, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x7D02) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x7D02, *c.R1.SP)
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

func TestED76(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCABF
  *c.R1.BC = 0xFF9A
  *c.R1.DE = 0xB98C
  *c.R1.HL = 0xA8E6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xFE8E
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x76
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCABF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCABF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFF9A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFF9A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB98C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB98C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA8E6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA8E6, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xFE8E) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xFE8E, *c.R1.SP)
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

func TestED78(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x58DD
  *c.R1.BC = 0xF206
  *c.R1.DE = 0x2D6A
  *c.R1.HL = 0xAF16
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x78
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF2A1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF2A1, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF206) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF206, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2D6A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2D6A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAF16) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAF16, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED79(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE000
  *c.R1.BC = 0x4243
  *c.R1.DE = 0x8F7F
  *c.R1.HL = 0xED90
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x79
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4243) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4243, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8F7F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8F7F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xED90) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xED90, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestED7A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x32FD
  *c.R1.BC = 0xD819
  *c.R1.DE = 0xD873
  *c.R1.HL = 0x8DCF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5D22
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x7A
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x32B8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x32B8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD819) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD819, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD873) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD873, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEAF2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEAF2, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x5D22) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5D22, *c.R1.SP)
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

func TestED7B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4F97
  *c.R1.BC = 0x24B7
  *c.R1.DE = 0xE105
  *c.R1.HL = 0x1BF2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5E17
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x7B
  mem[0x0002] = 0x50
  mem[0x0003] = 0x8C
  mem[0x8C50] = 0xD8
  mem[0x8C51] = 0x48
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4F97) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4F97, *c.R1.AF)
  }
  if (*c.R1.BC != 0x24B7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x24B7, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE105) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE105, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1BF2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1BF2, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x48D8) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x48D8, *c.R1.SP)
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

func TestED7C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD333
  *c.R1.BC = 0x29CA
  *c.R1.DE = 0x9622
  *c.R1.HL = 0xB452
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0BE6
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x7C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2D3B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2D3B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x29CA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x29CA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9622) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9622, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB452) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB452, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0BE6) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0BE6, *c.R1.SP)
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

func TestED7D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xECB6
  *c.R1.BC = 0x073E
  *c.R1.DE = 0xDC1E
  *c.R1.HL = 0x38D9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x66F0
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x7D
  mem[0x66F0] = 0x4F
  mem[0x66F1] = 0xFB
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0xECB6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xECB6, *c.R1.AF)
  }
  if (*c.R1.BC != 0x073E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x073E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDC1E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDC1E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x38D9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x38D9, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x66F2) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x66F2, *c.R1.SP)
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

func TestED7E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB246
  *c.R1.BC = 0x1A1A
  *c.R1.DE = 0x933A
  *c.R1.HL = 0x4B8B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x2242
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0x7E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB246) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB246, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1A1A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1A1A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x933A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x933A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4B8B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4B8B, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x2242) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x2242, *c.R1.SP)
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

func TestEDA0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1BC9
  *c.R1.BC = 0x3D11
  *c.R1.DE = 0x95C1
  *c.R1.HL = 0xD097
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA0
  mem[0xD097] = 0xB7
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1BE5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1BE5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3D10) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3D10, *c.R1.BC)
  }
  if (*c.R1.DE != 0x95C2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x95C2, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD098) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD098, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xECDB
  *c.R1.BC = 0x7666
  *c.R1.DE = 0x537F
  *c.R1.HL = 0x3BC3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA1
  mem[0x3BC3] = 0xB4
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEC0F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEC0F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7665) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7665, *c.R1.BC)
  }
  if (*c.R1.DE != 0x537F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x537F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3BC4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3BC4, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0121
  *c.R1.BC = 0x9A82
  *c.R1.DE = 0x5BBD
  *c.R1.HL = 0x2666
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA2
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x019F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x019F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9982) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9982, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5BBD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5BBD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2667) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2667, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA2_01(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0200
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x8000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA2
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0100) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0100, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8001) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8001, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA2_02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x569A
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x8000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA2
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x559A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x559A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8001) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8001, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA2_03(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0xABCC
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x8000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA2
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00BF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00BF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAACC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAACC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8001) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8001, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x42C5
  *c.R1.BC = 0x6334
  *c.R1.DE = 0x1E28
  *c.R1.HL = 0x32FA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x32FA] = 0xB3
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4233) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4233, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6234) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6234, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E28) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E28, *c.R1.DE)
  }
  if (*c.R1.HL != 0x32FB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x32FB, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_01(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0100
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FF] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0044) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0044, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0200) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0200, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0100
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0100
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x0100] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0040) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0040, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0101) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0101, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_03(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0100
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0107
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x0107] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0044) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0044, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0108) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0108, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_04(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0100
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FF] = 0x80
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0046) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0046, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0200) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0200, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_05(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0100
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FD] = 0x12
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0055) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0055, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x01FE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x01FE, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_06(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0100
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FE] = 0x12
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0051) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0051, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x01FF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x01FF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_07(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0200
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FF] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0100) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0100, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0200) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0200, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_08(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0800
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FE] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0004) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0004, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0700) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0700, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x01FF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x01FF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_09(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x8100
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FF] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0080) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0080, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0200) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0200, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_10(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x8200
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FF] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0084) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0084, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8100) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8100, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0200) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0200, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA3_11(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0xA900
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x01FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA3
  mem[0x01FF] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA800) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA800, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0200) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0200, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2A8E
  *c.R1.BC = 0x1607
  *c.R1.DE = 0x5938
  *c.R1.HL = 0x12E8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA8
  mem[0x12E8] = 0xD8
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2AA4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2AA4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1606) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1606, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5937) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5937, *c.R1.DE)
  }
  if (*c.R1.HL != 0x12E7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x12E7, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDA9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1495
  *c.R1.BC = 0xFB42
  *c.R1.DE = 0x0466
  *c.R1.HL = 0x0DBE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xA9
  mem[0x0DBE] = 0x89
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x14BF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x14BF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFB41) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFB41, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0466) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0466, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0DBD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0DBD, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDAA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2042
  *c.R1.BC = 0xD791
  *c.R1.DE = 0xA912
  *c.R1.HL = 0xA533
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xAA
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2097) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2097, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD691) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD691, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA912) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA912, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA532) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA532, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDAA_01(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0101
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x8000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xAA
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0040) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0040, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0001) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0001, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7FFF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7FFF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDAA_02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x56AA
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x8000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xAA
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x55AA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x55AA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7FFF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7FFF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDAA_03(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0xABCC
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x8000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xAA
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00BF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00BF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAACC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAACC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7FFF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7FFF, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDAB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0037
  *c.R1.BC = 0xF334
  *c.R1.DE = 0xD3E1
  *c.R1.HL = 0x199F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xAB
  mem[0x199F] = 0x49
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00A4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00A4, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF234) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF234, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD3E1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD3E1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x199E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x199E, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDAB_01(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x5800
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x007A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xAB
  mem[0x007A] = 0x7F
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5700) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5700, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0079) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0079, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDAB_02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0xAB00
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x00F1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xAB
  mem[0x00F1] = 0xCD
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00BF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00BF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAA00) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAA00, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x00F0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x00F0, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
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
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func TestEDB0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1045
  *c.R1.BC = 0x0010
  *c.R1.DE = 0xAAD8
  *c.R1.HL = 0x558E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xB0
  mem[0x558E] = 0x53
  mem[0x558F] = 0x94
  mem[0x5590] = 0x30
  mem[0x5591] = 0x05
  mem[0x5592] = 0x44
  mem[0x5593] = 0x24
  mem[0x5594] = 0x22
  mem[0x5595] = 0xB9
  mem[0x5596] = 0xE9
  mem[0x5597] = 0x77
  mem[0x5598] = 0x23
  mem[0x5599] = 0x71
  mem[0x559A] = 0xE2
  mem[0x559B] = 0x5C
  mem[0x559C] = 0xFB
  mem[0x559D] = 0x49
  for c.TStates < 331 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1049) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1049, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAAE8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAAE8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x559E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x559E, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x20) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x20, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 331) {
    t.Errorf("TStates mismatch: %d expected, got %d", 331, c.TStates)
  }
}

func TestEDB1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF4DD
  *c.R1.BC = 0x0008
  *c.R1.DE = 0xE4E0
  *c.R1.HL = 0x9825
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xB1
  mem[0x9825] = 0x50
  mem[0x9826] = 0xE5
  mem[0x9827] = 0x41
  mem[0x9828] = 0xF4
  mem[0x9829] = 0x01
  mem[0x982A] = 0x9F
  mem[0x982B] = 0x11
  mem[0x982C] = 0x85
  for c.TStates < 79 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF447) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF447, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0004) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0004, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE4E0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE4E0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9829) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9829, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x08) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x08, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 79) {
    t.Errorf("TStates mismatch: %d expected, got %d", 79, c.TStates)
  }
}

func TestEDB2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8A34
  *c.R1.BC = 0x0A40
  *c.R1.DE = 0xD98C
  *c.R1.HL = 0x37CE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xB2
  for c.TStates < 205 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8A40) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8A40, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0040) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0040, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD98C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD98C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x37D8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x37D8, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x14) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x14, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 205) {
    t.Errorf("TStates mismatch: %d expected, got %d", 205, c.TStates)
  }
}

func TestEDB3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x34AB
  *c.R1.BC = 0x03E0
  *c.R1.DE = 0x41B9
  *c.R1.HL = 0x1D7C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xB3
  mem[0x1D7C] = 0x9D
  mem[0x1D7D] = 0x24
  mem[0x1D7E] = 0xAA
  for c.TStates < 58 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3453) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3453, *c.R1.AF)
  }
  if (*c.R1.BC != 0x00E0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x00E0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x41B9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x41B9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1D7F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1D7F, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x06) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x06, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 58) {
    t.Errorf("TStates mismatch: %d expected, got %d", 58, c.TStates)
  }
}

func TestEDB8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE553
  *c.R1.BC = 0x0008
  *c.R1.DE = 0x68E8
  *c.R1.HL = 0x4DCF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xB8
  mem[0x4DC8] = 0x29
  mem[0x4DC9] = 0x85
  mem[0x4DCA] = 0xA7
  mem[0x4DCB] = 0xC3
  mem[0x4DCC] = 0x55
  mem[0x4DCD] = 0x74
  mem[0x4DCE] = 0x23
  mem[0x4DCF] = 0x0A
  for c.TStates < 163 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE569) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE569, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x68E0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x68E0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4DC7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4DC7, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x10) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x10, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 163) {
    t.Errorf("TStates mismatch: %d expected, got %d", 163, c.TStates)
  }
}

func TestEDB9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFFCD
  *c.R1.BC = 0x0008
  *c.R1.DE = 0xA171
  *c.R1.HL = 0xC749
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xB9
  mem[0xC742] = 0xC6
  mem[0xC743] = 0x09
  mem[0xC744] = 0x85
  mem[0xC745] = 0xEC
  mem[0xC746] = 0x5A
  mem[0xC747] = 0x01
  mem[0xC748] = 0x4E
  mem[0xC749] = 0x6C
  for c.TStates < 163 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF0B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF0B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA171) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA171, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC741) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC741, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x10) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x10, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 163) {
    t.Errorf("TStates mismatch: %d expected, got %d", 163, c.TStates)
  }
}

func TestEDBA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2567
  *c.R1.BC = 0x069F
  *c.R1.DE = 0xD40D
  *c.R1.HL = 0x6B55
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xBA
  for c.TStates < 121 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2540) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2540, *c.R1.AF)
  }
  if (*c.R1.BC != 0x009F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x009F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD40D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD40D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6B4F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6B4F, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x0C) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x0C, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 121) {
    t.Errorf("TStates mismatch: %d expected, got %d", 121, c.TStates)
  }
}

func TestEDBB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x09C4
  *c.R1.BC = 0x043B
  *c.R1.DE = 0xBE49
  *c.R1.HL = 0x1DD0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xED
  mem[0x0001] = 0xBB
  mem[0x1DCD] = 0xF9
  mem[0x1DCE] = 0x71
  mem[0x1DCF] = 0xC5
  mem[0x1DD0] = 0xB6
  for c.TStates < 79 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0957) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0957, *c.R1.AF)
  }
  if (*c.R1.BC != 0x003B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x003B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBE49) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBE49, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1DCC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1DCC, *c.R1.HL)
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
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x08) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x08, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 79) {
    t.Errorf("TStates mismatch: %d expected, got %d", 79, c.TStates)
  }
}