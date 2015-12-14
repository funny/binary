package binary

import "io"

var _ BinaryReader = (*Buffer)(nil)
var _ BinaryReader = (*Reader)(nil)
var _ BinaryWriter = (*Buffer)(nil)
var _ BinaryWriter = (*Writer)(nil)

type BinaryReader interface {
	Error() error

	io.Reader
	io.ByteReader

	ReadBytes(n int) []byte
	ReadString(n int) string

	ReadUvarint() uint64
	ReadVarint() int64

	ReadIntBE() int
	ReadIntLE() int
	ReadUintBE() uint
	ReadUintLE() uint

	ReadInt8() int8
	ReadUint8() uint8

	ReadInt16BE() int16
	ReadInt16LE() int16
	ReadUint16BE() uint16
	ReadUint16LE() uint16

	ReadInt24BE() int32
	ReadInt24LE() int32
	ReadUint24BE() uint32
	ReadUint24LE() uint32

	ReadInt32BE() int32
	ReadInt32LE() int32
	ReadUint32BE() uint32
	ReadUint32LE() uint32

	ReadInt40BE() int64
	ReadInt40LE() int64
	ReadUint40BE() uint64
	ReadUint40LE() uint64

	ReadInt48BE() int64
	ReadInt48LE() int64
	ReadUint48BE() uint64
	ReadUint48LE() uint64

	ReadInt56BE() int64
	ReadInt56LE() int64
	ReadUint56BE() uint64
	ReadUint56LE() uint64

	ReadInt64BE() int64
	ReadInt64LE() int64
	ReadUint64BE() uint64
	ReadUint64LE() uint64

	ReadFloat32BE() float32
	ReadFloat32LE() float32
	ReadFloat64BE() float64
	ReadFloat64LE() float64
}

type BinaryWriter interface {
	Error() error

	io.Writer

	WriteBytes(b []byte)
	WriteString(s string)

	WriteUvarint(v uint64)
	WriteVarint(v int64)

	WriteIntBE(v int)
	WriteIntLE(v int)
	WriteUintBE(v uint)
	WriteUintLE(v uint)

	WriteInt8(v int8)
	WriteUint8(v uint8)

	WriteInt16BE(v int16)
	WriteInt16LE(v int16)
	WriteUint16BE(v uint16)
	WriteUint16LE(v uint16)

	WriteInt24BE(v int32)
	WriteInt24LE(v int32)
	WriteUint24BE(v uint32)
	WriteUint24LE(v uint32)

	WriteInt32BE(v int32)
	WriteInt32LE(v int32)
	WriteUint32BE(v uint32)
	WriteUint32LE(v uint32)

	WriteInt40BE(v int64)
	WriteInt40LE(v int64)
	WriteUint40BE(v uint64)
	WriteUint40LE(v uint64)

	WriteInt48BE(v int64)
	WriteInt48LE(v int64)
	WriteUint48BE(v uint64)
	WriteUint48LE(v uint64)

	WriteInt56BE(v int64)
	WriteInt56LE(v int64)
	WriteUint56BE(v uint64)
	WriteUint56LE(v uint64)

	WriteInt64BE(v int64)
	WriteInt64LE(v int64)
	WriteUint64BE(v uint64)
	WriteUint64LE(v uint64)

	WriteFloat32BE(v float32)
	WriteFloat32LE(v float32)
	WriteFloat64BE(v float64)
	WriteFloat64LE(v float64)
}
