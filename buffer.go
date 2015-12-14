package binary

import "errors"

var ErrBufferFull = errors.New("funny/binary.Buffer: buffer full")

type Buffer struct {
	Data     []byte
	ReadPos  int
	WritePos int
}

func (buf *Buffer) Error() error {
	return nil
}

func (buf *Buffer) Read(b []byte) (int, error) {
	n := copy(b, buf.Data[buf.ReadPos:])
	buf.ReadPos += n
	return n, nil
}

func (buf *Buffer) ReadByte() (byte, error) {
	return buf.ReadUint8(), nil
}

func (buf *Buffer) ReadBytes(n int) []byte {
	b := make([]byte, n)
	copy(b, buf.Data[buf.ReadPos:buf.ReadPos+n])
	buf.ReadPos += n
	return b
}

func (buf *Buffer) ReadString(n int) string {
	s := string(buf.Data[buf.ReadPos : buf.ReadPos+n])
	buf.ReadPos += n
	return s
}

func (buf *Buffer) ReadUvarint() uint64 {
	v, n := GetUvarint(buf.Data[buf.ReadPos:])
	buf.ReadPos += n
	return v
}

func (buf *Buffer) ReadVarint() int64 {
	v, n := GetVarint(buf.Data[buf.ReadPos:])
	buf.ReadPos += n
	return v
}

func (buf *Buffer) ReadUint8() (v uint8) {
	v = uint8(buf.Data[buf.ReadPos])
	buf.ReadPos += 1
	return
}

func (buf *Buffer) ReadUint16BE() (v uint16) {
	v = GetUint16BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 2
	return
}

func (buf *Buffer) ReadUint16LE() (v uint16) {
	v = GetUint16LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 2
	return
}

func (buf *Buffer) ReadUint24BE() (v uint32) {
	v = GetUint24BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 3
	return
}

func (buf *Buffer) ReadUint24LE() (v uint32) {
	v = GetUint24LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 3
	return
}

func (buf *Buffer) ReadUint32BE() (v uint32) {
	v = GetUint32BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 4
	return
}

func (buf *Buffer) ReadUint32LE() (v uint32) {
	v = GetUint32LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 4
	return
}

func (buf *Buffer) ReadUint40BE() (v uint64) {
	v = GetUint40BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 5
	return
}

func (buf *Buffer) ReadUint40LE() (v uint64) {
	v = GetUint40LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 5
	return
}

func (buf *Buffer) ReadUint48BE() (v uint64) {
	v = GetUint48BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 6
	return
}

func (buf *Buffer) ReadUint48LE() (v uint64) {
	v = GetUint48LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 6
	return
}

func (buf *Buffer) ReadUint56BE() (v uint64) {
	v = GetUint56BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 7
	return
}

func (buf *Buffer) ReadUint56LE() (v uint64) {
	v = GetUint56LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 7
	return
}

func (buf *Buffer) ReadUint64BE() (v uint64) {
	v = GetUint64BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 8
	return
}

func (buf *Buffer) ReadUint64LE() (v uint64) {
	v = GetUint64LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 8
	return
}

func (buf *Buffer) ReadFloat32BE() (v float32) {
	v = GetFloat32BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 4
	return
}

func (buf *Buffer) ReadFloat32LE() (v float32) {
	v = GetFloat32LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 4
	return
}

func (buf *Buffer) ReadFloat64BE() (v float64) {
	v = GetFloat64BE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 8
	return
}

func (buf *Buffer) ReadFloat64LE() (v float64) {
	v = GetFloat64LE(buf.Data[buf.ReadPos:])
	buf.ReadPos += 8
	return
}

func (buf *Buffer) ReadInt8() int8     { return int8(buf.ReadUint8()) }
func (buf *Buffer) ReadInt16BE() int16 { return int16(buf.ReadUint16BE()) }
func (buf *Buffer) ReadInt16LE() int16 { return int16(buf.ReadUint16LE()) }
func (buf *Buffer) ReadInt24BE() int32 { return int32(buf.ReadUint24BE()) }
func (buf *Buffer) ReadInt24LE() int32 { return int32(buf.ReadUint24LE()) }
func (buf *Buffer) ReadInt32BE() int32 { return int32(buf.ReadUint32BE()) }
func (buf *Buffer) ReadInt32LE() int32 { return int32(buf.ReadUint32LE()) }
func (buf *Buffer) ReadInt40BE() int64 { return int64(buf.ReadUint40BE()) }
func (buf *Buffer) ReadInt40LE() int64 { return int64(buf.ReadUint40LE()) }
func (buf *Buffer) ReadInt48BE() int64 { return int64(buf.ReadUint48BE()) }
func (buf *Buffer) ReadInt48LE() int64 { return int64(buf.ReadUint48LE()) }
func (buf *Buffer) ReadInt56BE() int64 { return int64(buf.ReadUint56BE()) }
func (buf *Buffer) ReadInt56LE() int64 { return int64(buf.ReadUint56LE()) }
func (buf *Buffer) ReadInt64BE() int64 { return int64(buf.ReadUint64BE()) }
func (buf *Buffer) ReadInt64LE() int64 { return int64(buf.ReadUint64LE()) }
func (buf *Buffer) ReadIntBE() int     { return int(buf.ReadUint64BE()) }
func (buf *Buffer) ReadIntLE() int     { return int(buf.ReadUint64LE()) }
func (buf *Buffer) ReadUintBE() uint   { return uint(buf.ReadUint64BE()) }
func (buf *Buffer) ReadUintLE() uint   { return uint(buf.ReadUint64LE()) }

