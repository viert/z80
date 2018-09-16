package z80

import "testing"


func TestDD00(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
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
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x00
  mem[0x0002] = 0x00
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
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
  if (c.R != 0x03) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x03, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func TestDD09(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0D05
  *c.R1.BC = 0x1426
  *c.R1.DE = 0x53CE
  *c.R1.HL = 0x41E3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9EC0
  *c.R1.IY = 0x5C89
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x09
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0D34) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0D34, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1426) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1426, *c.R1.BC)
  }
  if (*c.R1.DE != 0x53CE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x53CE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x41E3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x41E3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB2E6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB2E6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5C89) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5C89, *c.R1.IY)
  }
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

func TestDD19(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1911
  *c.R1.BC = 0x0E0B
  *c.R1.DE = 0x2724
  *c.R1.HL = 0xBE62
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x824F
  *c.R1.IY = 0x760B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x19
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1928) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1928, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0E0B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0E0B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2724) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2724, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBE62) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBE62, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA973) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA973, *c.R1.IX)
  }
  if (*c.R1.IY != 0x760B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x760B, *c.R1.IY)
  }
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

func TestDD21(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC935
  *c.R1.BC = 0x4353
  *c.R1.DE = 0xBD22
  *c.R1.HL = 0x94D5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDADE
  *c.R1.IY = 0xAAD6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x21
  mem[0x0002] = 0xF2
  mem[0x0003] = 0x7C
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC935) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC935, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4353) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4353, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBD22) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBD22, *c.R1.DE)
  }
  if (*c.R1.HL != 0x94D5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x94D5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7CF2) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7CF2, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAAD6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAAD6, *c.R1.IY)
  }
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

func TestDD22(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5B1D
  *c.R1.BC = 0x45A1
  *c.R1.DE = 0x6DE8
  *c.R1.HL = 0x39D3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEBE7
  *c.R1.IY = 0x05B0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x22
  mem[0x0002] = 0x4F
  mem[0x0003] = 0xAD
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5B1D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5B1D, *c.R1.AF)
  }
  if (*c.R1.BC != 0x45A1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x45A1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6DE8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6DE8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x39D3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x39D3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEBE7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEBE7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x05B0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x05B0, *c.R1.IY)
  }
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

func TestDD23(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9095
  *c.R1.BC = 0xAC3C
  *c.R1.DE = 0x4D90
  *c.R1.HL = 0x379B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD50B
  *c.R1.IY = 0xA157
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x23
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9095) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9095, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAC3C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAC3C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4D90) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4D90, *c.R1.DE)
  }
  if (*c.R1.HL != 0x379B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x379B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD50C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD50C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA157) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA157, *c.R1.IY)
  }
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

func TestDD24(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0698
  *c.R1.BC = 0xDCD0
  *c.R1.DE = 0xA31B
  *c.R1.HL = 0xD527
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8CDA
  *c.R1.IY = 0xB096
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x24
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0688) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0688, *c.R1.AF)
  }
  if (*c.R1.BC != 0xDCD0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xDCD0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA31B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA31B, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD527) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD527, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8DDA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8DDA, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB096) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB096, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD25(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5ACC
  *c.R1.BC = 0x206B
  *c.R1.DE = 0xED10
  *c.R1.HL = 0x6EAB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBB3C
  *c.R1.IY = 0x5EBD
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x25
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5AAA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5AAA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x206B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x206B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xED10) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xED10, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6EAB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6EAB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBA3C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBA3C, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5EBD) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5EBD, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD26(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9522
  *c.R1.BC = 0xEDE0
  *c.R1.DE = 0xA352
  *c.R1.HL = 0xADEA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5F40
  *c.R1.IY = 0x82E1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x26
  mem[0x0002] = 0xAD
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9522) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9522, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEDE0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEDE0, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA352) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA352, *c.R1.DE)
  }
  if (*c.R1.HL != 0xADEA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xADEA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAD40) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAD40, *c.R1.IX)
  }
  if (*c.R1.IY != 0x82E1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x82E1, *c.R1.IY)
  }
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

func TestDD29(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAC80
  *c.R1.BC = 0x0F0E
  *c.R1.DE = 0x72C8
  *c.R1.HL = 0x1F2A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5195
  *c.R1.IY = 0x7D8A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x29
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0xACA0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xACA0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F0E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F0E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x72C8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x72C8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1F2A) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1F2A, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA32A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA32A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7D8A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7D8A, *c.R1.IY)
  }
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

func TestDD2A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3D36
  *c.R1.BC = 0xB24E
  *c.R1.DE = 0xBDBC
  *c.R1.HL = 0xCA4E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBA65
  *c.R1.IY = 0xE7CE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x2A
  mem[0x0002] = 0xBC
  mem[0x0003] = 0x40
  mem[0x40BC] = 0xB5
  mem[0x40BD] = 0x30
  for c.TStates < 20 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3D36) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3D36, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB24E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB24E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBDBC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBDBC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCA4E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCA4E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x30B5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x30B5, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE7CE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE7CE, *c.R1.IY)
  }
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

func TestDD2B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAD4B
  *c.R1.BC = 0xD5E6
  *c.R1.DE = 0x9377
  *c.R1.HL = 0xF132
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7A17
  *c.R1.IY = 0x2188
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x2B
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAD4B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAD4B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD5E6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD5E6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9377) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9377, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF132) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF132, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7A16) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7A16, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2188) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2188, *c.R1.IY)
  }
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

