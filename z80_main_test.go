package z80

import "testing"

var (
  c *Context
  mem []byte
)

func TestMain(m *testing.M) {
  mem = make([]byte, 65536)
  c = NewContext(false)
  c.MemoryRead = func(address uint16) byte { return mem[address] }
  c.MemoryWrite = func(address uint16, value byte) { mem[address] = value }
  c.IoRead = func(address uint16) byte { return byte(address >> 8) }
  c.IoWrite = func(address uint16, value byte) {}
  m.Run()
}


func Test10(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0800
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
  mem[0x0000] = 0x00
  mem[0x0001] = 0x10
  mem[0x0002] = 0xFD
  mem[0x0003] = 0x0C
  for c.TStates < 135 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0001) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0001, *c.R1.BC)
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
  if (c.R != 0x11) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x11, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 135) {
    t.Errorf("TStates mismatch: %d expected, got %d", 135, c.TStates)
  }
}

func Test11(t *testing.T) {
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
  mem[0x0000] = 0x11
  mem[0x0001] = 0x9A
  mem[0x0002] = 0xBC
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0xBC9A) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xBC9A, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func Test12(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5600
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x8000
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
  mem[0x0000] = 0x12
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5600, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8000, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test13(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0xDEF0
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
  mem[0x0000] = 0x13
  for c.TStates < 6 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0xDEF1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xDEF1, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test14(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x2700
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
  mem[0x0000] = 0x14
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0028) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0028, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x2800) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x2800, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test15(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x1000
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
  mem[0x0000] = 0x15
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x001A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x001A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0F00) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0F00, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test16(t *testing.T) {
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
  mem[0x0000] = 0x16
  mem[0x0001] = 0x12
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x1200) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x1200, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test17(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0801
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
  mem[0x0000] = 0x17
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1100, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test18(t *testing.T) {
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
  mem[0x0000] = 0x18
  mem[0x0001] = 0x40
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func Test19(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x3456
  *c.R1.HL = 0x789A
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x19
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0028) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0028, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3456) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3456, *c.R1.DE)
  }
  if (*c.R1.HL != 0xACF0) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xACF0, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func Test21(t *testing.T) {
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
  mem[0x0000] = 0x21
  mem[0x0001] = 0x28
  mem[0x0002] = 0xED
  for c.TStates < 10 {
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
  if (*c.R1.HL != 0xED28) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xED28, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func Test22(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0xC64C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x22
  mem[0x0001] = 0xB0
  mem[0x0002] = 0xC3
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
  if (*c.R1.HL != 0xC64C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xC64C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func Test23(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x9C4E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x23
  for c.TStates < 6 {
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
  if (*c.R1.HL != 0x9C4F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9C4F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test24(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x7200
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x24
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0020) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0020, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7300) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7300, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test25(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0xA500
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x25
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00A2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00A2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA400) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA400, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test26(t *testing.T) {
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
  mem[0x0000] = 0x26
  mem[0x0001] = 0x3A
  for c.TStates < 7 {
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
  if (*c.R1.HL != 0x3A00) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x3A00, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test27(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x1F00
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
  mem[0x0000] = 0x27
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2530) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2530, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test29(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0xCDFA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x29
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0019) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0019, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9BF4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9BF4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func Test31(t *testing.T) {
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
  mem[0x0000] = 0x31
  mem[0x0001] = 0xD4
  mem[0x0002] = 0x61
  for c.TStates < 10 {
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
  if (*c.R1.SP != 0x61D4) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x61D4, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func Test32(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0E00
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
  mem[0x0000] = 0x32
  mem[0x0001] = 0xAC
  mem[0x0002] = 0xAD
  for c.TStates < 13 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0E00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0E00, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 13) {
    t.Errorf("TStates mismatch: %d expected, got %d", 13, c.TStates)
  }
}

func Test33(t *testing.T) {
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
  *c.R1.SP = 0xA55A
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x33
  for c.TStates < 6 {
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
  if (*c.R1.SP != 0xA55B) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xA55B, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test34(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0xFE1D
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x34
  mem[0xFE1D] = 0xFD
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0xFE1D) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xFE1D, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func Test35(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x470C
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x35
  mem[0x470C] = 0x82
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0082) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0082, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x470C) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x470C, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func Test36(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x7D29
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x36
  mem[0x0001] = 0x7C
  for c.TStates < 10 {
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
  if (*c.R1.HL != 0x7D29) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7D29, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func Test37(t *testing.T) {
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
  mem[0x0000] = 0x37
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0001) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0001, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test39(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x1AEF
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xC534
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x29
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0030) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0030, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x35DE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x35DE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xC534) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xC534, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func Test40(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x40
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test41(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x41
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9898) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9898, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test42(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x42
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x9098) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x9098, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test43(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x43
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xD898) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xD898, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test44(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x44
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xA198) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xA198, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test45(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x45
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x6998) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x6998, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test46(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x46
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5098) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5098, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test47(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x47
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0298) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0298, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test48(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x48
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCFCF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCFCF, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test49(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x49
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test50(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x50
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0xCFD8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xCFD8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test51(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x51
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x98D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x98D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test52(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x52
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test53(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x53
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0xD8D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xD8D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test54(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x54
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0xA1D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xA1D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test55(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x55
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x69D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x69D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test56(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x56
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x50D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x50D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test57(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x57
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x02D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x02D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test58(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x58
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90CF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90CF, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test59(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x59
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9098) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9098, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test60(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x60
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xCF69) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCF69, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test61(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x61
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9869) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9869, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test62(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x62
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x9069) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9069, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test63(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x63
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xD869) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xD869, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test64(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x64
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test65(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x65
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x6969) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x6969, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test66(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x66
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x5069) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x5069, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test67(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x67
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0269) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0269, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test68(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x68
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA1CF) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA1CF, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test69(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x69
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA198) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA198, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test70(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x70
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test71(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x71
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test72(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x72
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test73(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x73
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test74(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x74
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test75(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x75
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test76(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x76
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test77(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x77
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test78(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x78
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCF00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCF00, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test79(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x79
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9800, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test80(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x80
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0411) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0411, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test81(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x81
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3031) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3031, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test82(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x82
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1501) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1501, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test83(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x83
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0211) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0211, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test84(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x84
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD191) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD191, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test85(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x85
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9B89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9B89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test86(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x86
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E29) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E29, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test87(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x87
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEAA9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEAA9, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test88(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x88
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0411) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0411, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test89(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x89
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3031) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3031, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test90(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x90
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE6B2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE6B2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test91(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x91
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBABA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBABA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test92(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x92
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD582) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD582, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test93(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x93
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE8BA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE8BA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test94(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x94
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x191A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x191A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test95(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x95
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4F1A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4F1A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test96(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x96
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xACBA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xACBA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test97(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x97
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0042) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0042, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test98(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x98
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE6B2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE6B2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test99(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x99
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBABA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBABA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test00(t *testing.T) {
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
  mem[0x0000] = 0x00
  for c.TStates < 4 {
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test01(t *testing.T) {
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
  mem[0x0000] = 0x01
  mem[0x0001] = 0x12
  mem[0x0002] = 0x34
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x3412) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x3412, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func Test02(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x5600
  *c.R1.BC = 0x0001
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
  mem[0x0000] = 0x02
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5600, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0001) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0001, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test03(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x789A
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
  mem[0x0000] = 0x03
  for c.TStates < 6 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x789B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x789B, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test04(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0xFF00
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
  mem[0x0000] = 0x04
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0050) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0050, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test05(t *testing.T) {
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
  mem[0x0000] = 0x05
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00BA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00BA, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFF00) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFF00, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test06(t *testing.T) {
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
  mem[0x0000] = 0x06
  mem[0x0001] = 0xBC
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xBC00) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xBC00, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test07(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8800
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
  mem[0x0000] = 0x07
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1101) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1101, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test08(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xDEF0
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x1234
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x08
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1234) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1234, *c.R1.AF)
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
  if (*c.R2.AF != 0xDEF0) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0xDEF0, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test09(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x5678
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x9ABC
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x09
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0030) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0030, *c.R1.AF)
  }
  if (*c.R1.BC != 0x5678) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x5678, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0xF134) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xF134, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func Test0A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0001
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
  mem[0x0000] = 0x0A
  mem[0x0001] = 0xDE
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xDE00) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xDE00, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0001) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0001, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test0B(t *testing.T) {
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
  mem[0x0000] = 0x0B
  for c.TStates < 6 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xFFFF) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xFFFF, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test0C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x007F
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
  mem[0x0000] = 0x0C
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0094) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0094, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0080) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0080, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test0D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0080
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
  mem[0x0000] = 0x0D
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x003E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x003E, *c.R1.AF)
  }
  if (*c.R1.BC != 0x007F) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x007F, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test0E(t *testing.T) {
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
  mem[0x0000] = 0x0E
  mem[0x0001] = 0xF0
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x00F0) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x00F0, *c.R1.BC)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test0F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4100
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
  mem[0x0000] = 0x0F
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA021) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA021, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test1A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x8000
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
  mem[0x0000] = 0x1A
  mem[0x8000] = 0x13
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1300) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1300, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x8000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x8000, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test1B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0xE5D4
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
  mem[0x0000] = 0x1B
  for c.TStates < 6 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE5D3) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE5D3, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test1C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x00AA
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
  mem[0x0000] = 0x1C
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x00AB) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x00AB, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test1D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x00AA
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
  mem[0x0000] = 0x1D
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00AA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00AA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x00A9) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x00A9, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test1E(t *testing.T) {
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
  mem[0x0000] = 0x1E
  mem[0x0001] = 0xEF
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x00EF) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x00EF, *c.R1.DE)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test1F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x01C4
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
  mem[0x0000] = 0x1F
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00C5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00C5, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test20_1(t *testing.T) {
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
  mem[0x0000] = 0x20
  mem[0x0001] = 0x40
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func Test20_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0040
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
  mem[0x0000] = 0x20
  mem[0x0001] = 0x40
  for c.TStates < 7 {
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test27_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x9A02
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
  mem[0x0000] = 0x27
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3423) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3423, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test28_1(t *testing.T) {
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
  mem[0x0000] = 0x28
  mem[0x0001] = 0x8E
  for c.TStates < 7 {
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test28_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0040
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
  mem[0x0000] = 0x28
  mem[0x0001] = 0x8E
  for c.TStates < 12 {
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func Test2A(t *testing.T) {
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
  mem[0x0000] = 0x2A
  mem[0x0001] = 0x45
  mem[0x0002] = 0xAC
  mem[0xAC45] = 0xC4
  mem[0xAC46] = 0xDE
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
  if (*c.R1.HL != 0xDEC4) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDEC4, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 16) {
    t.Errorf("TStates mismatch: %d expected, got %d", 16, c.TStates)
  }
}

func Test2B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x9E66
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x2B
  for c.TStates < 6 {
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
  if (*c.R1.HL != 0x9E65) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x9E65, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test2C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0026
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x2C
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0020) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0020, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0027) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0027, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test2D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0032
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x2D
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0022) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0022, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x0000) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x0000, *c.R1.DE)
  }
  if (*c.R1.HL != 0x0031) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0031, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test2E(t *testing.T) {
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
  mem[0x0000] = 0x2E
  mem[0x0001] = 0x18
  for c.TStates < 7 {
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
  if (*c.R1.HL != 0x0018) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x0018, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test2F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x8900
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
  mem[0x0000] = 0x2F
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7632) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7632, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test30_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0036
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
  mem[0x0000] = 0x30
  mem[0x0001] = 0x50
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0036) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0036, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func Test30_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0037
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
  mem[0x0000] = 0x30
  mem[0x0001] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0037) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0037, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test37_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x00FF
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
  mem[0x0000] = 0x37
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00C5) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00C5, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test37_2(t *testing.T) {
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
  mem[0x0000] = 0x37
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFF29) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFF29, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test37_3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xFFFF
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
  mem[0x0000] = 0x37
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFFED) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFFED, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test38_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x00B2
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
  mem[0x0000] = 0x38
  mem[0x0001] = 0x66
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00B2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00B2, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test38_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x00B3
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
  mem[0x0000] = 0x38
  mem[0x0001] = 0x66
  for c.TStates < 12 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00B3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00B3, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 12) {
    t.Errorf("TStates mismatch: %d expected, got %d", 12, c.TStates)
  }
}

func Test3A(t *testing.T) {
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
  mem[0x0000] = 0x3A
  mem[0x0001] = 0x52
  mem[0x0002] = 0x99
  mem[0x9952] = 0x28
  for c.TStates < 13 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2800, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 13) {
    t.Errorf("TStates mismatch: %d expected, got %d", 13, c.TStates)
  }
}

func Test3B(t *testing.T) {
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
  *c.R1.SP = 0x9D36
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x3B
  for c.TStates < 6 {
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
  if (*c.R1.SP != 0x9D35) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x9D35, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func Test3C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCF00
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
  mem[0x0000] = 0x3C
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD090) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD090, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test3D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xEA00
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
  mem[0x0000] = 0x3D
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE9AA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE9AA, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test3E(t *testing.T) {
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
  mem[0x0000] = 0x3E
  mem[0x0001] = 0xD6
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD600) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD600, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test3F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x005B
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
  mem[0x0000] = 0x3F
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0050) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0050, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test4A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x4A
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF90) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF90, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test4B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x4B
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCFD8) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCFD8, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test4C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x4C
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCFA1) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCFA1, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test4D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x4D
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF69) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF69, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test4E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x4E
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF50) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF50, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test4F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x4F
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF02) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF02, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test5A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x5A
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9090) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9090, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test5B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x5B
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test5C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x5C
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90A1) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90A1, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test5D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x5D
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9069) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9069, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test5E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x5E
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9050) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9050, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test5F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x5F
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x9002) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x9002, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test6A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x6A
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA190) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA190, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test6B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x6B
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA1D8) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA1D8, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test6C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x6C
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA1A1) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA1A1, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test6D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x6D
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test6E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x6E
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA150) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA150, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test6F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x6F
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA102) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA102, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test7A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x7A
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test7B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x7B
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD800) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD800, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test7C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x7C
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA100, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test7D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x7D
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6900) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6900, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test7E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x7E
  mem[0xA169] = 0x50
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test7F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0200
  *c.R1.BC = 0xCF98
  *c.R1.DE = 0x90D8
  *c.R1.HL = 0xA169
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x7F
  mem[0xA169] = 0x50
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0200, *c.R1.AF)
  }
  if (*c.R1.BC != 0xCF98) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xCF98, *c.R1.BC)
  }
  if (*c.R1.DE != 0x90D8) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x90D8, *c.R1.DE)
  }
  if (*c.R1.HL != 0xA169) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xA169, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test8A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x8A
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1501) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1501, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test8B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x8B
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0211) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0211, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test8C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x8C
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD191) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD191, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test8D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x8D
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x9B89) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x9B89, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test8E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x8E
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3E29) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3E29, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test8F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x8F
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEAA9) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEAA9, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test9A(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x9A
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD582) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD582, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test9B(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x9B
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE8BA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE8BA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test9C(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x9C
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x191A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x191A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test9D(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x9D
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4F1A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4F1A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func Test9E(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x9E
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xACBA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xACBA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func Test9F(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0x9F
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0042) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0042, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA0
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0514) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0514, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA1
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3130) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3130, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA2
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2030) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2030, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA3
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0514) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0514, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA4
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD494) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD494, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA5
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA4B0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA4B0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA6
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4114) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4114, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestA7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA7
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF5B4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF5B4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA8
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFAAC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFAAC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestA9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xA9
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xCE88) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xCE88, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestAA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xAA
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xD580) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xD580, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestAB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xAB
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF8A8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF8A8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestAC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xAC
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x2928) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x2928, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestAD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xAD
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5304) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5304, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestAE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xAE
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xBCA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xBCA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestAF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xAF
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0044) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0044, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB0(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB0
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFFAC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFFAC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB1
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFFAC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFFAC, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB2
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF5A4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF5A4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB3
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFDA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFDA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB4(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB4
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFDA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFDA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB5
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF7A0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF7A0, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB6
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xFDA8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xFDA8, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestB7(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB7
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF5A4) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF5A4, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB8(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB8
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF59A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF59A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestB9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xB9
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF5BA) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF5BA, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestBA(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xBA
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF5A2) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF5A2, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestBB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xBB
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF59A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF59A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestBC(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xBC
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF51A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF51A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestBD(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xBD
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF532) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF532, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestBE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xBE
  mem[0xDCA6] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF59A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF59A, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestBF(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xF500
  *c.R1.BC = 0x0F3B
  *c.R1.DE = 0x200D
  *c.R1.HL = 0xDCA6
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xBF
  mem[0xDCA6] = 0x49
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0xF562) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xF562, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0F3B) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0F3B, *c.R1.BC)
  }
  if (*c.R1.DE != 0x200D) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x200D, *c.R1.DE)
  }
  if (*c.R1.HL != 0xDCA6) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xDCA6, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestC0_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestC0_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x00D8
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00D8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00D8, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestC1(t *testing.T) {
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
  *c.R1.SP = 0x4143
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC1
  mem[0x4143] = 0xCE
  mem[0x4144] = 0xE8
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0xE8CE) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xE8CE, *c.R1.BC)
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
  if (*c.R1.SP != 0x4145) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x4145, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestC2_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xC2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestC2_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x00C7
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
  mem[0x0000] = 0xC2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00C7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00C7, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestC3(t *testing.T) {
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
  mem[0x0000] = 0xC3
  mem[0x0001] = 0xED
  mem[0x0002] = 0x7C
  for c.TStates < 10 {
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestC4_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestC4_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x004E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x004E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x004E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestC5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x53E3
  *c.R1.BC = 0x1459
  *c.R1.DE = 0x775F
  *c.R1.HL = 0x1A2F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xEC12
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC5
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x53E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x53E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1459) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1459, *c.R1.BC)
  }
  if (*c.R1.DE != 0x775F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x775F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1A2F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1A2F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xEC10) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xEC10, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestC6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xCA00
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
  mem[0x0000] = 0xC6
  mem[0x0001] = 0x6F
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x3939) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x3939, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestC7(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xC7
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestC8_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestC8_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x00D8
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00D8) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00D8, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestC9(t *testing.T) {
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
  *c.R1.SP = 0x887E
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xC9
  mem[0x887E] = 0x36
  mem[0x887F] = 0x11
  for c.TStates < 10 {
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
  if (*c.R1.SP != 0x8880) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x8880, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestCA_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xCA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestCA_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x00C7
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
  mem[0x0000] = 0xCA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x00C7) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x00C7, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestCC_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x004E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xCC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x004E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x004E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestCC_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xCC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestCD(t *testing.T) {
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
  *c.R1.SP = 0xB07D
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xCD
  mem[0x0001] = 0x5D
  mem[0x0002] = 0x3A
  for c.TStates < 17 {
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
  if (*c.R1.SP != 0xB07B) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xB07B, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestCE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x60F5
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
  mem[0x0000] = 0xCE
  mem[0x0001] = 0xB2
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x1301) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x1301, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestCF(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xCF
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD0_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD0_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0099
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0099) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0099, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestD1(t *testing.T) {
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
  *c.R1.SP = 0x4143
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD1
  mem[0x4143] = 0xCE
  mem[0x4144] = 0xE8
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0xE8CE) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0xE8CE, *c.R1.DE)
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
  if (*c.R1.SP != 0x4145) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x4145, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestD2_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0086
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
  mem[0x0000] = 0xD2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0086) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0086, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestD2_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xD2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestD3_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA200
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
  mem[0x0000] = 0xD3
  mem[0x0001] = 0xED
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA200, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD3_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4200
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
  mem[0x0000] = 0xD3
  mem[0x0001] = 0xEC
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4200, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD3_3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4200
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
  mem[0x0000] = 0xD3
  mem[0x0001] = 0xED
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4200, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xA200
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
  mem[0x0000] = 0xD3
  mem[0x0001] = 0xEC
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA200) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA200, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD4_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestD4_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000F
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000F, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestD5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x53E3
  *c.R1.BC = 0x1459
  *c.R1.DE = 0x775F
  *c.R1.HL = 0x1A2F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xEC12
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD5
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x53E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x53E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1459) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1459, *c.R1.BC)
  }
  if (*c.R1.DE != 0x775F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x775F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1A2F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1A2F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xEC10) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xEC10, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3900
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
  mem[0x0000] = 0xD6
  mem[0x0001] = 0xDF
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x5A1B) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x5A1B, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestD7(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xD7
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD8_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestD8_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0099
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0099) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0099, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestD9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x4D94
  *c.R1.BC = 0xE07A
  *c.R1.DE = 0xE35B
  *c.R1.HL = 0x9D64
  *c.R2.AF = 0x1A64
  *c.R2.BC = 0xC930
  *c.R2.DE = 0x3D01
  *c.R2.HL = 0x7D02
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xD9
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4D94) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4D94, *c.R1.AF)
  }
  if (*c.R1.BC != 0xC930) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0xC930, *c.R1.BC)
  }
  if (*c.R1.DE != 0x3D01) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x3D01, *c.R1.DE)
  }
  if (*c.R1.HL != 0x7D02) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x7D02, *c.R1.HL)
  }
  if (*c.R2.AF != 0x1A64) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x1A64, *c.R2.AF)
  }
  if (*c.R2.BC != 0xE07A) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0xE07A, *c.R2.BC)
  }
  if (*c.R2.DE != 0xE35B) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0xE35B, *c.R2.DE)
  }
  if (*c.R2.HL != 0x9D64) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x9D64, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestDA_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xDA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestDA_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0086
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
  mem[0x0000] = 0xDA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0086) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0086, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestDB_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC100
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
  mem[0x0000] = 0xDB
  mem[0x0001] = 0xE3
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC100, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestDB_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7100
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
  mem[0x0000] = 0xDB
  mem[0x0001] = 0xE2
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7100, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestDB_3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7100
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
  mem[0x0000] = 0xDB
  mem[0x0001] = 0xE3
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x7100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x7100, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestDB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xC100
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
  mem[0x0000] = 0xDB
  mem[0x0001] = 0xE2
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0xC100) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xC100, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestDC_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000F
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000F) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000F, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestDC_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xDC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestDE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0xE78D
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
  mem[0x0000] = 0xDE
  mem[0x0001] = 0xA1
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4502) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4502, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestDF(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xDF
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestE0_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestE0_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x009C
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x009C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x009C, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestE1(t *testing.T) {
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
  *c.R1.SP = 0x4143
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE1
  mem[0x4143] = 0xCE
  mem[0x4144] = 0xE8
  for c.TStates < 10 {
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
  if (*c.R1.HL != 0xE8CE) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE8CE, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x4145) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x4145, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestE2_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0083
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
  mem[0x0000] = 0xE2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0083) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0083, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestE2_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xE2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestE3(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x4D22
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0373
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE3
  mem[0x0373] = 0x8E
  mem[0x0374] = 0xE1
  for c.TStates < 19 {
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
  if (*c.R1.HL != 0xE18E) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xE18E, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0373) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0373, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 19) {
    t.Errorf("TStates mismatch: %d expected, got %d", 19, c.TStates)
  }
}

