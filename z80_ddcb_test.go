package z80

import "testing"


func TestDDCB00(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3C65
  *c.R1.BC = 0xF0E4
  *c.R1.DE = 0x09D1
  *c.R1.HL = 0x646B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1DA1
  *c.R1.IY = 0xF08F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0D
  mem[0x0003] = 0x00
  mem[0x1DAE] = 0xA1
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C01) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C01, *c.R1.AF)
  }
  if (*c.R1.BC != 0x43E4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x43E4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x09D1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x09D1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x646B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x646B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1DA1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1DA1, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF08F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF08F, *c.R1.IY)
  }
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

func TestDDCB01(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF68F
  *c.R1.BC = 0xE33B
  *c.R1.DE = 0x2D4A
  *c.R1.HL = 0x7725
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x28FD
  *c.R1.IY = 0xF31B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB7
  mem[0x0003] = 0x01
  mem[0x28B4] = 0xE3
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF681) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF681, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE3C7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE3C7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2D4A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2D4A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7725) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7725, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x28FD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x28FD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF31B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF31B, *c.R1.IY)
  }
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

func TestDDCB02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE20C
  *c.R1.BC = 0x836E
  *c.R1.DE = 0x513A
  *c.R1.HL = 0xF840
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC796
  *c.R1.IY = 0xAE9B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x91
  mem[0x0003] = 0x02
  mem[0xC727] = 0x8D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE20D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE20D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x836E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x836E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1B3A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1B3A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF840) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF840, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC796) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC796, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAE9B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAE9B, *c.R1.IY)
  }
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

func TestDDCB03(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6224
  *c.R1.BC = 0x3571
  *c.R1.DE = 0xC519
  *c.R1.HL = 0x48DC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x041E
  *c.R1.IY = 0xC07B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x48
  mem[0x0003] = 0x03
  mem[0x0466] = 0x78
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x62A4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x62A4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3571) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3571, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC5F0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC5F0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x48DC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x48DC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x041E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x041E, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC07B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC07B, *c.R1.IY)
  }
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

func TestDDCB04(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB310
  *c.R1.BC = 0xBFC4
  *c.R1.DE = 0x64AF
  *c.R1.HL = 0xD622
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5949
  *c.R1.IY = 0xA989
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x48
  mem[0x0003] = 0x04
  mem[0x5991] = 0x68
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB380) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB380, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBFC4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBFC4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x64AF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x64AF, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD022) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD022, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5949) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5949, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA989) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA989, *c.R1.IY)
  }
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

func TestDDCB05(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4954
  *c.R1.BC = 0xBB04
  *c.R1.DE = 0x56EC
  *c.R1.HL = 0x9D58
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0077
  *c.R1.IY = 0x1349
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFF
  mem[0x0003] = 0x05
  mem[0x0076] = 0x95
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x492D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x492D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBB04) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBB04, *c.R1.BC)
  }
  if (*c.R1.DE != 0x56EC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x56EC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9D2B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9D2B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0077) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0077, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1349) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1349, *c.R1.IY)
  }
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

func TestDDCB06(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0CF4
  *c.R1.BC = 0xF636
  *c.R1.DE = 0x90A6
  *c.R1.HL = 0x6117
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5421
  *c.R1.IY = 0x90EE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x07
  mem[0x0003] = 0x06
  mem[0x5428] = 0x97
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0C29) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0C29, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF636) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF636, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90A6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90A6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6117) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6117, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5421) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5421, *c.R1.IX)
  }
  if (*c.R1.IY != 0x90EE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x90EE, *c.R1.IY)
  }
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

func TestDDCB07(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6F4D
  *c.R1.BC = 0x9CA3
  *c.R1.DE = 0xBDF6
  *c.R1.HL = 0xED50
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9803
  *c.R1.IY = 0x55F9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x42
  mem[0x0003] = 0x07
  mem[0x9845] = 0xAE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5D09) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5D09, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9CA3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9CA3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBDF6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBDF6, *c.R1.DE)
  }
  if (*c.R1.HL != 0xED50) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xED50, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9803) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9803, *c.R1.IX)
  }
  if (*c.R1.IY != 0x55F9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x55F9, *c.R1.IY)
  }
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

func TestDDCB08(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x02F4
  *c.R1.BC = 0x1C66
  *c.R1.DE = 0x6023
  *c.R1.HL = 0xAE06
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEF40
  *c.R1.IY = 0xB006
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0A
  mem[0x0003] = 0x08
  mem[0xEF4A] = 0xDA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0228) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0228, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6D66) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6D66, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6023) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6023, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAE06) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAE06, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEF40) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEF40, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB006) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB006, *c.R1.IY)
  }
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

func TestDDCB09(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9825
  *c.R1.BC = 0x9258
  *c.R1.DE = 0x54D5
  *c.R1.HL = 0x5E1E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9D0B
  *c.R1.IY = 0x6E58
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3B
  mem[0x0003] = 0x09
  mem[0x9D46] = 0x6F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x98A5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x98A5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x92B7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x92B7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x54D5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x54D5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E1E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E1E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9D0B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9D0B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6E58) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6E58, *c.R1.IY)
  }
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

func TestDDCB0A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD2DD
  *c.R1.BC = 0x6AAC
  *c.R1.DE = 0xE789
  *c.R1.HL = 0x9293
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1FB4
  *c.R1.IY = 0x2498
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x83
  mem[0x0003] = 0x0A
  mem[0x1F37] = 0x78
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD22C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD22C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6AAC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6AAC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3C89) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3C89, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9293) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9293, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1FB4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1FB4, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2498) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2498, *c.R1.IY)
  }
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

func TestDDCB0B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB82C
  *c.R1.BC = 0xB284
  *c.R1.DE = 0x23F8
  *c.R1.HL = 0x7E7D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCD09
  *c.R1.IY = 0x6A03
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFA
  mem[0x0003] = 0x0B
  mem[0xCD03] = 0x92
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB808) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB808, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB284) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB284, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2349) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2349, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7E7D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7E7D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCD09) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCD09, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6A03) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6A03, *c.R1.IY)
  }
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

func TestDDCB0C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDF8B
  *c.R1.BC = 0xB6CC
  *c.R1.DE = 0xEE8D
  *c.R1.HL = 0x855A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBF6B
  *c.R1.IY = 0x9B7D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x79
  mem[0x0003] = 0x0C
  mem[0xBFE4] = 0x0D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDF81) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDF81, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB6CC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB6CC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEE8D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEE8D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x865A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x865A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBF6B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBF6B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9B7D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9B7D, *c.R1.IY)
  }
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

func TestDDCB0D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBAE3
  *c.R1.BC = 0xCEEC
  *c.R1.DE = 0xBBAA
  *c.R1.HL = 0xB65E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x88BD
  *c.R1.IY = 0x503E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE4
  mem[0x0003] = 0x0D
  mem[0x88A1] = 0x1F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBA89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBA89, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCEEC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCEEC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBBAA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBBAA, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB68F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB68F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
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
  if (*c.R1.IY != 0x503E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x503E, *c.R1.IY)
  }
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

func TestDDCB0E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1C36
  *c.R1.BC = 0x890B
  *c.R1.DE = 0x7830
  *c.R1.HL = 0x060C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFD49
  *c.R1.IY = 0x5D07
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC6
  mem[0x0003] = 0x0E
  mem[0xFD0F] = 0xAD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1C81) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1C81, *c.R1.AF)
  }
  if (*c.R1.BC != 0x890B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x890B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7830) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7830, *c.R1.DE)
  }
  if (*c.R1.HL != 0x060C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x060C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFD49) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFD49, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5D07) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5D07, *c.R1.IY)
  }
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

func TestDDCB0F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF5A7
  *c.R1.BC = 0xFAD4
  *c.R1.DE = 0xFA4B
  *c.R1.HL = 0x9C53
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7447
  *c.R1.IY = 0x2267
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x57
  mem[0x0003] = 0x0F
  mem[0x749E] = 0xF8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7C28) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7C28, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFAD4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFAD4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFA4B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFA4B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9C53) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9C53, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7447) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7447, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2267) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2267, *c.R1.IY)
  }
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

func TestDDCB10(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF3AF
  *c.R1.BC = 0xBA1F
  *c.R1.DE = 0x5387
  *c.R1.HL = 0x926E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBBA2
  *c.R1.IY = 0xCA47
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4F
  mem[0x0003] = 0x10
  mem[0xBBF1] = 0x45
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF38C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF38C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8B1F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8B1F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5387) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5387, *c.R1.DE)
  }
  if (*c.R1.HL != 0x926E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x926E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBBA2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBBA2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCA47) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCA47, *c.R1.IY)
  }
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

func TestDDCB11(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2A69
  *c.R1.BC = 0xD604
  *c.R1.DE = 0xA9AA
  *c.R1.HL = 0x5B52
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1809
  *c.R1.IY = 0xD275
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEB
  mem[0x0003] = 0x11
  mem[0x17F4] = 0xD9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2AA1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2AA1, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD6B3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD6B3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA9AA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA9AA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5B52) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5B52, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1809) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1809, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD275) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD275, *c.R1.IY)
  }
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

func TestDDCB12(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9287
  *c.R1.BC = 0xC479
  *c.R1.DE = 0x26D1
  *c.R1.HL = 0x10CE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC0FB
  *c.R1.IY = 0x2777
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA6
  mem[0x0003] = 0x12
  mem[0xC0A1] = 0xE2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9285) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9285, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC479) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC479, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC5D1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC5D1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x10CE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x10CE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC0FB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC0FB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2777) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2777, *c.R1.IY)
  }
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

func TestDDCB13(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA507
  *c.R1.BC = 0x580A
  *c.R1.DE = 0xA48F
  *c.R1.HL = 0x11CD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5AC4
  *c.R1.IY = 0xCCC7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFF
  mem[0x0003] = 0x13
  mem[0x5AC3] = 0xA7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA509) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA509, *c.R1.AF)
  }
  if (*c.R1.BC != 0x580A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x580A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA44F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA44F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x11CD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x11CD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5AC4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5AC4, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCCC7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCCC7, *c.R1.IY)
  }
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

func TestDDCB14(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x294B
  *c.R1.BC = 0x5B89
  *c.R1.DE = 0x8467
  *c.R1.HL = 0x0430
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0977
  *c.R1.IY = 0xC4E8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDD
  mem[0x0003] = 0x14
  mem[0x0954] = 0x85
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2909) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2909, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5B89) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5B89, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8467) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8467, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0B30) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0B30, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0977) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0977, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC4E8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC4E8, *c.R1.IY)
  }
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

func TestDDCB15(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1FD1
  *c.R1.BC = 0x6D53
  *c.R1.DE = 0x5B7C
  *c.R1.HL = 0xA134
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEDE9
  *c.R1.IY = 0xA85C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x07
  mem[0x0003] = 0x15
  mem[0xEDF0] = 0x0E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1F0C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1F0C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6D53) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6D53, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5B7C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5B7C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA11D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA11D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEDE9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEDE9, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA85C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA85C, *c.R1.IY)
  }
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

func TestDDCB16(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDA70
  *c.R1.BC = 0xA1E4
  *c.R1.DE = 0x00B0
  *c.R1.HL = 0x92C8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x16BE
  *c.R1.IY = 0x2C95
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x45
  mem[0x0003] = 0x16
  mem[0x1703] = 0x5B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDAA0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDAA0, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA1E4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA1E4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x00B0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x00B0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x92C8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x92C8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x16BE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x16BE, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2C95) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2C95, *c.R1.IY)
  }
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

func TestDDCB17(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3300
  *c.R1.BC = 0xCBD1
  *c.R1.DE = 0x4E1A
  *c.R1.HL = 0xCD27
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB8C9
  *c.R1.IY = 0xE6D4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1C
  mem[0x0003] = 0x17
  mem[0xB8E5] = 0x7E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFCAC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFCAC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCBD1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCBD1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4E1A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4E1A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCD27) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCD27, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB8C9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB8C9, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE6D4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE6D4, *c.R1.IY)
  }
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

func TestDDCB18(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD980
  *c.R1.BC = 0x4EB5
  *c.R1.DE = 0x9CF9
  *c.R1.HL = 0xB9F1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA189
  *c.R1.IY = 0xBD7C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0E
  mem[0x0003] = 0x18
  mem[0xA197] = 0x90
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD90C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD90C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x48B5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x48B5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9CF9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9CF9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB9F1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB9F1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA189) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA189, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBD7C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBD7C, *c.R1.IY)
  }
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

func TestDDCB19(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x23B7
  *c.R1.BC = 0x595A
  *c.R1.DE = 0xA756
  *c.R1.HL = 0xCF2E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF0E7
  *c.R1.IY = 0x26E4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA3
  mem[0x0003] = 0x19
  mem[0xF08A] = 0x37
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2389) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2389, *c.R1.AF)
  }
  if (*c.R1.BC != 0x599B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x599B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA756) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA756, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCF2E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCF2E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF0E7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF0E7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x26E4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x26E4, *c.R1.IY)
  }
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

func TestDDCB1A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8B52
  *c.R1.BC = 0x7E45
  *c.R1.DE = 0xBD0F
  *c.R1.HL = 0x37A6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDE61
  *c.R1.IY = 0x9CD9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAC
  mem[0x0003] = 0x1A
  mem[0xDE0D] = 0xCC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8B24) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8B24, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7E45) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7E45, *c.R1.BC)
  }
  if (*c.R1.DE != 0x660F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x660F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x37A6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x37A6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDE61) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDE61, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9CD9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9CD9, *c.R1.IY)
  }
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

func TestDDCB1B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5C79
  *c.R1.BC = 0x1414
  *c.R1.DE = 0x811C
  *c.R1.HL = 0x5881
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB7C3
  *c.R1.IY = 0xD14F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x05
  mem[0x0003] = 0x1B
  mem[0xB7C8] = 0x91
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5C89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5C89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1414) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1414, *c.R1.BC)
  }
  if (*c.R1.DE != 0x81C8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x81C8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5881) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5881, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB7C3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB7C3, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD14F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD14F, *c.R1.IY)
  }
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

func TestDDCB1C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFAFC
  *c.R1.BC = 0x6277
  *c.R1.DE = 0x8B67
  *c.R1.HL = 0xD423
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFEF9
  *c.R1.IY = 0x4A66
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFF
  mem[0x0003] = 0x1C
  mem[0xFEF8] = 0x61
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFA25) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFA25, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6277) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6277, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8B67) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8B67, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3023) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3023, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFEF9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFEF9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4A66) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4A66, *c.R1.IY)
  }
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

func TestDDCB1D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x76A5
  *c.R1.BC = 0x324E
  *c.R1.DE = 0xE641
  *c.R1.HL = 0x58F9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5B63
  *c.R1.IY = 0xE18B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3A
  mem[0x0003] = 0x1D
  mem[0x5B9D] = 0xF3
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x76AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x76AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0x324E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x324E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE641) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE641, *c.R1.DE)
  }
  if (*c.R1.HL != 0x58F9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x58F9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5B63) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5B63, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE18B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE18B, *c.R1.IY)
  }
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

func TestDDCB1E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC5D9
  *c.R1.BC = 0xCD58
  *c.R1.DE = 0x8967
  *c.R1.HL = 0xF074
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x75B4
  *c.R1.IY = 0x693A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xCE
  mem[0x0003] = 0x1E
  mem[0x7582] = 0x91
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC589) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC589, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCD58) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCD58, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8967) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8967, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF074) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF074, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x75B4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x75B4, *c.R1.IX)
  }
  if (*c.R1.IY != 0x693A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x693A, *c.R1.IY)
  }
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

func TestDDCB1F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD28F
  *c.R1.BC = 0x7F6D
  *c.R1.DE = 0x2058
  *c.R1.HL = 0x63E3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1D9B
  *c.R1.IY = 0xBABA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA8
  mem[0x0003] = 0x1F
  mem[0x1D43] = 0xB4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDA88) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDA88, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7F6D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7F6D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2058) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2058, *c.R1.DE)
  }
  if (*c.R1.HL != 0x63E3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x63E3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1D9B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1D9B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBABA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBABA, *c.R1.IY)
  }
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