func TestDD2C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8838
  *c.R1.BC = 0xF2F3
  *c.R1.DE = 0xD277
  *c.R1.HL = 0x9153
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC62F
  *c.R1.IY = 0xB002
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x2C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8830) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8830, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF2F3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF2F3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD277) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD277, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9153) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9153, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC630) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC630, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB002) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB002, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD2D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x39BC
  *c.R1.BC = 0xB23C
  *c.R1.DE = 0x6E11
  *c.R1.HL = 0x5A49
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0267
  *c.R1.IY = 0xAB03
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x2D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3922) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3922, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB23C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB23C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6E11) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6E11, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5A49) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5A49, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0266) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0266, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAB03) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAB03, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD2E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9ACA
  *c.R1.BC = 0xA04A
  *c.R1.DE = 0xB49F
  *c.R1.HL = 0xA4A6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBD90
  *c.R1.IY = 0x38A1
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x2E
  mem[0x0002] = 0x1C
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9ACA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9ACA, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA04A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA04A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB49F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB49F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA4A6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA4A6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
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
  if (*c.R1.IY != 0x38A1) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x38A1, *c.R1.IY)
  }
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

func TestDD34(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8304
  *c.R1.BC = 0xD1FC
  *c.R1.DE = 0xB80B
  *c.R1.HL = 0x8082
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xDEA9
  *c.R1.IY = 0x6FD8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x34
  mem[0x0002] = 0xE6
  mem[0xDE8F] = 0x57
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8308) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8308, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD1FC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD1FC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB80B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB80B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8082) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8082, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xDEA9) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xDEA9, *c.R1.IX)
  }
  if (*c.R1.IY != 0x6FD8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x6FD8, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 23) {
    t.Errorf("TStates mismatch: %d expected, got %d", 23, c.TStates)
  }
}

func TestDD35(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8681
  *c.R1.BC = 0x4641
  *c.R1.DE = 0x1EF6
  *c.R1.HL = 0x10AB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC733
  *c.R1.IY = 0x8EC4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x35
  mem[0x0002] = 0x60
  mem[0xC793] = 0xF7
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x86A3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x86A3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x4641) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x4641, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1EF6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1EF6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x10AB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x10AB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC733) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC733, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8EC4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8EC4, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 23) {
    t.Errorf("TStates mismatch: %d expected, got %d", 23, c.TStates)
  }
}

func TestDD36(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x76DC
  *c.R1.BC = 0x2530
  *c.R1.DE = 0x5158
  *c.R1.HL = 0x877D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB5C6
  *c.R1.IY = 0x8D3C
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x36
  mem[0x0002] = 0x35
  mem[0x0003] = 0xB5
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x76DC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x76DC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2530) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2530, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5158) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5158, *c.R1.DE)
  }
  if (*c.R1.HL != 0x877D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x877D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB5C6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB5C6, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8D3C) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8D3C, *c.R1.IY)
  }
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

func TestDD39(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x875B
  *c.R1.BC = 0xA334
  *c.R1.DE = 0xD79D
  *c.R1.HL = 0x59E4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB11A
  *c.R1.IY = 0x4C88
  *c.R1.SP = 0xFA4A
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x39
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8769) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8769, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA334) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA334, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD79D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD79D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x59E4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x59E4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAB64) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAB64, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4C88) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4C88, *c.R1.IY)
  }
  if (*c.R1.SP != 0xFA4A) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xFA4A, *c.R1.SP)
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

func TestDD44(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB37E
  *c.R1.BC = 0xCBB0
  *c.R1.DE = 0x36E8
  *c.R1.HL = 0x3F45
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2702
  *c.R1.IY = 0xB3B9
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x44
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xB37E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xB37E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x27B0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x27B0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x36E8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x36E8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3F45) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3F45, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2702) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2702, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB3B9) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB3B9, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD45(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4E10
  *c.R1.BC = 0x5C6D
  *c.R1.DE = 0xD11D
  *c.R1.HL = 0x1736
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7298
  *c.R1.IY = 0x2D10
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x45
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4E10) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4E10, *c.R1.AF)
  }
  if (*c.R1.BC != 0x986D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x986D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD11D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD11D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1736) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1736, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7298) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7298, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2D10) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2D10, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD46(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC758
  *c.R1.BC = 0xBF29
  *c.R1.DE = 0x66F2
  *c.R1.HL = 0x29EF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5CC7
  *c.R1.IY = 0x407D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x46
  mem[0x0002] = 0x68
  mem[0x5D2F] = 0x8D
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC758) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC758, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8D29) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8D29, *c.R1.BC)
  }
  if (*c.R1.DE != 0x66F2) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x66F2, *c.R1.DE)
  }
  if (*c.R1.HL != 0x29EF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x29EF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5CC7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5CC7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x407D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x407D, *c.R1.IY)
  }
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

