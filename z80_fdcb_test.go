package z80

import "testing"


func TestFDCB00(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x85AC
  *c.R1.BC = 0x46D0
  *c.R1.DE = 0xA135
  *c.R1.HL = 0x20C5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB8DE
  *c.R1.IY = 0x2776
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0B
  mem[0x0003] = 0x00
  mem[0x2781] = 0x50
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x85A4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x85A4, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA0D0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA0D0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA135) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA135, *c.R1.DE)
  }
  if (*c.R1.HL != 0x20C5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x20C5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB8DE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB8DE, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2776) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2776, *c.R1.IY)
  }
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

func TestFDCB01(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x577C
  *c.R1.BC = 0x2B76
  *c.R1.DE = 0x3576
  *c.R1.HL = 0x280A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAE22
  *c.R1.IY = 0x5C35
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC8
  mem[0x0003] = 0x01
  mem[0x5BFD] = 0xCB
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5781) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5781, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2B97) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2B97, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3576) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3576, *c.R1.DE)
  }
  if (*c.R1.HL != 0x280A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x280A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAE22) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAE22, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5C35) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5C35, *c.R1.IY)
  }
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

func TestFDCB02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDC23
  *c.R1.BC = 0x2B37
  *c.R1.DE = 0x83C8
  *c.R1.HL = 0x5DD9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB2D2
  *c.R1.IY = 0x3DF2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x14
  mem[0x0003] = 0x02
  mem[0x3E06] = 0x58
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDCA0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDCA0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2B37) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2B37, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB0C8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB0C8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5DD9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5DD9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB2D2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB2D2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3DF2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3DF2, *c.R1.IY)
  }
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

func TestFDCB03(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x57EE
  *c.R1.BC = 0xC179
  *c.R1.DE = 0xB2B6
  *c.R1.HL = 0x7058
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3F2E
  *c.R1.IY = 0x57E7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3A
  mem[0x0003] = 0x03
  mem[0x5821] = 0x1A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5720) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5720, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC179) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC179, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB234) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB234, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7058) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7058, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3F2E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3F2E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x57E7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x57E7, *c.R1.IY)
  }
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

func TestFDCB04(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xED18
  *c.R1.BC = 0x3F03
  *c.R1.DE = 0x3327
  *c.R1.HL = 0xF35A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCBF2
  *c.R1.IY = 0x5071
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x67
  mem[0x0003] = 0x04
  mem[0x50D8] = 0x92
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xED21) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xED21, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3F03) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3F03, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3327) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3327, *c.R1.DE)
  }
  if (*c.R1.HL != 0x255A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x255A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCBF2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCBF2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5071) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5071, *c.R1.IY)
  }
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

func TestFDCB05(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7A39
  *c.R1.BC = 0x0858
  *c.R1.DE = 0xDB6C
  *c.R1.HL = 0xDBE0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x157A
  *c.R1.IY = 0xB25B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1E
  mem[0x0003] = 0x05
  mem[0xB279] = 0x66
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7A8C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7A8C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0858) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0858, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDB6C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDB6C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDBCC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDBCC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x157A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x157A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB25B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB25B, *c.R1.IY)
  }
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

func TestFDCB06(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF285
  *c.R1.BC = 0x89A2
  *c.R1.DE = 0xE78F
  *c.R1.HL = 0xEF74
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x140D
  *c.R1.IY = 0xFF27
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x72
  mem[0x0003] = 0x06
  mem[0xFF99] = 0xF1
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF2A1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF2A1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x89A2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x89A2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE78F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE78F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEF74) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEF74, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x140D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x140D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFF27) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFF27, *c.R1.IY)
  }
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

func TestFDCB07(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8CCE
  *c.R1.BC = 0xF3A7
  *c.R1.DE = 0x3A6E
  *c.R1.HL = 0x8F0A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8423
  *c.R1.IY = 0x07EB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x24
  mem[0x0003] = 0x07
  mem[0x080F] = 0xAE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5D09) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5D09, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF3A7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF3A7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3A6E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3A6E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8F0A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8F0A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8423) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8423, *c.R1.IX)
  }
  if (*c.R1.IY != 0x07EB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x07EB, *c.R1.IY)
  }
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

func TestFDCB08(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA611
  *c.R1.BC = 0xE8EC
  *c.R1.DE = 0xC958
  *c.R1.HL = 0x7BDA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x194D
  *c.R1.IY = 0x6137
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x25
  mem[0x0003] = 0x08
  mem[0x615C] = 0x83
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA681) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA681, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC1EC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC1EC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC958) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC958, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7BDA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7BDA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x194D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x194D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6137) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6137, *c.R1.IY)
  }
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

func TestFDCB09(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x54B1
  *c.R1.BC = 0xFA1A
  *c.R1.DE = 0x84E8
  *c.R1.HL = 0x4FA5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1AD3
  *c.R1.IY = 0x19DA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA0
  mem[0x0003] = 0x09
  mem[0x197A] = 0x27
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5485) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5485, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFA93) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFA93, *c.R1.BC)
  }
  if (*c.R1.DE != 0x84E8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x84E8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4FA5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4FA5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1AD3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1AD3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x19DA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x19DA, *c.R1.IY)
  }
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

func TestFDCB0A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB3EF
  *c.R1.BC = 0xA2BB
  *c.R1.DE = 0xE5D6
  *c.R1.HL = 0x9617
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF946
  *c.R1.IY = 0xEEF6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE1
  mem[0x0003] = 0x0A
  mem[0xEED7] = 0x19
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB389) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB389, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA2BB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA2BB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8CD6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8CD6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9617) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9617, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF946) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF946, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEEF6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEEF6, *c.R1.IY)
  }
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

func TestFDCB0B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAE10
  *c.R1.BC = 0x8C4E
  *c.R1.DE = 0xE159
  *c.R1.HL = 0x1C54
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE108
  *c.R1.IY = 0xC68F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0C
  mem[0x0003] = 0x0B
  mem[0xC69B] = 0xF2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAE28) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAE28, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8C4E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8C4E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE179) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE179, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1C54) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1C54, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE108) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE108, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC68F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC68F, *c.R1.IY)
  }
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

func TestFDCB0C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8719
  *c.R1.BC = 0x6B16
  *c.R1.DE = 0x4C3B
  *c.R1.HL = 0x180A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x175A
  *c.R1.IY = 0x8C9D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD7
  mem[0x0003] = 0x0C
  mem[0x8C74] = 0xAE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8700) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8700, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6B16) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6B16, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4C3B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4C3B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x570A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x570A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x175A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x175A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8C9D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8C9D, *c.R1.IY)
  }
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

func TestFDCB0D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1204
  *c.R1.BC = 0xE0CB
  *c.R1.DE = 0x3AB1
  *c.R1.HL = 0x2416
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1DE4
  *c.R1.IY = 0xFE2D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x11
  mem[0x0003] = 0x0D
  mem[0xFE3E] = 0x1B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x128D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x128D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE0CB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE0CB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3AB1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3AB1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x248D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x248D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1DE4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1DE4, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFE2D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFE2D, *c.R1.IY)
  }
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

func TestFDCB0E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8DA4
  *c.R1.BC = 0x8F91
  *c.R1.DE = 0xFC5A
  *c.R1.HL = 0x5E2C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB2F2
  *c.R1.IY = 0xF223
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0C
  mem[0x0003] = 0x0E
  mem[0xF22F] = 0xF7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8DA9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8DA9, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8F91) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8F91, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFC5A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFC5A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E2C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E2C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB2F2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB2F2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF223) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF223, *c.R1.IY)
  }
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

func TestFDCB0F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFBB0
  *c.R1.BC = 0x2AC9
  *c.R1.DE = 0xEC6B
  *c.R1.HL = 0x6511
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC93A
  *c.R1.IY = 0xCE38
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x15
  mem[0x0003] = 0x0F
  mem[0xCE4D] = 0x44
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2224) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2224, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2AC9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2AC9, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEC6B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEC6B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6511) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6511, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC93A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC93A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCE38) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCE38, *c.R1.IY)
  }
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

func TestFDCB10(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x259D
  *c.R1.BC = 0x3852
  *c.R1.DE = 0x590D
  *c.R1.HL = 0xAC66
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x144F
  *c.R1.IY = 0x42A2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7A
  mem[0x0003] = 0x10
  mem[0x431C] = 0x1C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x252C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x252C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3952) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3952, *c.R1.BC)
  }
  if (*c.R1.DE != 0x590D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x590D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAC66) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAC66, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x144F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x144F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x42A2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x42A2, *c.R1.IY)
  }
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

func TestFDCB11(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBC60
  *c.R1.BC = 0x61C1
  *c.R1.DE = 0xF5F8
  *c.R1.HL = 0xAF24
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4019
  *c.R1.IY = 0x9C90
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7B
  mem[0x0003] = 0x11
  mem[0x9D0B] = 0x5E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBCA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBCA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x61BC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x61BC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF5F8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF5F8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAF24) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAF24, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4019) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4019, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9C90) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9C90, *c.R1.IY)
  }
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

func TestFDCB12(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4E45
  *c.R1.BC = 0x3A25
  *c.R1.DE = 0x3417
  *c.R1.HL = 0xBCC7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0D7E
  *c.R1.IY = 0x8537
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x61
  mem[0x0003] = 0x12
  mem[0x8598] = 0xA7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4E09) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4E09, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3A25) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3A25, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4F17) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4F17, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBCC7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBCC7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0D7E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0D7E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8537) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8537, *c.R1.IY)
  }
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

func TestFDCB13(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB224
  *c.R1.BC = 0xB79B
  *c.R1.DE = 0x84F1
  *c.R1.HL = 0xFF7D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x414C
  *c.R1.IY = 0xE798
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB3
  mem[0x0003] = 0x13
  mem[0xE74B] = 0xB3
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB225) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB225, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB79B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB79B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8466) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8466, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFF7D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFF7D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x414C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x414C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE798) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE798, *c.R1.IY)
  }
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

func TestFDCB14(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xABBB
  *c.R1.BC = 0x451A
  *c.R1.DE = 0xFC65
  *c.R1.HL = 0x14A1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0F4D
  *c.R1.IY = 0xD93C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC4
  mem[0x0003] = 0x14
  mem[0xD900] = 0x06
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAB08) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAB08, *c.R1.AF)
  }
  if (*c.R1.BC != 0x451A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x451A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFC65) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFC65, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0DA1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0DA1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0F4D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0F4D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD93C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD93C, *c.R1.IY)
  }
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

func TestFDCB15(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2864
  *c.R1.BC = 0x9532
  *c.R1.DE = 0x8631
  *c.R1.HL = 0x751C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE327
  *c.R1.IY = 0x2D7B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x17
  mem[0x0003] = 0x15
  mem[0x2D92] = 0x12
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2824) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2824, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9532) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9532, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8631) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8631, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7524) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7524, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE327) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE327, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2D7B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2D7B, *c.R1.IY)
  }
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

func TestFDCB16(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0C3C
  *c.R1.BC = 0xDCD7
  *c.R1.DE = 0xADCC
  *c.R1.HL = 0x196D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x87E2
  *c.R1.IY = 0xF0B4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x23
  mem[0x0003] = 0x16
  mem[0xF0D7] = 0x89
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0C05) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0C05, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDCD7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDCD7, *c.R1.BC)
  }
  if (*c.R1.DE != 0xADCC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xADCC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x196D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x196D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x87E2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x87E2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF0B4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF0B4, *c.R1.IY)
  }
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

func TestFDCB17(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAF5B
  *c.R1.BC = 0xD016
  *c.R1.DE = 0x066E
  *c.R1.HL = 0x6638
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5E92
  *c.R1.IY = 0x2013
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8A
  mem[0x0003] = 0x17
  mem[0x1F9D] = 0xB8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7125) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7125, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD016) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD016, *c.R1.BC)
  }
  if (*c.R1.DE != 0x066E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x066E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6638) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6638, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5E92) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5E92, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2013) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2013, *c.R1.IY)
  }
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

func TestFDCB18(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x23F3
  *c.R1.BC = 0x4517
  *c.R1.DE = 0x16E0
  *c.R1.HL = 0x6894
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB908
  *c.R1.IY = 0x3216
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC0
  mem[0x0003] = 0x18
  mem[0x31D6] = 0xFA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x23A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x23A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFD17) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFD17, *c.R1.BC)
  }
  if (*c.R1.DE != 0x16E0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x16E0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6894) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6894, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB908) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB908, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3216) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3216, *c.R1.IY)
  }
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

func TestFDCB19(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x11ED
  *c.R1.BC = 0xC2B8
  *c.R1.DE = 0xA9F3
  *c.R1.HL = 0x2014
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6DB0
  *c.R1.IY = 0x4D2E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA2
  mem[0x0003] = 0x19
  mem[0x4CD0] = 0x4B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x11A5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x11A5, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC2A5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC2A5, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA9F3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA9F3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2014) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2014, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6DB0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6DB0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4D2E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4D2E, *c.R1.IY)
  }
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

func TestFDCB1A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBC5C
  *c.R1.BC = 0x6168
  *c.R1.DE = 0xE541
  *c.R1.HL = 0xB630
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0207
  *c.R1.IY = 0x40D3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x78
  mem[0x0003] = 0x1A
  mem[0x414B] = 0x44
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBC24) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBC24, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6168) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6168, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2241) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2241, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB630) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB630, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0207) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0207, *c.R1.IX)
  }
  if (*c.R1.IY != 0x40D3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x40D3, *c.R1.IY)
  }
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

func TestFDCB1B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7A28
  *c.R1.BC = 0x1286
  *c.R1.DE = 0xFE50
  *c.R1.HL = 0xC42D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE290
  *c.R1.IY = 0x71B0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x16
  mem[0x0003] = 0x1B
  mem[0x71C6] = 0xB8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7A0C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7A0C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1286) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1286, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFE5C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFE5C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC42D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC42D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE290) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE290, *c.R1.IX)
  }
  if (*c.R1.IY != 0x71B0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x71B0, *c.R1.IY)
  }
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

func TestFDCB1C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x932B
  *c.R1.BC = 0x097B
  *c.R1.DE = 0x6928
  *c.R1.HL = 0x83A3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFF2D
  *c.R1.IY = 0xDF62
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x86
  mem[0x0003] = 0x1C
  mem[0xDEE8] = 0x8F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9381) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9381, *c.R1.AF)
  }
  if (*c.R1.BC != 0x097B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x097B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6928) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6928, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC7A3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC7A3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFF2D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFF2D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDF62) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDF62, *c.R1.IY)
  }
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

func TestFDCB1D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x97B1
  *c.R1.BC = 0x2B30
  *c.R1.DE = 0x2645
  *c.R1.HL = 0x04EF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x186A
  *c.R1.IY = 0xD667
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x27
  mem[0x0003] = 0x1D
  mem[0xD68E] = 0xB7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x978D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x978D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2B30) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2B30, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2645) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2645, *c.R1.DE)
  }
  if (*c.R1.HL != 0x04DB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x04DB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x186A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x186A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD667) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD667, *c.R1.IY)
  }
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

func TestFDCB1E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2F39
  *c.R1.BC = 0x2470
  *c.R1.DE = 0xB521
  *c.R1.HL = 0x6CA3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1066
  *c.R1.IY = 0xDA38
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3A
  mem[0x0003] = 0x1E
  mem[0xDA72] = 0x25
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2F81) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2F81, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2470) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2470, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB521) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB521, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6CA3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6CA3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1066) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1066, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDA38) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDA38, *c.R1.IY)
  }
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

func TestFDCB1F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4CDD
  *c.R1.BC = 0x49A3
  *c.R1.DE = 0xDA18
  *c.R1.HL = 0x3AFD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA4F1
  *c.R1.IY = 0x2095
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7B
  mem[0x0003] = 0x1F
  mem[0x2110] = 0x04
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8284) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8284, *c.R1.AF)
  }
  if (*c.R1.BC != 0x49A3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x49A3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDA18) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDA18, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3AFD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3AFD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA4F1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA4F1, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2095) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2095, *c.R1.IY)
  }
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