func TestDDCB20(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4CE5
  *c.R1.BC = 0x739E
  *c.R1.DE = 0xDC6C
  *c.R1.HL = 0x18F4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDC39
  *c.R1.IY = 0x8B0C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE8
  mem[0x0003] = 0x20
  mem[0xDC21] = 0x0E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4C08) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4C08, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1C9E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1C9E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDC6C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDC6C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x18F4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x18F4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDC39) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDC39, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8B0C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8B0C, *c.R1.IY)
  }
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

func TestDDCB21(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD29D
  *c.R1.BC = 0x66DD
  *c.R1.DE = 0x23EF
  *c.R1.HL = 0x9096
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3494
  *c.R1.IY = 0xB6C3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9E
  mem[0x0003] = 0x21
  mem[0x3432] = 0xF7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD2AD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD2AD, *c.R1.AF)
  }
  if (*c.R1.BC != 0x66EE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x66EE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x23EF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x23EF, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9096) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9096, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3494) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3494, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB6C3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB6C3, *c.R1.IY)
  }
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

func TestDDCB22(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFB5D
  *c.R1.BC = 0xE0D0
  *c.R1.DE = 0x7C02
  *c.R1.HL = 0xB4B7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBD3F
  *c.R1.IY = 0x385B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x43
  mem[0x0003] = 0x22
  mem[0xBD82] = 0x9F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFB29) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFB29, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE0D0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE0D0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3E02) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3E02, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB4B7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB4B7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBD3F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBD3F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x385B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x385B, *c.R1.IY)
  }
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

func TestDDCB23(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC359
  *c.R1.BC = 0x68B6
  *c.R1.DE = 0xDA84
  *c.R1.HL = 0xB990
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x22DD
  *c.R1.IY = 0xBD27
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC1
  mem[0x0003] = 0x23
  mem[0x229E] = 0xE0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC385) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC385, *c.R1.AF)
  }
  if (*c.R1.BC != 0x68B6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x68B6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDAC0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDAC0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB990) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB990, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x22DD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x22DD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBD27) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBD27, *c.R1.IY)
  }
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

func TestDDCB24(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBAF5
  *c.R1.BC = 0x7B0B
  *c.R1.DE = 0x560B
  *c.R1.HL = 0x7C33
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x31F1
  *c.R1.IY = 0xDDBD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE8
  mem[0x0003] = 0x24
  mem[0x31D9] = 0xC3
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBA81) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBA81, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7B0B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7B0B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x560B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x560B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8633) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8633, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x31F1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x31F1, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDDBD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDDBD, *c.R1.IY)
  }
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

func TestDDCB25(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x43BB
  *c.R1.BC = 0xA21B
  *c.R1.DE = 0x2347
  *c.R1.HL = 0xAE4A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCC63
  *c.R1.IY = 0xFC94
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC1
  mem[0x0003] = 0x25
  mem[0xCC24] = 0xEB
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4381) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4381, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA21B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA21B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2347) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2347, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAED6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAED6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCC63) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCC63, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFC94) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFC94, *c.R1.IY)
  }
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

func TestDDCB26(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2065
  *c.R1.BC = 0xFF37
  *c.R1.DE = 0xE41F
  *c.R1.HL = 0x70E7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6528
  *c.R1.IY = 0xA0D5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF7
  mem[0x0003] = 0x26
  mem[0x651F] = 0x89
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2005) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2005, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFF37) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFF37, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE41F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE41F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x70E7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x70E7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6528) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6528, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA0D5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA0D5, *c.R1.IY)
  }
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

func TestDDCB27(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA806
  *c.R1.BC = 0x5669
  *c.R1.DE = 0x1BEE
  *c.R1.HL = 0xF62C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1F69
  *c.R1.IY = 0x3418
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC3
  mem[0x0003] = 0x27
  mem[0x1F2C] = 0xAC
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5809) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5809, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5669) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5669, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1BEE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1BEE, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF62C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF62C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1F69) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1F69, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3418) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3418, *c.R1.IY)
  }
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

func TestDDCB28(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7AFD
  *c.R1.BC = 0x64B8
  *c.R1.DE = 0x51F7
  *c.R1.HL = 0x7164
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x999B
  *c.R1.IY = 0x8857
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB6
  mem[0x0003] = 0x28
  mem[0x9951] = 0x24
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7A04) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7A04, *c.R1.AF)
  }
  if (*c.R1.BC != 0x12B8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x12B8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x51F7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x51F7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7164) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7164, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x999B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x999B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8857) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8857, *c.R1.IY)
  }
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

func TestDDCB29(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0404
  *c.R1.BC = 0xB794
  *c.R1.DE = 0x323F
  *c.R1.HL = 0xFD34
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x20E7
  *c.R1.IY = 0xC753
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9C
  mem[0x0003] = 0x29
  mem[0x2083] = 0x82
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0480) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0480, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB7C1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB7C1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x323F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x323F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFD34) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFD34, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x20E7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x20E7, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC753) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC753, *c.R1.IY)
  }
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

func TestDDCB2A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4524
  *c.R1.BC = 0xAFDE
  *c.R1.DE = 0x0C08
  *c.R1.HL = 0x75D7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9505
  *c.R1.IY = 0xB624
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD8
  mem[0x0003] = 0x2A
  mem[0x94DD] = 0x7C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4528) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4528, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAFDE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAFDE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3E08) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3E08, *c.R1.DE)
  }
  if (*c.R1.HL != 0x75D7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x75D7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9505) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9505, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB624) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB624, *c.R1.IY)
  }
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

func TestDDCB2B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8324
  *c.R1.BC = 0xE290
  *c.R1.DE = 0x26BE
  *c.R1.HL = 0x7DDD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB484
  *c.R1.IY = 0x571C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBD
  mem[0x0003] = 0x2B
  mem[0xB441] = 0x44
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8324) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8324, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE290) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE290, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2622) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2622, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7DDD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7DDD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB484) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB484, *c.R1.IX)
  }
  if (*c.R1.IY != 0x571C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x571C, *c.R1.IY)
  }
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

func TestDDCB2C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC688
  *c.R1.BC = 0x0C94
  *c.R1.DE = 0x6E4B
  *c.R1.HL = 0x7DC7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFE28
  *c.R1.IY = 0xDC80
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2C
  mem[0x0003] = 0x2C
  mem[0xFE54] = 0x81
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC685) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC685, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0C94) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0C94, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6E4B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6E4B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC0C7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC0C7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFE28) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFE28, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDC80) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDC80, *c.R1.IY)
  }
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

func TestDDCB2D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCE28
  *c.R1.BC = 0xD2AE
  *c.R1.DE = 0xC9BE
  *c.R1.HL = 0x4236
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB4ED
  *c.R1.IY = 0x6DE3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9B
  mem[0x0003] = 0x2D
  mem[0xB488] = 0x44
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCE24) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCE24, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD2AE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD2AE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC9BE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC9BE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4222) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4222, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB4ED) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB4ED, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6DE3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6DE3, *c.R1.IY)
  }
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

func TestDDCB2E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x50B0
  *c.R1.BC = 0xDE74
  *c.R1.DE = 0xECA8
  *c.R1.HL = 0x83FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x69D8
  *c.R1.IY = 0x75C7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3D
  mem[0x0003] = 0x2E
  mem[0x6A15] = 0x05
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5001) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5001, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDE74) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDE74, *c.R1.BC)
  }
  if (*c.R1.DE != 0xECA8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xECA8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x83FF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x83FF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x69D8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x69D8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x75C7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x75C7, *c.R1.IY)
  }
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

func TestDDCB2F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAEC6
  *c.R1.BC = 0x759B
  *c.R1.DE = 0x3059
  *c.R1.HL = 0x01B9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7A30
  *c.R1.IY = 0xDD56
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD3
  mem[0x0003] = 0x2F
  mem[0x7A03] = 0xF2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF9AC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF9AC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x759B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x759B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3059) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3059, *c.R1.DE)
  }
  if (*c.R1.HL != 0x01B9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x01B9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7A30) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7A30, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDD56) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDD56, *c.R1.IY)
  }
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

func TestDDCB30(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3C89
  *c.R1.BC = 0x96AD
  *c.R1.DE = 0x9CC7
  *c.R1.HL = 0xA68C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEEE8
  *c.R1.IY = 0x5A80
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDF
  mem[0x0003] = 0x30
  mem[0xEEC7] = 0x32
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C24) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C24, *c.R1.AF)
  }
  if (*c.R1.BC != 0x65AD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x65AD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9CC7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9CC7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA68C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA68C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEEE8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEEE8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5A80) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5A80, *c.R1.IY)
  }
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

func TestDDCB31(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEBF5
  *c.R1.BC = 0x41E9
  *c.R1.DE = 0x929B
  *c.R1.HL = 0x7D47
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF22D
  *c.R1.IY = 0x8943
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x49
  mem[0x0003] = 0x31
  mem[0xF276] = 0xCD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEB89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEB89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x419B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x419B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x929B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x929B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7D47) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7D47, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF22D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF22D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8943) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8943, *c.R1.IY)
  }
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

func TestDDCB32(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9A1B
  *c.R1.BC = 0xAA64
  *c.R1.DE = 0x4209
  *c.R1.HL = 0x01AD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x579F
  *c.R1.IY = 0xEC4C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE0
  mem[0x0003] = 0x32
  mem[0x577F] = 0xE2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9A85) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9A85, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAA64) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAA64, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC509) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC509, *c.R1.DE)
  }
  if (*c.R1.HL != 0x01AD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x01AD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x579F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x579F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEC4C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEC4C, *c.R1.IY)
  }
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

func TestDDCB33(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB8B1
  *c.R1.BC = 0xB854
  *c.R1.DE = 0x524F
  *c.R1.HL = 0x9599
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEFAC
  *c.R1.IY = 0xD9EC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC9
  mem[0x0003] = 0x33
  mem[0xEF75] = 0x0B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB804) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB804, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB854) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB854, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5217) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5217, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9599) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9599, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEFAC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEFAC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD9EC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD9EC, *c.R1.IY)
  }
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

func TestDDCB34(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCD3C
  *c.R1.BC = 0x4432
  *c.R1.DE = 0x20D4
  *c.R1.HL = 0x0B3E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAB48
  *c.R1.IY = 0xC95F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x49
  mem[0x0003] = 0x34
  mem[0xAB91] = 0xEF
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCD89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCD89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4432) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4432, *c.R1.BC)
  }
  if (*c.R1.DE != 0x20D4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x20D4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDF3E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDF3E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAB48) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAB48, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC95F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC95F, *c.R1.IY)
  }
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

func TestDDCB35(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDEB1
  *c.R1.BC = 0xC6FC
  *c.R1.DE = 0x696D
  *c.R1.HL = 0x150D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEB1A
  *c.R1.IY = 0x4A12
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB9
  mem[0x0003] = 0x35
  mem[0xEAD3] = 0x8F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDE09) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDE09, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC6FC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC6FC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x696D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x696D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x151F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x151F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEB1A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEB1A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4A12) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4A12, *c.R1.IY)
  }
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

func TestDDCB36(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3D81
  *c.R1.BC = 0x443B
  *c.R1.DE = 0xFF21
  *c.R1.HL = 0x63E3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x132E
  *c.R1.IY = 0xFB39
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB4
  mem[0x0003] = 0x36
  mem[0x12E2] = 0x02
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3D04) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3D04, *c.R1.AF)
  }
  if (*c.R1.BC != 0x443B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x443B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFF21) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFF21, *c.R1.DE)
  }
  if (*c.R1.HL != 0x63E3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x63E3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x132E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x132E, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFB39) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFB39, *c.R1.IY)
  }
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

func TestDDCB37(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x72D9
  *c.R1.BC = 0xBFC9
  *c.R1.DE = 0xA69A
  *c.R1.HL = 0xEC0B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5077
  *c.R1.IY = 0x4E3E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC6
  mem[0x0003] = 0x37
  mem[0x503D] = 0x3D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7B2C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7B2C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBFC9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBFC9, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA69A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA69A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEC0B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEC0B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5077) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5077, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4E3E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4E3E, *c.R1.IY)
  }
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

func TestDDCB38(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3C64
  *c.R1.BC = 0xB1EE
  *c.R1.DE = 0x38E1
  *c.R1.HL = 0xAE9F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF695
  *c.R1.IY = 0x44B3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8E
  mem[0x0003] = 0x38
  mem[0xF623] = 0x5E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C28) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C28, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2FEE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2FEE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x38E1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x38E1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAE9F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAE9F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF695) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF695, *c.R1.IX)
  }
  if (*c.R1.IY != 0x44B3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x44B3, *c.R1.IY)
  }
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

func TestDDCB39(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x05D6
  *c.R1.BC = 0x9AAD
  *c.R1.DE = 0xA2DB
  *c.R1.HL = 0xDF75
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA895
  *c.R1.IY = 0xE243
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDC
  mem[0x0003] = 0x39
  mem[0xA871] = 0x83
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0505) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0505, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9A41) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9A41, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA2DB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA2DB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDF75) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDF75, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA895) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA895, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE243) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE243, *c.R1.IY)
  }
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

func TestDDCB3A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0E22
  *c.R1.BC = 0x0B9F
  *c.R1.DE = 0x873B
  *c.R1.HL = 0xC01D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2591
  *c.R1.IY = 0x49C3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0D
  mem[0x0003] = 0x3A
  mem[0x259E] = 0x89
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0E05) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0E05, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0B9F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0B9F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x443B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x443B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC01D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC01D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2591) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2591, *c.R1.IX)
  }
  if (*c.R1.IY != 0x49C3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x49C3, *c.R1.IY)
  }
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

func TestDDCB3B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1BD9
  *c.R1.BC = 0xC795
  *c.R1.DE = 0xD8AE
  *c.R1.HL = 0x7CCF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6FED
  *c.R1.IY = 0x09DC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x20
  mem[0x0003] = 0x3B
  mem[0x700D] = 0xA9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1B01) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1B01, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC795) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC795, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD854) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD854, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7CCF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7CCF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6FED) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6FED, *c.R1.IX)
  }
  if (*c.R1.IY != 0x09DC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x09DC, *c.R1.IY)
  }
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

func TestDDCB3C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB651
  *c.R1.BC = 0xBDF7
  *c.R1.DE = 0xFCA3
  *c.R1.HL = 0x7529
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF53B
  *c.R1.IY = 0x018B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE1
  mem[0x0003] = 0x3C
  mem[0xF51C] = 0xD0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB628) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB628, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBDF7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBDF7, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFCA3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFCA3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6829) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6829, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF53B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF53B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x018B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x018B, *c.R1.IY)
  }
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

func TestDDCB3D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2A2D
  *c.R1.BC = 0x6E6E
  *c.R1.DE = 0xCFBD
  *c.R1.HL = 0x1DB5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0320
  *c.R1.IY = 0x6AB0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBE
  mem[0x0003] = 0x3D
  mem[0x02DE] = 0x58
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2A28) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2A28, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6E6E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6E6E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCFBD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCFBD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1D2C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1D2C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0320) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0320, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6AB0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6AB0, *c.R1.IY)
  }
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

func TestDDCB3E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x39B8
  *c.R1.BC = 0xB26E
  *c.R1.DE = 0xB670
  *c.R1.HL = 0xB8A2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x784A
  *c.R1.IY = 0x7840
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0A
  mem[0x0003] = 0x3E
  mem[0x7854] = 0x5D
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x392D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x392D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB26E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB26E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB670) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB670, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB8A2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB8A2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x784A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x784A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7840) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7840, *c.R1.IY)
  }
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

func TestDDCB3F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2A17
  *c.R1.BC = 0x429D
  *c.R1.DE = 0xD8C0
  *c.R1.HL = 0xE069
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3488
  *c.R1.IY = 0x7150
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x31
  mem[0x0003] = 0x3F
  mem[0x34B9] = 0x04
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x429D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x429D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD8C0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD8C0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE069) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE069, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3488) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3488, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7150) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7150, *c.R1.IY)
  }
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

func TestDDCB40(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x119B
  *c.R1.BC = 0xF6BA
  *c.R1.DE = 0x079E
  *c.R1.HL = 0x0E41
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8C01
  *c.R1.IY = 0xCD21
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBD
  mem[0x0003] = 0x40
  mem[0x8BBE] = 0xE7
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1119) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1119, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF6BA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF6BA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x079E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x079E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0E41) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0E41, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8C01) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8C01, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCD21) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCD21, *c.R1.IY)
  }
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