func TestDD4C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE15C
  *c.R1.BC = 0x75EC
  *c.R1.DE = 0x7531
  *c.R1.HL = 0xAE9E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3ED8
  *c.R1.IY = 0x03B7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x4C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE15C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE15C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x753E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x753E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7531) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7531, *c.R1.DE)
  }
  if (*c.R1.HL != 0xAE9E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xAE9E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3ED8) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3ED8, *c.R1.IX)
  }
  if (*c.R1.IY != 0x03B7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x03B7, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD4D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x469E
  *c.R1.BC = 0x7864
  *c.R1.DE = 0x6A5A
  *c.R1.HL = 0x00E2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA1AA
  *c.R1.IY = 0x0D6F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x4D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x469E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x469E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x78AA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x78AA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6A5A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6A5A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x00E2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x00E2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA1AA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA1AA, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0D6F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0D6F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD4E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7BF7
  *c.R1.BC = 0x6605
  *c.R1.DE = 0x8D55
  *c.R1.HL = 0xDEF2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD94B
  *c.R1.IY = 0x17FB
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x4E
  mem[0x0002] = 0x2E
  mem[0xD979] = 0x76
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7BF7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7BF7, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6676) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6676, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8D55) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8D55, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDEF2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDEF2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD94B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD94B, *c.R1.IX)
  }
  if (*c.R1.IY != 0x17FB) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x17FB, *c.R1.IY)
  }
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

func TestDD54(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8376
  *c.R1.BC = 0x0D13
  *c.R1.DE = 0xC767
  *c.R1.HL = 0x3119
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x4B6D
  *c.R1.IY = 0x030B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x54
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8376) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8376, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0D13) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0D13, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4B67) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4B67, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3119) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3119, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4B6D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4B6D, *c.R1.IX)
  }
  if (*c.R1.IY != 0x030B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x030B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD55(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFF78
  *c.R1.BC = 0x85E3
  *c.R1.DE = 0x566B
  *c.R1.HL = 0x8F3A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD7D7
  *c.R1.IY = 0x4E0B
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x55
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF78) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF78, *c.R1.AF)
  }
  if (*c.R1.BC != 0x85E3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x85E3, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD76B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD76B, *c.R1.DE)
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
  if (*c.R1.IX != 0xD7D7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD7D7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4E0B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4E0B, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD56(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x97B3
  *c.R1.BC = 0xB617
  *c.R1.DE = 0xBB50
  *c.R1.HL = 0x81D1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA306
  *c.R1.IY = 0x7A49
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x56
  mem[0x0002] = 0xF4
  mem[0xA2FA] = 0xDE
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x97B3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x97B3, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB617) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB617, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDE50) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDE50, *c.R1.DE)
  }
  if (*c.R1.HL != 0x81D1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x81D1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA306) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA306, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7A49) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7A49, *c.R1.IY)
  }
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

func TestDD5C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xAF82
  *c.R1.BC = 0x24BF
  *c.R1.DE = 0x2793
  *c.R1.HL = 0xF925
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xF9A3
  *c.R1.IY = 0x0B82
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x5C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAF82) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAF82, *c.R1.AF)
  }
  if (*c.R1.BC != 0x24BF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x24BF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x27F9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x27F9, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF925) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF925, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xF9A3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xF9A3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0B82) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0B82, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD5D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x36CB
  *c.R1.BC = 0x97A9
  *c.R1.DE = 0x400D
  *c.R1.HL = 0x30FE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3340
  *c.R1.IY = 0xB3ED
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x5D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x36CB) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x36CB, *c.R1.AF)
  }
  if (*c.R1.BC != 0x97A9) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x97A9, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4040) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4040, *c.R1.DE)
  }
  if (*c.R1.HL != 0x30FE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x30FE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3340) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3340, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB3ED) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB3ED, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD5E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA220
  *c.R1.BC = 0x389D
  *c.R1.DE = 0x2FF8
  *c.R1.HL = 0x368C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8D32
  *c.R1.IY = 0x3512
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x5E
  mem[0x0002] = 0x8F
  mem[0x8CC1] = 0xCE
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA220) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA220, *c.R1.AF)
  }
  if (*c.R1.BC != 0x389D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x389D, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2FCE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2FCE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x368C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x368C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8D32) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8D32, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3512) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3512, *c.R1.IY)
  }
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

func TestDD60(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2392
  *c.R1.BC = 0x7F6A
  *c.R1.DE = 0x3DC0
  *c.R1.HL = 0xCEFB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x44A0
  *c.R1.IY = 0xC424
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x60
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2392) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2392, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7F6A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7F6A, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3DC0) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3DC0, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCEFB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCEFB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7FA0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7FA0, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC424) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC424, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD61(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x76ED
  *c.R1.BC = 0x268C
  *c.R1.DE = 0xD5C8
  *c.R1.HL = 0xBAB0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB650
  *c.R1.IY = 0x0A93
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x61
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x76ED) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x76ED, *c.R1.AF)
  }
  if (*c.R1.BC != 0x268C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x268C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD5C8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD5C8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xBAB0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xBAB0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8C50) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8C50, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0A93) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0A93, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD62(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4C6F
  *c.R1.BC = 0xB482
  *c.R1.DE = 0xFEF4
  *c.R1.HL = 0x62E7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6E25
  *c.R1.IY = 0x9655
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x62
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4C6F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4C6F, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB482) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB482, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFEF4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFEF4, *c.R1.DE)
  }
  if (*c.R1.HL != 0x62E7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x62E7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFE25) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFE25, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9655) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9655, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD63(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6E9A
  *c.R1.BC = 0x5499
  *c.R1.DE = 0x3C8F
  *c.R1.HL = 0x1F64
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBF35
  *c.R1.IY = 0x0DF7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x63
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6E9A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6E9A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5499) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5499, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3C8F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3C8F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1F64) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1F64, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8F35) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8F35, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0DF7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0DF7, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD64(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x47F6
  *c.R1.BC = 0x1B7A
  *c.R1.DE = 0xA55E
  *c.R1.HL = 0x2FC2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xEFC7
  *c.R1.IY = 0xACA0
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x64
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x47F6) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x47F6, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1B7A) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1B7A, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA55E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA55E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2FC2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2FC2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xEFC7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xEFC7, *c.R1.IX)
  }
  if (*c.R1.IY != 0xACA0) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xACA0, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD65(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD786
  *c.R1.BC = 0x7D1D
  *c.R1.DE = 0xB659
  *c.R1.HL = 0x77E8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x58FA
  *c.R1.IY = 0x006D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x65
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD786) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD786, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7D1D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7D1D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB659) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB659, *c.R1.DE)
  }
  if (*c.R1.HL != 0x77E8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x77E8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFAFA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFAFA, *c.R1.IX)
  }
  if (*c.R1.IY != 0x006D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x006D, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD66(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x84C2
  *c.R1.BC = 0x79B1
  *c.R1.DE = 0xCA4A
  *c.R1.HL = 0xAAA0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xCE5D
  *c.R1.IY = 0xDD2D
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x66
  mem[0x0002] = 0xB5
  mem[0xCE12] = 0x03
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x84C2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x84C2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x79B1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x79B1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCA4A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCA4A, *c.R1.DE)
  }
  if (*c.R1.HL != 0x03A0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x03A0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xCE5D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xCE5D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xDD2D) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xDD2D, *c.R1.IY)
  }
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