func TestFDCB20(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3D74
  *c.R1.BC = 0x3A8F
  *c.R1.DE = 0x206F
  *c.R1.HL = 0x8894
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDDAB
  *c.R1.IY = 0xDA25
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7A
  mem[0x0003] = 0x20
  mem[0xDA9F] = 0x89
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3D05) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3D05, *c.R1.AF)
  }
  if (*c.R1.BC != 0x128F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x128F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x206F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x206F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8894) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8894, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDDAB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDDAB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDA25) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDA25, *c.R1.IY)
  }
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

func TestFDCB21(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1674
  *c.R1.BC = 0x6025
  *c.R1.DE = 0x641A
  *c.R1.HL = 0x6598
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x473B
  *c.R1.IY = 0xDE36
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7B
  mem[0x0003] = 0x21
  mem[0xDEB1] = 0x23
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1600, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6046) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6046, *c.R1.BC)
  }
  if (*c.R1.DE != 0x641A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x641A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6598) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6598, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
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
  if (*c.R1.IY != 0xDE36) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDE36, *c.R1.IY)
  }
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

func TestFDCB22(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xADA9
  *c.R1.BC = 0xEFB2
  *c.R1.DE = 0x6F03
  *c.R1.HL = 0xE732
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC11D
  *c.R1.IY = 0x8926
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9A
  mem[0x0003] = 0x22
  mem[0x88C0] = 0xD4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xADA9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xADA9, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEFB2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEFB2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA803) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA803, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE732) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE732, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC11D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC11D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8926) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8926, *c.R1.IY)
  }
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

func TestFDCB23(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x21E9
  *c.R1.BC = 0xD678
  *c.R1.DE = 0xA71B
  *c.R1.HL = 0x25D7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4CA8
  *c.R1.IY = 0x5255
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF5
  mem[0x0003] = 0x23
  mem[0x524A] = 0x65
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x218C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x218C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD678) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD678, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA7CA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA7CA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x25D7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x25D7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4CA8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4CA8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5255) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5255, *c.R1.IY)
  }
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

func TestFDCB24(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1C51
  *c.R1.BC = 0xDA3E
  *c.R1.DE = 0xCC7C
  *c.R1.HL = 0xCB19
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x572C
  *c.R1.IY = 0xAFFE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB4
  mem[0x0003] = 0x24
  mem[0xAFB2] = 0x7E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1CAC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1CAC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDA3E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDA3E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCC7C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCC7C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFC19) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFC19, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x572C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x572C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAFFE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAFFE, *c.R1.IY)
  }
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

func TestFDCB25(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x954E
  *c.R1.BC = 0x097C
  *c.R1.DE = 0xA341
  *c.R1.HL = 0x89E0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x435D
  *c.R1.IY = 0x23E9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA6
  mem[0x0003] = 0x25
  mem[0x238F] = 0x26
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9508) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9508, *c.R1.AF)
  }
  if (*c.R1.BC != 0x097C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x097C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA341) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA341, *c.R1.DE)
  }
  if (*c.R1.HL != 0x894C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x894C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x435D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x435D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x23E9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x23E9, *c.R1.IY)
  }
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

func TestFDCB26(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5844
  *c.R1.BC = 0x0E19
  *c.R1.DE = 0xD277
  *c.R1.HL = 0xBF7F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6504
  *c.R1.IY = 0xD4E4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBD
  mem[0x0003] = 0x26
  mem[0xD4A1] = 0xBF
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x582D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x582D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0E19) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0E19, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD277) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD277, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBF7F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBF7F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6504) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6504, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD4E4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD4E4, *c.R1.IY)
  }
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

func TestFDCB27(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8E0D
  *c.R1.BC = 0x8C06
  *c.R1.DE = 0x2C4C
  *c.R1.HL = 0xD7C8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9239
  *c.R1.IY = 0x8D42
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x59
  mem[0x0003] = 0x27
  mem[0x8D9B] = 0xA7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4E0D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4E0D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8C06) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8C06, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2C4C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2C4C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD7C8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD7C8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9239) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9239, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8D42) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8D42, *c.R1.IY)
  }
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

func TestFDCB28(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4122
  *c.R1.BC = 0xAF9B
  *c.R1.DE = 0x7745
  *c.R1.HL = 0x76F5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA1BB
  *c.R1.IY = 0xAB43
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x83
  mem[0x0003] = 0x28
  mem[0xAAC6] = 0x5D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x412D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x412D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2E9B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2E9B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7745) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7745, *c.R1.DE)
  }
  if (*c.R1.HL != 0x76F5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x76F5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA1BB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA1BB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAB43) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAB43, *c.R1.IY)
  }
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

func TestFDCB29(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0B21
  *c.R1.BC = 0xAFFD
  *c.R1.DE = 0xFEA6
  *c.R1.HL = 0x9478
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x32BB
  *c.R1.IY = 0x0343
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7D
  mem[0x0003] = 0x29
  mem[0x03C0] = 0x84
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0B80) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0B80, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAFC2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAFC2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFEA6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFEA6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9478) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9478, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x32BB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x32BB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0343) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0343, *c.R1.IY)
  }
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

func TestFDCB2A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF236
  *c.R1.BC = 0x8C31
  *c.R1.DE = 0x5932
  *c.R1.HL = 0x7FEB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7DB7
  *c.R1.IY = 0xABE7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF9
  mem[0x0003] = 0x2A
  mem[0xABE0] = 0xDD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF2AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF2AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8C31) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8C31, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEE32) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEE32, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7FEB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7FEB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7DB7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7DB7, *c.R1.IX)
  }
  if (*c.R1.IY != 0xABE7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xABE7, *c.R1.IY)
  }
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

func TestFDCB2B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2450
  *c.R1.BC = 0x6945
  *c.R1.DE = 0xDCFC
  *c.R1.HL = 0xD643
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5BE1
  *c.R1.IY = 0x4A94
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4B
  mem[0x0003] = 0x2B
  mem[0x4ADF] = 0x49
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2425) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2425, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6945) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6945, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDC24) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDC24, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD643) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD643, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5BE1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5BE1, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4A94) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4A94, *c.R1.IY)
  }
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

func TestFDCB2C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x117F
  *c.R1.BC = 0xB32B
  *c.R1.DE = 0xE530
  *c.R1.HL = 0x255A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2416
  *c.R1.IY = 0xCCD1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE6
  mem[0x0003] = 0x2C
  mem[0xCCB7] = 0x3C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x110C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x110C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB32B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB32B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE530) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE530, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1E5A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1E5A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2416) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2416, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCCD1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCCD1, *c.R1.IY)
  }
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

func TestFDCB2D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD0C3
  *c.R1.BC = 0x344B
  *c.R1.DE = 0x1BB0
  *c.R1.HL = 0x3EAB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFE11
  *c.R1.IY = 0xE4E6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5F
  mem[0x0003] = 0x2D
  mem[0xE545] = 0x78
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD02C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD02C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x344B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x344B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1BB0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1BB0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3E3C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3E3C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFE11) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFE11, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE4E6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE4E6, *c.R1.IY)
  }
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

func TestFDCB2E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF4EE
  *c.R1.BC = 0xB832
  *c.R1.DE = 0x4B7F
  *c.R1.HL = 0xE2B7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9386
  *c.R1.IY = 0x42FD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x06
  mem[0x0003] = 0x2E
  mem[0x4303] = 0xAD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF481) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF481, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB832) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB832, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4B7F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4B7F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE2B7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE2B7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9386) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9386, *c.R1.IX)
  }
  if (*c.R1.IY != 0x42FD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x42FD, *c.R1.IY)
  }
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

func TestFDCB2F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF86
  *c.R1.BC = 0xF2C2
  *c.R1.DE = 0x9F2F
  *c.R1.HL = 0xC946
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5FE0
  *c.R1.IY = 0x16B8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x29
  mem[0x0003] = 0x2F
  mem[0x16E1] = 0x18
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0C0C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0C0C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF2C2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF2C2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9F2F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9F2F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC946) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC946, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5FE0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5FE0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x16B8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x16B8, *c.R1.IY)
  }
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

func TestFDCB30(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xACF6
  *c.R1.BC = 0xE832
  *c.R1.DE = 0xF9ED
  *c.R1.HL = 0xCABC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFABD
  *c.R1.IY = 0xD646
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1B
  mem[0x0003] = 0x30
  mem[0xD661] = 0xA5
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAC0D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAC0D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4B32) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4B32, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF9ED) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF9ED, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCABC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCABC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFABD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFABD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD646) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD646, *c.R1.IY)
  }
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

func TestFDCB31(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2B96
  *c.R1.BC = 0x5134
  *c.R1.DE = 0x83A7
  *c.R1.HL = 0x7EEE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7750
  *c.R1.IY = 0xBFE0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF0
  mem[0x0003] = 0x31
  mem[0xBFD0] = 0xF1
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2BA1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2BA1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x51E3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x51E3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x83A7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x83A7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7EEE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7EEE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7750) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7750, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBFE0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBFE0, *c.R1.IY)
  }
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

func TestFDCB32(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB2BC
  *c.R1.BC = 0xA4B1
  *c.R1.DE = 0xB685
  *c.R1.HL = 0xF66E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA9A1
  *c.R1.IY = 0x5ADE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC5
  mem[0x0003] = 0x32
  mem[0x5AA3] = 0x59
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB2A0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB2A0, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA4B1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA4B1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB385) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB385, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF66E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF66E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA9A1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA9A1, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5ADE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5ADE, *c.R1.IY)
  }
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

func TestFDCB33(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9C6D
  *c.R1.BC = 0x2C90
  *c.R1.DE = 0xD0A9
  *c.R1.HL = 0x2BE3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2691
  *c.R1.IY = 0x1964
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7F
  mem[0x0003] = 0x33
  mem[0x19E3] = 0xDA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9CA1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9CA1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2C90) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2C90, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD0B5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD0B5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2BE3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2BE3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2691) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2691, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1964) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1964, *c.R1.IY)
  }
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

func TestFDCB34(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6029
  *c.R1.BC = 0xFBCD
  *c.R1.DE = 0x5348
  *c.R1.HL = 0xF947
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5338
  *c.R1.IY = 0x5696
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD2
  mem[0x0003] = 0x34
  mem[0x5668] = 0xD4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x60AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x60AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFBCD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFBCD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5348) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5348, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA947) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA947, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5338) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5338, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5696) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5696, *c.R1.IY)
  }
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

func TestFDCB35(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x96A9
  *c.R1.BC = 0x21C6
  *c.R1.DE = 0x4CB6
  *c.R1.HL = 0xB40B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x673A
  *c.R1.IY = 0x00F8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x71
  mem[0x0003] = 0x35
  mem[0x0169] = 0x0B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9604) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9604, *c.R1.AF)
  }
  if (*c.R1.BC != 0x21C6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x21C6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4CB6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4CB6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB417) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB417, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x673A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x673A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x00F8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x00F8, *c.R1.IY)
  }
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

func TestFDCB36(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDC6F
  *c.R1.BC = 0x0892
  *c.R1.DE = 0x3CC7
  *c.R1.HL = 0x1494
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8598
  *c.R1.IY = 0x1ADE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDA
  mem[0x0003] = 0x36
  mem[0x1AB8] = 0x3C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDC28) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDC28, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0892) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0892, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3CC7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3CC7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1494) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1494, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8598) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8598, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1ADE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1ADE, *c.R1.IY)
  }
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

func TestFDCB37(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD2B3
  *c.R1.BC = 0x4524
  *c.R1.DE = 0x208F
  *c.R1.HL = 0x076F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAD10
  *c.R1.IY = 0xE7EC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xCB
  mem[0x0003] = 0x37
  mem[0xE7B7] = 0x9F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3F2D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3F2D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4524) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4524, *c.R1.BC)
  }
  if (*c.R1.DE != 0x208F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x208F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x076F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x076F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAD10) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAD10, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE7EC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE7EC, *c.R1.IY)
  }
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

func TestFDCB38(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4F07
  *c.R1.BC = 0x0050
  *c.R1.DE = 0x40C6
  *c.R1.HL = 0x4FB7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF37E
  *c.R1.IY = 0xD096
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8E
  mem[0x0003] = 0x38
  mem[0xD024] = 0x0D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4F05) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4F05, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0650) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0650, *c.R1.BC)
  }
  if (*c.R1.DE != 0x40C6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x40C6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4FB7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4FB7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF37E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF37E, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD096) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD096, *c.R1.IY)
  }
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

func TestFDCB39(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBCC2
  *c.R1.BC = 0xF5B5
  *c.R1.DE = 0x8DEE
  *c.R1.HL = 0xE514
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x48BC
  *c.R1.IY = 0xF433
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7F
  mem[0x0003] = 0x39
  mem[0xF4B2] = 0xF5
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBC29) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBC29, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF57A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF57A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8DEE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8DEE, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE514) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE514, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x48BC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x48BC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF433) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF433, *c.R1.IY)
  }
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

func TestFDCB3A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD012
  *c.R1.BC = 0x2EF5
  *c.R1.DE = 0x2910
  *c.R1.HL = 0x9CA5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB155
  *c.R1.IY = 0xCB03
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1D
  mem[0x0003] = 0x3A
  mem[0xCB20] = 0xA8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2EF5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2EF5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5410) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5410, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9CA5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9CA5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB155) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB155, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCB03) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCB03, *c.R1.IY)
  }
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

func TestFDCB3B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x503D
  *c.R1.BC = 0xA85B
  *c.R1.DE = 0xCFBB
  *c.R1.HL = 0xDE8C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9C5B
  *c.R1.IY = 0xD263
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x05
  mem[0x0003] = 0x3B
  mem[0xD268] = 0xB2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x500C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x500C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA85B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA85B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCF59) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCF59, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDE8C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDE8C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9C5B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9C5B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD263) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD263, *c.R1.IY)
  }
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

func TestFDCB3C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x97F0
  *c.R1.BC = 0x4456
  *c.R1.DE = 0x0B52
  *c.R1.HL = 0xFDAD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6D2A
  *c.R1.IY = 0xA80F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAE
  mem[0x0003] = 0x3C
  mem[0xA7BD] = 0x96
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x970C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x970C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4456) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4456, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0B52) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0B52, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4BAD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4BAD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6D2A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6D2A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA80F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA80F, *c.R1.IY)
  }
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

func TestFDCB3D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7D44
  *c.R1.BC = 0x9303
  *c.R1.DE = 0xE12B
  *c.R1.HL = 0xBFF6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4C0F
  *c.R1.IY = 0xE52A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x13
  mem[0x0003] = 0x3D
  mem[0xE53D] = 0xFB
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7D2D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7D2D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9303) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9303, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE12B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE12B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBF7D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBF7D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4C0F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4C0F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE52A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE52A, *c.R1.IY)
  }
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

func TestFDCB3E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0D95
  *c.R1.BC = 0x3E02
  *c.R1.DE = 0x8F74
  *c.R1.HL = 0x0F82
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x85DF
  *c.R1.IY = 0xB2D1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2E
  mem[0x0003] = 0x3E
  mem[0xB2FF] = 0x50
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0D2C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0D2C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3E02) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3E02, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8F74) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8F74, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0F82) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0F82, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x85DF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x85DF, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB2D1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB2D1, *c.R1.IY)
  }
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

func TestFDCB3F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x89E3
  *c.R1.BC = 0x12F6
  *c.R1.DE = 0x426C
  *c.R1.HL = 0x52D4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD9F7
  *c.R1.IY = 0xC1AC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x21
  mem[0x0003] = 0x3F
  mem[0xC1CD] = 0x78
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C2C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C2C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x12F6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x12F6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x426C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x426C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x52D4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x52D4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD9F7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD9F7, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC1AC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC1AC, *c.R1.IY)
  }
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

func TestFDCB40(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5408
  *c.R1.BC = 0x2C34
  *c.R1.DE = 0x6784
  *c.R1.HL = 0xB376
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8FF9
  *c.R1.IY = 0x4195
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3B
  mem[0x0003] = 0x40
  mem[0x41D0] = 0x0D
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5410) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5410, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2C34) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2C34, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6784) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6784, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB376) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB376, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8FF9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8FF9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4195) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4195, *c.R1.IY)
  }
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