func TestDDCB41(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x22B3
  *c.R1.BC = 0xC4B0
  *c.R1.DE = 0x575B
  *c.R1.HL = 0x66B4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCDCF
  *c.R1.IY = 0xA25C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x52
  mem[0x0003] = 0x41
  mem[0xCE21] = 0x75
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2219) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2219, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC4B0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC4B0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x575B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x575B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x66B4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x66B4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCDCF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCDCF, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA25C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA25C, *c.R1.IY)
  }
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

func TestDDCB42(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAF5E
  *c.R1.BC = 0x7720
  *c.R1.DE = 0xAA95
  *c.R1.HL = 0x3B0A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF03A
  *c.R1.IY = 0x856A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x1E
  mem[0x0003] = 0x42
  mem[0xF058] = 0x90
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAF74) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAF74, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7720) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7720, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAA95) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAA95, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3B0A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3B0A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF03A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF03A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x856A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x856A, *c.R1.IY)
  }
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

func TestDDCB43(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7FA6
  *c.R1.BC = 0xB699
  *c.R1.DE = 0x5E71
  *c.R1.HL = 0x1827
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE8B6
  *c.R1.IY = 0x96A8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBC
  mem[0x0003] = 0x43
  mem[0xE872] = 0x6B
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7F38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7F38, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB699) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB699, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5E71) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5E71, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1827) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1827, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE8B6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE8B6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x96A8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x96A8, *c.R1.IY)
  }
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

func TestDDCB44(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5FAA
  *c.R1.BC = 0xDE05
  *c.R1.DE = 0x12FD
  *c.R1.HL = 0xF73B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEE0A
  *c.R1.IY = 0x6634
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE8
  mem[0x0003] = 0x44
  mem[0xEDF2] = 0x62
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5F7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5F7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDE05) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDE05, *c.R1.BC)
  }
  if (*c.R1.DE != 0x12FD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x12FD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF73B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF73B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
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
  if (*c.R1.IY != 0x6634) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6634, *c.R1.IY)
  }
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

func TestDDCB45(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEAC7
  *c.R1.BC = 0x699C
  *c.R1.DE = 0x47D3
  *c.R1.HL = 0x89C3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA2BE
  *c.R1.IY = 0xD81E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x02
  mem[0x0003] = 0x45
  mem[0xA2C0] = 0x55
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEA31) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEA31, *c.R1.AF)
  }
  if (*c.R1.BC != 0x699C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x699C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x47D3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x47D3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x89C3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x89C3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA2BE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA2BE, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD81E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD81E, *c.R1.IY)
  }
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

func TestDDCB46(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x60DE
  *c.R1.BC = 0xAC1D
  *c.R1.DE = 0x4173
  *c.R1.HL = 0xF92A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA39F
  *c.R1.IY = 0x12E5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE2
  mem[0x0003] = 0x46
  mem[0xA381] = 0xD5
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6030) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6030, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAC1D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAC1D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4173) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4173, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF92A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF92A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA39F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA39F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x12E5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x12E5, *c.R1.IY)
  }
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

func TestDDCB47(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1B1A
  *c.R1.BC = 0xF7C0
  *c.R1.DE = 0x22F6
  *c.R1.HL = 0x5253
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5227
  *c.R1.IY = 0x919D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7A
  mem[0x0003] = 0x47
  mem[0x52A1] = 0x6A
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1B54) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1B54, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF7C0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF7C0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x22F6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x22F6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5253) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5253, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5227) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5227, *c.R1.IX)
  }
  if (*c.R1.IY != 0x919D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x919D, *c.R1.IY)
  }
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

func TestDDCB48(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x721A
  *c.R1.BC = 0x4509
  *c.R1.DE = 0xD68F
  *c.R1.HL = 0x3B3D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2746
  *c.R1.IY = 0x7F97
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x13
  mem[0x0003] = 0x48
  mem[0x2759] = 0xA8
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7274) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7274, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4509) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4509, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD68F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD68F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3B3D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3B3D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2746) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2746, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7F97) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7F97, *c.R1.IY)
  }
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

func TestDDCB49(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7FE9
  *c.R1.BC = 0xDA22
  *c.R1.DE = 0xEA9C
  *c.R1.HL = 0xF480
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x41C6
  *c.R1.IY = 0x75A9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x94
  mem[0x0003] = 0x49
  mem[0x415A] = 0x26
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7F11) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7F11, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDA22) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDA22, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEA9C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEA9C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF480) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF480, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x41C6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x41C6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x75A9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x75A9, *c.R1.IY)
  }
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

func TestDDCB4A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF16D
  *c.R1.BC = 0xE6C3
  *c.R1.DE = 0x5A42
  *c.R1.HL = 0x8B21
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBFEB
  *c.R1.IY = 0xE383
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3B
  mem[0x0003] = 0x4A
  mem[0xC026] = 0xB5
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF155) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF155, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE6C3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE6C3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5A42) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5A42, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8B21) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8B21, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBFEB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBFEB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE383) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE383, *c.R1.IY)
  }
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

func TestDDCB4B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1050
  *c.R1.BC = 0x880A
  *c.R1.DE = 0x52B2
  *c.R1.HL = 0xFB1B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC239
  *c.R1.IY = 0x6B40
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB0
  mem[0x0003] = 0x4B
  mem[0xC1E9] = 0x18
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1054) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1054, *c.R1.AF)
  }
  if (*c.R1.BC != 0x880A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x880A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x52B2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x52B2, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFB1B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFB1B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC239) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC239, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6B40) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6B40, *c.R1.IY)
  }
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

func TestDDCB4C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0538
  *c.R1.BC = 0xBC63
  *c.R1.DE = 0xF081
  *c.R1.HL = 0x0A55
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x874C
  *c.R1.IY = 0x80A3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x97
  mem[0x0003] = 0x4C
  mem[0x86E3] = 0x63
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0510) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0510, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBC63) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBC63, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF081) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF081, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0A55) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0A55, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x874C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x874C, *c.R1.IX)
  }
  if (*c.R1.IY != 0x80A3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x80A3, *c.R1.IY)
  }
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

func TestDDCB4D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7F8C
  *c.R1.BC = 0x32B4
  *c.R1.DE = 0x03D5
  *c.R1.HL = 0xEF66
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7D2A
  *c.R1.IY = 0x03BC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x15
  mem[0x0003] = 0x4D
  mem[0x7D3F] = 0x60
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7F7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7F7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x32B4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x32B4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x03D5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x03D5, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEF66) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEF66, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7D2A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7D2A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x03BC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x03BC, *c.R1.IY)
  }
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

func TestDDCB4E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7C67
  *c.R1.BC = 0xFA92
  *c.R1.DE = 0xB4D0
  *c.R1.HL = 0x9F23
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEADE
  *c.R1.IY = 0x1785
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB0
  mem[0x0003] = 0x4E
  mem[0xEA8E] = 0x3B
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7C39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7C39, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFA92) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFA92, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB4D0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB4D0, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9F23) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9F23, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEADE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEADE, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1785) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1785, *c.R1.IY)
  }
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

func TestDDCB4F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x725C
  *c.R1.BC = 0x257B
  *c.R1.DE = 0xDB73
  *c.R1.HL = 0x2478
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x88C0
  *c.R1.IY = 0xF151
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8B
  mem[0x0003] = 0x4F
  mem[0x884B] = 0x4C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x725C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x725C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x257B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x257B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDB73) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDB73, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2478) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2478, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x88C0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x88C0, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF151) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF151, *c.R1.IY)
  }
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

func TestDDCB50(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x35F4
  *c.R1.BC = 0x8E51
  *c.R1.DE = 0x406C
  *c.R1.HL = 0x2E3C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDAF2
  *c.R1.IY = 0x413C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x12
  mem[0x0003] = 0x50
  mem[0xDB04] = 0x00
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x355C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x355C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8E51) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8E51, *c.R1.BC)
  }
  if (*c.R1.DE != 0x406C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x406C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2E3C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2E3C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDAF2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDAF2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x413C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x413C, *c.R1.IY)
  }
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

func TestDDCB51(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA630
  *c.R1.BC = 0xBA85
  *c.R1.DE = 0xC88C
  *c.R1.HL = 0xE86C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x84B2
  *c.R1.IY = 0xCD8E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x18
  mem[0x0003] = 0x51
  mem[0x84CA] = 0x1C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA610) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA610, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBA85) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBA85, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC88C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC88C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE86C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE86C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x84B2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x84B2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCD8E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCD8E, *c.R1.IY)
  }
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

func TestDDCB52(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCB88
  *c.R1.BC = 0x1220
  *c.R1.DE = 0x1103
  *c.R1.HL = 0xA868
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6156
  *c.R1.IY = 0xCFAC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x42
  mem[0x0003] = 0x52
  mem[0x6198] = 0x53
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCB74) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCB74, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1220) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1220, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1103) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1103, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA868) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA868, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6156) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6156, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCFAC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCFAC, *c.R1.IY)
  }
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

func TestDDCB53(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5EB3
  *c.R1.BC = 0x569E
  *c.R1.DE = 0xF76D
  *c.R1.HL = 0x88C6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAE45
  *c.R1.IY = 0x623E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE3
  mem[0x0003] = 0x53
  mem[0xAE28] = 0xD6
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5E39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5E39, *c.R1.AF)
  }
  if (*c.R1.BC != 0x569E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x569E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF76D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF76D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x88C6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x88C6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAE45) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAE45, *c.R1.IX)
  }
  if (*c.R1.IY != 0x623E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x623E, *c.R1.IY)
  }
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

func TestDDCB54(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC3C9
  *c.R1.BC = 0x76FE
  *c.R1.DE = 0xF1FF
  *c.R1.HL = 0x416E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEFD5
  *c.R1.IY = 0x7576
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7D
  mem[0x0003] = 0x54
  mem[0xF052] = 0x5D
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC331) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC331, *c.R1.AF)
  }
  if (*c.R1.BC != 0x76FE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x76FE, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF1FF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF1FF, *c.R1.DE)
  }
  if (*c.R1.HL != 0x416E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x416E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEFD5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEFD5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7576) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7576, *c.R1.IY)
  }
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

func TestDDCB55(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7068
  *c.R1.BC = 0xDCD0
  *c.R1.DE = 0x8345
  *c.R1.HL = 0xD498
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF352
  *c.R1.IY = 0xA88B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x88
  mem[0x0003] = 0x55
  mem[0xF2DA] = 0x03
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7074) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7074, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDCD0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDCD0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8345) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8345, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD498) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD498, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
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
  if (*c.R1.IY != 0xA88B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA88B, *c.R1.IY)
  }
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

func TestDDCB56(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9128
  *c.R1.BC = 0x2CB8
  *c.R1.DE = 0x571C
  *c.R1.HL = 0xF4FD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6D30
  *c.R1.IY = 0xAEC2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x57
  mem[0x0003] = 0x56
  mem[0x6D87] = 0x61
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x917C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x917C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2CB8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2CB8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x571C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x571C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF4FD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF4FD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6D30) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6D30, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAEC2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAEC2, *c.R1.IY)
  }
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

func TestDDCB57(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3CA7
  *c.R1.BC = 0x541A
  *c.R1.DE = 0x027C
  *c.R1.HL = 0xC0B4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5833
  *c.R1.IY = 0x160A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x06
  mem[0x0003] = 0x57
  mem[0x5839] = 0x1D
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C19) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C19, *c.R1.AF)
  }
  if (*c.R1.BC != 0x541A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x541A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x027C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x027C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC0B4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC0B4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5833) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5833, *c.R1.IX)
  }
  if (*c.R1.IY != 0x160A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x160A, *c.R1.IY)
  }
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

func TestDDCB58(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC650
  *c.R1.BC = 0xE1A8
  *c.R1.DE = 0x9D6C
  *c.R1.HL = 0xBEC3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6A46
  *c.R1.IY = 0xB66C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x83
  mem[0x0003] = 0x58
  mem[0x69C9] = 0x0F
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC638) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC638, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE1A8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE1A8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9D6C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9D6C, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBEC3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBEC3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6A46) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6A46, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB66C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB66C, *c.R1.IY)
  }
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

func TestDDCB59(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAD07
  *c.R1.BC = 0x9BDA
  *c.R1.DE = 0xB7EE
  *c.R1.HL = 0x63C4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9195
  *c.R1.IY = 0x9703
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDB
  mem[0x0003] = 0x59
  mem[0x9170] = 0x10
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAD55) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAD55, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9BDA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9BDA, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB7EE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB7EE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x63C4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x63C4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9195) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9195, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9703) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9703, *c.R1.IY)
  }
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

func TestDDCB5A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x80C0
  *c.R1.BC = 0x5105
  *c.R1.DE = 0x36B0
  *c.R1.HL = 0xA37C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0DE0
  *c.R1.IY = 0xCE7F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD1
  mem[0x0003] = 0x5A
  mem[0x0DB1] = 0xBE
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8018) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8018, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5105) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5105, *c.R1.BC)
  }
  if (*c.R1.DE != 0x36B0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x36B0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA37C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA37C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0DE0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0DE0, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCE7F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCE7F, *c.R1.IY)
  }
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

func TestDDCB5B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2A8D
  *c.R1.BC = 0x083D
  *c.R1.DE = 0x1409
  *c.R1.HL = 0x06BA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x62AD
  *c.R1.IY = 0xBAFF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD5
  mem[0x0003] = 0x5B
  mem[0x6282] = 0x67
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2A75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2A75, *c.R1.AF)
  }
  if (*c.R1.BC != 0x083D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x083D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1409) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1409, *c.R1.DE)
  }
  if (*c.R1.HL != 0x06BA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x06BA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x62AD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x62AD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBAFF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBAFF, *c.R1.IY)
  }
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

func TestDDCB5C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4CA4
  *c.R1.BC = 0xE502
  *c.R1.DE = 0xD23C
  *c.R1.HL = 0x6DA8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9DC6
  *c.R1.IY = 0x6F04
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5C
  mem[0x0003] = 0x5C
  mem[0x9E22] = 0xC9
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4C18) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4C18, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE502) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE502, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD23C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD23C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6DA8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6DA8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9DC6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9DC6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6F04) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6F04, *c.R1.IY)
  }
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

func TestDDCB5D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7E39
  *c.R1.BC = 0x511B
  *c.R1.DE = 0x3CFA
  *c.R1.HL = 0x60D3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD193
  *c.R1.IY = 0x3FE9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFF
  mem[0x0003] = 0x5D
  mem[0xD192] = 0x0D
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7E11) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7E11, *c.R1.AF)
  }
  if (*c.R1.BC != 0x511B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x511B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3CFA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3CFA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x60D3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x60D3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD193) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD193, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3FE9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3FE9, *c.R1.IY)
  }
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

func TestDDCB5E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCEF1
  *c.R1.BC = 0x0235
  *c.R1.DE = 0xE2B1
  *c.R1.HL = 0x7A4C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xED14
  *c.R1.IY = 0xD0D6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x62
  mem[0x0003] = 0x5E
  mem[0xED76] = 0xA7
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCE7D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCE7D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0235) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0235, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE2B1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE2B1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7A4C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7A4C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xED14) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xED14, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD0D6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD0D6, *c.R1.IY)
  }
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

func TestDDCB5F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x094F
  *c.R1.BC = 0x20A8
  *c.R1.DE = 0x52E1
  *c.R1.HL = 0xD783
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDF46
  *c.R1.IY = 0xDA41
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3F
  mem[0x0003] = 0x5F
  mem[0xDF85] = 0x9E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0919) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0919, *c.R1.AF)
  }
  if (*c.R1.BC != 0x20A8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x20A8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x52E1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x52E1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD783) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD783, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDF46) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDF46, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDA41) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDA41, *c.R1.IY)
  }
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

func TestDDCB60(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x42CE
  *c.R1.BC = 0x0713
  *c.R1.DE = 0xDC90
  *c.R1.HL = 0x2C89
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x32A2
  *c.R1.IY = 0xC4D4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x65
  mem[0x0003] = 0x60
  mem[0x3307] = 0x2E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4274) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4274, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0713) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0713, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDC90) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDC90, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2C89) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2C89, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x32A2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x32A2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC4D4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC4D4, *c.R1.IY)
  }
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

