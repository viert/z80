package z80

import "testing"


func TestCB00(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDA00
  *c.R1.BC = 0xE479
  *c.R1.DE = 0x552E
  *c.R1.HL = 0xA806
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x00
  mem[0xA806] = 0x76
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDA8D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDA8D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC979) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC979, *c.R1.BC)
  }
  if (*c.R1.DE != 0x552E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x552E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA806) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA806, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB01(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1000
  *c.R1.BC = 0xB379
  *c.R1.DE = 0xB480
  *c.R1.HL = 0xEF65
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x01
  mem[0xEF65] = 0xFB
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x10A0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x10A0, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB3F2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB3F2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB480) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB480, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEF65) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEF65, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2E00
  *c.R1.BC = 0x9ADF
  *c.R1.DE = 0xAE6E
  *c.R1.HL = 0xA7F2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x02
  mem[0xA7F2] = 0x4A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2E09) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2E09, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9ADF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9ADF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5D6E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5D6E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA7F2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA7F2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB03(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6800
  *c.R1.BC = 0x9995
  *c.R1.DE = 0xDE3F
  *c.R1.HL = 0xCA71
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x03
  mem[0xCA71] = 0xE7
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x682C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x682C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9995) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9995, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDE7E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDE7E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCA71) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCA71, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB04(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8C00
  *c.R1.BC = 0xBEEA
  *c.R1.DE = 0x0CE4
  *c.R1.HL = 0x67B0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x04
  mem[0x67B0] = 0xCD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8C88) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8C88, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBEEA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBEEA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0CE4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0CE4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCEB0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCEB0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB05(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3600
  *c.R1.BC = 0xE19F
  *c.R1.DE = 0x78C9
  *c.R1.HL = 0xCB32
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x05
  mem[0xCB32] = 0x1B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3620) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3620, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE19F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE19F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x78C9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x78C9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCB64) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCB64, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB06(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8A00
  *c.R1.BC = 0xDB02
  *c.R1.DE = 0x8FB1
  *c.R1.HL = 0x5B04
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x06
  mem[0x5B04] = 0xD4
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8AAD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8AAD, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDB02) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDB02, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8FB1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8FB1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5B04) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5B04, *c.R1.HL)
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

func TestCB07(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6D00
  *c.R1.BC = 0x19CF
  *c.R1.DE = 0x7259
  *c.R1.HL = 0xDCAA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x07
  mem[0xDCAA] = 0x8D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDA88) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDA88, *c.R1.AF)
  }
  if (*c.R1.BC != 0x19CF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x19CF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7259) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7259, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCAA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCAA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB08(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8000
  *c.R1.BC = 0xCDB5
  *c.R1.DE = 0x818E
  *c.R1.HL = 0x2EE2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x08
  mem[0x2EE2] = 0x53
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x80A1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x80A1, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE6B5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE6B5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x818E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x818E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2EE2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2EE2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB09(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1800
  *c.R1.BC = 0x125C
  *c.R1.DE = 0xDD97
  *c.R1.HL = 0x59C6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x09
  mem[0x59C6] = 0x9E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x182C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x182C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x122E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x122E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDD97) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDD97, *c.R1.DE)
  }
  if (*c.R1.HL != 0x59C6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x59C6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB0A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1200
  *c.R1.BC = 0x3BA1
  *c.R1.DE = 0x7724
  *c.R1.HL = 0x63AD
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x0A
  mem[0x63AD] = 0x96
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x12AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x12AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3BA1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3BA1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBB24) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBB24, *c.R1.DE)
  }
  if (*c.R1.HL != 0x63AD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x63AD, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB0B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7600
  *c.R1.BC = 0x2ABF
  *c.R1.DE = 0xB626
  *c.R1.HL = 0x0289
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x0B
  mem[0x0289] = 0x37
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7600, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2ABF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2ABF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB613) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB613, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0289) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0289, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB0C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0E00
  *c.R1.BC = 0x6FC5
  *c.R1.DE = 0x2F12
  *c.R1.HL = 0x34D9
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x0C
  mem[0x34D9] = 0x50
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0E08) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0E08, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6FC5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6FC5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2F12) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2F12, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1AD9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1AD9, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB0D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6300
  *c.R1.BC = 0x95A3
  *c.R1.DE = 0xFCD2
  *c.R1.HL = 0x519A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x0D
  mem[0x519A] = 0x7A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x630C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x630C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x95A3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x95A3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFCD2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFCD2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x514D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x514D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB0E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFC00
  *c.R1.BC = 0xADF9
  *c.R1.DE = 0x4925
  *c.R1.HL = 0x543E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x0E
  mem[0x543E] = 0xD2
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFC2C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFC2C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xADF9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xADF9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4925) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4925, *c.R1.DE)
  }
  if (*c.R1.HL != 0x543E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x543E, *c.R1.HL)
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

func TestCB0F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC300
  *c.R1.BC = 0x18F3
  *c.R1.DE = 0x41B8
  *c.R1.HL = 0x070B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x0F
  mem[0x070B] = 0x86
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE1A5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE1A5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x18F3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x18F3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x41B8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x41B8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x070B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x070B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB10(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF800
  *c.R1.BC = 0xDC25
  *c.R1.DE = 0x33B3
  *c.R1.HL = 0x0D74
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x10
  mem[0x0D74] = 0x3D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF8AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF8AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB825) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB825, *c.R1.BC)
  }
  if (*c.R1.DE != 0x33B3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x33B3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0D74) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0D74, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB11(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6500
  *c.R1.BC = 0xE25C
  *c.R1.DE = 0x4B8A
  *c.R1.HL = 0xED42
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x11
  mem[0xED42] = 0xB7
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x65AC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x65AC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE2B8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE2B8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4B8A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4B8A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xED42) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xED42, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB12(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7700
  *c.R1.BC = 0x1384
  *c.R1.DE = 0x0F50
  *c.R1.HL = 0x29C6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x12
  mem[0x29C6] = 0x88
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x770C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x770C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1384) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1384, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E50) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E50, *c.R1.DE)
  }
  if (*c.R1.HL != 0x29C6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x29C6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB13(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCE00
  *c.R1.BC = 0x9F17
  *c.R1.DE = 0xE128
  *c.R1.HL = 0x3ED7
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x13
  mem[0x3ED7] = 0xEA
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCE04) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCE04, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9F17) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9F17, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE150) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE150, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3ED7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3ED7, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB14(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB200
  *c.R1.BC = 0x541A
  *c.R1.DE = 0x60C7
  *c.R1.HL = 0x7C9A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x14
  mem[0x7C9A] = 0x0F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB2A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB2A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x541A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x541A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x60C7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x60C7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF89A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF89A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB15(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2D00
  *c.R1.BC = 0xC1DF
  *c.R1.DE = 0x6EAB
  *c.R1.HL = 0x03E2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x15
  mem[0x03E2] = 0xBC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2D81) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2D81, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC1DF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC1DF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6EAB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6EAB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x03C4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x03C4, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB16(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3600
  *c.R1.BC = 0x3B53
  *c.R1.DE = 0x1A4A
  *c.R1.HL = 0x684E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x16
  mem[0x684E] = 0xC3
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3681) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3681, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3B53) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3B53, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1A4A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1A4A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x684E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x684E, *c.R1.HL)
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

func TestCB17(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5400
  *c.R1.BC = 0xD090
  *c.R1.DE = 0xF60D
  *c.R1.HL = 0x0FA2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x17
  mem[0x0FA2] = 0x23
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA8A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA8A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD090) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD090, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF60D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF60D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0FA2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0FA2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB18(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8600
  *c.R1.BC = 0xC658
  *c.R1.DE = 0x755F
  *c.R1.HL = 0x9596
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x18
  mem[0x9596] = 0xB6
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8624) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8624, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6358) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6358, *c.R1.BC)
  }
  if (*c.R1.DE != 0x755F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x755F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9596) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9596, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB19(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9600
  *c.R1.BC = 0xBEB3
  *c.R1.DE = 0x7C22
  *c.R1.HL = 0x71C8
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x19
  mem[0x71C8] = 0x85
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x960D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x960D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBE59) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBE59, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7C22) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7C22, *c.R1.DE)
  }
  if (*c.R1.HL != 0x71C8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x71C8, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB1A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3900
  *c.R1.BC = 0x882F
  *c.R1.DE = 0x543B
  *c.R1.HL = 0x5279
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x1A
  mem[0x5279] = 0x26
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3928) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3928, *c.R1.AF)
  }
  if (*c.R1.BC != 0x882F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x882F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2A3B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2A3B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5279) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5279, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB1B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9E00
  *c.R1.BC = 0xB338
  *c.R1.DE = 0x876C
  *c.R1.HL = 0xE8B4
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x1B
  mem[0xE8B4] = 0xB9
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9E24) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9E24, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB338) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB338, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8736) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8736, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE8B4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE8B4, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB1C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4B00
  *c.R1.BC = 0xB555
  *c.R1.DE = 0x238F
  *c.R1.HL = 0x311D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x1C
  mem[0x311D] = 0x11
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4B0D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4B0D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB555) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB555, *c.R1.BC)
  }
  if (*c.R1.DE != 0x238F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x238F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x181D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x181D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB1D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2100
  *c.R1.BC = 0x3D7E
  *c.R1.DE = 0x5E39
  *c.R1.HL = 0xE451
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x1D
  mem[0xE451] = 0x47
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x212D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x212D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3D7E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3D7E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5E39) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5E39, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE428) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE428, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB1E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5E00
  *c.R1.BC = 0x66B9
  *c.R1.DE = 0x80DC
  *c.R1.HL = 0x00EF
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x1E
  mem[0x00EF] = 0x91
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5E0D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5E0D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x66B9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x66B9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x80DC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x80DC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x00EF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x00EF, *c.R1.HL)
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