func TestFDCB41(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8C35
  *c.R1.BC = 0x5A58
  *c.R1.DE = 0xB71C
  *c.R1.HL = 0x6777
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDECA
  *c.R1.IY = 0x03CB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xCC
  mem[0x0003] = 0x41
  mem[0x0397] = 0xE9
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8C11) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8C11, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5A58) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5A58, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB71C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB71C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6777) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6777, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDECA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDECA, *c.R1.IX)
  }
  if (*c.R1.IY != 0x03CB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x03CB, *c.R1.IY)
  }
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

func TestFDCB42(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5535
  *c.R1.BC = 0x9C29
  *c.R1.DE = 0x2FEB
  *c.R1.HL = 0x97FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7F17
  *c.R1.IY = 0x9F56
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x01
  mem[0x0003] = 0x42
  mem[0x9F57] = 0xA8
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x555D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x555D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9C29) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9C29, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2FEB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2FEB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x97FF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x97FF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7F17) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7F17, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9F56) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9F56, *c.R1.IY)
  }
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

func TestFDCB43(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB404
  *c.R1.BC = 0xE58C
  *c.R1.DE = 0xE62E
  *c.R1.HL = 0x2A32
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7130
  *c.R1.IY = 0x1FD1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x08
  mem[0x0003] = 0x43
  mem[0x1FD9] = 0xAA
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB45C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB45C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE58C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE58C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE62E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE62E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2A32) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2A32, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7130) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7130, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1FD1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1FD1, *c.R1.IY)
  }
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

func TestFDCB44(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA954
  *c.R1.BC = 0x68F4
  *c.R1.DE = 0x9FA4
  *c.R1.HL = 0x7F66
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0209
  *c.R1.IY = 0xF4F3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x03
  mem[0x0003] = 0x44
  mem[0xF4F6] = 0x89
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA930) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA930, *c.R1.AF)
  }
  if (*c.R1.BC != 0x68F4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x68F4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9FA4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9FA4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7F66) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7F66, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0209) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0209, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF4F3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF4F3, *c.R1.IY)
  }
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

func TestFDCB45(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x73E5
  *c.R1.BC = 0x8DDE
  *c.R1.DE = 0x5E4F
  *c.R1.HL = 0x84A7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4E24
  *c.R1.IY = 0x93ED
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8D
  mem[0x0003] = 0x45
  mem[0x937A] = 0x8D
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7311) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7311, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8DDE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8DDE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5E4F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5E4F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x84A7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x84A7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4E24) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4E24, *c.R1.IX)
  }
  if (*c.R1.IY != 0x93ED) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x93ED, *c.R1.IY)
  }
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

func TestFDCB46(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0E5A
  *c.R1.BC = 0xB1F9
  *c.R1.DE = 0x475F
  *c.R1.HL = 0xEBFC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7765
  *c.R1.IY = 0x63B1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8C
  mem[0x0003] = 0x46
  mem[0x633D] = 0xFE
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0E74) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0E74, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB1F9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB1F9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x475F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x475F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEBFC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEBFC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7765) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7765, *c.R1.IX)
  }
  if (*c.R1.IY != 0x63B1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x63B1, *c.R1.IY)
  }
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

func TestFDCB47(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9B3D
  *c.R1.BC = 0x7F38
  *c.R1.DE = 0x0753
  *c.R1.HL = 0xD5E7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB9C3
  *c.R1.IY = 0x6E0E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x96
  mem[0x0003] = 0x47
  mem[0x6DA4] = 0xD6
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9B7D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9B7D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7F38) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7F38, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0753) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0753, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD5E7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD5E7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB9C3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB9C3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6E0E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6E0E, *c.R1.IY)
  }
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

func TestFDCB48(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7D94
  *c.R1.BC = 0x50A9
  *c.R1.DE = 0x2511
  *c.R1.HL = 0x8F9F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB612
  *c.R1.IY = 0xABA9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x44
  mem[0x0003] = 0x48
  mem[0xABED] = 0xB0
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7D7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7D7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x50A9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x50A9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2511) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2511, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8F9F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8F9F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB612) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB612, *c.R1.IX)
  }
  if (*c.R1.IY != 0xABA9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xABA9, *c.R1.IY)
  }
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

func TestFDCB49(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x691E
  *c.R1.BC = 0x3A39
  *c.R1.DE = 0xB834
  *c.R1.HL = 0x74B6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0EB7
  *c.R1.IY = 0x3E21
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4E
  mem[0x0003] = 0x49
  mem[0x3E6F] = 0xA9
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x697C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x697C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3A39) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3A39, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB834) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB834, *c.R1.DE)
  }
  if (*c.R1.HL != 0x74B6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x74B6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0EB7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0EB7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3E21) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3E21, *c.R1.IY)
  }
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

func TestFDCB4A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x31E3
  *c.R1.BC = 0x68E0
  *c.R1.DE = 0xFE2F
  *c.R1.HL = 0xA2C4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAC96
  *c.R1.IY = 0xE7DB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x52
  mem[0x0003] = 0x4A
  mem[0xE82D] = 0xDA
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3139) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3139, *c.R1.AF)
  }
  if (*c.R1.BC != 0x68E0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x68E0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFE2F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFE2F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA2C4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA2C4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAC96) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAC96, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE7DB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE7DB, *c.R1.IY)
  }
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

func TestFDCB4B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x09A1
  *c.R1.BC = 0x2453
  *c.R1.DE = 0x9186
  *c.R1.HL = 0xA32A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x71AF
  *c.R1.IY = 0x883F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEA
  mem[0x0003] = 0x4B
  mem[0x8829] = 0x4E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0919) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0919, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2453) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2453, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9186) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9186, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA32A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA32A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x71AF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x71AF, *c.R1.IX)
  }
  if (*c.R1.IY != 0x883F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x883F, *c.R1.IY)
  }
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

func TestFDCB4C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4A52
  *c.R1.BC = 0x1E5B
  *c.R1.DE = 0xBE2E
  *c.R1.HL = 0x3EE4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAF79
  *c.R1.IY = 0x7F22
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEE
  mem[0x0003] = 0x4C
  mem[0x7F10] = 0x70
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4A7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4A7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1E5B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1E5B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBE2E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBE2E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3EE4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3EE4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAF79) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAF79, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7F22) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7F22, *c.R1.IY)
  }
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

func TestFDCB4D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9F87
  *c.R1.BC = 0x6C8F
  *c.R1.DE = 0x34F4
  *c.R1.HL = 0x5A79
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD3CC
  *c.R1.IY = 0xA770
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x29
  mem[0x0003] = 0x4D
  mem[0xA799] = 0x78
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9F75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9F75, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6C8F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6C8F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x34F4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x34F4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5A79) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5A79, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD3CC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD3CC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA770) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA770, *c.R1.IY)
  }
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

func TestFDCB4E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x30CB
  *c.R1.BC = 0x5626
  *c.R1.DE = 0x52BC
  *c.R1.HL = 0x5503
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x303B
  *c.R1.IY = 0xE1C8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x20
  mem[0x0003] = 0x4E
  mem[0xE1E8] = 0xAA
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3031) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3031, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5626) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5626, *c.R1.BC)
  }
  if (*c.R1.DE != 0x52BC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x52BC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5503) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5503, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x303B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x303B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE1C8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE1C8, *c.R1.IY)
  }
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

func TestFDCB4F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6088
  *c.R1.BC = 0xE079
  *c.R1.DE = 0x7152
  *c.R1.HL = 0x671F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8C22
  *c.R1.IY = 0x1CF8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9D
  mem[0x0003] = 0x4F
  mem[0x1C95] = 0x18
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x605C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x605C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE079) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE079, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7152) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7152, *c.R1.DE)
  }
  if (*c.R1.HL != 0x671F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x671F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8C22) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8C22, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1CF8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1CF8, *c.R1.IY)
  }
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

func TestFDCB50(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8CDE
  *c.R1.BC = 0x1409
  *c.R1.DE = 0x6D69
  *c.R1.HL = 0xE5B2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4A0C
  *c.R1.IY = 0xC75F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6B
  mem[0x0003] = 0x50
  mem[0xC7CA] = 0xFE
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8C10) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8C10, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1409) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1409, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6D69) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6D69, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE5B2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE5B2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4A0C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4A0C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC75F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC75F, *c.R1.IY)
  }
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

func TestFDCB51(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8F59
  *c.R1.BC = 0x40CB
  *c.R1.DE = 0x9543
  *c.R1.HL = 0x9B3A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1942
  *c.R1.IY = 0x3495
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x86
  mem[0x0003] = 0x51
  mem[0x341B] = 0x13
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8F75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8F75, *c.R1.AF)
  }
  if (*c.R1.BC != 0x40CB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x40CB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9543) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9543, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9B3A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9B3A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1942) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1942, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3495) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3495, *c.R1.IY)
  }
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

func TestFDCB52(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8905
  *c.R1.BC = 0x3E41
  *c.R1.DE = 0x7AB4
  *c.R1.HL = 0x37F6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF82D
  *c.R1.IY = 0x8B0D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE6
  mem[0x0003] = 0x52
  mem[0x8AF3] = 0x87
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8919) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8919, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3E41) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3E41, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7AB4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7AB4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x37F6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x37F6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF82D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF82D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8B0D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8B0D, *c.R1.IY)
  }
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

func TestFDCB53(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEFDE
  *c.R1.BC = 0xE345
  *c.R1.DE = 0x09A3
  *c.R1.HL = 0xF0B2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC378
  *c.R1.IY = 0x7EE1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD1
  mem[0x0003] = 0x53
  mem[0x7EB2] = 0xE4
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEF38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEF38, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE345) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE345, *c.R1.BC)
  }
  if (*c.R1.DE != 0x09A3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x09A3, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF0B2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF0B2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC378) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC378, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7EE1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7EE1, *c.R1.IY)
  }
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

func TestFDCB54(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x72A6
  *c.R1.BC = 0xCB82
  *c.R1.DE = 0xD966
  *c.R1.HL = 0x2FC6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3C00
  *c.R1.IY = 0x5B6B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x08
  mem[0x0003] = 0x54
  mem[0x5B73] = 0x07
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7218) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7218, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCB82) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCB82, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD966) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD966, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2FC6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2FC6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3C00) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3C00, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5B6B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5B6B, *c.R1.IY)
  }
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

func TestFDCB55(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x855C
  *c.R1.BC = 0xC23B
  *c.R1.DE = 0x6AAB
  *c.R1.HL = 0x9B00
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFE93
  *c.R1.IY = 0xB4B2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x54
  mem[0x0003] = 0x55
  mem[0xB506] = 0x46
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8530) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8530, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC23B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC23B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6AAB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6AAB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9B00) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9B00, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFE93) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFE93, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB4B2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB4B2, *c.R1.IY)
  }
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

func TestFDCB56(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF5AD
  *c.R1.BC = 0xF9F6
  *c.R1.DE = 0x1E8C
  *c.R1.HL = 0x9E08
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x716A
  *c.R1.IY = 0x6932
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6F
  mem[0x0003] = 0x56
  mem[0x69A1] = 0xDF
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF539) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF539, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF9F6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF9F6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E8C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E8C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9E08) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9E08, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x716A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x716A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6932) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6932, *c.R1.IY)
  }
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

func TestFDCB57(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x37D7
  *c.R1.BC = 0xB7DC
  *c.R1.DE = 0xBE1C
  *c.R1.HL = 0x38EA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5E82
  *c.R1.IY = 0xA3BB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3C
  mem[0x0003] = 0x57
  mem[0xA3F7] = 0x6C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3731) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3731, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB7DC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB7DC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBE1C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBE1C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x38EA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x38EA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5E82) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5E82, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA3BB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA3BB, *c.R1.IY)
  }
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

func TestFDCB58(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x752C
  *c.R1.BC = 0x7296
  *c.R1.DE = 0x3EA5
  *c.R1.HL = 0x1143
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD7CC
  *c.R1.IY = 0x1E94
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4E
  mem[0x0003] = 0x58
  mem[0x1EE2] = 0xF6
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x755C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x755C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7296) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7296, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3EA5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3EA5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1143) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1143, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD7CC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD7CC, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1E94) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1E94, *c.R1.IY)
  }
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

func TestFDCB59(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8056
  *c.R1.BC = 0xBF2A
  *c.R1.DE = 0x1809
  *c.R1.HL = 0xED31
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFE2B
  *c.R1.IY = 0xFAD3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2E
  mem[0x0003] = 0x59
  mem[0xFB01] = 0x6F
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8038) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8038, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBF2A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBF2A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1809) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1809, *c.R1.DE)
  }
  if (*c.R1.HL != 0xED31) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xED31, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFE2B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFE2B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFAD3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFAD3, *c.R1.IY)
  }
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

func TestFDCB5A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCC74
  *c.R1.BC = 0xA108
  *c.R1.DE = 0x65D4
  *c.R1.HL = 0x6F66
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0008
  *c.R1.IY = 0x7BB8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x88
  mem[0x0003] = 0x5A
  mem[0x7B40] = 0x6E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCC38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCC38, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA108) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA108, *c.R1.BC)
  }
  if (*c.R1.DE != 0x65D4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x65D4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6F66) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6F66, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0008) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0008, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7BB8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7BB8, *c.R1.IY)
  }
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

func TestFDCB5B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5CF1
  *c.R1.BC = 0xB3BD
  *c.R1.DE = 0x25BD
  *c.R1.HL = 0x98CF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2BA1
  *c.R1.IY = 0x315C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE7
  mem[0x0003] = 0x5B
  mem[0x3143] = 0xB1
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5C75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5C75, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB3BD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB3BD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x25BD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x25BD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x98CF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x98CF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2BA1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2BA1, *c.R1.IX)
  }
  if (*c.R1.IY != 0x315C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x315C, *c.R1.IY)
  }
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

func TestFDCB5C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB3E0
  *c.R1.BC = 0xD43D
  *c.R1.DE = 0xD9C0
  *c.R1.HL = 0xB04D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x21A9
  *c.R1.IY = 0x543E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x74
  mem[0x0003] = 0x5C
  mem[0x54B2] = 0xE3
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB354) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB354, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD43D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD43D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD9C0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD9C0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB04D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB04D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x21A9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x21A9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x543E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x543E, *c.R1.IY)
  }
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

func TestFDCB5D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9F49
  *c.R1.BC = 0x43DD
  *c.R1.DE = 0xCCB3
  *c.R1.HL = 0x085A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF130
  *c.R1.IY = 0x3B84
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDC
  mem[0x0003] = 0x5D
  mem[0x3B60] = 0xEF
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9F39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9F39, *c.R1.AF)
  }
  if (*c.R1.BC != 0x43DD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x43DD, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCCB3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCCB3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x085A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x085A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF130) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF130, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3B84) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3B84, *c.R1.IY)
  }
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

func TestFDCB5E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6F89
  *c.R1.BC = 0xEFF5
  *c.R1.DE = 0x993B
  *c.R1.HL = 0x22B5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0F30
  *c.R1.IY = 0xE165
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE2
  mem[0x0003] = 0x5E
  mem[0xE147] = 0x17
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6F75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6F75, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEFF5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEFF5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x993B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x993B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x22B5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x22B5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0F30) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0F30, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE165) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE165, *c.R1.IY)
  }
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

func TestFDCB5F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD72A
  *c.R1.BC = 0xA57A
  *c.R1.DE = 0xACA6
  *c.R1.HL = 0x667E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5C33
  *c.R1.IY = 0xF81B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAB
  mem[0x0003] = 0x5F
  mem[0xF7C6] = 0xE2
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD774) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD774, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA57A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA57A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xACA6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xACA6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x667E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x667E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5C33) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5C33, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF81B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF81B, *c.R1.IY)
  }
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