func TestE4_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000A
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000A, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestE4_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestE5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x53E3
  *c.R1.BC = 0x1459
  *c.R1.DE = 0x775F
  *c.R1.HL = 0x1A2F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xEC12
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE5
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x53E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x53E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1459) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1459, *c.R1.BC)
  }
  if (*c.R1.DE != 0x775F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x775F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1A2F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1A2F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xEC10) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xEC10, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestE6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x7500
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
  mem[0x0000] = 0xE6
  mem[0x0001] = 0x49
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x4114) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x4114, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestE7(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xE7
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestE8_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestE8_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x009C
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x009C) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x009C, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestE9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0xCABA
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xE9
  for c.TStates < 4 {
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
  if (*c.R1.HL != 0xCABA) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCABA, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestEA_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xEA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestEA_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0083
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
  mem[0x0000] = 0xEA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0083) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0083, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestEB(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0xB879
  *c.R1.HL = 0x942E
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xEB
  for c.TStates < 4 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0000) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0000, *c.R1.AF)
  }
  if (*c.R1.BC != 0x0000) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x0000, *c.R1.BC)
  }
  if (*c.R1.DE != 0x942E) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x942E, *c.R1.DE)
  }
  if (*c.R1.HL != 0xB879) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xB879, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0x0000) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x0000, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestEC_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xEC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestEC_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000A
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xEC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000A) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000A, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestEE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x3E00
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
  mem[0x0000] = 0xEE
  mem[0x0001] = 0xD0
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xEEAC) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xEEAC, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestEF(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xEF
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestF0_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0018
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0018) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0018, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestF0_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF0
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestF1(t *testing.T) {
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
  *c.R1.SP = 0x4143
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF1
  mem[0x4143] = 0xCE
  mem[0x4144] = 0xE8
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0xE8CE) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xE8CE, *c.R1.AF)
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
  if (*c.R1.SP != 0x4145) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x4145, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestF2_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0007
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
  mem[0x0000] = 0xF2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0007) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0007, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestF2_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xF2
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestF3(t *testing.T) {
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
  mem[0x0000] = 0xF3
  for c.TStates < 4 {
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestF4_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestF4_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x008E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF4
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x008E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x008E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestF5(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x53E3
  *c.R1.BC = 0x1459
  *c.R1.DE = 0x775F
  *c.R1.HL = 0x1A2F
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0xEC12
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF5
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x53E3) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x53E3, *c.R1.AF)
  }
  if (*c.R1.BC != 0x1459) {
    t.Errorf("BC mismatch: %04X expected, got %04X", 0x1459, *c.R1.BC)
  }
  if (*c.R1.DE != 0x775F) {
    t.Errorf("DE mismatch: %04X expected, got %04X", 0x775F, *c.R1.DE)
  }
  if (*c.R1.HL != 0x1A2F) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0x1A2F, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xEC10) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xEC10, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestF6(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0600
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
  mem[0x0000] = 0xF6
  mem[0x0001] = 0xA7
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0xA7A0) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0xA7A0, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestF7(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xF7
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestF8_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0018
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 5 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0018) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0018, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F7) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F7, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 5) {
    t.Errorf("TStates mismatch: %d expected, got %d", 5, c.TStates)
  }
}