func TestCB1F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xED00
  *c.R1.BC = 0xB838
  *c.R1.DE = 0x8E18
  *c.R1.HL = 0xACE7
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x1F
  mem[0xACE7] = 0x82
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7621) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7621, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB838) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB838, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8E18) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8E18, *c.R1.DE)
  }
  if (*c.R1.HL != 0xACE7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xACE7, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB20(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC700
  *c.R1.BC = 0x0497
  *c.R1.DE = 0xD72B
  *c.R1.HL = 0xCCB6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x20
  mem[0xCCB6] = 0x1A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC708) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC708, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0897) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0897, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD72B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD72B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCCB6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCCB6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB21(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2200
  *c.R1.BC = 0x5CF4
  *c.R1.DE = 0x938E
  *c.R1.HL = 0x37A8
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x21
  mem[0x37A8] = 0xDD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x22AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x22AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5CE8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5CE8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x938E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x938E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x37A8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x37A8, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB22(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8500
  *c.R1.BC = 0x0950
  *c.R1.DE = 0xE7E8
  *c.R1.HL = 0x0641
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x22
  mem[0x0641] = 0x4D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8589) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8589, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0950) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0950, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCEE8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCEE8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0641) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0641, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB23(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2100
  *c.R1.BC = 0x2A7C
  *c.R1.DE = 0x37D0
  *c.R1.HL = 0xAA59
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x23
  mem[0xAA59] = 0xC1
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x21A5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x21A5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2A7C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2A7C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x37A0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x37A0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAA59) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAA59, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB24(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFB00
  *c.R1.BC = 0xB9DE
  *c.R1.DE = 0x7014
  *c.R1.HL = 0x84B6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x24
  mem[0x84B6] = 0x80
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFB09) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFB09, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB9DE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB9DE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7014) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7014, *c.R1.DE)
  }
  if (*c.R1.HL != 0x08B6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x08B6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB25(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1500
  *c.R1.BC = 0x6BBC
  *c.R1.DE = 0x894E
  *c.R1.HL = 0x85BC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x25
  mem[0x85BC] = 0xEF
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x152D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x152D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6BBC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6BBC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x894E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x894E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8578) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8578, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB26(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0A00
  *c.R1.BC = 0x372E
  *c.R1.DE = 0xE315
  *c.R1.HL = 0x283A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x26
  mem[0x283A] = 0xEE
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0A89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0A89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x372E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x372E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE315) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE315, *c.R1.DE)
  }
  if (*c.R1.HL != 0x283A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x283A, *c.R1.HL)
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

func TestCB27(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBF00
  *c.R1.BC = 0xBDBA
  *c.R1.DE = 0x67AB
  *c.R1.HL = 0x5EA2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x27
  mem[0x5EA2] = 0xBD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7E2D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7E2D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBDBA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBDBA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x67AB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x67AB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5EA2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5EA2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB28(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC000
  *c.R1.BC = 0x0435
  *c.R1.DE = 0x3E0F
  *c.R1.HL = 0x021B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x28
  mem[0x021B] = 0x90
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0235) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0235, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3E0F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3E0F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x021B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x021B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB29(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0600
  *c.R1.BC = 0xF142
  *c.R1.DE = 0x6ADA
  *c.R1.HL = 0xC306
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x29
  mem[0xC306] = 0x5C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0624) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0624, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF121) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF121, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6ADA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6ADA, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC306) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC306, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB2A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3000
  *c.R1.BC = 0xEC3A
  *c.R1.DE = 0x7F7D
  *c.R1.HL = 0x3473
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x2A
  mem[0x3473] = 0x34
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x302D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x302D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEC3A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEC3A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3F7D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3F7D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3473) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3473, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB2B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE000
  *c.R1.BC = 0xCCF0
  *c.R1.DE = 0xBBDA
  *c.R1.HL = 0xB78A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x2B
  mem[0xB78A] = 0xAB
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE0AC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE0AC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCCF0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCCF0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBBED) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBBED, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB78A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB78A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB2C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5B00
  *c.R1.BC = 0x25C0
  *c.R1.DE = 0x996D
  *c.R1.HL = 0x1E7B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x2C
  mem[0x1E7B] = 0x2C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5B0C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5B0C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x25C0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x25C0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x996D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x996D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0F7B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0F7B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB2D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5E00
  *c.R1.BC = 0xC51B
  *c.R1.DE = 0x58E3
  *c.R1.HL = 0x78EA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x2D
  mem[0x78EA] = 0x85
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5EA4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5EA4, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC51B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC51B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x58E3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x58E3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x78F5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x78F5, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB2E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3900
  *c.R1.BC = 0xA2CD
  *c.R1.DE = 0x0629
  *c.R1.HL = 0x24BF
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x2E
  mem[0x24BF] = 0xB5
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3989) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3989, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA2CD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA2CD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0629) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0629, *c.R1.DE)
  }
  if (*c.R1.HL != 0x24BF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x24BF, *c.R1.HL)
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

func TestCB2F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAA00
  *c.R1.BC = 0xA194
  *c.R1.DE = 0xD0E3
  *c.R1.HL = 0x5C65
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x2F
  mem[0x5C65] = 0xC9
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD580) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD580, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA194) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA194, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD0E3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD0E3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5C65) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5C65, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB30(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCD00
  *c.R1.BC = 0x7A81
  *c.R1.DE = 0xD67B
  *c.R1.HL = 0x656B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x30
  mem[0x656B] = 0x32
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCDA4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCDA4, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF581) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF581, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD67B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD67B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x656B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x656B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB31(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2800
  *c.R1.BC = 0xE7FA
  *c.R1.DE = 0x6D8C
  *c.R1.HL = 0x75A4
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x31
  mem[0x75A4] = 0x0C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x28A5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x28A5, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE7F5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE7F5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6D8C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6D8C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x75A4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x75A4, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB32(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1300
  *c.R1.BC = 0x3F36
  *c.R1.DE = 0xF608
  *c.R1.HL = 0x5E56
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x32
  mem[0x5E56] = 0x8D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x13AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x13AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3F36) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3F36, *c.R1.BC)
  }
  if (*c.R1.DE != 0xED08) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xED08, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E56) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E56, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB33(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD500
  *c.R1.BC = 0x9720
  *c.R1.DE = 0x7644
  *c.R1.HL = 0x038F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x33
  mem[0x038F] = 0xBA
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD588) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD588, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9720) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9720, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7689) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7689, *c.R1.DE)
  }
  if (*c.R1.HL != 0x038F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x038F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB34(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1200
  *c.R1.BC = 0x77F6
  *c.R1.DE = 0x0206
  *c.R1.HL = 0xFB38
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x34
  mem[0xFB38] = 0x07
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x12A1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x12A1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x77F6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x77F6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0206) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0206, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF738) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF738, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB35(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3C00
  *c.R1.BC = 0xFD68
  *c.R1.DE = 0xEA91
  *c.R1.HL = 0x7861
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x35
  mem[0x7861] = 0x72
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C84) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C84, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFD68) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFD68, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEA91) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEA91, *c.R1.DE)
  }
  if (*c.R1.HL != 0x78C3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x78C3, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB36(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8A00
  *c.R1.BC = 0x1185
  *c.R1.DE = 0x1DDE
  *c.R1.HL = 0x6D38
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x36
  mem[0x6D38] = 0xF1
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8AA1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8AA1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1185) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1185, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1DDE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1DDE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6D38) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6D38, *c.R1.HL)
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

func TestCB37(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4300
  *c.R1.BC = 0xD7BC
  *c.R1.DE = 0x9133
  *c.R1.HL = 0x6E56
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x37
  mem[0x6E56] = 0xF8
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8784) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8784, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD7BC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD7BC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9133) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9133, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6E56) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6E56, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB38(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDF00
  *c.R1.BC = 0x7C1B
  *c.R1.DE = 0x9F9F
  *c.R1.HL = 0x4FF2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x38
  mem[0x4FF2] = 0xAA
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDF28) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDF28, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3E1B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3E1B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9F9F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9F9F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4FF2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4FF2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB39(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6600
  *c.R1.BC = 0xB702
  *c.R1.DE = 0x14F5
  *c.R1.HL = 0x3C17
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x39
  mem[0x3C17] = 0x61
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6600, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB701) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB701, *c.R1.BC)
  }
  if (*c.R1.DE != 0x14F5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x14F5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3C17) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3C17, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB3A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD100
  *c.R1.BC = 0x5C5F
  *c.R1.DE = 0xE42E
  *c.R1.HL = 0xF1B1
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x3A
  mem[0xF1B1] = 0x6E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD124) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD124, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5C5F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5C5F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x722E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x722E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF1B1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF1B1, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB3B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB200
  *c.R1.BC = 0x38C8
  *c.R1.DE = 0xA560
  *c.R1.HL = 0x7419
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x3B
  mem[0x7419] = 0x11
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB224) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB224, *c.R1.AF)
  }
  if (*c.R1.BC != 0x38C8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x38C8, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA530) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA530, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7419) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7419, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB3C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7800
  *c.R1.BC = 0xCFAE
  *c.R1.DE = 0x66D8
  *c.R1.HL = 0x2AD8
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x3C
  mem[0x2AD8] = 0x8D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7800, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCFAE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCFAE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x66D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x66D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x15D8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x15D8, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB3D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE600
  *c.R1.BC = 0xDCDA
  *c.R1.DE = 0x06AA
  *c.R1.HL = 0x46CD
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x3D
  mem[0x46CD] = 0xF9
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE625) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE625, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDCDA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDCDA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x06AA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x06AA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4666) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4666, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB3E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA900
  *c.R1.BC = 0x6A34
  *c.R1.DE = 0xE8D0
  *c.R1.HL = 0xA96C
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x3E
  mem[0xA96C] = 0xA0
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA904) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA904, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6A34) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6A34, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE8D0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE8D0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA96C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA96C, *c.R1.HL)
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