func TestDDCB61(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1B36
  *c.R1.BC = 0x1403
  *c.R1.DE = 0x8B9B
  *c.R1.HL = 0xC221
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x36CB
  *c.R1.IY = 0x93D4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA8
  mem[0x0003] = 0x61
  mem[0x3673] = 0xBC
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1B30) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1B30, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1403) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1403, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8B9B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8B9B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC221) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC221, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x36CB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x36CB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x93D4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x93D4, *c.R1.IY)
  }
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

func TestDDCB62(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x361B
  *c.R1.BC = 0x4055
  *c.R1.DE = 0x650A
  *c.R1.HL = 0x3F98
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0ACC
  *c.R1.IY = 0xA102
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD9
  mem[0x0003] = 0x62
  mem[0x0AA5] = 0xEA
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x365D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x365D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4055) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4055, *c.R1.BC)
  }
  if (*c.R1.DE != 0x650A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x650A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3F98) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3F98, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0ACC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0ACC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA102) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA102, *c.R1.IY)
  }
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

func TestDDCB63(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6548
  *c.R1.BC = 0x08DF
  *c.R1.DE = 0x3CEB
  *c.R1.HL = 0x6D24
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE679
  *c.R1.IY = 0xF98E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x79
  mem[0x0003] = 0x63
  mem[0xE6F2] = 0x83
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6574) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6574, *c.R1.AF)
  }
  if (*c.R1.BC != 0x08DF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x08DF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3CEB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3CEB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6D24) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6D24, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE679) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE679, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF98E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF98E, *c.R1.IY)
  }
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

func TestDDCB64(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3C22
  *c.R1.BC = 0xE2A7
  *c.R1.DE = 0x6DA9
  *c.R1.HL = 0xC346
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xECFB
  *c.R1.IY = 0x85B6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x71
  mem[0x0003] = 0x64
  mem[0xED6C] = 0x52
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3C38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3C38, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE2A7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE2A7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6DA9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6DA9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC346) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC346, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xECFB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xECFB, *c.R1.IX)
  }
  if (*c.R1.IY != 0x85B6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x85B6, *c.R1.IY)
  }
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

func TestDDCB65(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x09BD
  *c.R1.BC = 0x0ABB
  *c.R1.DE = 0x3AFA
  *c.R1.HL = 0x91F5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7779
  *c.R1.IY = 0xAEF5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x67
  mem[0x0003] = 0x65
  mem[0x77E0] = 0xF5
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0931) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0931, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0ABB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0ABB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3AFA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3AFA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x91F5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x91F5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7779) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7779, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAEF5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAEF5, *c.R1.IY)
  }
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

func TestDDCB66(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCCBC
  *c.R1.BC = 0xD301
  *c.R1.DE = 0x9B66
  *c.R1.HL = 0x40FB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEE15
  *c.R1.IY = 0x0D23
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x63
  mem[0x0003] = 0x66
  mem[0xEE78] = 0x70
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCC38) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCC38, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD301) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD301, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9B66) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9B66, *c.R1.DE)
  }
  if (*c.R1.HL != 0x40FB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x40FB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEE15) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEE15, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0D23) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0D23, *c.R1.IY)
  }
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

func TestDDCB67(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xECCB
  *c.R1.BC = 0x342F
  *c.R1.DE = 0xBE3E
  *c.R1.HL = 0xA79B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEEA1
  *c.R1.IY = 0xDFAE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD7
  mem[0x0003] = 0x67
  mem[0xEE78] = 0x06
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEC7D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEC7D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x342F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x342F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBE3E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBE3E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA79B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA79B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEEA1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEEA1, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDFAE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDFAE, *c.R1.IY)
  }
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

func TestDDCB68(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8E51
  *c.R1.BC = 0x0063
  *c.R1.DE = 0x49AD
  *c.R1.HL = 0xB7D4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE968
  *c.R1.IY = 0x864E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB1
  mem[0x0003] = 0x68
  mem[0xE919] = 0x20
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8E39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8E39, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0063) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0063, *c.R1.BC)
  }
  if (*c.R1.DE != 0x49AD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x49AD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB7D4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB7D4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE968) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE968, *c.R1.IX)
  }
  if (*c.R1.IY != 0x864E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x864E, *c.R1.IY)
  }
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

func TestDDCB69(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9F11
  *c.R1.BC = 0x42B5
  *c.R1.DE = 0x74FE
  *c.R1.HL = 0x1116
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x33F4
  *c.R1.IY = 0x46C2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE8
  mem[0x0003] = 0x69
  mem[0x33DC] = 0x4F
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9F75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9F75, *c.R1.AF)
  }
  if (*c.R1.BC != 0x42B5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x42B5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x74FE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x74FE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1116) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1116, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x33F4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x33F4, *c.R1.IX)
  }
  if (*c.R1.IY != 0x46C2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x46C2, *c.R1.IY)
  }
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

func TestDDCB6A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4632
  *c.R1.BC = 0x0BD8
  *c.R1.DE = 0x0018
  *c.R1.HL = 0x1AC3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x86B6
  *c.R1.IY = 0x1DD2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x33
  mem[0x0003] = 0x6A
  mem[0x86E9] = 0x1C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4654) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4654, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0BD8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0BD8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0018) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0018, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1AC3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1AC3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x86B6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x86B6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1DD2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1DD2, *c.R1.IY)
  }
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

func TestDDCB6B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7A76
  *c.R1.BC = 0xF79F
  *c.R1.DE = 0xA78E
  *c.R1.HL = 0xF867
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x187B
  *c.R1.IY = 0x0023
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x11
  mem[0x0003] = 0x6B
  mem[0x188C] = 0xBC
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7A18) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7A18, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF79F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF79F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA78E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA78E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF867) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF867, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x187B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x187B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0023) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0023, *c.R1.IY)
  }
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

func TestDDCB6C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDD91
  *c.R1.BC = 0x1F1E
  *c.R1.DE = 0xC1E1
  *c.R1.HL = 0x0EA7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3E21
  *c.R1.IY = 0xF544
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5E
  mem[0x0003] = 0x6C
  mem[0x3E7F] = 0x2A
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDD39) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDD39, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1F1E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1F1E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC1E1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC1E1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0EA7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0EA7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3E21) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3E21, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF544) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF544, *c.R1.IY)
  }
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

func TestDDCB6D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDEBF
  *c.R1.BC = 0x9AE4
  *c.R1.DE = 0xFD24
  *c.R1.HL = 0xB3C2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE314
  *c.R1.IY = 0xAD84
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDD
  mem[0x0003] = 0x6D
  mem[0xE2F1] = 0x41
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDE75) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDE75, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9AE4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9AE4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFD24) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFD24, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB3C2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB3C2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE314) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE314, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAD84) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAD84, *c.R1.IY)
  }
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

func TestDDCB6E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCA75
  *c.R1.BC = 0x9F16
  *c.R1.DE = 0xC700
  *c.R1.HL = 0x1DCE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3086
  *c.R1.IY = 0xD68E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB2
  mem[0x0003] = 0x6E
  mem[0x3038] = 0x3F
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCA31) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCA31, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9F16) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9F16, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC700) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC700, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1DCE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1DCE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3086) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3086, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD68E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD68E, *c.R1.IY)
  }
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

func TestDDCB6F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD4CD
  *c.R1.BC = 0x0B39
  *c.R1.DE = 0x3E2E
  *c.R1.HL = 0xC06E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFC1B
  *c.R1.IY = 0xD592
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBE
  mem[0x0003] = 0x6F
  mem[0xFBD9] = 0x56
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD47D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD47D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0B39) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0B39, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3E2E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3E2E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC06E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC06E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFC1B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFC1B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD592) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD592, *c.R1.IY)
  }
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

func TestDDCB70(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF901
  *c.R1.BC = 0x09B8
  *c.R1.DE = 0x43F8
  *c.R1.HL = 0x2A76
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x042C
  *c.R1.IY = 0x7F2D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB5
  mem[0x0003] = 0x70
  mem[0x03E1] = 0x74
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF911) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF911, *c.R1.AF)
  }
  if (*c.R1.BC != 0x09B8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x09B8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x43F8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x43F8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2A76) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2A76, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x042C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x042C, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7F2D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7F2D, *c.R1.IY)
  }
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

func TestDDCB71(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAC78
  *c.R1.BC = 0x36AD
  *c.R1.DE = 0x34CB
  *c.R1.HL = 0xF950
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1B33
  *c.R1.IY = 0xAA23
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF7
  mem[0x0003] = 0x71
  mem[0x1B2A] = 0x08
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAC5C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAC5C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x36AD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x36AD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x34CB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x34CB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF950) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF950, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1B33) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1B33, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAA23) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAA23, *c.R1.IY)
  }
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

func TestDDCB72(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB1B3
  *c.R1.BC = 0xF1E4
  *c.R1.DE = 0x9984
  *c.R1.HL = 0xC7FB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCE25
  *c.R1.IY = 0xC5B6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x22
  mem[0x0003] = 0x72
  mem[0xCE47] = 0x08
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB15D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB15D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF1E4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF1E4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9984) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9984, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC7FB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC7FB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCE25) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCE25, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC5B6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC5B6, *c.R1.IY)
  }
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

func TestDDCB73(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x21BA
  *c.R1.BC = 0x592D
  *c.R1.DE = 0xF406
  *c.R1.HL = 0xE21F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6442
  *c.R1.IY = 0xCF58
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x12
  mem[0x0003] = 0x73
  mem[0x6454] = 0x3C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2174) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2174, *c.R1.AF)
  }
  if (*c.R1.BC != 0x592D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x592D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF406) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF406, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE21F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE21F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6442) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6442, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCF58) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCF58, *c.R1.IY)
  }
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

func TestDDCB74(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6642
  *c.R1.BC = 0x64C1
  *c.R1.DE = 0xDBE5
  *c.R1.HL = 0xEB48
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7DC1
  *c.R1.IY = 0xC1FB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x08
  mem[0x0003] = 0x74
  mem[0x7DC9] = 0xBE
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x667C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x667C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x64C1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x64C1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDBE5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDBE5, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEB48) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEB48, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7DC1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7DC1, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC1FB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC1FB, *c.R1.IY)
  }
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

func TestDDCB75(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8778
  *c.R1.BC = 0x580E
  *c.R1.DE = 0x00DD
  *c.R1.HL = 0xF4C6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x60AD
  *c.R1.IY = 0x9B60
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5B
  mem[0x0003] = 0x75
  mem[0x6108] = 0xCF
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8730) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8730, *c.R1.AF)
  }
  if (*c.R1.BC != 0x580E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x580E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x00DD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x00DD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF4C6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF4C6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x60AD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x60AD, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9B60) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9B60, *c.R1.IY)
  }
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

func TestDDCB76(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x65B8
  *c.R1.BC = 0x5CC2
  *c.R1.DE = 0x3058
  *c.R1.HL = 0xE258
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7E8A
  *c.R1.IY = 0xB296
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x73
  mem[0x0003] = 0x76
  mem[0x7EFD] = 0x1E
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x657C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x657C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5CC2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5CC2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3058) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3058, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE258) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE258, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7E8A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7E8A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB296) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB296, *c.R1.IY)
  }
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

func TestDDCB77(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE3A8
  *c.R1.BC = 0x47A0
  *c.R1.DE = 0xC510
  *c.R1.HL = 0xCF0A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0537
  *c.R1.IY = 0xB242
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7F
  mem[0x0003] = 0x77
  mem[0x05B6] = 0x97
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE354) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE354, *c.R1.AF)
  }
  if (*c.R1.BC != 0x47A0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x47A0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC510) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC510, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCF0A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCF0A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0537) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0537, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB242) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB242, *c.R1.IY)
  }
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

func TestDDCB78(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x424F
  *c.R1.BC = 0x24F6
  *c.R1.DE = 0x1632
  *c.R1.HL = 0x8A4F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9397
  *c.R1.IY = 0x846C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x70
  mem[0x0003] = 0x78
  mem[0x9407] = 0x76
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4255) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4255, *c.R1.AF)
  }
  if (*c.R1.BC != 0x24F6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x24F6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1632) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1632, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8A4F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8A4F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9397) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9397, *c.R1.IX)
  }
  if (*c.R1.IY != 0x846C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x846C, *c.R1.IY)
  }
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

func TestDDCB79(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE6A0
  *c.R1.BC = 0xEEAA
  *c.R1.DE = 0x41F7
  *c.R1.HL = 0x5DA2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x41DE
  *c.R1.IY = 0x4189
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC3
  mem[0x0003] = 0x79
  mem[0x41A1] = 0xB8
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE690) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE690, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEEAA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEEAA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x41F7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x41F7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5DA2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5DA2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x41DE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x41DE, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4189) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4189, *c.R1.IY)
  }
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

func TestDDCB7A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCABF
  *c.R1.BC = 0x56AA
  *c.R1.DE = 0x6A06
  *c.R1.HL = 0x6CD7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0AA9
  *c.R1.IY = 0x9812
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3F
  mem[0x0003] = 0x7A
  mem[0x0AE8] = 0xEB
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCA99) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCA99, *c.R1.AF)
  }
  if (*c.R1.BC != 0x56AA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x56AA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6A06) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6A06, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6CD7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6CD7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0AA9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0AA9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9812) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9812, *c.R1.IY)
  }
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

func TestDDCB7B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAE3F
  *c.R1.BC = 0x0227
  *c.R1.DE = 0x721F
  *c.R1.HL = 0x52A1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5040
  *c.R1.IY = 0xB98A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x88
  mem[0x0003] = 0x7B
  mem[0x4FC8] = 0x22
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAE5D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAE5D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0227) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0227, *c.R1.BC)
  }
  if (*c.R1.DE != 0x721F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x721F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x52A1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x52A1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5040) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5040, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB98A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB98A, *c.R1.IY)
  }
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

func TestDDCB7C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8A80
  *c.R1.BC = 0xA2F1
  *c.R1.DE = 0x239A
  *c.R1.HL = 0xD5CC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6883
  *c.R1.IY = 0xB050
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9E
  mem[0x0003] = 0x7C
  mem[0x6821] = 0x3A
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8A7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8A7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA2F1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA2F1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x239A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x239A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD5CC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD5CC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6883) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6883, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB050) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB050, *c.R1.IY)
  }
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

func TestDDCB7D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC37F
  *c.R1.BC = 0xCF33
  *c.R1.DE = 0x1010
  *c.R1.HL = 0x98E6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB021
  *c.R1.IY = 0x0356
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x29
  mem[0x0003] = 0x7D
  mem[0xB04A] = 0x2C
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC375) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC375, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF33) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF33, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1010) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1010, *c.R1.DE)
  }
  if (*c.R1.HL != 0x98E6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x98E6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB021) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB021, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0356) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0356, *c.R1.IY)
  }
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

func TestDDCB7E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9A25
  *c.R1.BC = 0x2F6E
  *c.R1.DE = 0x0D0D
  *c.R1.HL = 0xA83F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCEF0
  *c.R1.IY = 0x8C15
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4F
  mem[0x0003] = 0x7E
  mem[0xCF3F] = 0xF2
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9A99) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9A99, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2F6E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2F6E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0D0D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0D0D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA83F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA83F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCEF0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCEF0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8C15) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8C15, *c.R1.IY)
  }
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

func TestDDCB7F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x53B9
  *c.R1.BC = 0x1F4E
  *c.R1.DE = 0x4837
  *c.R1.HL = 0x21B6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5EC2
  *c.R1.IY = 0x80C3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x75
  mem[0x0003] = 0x7F
  mem[0x5F37] = 0xA2
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5399) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5399, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1F4E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1F4E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4837) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4837, *c.R1.DE)
  }
  if (*c.R1.HL != 0x21B6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x21B6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5EC2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5EC2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x80C3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x80C3, *c.R1.IY)
  }
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

func TestDDCB80(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6319
  *c.R1.BC = 0xBAF9
  *c.R1.DE = 0xC84B
  *c.R1.HL = 0xBCF2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xACC5
  *c.R1.IY = 0xA4ED
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x70
  mem[0x0003] = 0x80
  mem[0xAD35] = 0x30
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6319) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6319, *c.R1.AF)
  }
  if (*c.R1.BC != 0x30F9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x30F9, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC84B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC84B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBCF2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBCF2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xACC5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xACC5, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA4ED) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA4ED, *c.R1.IY)
  }
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