func TestDD67(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x967C
  *c.R1.BC = 0x511E
  *c.R1.DE = 0x336D
  *c.R1.HL = 0x40F6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x66E7
  *c.R1.IY = 0x5BE2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x67
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x967C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x967C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x511E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x511E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x336D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x336D, *c.R1.DE)
  }
  if (*c.R1.HL != 0x40F6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x40F6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x96E7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x96E7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5BE2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5BE2, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD68(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4A9D
  *c.R1.BC = 0xEFA8
  *c.R1.DE = 0xFEBD
  *c.R1.HL = 0x07E4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5FD8
  *c.R1.IY = 0xB23F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x68
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4A9D) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4A9D, *c.R1.AF)
  }
  if (*c.R1.BC != 0xEFA8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xEFA8, *c.R1.BC)
  }
  if (*c.R1.DE != 0xFEBD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xFEBD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x07E4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x07E4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5FEF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5FEF, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB23F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB23F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD69(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6466
  *c.R1.BC = 0x2142
  *c.R1.DE = 0x2523
  *c.R1.HL = 0x82B3
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x6479
  *c.R1.IY = 0x04A7
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x69
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6466) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6466, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2142) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2142, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2523) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2523, *c.R1.DE)
  }
  if (*c.R1.HL != 0x82B3) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x82B3, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
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
  if (*c.R1.IY != 0x04A7) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x04A7, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD6A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x401F
  *c.R1.BC = 0x61F1
  *c.R1.DE = 0x4B08
  *c.R1.HL = 0xFA88
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC37F
  *c.R1.IY = 0xD8F6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x6A
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x401F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x401F, *c.R1.AF)
  }
  if (*c.R1.BC != 0x61F1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x61F1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4B08) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4B08, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFA88) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFA88, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC34B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC34B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD8F6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD8F6, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD6B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6DC7
  *c.R1.BC = 0xE2AE
  *c.R1.DE = 0x40BD
  *c.R1.HL = 0xF3C0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2290
  *c.R1.IY = 0x2749
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x6B
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6DC7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6DC7, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE2AE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE2AE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x40BD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x40BD, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF3C0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF3C0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x22BD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x22BD, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2749) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2749, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD6C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3939
  *c.R1.BC = 0x90DA
  *c.R1.DE = 0x62DC
  *c.R1.HL = 0x7C31
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x412F
  *c.R1.IY = 0x7211
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x6C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3939) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3939, *c.R1.AF)
  }
  if (*c.R1.BC != 0x90DA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x90DA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x62DC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x62DC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7C31) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7C31, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x4141) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x4141, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7211) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7211, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD6D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3964
  *c.R1.BC = 0xFF3F
  *c.R1.DE = 0x23D4
  *c.R1.HL = 0xC7C7
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9B70
  *c.R1.IY = 0x20C6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x6D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3964) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3964, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFF3F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFF3F, *c.R1.BC)
  }
  if (*c.R1.DE != 0x23D4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x23D4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xC7C7) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC7C7, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9B70) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9B70, *c.R1.IX)
  }
  if (*c.R1.IY != 0x20C6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x20C6, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD6E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x223F
  *c.R1.BC = 0xF661
  *c.R1.DE = 0xB61C
  *c.R1.HL = 0x0F53
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC648
  *c.R1.IY = 0xFAE8
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x6E
  mem[0x0002] = 0x2C
  mem[0xC674] = 0x6B
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x223F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x223F, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF661) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF661, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB61C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB61C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0F6B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0F6B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC648) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC648, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFAE8) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFAE8, *c.R1.IY)
  }
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

func TestDD6F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6E84
  *c.R1.BC = 0x9CD4
  *c.R1.DE = 0xA293
  *c.R1.HL = 0x647D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0D0B
  *c.R1.IY = 0x4A56
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x6F
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6E84) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6E84, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9CD4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9CD4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA293) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA293, *c.R1.DE)
  }
  if (*c.R1.HL != 0x647D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x647D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0D6E) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0D6E, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4A56) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4A56, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD70(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xD09F
  *c.R1.BC = 0xFE00
  *c.R1.DE = 0x231E
  *c.R1.HL = 0x31EC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x05FA
  *c.R1.IY = 0xEA92
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x70
  mem[0x0002] = 0xF6
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD09F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD09F, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFE00) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFE00, *c.R1.BC)
  }
  if (*c.R1.DE != 0x231E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x231E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x31EC) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x31EC, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x05FA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x05FA, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEA92) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEA92, *c.R1.IY)
  }
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