func TestCB3F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF100
  *c.R1.BC = 0xCEEA
  *c.R1.DE = 0x721E
  *c.R1.HL = 0x77F0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x3F
  mem[0x77F0] = 0x7C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x782D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x782D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCEEA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCEEA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x721E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x721E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x77F0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x77F0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB40(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9E00
  *c.R1.BC = 0xBCB2
  *c.R1.DE = 0xEFAA
  *c.R1.HL = 0x505F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x40
  mem[0x505F] = 0x59
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9E7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9E7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBCB2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBCB2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEFAA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEFAA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x505F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x505F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB41(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9E00
  *c.R1.BC = 0x1B43
  *c.R1.DE = 0x954E
  *c.R1.HL = 0x7BE9
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x41
  mem[0x7BE9] = 0xF7
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9E10) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9E10, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1B43) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1B43, *c.R1.BC)
  }
  if (*c.R1.DE != 0x954E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x954E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7BE9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7BE9, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB42(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF200
  *c.R1.BC = 0xDD12
  *c.R1.DE = 0x7D4F
  *c.R1.HL = 0x551F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x42
  mem[0x551F] = 0xC9
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF238) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF238, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDD12) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDD12, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7D4F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7D4F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x551F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x551F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB43(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAD00
  *c.R1.BC = 0xC3B3
  *c.R1.DE = 0xF1D0
  *c.R1.HL = 0xBAB4
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x43
  mem[0xBAB4] = 0x76
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAD54) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAD54, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC3B3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC3B3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF1D0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF1D0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBAB4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBAB4, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB44(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB700
  *c.R1.BC = 0xC829
  *c.R1.DE = 0x27E3
  *c.R1.HL = 0x5B92
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x44
  mem[0x5B92] = 0x78
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB718) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB718, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC829) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC829, *c.R1.BC)
  }
  if (*c.R1.DE != 0x27E3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x27E3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5B92) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5B92, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB45(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7700
  *c.R1.BC = 0x68EE
  *c.R1.DE = 0x0C77
  *c.R1.HL = 0x409B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x45
  mem[0x409B] = 0x64
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7718) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7718, *c.R1.AF)
  }
  if (*c.R1.BC != 0x68EE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x68EE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0C77) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0C77, *c.R1.DE)
  }
  if (*c.R1.HL != 0x409B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x409B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB46(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7200
  *c.R1.BC = 0x7AE3
  *c.R1.DE = 0xA11E
  *c.R1.HL = 0x6131
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x46
  mem[0x6131] = 0xD5
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7210) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7210, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7AE3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7AE3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA11E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA11E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6131) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6131, *c.R1.HL)
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

func TestCB47_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x47
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB47(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1000
  *c.R1.BC = 0xD8CA
  *c.R1.DE = 0xE2C4
  *c.R1.HL = 0x8A8C
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x47
  mem[0x8A8C] = 0x0E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1054) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1054, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD8CA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD8CA, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE2C4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE2C4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8A8C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8A8C, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB48(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA900
  *c.R1.BC = 0x6264
  *c.R1.DE = 0xE833
  *c.R1.HL = 0x6DE0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x48
  mem[0x6DE0] = 0x8C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA930) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA930, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6264) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6264, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE833) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE833, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6DE0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6DE0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB49(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6C00
  *c.R1.BC = 0xD0F7
  *c.R1.DE = 0x1DB7
  *c.R1.HL = 0xA040
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x49
  mem[0xA040] = 0x5F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6C30) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6C30, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD0F7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD0F7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1DB7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1DB7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA040) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA040, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB4A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4F00
  *c.R1.BC = 0xF04C
  *c.R1.DE = 0x5B29
  *c.R1.HL = 0x77A4
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x4A
  mem[0x77A4] = 0x96
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4F18) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4F18, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF04C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF04C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5B29) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5B29, *c.R1.DE)
  }
  if (*c.R1.HL != 0x77A4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x77A4, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB4B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5500
  *c.R1.BC = 0x9848
  *c.R1.DE = 0x095F
  *c.R1.HL = 0x40CA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x4B
  mem[0x40CA] = 0x8A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5518) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5518, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9848) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9848, *c.R1.BC)
  }
  if (*c.R1.DE != 0x095F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x095F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x40CA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x40CA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB4C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8800
  *c.R1.BC = 0x0521
  *c.R1.DE = 0xBF31
  *c.R1.HL = 0x6D5D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x4C
  mem[0x6D5D] = 0xE7
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x887C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x887C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0521) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0521, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBF31) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBF31, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6D5D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6D5D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB4D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF900
  *c.R1.BC = 0x27D0
  *c.R1.DE = 0x0F7E
  *c.R1.HL = 0x158D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x4D
  mem[0x158D] = 0xE0
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF95C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF95C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x27D0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x27D0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0F7E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0F7E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x158D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x158D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB4E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2600
  *c.R1.BC = 0x9207
  *c.R1.DE = 0x459A
  *c.R1.HL = 0xADA3
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x4E
  mem[0xADA3] = 0x5B
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2618) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2618, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9207) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9207, *c.R1.BC)
  }
  if (*c.R1.DE != 0x459A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x459A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xADA3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xADA3, *c.R1.HL)
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

func TestCB4F_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x4F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB4F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1700
  *c.R1.BC = 0x2DC1
  *c.R1.DE = 0xACA2
  *c.R1.HL = 0x0BCC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x4F
  mem[0x0BCC] = 0xA3
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1710) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1710, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2DC1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2DC1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xACA2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xACA2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0BCC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0BCC, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB50(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2300
  *c.R1.BC = 0x2749
  *c.R1.DE = 0x1012
  *c.R1.HL = 0x84D2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x50
  mem[0x84D2] = 0x6A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2330) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2330, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2749) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2749, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1012) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1012, *c.R1.DE)
  }
  if (*c.R1.HL != 0x84D2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x84D2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB51(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2200
  *c.R1.BC = 0xB7DB
  *c.R1.DE = 0xE19D
  *c.R1.HL = 0xAAFC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x51
  mem[0xAAFC] = 0xA6
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x225C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x225C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB7DB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB7DB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE19D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE19D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAAFC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAAFC, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB52(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8B00
  *c.R1.BC = 0xFF7A
  *c.R1.DE = 0xB0FF
  *c.R1.HL = 0xAC44
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x52
  mem[0xAC44] = 0x00
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8B74) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8B74, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFF7A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFF7A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB0FF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB0FF, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAC44) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAC44, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB53(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6000
  *c.R1.BC = 0x31A1
  *c.R1.DE = 0xA4F4
  *c.R1.HL = 0x7C75
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x53
  mem[0x7C75] = 0xAB
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6030) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6030, *c.R1.AF)
  }
  if (*c.R1.BC != 0x31A1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x31A1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA4F4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA4F4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7C75) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7C75, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB54(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3800
  *c.R1.BC = 0x7CCC
  *c.R1.DE = 0x89CC
  *c.R1.HL = 0x1999
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x54
  mem[0x1999] = 0x98
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x385C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x385C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7CCC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7CCC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x89CC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x89CC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1999) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1999, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB55(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF900
  *c.R1.BC = 0x1F79
  *c.R1.DE = 0x19CD
  *c.R1.HL = 0xFB4B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x55
  mem[0xFB4B] = 0x0B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF95C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF95C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1F79) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1F79, *c.R1.BC)
  }
  if (*c.R1.DE != 0x19CD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x19CD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFB4B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFB4B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB56(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1500
  *c.R1.BC = 0x2BFE
  *c.R1.DE = 0xE3B5
  *c.R1.HL = 0xBBF9
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x56
  mem[0xBBF9] = 0x10
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1554) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1554, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2BFE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2BFE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE3B5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE3B5, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBBF9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBBF9, *c.R1.HL)
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

func TestCB57_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x57
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB57(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6600
  *c.R1.BC = 0xAF32
  *c.R1.DE = 0x532A
  *c.R1.HL = 0xDA50
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x57
  mem[0xDA50] = 0x30
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6630) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6630, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAF32) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAF32, *c.R1.BC)
  }
  if (*c.R1.DE != 0x532A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x532A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDA50) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDA50, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB58(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5000
  *c.R1.BC = 0x1AEE
  *c.R1.DE = 0x2E47
  *c.R1.HL = 0x1479
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x58
  mem[0x1479] = 0xA0
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5018) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5018, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1AEE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1AEE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2E47) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2E47, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1479) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1479, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB59(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7200
  *c.R1.BC = 0x5E68
  *c.R1.DE = 0xFF28
  *c.R1.HL = 0x2075
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x59
  mem[0x2075] = 0xC1
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7238) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7238, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5E68) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5E68, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFF28) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFF28, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2075) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2075, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB5A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEB00
  *c.R1.BC = 0xFEA7
  *c.R1.DE = 0x17D1
  *c.R1.HL = 0xD99B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x5A
  mem[0xD99B] = 0xE8
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEB54) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEB54, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFEA7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFEA7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x17D1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x17D1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD99B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD99B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB5B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6B00
  *c.R1.BC = 0x6F2C
  *c.R1.DE = 0x3FE3
  *c.R1.HL = 0x1691
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x5B
  mem[0x1691] = 0xC7
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6B74) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6B74, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6F2C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6F2C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3FE3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3FE3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1691) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1691, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB5C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3300
  *c.R1.BC = 0xA7E7
  *c.R1.DE = 0x2077
  *c.R1.HL = 0x13E9
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x5C
  mem[0x13E9] = 0xAE
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3354) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3354, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA7E7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA7E7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2077) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2077, *c.R1.DE)
  }
  if (*c.R1.HL != 0x13E9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x13E9, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB5D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC100
  *c.R1.BC = 0xAFCC
  *c.R1.DE = 0xC8B1
  *c.R1.HL = 0xEE49
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x5D
  mem[0xEE49] = 0xA6
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC118) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC118, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAFCC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAFCC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC8B1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC8B1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEE49) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEE49, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB5E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3000
  *c.R1.BC = 0xAD43
  *c.R1.DE = 0x16C1
  *c.R1.HL = 0x349A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x5E
  mem[0x349A] = 0x3C
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3038) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3038, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAD43) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAD43, *c.R1.BC)
  }
  if (*c.R1.DE != 0x16C1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x16C1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x349A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x349A, *c.R1.HL)
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