func TestDDCB81(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFAE1
  *c.R1.BC = 0x5AE5
  *c.R1.DE = 0x9502
  *c.R1.HL = 0xDC9B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBDD3
  *c.R1.IY = 0x1A52
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2A
  mem[0x0003] = 0x81
  mem[0xBDFD] = 0x24
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFAE1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFAE1, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5A24) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5A24, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9502) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9502, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDC9B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDC9B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBDD3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBDD3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1A52) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1A52, *c.R1.IY)
  }
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

func TestDDCB82(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDAF6
  *c.R1.BC = 0x3260
  *c.R1.DE = 0xF1AC
  *c.R1.HL = 0x1D47
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5E74
  *c.R1.IY = 0x35E2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9A
  mem[0x0003] = 0x82
  mem[0x5E0E] = 0x51
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDAF6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDAF6, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3260) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3260, *c.R1.BC)
  }
  if (*c.R1.DE != 0x50AC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x50AC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1D47) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1D47, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5E74) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5E74, *c.R1.IX)
  }
  if (*c.R1.IY != 0x35E2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x35E2, *c.R1.IY)
  }
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

func TestDDCB83(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8E7C
  *c.R1.BC = 0x5586
  *c.R1.DE = 0x8C92
  *c.R1.HL = 0xFB00
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3441
  *c.R1.IY = 0xD365
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0E
  mem[0x0003] = 0x83
  mem[0x344F] = 0x01
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8E7C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8E7C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5586) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5586, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8C00) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8C00, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFB00) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFB00, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3441) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3441, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD365) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD365, *c.R1.IY)
  }
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

func TestDDCB84(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC1B3
  *c.R1.BC = 0x4874
  *c.R1.DE = 0xC535
  *c.R1.HL = 0x0E1C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0123
  *c.R1.IY = 0xDD28
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x47
  mem[0x0003] = 0x84
  mem[0x016A] = 0xB0
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC1B3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC1B3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4874) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4874, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC535) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC535, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB01C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB01C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0123) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0123, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDD28) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDD28, *c.R1.IY)
  }
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

func TestDDCB85(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0928
  *c.R1.BC = 0xB0DB
  *c.R1.DE = 0x4E07
  *c.R1.HL = 0xA7B7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0BA3
  *c.R1.IY = 0xC61C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6C
  mem[0x0003] = 0x85
  mem[0x0C0F] = 0xDE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0928) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0928, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB0DB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB0DB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4E07) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4E07, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA7DE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA7DE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0BA3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0BA3, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC61C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC61C, *c.R1.IY)
  }
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

func TestDDCB86(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4515
  *c.R1.BC = 0xDE09
  *c.R1.DE = 0x3CE7
  *c.R1.HL = 0x1FDE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x10C5
  *c.R1.IY = 0x33ED
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x5C
  mem[0x0003] = 0x86
  mem[0x1121] = 0x7C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4515) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4515, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDE09) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDE09, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3CE7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3CE7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1FDE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1FDE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x10C5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x10C5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x33ED) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x33ED, *c.R1.IY)
  }
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

func TestDDCB87(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD05E
  *c.R1.BC = 0xA733
  *c.R1.DE = 0xD1DD
  *c.R1.HL = 0x1603
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEDE6
  *c.R1.IY = 0xE5FB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x02
  mem[0x0003] = 0x87
  mem[0xEDE8] = 0xC4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC45E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC45E, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA733) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA733, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD1DD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD1DD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1603) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1603, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEDE6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEDE6, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE5FB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE5FB, *c.R1.IY)
  }
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

func TestDDCB88(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE4FA
  *c.R1.BC = 0x3325
  *c.R1.DE = 0xC266
  *c.R1.HL = 0x1B13
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x878E
  *c.R1.IY = 0xE695
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9B
  mem[0x0003] = 0x88
  mem[0x8729] = 0x7C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE4FA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE4FA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7C25) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7C25, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC266) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC266, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1B13) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1B13, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x878E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x878E, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE695) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE695, *c.R1.IY)
  }
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

func TestDDCB89(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x933B
  *c.R1.BC = 0x6FDD
  *c.R1.DE = 0xA3A8
  *c.R1.HL = 0x2634
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8F3E
  *c.R1.IY = 0x7727
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2B
  mem[0x0003] = 0x89
  mem[0x8F69] = 0xCF
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x933B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x933B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6FCD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6FCD, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA3A8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA3A8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2634) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2634, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8F3E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8F3E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7727) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7727, *c.R1.IY)
  }
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

func TestDDCB8A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6759
  *c.R1.BC = 0xAD1E
  *c.R1.DE = 0x5D71
  *c.R1.HL = 0xCE52
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x39A9
  *c.R1.IY = 0x38A0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0A
  mem[0x0003] = 0x8A
  mem[0x39B3] = 0xEA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6759) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6759, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAD1E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAD1E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE871) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE871, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCE52) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCE52, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x39A9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x39A9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x38A0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x38A0, *c.R1.IY)
  }
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

func TestDDCB8B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3DA2
  *c.R1.BC = 0x1833
  *c.R1.DE = 0x03C1
  *c.R1.HL = 0x07E9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1685
  *c.R1.IY = 0xD790
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x62
  mem[0x0003] = 0x8B
  mem[0x16E7] = 0x8A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3DA2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3DA2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1833) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1833, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0388) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0388, *c.R1.DE)
  }
  if (*c.R1.HL != 0x07E9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x07E9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1685) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1685, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD790) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD790, *c.R1.IY)
  }
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

func TestDDCB8C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA625
  *c.R1.BC = 0xED31
  *c.R1.DE = 0x3946
  *c.R1.HL = 0x32DC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC6A2
  *c.R1.IY = 0x7AD6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE8
  mem[0x0003] = 0x8C
  mem[0xC68A] = 0x3E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA625) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA625, *c.R1.AF)
  }
  if (*c.R1.BC != 0xED31) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xED31, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3946) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3946, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3CDC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3CDC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC6A2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC6A2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7AD6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7AD6, *c.R1.IY)
  }
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

func TestDDCB8D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x016B
  *c.R1.BC = 0x5802
  *c.R1.DE = 0xA683
  *c.R1.HL = 0x2549
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x22E6
  *c.R1.IY = 0x33BB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xCC
  mem[0x0003] = 0x8D
  mem[0x22B2] = 0x9E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x016B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x016B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5802) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5802, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA683) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA683, *c.R1.DE)
  }
  if (*c.R1.HL != 0x259C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x259C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x22E6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x22E6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x33BB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x33BB, *c.R1.IY)
  }
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

func TestDDCB8E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF4F4
  *c.R1.BC = 0xF3A8
  *c.R1.DE = 0x2843
  *c.R1.HL = 0x82CB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD2E8
  *c.R1.IY = 0xD367
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0A
  mem[0x0003] = 0x8E
  mem[0xD2F2] = 0x03
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF4F4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF4F4, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF3A8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF3A8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2843) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2843, *c.R1.DE)
  }
  if (*c.R1.HL != 0x82CB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x82CB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD2E8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD2E8, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD367) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD367, *c.R1.IY)
  }
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

func TestDDCB8F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6B1A
  *c.R1.BC = 0x8AE2
  *c.R1.DE = 0x269B
  *c.R1.HL = 0xCB2F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3FFE
  *c.R1.IY = 0x75DD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7B
  mem[0x0003] = 0x8F
  mem[0x4079] = 0x96
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x941A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x941A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8AE2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8AE2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x269B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x269B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCB2F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCB2F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3FFE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3FFE, *c.R1.IX)
  }
  if (*c.R1.IY != 0x75DD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x75DD, *c.R1.IY)
  }
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

func TestDDCB90(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC167
  *c.R1.BC = 0x3DFC
  *c.R1.DE = 0x42E7
  *c.R1.HL = 0x9E14
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB501
  *c.R1.IY = 0x84FE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x04
  mem[0x0003] = 0x90
  mem[0xB505] = 0x46
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC167) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC167, *c.R1.AF)
  }
  if (*c.R1.BC != 0x42FC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x42FC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x42E7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x42E7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9E14) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9E14, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB501) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB501, *c.R1.IX)
  }
  if (*c.R1.IY != 0x84FE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x84FE, *c.R1.IY)
  }
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

func TestDDCB91(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE85E
  *c.R1.BC = 0xCC89
  *c.R1.DE = 0xD249
  *c.R1.HL = 0xEA3B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC987
  *c.R1.IY = 0xC4D1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x11
  mem[0x0003] = 0x91
  mem[0xC998] = 0x83
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE85E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE85E, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCC83) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCC83, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD249) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD249, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEA3B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEA3B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC987) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC987, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC4D1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC4D1, *c.R1.IY)
  }
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

func TestDDCB92(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x28A3
  *c.R1.BC = 0x85FF
  *c.R1.DE = 0xAB28
  *c.R1.HL = 0x47A5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9166
  *c.R1.IY = 0xE755
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4B
  mem[0x0003] = 0x92
  mem[0x91B1] = 0xAA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x28A3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x28A3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x85FF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x85FF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAA28) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAA28, *c.R1.DE)
  }
  if (*c.R1.HL != 0x47A5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x47A5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9166) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9166, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE755) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE755, *c.R1.IY)
  }
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

func TestDDCB93(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x58AC
  *c.R1.BC = 0xC88B
  *c.R1.DE = 0x6D24
  *c.R1.HL = 0xDBDD
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAC2E
  *c.R1.IY = 0x5199
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x03
  mem[0x0003] = 0x93
  mem[0xAC31] = 0x93
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x58AC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x58AC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC88B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC88B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6D93) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6D93, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDBDD) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDBDD, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAC2E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAC2E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5199) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5199, *c.R1.IY)
  }
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

func TestDDCB94(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE38D
  *c.R1.BC = 0x35A5
  *c.R1.DE = 0x8D07
  *c.R1.HL = 0xBFB8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5E84
  *c.R1.IY = 0x5F24
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x11
  mem[0x0003] = 0x94
  mem[0x5E95] = 0xB7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE38D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE38D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x35A5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x35A5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8D07) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8D07, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB3B8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB3B8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5E84) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5E84, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5F24) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5F24, *c.R1.IY)
  }
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

func TestDDCB95(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x41F4
  *c.R1.BC = 0x9536
  *c.R1.DE = 0xDD7D
  *c.R1.HL = 0x4948
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFB74
  *c.R1.IY = 0xF17D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE6
  mem[0x0003] = 0x95
  mem[0xFB5A] = 0xC6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x41F4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x41F4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9536) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9536, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDD7D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDD7D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x49C2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x49C2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFB74) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFB74, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF17D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF17D, *c.R1.IY)
  }
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

func TestDDCB96(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4A9E
  *c.R1.BC = 0x42EF
  *c.R1.DE = 0x32D7
  *c.R1.HL = 0x18CF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7A81
  *c.R1.IY = 0xBB1D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD5
  mem[0x0003] = 0x96
  mem[0x7A56] = 0xAE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4A9E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4A9E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x42EF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x42EF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x32D7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x32D7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x18CF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x18CF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7A81) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7A81, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBB1D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBB1D, *c.R1.IY)
  }
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

func TestDDCB97(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9AD3
  *c.R1.BC = 0x89F0
  *c.R1.DE = 0x73C7
  *c.R1.HL = 0x0B1A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x847C
  *c.R1.IY = 0x4B86
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x92
  mem[0x0003] = 0x97
  mem[0x840E] = 0x23
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x23D3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x23D3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x89F0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x89F0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x73C7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x73C7, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0B1A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0B1A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x847C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x847C, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4B86) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4B86, *c.R1.IY)
  }
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

func TestDDCB98(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6E22
  *c.R1.BC = 0xB9FD
  *c.R1.DE = 0x9FDC
  *c.R1.HL = 0x3AED
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x041E
  *c.R1.IY = 0xFD79
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDC
  mem[0x0003] = 0x98
  mem[0x03FA] = 0x58
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6E22) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6E22, *c.R1.AF)
  }
  if (*c.R1.BC != 0x50FD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x50FD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9FDC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9FDC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3AED) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3AED, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x041E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x041E, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFD79) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFD79, *c.R1.IY)
  }
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

func TestDDCB99(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA132
  *c.R1.BC = 0x3891
  *c.R1.DE = 0x1515
  *c.R1.HL = 0x2830
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x09FD
  *c.R1.IY = 0x0473
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6D
  mem[0x0003] = 0x99
  mem[0x0A6A] = 0xCE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA132) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA132, *c.R1.AF)
  }
  if (*c.R1.BC != 0x38C6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x38C6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1515) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1515, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2830) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2830, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x09FD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x09FD, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0473) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0473, *c.R1.IY)
  }
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

func TestDDCB9A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x783D
  *c.R1.BC = 0x8F69
  *c.R1.DE = 0x91C4
  *c.R1.HL = 0xE38F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x68A8
  *c.R1.IY = 0x391D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8A
  mem[0x0003] = 0x9A
  mem[0x6832] = 0xA8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x783D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x783D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8F69) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8F69, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA0C4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA0C4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE38F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE38F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x68A8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x68A8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x391D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x391D, *c.R1.IY)
  }
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

func TestDDCB9B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x955A
  *c.R1.BC = 0xC7B0
  *c.R1.DE = 0x53B3
  *c.R1.HL = 0xAEC6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x06EF
  *c.R1.IY = 0xE991
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x97
  mem[0x0003] = 0x9B
  mem[0x0686] = 0x62
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x955A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x955A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC7B0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC7B0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5362) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5362, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAEC6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAEC6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x06EF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x06EF, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE991) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE991, *c.R1.IY)
  }
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

func TestDDCB9C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAF69
  *c.R1.BC = 0xF896
  *c.R1.DE = 0xE791
  *c.R1.HL = 0xA2EE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x847B
  *c.R1.IY = 0x59ED
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x54
  mem[0x0003] = 0x9C
  mem[0x84CF] = 0x1B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAF69) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAF69, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF896) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF896, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE791) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE791, *c.R1.DE)
  }
  if (*c.R1.HL != 0x13EE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x13EE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x847B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x847B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x59ED) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x59ED, *c.R1.IY)
  }
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

func TestDDCB9D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7D1E
  *c.R1.BC = 0x5009
  *c.R1.DE = 0x1248
  *c.R1.HL = 0x380C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE920
  *c.R1.IY = 0x4FE6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0F
  mem[0x0003] = 0x9D
  mem[0xE92F] = 0xE8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7D1E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7D1E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5009) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5009, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1248) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1248, *c.R1.DE)
  }
  if (*c.R1.HL != 0x38E0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x38E0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE920) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE920, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4FE6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4FE6, *c.R1.IY)
  }
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

func TestDDCB9E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC207
  *c.R1.BC = 0xB47C
  *c.R1.DE = 0x0E16
  *c.R1.HL = 0xE17F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD8BB
  *c.R1.IY = 0xBB99
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB5
  mem[0x0003] = 0x9E
  mem[0xD870] = 0xEE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC207) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC207, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB47C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB47C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0E16) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0E16, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE17F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE17F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD8BB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD8BB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBB99) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBB99, *c.R1.IY)
  }
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

func TestDDCB9F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC26B
  *c.R1.BC = 0x7537
  *c.R1.DE = 0x46BB
  *c.R1.HL = 0x13C0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE63C
  *c.R1.IY = 0x1D98
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB8
  mem[0x0003] = 0x9F
  mem[0xE5F4] = 0xA6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA66B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA66B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7537) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7537, *c.R1.BC)
  }
  if (*c.R1.DE != 0x46BB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x46BB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x13C0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x13C0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE63C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE63C, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1D98) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1D98, *c.R1.IY)
  }
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

func TestDDCBA0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0BBE
  *c.R1.BC = 0x8500
  *c.R1.DE = 0x8609
  *c.R1.HL = 0x5352
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA2F0
  *c.R1.IY = 0xDA02
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x02
  mem[0x0003] = 0xA0
  mem[0xA2F2] = 0x39
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0BBE) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0BBE, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2900) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2900, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8609) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8609, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5352) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5352, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA2F0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA2F0, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDA02) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDA02, *c.R1.IY)
  }
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