func TestFDCB60(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x15E9
  *c.R1.BC = 0x8D30
  *c.R1.DE = 0x43F4
  *c.R1.HL = 0xC65E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1E34
  *c.R1.IY = 0x8C44
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x85
  mem[0x0003] = 0x60
  mem[0x8BC9] = 0xB9
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1519) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1519, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8D30) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8D30, *c.R1.BC)
  }
  if (*c.R1.DE != 0x43F4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x43F4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC65E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC65E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1E34) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1E34, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8C44) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8C44, *c.R1.IY)
  }
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

func TestFDCB61(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7BD1
  *c.R1.BC = 0xD421
  *c.R1.DE = 0x5570
  *c.R1.HL = 0xCB85
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x32EC
  *c.R1.IY = 0x92E4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBE
  mem[0x0003] = 0x61
  mem[0x92A2] = 0x28
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7B55) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7B55, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD421) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD421, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5570) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5570, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCB85) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCB85, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x32EC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x32EC, *c.R1.IX)
  }
  if (*c.R1.IY != 0x92E4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x92E4, *c.R1.IY)
  }
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

func TestFDCB62(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBA2F
  *c.R1.BC = 0x4FBB
  *c.R1.DE = 0x67A7
  *c.R1.HL = 0xC5DB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x470B
  *c.R1.IY = 0x7EB1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9D
  mem[0x0003] = 0x62
  mem[0x7E4E] = 0x1A
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBA39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBA39, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4FBB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4FBB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x67A7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x67A7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC5DB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC5DB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x470B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x470B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7EB1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7EB1, *c.R1.IY)
  }
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

func TestFDCB63(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC0A1
  *c.R1.BC = 0x2CC2
  *c.R1.DE = 0xCE12
  *c.R1.HL = 0xE77C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x71C5
  *c.R1.IY = 0x1713
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF4
  mem[0x0003] = 0x63
  mem[0x1707] = 0x3B
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC011) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC011, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2CC2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2CC2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCE12) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCE12, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE77C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE77C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x71C5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x71C5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1713) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1713, *c.R1.IY)
  }
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

func TestFDCB64(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0C1F
  *c.R1.BC = 0x7847
  *c.R1.DE = 0x2494
  *c.R1.HL = 0x71EB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x315C
  *c.R1.IY = 0xB336
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x35
  mem[0x0003] = 0x64
  mem[0xB36B] = 0x8C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0C75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0C75, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7847) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7847, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2494) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2494, *c.R1.DE)
  }
  if (*c.R1.HL != 0x71EB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x71EB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x315C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x315C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB336) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB336, *c.R1.IY)
  }
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

func TestFDCB65(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5245
  *c.R1.BC = 0xA82D
  *c.R1.DE = 0x1112
  *c.R1.HL = 0x8F09
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x672A
  *c.R1.IY = 0x89F4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x37
  mem[0x0003] = 0x65
  mem[0x8A2B] = 0x08
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x525D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x525D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA82D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA82D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1112) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1112, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8F09) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8F09, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x672A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x672A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x89F4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x89F4, *c.R1.IY)
  }
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

func TestFDCB66(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x583F
  *c.R1.BC = 0xC13E
  *c.R1.DE = 0xB136
  *c.R1.HL = 0x6BC5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3EF9
  *c.R1.IY = 0x6948
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9D
  mem[0x0003] = 0x66
  mem[0x68E5] = 0x90
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5839) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5839, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC13E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC13E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB136) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB136, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6BC5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6BC5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3EF9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3EF9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6948) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6948, *c.R1.IY)
  }
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

func TestFDCB67(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x31B6
  *c.R1.BC = 0x0F7D
  *c.R1.DE = 0x48B5
  *c.R1.HL = 0xCC5F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2103
  *c.R1.IY = 0x6572
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xCB
  mem[0x0003] = 0x67
  mem[0x653D] = 0x15
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3130) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3130, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F7D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F7D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x48B5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x48B5, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCC5F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCC5F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2103) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2103, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6572) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6572, *c.R1.IY)
  }
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

func TestFDCB68(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE330
  *c.R1.BC = 0x39FB
  *c.R1.DE = 0xA03A
  *c.R1.HL = 0x59BC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE04A
  *c.R1.IY = 0x03BE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xCA
  mem[0x0003] = 0x68
  mem[0x0388] = 0x83
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE354) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE354, *c.R1.AF)
  }
  if (*c.R1.BC != 0x39FB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x39FB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA03A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA03A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x59BC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x59BC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE04A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE04A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x03BE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x03BE, *c.R1.IY)
  }
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

func TestFDCB69(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1896
  *c.R1.BC = 0x5BC2
  *c.R1.DE = 0xD4D9
  *c.R1.HL = 0x4E8A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3716
  *c.R1.IY = 0xA603
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE2
  mem[0x0003] = 0x69
  mem[0xA5E5] = 0x01
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1874) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1874, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5BC2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5BC2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD4D9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD4D9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4E8A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4E8A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3716) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3716, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA603) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA603, *c.R1.IY)
  }
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

func TestFDCB6A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5BC9
  *c.R1.BC = 0x0099
  *c.R1.DE = 0x34F8
  *c.R1.HL = 0x3E96
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF251
  *c.R1.IY = 0x93BE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAE
  mem[0x0003] = 0x6A
  mem[0x936C] = 0x33
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5B11) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5B11, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0099) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0099, *c.R1.BC)
  }
  if (*c.R1.DE != 0x34F8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x34F8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3E96) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3E96, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF251) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF251, *c.R1.IX)
  }
  if (*c.R1.IY != 0x93BE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x93BE, *c.R1.IY)
  }
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

func TestFDCB6B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBBE5
  *c.R1.BC = 0x9E6C
  *c.R1.DE = 0xABD1
  *c.R1.HL = 0x515F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x73DB
  *c.R1.IY = 0xAA2F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1F
  mem[0x0003] = 0x6B
  mem[0xAA4E] = 0x7C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBB39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBB39, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9E6C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9E6C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xABD1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xABD1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x515F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x515F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x73DB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x73DB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAA2F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAA2F, *c.R1.IY)
  }
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

func TestFDCB6C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x144B
  *c.R1.BC = 0x3AF2
  *c.R1.DE = 0x8F80
  *c.R1.HL = 0x7BE5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC379
  *c.R1.IY = 0x86BA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0D
  mem[0x0003] = 0x6C
  mem[0x86C7] = 0x25
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1411) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1411, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3AF2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3AF2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8F80) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8F80, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7BE5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7BE5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC379) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC379, *c.R1.IX)
  }
  if (*c.R1.IY != 0x86BA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x86BA, *c.R1.IY)
  }
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

func TestFDCB6D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6392
  *c.R1.BC = 0xD077
  *c.R1.DE = 0x668D
  *c.R1.HL = 0x6E4A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB0A8
  *c.R1.IY = 0x62C8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF0
  mem[0x0003] = 0x6D
  mem[0x62B8] = 0xE3
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6330) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6330, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD077) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD077, *c.R1.BC)
  }
  if (*c.R1.DE != 0x668D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x668D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6E4A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6E4A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB0A8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB0A8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x62C8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x62C8, *c.R1.IY)
  }
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

func TestFDCB6E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2DA0
  *c.R1.BC = 0xF872
  *c.R1.DE = 0x692D
  *c.R1.HL = 0x92C4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x36B5
  *c.R1.IY = 0x4210
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x93
  mem[0x0003] = 0x6E
  mem[0x41A3] = 0x1E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2D54) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2D54, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF872) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF872, *c.R1.BC)
  }
  if (*c.R1.DE != 0x692D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x692D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x92C4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x92C4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x36B5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x36B5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4210) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4210, *c.R1.IY)
  }
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

func TestFDCB6F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDF7B
  *c.R1.BC = 0xC7AA
  *c.R1.DE = 0x9002
  *c.R1.HL = 0x86B8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1347
  *c.R1.IY = 0x004E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x20
  mem[0x0003] = 0x6F
  mem[0x006E] = 0x37
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDF11) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDF11, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC7AA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC7AA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9002) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9002, *c.R1.DE)
  }
  if (*c.R1.HL != 0x86B8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x86B8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1347) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1347, *c.R1.IX)
  }
  if (*c.R1.IY != 0x004E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x004E, *c.R1.IY)
  }
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

func TestFDCB70(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6EA9
  *c.R1.BC = 0x018D
  *c.R1.DE = 0x5075
  *c.R1.HL = 0xCF4E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCD2B
  *c.R1.IY = 0x3E68
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD9
  mem[0x0003] = 0x70
  mem[0x3E41] = 0xC9
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6E39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6E39, *c.R1.AF)
  }
  if (*c.R1.BC != 0x018D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x018D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5075) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5075, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCF4E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCF4E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCD2B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCD2B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3E68) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3E68, *c.R1.IY)
  }
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

func TestFDCB71(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1B48
  *c.R1.BC = 0xE3AF
  *c.R1.DE = 0x94D5
  *c.R1.HL = 0x0996
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCAD5
  *c.R1.IY = 0x999A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x27
  mem[0x0003] = 0x71
  mem[0x99C1] = 0x3E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1B5C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1B5C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE3AF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE3AF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x94D5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x94D5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0996) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0996, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCAD5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCAD5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x999A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x999A, *c.R1.IY)
  }
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

func TestFDCB72(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE83B
  *c.R1.BC = 0x26B1
  *c.R1.DE = 0x8608
  *c.R1.HL = 0xF3CB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6323
  *c.R1.IY = 0xFD31
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x98
  mem[0x0003] = 0x72
  mem[0xFCC9] = 0x4F
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE839) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE839, *c.R1.AF)
  }
  if (*c.R1.BC != 0x26B1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x26B1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8608) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8608, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF3CB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF3CB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6323) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6323, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFD31) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFD31, *c.R1.IY)
  }
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

func TestFDCB73(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x101B
  *c.R1.BC = 0x446C
  *c.R1.DE = 0xC2F9
  *c.R1.HL = 0xB9B1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0820
  *c.R1.IY = 0xF5D8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7A
  mem[0x0003] = 0x73
  mem[0xF652] = 0x31
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1075) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1075, *c.R1.AF)
  }
  if (*c.R1.BC != 0x446C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x446C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC2F9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC2F9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB9B1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB9B1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0820) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0820, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF5D8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF5D8, *c.R1.IY)
  }
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

func TestFDCB74(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6847
  *c.R1.BC = 0x38C2
  *c.R1.DE = 0x0EA4
  *c.R1.HL = 0x0825
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD255
  *c.R1.IY = 0x5E4A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4B
  mem[0x0003] = 0x74
  mem[0x5E95] = 0xFE
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6819) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6819, *c.R1.AF)
  }
  if (*c.R1.BC != 0x38C2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x38C2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0EA4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0EA4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0825) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0825, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD255) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD255, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5E4A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5E4A, *c.R1.IY)
  }
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

func TestFDCB75(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x56F2
  *c.R1.BC = 0xC034
  *c.R1.DE = 0x6E11
  *c.R1.HL = 0xD35E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE702
  *c.R1.IY = 0x60BE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x57
  mem[0x0003] = 0x75
  mem[0x6115] = 0x21
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5674) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5674, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC034) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC034, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6E11) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6E11, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD35E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD35E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE702) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE702, *c.R1.IX)
  }
  if (*c.R1.IY != 0x60BE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x60BE, *c.R1.IY)
  }
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

func TestFDCB76(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7375
  *c.R1.BC = 0xCAFF
  *c.R1.DE = 0xDD80
  *c.R1.HL = 0xC8ED
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7E39
  *c.R1.IY = 0x6623
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x53
  mem[0x0003] = 0x76
  mem[0x6676] = 0x3A
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7375) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7375, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCAFF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCAFF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDD80) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDD80, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC8ED) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC8ED, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7E39) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7E39, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6623) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6623, *c.R1.IY)
  }
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

func TestFDCB77(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAB10
  *c.R1.BC = 0x983E
  *c.R1.DE = 0x0BDC
  *c.R1.HL = 0x3B46
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAE51
  *c.R1.IY = 0x8841
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x02
  mem[0x0003] = 0x77
  mem[0x8843] = 0xD8
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAB18) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAB18, *c.R1.AF)
  }
  if (*c.R1.BC != 0x983E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x983E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0BDC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0BDC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3B46) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3B46, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAE51) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAE51, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8841) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8841, *c.R1.IY)
  }
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

func TestFDCB78(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2765
  *c.R1.BC = 0xCE2F
  *c.R1.DE = 0x4824
  *c.R1.HL = 0x6930
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAE69
  *c.R1.IY = 0xFECB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7D
  mem[0x0003] = 0x78
  mem[0xFF48] = 0xEC
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x27B9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x27B9, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCE2F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCE2F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4824) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4824, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6930) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6930, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAE69) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAE69, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFECB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFECB, *c.R1.IY)
  }
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

func TestFDCB79(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB428
  *c.R1.BC = 0x6355
  *c.R1.DE = 0x7896
  *c.R1.HL = 0x8A7C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9090
  *c.R1.IY = 0x1CAE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x23
  mem[0x0003] = 0x79
  mem[0x1CD1] = 0x87
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB498) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB498, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6355) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6355, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7896) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7896, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8A7C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8A7C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9090) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9090, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1CAE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1CAE, *c.R1.IY)
  }
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

func TestFDCB7A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x59F4
  *c.R1.BC = 0xCA21
  *c.R1.DE = 0x1482
  *c.R1.HL = 0x3FAE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC6C9
  *c.R1.IY = 0xD923
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x42
  mem[0x0003] = 0x7A
  mem[0xD965] = 0xB3
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5998) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5998, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCA21) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCA21, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1482) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1482, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3FAE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3FAE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC6C9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC6C9, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD923) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD923, *c.R1.IY)
  }
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

func TestFDCB7B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6314
  *c.R1.BC = 0x0240
  *c.R1.DE = 0x5EFA
  *c.R1.HL = 0x5E7B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3E50
  *c.R1.IY = 0x0A83
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x17
  mem[0x0003] = 0x7B
  mem[0x0A9A] = 0xBD
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6398) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6398, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0240) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0240, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5EFA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5EFA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E7B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E7B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3E50) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3E50, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0A83) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0A83, *c.R1.IY)
  }
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

func TestFDCB7C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x22A6
  *c.R1.BC = 0xAFF4
  *c.R1.DE = 0xB89B
  *c.R1.HL = 0x4DCA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0AC2
  *c.R1.IY = 0xD371
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF1
  mem[0x0003] = 0x7C
  mem[0xD362] = 0x1B
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2254) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2254, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAFF4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAFF4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB89B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB89B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4DCA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4DCA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0AC2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0AC2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD371) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD371, *c.R1.IY)
  }
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

func TestFDCB7D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1C95
  *c.R1.BC = 0xD615
  *c.R1.DE = 0x825A
  *c.R1.HL = 0x5E64
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x32FB
  *c.R1.IY = 0xAC3B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9F
  mem[0x0003] = 0x7D
  mem[0xABDA] = 0x8A
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1CB9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1CB9, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD615) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD615, *c.R1.BC)
  }
  if (*c.R1.DE != 0x825A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x825A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E64) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E64, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x32FB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x32FB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAC3B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAC3B, *c.R1.IY)
  }
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

func TestFDCB7E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x503C
  *c.R1.BC = 0x8DFE
  *c.R1.DE = 0x1019
  *c.R1.HL = 0x6778
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF7DF
  *c.R1.IY = 0x9484
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x40
  mem[0x0003] = 0x7E
  mem[0x94C4] = 0x9E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5090) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5090, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8DFE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8DFE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1019) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1019, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6778) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6778, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF7DF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF7DF, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9484) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9484, *c.R1.IY)
  }
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

func TestFDCB7F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1B07
  *c.R1.BC = 0x9EC3
  *c.R1.DE = 0x14BE
  *c.R1.HL = 0x5EBE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1178
  *c.R1.IY = 0xCE69
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA2
  mem[0x0003] = 0x7F
  mem[0xCE0B] = 0x47
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1B5D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1B5D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9EC3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9EC3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x14BE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x14BE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5EBE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5EBE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1178) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1178, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCE69) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCE69, *c.R1.IY)
  }
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