func TestCB5F_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x5F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB5F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8C00
  *c.R1.BC = 0x1B67
  *c.R1.DE = 0x2314
  *c.R1.HL = 0x6133
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x5F
  mem[0x6133] = 0x90
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8C18) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8C18, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1B67) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1B67, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2314) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2314, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6133) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6133, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB60(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9900
  *c.R1.BC = 0x34B5
  *c.R1.DE = 0x0FD8
  *c.R1.HL = 0x5273
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x60
  mem[0x5273] = 0x0A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9930) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9930, *c.R1.AF)
  }
  if (*c.R1.BC != 0x34B5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x34B5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0FD8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0FD8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5273) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5273, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB61(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD100
  *c.R1.BC = 0x219F
  *c.R1.DE = 0x3BB4
  *c.R1.HL = 0x7C44
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x61
  mem[0x7C44] = 0x77
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD118) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD118, *c.R1.AF)
  }
  if (*c.R1.BC != 0x219F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x219F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3BB4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3BB4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7C44) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7C44, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB62(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAF00
  *c.R1.BC = 0xBDF8
  *c.R1.DE = 0xC536
  *c.R1.HL = 0x8CC5
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x62
  mem[0x8CC5] = 0xAF
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAF54) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAF54, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBDF8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBDF8, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC536) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC536, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8CC5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8CC5, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB63(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2A00
  *c.R1.BC = 0x5E16
  *c.R1.DE = 0xF627
  *c.R1.HL = 0x84CA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x63
  mem[0x84CA] = 0xE6
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2A74) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2A74, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5E16) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5E16, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF627) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF627, *c.R1.DE)
  }
  if (*c.R1.HL != 0x84CA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x84CA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB64(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA900
  *c.R1.BC = 0xA365
  *c.R1.DE = 0xC00B
  *c.R1.HL = 0xEA94
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x64
  mem[0xEA94] = 0x0C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA97C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA97C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA365) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA365, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC00B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC00B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEA94) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEA94, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB65(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1800
  *c.R1.BC = 0x8D58
  *c.R1.DE = 0x4256
  *c.R1.HL = 0x427A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x65
  mem[0x427A] = 0xEE
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1838) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1838, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8D58) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8D58, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4256) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4256, *c.R1.DE)
  }
  if (*c.R1.HL != 0x427A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x427A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB66(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4C00
  *c.R1.BC = 0x3EF7
  *c.R1.DE = 0xE544
  *c.R1.HL = 0xA44F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x66
  mem[0xA44F] = 0xD2
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4C10) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4C10, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3EF7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3EF7, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE544) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE544, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA44F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA44F, *c.R1.HL)
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

func TestCB67_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x67
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB67(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8600
  *c.R1.BC = 0x5E92
  *c.R1.DE = 0x2986
  *c.R1.HL = 0x394D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x67
  mem[0x394D] = 0x10
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8654) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8654, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5E92) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5E92, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2986) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2986, *c.R1.DE)
  }
  if (*c.R1.HL != 0x394D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x394D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB68(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD700
  *c.R1.BC = 0x0F6A
  *c.R1.DE = 0x18A6
  *c.R1.HL = 0xDDD2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x68
  mem[0xDDD2] = 0x16
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD75C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD75C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F6A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F6A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x18A6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x18A6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDDD2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDDD2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB69(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDA00
  *c.R1.BC = 0x691B
  *c.R1.DE = 0x7C79
  *c.R1.HL = 0x1DBA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x69
  mem[0x1DBA] = 0x8A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDA5C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDA5C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x691B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x691B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7C79) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7C79, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1DBA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1DBA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB6A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2200
  *c.R1.BC = 0x13E8
  *c.R1.DE = 0x86D4
  *c.R1.HL = 0x4E09
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x6A
  mem[0x4E09] = 0xD5
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2254) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2254, *c.R1.AF)
  }
  if (*c.R1.BC != 0x13E8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x13E8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x86D4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x86D4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4E09) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4E09, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB6B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAF00
  *c.R1.BC = 0x5123
  *c.R1.DE = 0x7635
  *c.R1.HL = 0x1CA9
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x6B
  mem[0x1CA9] = 0x86
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAF30) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAF30, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5123) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5123, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7635) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7635, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1CA9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1CA9, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB6C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4300
  *c.R1.BC = 0xFAA6
  *c.R1.DE = 0xABC2
  *c.R1.HL = 0x5605
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x6C
  mem[0x5605] = 0x2B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4354) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4354, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFAA6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFAA6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xABC2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xABC2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5605) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5605, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB6D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7F00
  *c.R1.BC = 0xF099
  *c.R1.DE = 0xD435
  *c.R1.HL = 0xD9AD
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x6D
  mem[0xD9AD] = 0x4E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7F38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7F38, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF099) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF099, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD435) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD435, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD9AD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD9AD, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB6E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4A00
  *c.R1.BC = 0x08C9
  *c.R1.DE = 0x8177
  *c.R1.HL = 0xD8BA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x6E
  mem[0xD8BA] = 0x31
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4A30) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4A30, *c.R1.AF)
  }
  if (*c.R1.BC != 0x08C9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x08C9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8177) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8177, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD8BA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD8BA, *c.R1.HL)
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

func TestCB6F_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x6F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB6F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA100
  *c.R1.BC = 0x8C80
  *c.R1.DE = 0x4678
  *c.R1.HL = 0x4D34
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x6F
  mem[0x4D34] = 0x78
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA130) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA130, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8C80) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8C80, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4678) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4678, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4D34) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4D34, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB70(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1900
  *c.R1.BC = 0x958A
  *c.R1.DE = 0x5DAB
  *c.R1.HL = 0xF913
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x70
  mem[0xF913] = 0xCF
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1954) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1954, *c.R1.AF)
  }
  if (*c.R1.BC != 0x958A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x958A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5DAB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5DAB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF913) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF913, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB71(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3D00
  *c.R1.BC = 0x095E
  *c.R1.DE = 0xD6DF
  *c.R1.HL = 0x42FE
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x71
  mem[0x42FE] = 0x24
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3D18) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3D18, *c.R1.AF)
  }
  if (*c.R1.BC != 0x095E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x095E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD6DF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD6DF, *c.R1.DE)
  }
  if (*c.R1.HL != 0x42FE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x42FE, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB72(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA500
  *c.R1.BC = 0xC0BF
  *c.R1.DE = 0x4C8D
  *c.R1.HL = 0xAD11
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x72
  mem[0xAD11] = 0x3B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA518) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA518, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC0BF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC0BF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4C8D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4C8D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAD11) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAD11, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB73(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF200
  *c.R1.BC = 0x49A6
  *c.R1.DE = 0xB279
  *c.R1.HL = 0x2ECC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x73
  mem[0x2ECC] = 0xE0
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF238) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF238, *c.R1.AF)
  }
  if (*c.R1.BC != 0x49A6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x49A6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB279) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB279, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2ECC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2ECC, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB74(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0500
  *c.R1.BC = 0x445E
  *c.R1.DE = 0x05E9
  *c.R1.HL = 0x983D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x74
  mem[0x983D] = 0xFA
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x055C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x055C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x445E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x445E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x05E9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x05E9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x983D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x983D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB75(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6B00
  *c.R1.BC = 0x83C6
  *c.R1.DE = 0x635A
  *c.R1.HL = 0xD18D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x75
  mem[0xD18D] = 0x11
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6B5C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6B5C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x83C6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x83C6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x635A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x635A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD18D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD18D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB76(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF800
  *c.R1.BC = 0x3057
  *c.R1.DE = 0x3629
  *c.R1.HL = 0xBC71
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x76
  mem[0xBC71] = 0x18
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF85C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF85C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3057) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3057, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3629) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3629, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBC71) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBC71, *c.R1.HL)
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

func TestCB77_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x77
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB77(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9200
  *c.R1.BC = 0xD6F8
  *c.R1.DE = 0x5100
  *c.R1.HL = 0x736D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x77
  mem[0x736D] = 0x36
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9254) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9254, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD6F8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD6F8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5100) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5100, *c.R1.DE)
  }
  if (*c.R1.HL != 0x736D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x736D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB78(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7200
  *c.R1.BC = 0x1CF8
  *c.R1.DE = 0x8D2B
  *c.R1.HL = 0xC76A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x78
  mem[0xC76A] = 0x1F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x725C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x725C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1CF8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1CF8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8D2B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8D2B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC76A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC76A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB79(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA800
  *c.R1.BC = 0x809E
  *c.R1.DE = 0x1124
  *c.R1.HL = 0x39E8
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x79
  mem[0x39E8] = 0x98
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA898) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA898, *c.R1.AF)
  }
  if (*c.R1.BC != 0x809E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x809E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1124) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1124, *c.R1.DE)
  }
  if (*c.R1.HL != 0x39E8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x39E8, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB7A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5800
  *c.R1.BC = 0x7D24
  *c.R1.DE = 0x63E1
  *c.R1.HL = 0xD9AF
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x7A
  mem[0xD9AF] = 0xED
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5874) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5874, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7D24) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7D24, *c.R1.BC)
  }
  if (*c.R1.DE != 0x63E1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x63E1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD9AF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD9AF, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB7B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0300
  *c.R1.BC = 0x50AB
  *c.R1.DE = 0x05BD
  *c.R1.HL = 0x6BD0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x7B
  mem[0x6BD0] = 0xA5
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x03B8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x03B8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x50AB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x50AB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x05BD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x05BD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6BD0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6BD0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB7C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAD00
  *c.R1.BC = 0xF77B
  *c.R1.DE = 0x55AE
  *c.R1.HL = 0x063B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x7C
  mem[0x063B] = 0x34
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAD54) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAD54, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF77B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF77B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x55AE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x55AE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x063B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x063B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB7D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8200
  *c.R1.BC = 0xB792
  *c.R1.DE = 0x38CB
  *c.R1.HL = 0x5F9B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x7D
  mem[0x5F9B] = 0x97
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8298) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8298, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB792) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB792, *c.R1.BC)
  }
  if (*c.R1.DE != 0x38CB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x38CB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5F9B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5F9B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB7E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4200
  *c.R1.BC = 0x3B91
  *c.R1.DE = 0xF59C
  *c.R1.HL = 0xA25E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x7E
  mem[0xA25E] = 0xD7
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4290) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4290, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3B91) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3B91, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF59C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF59C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA25E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA25E, *c.R1.HL)
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