func TestDDCBA1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAD0A
  *c.R1.BC = 0xAA76
  *c.R1.DE = 0x0F2D
  *c.R1.HL = 0x832C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x45BB
  *c.R1.IY = 0xA22D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF5
  mem[0x0003] = 0xA1
  mem[0x45B0] = 0xD2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAD0A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAD0A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAAC2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAAC2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0F2D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0F2D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x832C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x832C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x45BB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x45BB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA22D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA22D, *c.R1.IY)
  }
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

func TestDDCBA2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF586
  *c.R1.BC = 0x4A7D
  *c.R1.DE = 0xA5AB
  *c.R1.HL = 0x26FC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x628B
  *c.R1.IY = 0x6C4D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0E
  mem[0x0003] = 0xA2
  mem[0x6299] = 0xA1
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF586) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF586, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4A7D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4A7D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA1AB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA1AB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x26FC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x26FC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x628B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x628B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6C4D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6C4D, *c.R1.IY)
  }
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

func TestDDCBA3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDE5B
  *c.R1.BC = 0xA284
  *c.R1.DE = 0xD40E
  *c.R1.HL = 0xC92D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x040D
  *c.R1.IY = 0x12C0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2E
  mem[0x0003] = 0xA3
  mem[0x043B] = 0x04
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDE5B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDE5B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA284) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA284, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD404) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD404, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC92D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC92D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x040D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x040D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x12C0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x12C0, *c.R1.IY)
  }
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

func TestDDCBA4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDFAA
  *c.R1.BC = 0xAE40
  *c.R1.DE = 0x02C3
  *c.R1.HL = 0xE0B5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFE4D
  *c.R1.IY = 0xFAA3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x03
  mem[0x0003] = 0xA4
  mem[0xFE50] = 0x27
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDFAA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDFAA, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAE40) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAE40, *c.R1.BC)
  }
  if (*c.R1.DE != 0x02C3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x02C3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x27B5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x27B5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFE4D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFE4D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFAA3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFAA3, *c.R1.IY)
  }
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

func TestDDCBA5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1A15
  *c.R1.BC = 0x04CB
  *c.R1.DE = 0x4352
  *c.R1.HL = 0xEE39
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7B27
  *c.R1.IY = 0x38A0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF6
  mem[0x0003] = 0xA5
  mem[0x7B1D] = 0x6B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1A15) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1A15, *c.R1.AF)
  }
  if (*c.R1.BC != 0x04CB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x04CB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4352) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4352, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEE6B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEE6B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7B27) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7B27, *c.R1.IX)
  }
  if (*c.R1.IY != 0x38A0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x38A0, *c.R1.IY)
  }
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

func TestDDCBA6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5E46
  *c.R1.BC = 0xB98A
  *c.R1.DE = 0xB822
  *c.R1.HL = 0x04CA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAE1B
  *c.R1.IY = 0x8730
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x27
  mem[0x0003] = 0xA6
  mem[0xAE42] = 0x8F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5E46) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5E46, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB98A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB98A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB822) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB822, *c.R1.DE)
  }
  if (*c.R1.HL != 0x04CA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x04CA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAE1B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAE1B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8730) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8730, *c.R1.IY)
  }
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

func TestDDCBA7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0EED
  *c.R1.BC = 0x7B11
  *c.R1.DE = 0x8CB0
  *c.R1.HL = 0xEB3D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5EC8
  *c.R1.IY = 0x97CF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF2
  mem[0x0003] = 0xA7
  mem[0x5EBA] = 0x87
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x87ED) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x87ED, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7B11) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7B11, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8CB0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8CB0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEB3D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEB3D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5EC8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5EC8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x97CF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x97CF, *c.R1.IY)
  }
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

func TestDDCBA8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5173
  *c.R1.BC = 0x3089
  *c.R1.DE = 0x070D
  *c.R1.HL = 0xE8F9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE84F
  *c.R1.IY = 0x55F0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD0
  mem[0x0003] = 0xA8
  mem[0xE81F] = 0x7E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5173) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5173, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5E89) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5E89, *c.R1.BC)
  }
  if (*c.R1.DE != 0x070D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x070D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE8F9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE8F9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE84F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE84F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x55F0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x55F0, *c.R1.IY)
  }
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

func TestDDCBA9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4FB8
  *c.R1.BC = 0xCCB5
  *c.R1.DE = 0x3E9A
  *c.R1.HL = 0x2673
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0FDD
  *c.R1.IY = 0xAEF2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9D
  mem[0x0003] = 0xA9
  mem[0x0F7A] = 0x1F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4FB8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4FB8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCC1F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCC1F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3E9A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3E9A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2673) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2673, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0FDD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0FDD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAEF2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAEF2, *c.R1.IY)
  }
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

func TestDDCBAA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFE76
  *c.R1.BC = 0x6F96
  *c.R1.DE = 0x3FEB
  *c.R1.HL = 0x0B21
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6747
  *c.R1.IY = 0x07BA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9F
  mem[0x0003] = 0xAA
  mem[0x66E6] = 0x50
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFE76) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFE76, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6F96) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6F96, *c.R1.BC)
  }
  if (*c.R1.DE != 0x50EB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x50EB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0B21) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0B21, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6747) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6747, *c.R1.IX)
  }
  if (*c.R1.IY != 0x07BA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x07BA, *c.R1.IY)
  }
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

func TestDDCBAB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2EB4
  *c.R1.BC = 0x36F1
  *c.R1.DE = 0x8F44
  *c.R1.HL = 0x36AF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6682
  *c.R1.IY = 0x9D60
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x16
  mem[0x0003] = 0xAB
  mem[0x6698] = 0xEB
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2EB4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2EB4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x36F1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x36F1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8FCB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8FCB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x36AF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x36AF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6682) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6682, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9D60) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9D60, *c.R1.IY)
  }
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

func TestDDCBAC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAF32
  *c.R1.BC = 0x8CA8
  *c.R1.DE = 0x6558
  *c.R1.HL = 0x06D9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA4DD
  *c.R1.IY = 0xCD1F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC4
  mem[0x0003] = 0xAC
  mem[0xA4A1] = 0x44
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAF32) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAF32, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8CA8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8CA8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6558) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6558, *c.R1.DE)
  }
  if (*c.R1.HL != 0x44D9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x44D9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA4DD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA4DD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCD1F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCD1F, *c.R1.IY)
  }
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

func TestDDCBAD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFCC9
  *c.R1.BC = 0x69A7
  *c.R1.DE = 0x0EED
  *c.R1.HL = 0xEAB5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEEF5
  *c.R1.IY = 0x3ED2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x49
  mem[0x0003] = 0xAD
  mem[0xEF3E] = 0x76
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFCC9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFCC9, *c.R1.AF)
  }
  if (*c.R1.BC != 0x69A7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x69A7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0EED) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0EED, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEA56) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEA56, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEEF5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEEF5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3ED2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3ED2, *c.R1.IY)
  }
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

func TestDDCBAE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5F7A
  *c.R1.BC = 0x9C20
  *c.R1.DE = 0xF013
  *c.R1.HL = 0xC4B7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB306
  *c.R1.IY = 0x15DD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6E
  mem[0x0003] = 0xAE
  mem[0xB374] = 0x5A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5F7A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5F7A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9C20) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9C20, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF013) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF013, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC4B7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC4B7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB306) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB306, *c.R1.IX)
  }
  if (*c.R1.IY != 0x15DD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x15DD, *c.R1.IY)
  }
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

func TestDDCBAF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB11E
  *c.R1.BC = 0x2583
  *c.R1.DE = 0x51FA
  *c.R1.HL = 0xD427
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3619
  *c.R1.IY = 0x9CEF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC2
  mem[0x0003] = 0xAF
  mem[0x35DB] = 0x15
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x151E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x151E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2583) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2583, *c.R1.BC)
  }
  if (*c.R1.DE != 0x51FA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x51FA, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD427) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD427, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3619) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3619, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9CEF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9CEF, *c.R1.IY)
  }
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

func TestDDCBB0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF43E
  *c.R1.BC = 0xCE57
  *c.R1.DE = 0x3BF3
  *c.R1.HL = 0x0933
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x58D7
  *c.R1.IY = 0xD89F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x47
  mem[0x0003] = 0xB0
  mem[0x591E] = 0x1E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF43E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF43E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1E57) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1E57, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3BF3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3BF3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0933) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0933, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x58D7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x58D7, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD89F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD89F, *c.R1.IY)
  }
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

func TestDDCBB1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x35EF
  *c.R1.BC = 0xBBBC
  *c.R1.DE = 0xDB46
  *c.R1.HL = 0x046C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xADD2
  *c.R1.IY = 0x2B6E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x86
  mem[0x0003] = 0xB1
  mem[0xAD58] = 0x46
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x35EF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x35EF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBB06) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBB06, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDB46) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDB46, *c.R1.DE)
  }
  if (*c.R1.HL != 0x046C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x046C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xADD2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xADD2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2B6E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2B6E, *c.R1.IY)
  }
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

func TestDDCBB2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC26C
  *c.R1.BC = 0xFD32
  *c.R1.DE = 0x9B7F
  *c.R1.HL = 0xAB6C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE7D0
  *c.R1.IY = 0x501F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x70
  mem[0x0003] = 0xB2
  mem[0xE840] = 0x48
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC26C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC26C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFD32) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFD32, *c.R1.BC)
  }
  if (*c.R1.DE != 0x087F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x087F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAB6C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAB6C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE7D0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE7D0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x501F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x501F, *c.R1.IY)
  }
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

func TestDDCBB3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x36CA
  *c.R1.BC = 0xB434
  *c.R1.DE = 0xE212
  *c.R1.HL = 0xF805
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x53FB
  *c.R1.IY = 0xB191
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDE
  mem[0x0003] = 0xB3
  mem[0x53D9] = 0x06
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x36CA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x36CA, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB434) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB434, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE206) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE206, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF805) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF805, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x53FB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x53FB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB191) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB191, *c.R1.IY)
  }
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

func TestDDCBB4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0A1C
  *c.R1.BC = 0xAB67
  *c.R1.DE = 0x9CA1
  *c.R1.HL = 0x2F98
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5066
  *c.R1.IY = 0x320C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6B
  mem[0x0003] = 0xB4
  mem[0x50D1] = 0xDD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0A1C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0A1C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAB67) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAB67, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9CA1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9CA1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9D98) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9D98, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5066) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5066, *c.R1.IX)
  }
  if (*c.R1.IY != 0x320C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x320C, *c.R1.IY)
  }
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

func TestDDCBB5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFD6D
  *c.R1.BC = 0x51C9
  *c.R1.DE = 0x16D6
  *c.R1.HL = 0x1373
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x146E
  *c.R1.IY = 0x2148
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEC
  mem[0x0003] = 0xB5
  mem[0x145A] = 0xD6
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFD6D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFD6D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x51C9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x51C9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x16D6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x16D6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1396) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1396, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x146E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x146E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2148) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2148, *c.R1.IY)
  }
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

func TestDDCBB6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1D0B
  *c.R1.BC = 0x04E8
  *c.R1.DE = 0x109E
  *c.R1.HL = 0x1DDE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8772
  *c.R1.IY = 0x8661
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x15
  mem[0x0003] = 0xB6
  mem[0x8787] = 0x8C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1D0B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1D0B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x04E8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x04E8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x109E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x109E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1DDE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1DDE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8772) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8772, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8661) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8661, *c.R1.IY)
  }
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

func TestDDCBB7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF012
  *c.R1.BC = 0xB87E
  *c.R1.DE = 0x65BA
  *c.R1.HL = 0xA5C8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6120
  *c.R1.IY = 0x789D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD3
  mem[0x0003] = 0xB7
  mem[0x60F3] = 0x54
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1412) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1412, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB87E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB87E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x65BA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x65BA, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA5C8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA5C8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6120) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6120, *c.R1.IX)
  }
  if (*c.R1.IY != 0x789D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x789D, *c.R1.IY)
  }
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

func TestDDCBB8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8EAE
  *c.R1.BC = 0x4A53
  *c.R1.DE = 0xBFA1
  *c.R1.HL = 0x5E7E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0BF6
  *c.R1.IY = 0x1E35
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x13
  mem[0x0003] = 0xB8
  mem[0x0C09] = 0x87
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8EAE) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8EAE, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0753) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0753, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBFA1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBFA1, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E7E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E7E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0BF6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0BF6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1E35) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1E35, *c.R1.IY)
  }
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

func TestDDCBB9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5FB7
  *c.R1.BC = 0xA81E
  *c.R1.DE = 0xE2D2
  *c.R1.HL = 0x4117
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0564
  *c.R1.IY = 0x48A1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x56
  mem[0x0003] = 0xB9
  mem[0x05BA] = 0xC8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5FB7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5FB7, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA848) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA848, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE2D2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE2D2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4117) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4117, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0564) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0564, *c.R1.IX)
  }
  if (*c.R1.IY != 0x48A1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x48A1, *c.R1.IY)
  }
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

func TestDDCBBA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7F6A
  *c.R1.BC = 0x47FE
  *c.R1.DE = 0xCE45
  *c.R1.HL = 0x75DE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF5E0
  *c.R1.IY = 0x032C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x86
  mem[0x0003] = 0xBA
  mem[0xF566] = 0x30
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7F6A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7F6A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x47FE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x47FE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3045) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3045, *c.R1.DE)
  }
  if (*c.R1.HL != 0x75DE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x75DE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF5E0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF5E0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x032C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x032C, *c.R1.IY)
  }
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

func TestDDCBBB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC7E3
  *c.R1.BC = 0xE49E
  *c.R1.DE = 0x9EC5
  *c.R1.HL = 0x07E7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBD31
  *c.R1.IY = 0x9D5F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEF
  mem[0x0003] = 0xBB
  mem[0xBD20] = 0xC9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC7E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC7E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE49E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE49E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9E49) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9E49, *c.R1.DE)
  }
  if (*c.R1.HL != 0x07E7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x07E7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBD31) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBD31, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9D5F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9D5F, *c.R1.IY)
  }
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

func TestDDCBBC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB430
  *c.R1.BC = 0x7AC7
  *c.R1.DE = 0xB45F
  *c.R1.HL = 0xFBF7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x638E
  *c.R1.IY = 0x3173
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC0
  mem[0x0003] = 0xBC
  mem[0x634E] = 0x28
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB430) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB430, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7AC7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7AC7, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB45F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB45F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x28F7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x28F7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x638E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x638E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3173) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3173, *c.R1.IY)
  }
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

func TestDDCBBD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4E71
  *c.R1.BC = 0x6FFA
  *c.R1.DE = 0xA3F9
  *c.R1.HL = 0xA2E5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE3C4
  *c.R1.IY = 0x02D4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xB9
  mem[0x0003] = 0xBD
  mem[0xE37D] = 0xDD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4E71) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4E71, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6FFA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6FFA, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA3F9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA3F9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA25D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA25D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE3C4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE3C4, *c.R1.IX)
  }
  if (*c.R1.IY != 0x02D4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x02D4, *c.R1.IY)
  }
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

func TestDDCBBE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4AF8
  *c.R1.BC = 0x99A5
  *c.R1.DE = 0xD6FD
  *c.R1.HL = 0x7A16
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x58D3
  *c.R1.IY = 0xCE54
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4D
  mem[0x0003] = 0xBE
  mem[0x5920] = 0xE8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4AF8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4AF8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x99A5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x99A5, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD6FD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD6FD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7A16) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7A16, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x58D3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x58D3, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCE54) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCE54, *c.R1.IY)
  }
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

func TestDDCBBF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6E31
  *c.R1.BC = 0x0320
  *c.R1.DE = 0x134B
  *c.R1.HL = 0x77C3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1734
  *c.R1.IY = 0xBC2D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x26
  mem[0x0003] = 0xBF
  mem[0x175A] = 0xE2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6231) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6231, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0320) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0320, *c.R1.BC)
  }
  if (*c.R1.DE != 0x134B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x134B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x77C3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x77C3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1734) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1734, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBC2D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBC2D, *c.R1.IY)
  }
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

func TestDDCBC0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x75BE
  *c.R1.BC = 0x2B93
  *c.R1.DE = 0x093D
  *c.R1.HL = 0x1128
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x792E
  *c.R1.IY = 0x31F7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x00
  mem[0x0003] = 0xC0
  mem[0x792E] = 0x92
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x75BE) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x75BE, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9393) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9393, *c.R1.BC)
  }
  if (*c.R1.DE != 0x093D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x093D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1128) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1128, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x792E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x792E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x31F7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x31F7, *c.R1.IY)
  }
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