func (buf *Buffer) Take(n int) (data []byte) {
	data = buf.Data[buf.WritePos : buf.WritePos+n]
	buf.WritePos += n
	return
}

func (buf *Buffer) Write(b []byte) (int, error) {
	n := copy(buf.Data[buf.WritePos:], b)
	buf.WritePos += n
	if n != len(b) {
		return n, ErrBufferFull
	}
	return n, nil
}

func (buf *Buffer) WriteBytes(b []byte) {
	n := copy(buf.Data[buf.WritePos:], b)
	buf.WritePos += n
	if n != len(b) {
		panic(ErrBufferFull)
	}
}

func (buf *Buffer) WriteString(s string) {
	n := copy(buf.Data[buf.WritePos:], s)
	buf.WritePos += n
	if n != len(s) {
		panic(ErrBufferFull)
	}
}

func (buf *Buffer) WriteUvarint(v uint64) {
	buf.WritePos += PutUvarint(buf.Data[buf.WritePos:], v)
}

func (buf *Buffer) WriteVarint(v int64) {
	buf.WritePos += PutVarint(buf.Data[buf.WritePos:], v)
}

func (buf *Buffer) WriteUint8(v uint8) {
	buf.Data[buf.WritePos] = byte(v)
	buf.WritePos += 1
}

func (buf *Buffer) WriteUint16BE(v uint16) {
	PutUint16BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 2
}

func (buf *Buffer) WriteUint16LE(v uint16) {
	PutUint16LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 2
}
func (buf *Buffer) WriteUint24BE(v uint32) {
	PutUint24BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 3
}

func (buf *Buffer) WriteUint24LE(v uint32) {
	PutUint24LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 3
}

func (buf *Buffer) WriteUint32BE(v uint32) {
	PutUint32BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 4
}

func (buf *Buffer) WriteUint32LE(v uint32) {
	PutUint32LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 4
}

func (buf *Buffer) WriteUint40BE(v uint64) {
	PutUint40BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 5
}

func (buf *Buffer) WriteUint40LE(v uint64) {
	PutUint40LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 5
}

func (buf *Buffer) WriteUint48BE(v uint64) {
	PutUint48BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 6
}

func (buf *Buffer) WriteUint48LE(v uint64) {
	PutUint48LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 6
}

func (buf *Buffer) WriteUint56BE(v uint64) {
	PutUint56BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 7
}

func (buf *Buffer) WriteUint56LE(v uint64) {
	PutUint56LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 7
}

func (buf *Buffer) WriteUint64BE(v uint64) {
	PutUint64BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 8
}

func (buf *Buffer) WriteUint64LE(v uint64) {
	PutUint64LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 8
}

func (buf *Buffer) WriteFloat32BE(v float32) {
	PutFloat32BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 4
}

func (buf *Buffer) WriteFloat32LE(v float32) {
	PutFloat32LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 4
}

func (buf *Buffer) WriteFloat64BE(v float64) {
	PutFloat64BE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 8
}

func (buf *Buffer) WriteFloat64LE(v float64) {
	PutFloat64LE(buf.Data[buf.WritePos:], v)
	buf.WritePos += 8
}

func (buf *Buffer) WriteInt8(v int8)     { buf.WriteUint8(uint8(v)) }
func (buf *Buffer) WriteInt16BE(v int16) { buf.WriteUint16BE(uint16(v)) }
func (buf *Buffer) WriteInt16LE(v int16) { buf.WriteUint16LE(uint16(v)) }
func (buf *Buffer) WriteInt24BE(v int32) { buf.WriteUint24BE(uint32(v)) }
func (buf *Buffer) WriteInt24LE(v int32) { buf.WriteUint24LE(uint32(v)) }
func (buf *Buffer) WriteInt32BE(v int32) { buf.WriteUint32BE(uint32(v)) }
func (buf *Buffer) WriteInt32LE(v int32) { buf.WriteUint32LE(uint32(v)) }
func (buf *Buffer) WriteInt40BE(v int64) { buf.WriteUint40BE(uint64(v)) }
func (buf *Buffer) WriteInt40LE(v int64) { buf.WriteUint40LE(uint64(v)) }
func (buf *Buffer) WriteInt48BE(v int64) { buf.WriteUint48BE(uint64(v)) }
func (buf *Buffer) WriteInt48LE(v int64) { buf.WriteUint48LE(uint64(v)) }
func (buf *Buffer) WriteInt56BE(v int64) { buf.WriteUint56BE(uint64(v)) }
func (buf *Buffer) WriteInt56LE(v int64) { buf.WriteUint56LE(uint64(v)) }
func (buf *Buffer) WriteInt64BE(v int64) { buf.WriteUint64BE(uint64(v)) }
func (buf *Buffer) WriteInt64LE(v int64) { buf.WriteUint64LE(uint64(v)) }
func (buf *Buffer) WriteIntBE(v int)     { buf.WriteUint64BE(uint64(v)) }
func (buf *Buffer) WriteIntLE(v int)     { buf.WriteUint64LE(uint64(v)) }
func (buf *Buffer) WriteUintBE(v uint)   { buf.WriteUint64BE(uint64(v)) }
func (buf *Buffer) WriteUintLE(v uint)   { buf.WriteUint64LE(uint64(v)) }