func TestCB7F_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF00
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x7F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFFB8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFFB8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0000) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0000, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB7F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6A00
  *c.R1.BC = 0x84EC
  *c.R1.DE = 0xCF4E
  *c.R1.HL = 0x185B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x7F
  mem[0x185B] = 0xF1
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6A7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6A7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x84EC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x84EC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCF4E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCF4E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x185B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x185B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB80(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8F00
  *c.R1.BC = 0x702F
  *c.R1.DE = 0x17BD
  *c.R1.HL = 0xA706
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x80
  mem[0xA706] = 0x0A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8F00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8F00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x702F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x702F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x17BD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x17BD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA706) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA706, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB81(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAE00
  *c.R1.BC = 0x947F
  *c.R1.DE = 0x7153
  *c.R1.HL = 0x6616
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x81
  mem[0x6616] = 0x74
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAE00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAE00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x947E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x947E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7153) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7153, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6616) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6616, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB82(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8100
  *c.R1.BC = 0xBED2
  *c.R1.DE = 0xC719
  *c.R1.HL = 0x4572
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x82
  mem[0x4572] = 0x2F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8100, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBED2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBED2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC619) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC619, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4572) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4572, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB83(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE600
  *c.R1.BC = 0x63A2
  *c.R1.DE = 0xCCF7
  *c.R1.HL = 0xAE9A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x83
  mem[0xAE9A] = 0x16
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE600, *c.R1.AF)
  }
  if (*c.R1.BC != 0x63A2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x63A2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCCF6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCCF6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAE9A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAE9A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB84(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCE00
  *c.R1.BC = 0xE0CC
  *c.R1.DE = 0xD305
  *c.R1.HL = 0xD6C0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x84
  mem[0xD6C0] = 0x72
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCE00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCE00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE0CC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE0CC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD305) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD305, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD6C0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD6C0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB85(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF300
  *c.R1.BC = 0xED79
  *c.R1.DE = 0x9DB7
  *c.R1.HL = 0xDDA0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x85
  mem[0xDDA0] = 0x8A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF300, *c.R1.AF)
  }
  if (*c.R1.BC != 0xED79) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xED79, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9DB7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9DB7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDDA0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDDA0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB86(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2A00
  *c.R1.BC = 0xB0B9
  *c.R1.DE = 0x9426
  *c.R1.HL = 0x1B48
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x86
  mem[0x1B48] = 0x62
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2A00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2A00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB0B9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB0B9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9426) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9426, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1B48) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1B48, *c.R1.HL)
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

func TestCB87(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1100
  *c.R1.BC = 0x86DC
  *c.R1.DE = 0x1798
  *c.R1.HL = 0xDFC5
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x87
  mem[0xDFC5] = 0xDE
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x86DC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x86DC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1798) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1798, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDFC5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDFC5, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB88(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE300
  *c.R1.BC = 0x8A21
  *c.R1.DE = 0xE33E
  *c.R1.HL = 0x674D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x88
  mem[0x674D] = 0x5F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8821) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8821, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE33E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE33E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x674D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x674D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB89(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6000
  *c.R1.BC = 0xD186
  *c.R1.DE = 0xC5B6
  *c.R1.HL = 0x1BD7
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x89
  mem[0x1BD7] = 0xF2
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD184) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD184, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC5B6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC5B6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1BD7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1BD7, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB8A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3E00
  *c.R1.BC = 0x5FCD
  *c.R1.DE = 0x0B38
  *c.R1.HL = 0xB98E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x8A
  mem[0xB98E] = 0x2F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5FCD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5FCD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0938) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0938, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB98E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB98E, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB8B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6500
  *c.R1.BC = 0x040E
  *c.R1.DE = 0x103F
  *c.R1.HL = 0x4A07
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x8B
  mem[0x4A07] = 0x3F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6500, *c.R1.AF)
  }
  if (*c.R1.BC != 0x040E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x040E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x103D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x103D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4A07) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4A07, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB8C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF800
  *c.R1.BC = 0x6D27
  *c.R1.DE = 0x9BDF
  *c.R1.HL = 0xDAEF
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x8C
  mem[0xDAEF] = 0x0C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF800, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6D27) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6D27, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9BDF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9BDF, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD8EF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD8EF, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB8D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3E00
  *c.R1.BC = 0x5469
  *c.R1.DE = 0x2C28
  *c.R1.HL = 0xBD72
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x8D
  mem[0xBD72] = 0x13
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5469) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5469, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2C28) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2C28, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBD70) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBD70, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB8E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1F00
  *c.R1.BC = 0x140B
  *c.R1.DE = 0xB492
  *c.R1.HL = 0x63A7
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x8E
  mem[0x63A7] = 0xD4
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1F00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1F00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x140B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x140B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB492) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB492, *c.R1.DE)
  }
  if (*c.R1.HL != 0x63A7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x63A7, *c.R1.HL)
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

func TestCB8F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2500
  *c.R1.BC = 0xC522
  *c.R1.DE = 0xCA46
  *c.R1.HL = 0x1C1A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x8F
  mem[0x1C1A] = 0x37
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2500, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC522) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC522, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCA46) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCA46, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1C1A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1C1A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB90(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5700
  *c.R1.BC = 0x595C
  *c.R1.DE = 0x4F0A
  *c.R1.HL = 0xC73C
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x90
  mem[0xC73C] = 0xA2
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5700) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5700, *c.R1.AF)
  }
  if (*c.R1.BC != 0x595C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x595C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4F0A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4F0A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC73C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC73C, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB91(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5E00
  *c.R1.BC = 0x8F26
  *c.R1.DE = 0xA735
  *c.R1.HL = 0x97E0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x91
  mem[0x97E0] = 0x5E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8F22) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8F22, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA735) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA735, *c.R1.DE)
  }
  if (*c.R1.HL != 0x97E0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x97E0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB92(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3300
  *c.R1.BC = 0x7D9F
  *c.R1.DE = 0x87D0
  *c.R1.HL = 0x83D0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x92
  mem[0x83D0] = 0x2B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7D9F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7D9F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x83D0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x83D0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x83D0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x83D0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB93(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC200
  *c.R1.BC = 0x4E05
  *c.R1.DE = 0xB3F8
  *c.R1.HL = 0x2234
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x93
  mem[0x2234] = 0xA0
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4E05) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4E05, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB3F8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB3F8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2234) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2234, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB94(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEE00
  *c.R1.BC = 0x8F4B
  *c.R1.DE = 0x2831
  *c.R1.HL = 0xD6A6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x94
  mem[0xD6A6] = 0xD0
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEE00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEE00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8F4B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8F4B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2831) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2831, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD2A6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD2A6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB95(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3C00
  *c.R1.BC = 0x6AF2
  *c.R1.DE = 0xB25D
  *c.R1.HL = 0x36FF
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x95
  mem[0x36FF] = 0xCD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6AF2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6AF2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB25D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB25D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x36FB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x36FB, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB96(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7600
  *c.R1.BC = 0xB027
  *c.R1.DE = 0xD0A5
  *c.R1.HL = 0x3324
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x96
  mem[0x3324] = 0x21
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7600, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB027) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB027, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD0A5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD0A5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3324) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3324, *c.R1.HL)
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

func TestCB97(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1600
  *c.R1.BC = 0xAD09
  *c.R1.DE = 0x7902
  *c.R1.HL = 0x97BC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x97
  mem[0x97BC] = 0x75
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAD09) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAD09, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7902) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7902, *c.R1.DE)
  }
  if (*c.R1.HL != 0x97BC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x97BC, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB98(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3400
  *c.R1.BC = 0xB61C
  *c.R1.DE = 0x771D
  *c.R1.HL = 0x5D5E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x98
  mem[0x5D5E] = 0xA4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3400, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB61C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB61C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x771D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x771D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5D5E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5D5E, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB99(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5100
  *c.R1.BC = 0x65BE
  *c.R1.DE = 0x1359
  *c.R1.HL = 0x8BEC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x99
  mem[0x8BEC] = 0x0B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5100, *c.R1.AF)
  }
  if (*c.R1.BC != 0x65B6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x65B6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1359) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1359, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8BEC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8BEC, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB9A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6400
  *c.R1.BC = 0x976D
  *c.R1.DE = 0x4C25
  *c.R1.HL = 0xDCB2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x9A
  mem[0xDCB2] = 0x09
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6400, *c.R1.AF)
  }
  if (*c.R1.BC != 0x976D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x976D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4425) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4425, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCB2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCB2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB9B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA100
  *c.R1.BC = 0xB58A
  *c.R1.DE = 0xD264
  *c.R1.HL = 0x2BD6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x9B
  mem[0x2BD6] = 0xD3
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA100, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB58A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB58A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD264) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD264, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2BD6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2BD6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB9C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD800
  *c.R1.BC = 0x63D6
  *c.R1.DE = 0xAC7B
  *c.R1.HL = 0xC7A0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x9C
  mem[0xC7A0] = 0x75
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD800, *c.R1.AF)
  }
  if (*c.R1.BC != 0x63D6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x63D6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAC7B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAC7B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC7A0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC7A0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB9D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0D00
  *c.R1.BC = 0xD840
  *c.R1.DE = 0x0810
  *c.R1.HL = 0x0800
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x9D
  mem[0x0800] = 0xCD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0D00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0D00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD840) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD840, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0810) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0810, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0800) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0800, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCB9E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3B00
  *c.R1.BC = 0xEBBF
  *c.R1.DE = 0x9434
  *c.R1.HL = 0x3A65
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x9E
  mem[0x3A65] = 0x2A
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3B00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3B00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEBBF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEBBF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9434) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9434, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3A65) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3A65, *c.R1.HL)
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