func TestFDCB80(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE196
  *c.R1.BC = 0x72EA
  *c.R1.DE = 0x507E
  *c.R1.HL = 0x6457
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAB75
  *c.R1.IY = 0x920D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8B
  mem[0x0003] = 0x80
  mem[0x9198] = 0xA9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE196) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE196, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA8EA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA8EA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x507E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x507E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6457) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6457, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAB75) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAB75, *c.R1.IX)
  }
  if (*c.R1.IY != 0x920D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x920D, *c.R1.IY)
  }
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

func TestFDCB81(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3D3D
  *c.R1.BC = 0xB255
  *c.R1.DE = 0x8759
  *c.R1.HL = 0x0CB0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE078
  *c.R1.IY = 0x82A5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x55
  mem[0x0003] = 0x81
  mem[0x82FA] = 0xFA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3D3D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3D3D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB2FA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB2FA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8759) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8759, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0CB0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0CB0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE078) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE078, *c.R1.IX)
  }
  if (*c.R1.IY != 0x82A5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x82A5, *c.R1.IY)
  }
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

func TestFDCB82(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4E10
  *c.R1.BC = 0x5D8D
  *c.R1.DE = 0x27A0
  *c.R1.HL = 0xFFFF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEE0A
  *c.R1.IY = 0x5DD8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9C
  mem[0x0003] = 0x82
  mem[0x5D74] = 0x9D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4E10) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4E10, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5D8D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5D8D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9CA0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9CA0, *c.R1.DE)
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
  if (*c.R1.IX != 0xEE0A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEE0A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5DD8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5DD8, *c.R1.IY)
  }
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

func TestFDCB83(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3C7F
  *c.R1.BC = 0xFD81
  *c.R1.DE = 0x47FB
  *c.R1.HL = 0x9F12
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCBF9
  *c.R1.IY = 0x374A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x28
  mem[0x0003] = 0x83
  mem[0x3772] = 0xD5
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C7F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C7F, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFD81) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFD81, *c.R1.BC)
  }
  if (*c.R1.DE != 0x47D4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x47D4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9F12) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9F12, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCBF9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCBF9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x374A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x374A, *c.R1.IY)
  }
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

func TestFDCB84(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6872
  *c.R1.BC = 0x81B1
  *c.R1.DE = 0x1E7A
  *c.R1.HL = 0xE37E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9B4C
  *c.R1.IY = 0xF1C3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAA
  mem[0x0003] = 0x84
  mem[0xF16D] = 0xEA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6872) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6872, *c.R1.AF)
  }
  if (*c.R1.BC != 0x81B1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x81B1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E7A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E7A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEA7E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEA7E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9B4C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9B4C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF1C3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF1C3, *c.R1.IY)
  }
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

func TestFDCB85(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x25B3
  *c.R1.BC = 0x5694
  *c.R1.DE = 0x57CD
  *c.R1.HL = 0xF34D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8ED2
  *c.R1.IY = 0x0433
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6C
  mem[0x0003] = 0x85
  mem[0x049F] = 0xE0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x25B3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x25B3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5694) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5694, *c.R1.BC)
  }
  if (*c.R1.DE != 0x57CD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x57CD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF3E0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF3E0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8ED2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8ED2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0433) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0433, *c.R1.IY)
  }
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

func TestFDCB86(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x152B
  *c.R1.BC = 0x8CE1
  *c.R1.DE = 0x818D
  *c.R1.HL = 0x40F2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9B7A
  *c.R1.IY = 0x2A50
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7E
  mem[0x0003] = 0x86
  mem[0x2ACE] = 0x36
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x152B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x152B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8CE1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8CE1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x818D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x818D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x40F2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x40F2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9B7A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9B7A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2A50) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2A50, *c.R1.IY)
  }
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

func TestFDCB87(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFE1D
  *c.R1.BC = 0x5353
  *c.R1.DE = 0x618D
  *c.R1.HL = 0x3266
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1A53
  *c.R1.IY = 0x246A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x59
  mem[0x0003] = 0x87
  mem[0x24C3] = 0x65
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x641D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x641D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5353) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5353, *c.R1.BC)
  }
  if (*c.R1.DE != 0x618D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x618D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3266) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3266, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1A53) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1A53, *c.R1.IX)
  }
  if (*c.R1.IY != 0x246A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x246A, *c.R1.IY)
  }
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

func TestFDCB88(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7D14
  *c.R1.BC = 0xA0EC
  *c.R1.DE = 0x1E47
  *c.R1.HL = 0x76E1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3871
  *c.R1.IY = 0xC60D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD4
  mem[0x0003] = 0x88
  mem[0xC5E1] = 0xD6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7D14) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7D14, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD4EC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD4EC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E47) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E47, *c.R1.DE)
  }
  if (*c.R1.HL != 0x76E1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x76E1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3871) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3871, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC60D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC60D, *c.R1.IY)
  }
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

func TestFDCB89(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x86C3
  *c.R1.BC = 0x50A6
  *c.R1.DE = 0x8592
  *c.R1.HL = 0xD6CA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x947B
  *c.R1.IY = 0x0A01
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC3
  mem[0x0003] = 0x89
  mem[0x09C4] = 0xB0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x86C3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x86C3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x50B0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x50B0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8592) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8592, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD6CA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD6CA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x947B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x947B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0A01) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0A01, *c.R1.IY)
  }
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

func TestFDCB8A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x599C
  *c.R1.BC = 0x961A
  *c.R1.DE = 0x55F9
  *c.R1.HL = 0x8470
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD2A5
  *c.R1.IY = 0xD4D2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF9
  mem[0x0003] = 0x8A
  mem[0xD4CB] = 0xD8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x599C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x599C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x961A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x961A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD8F9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD8F9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8470) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8470, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD2A5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD2A5, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD4D2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD4D2, *c.R1.IY)
  }
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

func TestFDCB8B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2715
  *c.R1.BC = 0xA209
  *c.R1.DE = 0xAB47
  *c.R1.HL = 0x3EAC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF352
  *c.R1.IY = 0xC71E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xED
  mem[0x0003] = 0x8B
  mem[0xC70B] = 0xDC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2715) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2715, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA209) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA209, *c.R1.BC)
  }
  if (*c.R1.DE != 0xABDC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xABDC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3EAC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3EAC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF352) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF352, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC71E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC71E, *c.R1.IY)
  }
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

func TestFDCB8C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2818
  *c.R1.BC = 0x4259
  *c.R1.DE = 0xA9B0
  *c.R1.HL = 0xE7A0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6471
  *c.R1.IY = 0xA202
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x97
  mem[0x0003] = 0x8C
  mem[0xA199] = 0x67
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2818) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2818, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4259) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4259, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA9B0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA9B0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x65A0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x65A0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6471) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6471, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA202) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA202, *c.R1.IY)
  }
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

func TestFDCB8D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x14E3
  *c.R1.BC = 0xC330
  *c.R1.DE = 0x9AA2
  *c.R1.HL = 0x8418
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0D4F
  *c.R1.IY = 0x5669
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC9
  mem[0x0003] = 0x8D
  mem[0x5632] = 0x9A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x14E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x14E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC330) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC330, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9AA2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9AA2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8498) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8498, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0D4F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0D4F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5669) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5669, *c.R1.IY)
  }
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

func TestFDCB8E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCB79
  *c.R1.BC = 0x0FFF
  *c.R1.DE = 0xB244
  *c.R1.HL = 0xC902
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6246
  *c.R1.IY = 0x4C81
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC2
  mem[0x0003] = 0x8E
  mem[0x4C43] = 0x7F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCB79) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCB79, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0FFF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0FFF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB244) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB244, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC902) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC902, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6246) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6246, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4C81) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4C81, *c.R1.IY)
  }
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

func TestFDCB8F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x66B4
  *c.R1.BC = 0x5FBB
  *c.R1.DE = 0x6C9B
  *c.R1.HL = 0xD0E3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAC5A
  *c.R1.IY = 0x6B51
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD4
  mem[0x0003] = 0x8F
  mem[0x6B25] = 0x59
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x59B4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x59B4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5FBB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5FBB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6C9B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6C9B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD0E3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD0E3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAC5A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAC5A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6B51) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6B51, *c.R1.IY)
  }
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

func TestFDCB90(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1305
  *c.R1.BC = 0x1CE1
  *c.R1.DE = 0xD627
  *c.R1.HL = 0x7402
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB470
  *c.R1.IY = 0xD7F5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFD
  mem[0x0003] = 0x90
  mem[0xD7F2] = 0x70
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1305) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1305, *c.R1.AF)
  }
  if (*c.R1.BC != 0x70E1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x70E1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD627) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD627, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7402) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7402, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB470) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB470, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD7F5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD7F5, *c.R1.IY)
  }
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

func TestFDCB91(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x10DF
  *c.R1.BC = 0xC48F
  *c.R1.DE = 0x0213
  *c.R1.HL = 0xFC7E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBFAB
  *c.R1.IY = 0x47D2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBF
  mem[0x0003] = 0x91
  mem[0x4791] = 0x0E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x10DF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x10DF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC40A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC40A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0213) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0213, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFC7E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFC7E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBFAB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBFAB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x47D2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x47D2, *c.R1.IY)
  }
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

func TestFDCB92(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6A11
  *c.R1.BC = 0xF89E
  *c.R1.DE = 0xF49D
  *c.R1.HL = 0xC115
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBC5D
  *c.R1.IY = 0x313A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0B
  mem[0x0003] = 0x92
  mem[0x3145] = 0xF6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6A11) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6A11, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF89E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF89E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF29D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF29D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC115) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC115, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBC5D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBC5D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x313A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x313A, *c.R1.IY)
  }
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

func TestFDCB93(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x61E5
  *c.R1.BC = 0xCC2C
  *c.R1.DE = 0x959A
  *c.R1.HL = 0xB52B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFA64
  *c.R1.IY = 0x2940
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x52
  mem[0x0003] = 0x93
  mem[0x2992] = 0x38
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x61E5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x61E5, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCC2C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCC2C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9538) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9538, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB52B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB52B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFA64) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFA64, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2940) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2940, *c.R1.IY)
  }
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

func TestFDCB94(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x31B4
  *c.R1.BC = 0x3E5A
  *c.R1.DE = 0xFB3D
  *c.R1.HL = 0xAB83
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA801
  *c.R1.IY = 0xFE1C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x95
  mem[0x0003] = 0x94
  mem[0xFDB1] = 0x48
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x31B4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x31B4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3E5A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3E5A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFB3D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFB3D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4883) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4883, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA801) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA801, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFE1C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFE1C, *c.R1.IY)
  }
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

func TestFDCB95(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x337E
  *c.R1.BC = 0x63A7
  *c.R1.DE = 0x2918
  *c.R1.HL = 0xED6B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB12C
  *c.R1.IY = 0xE776
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x90
  mem[0x0003] = 0x95
  mem[0xE706] = 0xEB
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x337E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x337E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x63A7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x63A7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2918) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2918, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEDEB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEDEB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB12C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB12C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE776) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE776, *c.R1.IY)
  }
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

func TestFDCB96(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5D99
  *c.R1.BC = 0xD9EC
  *c.R1.DE = 0xB6D0
  *c.R1.HL = 0x5ED5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5D9D
  *c.R1.IY = 0xE6CF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9E
  mem[0x0003] = 0x96
  mem[0xE66D] = 0xFC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5D99) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5D99, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD9EC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD9EC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB6D0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB6D0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5ED5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5ED5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5D9D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5D9D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE6CF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE6CF, *c.R1.IY)
  }
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

func TestFDCB97(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCCB6
  *c.R1.BC = 0x8406
  *c.R1.DE = 0x72C6
  *c.R1.HL = 0x1BA7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6DCA
  *c.R1.IY = 0x187F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x64
  mem[0x0003] = 0x97
  mem[0x18E3] = 0x9D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x99B6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x99B6, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8406) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8406, *c.R1.BC)
  }
  if (*c.R1.DE != 0x72C6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x72C6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1BA7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1BA7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6DCA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6DCA, *c.R1.IX)
  }
  if (*c.R1.IY != 0x187F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x187F, *c.R1.IY)
  }
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

func TestFDCB98(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0495
  *c.R1.BC = 0x312F
  *c.R1.DE = 0x8000
  *c.R1.HL = 0xB749
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE9CB
  *c.R1.IY = 0x43B8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDA
  mem[0x0003] = 0x98
  mem[0x4392] = 0x15
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0495) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0495, *c.R1.AF)
  }
  if (*c.R1.BC != 0x152F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x152F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8000, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB749) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB749, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE9CB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE9CB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x43B8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x43B8, *c.R1.IY)
  }
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

func TestFDCB99(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2824
  *c.R1.BC = 0xA485
  *c.R1.DE = 0xA30B
  *c.R1.HL = 0xB286
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x10B0
  *c.R1.IY = 0xD86C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x78
  mem[0x0003] = 0x99
  mem[0xD8E4] = 0xB5
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2824) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2824, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA4B5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA4B5, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA30B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA30B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB286) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB286, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x10B0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x10B0, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD86C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD86C, *c.R1.IY)
  }
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

func TestFDCB9A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB0CC
  *c.R1.BC = 0xC40C
  *c.R1.DE = 0xDC1A
  *c.R1.HL = 0x014A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2FF9
  *c.R1.IY = 0xD717
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9C
  mem[0x0003] = 0x9A
  mem[0xD6B3] = 0x9D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB0CC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB0CC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC40C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC40C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x951A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x951A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x014A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x014A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2FF9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2FF9, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD717) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD717, *c.R1.IY)
  }
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

func TestFDCB9B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD092
  *c.R1.BC = 0xA6C2
  *c.R1.DE = 0x7900
  *c.R1.HL = 0x5448
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFAB0
  *c.R1.IY = 0xCB1E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x83
  mem[0x0003] = 0x9B
  mem[0xCAA1] = 0x95
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD092) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD092, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA6C2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA6C2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7995) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7995, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5448) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5448, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFAB0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFAB0, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCB1E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCB1E, *c.R1.IY)
  }
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

func TestFDCB9C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB58D
  *c.R1.BC = 0x1ED1
  *c.R1.DE = 0xE93B
  *c.R1.HL = 0x9E0C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5605
  *c.R1.IY = 0x03B3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1E
  mem[0x0003] = 0x9C
  mem[0x03D1] = 0x78
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB58D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB58D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1ED1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1ED1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE93B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE93B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x700C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x700C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5605) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5605, *c.R1.IX)
  }
  if (*c.R1.IY != 0x03B3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x03B3, *c.R1.IY)
  }
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

func TestFDCB9D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC7E9
  *c.R1.BC = 0x18D3
  *c.R1.DE = 0x8EED
  *c.R1.HL = 0xBD7D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9A7F
  *c.R1.IY = 0xC087
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE6
  mem[0x0003] = 0x9D
  mem[0xC06D] = 0x53
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC7E9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC7E9, *c.R1.AF)
  }
  if (*c.R1.BC != 0x18D3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x18D3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8EED) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8EED, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBD53) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBD53, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9A7F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9A7F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC087) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC087, *c.R1.IY)
  }
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

func TestFDCB9E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x81C7
  *c.R1.BC = 0x71DF
  *c.R1.DE = 0x45D5
  *c.R1.HL = 0x0CA7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x648F
  *c.R1.IY = 0x41BD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEB
  mem[0x0003] = 0x9E
  mem[0x41A8] = 0x61
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x81C7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x81C7, *c.R1.AF)
  }
  if (*c.R1.BC != 0x71DF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x71DF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x45D5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x45D5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0CA7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0CA7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x648F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x648F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x41BD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x41BD, *c.R1.IY)
  }
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

func TestFDCB9F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEBF5
  *c.R1.BC = 0xDC9F
  *c.R1.DE = 0xD490
  *c.R1.HL = 0x15BE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0E12
  *c.R1.IY = 0x9D49
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x50
  mem[0x0003] = 0x9F
  mem[0x9D99] = 0x89
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x81F5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x81F5, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDC9F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDC9F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD490) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD490, *c.R1.DE)
  }
  if (*c.R1.HL != 0x15BE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x15BE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0E12) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0E12, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9D49) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9D49, *c.R1.IY)
  }
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