func TestDDCBC1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x313F
  *c.R1.BC = 0x8223
  *c.R1.DE = 0x5FCC
  *c.R1.HL = 0x42C8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDCCC
  *c.R1.IY = 0xD87B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF7
  mem[0x0003] = 0xC1
  mem[0xDCC3] = 0x1C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x313F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x313F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x821D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x821D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5FCC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5FCC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x42C8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x42C8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDCCC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDCCC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD87B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD87B, *c.R1.IY)
  }
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

func TestDDCBC2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA7E3
  *c.R1.BC = 0xBF55
  *c.R1.DE = 0xD27B
  *c.R1.HL = 0x0A9D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0CFA
  *c.R1.IY = 0xEA4E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x85
  mem[0x0003] = 0xC2
  mem[0x0C7F] = 0x30
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA7E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA7E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBF55) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBF55, *c.R1.BC)
  }
  if (*c.R1.DE != 0x317B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x317B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0A9D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0A9D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0CFA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0CFA, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEA4E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEA4E, *c.R1.IY)
  }
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

func TestDDCBC3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE076
  *c.R1.BC = 0x2760
  *c.R1.DE = 0x1EEC
  *c.R1.HL = 0x9968
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5426
  *c.R1.IY = 0xA1A0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x32
  mem[0x0003] = 0xC3
  mem[0x5458] = 0xDD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE076) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE076, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2760) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2760, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1EDD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1EDD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9968) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9968, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5426) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5426, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA1A0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA1A0, *c.R1.IY)
  }
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

func TestDDCBC4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA679
  *c.R1.BC = 0xCC05
  *c.R1.DE = 0x3F4D
  *c.R1.HL = 0xC899
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7ACD
  *c.R1.IY = 0x48D7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAE
  mem[0x0003] = 0xC4
  mem[0x7A7B] = 0x27
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA679) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA679, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCC05) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCC05, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3F4D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3F4D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2799) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2799, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7ACD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7ACD, *c.R1.IX)
  }
  if (*c.R1.IY != 0x48D7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x48D7, *c.R1.IY)
  }
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

func TestDDCBC5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDDFD
  *c.R1.BC = 0x64D4
  *c.R1.DE = 0x2671
  *c.R1.HL = 0x35E7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBA99
  *c.R1.IY = 0xBD98
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9C
  mem[0x0003] = 0xC5
  mem[0xBA35] = 0x20
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDDFD) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDDFD, *c.R1.AF)
  }
  if (*c.R1.BC != 0x64D4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x64D4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2671) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2671, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3521) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3521, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBA99) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBA99, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBD98) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBD98, *c.R1.IY)
  }
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

func TestDDCBC6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB324
  *c.R1.BC = 0xDC0C
  *c.R1.DE = 0x1E35
  *c.R1.HL = 0x8CD5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAB2C
  *c.R1.IY = 0xB6F3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC4
  mem[0x0003] = 0xC6
  mem[0xAAF0] = 0xB8
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB324) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB324, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDC0C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDC0C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E35) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E35, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8CD5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8CD5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAB2C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAB2C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB6F3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB6F3, *c.R1.IY)
  }
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

func TestDDCBC7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA254
  *c.R1.BC = 0x9E56
  *c.R1.DE = 0x6828
  *c.R1.HL = 0x3189
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x64CB
  *c.R1.IY = 0xDFAD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF8
  mem[0x0003] = 0xC7
  mem[0x64C3] = 0x94
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9554) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9554, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9E56) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9E56, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6828) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6828, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3189) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3189, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x64CB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x64CB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDFAD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDFAD, *c.R1.IY)
  }
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

func TestDDCBC8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8ACA
  *c.R1.BC = 0x139E
  *c.R1.DE = 0xE652
  *c.R1.HL = 0x248B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6E7A
  *c.R1.IY = 0x189A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x65
  mem[0x0003] = 0xC8
  mem[0x6EDF] = 0x8F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8ACA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8ACA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8F9E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8F9E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE652) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE652, *c.R1.DE)
  }
  if (*c.R1.HL != 0x248B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x248B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6E7A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6E7A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x189A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x189A, *c.R1.IY)
  }
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

func TestDDCBC9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF15F
  *c.R1.BC = 0x856E
  *c.R1.DE = 0xA21F
  *c.R1.HL = 0x8A59
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB670
  *c.R1.IY = 0x4F79
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xFB
  mem[0x0003] = 0xC9
  mem[0xB66B] = 0xB9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF15F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF15F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x85BB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x85BB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA21F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA21F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8A59) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8A59, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB670) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB670, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4F79) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4F79, *c.R1.IY)
  }
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

func TestDDCBCA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDFAB
  *c.R1.BC = 0xA031
  *c.R1.DE = 0x1D78
  *c.R1.HL = 0xAD3A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA887
  *c.R1.IY = 0x7334
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8A
  mem[0x0003] = 0xCA
  mem[0xA811] = 0x7E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDFAB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDFAB, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA031) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA031, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7E78) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7E78, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAD3A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAD3A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA887) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA887, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7334) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7334, *c.R1.IY)
  }
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

func TestDDCBCB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEBD6
  *c.R1.BC = 0x376E
  *c.R1.DE = 0xC346
  *c.R1.HL = 0xB10C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA447
  *c.R1.IY = 0x31D6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA4
  mem[0x0003] = 0xCB
  mem[0xA3EB] = 0x73
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEBD6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEBD6, *c.R1.AF)
  }
  if (*c.R1.BC != 0x376E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x376E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC373) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC373, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB10C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB10C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA447) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA447, *c.R1.IX)
  }
  if (*c.R1.IY != 0x31D6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x31D6, *c.R1.IY)
  }
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

func TestDDCBCC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0212
  *c.R1.BC = 0xDC46
  *c.R1.DE = 0x8F41
  *c.R1.HL = 0x854E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1F5A
  *c.R1.IY = 0x07CA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x65
  mem[0x0003] = 0xCC
  mem[0x1FBF] = 0x72
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0212) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0212, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDC46) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDC46, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8F41) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8F41, *c.R1.DE)
  }
  if (*c.R1.HL != 0x724E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x724E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1F5A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1F5A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x07CA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x07CA, *c.R1.IY)
  }
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

func TestDDCBCD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3344
  *c.R1.BC = 0xD73C
  *c.R1.DE = 0xD6B8
  *c.R1.HL = 0x929D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5376
  *c.R1.IY = 0x6D3A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE9
  mem[0x0003] = 0xCD
  mem[0x535F] = 0x1C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3344) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3344, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD73C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD73C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD6B8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD6B8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x921E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x921E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5376) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5376, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6D3A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6D3A, *c.R1.IY)
  }
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

func TestDDCBCE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9E47
  *c.R1.BC = 0xFC93
  *c.R1.DE = 0x9FFC
  *c.R1.HL = 0xAACE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0313
  *c.R1.IY = 0x7F66
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x85
  mem[0x0003] = 0xCE
  mem[0x0298] = 0x10
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9E47) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9E47, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFC93) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFC93, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9FFC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9FFC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAACE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAACE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0313) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0313, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7F66) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7F66, *c.R1.IY)
  }
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

func TestDDCBCF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x53E8
  *c.R1.BC = 0xD379
  *c.R1.DE = 0x87D5
  *c.R1.HL = 0x10B0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC5D0
  *c.R1.IY = 0x4F7F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xE2
  mem[0x0003] = 0xCF
  mem[0xC5B2] = 0xB5
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB7E8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB7E8, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD379) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD379, *c.R1.BC)
  }
  if (*c.R1.DE != 0x87D5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x87D5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x10B0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x10B0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC5D0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC5D0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4F7F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4F7F, *c.R1.IY)
  }
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

func TestDDCBD0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3278
  *c.R1.BC = 0x6114
  *c.R1.DE = 0xD25D
  *c.R1.HL = 0x1CF8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAD43
  *c.R1.IY = 0x99FC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7F
  mem[0x0003] = 0xD0
  mem[0xADC2] = 0x51
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3278) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3278, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5514) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5514, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD25D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD25D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1CF8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1CF8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAD43) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAD43, *c.R1.IX)
  }
  if (*c.R1.IY != 0x99FC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x99FC, *c.R1.IY)
  }
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

func TestDDCBD1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC0B8
  *c.R1.BC = 0x371A
  *c.R1.DE = 0x6472
  *c.R1.HL = 0xD92D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x10B2
  *c.R1.IY = 0x3074
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA6
  mem[0x0003] = 0xD1
  mem[0x1058] = 0x2C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC0B8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC0B8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x372C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x372C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6472) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6472, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD92D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD92D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x10B2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x10B2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3074) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3074, *c.R1.IY)
  }
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

func TestDDCBD2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5BB6
  *c.R1.BC = 0xCAA8
  *c.R1.DE = 0xE0DB
  *c.R1.HL = 0xAF84
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB9A1
  *c.R1.IY = 0x7B5F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x9C
  mem[0x0003] = 0xD2
  mem[0xB93D] = 0x9C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5BB6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5BB6, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCAA8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCAA8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9CDB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9CDB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAF84) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAF84, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB9A1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB9A1, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7B5F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7B5F, *c.R1.IY)
  }
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

func TestDDCBD3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDB6A
  *c.R1.BC = 0x4FE2
  *c.R1.DE = 0x9E52
  *c.R1.HL = 0xA034
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDA36
  *c.R1.IY = 0x88A0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBD
  mem[0x0003] = 0xD3
  mem[0xD9F3] = 0x60
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDB6A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDB6A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4FE2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4FE2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9E64) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9E64, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA034) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA034, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDA36) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDA36, *c.R1.IX)
  }
  if (*c.R1.IY != 0x88A0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x88A0, *c.R1.IY)
  }
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

func TestDDCBD4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCC1C
  *c.R1.BC = 0xB884
  *c.R1.DE = 0x6AD2
  *c.R1.HL = 0x1621
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEF26
  *c.R1.IY = 0x41DE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x58
  mem[0x0003] = 0xD4
  mem[0xEF7E] = 0x5E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCC1C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCC1C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB884) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB884, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6AD2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6AD2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5E21) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5E21, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEF26) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEF26, *c.R1.IX)
  }
  if (*c.R1.IY != 0x41DE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x41DE, *c.R1.IY)
  }
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

func TestDDCBD5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC41D
  *c.R1.BC = 0xC8B0
  *c.R1.DE = 0xCACB
  *c.R1.HL = 0x7687
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8DBC
  *c.R1.IY = 0xCC25
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x41
  mem[0x0003] = 0xD5
  mem[0x8DFD] = 0x71
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC41D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC41D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC8B0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC8B0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCACB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCACB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7675) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7675, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8DBC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8DBC, *c.R1.IX)
  }
  if (*c.R1.IY != 0xCC25) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xCC25, *c.R1.IY)
  }
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

func TestDDCBD6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x09EB
  *c.R1.BC = 0x769D
  *c.R1.DE = 0x7E07
  *c.R1.HL = 0x51F9
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5F03
  *c.R1.IY = 0x6280
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEA
  mem[0x0003] = 0xD6
  mem[0x5EED] = 0x73
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x09EB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x09EB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x769D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x769D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7E07) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7E07, *c.R1.DE)
  }
  if (*c.R1.HL != 0x51F9) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x51F9, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5F03) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5F03, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6280) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6280, *c.R1.IY)
  }
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

func TestDDCBD7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x241B
  *c.R1.BC = 0xEE10
  *c.R1.DE = 0xC152
  *c.R1.HL = 0x2F6D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE725
  *c.R1.IY = 0xC0D7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x80
  mem[0x0003] = 0xD7
  mem[0xE6A5] = 0x60
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x641B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x641B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEE10) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEE10, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC152) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC152, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2F6D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2F6D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE725) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE725, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC0D7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC0D7, *c.R1.IY)
  }
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

func TestDDCBD8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE3DC
  *c.R1.BC = 0x1981
  *c.R1.DE = 0xC97B
  *c.R1.HL = 0xCB42
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB30F
  *c.R1.IY = 0xB32A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4C
  mem[0x0003] = 0xD8
  mem[0xB35B] = 0x96
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE3DC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE3DC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9E81) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9E81, *c.R1.BC)
  }
  if (*c.R1.DE != 0xC97B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xC97B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCB42) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCB42, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB30F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB30F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB32A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB32A, *c.R1.IY)
  }
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

func TestDDCBD9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE9A0
  *c.R1.BC = 0xA7C7
  *c.R1.DE = 0xA476
  *c.R1.HL = 0x6057
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2642
  *c.R1.IY = 0x58A0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x52
  mem[0x0003] = 0xD9
  mem[0x2694] = 0xEF
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE9A0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE9A0, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA7EF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA7EF, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA476) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA476, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6057) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6057, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2642) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2642, *c.R1.IX)
  }
  if (*c.R1.IY != 0x58A0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x58A0, *c.R1.IY)
  }
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

func TestDDCBDA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6787
  *c.R1.BC = 0x26A7
  *c.R1.DE = 0xA194
  *c.R1.HL = 0x11D3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2D76
  *c.R1.IY = 0x7F80
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEA
  mem[0x0003] = 0xDA
  mem[0x2D60] = 0x82
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6787) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6787, *c.R1.AF)
  }
  if (*c.R1.BC != 0x26A7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x26A7, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8A94) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8A94, *c.R1.DE)
  }
  if (*c.R1.HL != 0x11D3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x11D3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2D76) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2D76, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7F80) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7F80, *c.R1.IY)
  }
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

func TestDDCBDB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF986
  *c.R1.BC = 0x6A4B
  *c.R1.DE = 0x6588
  *c.R1.HL = 0xD2C8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2B7D
  *c.R1.IY = 0x5847
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x4D
  mem[0x0003] = 0xDB
  mem[0x2BCA] = 0x10
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF986) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF986, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6A4B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6A4B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6518) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6518, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD2C8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD2C8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2B7D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2B7D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5847) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5847, *c.R1.IY)
  }
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

func TestDDCBDC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4C9E
  *c.R1.BC = 0xD94D
  *c.R1.DE = 0x9760
  *c.R1.HL = 0xB707
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7ED4
  *c.R1.IY = 0x5CC5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD3
  mem[0x0003] = 0xDC
  mem[0x7EA7] = 0x45
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4C9E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4C9E, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD94D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD94D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9760) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9760, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4D07) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4D07, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7ED4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7ED4, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5CC5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5CC5, *c.R1.IY)
  }
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

func TestDDCBDD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4B3B
  *c.R1.BC = 0xD351
  *c.R1.DE = 0x9BE9
  *c.R1.HL = 0x2310
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x58C1
  *c.R1.IY = 0xE430
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x6F
  mem[0x0003] = 0xDD
  mem[0x5930] = 0x20
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4B3B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4B3B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD351) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD351, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9BE9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9BE9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2328) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2328, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x58C1) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x58C1, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE430) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE430, *c.R1.IY)
  }
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

func TestDDCBDE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3B62
  *c.R1.BC = 0xCA1E
  *c.R1.DE = 0xA41A
  *c.R1.HL = 0x227A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x89D2
  *c.R1.IY = 0x7011
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x14
  mem[0x0003] = 0xDE
  mem[0x89E6] = 0x5E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3B62) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3B62, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCA1E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCA1E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA41A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA41A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x227A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x227A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x89D2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x89D2, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7011) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7011, *c.R1.IY)
  }
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

func TestDDCBDF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4C8A
  *c.R1.BC = 0x5B42
  *c.R1.DE = 0x50DD
  *c.R1.HL = 0x4BE0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD227
  *c.R1.IY = 0x4913
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEF
  mem[0x0003] = 0xDF
  mem[0xD216] = 0x72
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7A8A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7A8A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5B42) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5B42, *c.R1.BC)
  }
  if (*c.R1.DE != 0x50DD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x50DD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4BE0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4BE0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD227) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD227, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4913) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4913, *c.R1.IY)
  }
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

func TestDDCBE0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x440A
  *c.R1.BC = 0x713D
  *c.R1.DE = 0xACFC
  *c.R1.HL = 0xF762
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1C4B
  *c.R1.IY = 0xB6BA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x62
  mem[0x0003] = 0xE0
  mem[0x1CAD] = 0x46
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x440A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x440A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x563D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x563D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xACFC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xACFC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF762) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF762, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1C4B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1C4B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB6BA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB6BA, *c.R1.IY)
  }
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