func TestDD71(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEBEE
  *c.R1.BC = 0x151C
  *c.R1.DE = 0x05C7
  *c.R1.HL = 0xEE08
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x3722
  *c.R1.IY = 0x2EC6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x71
  mem[0x0002] = 0x23
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEBEE) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEBEE, *c.R1.AF)
  }
  if (*c.R1.BC != 0x151C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x151C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x05C7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x05C7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xEE08) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xEE08, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3722) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3722, *c.R1.IX)
  }
  if (*c.R1.IY != 0x2EC6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x2EC6, *c.R1.IY)
  }
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

func TestDD72(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x80C9
  *c.R1.BC = 0xAC1E
  *c.R1.DE = 0x63BD
  *c.R1.HL = 0x828B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8DFF
  *c.R1.IY = 0x94EF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x72
  mem[0x0002] = 0x93
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x80C9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x80C9, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAC1E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAC1E, *c.R1.BC)
  }
  if (*c.R1.DE != 0x63BD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x63BD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x828B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x828B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8DFF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8DFF, *c.R1.IX)
  }
  if (*c.R1.IY != 0x94EF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x94EF, *c.R1.IY)
  }
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

func TestDD73(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8F3E
  *c.R1.BC = 0xB5A3
  *c.R1.DE = 0x07DE
  *c.R1.HL = 0x0B0C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x79C6
  *c.R1.IY = 0xAE79
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x73
  mem[0x0002] = 0x57
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8F3E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8F3E, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB5A3) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB5A3, *c.R1.BC)
  }
  if (*c.R1.DE != 0x07DE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x07DE, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0B0C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0B0C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x79C6) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x79C6, *c.R1.IX)
  }
  if (*c.R1.IY != 0xAE79) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xAE79, *c.R1.IY)
  }
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

func TestDD74(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4AE0
  *c.R1.BC = 0x49C5
  *c.R1.DE = 0x3DEB
  *c.R1.HL = 0x0125
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5910
  *c.R1.IY = 0x429A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x74
  mem[0x0002] = 0xB9
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4AE0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4AE0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x49C5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x49C5, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3DEB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3DEB, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0125) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0125, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x5910) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x5910, *c.R1.IX)
  }
  if (*c.R1.IY != 0x429A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x429A, *c.R1.IY)
  }
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

func TestDD75(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5772
  *c.R1.BC = 0xE833
  *c.R1.DE = 0xB63E
  *c.R1.HL = 0x734F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAE4C
  *c.R1.IY = 0xE8C2
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x75
  mem[0x0002] = 0x30
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5772) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5772, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE833) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE833, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB63E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB63E, *c.R1.DE)
  }
  if (*c.R1.HL != 0x734F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x734F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAE4C) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAE4C, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE8C2) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE8C2, *c.R1.IY)
  }
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

func TestDD77(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDC56
  *c.R1.BC = 0xD893
  *c.R1.DE = 0x4116
  *c.R1.HL = 0xF2D2
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xA181
  *c.R1.IY = 0x3157
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x77
  mem[0x0002] = 0x8C
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDC56) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDC56, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD893) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD893, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4116) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4116, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF2D2) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF2D2, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xA181) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xA181, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3157) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3157, *c.R1.IY)
  }
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

func TestDD7C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7558
  *c.R1.BC = 0x7705
  *c.R1.DE = 0xAC92
  *c.R1.HL = 0xA6A1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x8CDE
  *c.R1.IY = 0x7507
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x7C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8C58) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8C58, *c.R1.AF)
  }
  if (*c.R1.BC != 0x7705) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x7705, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAC92) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAC92, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA6A1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA6A1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x8CDE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x8CDE, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7507) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7507, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD7D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6C18
  *c.R1.BC = 0x93FB
  *c.R1.DE = 0x6BDD
  *c.R1.HL = 0x3A10
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD7CB
  *c.R1.IY = 0xC0F6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x7D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCB18) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCB18, *c.R1.AF)
  }
  if (*c.R1.BC != 0x93FB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x93FB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6BDD) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6BDD, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3A10) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3A10, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD7CB) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD7CB, *c.R1.IX)
  }
  if (*c.R1.IY != 0xC0F6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xC0F6, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD7E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6A66
  *c.R1.BC = 0x1F77
  *c.R1.DE = 0x6220
  *c.R1.HL = 0x0C40
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x1CF4
  *c.R1.IY = 0x1A1F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x7E
  mem[0x0002] = 0xBC
  mem[0x1CB0] = 0x57
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5766) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5766, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1F77) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1F77, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6220) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6220, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0C40) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0C40, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x1CF4) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x1CF4, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1A1F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1A1F, *c.R1.IY)
  }
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