func TestCB9F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB200
  *c.R1.BC = 0xD1DE
  *c.R1.DE = 0xF991
  *c.R1.HL = 0x72F6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0x9F
  mem[0x72F6] = 0x72
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD1DE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD1DE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF991) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF991, *c.R1.DE)
  }
  if (*c.R1.HL != 0x72F6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x72F6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFA00
  *c.R1.BC = 0xD669
  *c.R1.DE = 0x71E1
  *c.R1.HL = 0xC80D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA0
  mem[0xC80D] = 0xC0
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFA00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFA00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC669) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC669, *c.R1.BC)
  }
  if (*c.R1.DE != 0x71E1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x71E1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC80D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC80D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8200
  *c.R1.BC = 0x75E4
  *c.R1.DE = 0xA0DE
  *c.R1.HL = 0xD0BA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA1
  mem[0xD0BA] = 0xBD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x75E4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x75E4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA0DE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA0DE, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD0BA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD0BA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDD00
  *c.R1.BC = 0x2B0D
  *c.R1.DE = 0x5554
  *c.R1.HL = 0x6FC0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA2
  mem[0x6FC0] = 0x61
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDD00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDD00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2B0D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2B0D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4554) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4554, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6FC0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6FC0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2200
  *c.R1.BC = 0x2F0D
  *c.R1.DE = 0x4D2C
  *c.R1.HL = 0x6666
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA3
  mem[0x6666] = 0x8E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2F0D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2F0D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4D2C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4D2C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6666) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6666, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD600
  *c.R1.BC = 0xD8ED
  *c.R1.DE = 0x9CD4
  *c.R1.HL = 0x8BB1
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA4
  mem[0x8BB1] = 0xBB
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD600, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD8ED) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD8ED, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9CD4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9CD4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8BB1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8BB1, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB400
  *c.R1.BC = 0xB393
  *c.R1.DE = 0x3E42
  *c.R1.HL = 0x88CA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA5
  mem[0x88CA] = 0x4F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB400, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB393) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB393, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3E42) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3E42, *c.R1.DE)
  }
  if (*c.R1.HL != 0x88CA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x88CA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0A00
  *c.R1.BC = 0x4C34
  *c.R1.DE = 0xF5A7
  *c.R1.HL = 0xE70D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA6
  mem[0xE70D] = 0x27
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0A00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0A00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4C34) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4C34, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF5A7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF5A7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE70D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE70D, *c.R1.HL)
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

func TestCBA7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4500
  *c.R1.BC = 0xAF61
  *c.R1.DE = 0x569A
  *c.R1.HL = 0xC77B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA7
  mem[0xC77B] = 0xFF
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4500, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAF61) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAF61, *c.R1.BC)
  }
  if (*c.R1.DE != 0x569A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x569A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC77B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC77B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6400
  *c.R1.BC = 0xF269
  *c.R1.DE = 0xBAE4
  *c.R1.HL = 0xC9E7
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA8
  mem[0xC9E7] = 0x46
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6400, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD269) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD269, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBAE4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBAE4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC9E7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC9E7, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBA9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE400
  *c.R1.BC = 0x7AD4
  *c.R1.DE = 0xBF0A
  *c.R1.HL = 0xCE0B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xA9
  mem[0xCE0B] = 0x39
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE400, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7AD4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7AD4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBF0A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBF0A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCE0B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCE0B, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBAA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCD00
  *c.R1.BC = 0xD249
  *c.R1.DE = 0x4159
  *c.R1.HL = 0xFED5
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xAA
  mem[0xFED5] = 0xB0
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCD00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCD00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD249) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD249, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4159) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4159, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFED5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFED5, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBAB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAC00
  *c.R1.BC = 0x939A
  *c.R1.DE = 0x5D9B
  *c.R1.HL = 0x0812
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xAB
  mem[0x0812] = 0xF2
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAC00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAC00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x939A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x939A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5D9B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5D9B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0812) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0812, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBAC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2400
  *c.R1.BC = 0x8A7D
  *c.R1.DE = 0x2CAC
  *c.R1.HL = 0xFFAA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xAC
  mem[0xFFAA] = 0x09
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2400, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8A7D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8A7D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2CAC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2CAC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDFAA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDFAA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBAD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6F00
  *c.R1.BC = 0x5FFB
  *c.R1.DE = 0x2360
  *c.R1.HL = 0xAE15
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xAD
  mem[0xAE15] = 0x30
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6F00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6F00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5FFB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5FFB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2360) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2360, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAE15) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAE15, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBAE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5A00
  *c.R1.BC = 0xAA17
  *c.R1.DE = 0x12F3
  *c.R1.HL = 0x190E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xAE
  mem[0x190E] = 0x66
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5A00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5A00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAA17) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAA17, *c.R1.BC)
  }
  if (*c.R1.DE != 0x12F3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x12F3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x190E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x190E, *c.R1.HL)
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

func TestCBAF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFC00
  *c.R1.BC = 0xBB3F
  *c.R1.DE = 0x8BB6
  *c.R1.HL = 0x5877
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xAF
  mem[0x5877] = 0x62
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDC00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDC00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBB3F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBB3F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8BB6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8BB6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5877) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5877, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB900
  *c.R1.BC = 0x7A79
  *c.R1.DE = 0x1AAA
  *c.R1.HL = 0xC3BA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB0
  mem[0xC3BA] = 0x4C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB900, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3A79) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3A79, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1AAA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1AAA, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC3BA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC3BA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4900
  *c.R1.BC = 0x63E4
  *c.R1.DE = 0xA544
  *c.R1.HL = 0x1190
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB1
  mem[0x1190] = 0xE3
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4900, *c.R1.AF)
  }
  if (*c.R1.BC != 0x63A4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x63A4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA544) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA544, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1190) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1190, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4D00
  *c.R1.BC = 0x2B03
  *c.R1.DE = 0x6B23
  *c.R1.HL = 0x6FF5
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB2
  mem[0x6FF5] = 0x04
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4D00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4D00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2B03) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2B03, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2B23) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2B23, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6FF5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6FF5, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8700
  *c.R1.BC = 0x857A
  *c.R1.DE = 0xE98B
  *c.R1.HL = 0x5CB1
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB3
  mem[0x5CB1] = 0x43
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8700) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8700, *c.R1.AF)
  }
  if (*c.R1.BC != 0x857A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x857A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE98B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE98B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5CB1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5CB1, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2B00
  *c.R1.BC = 0xB73E
  *c.R1.DE = 0x79C9
  *c.R1.HL = 0xE1BB
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB4
  mem[0xE1BB] = 0x78
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2B00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2B00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB73E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB73E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x79C9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x79C9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA1BB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA1BB, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9B00
  *c.R1.BC = 0xD879
  *c.R1.DE = 0x2EC9
  *c.R1.HL = 0x4BBA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB5
  mem[0x4BBA] = 0x70
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9B00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9B00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD879) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD879, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2EC9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2EC9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4BBA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4BBA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8600
  *c.R1.BC = 0x89BF
  *c.R1.DE = 0xDE4A
  *c.R1.HL = 0x4FAB
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB6
  mem[0x4FAB] = 0xA5
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8600, *c.R1.AF)
  }
  if (*c.R1.BC != 0x89BF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x89BF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDE4A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDE4A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4FAB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4FAB, *c.R1.HL)
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

func TestCBB7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2200
  *c.R1.BC = 0xFB8A
  *c.R1.DE = 0x3D6E
  *c.R1.HL = 0xD4A2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB7
  mem[0xD4A2] = 0xF2
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFB8A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFB8A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3D6E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3D6E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD4A2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD4A2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD000
  *c.R1.BC = 0x37C6
  *c.R1.DE = 0x225A
  *c.R1.HL = 0xD249
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB8
  mem[0xD249] = 0xC4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x37C6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x37C6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x225A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x225A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD249) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD249, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBB9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA500
  *c.R1.BC = 0x1B4A
  *c.R1.DE = 0xD584
  *c.R1.HL = 0x5DEE
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xB9
  mem[0x5DEE] = 0xCC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA500, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1B4A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1B4A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD584) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD584, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5DEE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5DEE, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBBA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6300
  *c.R1.BC = 0xA5FE
  *c.R1.DE = 0xF42B
  *c.R1.HL = 0x34C9
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xBA
  mem[0x34C9] = 0xBC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6300, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA5FE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA5FE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x742B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x742B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x34C9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x34C9, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBBB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1200
  *c.R1.BC = 0xF661
  *c.R1.DE = 0xAA4F
  *c.R1.HL = 0xCB30
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xBB
  mem[0xCB30] = 0xF4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF661) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF661, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAA4F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAA4F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCB30) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCB30, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBBC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9800
  *c.R1.BC = 0xADC3
  *c.R1.DE = 0x0B29
  *c.R1.HL = 0x7B6E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xBC
  mem[0x7B6E] = 0x45
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9800, *c.R1.AF)
  }
  if (*c.R1.BC != 0xADC3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xADC3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0B29) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0B29, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7B6E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7B6E, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBBD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD600
  *c.R1.BC = 0xA6E1
  *c.R1.DE = 0x8813
  *c.R1.HL = 0x10B8
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xBD
  mem[0x10B8] = 0x35
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD600, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA6E1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA6E1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8813) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8813, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1038) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1038, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBBE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCA00
  *c.R1.BC = 0xFF64
  *c.R1.DE = 0x1218
  *c.R1.HL = 0x77D5
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xBE
  mem[0x77D5] = 0xEA
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCA00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCA00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFF64) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFF64, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1218) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1218, *c.R1.DE)
  }
  if (*c.R1.HL != 0x77D5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x77D5, *c.R1.HL)
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