func TestF8_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0098
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x43F7
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF8
  mem[0x43F7] = 0xE9
  mem[0x43F8] = 0xAF
  for c.TStates < 11 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0098) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0098, *c.R1.AF)
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
  if (*c.R1.SP != 0x43F9) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x43F9, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}

func TestF9(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0000
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0xCE32
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x0000
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xF9
  for c.TStates < 6 {
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
  if (*c.R1.HL != 0xCE32) {
    t.Errorf("HL mismatch: %04X expected, got %04X", 0xCE32, *c.R1.HL)
  }
  if (*c.R2.AF != 0x0000) {
    t.Errorf("AF' mismatch: %04X expected, got %04X", 0x0000, *c.R2.AF)
  }
  if (*c.R2.BC != 0x0000) {
    t.Errorf("BC' mismatch: %04X expected, got %04X", 0x0000, *c.R2.BC)
  }
  if (*c.R2.DE != 0x0000) {
    t.Errorf("DE' mismatch: %04X expected, got %04X", 0x0000, *c.R2.DE)
  }
  if (*c.R2.HL != 0x0000) {
    t.Errorf("HL' mismatch: %04X expected, got %04X", 0x0000, *c.R2.HL)
  }
  if (*c.R1.IX != 0x0000) {
    t.Errorf("IX mismatch: %04X expected, got %04X", 0x0000, *c.R1.IX)
  }
  if (*c.R1.IY != 0x0000) {
    t.Errorf("IY mismatch: %04X expected, got %04X", 0x0000, *c.R1.IY)
  }
  if (*c.R1.SP != 0xCE32) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0xCE32, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 6) {
    t.Errorf("TStates mismatch: %d expected, got %d", 6, c.TStates)
  }
}