func TestDD84(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2E47
  *c.R1.BC = 0x1DE8
  *c.R1.DE = 0xB8B9
  *c.R1.HL = 0x78A6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9F1D
  *c.R1.IY = 0xB11F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x84
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCD98) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCD98, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1DE8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1DE8, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB8B9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB8B9, *c.R1.DE)
  }
  if (*c.R1.HL != 0x78A6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x78A6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9F1D) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9F1D, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB11F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB11F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD85(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB27A
  *c.R1.BC = 0xB1FF
  *c.R1.DE = 0x8D7B
  *c.R1.HL = 0x40C0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB513
  *c.R1.IY = 0x0688
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x85
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC580) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC580, *c.R1.AF)
  }
  if (*c.R1.BC != 0xB1FF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xB1FF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8D7B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8D7B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x40C0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x40C0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB513) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB513, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0688) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0688, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD86(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4EFA
  *c.R1.BC = 0xD085
  *c.R1.DE = 0x5BAC
  *c.R1.HL = 0xE364
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB5B5
  *c.R1.IY = 0xFE3A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x86
  mem[0x0002] = 0xC1
  mem[0xB576] = 0x5B
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA9BC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA9BC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD085) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD085, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5BAC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5BAC, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE364) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE364, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB5B5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB5B5, *c.R1.IX)
  }
  if (*c.R1.IY != 0xFE3A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xFE3A, *c.R1.IY)
  }
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

func TestDD8C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBC63
  *c.R1.BC = 0x8FDC
  *c.R1.DE = 0xEA8F
  *c.R1.HL = 0x9734
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0EB3
  *c.R1.IY = 0x1B54
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x8C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCB98) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCB98, *c.R1.AF)
  }
  if (*c.R1.BC != 0x8FDC) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x8FDC, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEA8F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEA8F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9734) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9734, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0EB3) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0EB3, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1B54) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1B54, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD8D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xB61F
  *c.R1.BC = 0x1C81
  *c.R1.DE = 0xB6FB
  *c.R1.HL = 0xD6E5
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x09BE
  *c.R1.IY = 0xA736
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x8D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7535) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7535, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1C81) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1C81, *c.R1.BC)
  }
  if (*c.R1.DE != 0xB6FB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xB6FB, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD6E5) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD6E5, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x09BE) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x09BE, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA736) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA736, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD8E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4ED4
  *c.R1.BC = 0x182D
  *c.R1.DE = 0xAB17
  *c.R1.HL = 0x94AE
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBB97
  *c.R1.IY = 0x87DA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x8E
  mem[0x0002] = 0x25
  mem[0xBBBC] = 0x32
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8094) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8094, *c.R1.AF)
  }
  if (*c.R1.BC != 0x182D) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x182D, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAB17) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAB17, *c.R1.DE)
  }
  if (*c.R1.HL != 0x94AE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x94AE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBB97) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBB97, *c.R1.IX)
  }
  if (*c.R1.IY != 0x87DA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x87DA, *c.R1.IY)
  }
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

func TestDD94(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7EF1
  *c.R1.BC = 0x9EFE
  *c.R1.DE = 0x6EA1
  *c.R1.HL = 0xFC55
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0A09
  *c.R1.IY = 0x89C5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x94
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7422) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7422, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9EFE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9EFE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x6EA1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x6EA1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFC55) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFC55, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0A09) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0A09, *c.R1.IX)
  }
  if (*c.R1.IY != 0x89C5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x89C5, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD95(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x2920
  *c.R1.BC = 0x59AB
  *c.R1.DE = 0x428C
  *c.R1.HL = 0x3A94
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x44FD
  *c.R1.IY = 0xF243
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x95
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2C3B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2C3B, *c.R1.AF)
  }
  if (*c.R1.BC != 0x59AB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x59AB, *c.R1.BC)
  }
  if (*c.R1.DE != 0x428C) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x428C, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3A94) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3A94, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x44FD) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x44FD, *c.R1.IX)
  }
  if (*c.R1.IY != 0xF243) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xF243, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD96(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9B76
  *c.R1.BC = 0x461F
  *c.R1.DE = 0xCED7
  *c.R1.HL = 0xDB3F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2C66
  *c.R1.IY = 0x9DBF
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x96
  mem[0x0002] = 0x5F
  mem[0x2CC5] = 0x49
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5206) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5206, *c.R1.AF)
  }
  if (*c.R1.BC != 0x461F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x461F, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCED7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCED7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDB3F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDB3F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2C66) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2C66, *c.R1.IX)
  }
  if (*c.R1.IY != 0x9DBF) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x9DBF, *c.R1.IY)
  }
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

func TestDD9C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFAF4
  *c.R1.BC = 0x670E
  *c.R1.DE = 0xAFCC
  *c.R1.HL = 0x8B34
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x285F
  *c.R1.IY = 0x1CAA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x9C
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD282) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD282, *c.R1.AF)
  }
  if (*c.R1.BC != 0x670E) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x670E, *c.R1.BC)
  }
  if (*c.R1.DE != 0xAFCC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xAFCC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x8B34) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x8B34, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x285F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x285F, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1CAA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1CAA, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD9D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF827
  *c.R1.BC = 0x0CDB
  *c.R1.DE = 0xDF32
  *c.R1.HL = 0xD0E4
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9B12
  *c.R1.IY = 0x7D07
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x9D
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE5A2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE5A2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0CDB) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0CDB, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDF32) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDF32, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD0E4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD0E4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9B12) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9B12, *c.R1.IX)
  }
  if (*c.R1.IY != 0x7D07) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x7D07, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDD9E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x938E
  *c.R1.BC = 0xF9C5
  *c.R1.DE = 0xCBC4
  *c.R1.HL = 0xCA21
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB4CC
  *c.R1.IY = 0x46FA
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0x9E
  mem[0x0002] = 0x14
  mem[0xB4E0] = 0xB5
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDE9B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDE9B, *c.R1.AF)
  }
  if (*c.R1.BC != 0xF9C5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xF9C5, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCBC4) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCBC4, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCA21) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCA21, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB4CC) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB4CC, *c.R1.IX)
  }
  if (*c.R1.IY != 0x46FA) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x46FA, *c.R1.IY)
  }
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