func TestCBBF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6800
  *c.R1.BC = 0x4845
  *c.R1.DE = 0x690A
  *c.R1.HL = 0x15DE
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xBF
  mem[0x15DE] = 0x1D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6800, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4845) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4845, *c.R1.BC)
  }
  if (*c.R1.DE != 0x690A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x690A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x15DE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x15DE, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE300
  *c.R1.BC = 0xEF71
  *c.R1.DE = 0xBFFB
  *c.R1.HL = 0xB3A1
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC0
  mem[0xB3A1] = 0x5C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE300, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEF71) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEF71, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBFFB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBFFB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB3A1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB3A1, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3200
  *c.R1.BC = 0x32A1
  *c.R1.DE = 0x59AB
  *c.R1.HL = 0x3343
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC1
  mem[0x3343] = 0xAA
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x32A1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x32A1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x59AB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x59AB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3343) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3343, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC700
  *c.R1.BC = 0xB159
  *c.R1.DE = 0xC023
  *c.R1.HL = 0xE1F3
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC2
  mem[0xE1F3] = 0x14
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC700) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC700, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB159) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB159, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC123) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC123, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE1F3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE1F3, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0400
  *c.R1.BC = 0xB463
  *c.R1.DE = 0xC211
  *c.R1.HL = 0x8F3A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC3
  mem[0x8F3A] = 0x81
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0400, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB463) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB463, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC211) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC211, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8F3A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8F3A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7E00
  *c.R1.BC = 0x545A
  *c.R1.DE = 0x6ECF
  *c.R1.HL = 0x5876
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC4
  mem[0x5876] = 0x9D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x545A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x545A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6ECF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6ECF, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5976) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5976, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4000
  *c.R1.BC = 0xC617
  *c.R1.DE = 0x079C
  *c.R1.HL = 0x4107
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC5
  mem[0x4107] = 0xCC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC617) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC617, *c.R1.BC)
  }
  if (*c.R1.DE != 0x079C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x079C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4107) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4107, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB800
  *c.R1.BC = 0x0373
  *c.R1.DE = 0xB807
  *c.R1.HL = 0xF0BE
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC6
  mem[0xF0BE] = 0x9C
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB800, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0373) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0373, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB807) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB807, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF0BE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF0BE, *c.R1.HL)
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

func TestCBC7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7700
  *c.R1.BC = 0x3681
  *c.R1.DE = 0x9B55
  *c.R1.HL = 0x583F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC7
  mem[0x583F] = 0x58
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7700) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7700, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3681) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3681, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9B55) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9B55, *c.R1.DE)
  }
  if (*c.R1.HL != 0x583F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x583F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7D00
  *c.R1.BC = 0xA772
  *c.R1.DE = 0x8682
  *c.R1.HL = 0x7CF3
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC8
  mem[0x7CF3] = 0x75
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7D00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7D00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA772) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA772, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8682) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8682, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7CF3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7CF3, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBC9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0B00
  *c.R1.BC = 0x67EE
  *c.R1.DE = 0x30E0
  *c.R1.HL = 0x72DB
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xC9
  mem[0x72DB] = 0x87
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0B00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0B00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x67EE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x67EE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x30E0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x30E0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x72DB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x72DB, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBCA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9C00
  *c.R1.BC = 0x9517
  *c.R1.DE = 0xCFBB
  *c.R1.HL = 0xFBC7
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xCA
  mem[0xFBC7] = 0x1A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9C00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9C00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9517) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9517, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCFBB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCFBB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFBC7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFBC7, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBCB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE800
  *c.R1.BC = 0x0F3D
  *c.R1.DE = 0x336F
  *c.R1.HL = 0xF70D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xCB
  mem[0xF70D] = 0xA1
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE800, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x336F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x336F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF70D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF70D, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBCC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFB00
  *c.R1.BC = 0x7981
  *c.R1.DE = 0x0BBB
  *c.R1.HL = 0x18FD
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xCC
  mem[0x18FD] = 0xFE
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFB00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFB00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7981) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7981, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0BBB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0BBB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1AFD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1AFD, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBCD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5500
  *c.R1.BC = 0x5E78
  *c.R1.DE = 0xBF34
  *c.R1.HL = 0x2602
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xCD
  mem[0x2602] = 0x2D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5500, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5E78) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5E78, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBF34) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBF34, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2602) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2602, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBCE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD500
  *c.R1.BC = 0xA111
  *c.R1.DE = 0xCB2A
  *c.R1.HL = 0x8EC6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xCE
  mem[0x8EC6] = 0xBF
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD500, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA111) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA111, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCB2A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCB2A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8EC6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8EC6, *c.R1.HL)
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

func TestCBCF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA200
  *c.R1.BC = 0x6BAF
  *c.R1.DE = 0x98B2
  *c.R1.HL = 0x98A0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xCF
  mem[0x98A0] = 0xD4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6BAF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6BAF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x98B2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x98B2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x98A0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x98A0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2300
  *c.R1.BC = 0x7BCB
  *c.R1.DE = 0x02E7
  *c.R1.HL = 0x1724
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD0
  mem[0x1724] = 0x30
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7FCB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7FCB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x02E7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x02E7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1724) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1724, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5300
  *c.R1.BC = 0x581F
  *c.R1.DE = 0xB775
  *c.R1.HL = 0x47F4
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD1
  mem[0x47F4] = 0xC7
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x581F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x581F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB775) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB775, *c.R1.DE)
  }
  if (*c.R1.HL != 0x47F4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x47F4, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6900
  *c.R1.BC = 0xC147
  *c.R1.DE = 0xB79C
  *c.R1.HL = 0x7528
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD2
  mem[0x7528] = 0x4F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6900, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC147) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC147, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB79C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB79C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7528) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7528, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAE00
  *c.R1.BC = 0xBBC4
  *c.R1.DE = 0xCE52
  *c.R1.HL = 0x5FBA
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD3
  mem[0x5FBA] = 0x3A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAE00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAE00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBBC4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBBC4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCE56) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCE56, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5FBA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5FBA, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD800
  *c.R1.BC = 0x6E1E
  *c.R1.DE = 0xAF6F
  *c.R1.HL = 0xBF2E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD4
  mem[0xBF2E] = 0x71
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD800, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6E1E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6E1E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAF6F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAF6F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBF2E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBF2E, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8400
  *c.R1.BC = 0xA19A
  *c.R1.DE = 0xD2FD
  *c.R1.HL = 0x8A77
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD5
  mem[0x8A77] = 0x52
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8400, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA19A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA19A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD2FD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD2FD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8A77) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8A77, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA900
  *c.R1.BC = 0xF5F3
  *c.R1.DE = 0x2180
  *c.R1.HL = 0x6029
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD6
  mem[0x6029] = 0xB7
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA900, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF5F3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF5F3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2180) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2180, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6029) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6029, *c.R1.HL)
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

func TestCBD7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB100
  *c.R1.BC = 0xC008
  *c.R1.DE = 0x8425
  *c.R1.HL = 0x290A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD7
  mem[0x290A] = 0x42
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB500, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC008) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC008, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8425) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8425, *c.R1.DE)
  }
  if (*c.R1.HL != 0x290A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x290A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8B00
  *c.R1.BC = 0x09C4
  *c.R1.DE = 0xDDF3
  *c.R1.HL = 0x6D7E
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD8
  mem[0x6D7E] = 0x6E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8B00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8B00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x09C4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x09C4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDDF3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDDF3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6D7E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6D7E, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBD9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3E00
  *c.R1.BC = 0x3E36
  *c.R1.DE = 0x30EC
  *c.R1.HL = 0xEFC6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xD9
  mem[0xEFC6] = 0x5B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3E3E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3E3E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x30EC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x30EC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEFC6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEFC6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBDA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD000
  *c.R1.BC = 0x3E8F
  *c.R1.DE = 0x28FE
  *c.R1.HL = 0x1C87
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xDA
  mem[0x1C87] = 0xB9
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3E8F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3E8F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x28FE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x28FE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1C87) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1C87, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBDB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1200
  *c.R1.BC = 0x977A
  *c.R1.DE = 0x8C49
  *c.R1.HL = 0xBC48
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xDB
  mem[0xBC48] = 0xEF
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x977A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x977A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8C49) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8C49, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBC48) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBC48, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBDC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8D00
  *c.R1.BC = 0x05DE
  *c.R1.DE = 0xF8D3
  *c.R1.HL = 0xB125
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xDC
  mem[0xB125] = 0x0E
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8D00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8D00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x05DE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x05DE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF8D3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF8D3, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB925) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB925, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBDD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC300
  *c.R1.BC = 0x08A9
  *c.R1.DE = 0x2BC8
  *c.R1.HL = 0x5B9F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xDD
  mem[0x5B9F] = 0x94
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x08A9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x08A9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2BC8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2BC8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5B9F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5B9F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBDE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1900
  *c.R1.BC = 0x900F
  *c.R1.DE = 0xD572
  *c.R1.HL = 0xBA03
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xDE
  mem[0xBA03] = 0x93
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1900, *c.R1.AF)
  }
  if (*c.R1.BC != 0x900F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x900F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD572) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD572, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBA03) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBA03, *c.R1.HL)
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

