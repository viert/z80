package z80

/*

  Go port of Gabriel Gambetta's brilliant libz80 library.
  Original C version can be found at https://github.com/ggambetta/libz80

*/

import (
	"unsafe"
)

type RegisterSet struct {
	store [14]byte
	A     *byte
	B     *byte
	C     *byte
	D     *byte
	E     *byte
	F     *byte
	H     *byte
	L     *byte
	IXh   *byte
	IXl   *byte
	IYh   *byte
	IYl   *byte
	AF    *uint16
	BC    *uint16
	DE    *uint16
	HL    *uint16
	IX    *uint16
	IY    *uint16
	SP    *uint16
}

func NewRegisterSet() *RegisterSet {
	r := new(RegisterSet)
	r.F = &r.store[0]
	r.A = &r.store[1]
	r.C = &r.store[2]
	r.B = &r.store[3]
	r.E = &r.store[4]
	r.D = &r.store[5]
	r.L = &r.store[6]
	r.H = &r.store[7]
	r.IXl = &r.store[8]
	r.IXh = &r.store[9]
	r.IYl = &r.store[10]
	r.IYh = &r.store[11]
	r.AF = (*uint16)(unsafe.Pointer(&r.store[0]))
	r.BC = (*uint16)(unsafe.Pointer(&r.store[2]))
	r.DE = (*uint16)(unsafe.Pointer(&r.store[4]))
	r.HL = (*uint16)(unsafe.Pointer(&r.store[6]))
	r.IX = (*uint16)(unsafe.Pointer(&r.store[8]))
	r.IY = (*uint16)(unsafe.Pointer(&r.store[10]))
	r.SP = (*uint16)(unsafe.Pointer(&r.store[12]))
	return r
}