func TestDDA4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x52F5
  *c.R1.BC = 0xBA53
  *c.R1.DE = 0xACFC
  *c.R1.HL = 0x9481
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x2F8B
  *c.R1.IY = 0xEDF6
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xA4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0210) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0210, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBA53) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBA53, *c.R1.BC)
  }
  if (*c.R1.DE != 0xACFC) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xACFC, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9481) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9481, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x2F8B) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x2F8B, *c.R1.IX)
  }
  if (*c.R1.IY != 0xEDF6) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xEDF6, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDA5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBAAF
  *c.R1.BC = 0xA675
  *c.R1.DE = 0xD757
  *c.R1.HL = 0xF1DB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xFDEF
  *c.R1.IY = 0xD8CE
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xA5
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xAABC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xAABC, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA675) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA675, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD757) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD757, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF1DB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF1DB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xFDEF) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xFDEF, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD8CE) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD8CE, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDA6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1DA4
  *c.R1.BC = 0x20C4
  *c.R1.DE = 0xEBC3
  *c.R1.HL = 0xDA8D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x7E95
  *c.R1.IY = 0x5E8A
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xA6
  mem[0x0002] = 0x41
  mem[0x7ED6] = 0xC7
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0514) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0514, *c.R1.AF)
  }
  if (*c.R1.BC != 0x20C4) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x20C4, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEBC3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEBC3, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDA8D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDA8D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x7E95) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x7E95, *c.R1.IX)
  }
  if (*c.R1.IY != 0x5E8A) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x5E8A, *c.R1.IY)
  }
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

func TestDDAC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEF15
  *c.R1.BC = 0x2A7C
  *c.R1.DE = 0x17E5
  *c.R1.HL = 0x3F6E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xAFFA
  *c.R1.IY = 0xA0B5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xAC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x2A7C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x2A7C, *c.R1.BC)
  }
  if (*c.R1.DE != 0x17E5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x17E5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3F6E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3F6E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xAFFA) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xAFFA, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA0B5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA0B5, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDAD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xBA2E
  *c.R1.BC = 0x6BA1
  *c.R1.DE = 0xEF1B
  *c.R1.HL = 0x5713
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBA38
  *c.R1.IY = 0xA708
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xAD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8284) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8284, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6BA1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6BA1, *c.R1.BC)
  }
  if (*c.R1.DE != 0xEF1B) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xEF1B, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5713) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5713, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBA38) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBA38, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA708) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA708, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDAE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8009
  *c.R1.BC = 0x3AD6
  *c.R1.DE = 0xA721
  *c.R1.HL = 0x2100
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xE909
  *c.R1.IY = 0x87B4
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xAE
  mem[0x0002] = 0x72
  mem[0xE97B] = 0xC3
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3AD6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3AD6, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA721) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA721, *c.R1.DE)
  }
  if (*c.R1.HL != 0x2100) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x2100, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xE909) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xE909, *c.R1.IX)
  }
  if (*c.R1.IY != 0x87B4) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x87B4, *c.R1.IY)
  }
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

func TestDDB4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1CCD
  *c.R1.BC = 0x29AA
  *c.R1.DE = 0x2E82
  *c.R1.HL = 0x4DC8
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9C04
  *c.R1.IY = 0x8BE3
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xB4
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9C8C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9C8C, *c.R1.AF)
  }
  if (*c.R1.BC != 0x29AA) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x29AA, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2E82) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2E82, *c.R1.DE)
  }
  if (*c.R1.HL != 0x4DC8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x4DC8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9C04) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9C04, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8BE3) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8BE3, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDB5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x46B4
  *c.R1.BC = 0xFC93
  *c.R1.DE = 0x7A06
  *c.R1.HL = 0x0518
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0AC5
  *c.R1.IY = 0x4150
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xB5
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC780) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC780, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFC93) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFC93, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7A06) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7A06, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0518) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0518, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0AC5) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0AC5, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4150) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4150, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDB6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5017
  *c.R1.BC = 0xAB81
  *c.R1.DE = 0x4287
  *c.R1.HL = 0x5EE1
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xC66F
  *c.R1.IY = 0xD6CC
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xB6
  mem[0x0002] = 0x31
  mem[0xC6A0] = 0x1C
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5C0C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5C0C, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAB81) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAB81, *c.R1.BC)
  }
  if (*c.R1.DE != 0x4287) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x4287, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5EE1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5EE1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xC66F) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xC66F, *c.R1.IX)
  }
  if (*c.R1.IY != 0xD6CC) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xD6CC, *c.R1.IY)
  }
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