func TestFA_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0087
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
  mem[0x0000] = 0xFA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0087) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0087, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestFA_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x0007
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
  mem[0x0000] = 0xFA
  mem[0x0001] = 0x1B
  mem[0x0002] = 0xE1
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x0007) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x0007, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestFB(t *testing.T) {
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
  mem[0x0000] = 0xFB
  for c.TStates < 4 {
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 4) {
    t.Errorf("TStates mismatch: %d expected, got %d", 4, c.TStates)
  }
}

func TestFC_1(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x008E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 17 {
      c.Execute();
  }
  if (*c.R1.AF != 0x008E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x008E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5696) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5696, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 17) {
    t.Errorf("TStates mismatch: %d expected, got %d", 17, c.TStates)
  }
}

func TestFC_2(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x000E
  *c.R1.BC = 0x0000
  *c.R1.DE = 0x0000
  *c.R1.HL = 0x0000
  *c.R2.AF = 0x0000
  *c.R2.BC = 0x0000
  *c.R2.DE = 0x0000
  *c.R2.HL = 0x0000
  *c.R1.IX = 0x0000
  *c.R1.IY = 0x0000
  *c.R1.SP = 0x5698
  c.PC = 0x0000
  c.I = 0x00
  c.R = 0x00
  mem[0x0000] = 0xFC
  mem[0x0001] = 0x61
  mem[0x0002] = 0x9C
  for c.TStates < 10 {
      c.Execute();
  }
  if (*c.R1.AF != 0x000E) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x000E, *c.R1.AF)
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
  if (*c.R1.SP != 0x5698) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5698, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 10) {
    t.Errorf("TStates mismatch: %d expected, got %d", 10, c.TStates)
  }
}