func TestFDCBA0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8CCB
  *c.R1.BC = 0x0057
  *c.R1.DE = 0xBC19
  *c.R1.HL = 0xE543
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8C5D
  *c.R1.IY = 0xD68D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x76
  mem[0x0003] = 0xA0
  mem[0xD703] = 0xD4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8CCB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8CCB, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC457) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC457, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBC19) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBC19, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE543) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE543, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8C5D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8C5D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD68D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD68D, *c.R1.IY)
  }
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

func TestFDCBA1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEEE6
  *c.R1.BC = 0x6DA4
  *c.R1.DE = 0x3A20
  *c.R1.HL = 0x8BBA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1DE7
  *c.R1.IY = 0x66C8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x31
  mem[0x0003] = 0xA1
  mem[0x66F9] = 0xEC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEEE6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEEE6, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6DEC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6DEC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3A20) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3A20, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8BBA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8BBA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1DE7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1DE7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x66C8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x66C8, *c.R1.IY)
  }
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

func TestFDCBA2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3F89
  *c.R1.BC = 0x5120
  *c.R1.DE = 0x0BD1
  *c.R1.HL = 0xE669
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2993
  *c.R1.IY = 0x04BF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0E
  mem[0x0003] = 0xA2
  mem[0x04CD] = 0x47
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3F89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3F89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5120) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5120, *c.R1.BC)
  }
  if (*c.R1.DE != 0x47D1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x47D1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE669) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE669, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2993) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2993, *c.R1.IX)
  }
  if (*c.R1.IY != 0x04BF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x04BF, *c.R1.IY)
  }
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

func TestFDCBA3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4439
  *c.R1.BC = 0x6B8B
  *c.R1.DE = 0x6178
  *c.R1.HL = 0x1246
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4CDB
  *c.R1.IY = 0xAD77
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x05
  mem[0x0003] = 0xA3
  mem[0xAD7C] = 0x59
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4439) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4439, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6B8B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6B8B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6149) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6149, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1246) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1246, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4CDB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4CDB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAD77) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAD77, *c.R1.IY)
  }
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

func TestFDCBA4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3385
  *c.R1.BC = 0x261E
  *c.R1.DE = 0xA487
  *c.R1.HL = 0xB3BD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4B8F
  *c.R1.IY = 0xC0CD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x66
  mem[0x0003] = 0xA4
  mem[0xC133] = 0xC5
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3385) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3385, *c.R1.AF)
  }
  if (*c.R1.BC != 0x261E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x261E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA487) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA487, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC5BD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC5BD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4B8F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4B8F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC0CD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC0CD, *c.R1.IY)
  }
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

func TestFDCBA5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6E70
  *c.R1.BC = 0xB7ED
  *c.R1.DE = 0x22CD
  *c.R1.HL = 0xAEDC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x46DE
  *c.R1.IY = 0xF1A1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA0
  mem[0x0003] = 0xA5
  mem[0xF141] = 0x44
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6E70) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6E70, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB7ED) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB7ED, *c.R1.BC)
  }
  if (*c.R1.DE != 0x22CD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x22CD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAE44) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAE44, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x46DE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x46DE, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF1A1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF1A1, *c.R1.IY)
  }
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

func TestFDCBA6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x814B
  *c.R1.BC = 0x6408
  *c.R1.DE = 0x3DCB
  *c.R1.HL = 0x971F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5716
  *c.R1.IY = 0x93F3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x76
  mem[0x0003] = 0xA6
  mem[0x9469] = 0xBC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x814B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x814B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6408) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6408, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3DCB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3DCB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x971F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x971F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5716) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5716, *c.R1.IX)
  }
  if (*c.R1.IY != 0x93F3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x93F3, *c.R1.IY)
  }
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

func TestFDCBA7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA4C2
  *c.R1.BC = 0x679E
  *c.R1.DE = 0xC313
  *c.R1.HL = 0x61DF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x67E6
  *c.R1.IY = 0x79C4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x66
  mem[0x0003] = 0xA7
  mem[0x7A2A] = 0x2E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2EC2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2EC2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x679E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x679E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC313) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC313, *c.R1.DE)
  }
  if (*c.R1.HL != 0x61DF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x61DF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x67E6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x67E6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x79C4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x79C4, *c.R1.IY)
  }
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

func TestFDCBA8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x537C
  *c.R1.BC = 0x1FED
  *c.R1.DE = 0x6CBB
  *c.R1.HL = 0xBD26
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC638
  *c.R1.IY = 0x0D46
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA9
  mem[0x0003] = 0xA8
  mem[0x0CEF] = 0xB7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x537C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x537C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x97ED) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x97ED, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6CBB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6CBB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBD26) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBD26, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC638) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC638, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0D46) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0D46, *c.R1.IY)
  }
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

func TestFDCBA9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBA5A
  *c.R1.BC = 0x3076
  *c.R1.DE = 0xCDD7
  *c.R1.HL = 0x298D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x59AB
  *c.R1.IY = 0x0F54
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2B
  mem[0x0003] = 0xA9
  mem[0x0F7F] = 0x8F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBA5A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBA5A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x308F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x308F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCDD7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCDD7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x298D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x298D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x59AB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x59AB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0F54) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0F54, *c.R1.IY)
  }
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

func TestFDCBAA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x406A
  *c.R1.BC = 0x2ED6
  *c.R1.DE = 0xFA8C
  *c.R1.HL = 0xC633
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x87CB
  *c.R1.IY = 0xB3D1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0B
  mem[0x0003] = 0xAA
  mem[0xB3DC] = 0x3A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x406A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x406A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2ED6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2ED6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1A8C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1A8C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC633) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC633, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x87CB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x87CB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB3D1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB3D1, *c.R1.IY)
  }
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

func TestFDCBAB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDA61
  *c.R1.BC = 0x0521
  *c.R1.DE = 0xA123
  *c.R1.HL = 0xC7FA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB71A
  *c.R1.IY = 0x8ECE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA9
  mem[0x0003] = 0xAB
  mem[0x8E77] = 0x1F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDA61) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDA61, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0521) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0521, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA11F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA11F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC7FA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC7FA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB71A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB71A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8ECE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8ECE, *c.R1.IY)
  }
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

func TestFDCBAC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x34A3
  *c.R1.BC = 0x81CE
  *c.R1.DE = 0x07D6
  *c.R1.HL = 0xF3A4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x430B
  *c.R1.IY = 0x0525
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x23
  mem[0x0003] = 0xAC
  mem[0x0548] = 0x9C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x34A3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x34A3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x81CE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x81CE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x07D6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x07D6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9CA4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9CA4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x430B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x430B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0525) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0525, *c.R1.IY)
  }
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

func TestFDCBAD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5010
  *c.R1.BC = 0x918E
  *c.R1.DE = 0xDDBC
  *c.R1.HL = 0x4F89
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x88C5
  *c.R1.IY = 0x948F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4E
  mem[0x0003] = 0xAD
  mem[0x94DD] = 0x37
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5010) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5010, *c.R1.AF)
  }
  if (*c.R1.BC != 0x918E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x918E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDDBC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDDBC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4F17) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4F17, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x88C5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x88C5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x948F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x948F, *c.R1.IY)
  }
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

func TestFDCBAE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEC0D
  *c.R1.BC = 0xB57E
  *c.R1.DE = 0x18C6
  *c.R1.HL = 0x7B01
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBAC6
  *c.R1.IY = 0x0C1D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0C
  mem[0x0003] = 0xAE
  mem[0x0C29] = 0xA9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEC0D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEC0D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB57E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB57E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x18C6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x18C6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7B01) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7B01, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBAC6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBAC6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0C1D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0C1D, *c.R1.IY)
  }
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

func TestFDCBAF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB322
  *c.R1.BC = 0x6731
  *c.R1.DE = 0xDAAD
  *c.R1.HL = 0x8D38
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDD8F
  *c.R1.IY = 0x26EB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0D
  mem[0x0003] = 0xAF
  mem[0x26F8] = 0x44
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4422) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4422, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6731) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6731, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDAAD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDAAD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8D38) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8D38, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDD8F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDD8F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x26EB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x26EB, *c.R1.IY)
  }
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

func TestFDCBB0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB984
  *c.R1.BC = 0x796C
  *c.R1.DE = 0x44B1
  *c.R1.HL = 0xFEF9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4069
  *c.R1.IY = 0xA0CB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5A
  mem[0x0003] = 0xB0
  mem[0xA125] = 0x76
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB984) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB984, *c.R1.AF)
  }
  if (*c.R1.BC != 0x366C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x366C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x44B1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x44B1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFEF9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFEF9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4069) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4069, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA0CB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA0CB, *c.R1.IY)
  }
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

func TestFDCBB1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x59C3
  *c.R1.BC = 0xAB13
  *c.R1.DE = 0x42EE
  *c.R1.HL = 0xB764
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8F7F
  *c.R1.IY = 0xF398
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x82
  mem[0x0003] = 0xB1
  mem[0xF31A] = 0x79
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x59C3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x59C3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAB39) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAB39, *c.R1.BC)
  }
  if (*c.R1.DE != 0x42EE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x42EE, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB764) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB764, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8F7F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8F7F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF398) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF398, *c.R1.IY)
  }
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

func TestFDCBB2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF310
  *c.R1.BC = 0xCEEC
  *c.R1.DE = 0xBBFB
  *c.R1.HL = 0x3569
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4A6F
  *c.R1.IY = 0x33F9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x23
  mem[0x0003] = 0xB2
  mem[0x341C] = 0x7B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF310) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF310, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCEEC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCEEC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3BFB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3BFB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3569) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3569, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4A6F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4A6F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x33F9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x33F9, *c.R1.IY)
  }
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

func TestFDCBB3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9C05
  *c.R1.BC = 0x0F92
  *c.R1.DE = 0xBD3B
  *c.R1.HL = 0x553D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC75E
  *c.R1.IY = 0x51D2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6C
  mem[0x0003] = 0xB3
  mem[0x523E] = 0x37
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9C05) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9C05, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F92) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F92, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBD37) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBD37, *c.R1.DE)
  }
  if (*c.R1.HL != 0x553D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x553D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC75E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC75E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x51D2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x51D2, *c.R1.IY)
  }
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

func TestFDCBB4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3E55
  *c.R1.BC = 0x1338
  *c.R1.DE = 0x638D
  *c.R1.HL = 0x353C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x44AD
  *c.R1.IY = 0x4D17
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC5
  mem[0x0003] = 0xB4
  mem[0x4CDC] = 0xE9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E55) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E55, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1338) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1338, *c.R1.BC)
  }
  if (*c.R1.DE != 0x638D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x638D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA93C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA93C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x44AD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x44AD, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4D17) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4D17, *c.R1.IY)
  }
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

func TestFDCBB5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2F3A
  *c.R1.BC = 0xB709
  *c.R1.DE = 0x4167
  *c.R1.HL = 0x57BE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB543
  *c.R1.IY = 0x8EDD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x50
  mem[0x0003] = 0xB5
  mem[0x8F2D] = 0x0F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2F3A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2F3A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB709) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB709, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4167) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4167, *c.R1.DE)
  }
  if (*c.R1.HL != 0x570F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x570F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB543) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB543, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8EDD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8EDD, *c.R1.IY)
  }
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

func TestFDCBB6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA887
  *c.R1.BC = 0x519B
  *c.R1.DE = 0xC91B
  *c.R1.HL = 0xCC91
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA416
  *c.R1.IY = 0x1E16
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3A
  mem[0x0003] = 0xB6
  mem[0x1E50] = 0x13
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA887) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA887, *c.R1.AF)
  }
  if (*c.R1.BC != 0x519B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x519B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC91B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC91B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCC91) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCC91, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA416) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA416, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1E16) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1E16, *c.R1.IY)
  }
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

func TestFDCBB7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1335
  *c.R1.BC = 0xA599
  *c.R1.DE = 0x9FBF
  *c.R1.HL = 0xC111
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8BC5
  *c.R1.IY = 0x00A9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC0
  mem[0x0003] = 0xB7
  mem[0x0069] = 0x38
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3835) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3835, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA599) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA599, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9FBF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9FBF, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC111) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC111, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8BC5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8BC5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x00A9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x00A9, *c.R1.IY)
  }
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

func TestFDCBB8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD146
  *c.R1.BC = 0x1138
  *c.R1.DE = 0x1A45
  *c.R1.HL = 0x8259
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6A03
  *c.R1.IY = 0xD087
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x31
  mem[0x0003] = 0xB8
  mem[0xD0B8] = 0x17
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD146) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD146, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1738) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1738, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1A45) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1A45, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8259) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8259, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6A03) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6A03, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD087) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD087, *c.R1.IY)
  }
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

func TestFDCBB9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x757B
  *c.R1.BC = 0x0B9E
  *c.R1.DE = 0x767B
  *c.R1.HL = 0x2AD1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1498
  *c.R1.IY = 0xB84E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3B
  mem[0x0003] = 0xB9
  mem[0xB889] = 0xB4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x757B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x757B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0B34) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0B34, *c.R1.BC)
  }
  if (*c.R1.DE != 0x767B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x767B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2AD1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2AD1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1498) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1498, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB84E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB84E, *c.R1.IY)
  }
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

func TestFDCBBA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x43EF
  *c.R1.BC = 0x1C58
  *c.R1.DE = 0xDDA3
  *c.R1.HL = 0x4519
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB67B
  *c.R1.IY = 0x383F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x38
  mem[0x0003] = 0xBA
  mem[0x3877] = 0xD6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x43EF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x43EF, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1C58) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1C58, *c.R1.BC)
  }
  if (*c.R1.DE != 0x56A3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x56A3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4519) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4519, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB67B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB67B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x383F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x383F, *c.R1.IY)
  }
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

func TestFDCBBB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDCCB
  *c.R1.BC = 0x7AB3
  *c.R1.DE = 0x7615
  *c.R1.HL = 0x4161
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2942
  *c.R1.IY = 0xE2FE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x07
  mem[0x0003] = 0xBB
  mem[0xE305] = 0x6E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDCCB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDCCB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7AB3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7AB3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x766E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x766E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4161) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4161, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2942) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2942, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE2FE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE2FE, *c.R1.IY)
  }
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

func TestFDCBBC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0E07
  *c.R1.BC = 0x34F5
  *c.R1.DE = 0x0995
  *c.R1.HL = 0xCC42
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9D42
  *c.R1.IY = 0xAF0C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF8
  mem[0x0003] = 0xBC
  mem[0xAF04] = 0xCF
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0E07) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0E07, *c.R1.AF)
  }
  if (*c.R1.BC != 0x34F5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x34F5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0995) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0995, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4F42) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4F42, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9D42) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9D42, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAF0C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAF0C, *c.R1.IY)
  }
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

func TestFDCBBD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x30EF
  *c.R1.BC = 0xE60C
  *c.R1.DE = 0x9BF0
  *c.R1.HL = 0xA1BF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBD1C
  *c.R1.IY = 0xDF0D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAA
  mem[0x0003] = 0xBD
  mem[0xDEB7] = 0x8D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x30EF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x30EF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE60C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE60C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9BF0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9BF0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA10D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA10D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBD1C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBD1C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDF0D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDF0D, *c.R1.IY)
  }
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

func TestFDCBBE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1133
  *c.R1.BC = 0xBEF6
  *c.R1.DE = 0x5059
  *c.R1.HL = 0x1089
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD558
  *c.R1.IY = 0x3D0F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC8
  mem[0x0003] = 0xBE
  mem[0x3CD7] = 0xA1
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1133) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1133, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBEF6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBEF6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5059) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5059, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1089) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1089, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD558) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD558, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3D0F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3D0F, *c.R1.IY)
  }
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

func TestFDCBBF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x83D6
  *c.R1.BC = 0xC893
  *c.R1.DE = 0x8DB8
  *c.R1.HL = 0x716B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0956
  *c.R1.IY = 0xBDE7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFD
  mem[0x0003] = 0xBF
  mem[0xBDE4] = 0xAC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2CD6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2CD6, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC893) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC893, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8DB8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8DB8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x716B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x716B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0956) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0956, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBDE7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBDE7, *c.R1.IY)
  }
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

