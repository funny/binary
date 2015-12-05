package binary

import (
	"encoding/binary"
	"github.com/funny/utest"
	"math/rand"
	"testing"
)

func GetUintLE(b []byte, n uint) (r uint64) {
	for i := uint(0); i < n; i++ {
		r |= uint64(b[i]) << (8 * i)
	}
	return
}

func PutUintLE(b []byte, n uint, v uint64) {
	for i := uint(0); i < n; i++ {
		b[i] = byte(v >> (8 * i))
	}
}

func GetUintBE(b []byte, n uint) (r uint64) {
	for i := uint(0); i < n; i++ {
		r |= uint64(b[i]) << (8 * (n - 1 - i))
	}
	return
}

func PutUintBE(b []byte, n uint, v uint64) {
	for i := uint(0); i < n; i++ {
		b[i] = byte(v >> (8 * (n - 1 - i)))
	}
}

func Test_PutGet(t *testing.T) {
	x := make([]byte, 8)

	PutUintLE(x, 2, uint64(0xAABB))
	utest.EqualNow(t, GetUintLE(x, 2), 0xAABB)
	utest.EqualNow(t, GetUint16LE(x), 0xAABB)
	utest.EqualNow(t, binary.LittleEndian.Uint16(x), 0xAABB)

	PutUintBE(x, 2, uint64(0xAABB))
	utest.EqualNow(t, GetUintBE(x, 2), 0xAABB)
	utest.EqualNow(t, GetUint16BE(x), 0xAABB)
	utest.EqualNow(t, binary.BigEndian.Uint16(x), 0xAABB)

	PutUintLE(x, 3, uint64(0xAABBCC))
	utest.EqualNow(t, GetUintLE(x, 3), 0xAABBCC)
	utest.EqualNow(t, GetUint24LE(x), 0xAABBCC)

	PutUintBE(x, 3, uint64(0xAABBCC))
	utest.EqualNow(t, GetUintBE(x, 3), 0xAABBCC)
	utest.EqualNow(t, GetUint24BE(x), 0xAABBCC)

	PutUintLE(x, 4, uint64(0xAABBCCDD))
	utest.EqualNow(t, GetUintLE(x, 4), 0xAABBCCDD)
	utest.EqualNow(t, GetUint32LE(x), 0xAABBCCDD)
	utest.EqualNow(t, binary.LittleEndian.Uint32(x), 0xAABBCCDD)

	PutUintBE(x, 4, uint64(0xAABBCCDD))
	utest.EqualNow(t, GetUintBE(x, 4), 0xAABBCCDD)
	utest.EqualNow(t, GetUint32BE(x), 0xAABBCCDD)
	utest.EqualNow(t, binary.BigEndian.Uint32(x), 0xAABBCCDD)

	PutUintLE(x, 5, uint64(0xAABBCCDDEE))
	utest.EqualNow(t, GetUintLE(x, 5), 0xAABBCCDDEE)
	utest.EqualNow(t, GetUint40LE(x), 0xAABBCCDDEE)

	PutUintBE(x, 5, uint64(0xAABBCCDDEE))
	utest.EqualNow(t, GetUintBE(x, 5), 0xAABBCCDDEE)
	utest.EqualNow(t, GetUint40BE(x), 0xAABBCCDDEE)

	PutUintLE(x, 6, uint64(0xAABBCCDDEEFF))
	utest.EqualNow(t, GetUintLE(x, 6), 0xAABBCCDDEEFF)
	utest.EqualNow(t, GetUint48LE(x), 0xAABBCCDDEEFF)

	PutUintBE(x, 6, uint64(0xAABBCCDDEEFF))
	utest.EqualNow(t, GetUintBE(x, 6), 0xAABBCCDDEEFF)
	utest.EqualNow(t, GetUint48BE(x), 0xAABBCCDDEEFF)

	PutUintLE(x, 7, uint64(0xAABBCCDDEEFF00))
	utest.EqualNow(t, GetUintLE(x, 7), 0xAABBCCDDEEFF00)
	utest.EqualNow(t, GetUint56LE(x), 0xAABBCCDDEEFF00)

	PutUintBE(x, 7, uint64(0xAABBCCDDEEFF00))
	utest.EqualNow(t, GetUintBE(x, 7), 0xAABBCCDDEEFF00)
	utest.EqualNow(t, GetUint56BE(x), 0xAABBCCDDEEFF00)

	PutUintLE(x, 8, uint64(0xAABBCCDDEEFF0011))
	utest.EqualNow(t, GetUintLE(x, 8), uint64(0xAABBCCDDEEFF0011))
	utest.EqualNow(t, GetUint64LE(x), uint64(0xAABBCCDDEEFF0011))
	utest.EqualNow(t, binary.LittleEndian.Uint64(x), uint64(0xAABBCCDDEEFF0011))

	PutUintBE(x, 8, uint64(0xAABBCCDDEEFF0011))
	utest.EqualNow(t, GetUintBE(x, 8), uint64(0xAABBCCDDEEFF0011))
	utest.EqualNow(t, GetUint64BE(x), uint64(0xAABBCCDDEEFF0011))
	utest.EqualNow(t, binary.BigEndian.Uint64(x), uint64(0xAABBCCDDEEFF0011))
}

func Test_VarintSize(t *testing.T) {
	b := make([]byte, MaxVarintLen64)
	for i := 0; i < 100000; i++ {
		v := rand.Int63()
		size1 := VarintSize(v)
		size2 := PutVarint(b, v)
		utest.EqualNow(t, size1, size2)
	}
}

func Test_UvarintSize(t *testing.T) {
	b := make([]byte, MaxVarintLen64)
	for i := 0; i < 100000; i++ {
		v := uint64(rand.Int63())
		size1 := UvarintSize(v)
		size2 := PutUvarint(b, v)
		utest.EqualNow(t, size1, size2)
	}
}
