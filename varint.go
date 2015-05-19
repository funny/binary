package binary

import (
	"encoding/binary"
	"io"
)

func UvarintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}

func VarintSize(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return UvarintSize(ux)
}

func GetUvarint(b []byte) (uint64, int) {
	return binary.Uvarint(b)
}

func PutUvarint(b []byte, v uint64) int {
	return binary.PutUvarint(b, v)
}

func GetVarint(b []byte) (int64, int) {
	return binary.Varint(b)
}

func PutVarint(b []byte, v int64) int {
	return binary.PutVarint(b, v)
}

func ReadUvarint(r io.ByteReader) (uint64, error) {
	return binary.ReadUvarint(r)
}

func ReadVarint(r io.ByteReader) (int64, error) {
	return binary.ReadVarint(r)
}