func TestFE(t *testing.T) {
  c.Reset()
  *c.R1.AF = 0x6900
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
  mem[0x0000] = 0xFE
  mem[0x0001] = 0x82
  for c.TStates < 7 {
      c.Execute();
  }
  if (*c.R1.AF != 0x6987) {
    t.Errorf("AF mismatch: %04X expected, got %04X", 0x6987, *c.R1.AF)
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
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 7) {
    t.Errorf("TStates mismatch: %d expected, got %d", 7, c.TStates)
  }
}

func TestFF(t *testing.T) {
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
  *c.R1.SP = 0x5507
  c.PC = 0x6D33
  c.I = 0x00
  c.R = 0x00
  mem[0x6D33] = 0xFF
  for c.TStates < 11 {
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
  if (*c.R1.SP != 0x5505) {
    t.Errorf("SP mismatch: %04X expected, got %04X", 0x5505, *c.R1.SP)
  }
  if (c.R != 0x01) {
    t.Errorf("R mismatch: %02X expected, got %02X", 0x01, c.R)
  }
  if (c.I != 0x00) {
    t.Errorf("I mismatch: %02X expected, got %02X", 0x00, c.I)
  }
  if (c.TStates != 11) {
    t.Errorf("TStates mismatch: %d expected, got %d", 11, c.TStates)
  }
}