func TestFDCBC0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3666
  *c.R1.BC = 0x676C
  *c.R1.DE = 0x35E5
  *c.R1.HL = 0xDB0A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEA93
  *c.R1.IY = 0x2B31
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0A
  mem[0x0003] = 0xC0
  mem[0x2B3B] = 0xEC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3666) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3666, *c.R1.AF)
  }
  if (*c.R1.BC != 0xED6C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xED6C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x35E5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x35E5, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDB0A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDB0A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEA93) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEA93, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2B31) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2B31, *c.R1.IY)
  }
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

func TestFDCBC1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3902
  *c.R1.BC = 0xD498
  *c.R1.DE = 0xAF62
  *c.R1.HL = 0x9821
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x48B8
  *c.R1.IY = 0xBD67
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x87
  mem[0x0003] = 0xC1
  mem[0xBCEE] = 0xEE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3902) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3902, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD4EF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD4EF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAF62) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAF62, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9821) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9821, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x48B8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x48B8, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBD67) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBD67, *c.R1.IY)
  }
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

func TestFDCBC2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAD26
  *c.R1.BC = 0x5A6D
  *c.R1.DE = 0x6762
  *c.R1.HL = 0x16C9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x495A
  *c.R1.IY = 0x5B2C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8D
  mem[0x0003] = 0xC2
  mem[0x5AB9] = 0xC2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAD26) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAD26, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5A6D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5A6D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC362) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC362, *c.R1.DE)
  }
  if (*c.R1.HL != 0x16C9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x16C9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x495A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x495A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5B2C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5B2C, *c.R1.IY)
  }
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

func TestFDCBC3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3E6C
  *c.R1.BC = 0x9A74
  *c.R1.DE = 0xA2EE
  *c.R1.HL = 0x9838
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEAFA
  *c.R1.IY = 0xE666
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5A
  mem[0x0003] = 0xC3
  mem[0xE6C0] = 0x4F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E6C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E6C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9A74) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9A74, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA24F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA24F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9838) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9838, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEAFA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEAFA, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE666) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE666, *c.R1.IY)
  }
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

func TestFDCBC4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBF68
  *c.R1.BC = 0xD00B
  *c.R1.DE = 0x5283
  *c.R1.HL = 0x51C2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x517C
  *c.R1.IY = 0x5D10
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x89
  mem[0x0003] = 0xC4
  mem[0x5C99] = 0x61
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBF68) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBF68, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD00B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD00B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5283) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5283, *c.R1.DE)
  }
  if (*c.R1.HL != 0x61C2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x61C2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x517C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x517C, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5D10) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5D10, *c.R1.IY)
  }
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

func TestFDCBC5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x127B
  *c.R1.BC = 0xDB6A
  *c.R1.DE = 0x00B9
  *c.R1.HL = 0x5138
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x98F6
  *c.R1.IY = 0x02BB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA9
  mem[0x0003] = 0xC5
  mem[0x0264] = 0xCD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x127B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x127B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDB6A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDB6A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x00B9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x00B9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x51CD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x51CD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x98F6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x98F6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x02BB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x02BB, *c.R1.IY)
  }
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

func TestFDCBC6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x35DA
  *c.R1.BC = 0x98C2
  *c.R1.DE = 0x3F57
  *c.R1.HL = 0x44A4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2771
  *c.R1.IY = 0x76C4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEE
  mem[0x0003] = 0xC6
  mem[0x76B2] = 0x82
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x35DA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x35DA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x98C2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x98C2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3F57) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3F57, *c.R1.DE)
  }
  if (*c.R1.HL != 0x44A4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x44A4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2771) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2771, *c.R1.IX)
  }
  if (*c.R1.IY != 0x76C4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x76C4, *c.R1.IY)
  }
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

func TestFDCBC7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x763F
  *c.R1.BC = 0xB86F
  *c.R1.DE = 0x12D3
  *c.R1.HL = 0x7E2D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD870
  *c.R1.IY = 0xF30B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9E
  mem[0x0003] = 0xC7
  mem[0xF2A9] = 0xD7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD73F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD73F, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB86F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB86F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x12D3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x12D3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7E2D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7E2D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD870) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD870, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF30B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF30B, *c.R1.IY)
  }
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

func TestFDCBC8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1F81
  *c.R1.BC = 0xC7C0
  *c.R1.DE = 0x85DA
  *c.R1.HL = 0x3CDD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD854
  *c.R1.IY = 0xC412
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x10
  mem[0x0003] = 0xC8
  mem[0xC422] = 0xE9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1F81) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1F81, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEBC0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEBC0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x85DA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x85DA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3CDD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3CDD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD854) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD854, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC412) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC412, *c.R1.IY)
  }
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

func TestFDCBC9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xED19
  *c.R1.BC = 0x3F88
  *c.R1.DE = 0x1370
  *c.R1.HL = 0xE084
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4FDD
  *c.R1.IY = 0x8B42
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x61
  mem[0x0003] = 0xC9
  mem[0x8BA3] = 0xB7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xED19) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xED19, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3FB7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3FB7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1370) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1370, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE084) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE084, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4FDD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4FDD, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8B42) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8B42, *c.R1.IY)
  }
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

func TestFDCBCA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC7E5
  *c.R1.BC = 0x233B
  *c.R1.DE = 0x2312
  *c.R1.HL = 0xF7F9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE417
  *c.R1.IY = 0x5190
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1A
  mem[0x0003] = 0xCA
  mem[0x51AA] = 0x90
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC7E5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC7E5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x233B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x233B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9212) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9212, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF7F9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF7F9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE417) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE417, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5190) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5190, *c.R1.IY)
  }
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

func TestFDCBCB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBDBA
  *c.R1.BC = 0xA964
  *c.R1.DE = 0xEA38
  *c.R1.HL = 0x9422
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFCA3
  *c.R1.IY = 0x9A72
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5E
  mem[0x0003] = 0xCB
  mem[0x9AD0] = 0x70
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBDBA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBDBA, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA964) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA964, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEA72) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEA72, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9422) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9422, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFCA3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFCA3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9A72) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9A72, *c.R1.IY)
  }
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

func TestFDCBCC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0F4F
  *c.R1.BC = 0x0261
  *c.R1.DE = 0x21B0
  *c.R1.HL = 0x2097
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x575D
  *c.R1.IY = 0x14F9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2D
  mem[0x0003] = 0xCC
  mem[0x1526] = 0x4E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0F4F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0F4F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0261) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0261, *c.R1.BC)
  }
  if (*c.R1.DE != 0x21B0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x21B0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4E97) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4E97, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x575D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x575D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x14F9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x14F9, *c.R1.IY)
  }
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

func TestFDCBCD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1B79
  *c.R1.BC = 0x8F9F
  *c.R1.DE = 0x31BF
  *c.R1.HL = 0x9CA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7ECB
  *c.R1.IY = 0xBBE9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA1
  mem[0x0003] = 0xCD
  mem[0xBB8A] = 0x66
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1B79) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1B79, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8F9F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8F9F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x31BF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x31BF, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9C66) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9C66, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7ECB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7ECB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBBE9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBBE9, *c.R1.IY)
  }
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

func TestFDCBCE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8E13
  *c.R1.BC = 0x968E
  *c.R1.DE = 0x1784
  *c.R1.HL = 0x0A0A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1E87
  *c.R1.IY = 0xB8A2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x36
  mem[0x0003] = 0xCE
  mem[0xB8D8] = 0x45
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8E13) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8E13, *c.R1.AF)
  }
  if (*c.R1.BC != 0x968E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x968E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1784) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1784, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0A0A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0A0A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1E87) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1E87, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB8A2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB8A2, *c.R1.IY)
  }
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

func TestFDCBCF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8D0A
  *c.R1.BC = 0xA073
  *c.R1.DE = 0xC4BA
  *c.R1.HL = 0x5B69
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3B47
  *c.R1.IY = 0xC29C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x70
  mem[0x0003] = 0xCF
  mem[0xC30C] = 0x7A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7A0A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7A0A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA073) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA073, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC4BA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC4BA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5B69) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5B69, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3B47) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3B47, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC29C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC29C, *c.R1.IY)
  }
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

func TestFDCBD0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE2BB
  *c.R1.BC = 0x8635
  *c.R1.DE = 0x650C
  *c.R1.HL = 0x689A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1294
  *c.R1.IY = 0x3BEB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBC
  mem[0x0003] = 0xD0
  mem[0x3BA7] = 0x20
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE2BB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE2BB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2435) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2435, *c.R1.BC)
  }
  if (*c.R1.DE != 0x650C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x650C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x689A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x689A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1294) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1294, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3BEB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3BEB, *c.R1.IY)
  }
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

func TestFDCBD1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5DF8
  *c.R1.BC = 0xF701
  *c.R1.DE = 0x9494
  *c.R1.HL = 0x4967
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAD00
  *c.R1.IY = 0x8C65
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x11
  mem[0x0003] = 0xD1
  mem[0x8C76] = 0xB9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5DF8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5DF8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF7BD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF7BD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9494) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9494, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4967) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4967, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAD00) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAD00, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8C65) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8C65, *c.R1.IY)
  }
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

func TestFDCBD2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9876
  *c.R1.BC = 0x4BD9
  *c.R1.DE = 0x3148
  *c.R1.HL = 0x665A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7EAC
  *c.R1.IY = 0xC051
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFB
  mem[0x0003] = 0xD2
  mem[0xC04C] = 0x51
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9876) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9876, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4BD9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4BD9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5548) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5548, *c.R1.DE)
  }
  if (*c.R1.HL != 0x665A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x665A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7EAC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7EAC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC051) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC051, *c.R1.IY)
  }
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

func TestFDCBD3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8F90
  *c.R1.BC = 0xBACD
  *c.R1.DE = 0xE87A
  *c.R1.HL = 0x538F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFE5A
  *c.R1.IY = 0x0A87
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3E
  mem[0x0003] = 0xD3
  mem[0x0AC5] = 0xE0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8F90) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8F90, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBACD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBACD, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE8E4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE8E4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x538F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x538F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFE5A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFE5A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0A87) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0A87, *c.R1.IY)
  }
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

func TestFDCBD4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x15E2
  *c.R1.BC = 0x1820
  *c.R1.DE = 0x5588
  *c.R1.HL = 0xE67F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7193
  *c.R1.IY = 0x9478
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x36
  mem[0x0003] = 0xD4
  mem[0x94AE] = 0x7D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x15E2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x15E2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1820) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1820, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5588) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5588, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7D7F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7D7F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7193) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7193, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9478) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9478, *c.R1.IY)
  }
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

func TestFDCBD5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1409
  *c.R1.BC = 0x6535
  *c.R1.DE = 0xC371
  *c.R1.HL = 0xABE2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2E10
  *c.R1.IY = 0x8608
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x48
  mem[0x0003] = 0xD5
  mem[0x8650] = 0x98
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1409) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1409, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6535) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6535, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC371) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC371, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAB9C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAB9C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2E10) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2E10, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8608) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8608, *c.R1.IY)
  }
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

func TestFDCBD6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7801
  *c.R1.BC = 0x78B6
  *c.R1.DE = 0xD191
  *c.R1.HL = 0x054A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2065
  *c.R1.IY = 0x6AA3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC9
  mem[0x0003] = 0xD6
  mem[0x6A6C] = 0x7C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7801) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7801, *c.R1.AF)
  }
  if (*c.R1.BC != 0x78B6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x78B6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD191) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD191, *c.R1.DE)
  }
  if (*c.R1.HL != 0x054A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x054A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2065) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2065, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6AA3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6AA3, *c.R1.IY)
  }
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

func TestFDCBD7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1B6A
  *c.R1.BC = 0x266E
  *c.R1.DE = 0x387F
  *c.R1.HL = 0x7FCB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1941
  *c.R1.IY = 0x36AB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBE
  mem[0x0003] = 0xD7
  mem[0x3669] = 0x95
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x956A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x956A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x266E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x266E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x387F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x387F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7FCB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7FCB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1941) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1941, *c.R1.IX)
  }
  if (*c.R1.IY != 0x36AB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x36AB, *c.R1.IY)
  }
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

func TestFDCBD8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7B1B
  *c.R1.BC = 0xA191
  *c.R1.DE = 0xEFEE
  *c.R1.HL = 0x55B9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF789
  *c.R1.IY = 0x43F8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBC
  mem[0x0003] = 0xD8
  mem[0x43B4] = 0xD8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7B1B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7B1B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD891) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD891, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEFEE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEFEE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x55B9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x55B9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF789) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF789, *c.R1.IX)
  }
  if (*c.R1.IY != 0x43F8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x43F8, *c.R1.IY)
  }
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

func TestFDCBD9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0FAF
  *c.R1.BC = 0x4EDA
  *c.R1.DE = 0xC556
  *c.R1.HL = 0x6ED3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3FC3
  *c.R1.IY = 0x0A66
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x16
  mem[0x0003] = 0xD9
  mem[0x0A7C] = 0xF4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0FAF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0FAF, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4EFC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4EFC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC556) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC556, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6ED3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6ED3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3FC3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3FC3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0A66) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0A66, *c.R1.IY)
  }
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

func TestFDCBDA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9EA1
  *c.R1.BC = 0x8186
  *c.R1.DE = 0xC045
  *c.R1.HL = 0xD6E0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x34D3
  *c.R1.IY = 0xD0F0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE8
  mem[0x0003] = 0xDA
  mem[0xD0D8] = 0x6B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9EA1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9EA1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8186) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8186, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6B45) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6B45, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD6E0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD6E0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x34D3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x34D3, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD0F0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD0F0, *c.R1.IY)
  }
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

func TestFDCBDB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5EE0
  *c.R1.BC = 0xBDEA
  *c.R1.DE = 0xD00E
  *c.R1.HL = 0x513F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x690A
  *c.R1.IY = 0x8C29
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7A
  mem[0x0003] = 0xDB
  mem[0x8CA3] = 0x15
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5EE0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5EE0, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBDEA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBDEA, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD01D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD01D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x513F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x513F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x690A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x690A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8C29) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8C29, *c.R1.IY)
  }
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

func TestFDCBDC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5CFA
  *c.R1.BC = 0x2E2B
  *c.R1.DE = 0x1D17
  *c.R1.HL = 0xDBF6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA4F2
  *c.R1.IY = 0x593A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x64
  mem[0x0003] = 0xDC
  mem[0x599E] = 0x15
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5CFA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5CFA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2E2B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2E2B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1D17) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1D17, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1DF6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1DF6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA4F2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA4F2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x593A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x593A, *c.R1.IY)
  }
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

func TestFDCBDD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8773
  *c.R1.BC = 0x70A6
  *c.R1.DE = 0x83CE
  *c.R1.HL = 0x52B8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x35DA
  *c.R1.IY = 0x1D94
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x75
  mem[0x0003] = 0xDD
  mem[0x1E09] = 0x28
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8773) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8773, *c.R1.AF)
  }
  if (*c.R1.BC != 0x70A6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x70A6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x83CE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x83CE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5228) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5228, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x35DA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x35DA, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1D94) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1D94, *c.R1.IY)
  }
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

func TestFDCBDE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8310
  *c.R1.BC = 0xFA01
  *c.R1.DE = 0x6C69
  *c.R1.HL = 0x252A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5291
  *c.R1.IY = 0xC9E0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x17
  mem[0x0003] = 0xDE
  mem[0xC9F7] = 0x41
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8310) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8310, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFA01) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFA01, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6C69) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6C69, *c.R1.DE)
  }
  if (*c.R1.HL != 0x252A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x252A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5291) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5291, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC9E0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC9E0, *c.R1.IY)
  }
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

func TestFDCBDF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x780D
  *c.R1.BC = 0xA722
  *c.R1.DE = 0xE78E
  *c.R1.HL = 0x50BA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9D67
  *c.R1.IY = 0xEAC3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x93
  mem[0x0003] = 0xDF
  mem[0xEA56] = 0xEF
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEF0D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEF0D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA722) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA722, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE78E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE78E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x50BA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x50BA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9D67) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9D67, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEAC3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEAC3, *c.R1.IY)
  }
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