func TestDDBC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x53E0
  *c.R1.BC = 0xAA98
  *c.R1.DE = 0xF7D7
  *c.R1.HL = 0xFA0C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBE7A
  *c.R1.IY = 0xA41F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xBC
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x53BF) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x53BF, *c.R1.AF)
  }
  if (*c.R1.BC != 0xAA98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xAA98, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF7D7) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF7D7, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFA0C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFA0C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBE7A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBE7A, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA41F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA41F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDBD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDC83
  *c.R1.BC = 0x80CE
  *c.R1.DE = 0x5D2F
  *c.R1.HL = 0xE999
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBB41
  *c.R1.IY = 0xA24F
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xBD
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDC82) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDC82, *c.R1.AF)
  }
  if (*c.R1.BC != 0x80CE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x80CE, *c.R1.BC)
  }
  if (*c.R1.DE != 0x5D2F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x5D2F, *c.R1.DE)
  }
  if (*c.R1.HL != 0xE999) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE999, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xBB41) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xBB41, *c.R1.IX)
  }
  if (*c.R1.IY != 0xA24F) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xA24F, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDBE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9838
  *c.R1.BC = 0xBFD5
  *c.R1.DE = 0xA299
  *c.R1.HL = 0xD34B
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x9332
  *c.R1.IY = 0xB1D5
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xBE
  mem[0x0002] = 0x48
  mem[0x937A] = 0x5B
  for c.TStates < 19 {
      c.Execute();
  }
  if (*c.R1.AF != 0x981E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x981E, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBFD5) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBFD5, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA299) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA299, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD34B) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD34B, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x9332) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x9332, *c.R1.IX)
  }
  if (*c.R1.IY != 0xB1D5) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xB1D5, *c.R1.IY)
  }
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

func TestDDE1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8A15
  *c.R1.BC = 0x6BF0
  *c.R1.DE = 0x0106
  *c.R1.HL = 0x3DD0
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x5DA4
  *c.R1.IY = 0x8716
  *c.R1.SP = 0x595F
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xE1
  mem[0x595F] = 0x9A
  mem[0x5960] = 0x09
  for c.TStates < 14 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8A15) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8A15, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6BF0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6BF0, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0106) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0106, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3DD0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3DD0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x099A) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x099A, *c.R1.IX)
  }
  if (*c.R1.IY != 0x8716) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x8716, *c.R1.IY)
  }
  if (*c.R1.SP != 0x5961) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5961, *c.R1.SP)
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

func TestDDE3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x068E
  *c.R1.BC = 0x58E6
  *c.R1.DE = 0x2713
  *c.R1.HL = 0x500F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xBE05
  *c.R1.IY = 0x4308
  *c.R1.SP = 0x57BD
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xE3
  mem[0x57BD] = 0x15
  mem[0x57BE] = 0x3F
  for c.TStates < 23 {
      c.Execute();
  }
  if (*c.R1.AF != 0x068E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x068E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x58E6) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x58E6, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2713) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2713, *c.R1.DE)
  }
  if (*c.R1.HL != 0x500F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x500F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x3F15) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x3F15, *c.R1.IX)
  }
  if (*c.R1.IY != 0x4308) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x4308, *c.R1.IY)
  }
  if (*c.R1.SP != 0x57BD) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x57BD, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 23) {
    t.Errorf("TStates mismatch: %d expected, got %d", 23, c.TStates)
  }
}

func TestDDE5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7462
  *c.R1.BC = 0x9B6C
  *c.R1.DE = 0xBFE5
  *c.R1.HL = 0x0330
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xB282
  *c.R1.IY = 0xE272
  *c.R1.SP = 0x0761
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xE5
  for c.TStates < 15 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7462) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7462, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9B6C) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9B6C, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBFE5) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBFE5, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0330) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0330, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xB282) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xB282, *c.R1.IX)
  }
  if (*c.R1.IY != 0xE272) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0xE272, *c.R1.IY)
  }
  if (*c.R1.SP != 0x075F) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x075F, *c.R1.SP)
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

func TestDDE9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x75A7
  *c.R1.BC = 0x139B
  *c.R1.DE = 0xF9A3
  *c.R1.HL = 0x94BB
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x64F0
  *c.R1.IY = 0x3433
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xE9
  for c.TStates < 8 {
      c.Execute();
  }
  if (*c.R1.AF != 0x75A7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x75A7, *c.R1.AF)
  }
  if (*c.R1.BC != 0x139B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x139B, *c.R1.BC)
  }
  if (*c.R1.DE != 0xF9A3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xF9A3, *c.R1.DE)
  }
  if (*c.R1.HL != 0x94BB) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x94BB, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x64F0) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x64F0, *c.R1.IX)
  }
  if (*c.R1.IY != 0x3433) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x3433, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x02) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x02, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 8) {
    t.Errorf("TStates mismatch: %d expected, got %d", 8, c.TStates)
  }
}

func TestDDF9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8709
  *c.R1.BC = 0x15DD
  *c.R1.DE = 0x7FA6
  *c.R1.HL = 0x3C5C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0xD3A7
  *c.R1.IY = 0x1D7B
  *c.R1.SP = 0xF67C
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xF9
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x8709) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x8709, *c.R1.AF)
  }
  if (*c.R1.BC != 0x15DD) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x15DD, *c.R1.BC)
  }
  if (*c.R1.DE != 0x7FA6) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x7FA6, *c.R1.DE)
  }
  if (*c.R1.HL != 0x3C5C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3C5C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0xD3A7) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0xD3A7, *c.R1.IX)
  }
  if (*c.R1.IY != 0x1D7B) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x1D7B, *c.R1.IY)
  }
  if (*c.R1.SP != 0xD3A7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xD3A7, *c.R1.SP)
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

func TestDDFD00(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
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
  mem[0x0000] = 0xDD
  mem[0x0001] = 0xFD
  mem[0x0002] = 0x00
  mem[0x0003] = 0x00
  for c.TStates < 16 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
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
  if (c.R != 0x04) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x04, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}