func TestCBDF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6700
  *c.R1.BC = 0x2745
  *c.R1.DE = 0x7E3D
  *c.R1.HL = 0x0FA1
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xDF
  mem[0x0FA1] = 0xC5
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6F00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6F00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2745) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2745, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7E3D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7E3D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0FA1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0FA1, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3E00
  *c.R1.BC = 0xD633
  *c.R1.DE = 0x9897
  *c.R1.HL = 0x3744
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE0
  mem[0x3744] = 0x54
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD633) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD633, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9897) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9897, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3744) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3744, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7D00
  *c.R1.BC = 0x50A6
  *c.R1.DE = 0x0136
  *c.R1.HL = 0x5334
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE1
  mem[0x5334] = 0x85
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7D00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7D00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x50B6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x50B6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0136) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0136, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5334) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5334, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD400
  *c.R1.BC = 0x6B45
  *c.R1.DE = 0xA192
  *c.R1.HL = 0x3A4C
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE2
  mem[0x3A4C] = 0x47
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD400, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6B45) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6B45, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB192) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB192, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3A4C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3A4C, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3B00
  *c.R1.BC = 0xD29C
  *c.R1.DE = 0x05E0
  *c.R1.HL = 0x2E78
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE3
  mem[0x2E78] = 0x48
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3B00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3B00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD29C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD29C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x05F0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x05F0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2E78) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2E78, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1E00
  *c.R1.BC = 0x7D5E
  *c.R1.DE = 0x846D
  *c.R1.HL = 0x0978
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE4
  mem[0x0978] = 0x84
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7D5E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7D5E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x846D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x846D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1978) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1978, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCA00
  *c.R1.BC = 0xDF0D
  *c.R1.DE = 0xD588
  *c.R1.HL = 0xB48F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE5
  mem[0xB48F] = 0xCF
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCA00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCA00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDF0D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDF0D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD588) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD588, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB49F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB49F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB300
  *c.R1.BC = 0x52C2
  *c.R1.DE = 0xDBFE
  *c.R1.HL = 0x9F9B
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE6
  mem[0x9F9B] = 0xF6
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x52C2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x52C2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDBFE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDBFE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9F9B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9F9B, *c.R1.HL)
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

func TestCBE7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8E00
  *c.R1.BC = 0xCF02
  *c.R1.DE = 0x67EF
  *c.R1.HL = 0xF2E0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE7
  mem[0xF2E0] = 0xCF
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9E00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF02) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF02, *c.R1.BC)
  }
  if (*c.R1.DE != 0x67EF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x67EF, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF2E0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF2E0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7100
  *c.R1.BC = 0xBB18
  *c.R1.DE = 0x66EC
  *c.R1.HL = 0x4A05
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE8
  mem[0x4A05] = 0xE6
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7100, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBB18) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBB18, *c.R1.BC)
  }
  if (*c.R1.DE != 0x66EC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x66EC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4A05) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4A05, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBE9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5700
  *c.R1.BC = 0x2897
  *c.R1.DE = 0x8F2F
  *c.R1.HL = 0xA4D0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xE9
  mem[0xA4D0] = 0xB2
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5700) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5700, *c.R1.AF)
  }
  if (*c.R1.BC != 0x28B7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x28B7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8F2F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8F2F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA4D0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA4D0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBEA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEC00
  *c.R1.BC = 0x304A
  *c.R1.DE = 0x60A1
  *c.R1.HL = 0xF32A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xEA
  mem[0xF32A] = 0x9C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEC00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEC00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x304A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x304A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x60A1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x60A1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF32A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF32A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBEB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF000
  *c.R1.BC = 0x532B
  *c.R1.DE = 0xA1BE
  *c.R1.HL = 0x1A1A
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xEB
  mem[0x1A1A] = 0x21
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x532B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x532B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA1BE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA1BE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1A1A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1A1A, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBEC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF200
  *c.R1.BC = 0xF0F3
  *c.R1.DE = 0xA816
  *c.R1.HL = 0xBA08
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xEC
  mem[0xBA08] = 0x82
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF0F3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF0F3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA816) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA816, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBA08) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBA08, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBED(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1300
  *c.R1.BC = 0x5127
  *c.R1.DE = 0xADAB
  *c.R1.HL = 0x2DEC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xED
  mem[0x2DEC] = 0xCB
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5127) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5127, *c.R1.BC)
  }
  if (*c.R1.DE != 0xADAB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xADAB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2DEC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2DEC, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBEE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9000
  *c.R1.BC = 0xB273
  *c.R1.DE = 0x50AE
  *c.R1.HL = 0xE90D
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xEE
  mem[0xE90D] = 0xF1
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB273) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB273, *c.R1.BC)
  }
  if (*c.R1.DE != 0x50AE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x50AE, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE90D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE90D, *c.R1.HL)
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

func TestCBEF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2500
  *c.R1.BC = 0x4281
  *c.R1.DE = 0xF0D4
  *c.R1.HL = 0x2C39
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xEF
  mem[0x2C39] = 0xC8
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2500, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4281) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4281, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF0D4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF0D4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2C39) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2C39, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFB00
  *c.R1.BC = 0x5802
  *c.R1.DE = 0x0C27
  *c.R1.HL = 0x6FF5
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF0
  mem[0x6FF5] = 0xF6
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFB00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFB00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5802) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5802, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0C27) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0C27, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6FF5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6FF5, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5500
  *c.R1.BC = 0xA103
  *c.R1.DE = 0x3FF5
  *c.R1.HL = 0x5E1C
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF1
  mem[0x5E1C] = 0x37
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5500) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5500, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA143) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA143, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3FF5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3FF5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E1C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E1C, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF000
  *c.R1.BC = 0x625A
  *c.R1.DE = 0xAF82
  *c.R1.HL = 0x9819
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF2
  mem[0x9819] = 0xE4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x625A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x625A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEF82) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEF82, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9819) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9819, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8600
  *c.R1.BC = 0xD7BD
  *c.R1.DE = 0x5D86
  *c.R1.HL = 0x263F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF3
  mem[0x263F] = 0xA1
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8600, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD7BD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD7BD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5DC6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5DC6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x263F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x263F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9400
  *c.R1.BC = 0x0243
  *c.R1.DE = 0x9EC1
  *c.R1.HL = 0x75D9
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF4
  mem[0x75D9] = 0x3F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9400) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9400, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0243) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0243, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9EC1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9EC1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x75D9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x75D9, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCE00
  *c.R1.BC = 0x2D42
  *c.R1.DE = 0x5E6A
  *c.R1.HL = 0x47E6
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF5
  mem[0x47E6] = 0xCE
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCE00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCE00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2D42) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2D42, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5E6A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5E6A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x47E6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x47E6, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7B00
  *c.R1.BC = 0xC2D7
  *c.R1.DE = 0x4492
  *c.R1.HL = 0xA9BC
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF6
  mem[0xA9BC] = 0xB1
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7B00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7B00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC2D7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC2D7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4492) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4492, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA9BC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA9BC, *c.R1.HL)
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

func TestCBF7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6D00
  *c.R1.BC = 0xABAF
  *c.R1.DE = 0x5B5D
  *c.R1.HL = 0x188C
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF7
  mem[0x188C] = 0x6C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6D00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6D00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xABAF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xABAF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5B5D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5B5D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x188C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x188C, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC600
  *c.R1.BC = 0xB812
  *c.R1.DE = 0xA037
  *c.R1.HL = 0xD2B0
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF8
  mem[0xD2B0] = 0xCB
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC600, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB812) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB812, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA037) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA037, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD2B0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD2B0, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBF9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEF00
  *c.R1.BC = 0xC5F2
  *c.R1.DE = 0x77A8
  *c.R1.HL = 0x0730
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xF9
  mem[0x0730] = 0xAE
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEF00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEF00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC5F2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC5F2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x77A8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x77A8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0730) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0730, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBFA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8700
  *c.R1.BC = 0x1581
  *c.R1.DE = 0x63E3
  *c.R1.HL = 0xED03
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xFA
  mem[0xED03] = 0x27
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8700) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8700, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1581) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1581, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE3E3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE3E3, *c.R1.DE)
  }
  if (*c.R1.HL != 0xED03) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xED03, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBFB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA300
  *c.R1.BC = 0x7D27
  *c.R1.DE = 0x97C3
  *c.R1.HL = 0xD1AE
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xFB
  mem[0xD1AE] = 0xF2
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7D27) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7D27, *c.R1.BC)
  }
  if (*c.R1.DE != 0x97C3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x97C3, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD1AE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD1AE, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBFC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEC00
  *c.R1.BC = 0x060A
  *c.R1.DE = 0x3EF6
  *c.R1.HL = 0x500F
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xFC
  mem[0x500F] = 0x94
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEC00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEC00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x060A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x060A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3EF6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3EF6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD00F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD00F, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBFD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1100
  *c.R1.BC = 0x231A
  *c.R1.DE = 0x8563
  *c.R1.HL = 0x28C5
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xFD
  mem[0x28C5] = 0xAB
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1100, *c.R1.AF)
  }
  if (*c.R1.BC != 0x231A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x231A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8563) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8563, *c.R1.DE)
  }
  if (*c.R1.HL != 0x28C5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x28C5, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestCBFE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5300
  *c.R1.BC = 0x4948
  *c.R1.DE = 0x89DD
  *c.R1.HL = 0x3A24
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xFE
  mem[0x3A24] = 0xC3
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4948) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4948, *c.R1.BC)
  }
  if (*c.R1.DE != 0x89DD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x89DD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3A24) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3A24, *c.R1.HL)
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

func TestCBFF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7900
  *c.R1.BC = 0x799B
  *c.R1.DE = 0x6CF7
  *c.R1.HL = 0xE3F2
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
  mem[0x0000] = 0xCB
  mem[0x0001] = 0xFF
  mem[0xE3F2] = 0x25
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF900, *c.R1.AF)
  }
  if (*c.R1.BC != 0x799B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x799B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6CF7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6CF7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE3F2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE3F2, *c.R1.HL)
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
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}