func TestFDCBE0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x10EF
  *c.R1.BC = 0x4101
  *c.R1.DE = 0x2CA5
  *c.R1.HL = 0xF752
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4747
  *c.R1.IY = 0x1507
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x56
  mem[0x0003] = 0xE0
  mem[0x155D] = 0xB9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x10EF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x10EF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB901) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB901, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2CA5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2CA5, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF752) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF752, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4747) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4747, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1507) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1507, *c.R1.IY)
  }
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

func TestFDCBE1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE4CB
  *c.R1.BC = 0x6F72
  *c.R1.DE = 0x1C11
  *c.R1.HL = 0x1426
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x189B
  *c.R1.IY = 0x0E0D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD1
  mem[0x0003] = 0xE1
  mem[0x0DDE] = 0x16
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE4CB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE4CB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6F16) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6F16, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1C11) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1C11, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1426) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1426, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x189B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x189B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0E0D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0E0D, *c.R1.IY)
  }
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

func TestFDCBE2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x11A9
  *c.R1.BC = 0xBAE8
  *c.R1.DE = 0x938B
  *c.R1.HL = 0xBAC4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD8ED
  *c.R1.IY = 0xE49C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x50
  mem[0x0003] = 0xE2
  mem[0xE4EC] = 0xC2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x11A9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x11A9, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBAE8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBAE8, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD28B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD28B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBAC4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBAC4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD8ED) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD8ED, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE49C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE49C, *c.R1.IY)
  }
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

func TestFDCBE3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8832
  *c.R1.BC = 0x952B
  *c.R1.DE = 0x02B2
  *c.R1.HL = 0x26EF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFB55
  *c.R1.IY = 0xADA8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xCA
  mem[0x0003] = 0xE3
  mem[0xAD72] = 0xBA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8832) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8832, *c.R1.AF)
  }
  if (*c.R1.BC != 0x952B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x952B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x02BA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x02BA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x26EF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x26EF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFB55) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFB55, *c.R1.IX)
  }
  if (*c.R1.IY != 0xADA8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xADA8, *c.R1.IY)
  }
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

func TestFDCBE4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3989
  *c.R1.BC = 0x4142
  *c.R1.DE = 0x89E2
  *c.R1.HL = 0x785B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0BF7
  *c.R1.IY = 0x5474
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x62
  mem[0x0003] = 0xE4
  mem[0x54D6] = 0x7B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3989) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3989, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4142) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4142, *c.R1.BC)
  }
  if (*c.R1.DE != 0x89E2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x89E2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7B5B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7B5B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0BF7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0BF7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5474) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5474, *c.R1.IY)
  }
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

func TestFDCBE5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE5C5
  *c.R1.BC = 0xB86D
  *c.R1.DE = 0x41BB
  *c.R1.HL = 0x315E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1A78
  *c.R1.IY = 0xA52D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDA
  mem[0x0003] = 0xE5
  mem[0xA507] = 0x4C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE5C5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE5C5, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB86D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB86D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x41BB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x41BB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x315C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x315C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1A78) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1A78, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA52D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA52D, *c.R1.IY)
  }
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

func TestFDCBE6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFD89
  *c.R1.BC = 0xD888
  *c.R1.DE = 0x1E2F
  *c.R1.HL = 0xDDF5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x42F5
  *c.R1.IY = 0x8B06
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x76
  mem[0x0003] = 0xE6
  mem[0x8B7C] = 0x45
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFD89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFD89, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD888) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD888, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E2F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E2F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDDF5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDDF5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x42F5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x42F5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8B06) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8B06, *c.R1.IY)
  }
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

func TestFDCBE7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2025
  *c.R1.BC = 0xD3E9
  *c.R1.DE = 0xD4B6
  *c.R1.HL = 0xAA30
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x88BD
  *c.R1.IY = 0xB597
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x11
  mem[0x0003] = 0xE7
  mem[0xB5A8] = 0xA6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB625) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB625, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD3E9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD3E9, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD4B6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD4B6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAA30) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAA30, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x88BD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x88BD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB597) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB597, *c.R1.IY)
  }
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

func TestFDCBE8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x514D
  *c.R1.BC = 0xC2AB
  *c.R1.DE = 0x37B5
  *c.R1.HL = 0x57DE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA4EC
  *c.R1.IY = 0x0A77
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xED
  mem[0x0003] = 0xE8
  mem[0x0A64] = 0xD0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x514D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x514D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF0AB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF0AB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x37B5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x37B5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x57DE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x57DE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA4EC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA4EC, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0A77) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0A77, *c.R1.IY)
  }
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

func TestFDCBE9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x974E
  *c.R1.BC = 0xD28E
  *c.R1.DE = 0xD5CB
  *c.R1.HL = 0x6BD4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x158A
  *c.R1.IY = 0xA84E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x35
  mem[0x0003] = 0xE9
  mem[0xA883] = 0x2F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x974E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x974E, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD22F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD22F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD5CB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD5CB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6BD4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6BD4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x158A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x158A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA84E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA84E, *c.R1.IY)
  }
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

func TestFDCBEA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3EF4
  *c.R1.BC = 0x3FC6
  *c.R1.DE = 0x4A44
  *c.R1.HL = 0xE9A4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC877
  *c.R1.IY = 0x7593
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x93
  mem[0x0003] = 0xEA
  mem[0x7526] = 0x1B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3EF4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3EF4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3FC6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3FC6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3B44) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3B44, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE9A4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE9A4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC877) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC877, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7593) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7593, *c.R1.IY)
  }
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

func TestFDCBEB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x798F
  *c.R1.BC = 0x5E9B
  *c.R1.DE = 0x940E
  *c.R1.HL = 0x2E52
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD6AD
  *c.R1.IY = 0x2411
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD0
  mem[0x0003] = 0xEB
  mem[0x23E1] = 0x47
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x798F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x798F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5E9B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5E9B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9467) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9467, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2E52) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2E52, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD6AD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD6AD, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2411) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2411, *c.R1.IY)
  }
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

func TestFDCBEC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x38A4
  *c.R1.BC = 0x07C0
  *c.R1.DE = 0x6CEE
  *c.R1.HL = 0xE715
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF160
  *c.R1.IY = 0xD2EB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF3
  mem[0x0003] = 0xEC
  mem[0xD2DE] = 0x49
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x38A4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x38A4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x07C0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x07C0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6CEE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6CEE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6915) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6915, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF160) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF160, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD2EB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD2EB, *c.R1.IY)
  }
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

func TestFDCBED(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE0BC
  *c.R1.BC = 0x70C1
  *c.R1.DE = 0xDE35
  *c.R1.HL = 0x81C5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD57F
  *c.R1.IY = 0x0EAB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x58
  mem[0x0003] = 0xED
  mem[0x0F03] = 0x10
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE0BC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE0BC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x70C1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x70C1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDE35) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDE35, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8130) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8130, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD57F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD57F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0EAB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0EAB, *c.R1.IY)
  }
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

func TestFDCBEE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5FCB
  *c.R1.BC = 0x9007
  *c.R1.DE = 0x1736
  *c.R1.HL = 0xACA8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4BAB
  *c.R1.IY = 0x42BC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x02
  mem[0x0003] = 0xEE
  mem[0x42BE] = 0xD0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5FCB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5FCB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9007) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9007, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1736) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1736, *c.R1.DE)
  }
  if (*c.R1.HL != 0xACA8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xACA8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4BAB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4BAB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x42BC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x42BC, *c.R1.IY)
  }
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

func TestFDCBEF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4EE3
  *c.R1.BC = 0xD344
  *c.R1.DE = 0xCB5B
  *c.R1.HL = 0xAEB5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDE5F
  *c.R1.IY = 0x2272
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x58
  mem[0x0003] = 0xEF
  mem[0x22CA] = 0x09
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x29E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x29E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD344) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD344, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCB5B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCB5B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAEB5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAEB5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDE5F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDE5F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2272) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2272, *c.R1.IY)
  }
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

func TestFDCBF0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1080
  *c.R1.BC = 0xB270
  *c.R1.DE = 0x1B5B
  *c.R1.HL = 0xA9B7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE89D
  *c.R1.IY = 0xEE9E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x45
  mem[0x0003] = 0xF0
  mem[0xEEE3] = 0x2C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1080) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1080, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6C70) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6C70, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1B5B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1B5B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA9B7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA9B7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE89D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE89D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEE9E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEE9E, *c.R1.IY)
  }
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

func TestFDCBF1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1702
  *c.R1.BC = 0xC43B
  *c.R1.DE = 0xD138
  *c.R1.HL = 0x316F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8067
  *c.R1.IY = 0x4783
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2F
  mem[0x0003] = 0xF1
  mem[0x47B2] = 0xDC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1702) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1702, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC4DC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC4DC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD138) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD138, *c.R1.DE)
  }
  if (*c.R1.HL != 0x316F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x316F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8067) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8067, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4783) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4783, *c.R1.IY)
  }
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

func TestFDCBF2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x732A
  *c.R1.BC = 0x4CD1
  *c.R1.DE = 0x77FE
  *c.R1.HL = 0x4814
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x42F1
  *c.R1.IY = 0xEA97
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2C
  mem[0x0003] = 0xF2
  mem[0xEAC3] = 0x5E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x732A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x732A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4CD1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4CD1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5EFE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5EFE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4814) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4814, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x42F1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x42F1, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEA97) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEA97, *c.R1.IY)
  }
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

func TestFDCBF3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6B97
  *c.R1.BC = 0x59D3
  *c.R1.DE = 0xF546
  *c.R1.HL = 0x7530
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6670
  *c.R1.IY = 0x7D90
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x38
  mem[0x0003] = 0xF3
  mem[0x7DC8] = 0x0C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6B97) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6B97, *c.R1.AF)
  }
  if (*c.R1.BC != 0x59D3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x59D3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF54C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF54C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7530) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7530, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6670) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6670, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7D90) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7D90, *c.R1.IY)
  }
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

func TestFDCBF4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7AF0
  *c.R1.BC = 0xA81F
  *c.R1.DE = 0x5D3A
  *c.R1.HL = 0x799B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE12B
  *c.R1.IY = 0x309C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD0
  mem[0x0003] = 0xF4
  mem[0x306C] = 0x0E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7AF0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7AF0, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA81F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA81F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5D3A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5D3A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4E9B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4E9B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE12B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE12B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x309C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x309C, *c.R1.IY)
  }
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

func TestFDCBF5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1370
  *c.R1.BC = 0xF6B2
  *c.R1.DE = 0xAAA2
  *c.R1.HL = 0x7F0A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC9F6
  *c.R1.IY = 0x6B1F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x55
  mem[0x0003] = 0xF5
  mem[0x6B74] = 0xF8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1370) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1370, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF6B2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF6B2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAAA2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAAA2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7FF8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7FF8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC9F6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC9F6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6B1F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6B1F, *c.R1.IY)
  }
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

func TestFDCBF6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7C43
  *c.R1.BC = 0xFCD1
  *c.R1.DE = 0x34BD
  *c.R1.HL = 0xF4AB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEF33
  *c.R1.IY = 0xC61A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x56
  mem[0x0003] = 0xF6
  mem[0xC670] = 0x5D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7C43) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7C43, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFCD1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFCD1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x34BD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x34BD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF4AB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF4AB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEF33) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEF33, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC61A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC61A, *c.R1.IY)
  }
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

func TestFDCBF7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE6DA
  *c.R1.BC = 0x231A
  *c.R1.DE = 0x7BB1
  *c.R1.HL = 0x800D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE37E
  *c.R1.IY = 0x5789
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9E
  mem[0x0003] = 0xF7
  mem[0x5727] = 0x66
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x66DA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x66DA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x231A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x231A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7BB1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7BB1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x800D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x800D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE37E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE37E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5789) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5789, *c.R1.IY)
  }
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

func TestFDCBF8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFA29
  *c.R1.BC = 0xEE74
  *c.R1.DE = 0xD7C4
  *c.R1.HL = 0xAFAF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x512C
  *c.R1.IY = 0xDE7A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x54
  mem[0x0003] = 0xF8
  mem[0xDECE] = 0x7A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFA29) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFA29, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFA74) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFA74, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD7C4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD7C4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAFAF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAFAF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x512C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x512C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDE7A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDE7A, *c.R1.IY)
  }
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

func TestFDCBF9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4662
  *c.R1.BC = 0xA71B
  *c.R1.DE = 0x5065
  *c.R1.HL = 0xED06
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x279E
  *c.R1.IY = 0x99E3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x30
  mem[0x0003] = 0xF9
  mem[0x9A13] = 0xC6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4662) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4662, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA7C6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA7C6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5065) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5065, *c.R1.DE)
  }
  if (*c.R1.HL != 0xED06) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xED06, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x279E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x279E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x99E3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x99E3, *c.R1.IY)
  }
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

func TestFDCBFA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9426
  *c.R1.BC = 0x53EC
  *c.R1.DE = 0x5016
  *c.R1.HL = 0x6C99
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8B99
  *c.R1.IY = 0xBD79
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x09
  mem[0x0003] = 0xFA
  mem[0xBD82] = 0xF4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9426) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9426, *c.R1.AF)
  }
  if (*c.R1.BC != 0x53EC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x53EC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF416) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF416, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6C99) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6C99, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8B99) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8B99, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBD79) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBD79, *c.R1.IY)
  }
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

func TestFDCBFB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5343
  *c.R1.BC = 0xB212
  *c.R1.DE = 0x09CA
  *c.R1.HL = 0xE3C6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCD2B
  *c.R1.IY = 0xF875
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBA
  mem[0x0003] = 0xFB
  mem[0xF82F] = 0xED
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5343) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5343, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB212) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB212, *c.R1.BC)
  }
  if (*c.R1.DE != 0x09ED) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x09ED, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE3C6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE3C6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCD2B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCD2B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF875) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF875, *c.R1.IY)
  }
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

func TestFDCBFC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0965
  *c.R1.BC = 0x4392
  *c.R1.DE = 0xCA25
  *c.R1.HL = 0x2BAA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF023
  *c.R1.IY = 0x6623
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x56
  mem[0x0003] = 0xFC
  mem[0x6679] = 0x65
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0965) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0965, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4392) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4392, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCA25) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCA25, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE5AA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE5AA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF023) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF023, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6623) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6623, *c.R1.IY)
  }
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

func TestFDCBFD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1751
  *c.R1.BC = 0x233C
  *c.R1.DE = 0x6214
  *c.R1.HL = 0xD119
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC415
  *c.R1.IY = 0x5D2B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x25
  mem[0x0003] = 0xFD
  mem[0x5D50] = 0x27
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1751) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1751, *c.R1.AF)
  }
  if (*c.R1.BC != 0x233C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x233C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6214) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6214, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD1A7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD1A7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC415) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC415, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5D2B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5D2B, *c.R1.IY)
  }
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

func TestFDCBFE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB4CF
  *c.R1.BC = 0x5639
  *c.R1.DE = 0x677B
  *c.R1.HL = 0x0CA2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDDC5
  *c.R1.IY = 0x4E4F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x88
  mem[0x0003] = 0xFE
  mem[0x4DD7] = 0x4A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB4CF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB4CF, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5639) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5639, *c.R1.BC)
  }
  if (*c.R1.DE != 0x677B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x677B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0CA2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0CA2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDDC5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDDC5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4E4F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4E4F, *c.R1.IY)
  }
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

func TestFDCBFF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF151
  *c.R1.BC = 0x13DA
  *c.R1.DE = 0x7C56
  *c.R1.HL = 0xF025
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2B36
  *c.R1.IY = 0x2AED
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE4
  mem[0x0003] = 0xFF
  mem[0x2AD1] = 0x97
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9751) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9751, *c.R1.AF)
  }
  if (*c.R1.BC != 0x13DA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x13DA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7C56) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7C56, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF025) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF025, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2B36) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2B36, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2AED) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2AED, *c.R1.IY)
  }
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