func TestDDCBE1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC219
  *c.R1.BC = 0xAA6B
  *c.R1.DE = 0xDFBF
  *c.R1.HL = 0x6F10
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB931
  *c.R1.IY = 0xD3D6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2E
  mem[0x0003] = 0xE1
  mem[0xB95F] = 0x75
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC219) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC219, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAA75) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAA75, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDFBF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDFBF, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6F10) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6F10, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB931) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB931, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD3D6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD3D6, *c.R1.IY)
  }
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

func TestDDCBE2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x66D7
  *c.R1.BC = 0xABD0
  *c.R1.DE = 0xCB48
  *c.R1.HL = 0x8054
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEF50
  *c.R1.IY = 0x9997
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x85
  mem[0x0003] = 0xE2
  mem[0xEED5] = 0x72
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x66D7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x66D7, *c.R1.AF)
  }
  if (*c.R1.BC != 0xABD0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xABD0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7248) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7248, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8054) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8054, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEF50) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEF50, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9997) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9997, *c.R1.IY)
  }
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

func TestDDCBE3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7013
  *c.R1.BC = 0xE7ED
  *c.R1.DE = 0x7E1C
  *c.R1.HL = 0x57FB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7EC6
  *c.R1.IY = 0x75EB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF4
  mem[0x0003] = 0xE3
  mem[0x7EBA] = 0x34
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7013) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7013, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE7ED) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE7ED, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7E34) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7E34, *c.R1.DE)
  }
  if (*c.R1.HL != 0x57FB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x57FB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7EC6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7EC6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x75EB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x75EB, *c.R1.IY)
  }
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

func TestDDCBE4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1108
  *c.R1.BC = 0x6E70
  *c.R1.DE = 0xF0AF
  *c.R1.HL = 0x2F0C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x95C7
  *c.R1.IY = 0x6501
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBF
  mem[0x0003] = 0xE4
  mem[0x9586] = 0x34
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1108) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1108, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6E70) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6E70, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF0AF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF0AF, *c.R1.DE)
  }
  if (*c.R1.HL != 0x340C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x340C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x95C7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x95C7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6501) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6501, *c.R1.IY)
  }
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

func TestDDCBE5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x57CC
  *c.R1.BC = 0x5511
  *c.R1.DE = 0x2696
  *c.R1.HL = 0xB83D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6AB0
  *c.R1.IY = 0x0E90
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF2
  mem[0x0003] = 0xE5
  mem[0x6AA2] = 0x2E
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x57CC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x57CC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5511) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5511, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2696) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2696, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB83E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB83E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6AB0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6AB0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0E90) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0E90, *c.R1.IY)
  }
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

func TestDDCBE6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x207A
  *c.R1.BC = 0xA441
  *c.R1.DE = 0x1E03
  *c.R1.HL = 0xAC60
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD866
  *c.R1.IY = 0x5FDC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x25
  mem[0x0003] = 0xE6
  mem[0xD88B] = 0x4C
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x207A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x207A, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA441) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA441, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E03) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E03, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAC60) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAC60, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD866) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD866, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5FDC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5FDC, *c.R1.IY)
  }
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

func TestDDCBE7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC3C5
  *c.R1.BC = 0x7FA9
  *c.R1.DE = 0x4E07
  *c.R1.HL = 0xE02D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2A1B
  *c.R1.IY = 0x55B7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF3
  mem[0x0003] = 0xE7
  mem[0x2A0E] = 0xEB
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFBC5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFBC5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7FA9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7FA9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4E07) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4E07, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE02D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE02D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2A1B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2A1B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x55B7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x55B7, *c.R1.IY)
  }
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

func TestDDCBE8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6D1C
  *c.R1.BC = 0xA0C4
  *c.R1.DE = 0x93F0
  *c.R1.HL = 0xA0B4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4BDA
  *c.R1.IY = 0x7761
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xF2
  mem[0x0003] = 0xE8
  mem[0x4BCC] = 0xBA
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6D1C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6D1C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBAC4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBAC4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x93F0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x93F0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA0B4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA0B4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4BDA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4BDA, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7761) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7761, *c.R1.IY)
  }
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

func TestDDCBE9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEBE5
  *c.R1.BC = 0x0C2C
  *c.R1.DE = 0x1A2A
  *c.R1.HL = 0x2720
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x72DD
  *c.R1.IY = 0xA354
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x8A
  mem[0x0003] = 0xE9
  mem[0x7267] = 0x0A
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEBE5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEBE5, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0C2A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0C2A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1A2A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1A2A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2720) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2720, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x72DD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x72DD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA354) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA354, *c.R1.IY)
  }
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

func TestDDCBEA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x42D2
  *c.R1.BC = 0xDA7A
  *c.R1.DE = 0x757F
  *c.R1.HL = 0x6DA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA7E9
  *c.R1.IY = 0xB933
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x04
  mem[0x0003] = 0xEA
  mem[0xA7ED] = 0x5F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x42D2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x42D2, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDA7A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDA7A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7F7F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7F7F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6DA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6DA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA7E9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA7E9, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB933) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB933, *c.R1.IY)
  }
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

func TestDDCBEB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE945
  *c.R1.BC = 0x10AA
  *c.R1.DE = 0xF5F8
  *c.R1.HL = 0x7647
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x16DF
  *c.R1.IY = 0x93FB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x24
  mem[0x0003] = 0xEB
  mem[0x1703] = 0xF3
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE945) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE945, *c.R1.AF)
  }
  if (*c.R1.BC != 0x10AA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x10AA, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF5F3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF5F3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7647) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7647, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x16DF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x16DF, *c.R1.IX)
  }
  if (*c.R1.IY != 0x93FB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x93FB, *c.R1.IY)
  }
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

func TestDDCBEC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7180
  *c.R1.BC = 0xBC85
  *c.R1.DE = 0x7DD3
  *c.R1.HL = 0xF467
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDD88
  *c.R1.IY = 0x6A41
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x60
  mem[0x0003] = 0xEC
  mem[0xDDE8] = 0x00
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7180) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7180, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBC85) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBC85, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7DD3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7DD3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2067) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2067, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDD88) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDD88, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6A41) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6A41, *c.R1.IY)
  }
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

func TestDDCBED(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6B2F
  *c.R1.BC = 0x9762
  *c.R1.DE = 0x1F0A
  *c.R1.HL = 0xDB61
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF772
  *c.R1.IY = 0x33E3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xBE
  mem[0x0003] = 0xED
  mem[0xF730] = 0x6B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6B2F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6B2F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9762) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9762, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1F0A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1F0A, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDB6B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDB6B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF772) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF772, *c.R1.IX)
  }
  if (*c.R1.IY != 0x33E3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x33E3, *c.R1.IY)
  }
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

func TestDDCBEE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x79EA
  *c.R1.BC = 0xDC8A
  *c.R1.DE = 0x7887
  *c.R1.HL = 0x3BAA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6C28
  *c.R1.IY = 0xABBC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xDE
  mem[0x0003] = 0xEE
  mem[0x6C06] = 0xBD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x79EA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x79EA, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDC8A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDC8A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7887) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7887, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3BAA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3BAA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6C28) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6C28, *c.R1.IX)
  }
  if (*c.R1.IY != 0xABBC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xABBC, *c.R1.IY)
  }
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

func TestDDCBEF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x46C3
  *c.R1.BC = 0x2FC2
  *c.R1.DE = 0x8690
  *c.R1.HL = 0xA836
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCC68
  *c.R1.IY = 0xA8CE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x30
  mem[0x0003] = 0xEF
  mem[0xCC98] = 0x11
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x31C3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x31C3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2FC2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2FC2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8690) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8690, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA836) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA836, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCC68) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCC68, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA8CE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA8CE, *c.R1.IY)
  }
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

func TestDDCBF0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB330
  *c.R1.BC = 0x4469
  *c.R1.DE = 0x362B
  *c.R1.HL = 0xB515
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x13C0
  *c.R1.IY = 0x6479
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x2F
  mem[0x0003] = 0xF0
  mem[0x13EF] = 0xAD
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB330) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB330, *c.R1.AF)
  }
  if (*c.R1.BC != 0xED69) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xED69, *c.R1.BC)
  }
  if (*c.R1.DE != 0x362B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x362B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB515) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB515, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x13C0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x13C0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6479) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6479, *c.R1.IY)
  }
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

func TestDDCBF1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x94C0
  *c.R1.BC = 0x9AB0
  *c.R1.DE = 0xA0FD
  *c.R1.HL = 0x7C1D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x47BA
  *c.R1.IY = 0x8C81
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x40
  mem[0x0003] = 0xF1
  mem[0x47FA] = 0x78
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x94C0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x94C0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9A78) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9A78, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA0FD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA0FD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7C1D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7C1D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x47BA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x47BA, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8C81) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8C81, *c.R1.IY)
  }
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

func TestDDCBF2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5302
  *c.R1.BC = 0x9204
  *c.R1.DE = 0x20EC
  *c.R1.HL = 0xD640
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC947
  *c.R1.IY = 0x4EF1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x0F
  mem[0x0003] = 0xF2
  mem[0xC956] = 0x21
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5302) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5302, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9204) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9204, *c.R1.BC)
  }
  if (*c.R1.DE != 0x61EC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x61EC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD640) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD640, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC947) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC947, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4EF1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4EF1, *c.R1.IY)
  }
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

func TestDDCBF3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9950
  *c.R1.BC = 0xA3D2
  *c.R1.DE = 0x5058
  *c.R1.HL = 0x5CCC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1D96
  *c.R1.IY = 0x7C75
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x06
  mem[0x0003] = 0xF3
  mem[0x1D9C] = 0xE4
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9950) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9950, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA3D2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA3D2, *c.R1.BC)
  }
  if (*c.R1.DE != 0x50E4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x50E4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5CCC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5CCC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1D96) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1D96, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7C75) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7C75, *c.R1.IY)
  }
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

func TestDDCBF4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3712
  *c.R1.BC = 0x1F99
  *c.R1.DE = 0x4863
  *c.R1.HL = 0x47DE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1702
  *c.R1.IY = 0xC042
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x3B
  mem[0x0003] = 0xF4
  mem[0x173D] = 0xE1
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3712) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3712, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1F99) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1F99, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4863) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4863, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE1DE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE1DE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1702) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1702, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC042) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC042, *c.R1.IY)
  }
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

func TestDDCBF5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD83F
  *c.R1.BC = 0x1EC9
  *c.R1.DE = 0xD0DA
  *c.R1.HL = 0x4173
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEB3F
  *c.R1.IY = 0x1EAD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x64
  mem[0x0003] = 0xF5
  mem[0xEBA3] = 0xC5
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD83F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD83F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1EC9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1EC9, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD0DA) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD0DA, *c.R1.DE)
  }
  if (*c.R1.HL != 0x41C5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x41C5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEB3F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEB3F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1EAD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1EAD, *c.R1.IY)
  }
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

func TestDDCBF6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4D6C
  *c.R1.BC = 0x93AC
  *c.R1.DE = 0x810D
  *c.R1.HL = 0xCFE1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDC5A
  *c.R1.IY = 0xC33C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x7B
  mem[0x0003] = 0xF6
  mem[0xDCD5] = 0xA2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4D6C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4D6C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x93AC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x93AC, *c.R1.BC)
  }
  if (*c.R1.DE != 0x810D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x810D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCFE1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCFE1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDC5A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDC5A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC33C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC33C, *c.R1.IY)
  }
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

func TestDDCBF7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFE40
  *c.R1.BC = 0x7887
  *c.R1.DE = 0xB9DE
  *c.R1.HL = 0xC013
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x301E
  *c.R1.IY = 0x9710
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xC3
  mem[0x0003] = 0xF7
  mem[0x2FE1] = 0xA9
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE940) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE940, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7887) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7887, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB9DE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB9DE, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC013) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC013, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x301E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x301E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9710) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9710, *c.R1.IY)
  }
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

func TestDDCBF8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8278
  *c.R1.BC = 0x21A4
  *c.R1.DE = 0x1E5C
  *c.R1.HL = 0x4952
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x427F
  *c.R1.IY = 0x41E1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x59
  mem[0x0003] = 0xF8
  mem[0x42D8] = 0x28
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8278) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8278, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA8A4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA8A4, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1E5C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1E5C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4952) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4952, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x427F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x427F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x41E1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x41E1, *c.R1.IY)
  }
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

func TestDDCBF9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB2DF
  *c.R1.BC = 0xE9B8
  *c.R1.DE = 0x56C3
  *c.R1.HL = 0x16FF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD88F
  *c.R1.IY = 0x0BAB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x55
  mem[0x0003] = 0xF9
  mem[0xD8E4] = 0x14
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB2DF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB2DF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE994) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE994, *c.R1.BC)
  }
  if (*c.R1.DE != 0x56C3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x56C3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x16FF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x16FF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD88F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD88F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0BAB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0BAB, *c.R1.IY)
  }
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

func TestDDCBFA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x01F1
  *c.R1.BC = 0xBC0D
  *c.R1.DE = 0xD476
  *c.R1.HL = 0x1510
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9420
  *c.R1.IY = 0x93A3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x74
  mem[0x0003] = 0xFA
  mem[0x9494] = 0xFE
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x01F1) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x01F1, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBC0D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBC0D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFE76) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFE76, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1510) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1510, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9420) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9420, *c.R1.IX)
  }
  if (*c.R1.IY != 0x93A3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x93A3, *c.R1.IY)
  }
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

func TestDDCBFB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x709B
  *c.R1.BC = 0x14EB
  *c.R1.DE = 0xEC1C
  *c.R1.HL = 0xB844
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3453
  *c.R1.IY = 0xF2B0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xAF
  mem[0x0003] = 0xFB
  mem[0x3402] = 0x02
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x709B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x709B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x14EB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x14EB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEC82) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEC82, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB844) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB844, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3453) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3453, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF2B0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF2B0, *c.R1.IY)
  }
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

func TestDDCBFC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6C89
  *c.R1.BC = 0xA96E
  *c.R1.DE = 0xD27B
  *c.R1.HL = 0xD6A7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6139
  *c.R1.IY = 0xB4C1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xA1
  mem[0x0003] = 0xFC
  mem[0x60DA] = 0x10
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6C89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6C89, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA96E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA96E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD27B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD27B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x90A7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x90A7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x6139) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x6139, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB4C1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB4C1, *c.R1.IY)
  }
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

func TestDDCBFD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFB3F
  *c.R1.BC = 0x83F6
  *c.R1.DE = 0x2094
  *c.R1.HL = 0x3349
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3ED0
  *c.R1.IY = 0x6F0E
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0x28
  mem[0x0003] = 0xFD
  mem[0x3EF8] = 0xC2
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFB3F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFB3F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x83F6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x83F6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2094) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2094, *c.R1.DE)
  }
  if (*c.R1.HL != 0x33C2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x33C2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3ED0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3ED0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6F0E) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6F0E, *c.R1.IY)
  }
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

func TestDDCBFE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFC42
  *c.R1.BC = 0x50B7
  *c.R1.DE = 0xE98D
  *c.R1.HL = 0x3E45
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x41B5
  *c.R1.IY = 0x3410
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xEC
  mem[0x0003] = 0xFE
  mem[0x41A1] = 0xA1
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFC42) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFC42, *c.R1.AF)
  }
  if (*c.R1.BC != 0x50B7) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x50B7, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE98D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE98D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3E45) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3E45, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x41B5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x41B5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3410) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3410, *c.R1.IY)
  }
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

func TestDDCBFF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE666
  *c.R1.BC = 0x94D2
  *c.R1.DE = 0xAC90
  *c.R1.HL = 0x8F45
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0655
  *c.R1.IY = 0xBA29
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xCB
  mem[0x0002] = 0xD3
  mem[0x0003] = 0xFF
  mem[0x0628] = 0x2B
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAB66) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAB66, *c.R1.AF)
  }
  if (*c.R1.BC != 0x94D2) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x94D2, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAC90) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAC90, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8F45) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8F45, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0655) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0655, *c.R1.IX)
  }
  if (*c.R1.IY != 0xBA29) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xBA29, *c.R1.IY)
  }
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