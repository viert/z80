package z80

import (
	"testing"
)

func TestDisasmLDDEFFFF(t *testing.T) {
	mem[0x0000] = 0x11
	mem[0x0001] = 0xFF
	mem[0x0002] = 0xFF
	dasm, addr := c.Disassemble(0x0000)
	expected := "LD DE,0xFFFFh"
	if addr != 3 {
		t.Errorf("%s next addr should be 3, got %d", expected, addr)
	}

	if dasm != expected {
		t.Errorf("Invalid disassembly string: %s", dasm)
	}
}

func TestDisasmINCIYd(t *testing.T) {
	mem[0x0000] = 0xFD
	mem[0x0001] = 0x34
	mem[0x0002] = 0xFC
	expected := "INC (IY-4)"
	dasm, addr := c.Disassemble(0x0000)
	if addr != 3 {
		t.Errorf("%s next addr should be 3, got %d", expected, addr)
	}

	if dasm != expected {
		t.Errorf("Invalid disassembly string: %s", dasm)
	}
}

func TestDisasmSet1IYd(t *testing.T) {
	mem[0x0000] = 0xFD
	mem[0x0001] = 0xCB
	mem[0x0002] = 0x01
	mem[0x0003] = 0xCE
	expected := "SET 1,(IY+1)"
	dasm, addr := c.Disassemble(0x0000)
	if addr != 4 {
		t.Errorf("%s next addr should be 4, got %d", expected, addr)
	}

	if dasm != expected {
		t.Errorf("Invalid disassembly string: %s", dasm)
	}
}
