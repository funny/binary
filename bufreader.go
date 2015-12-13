package binary

import (
	"bufio"
	"io"
)

type BufReader struct {
	data    []byte
	readPos int
	reader  *bufio.Reader
	remind  int
}

func (buf *BufReader) Reset(reader *bufio.Reader, n int) (err error) {
	if buffered := reader.Buffered(); n > buffered {
		buf.remind = n - buffered
		n = buffered
	} else {
		buf.remind = 0
	}
	var data []byte
	data, err = reader.Peek(n)
	if err != nil {
		return
	}
	_, err = reader.Discard(n)
	if err != nil {
		return
	}
	buf.data = data
	buf.reader = reader
	buf.readPos = 0
	return
}

func (buf *BufReader) readForward(n int) (b []byte) {
	if buf.readPos == len(buf.data) {
		if buf.remind > cap(buf.data) {
			buf.data = make([]byte, buf.remind, buf.remind+512)
		} else {
			buf.data = buf.data[:buf.remind]
		}
		if _, err := io.ReadFull(buf.reader, buf.data); err == nil {
			buf.readPos = 0
		}
	}
	b = buf.data[buf.readPos:]
	buf.readPos += n
	return
}

func (buf *BufReader) Read(b []byte) (int, error) {
	return copy(b, buf.readForward(len(b))), nil
}

func (buf *BufReader) ReadByte() (byte, error) {
	return buf.ReadUint8(), nil
}

func (buf *BufReader) ReadBytes(n int) (b []byte) {
	b = make([]byte, n)
	copy(b, buf.readForward(n))
	return
}

func (buf *BufReader) ReadString(n int) string {
	return string(buf.readForward(n))
}

func (buf *BufReader) ReadUvarint() uint64 {
	v, err := ReadUvarint(buf)
	if err != nil {
		panic(err)
	}
	return v
}

func (buf *BufReader) ReadVarint() int64 {
	v, err := ReadVarint(buf)
	if err != nil {
		panic(err)
	}
	return v
}

func (buf *BufReader) ReadUint8() uint8 {
	return uint8(buf.readForward(1)[0])
}

func (buf *BufReader) ReadUint16BE() uint16 {
	return GetUint16BE(buf.readForward(2))
}

func (buf *BufReader) ReadUint16LE() uint16 {
	return GetUint16LE(buf.readForward(2))
}

func (buf *BufReader) ReadUint24BE() uint32 {
	return GetUint24BE(buf.readForward(3))
}

func (buf *BufReader) ReadUint24LE() uint32 {
	return GetUint24LE(buf.readForward(3))
}

func (buf *BufReader) ReadUint32BE() uint32 {
	return GetUint32BE(buf.readForward(4))
}

func (buf *BufReader) ReadUint32LE() uint32 {
	return GetUint32LE(buf.readForward(4))
}

func (buf *BufReader) ReadUint40BE() uint64 {
	return GetUint40BE(buf.readForward(5))
}

func (buf *BufReader) ReadUint40LE() uint64 {
	return GetUint40LE(buf.readForward(5))
}

func (buf *BufReader) ReadUint48BE() uint64 {
	return GetUint48BE(buf.readForward(6))
}

func (buf *BufReader) ReadUint48LE() uint64 {
	return GetUint48LE(buf.readForward(6))
}

func (buf *BufReader) ReadUint56BE() uint64 {
	return GetUint56BE(buf.readForward(7))
}

func (buf *BufReader) ReadUint56LE() uint64 {
	return GetUint56LE(buf.readForward(7))
}

func (buf *BufReader) ReadUint64BE() uint64 {
	return GetUint64BE(buf.readForward(8))
}

func (buf *BufReader) ReadUint64LE() uint64 {
	return GetUint64LE(buf.readForward(8))
}

func (buf *BufReader) ReadFloat32BE() float32 {
	return GetFloat32BE(buf.readForward(4))
}

func (buf *BufReader) ReadFloat32LE() float32 {
	return GetFloat32LE(buf.readForward(4))
}

func (buf *BufReader) ReadFloat64BE() float64 {
	return GetFloat64BE(buf.readForward(8))
}

func (buf *BufReader) ReadFloat64LE() float64 {
	return GetFloat64LE(buf.readForward(8))
}

func (buf *BufReader) ReadInt8() int8     { return int8(buf.ReadUint8()) }
func (buf *BufReader) ReadInt16BE() int16 { return int16(buf.ReadUint16BE()) }
func (buf *BufReader) ReadInt16LE() int16 { return int16(buf.ReadUint16LE()) }
func (buf *BufReader) ReadInt24BE() int32 { return int32(buf.ReadUint24BE()) }
func (buf *BufReader) ReadInt24LE() int32 { return int32(buf.ReadUint24LE()) }
func (buf *BufReader) ReadInt32BE() int32 { return int32(buf.ReadUint32BE()) }
func (buf *BufReader) ReadInt32LE() int32 { return int32(buf.ReadUint32LE()) }
func (buf *BufReader) ReadInt40BE() int64 { return int64(buf.ReadUint40BE()) }
func (buf *BufReader) ReadInt40LE() int64 { return int64(buf.ReadUint40LE()) }
func (buf *BufReader) ReadInt48BE() int64 { return int64(buf.ReadUint48BE()) }
func (buf *BufReader) ReadInt48LE() int64 { return int64(buf.ReadUint48LE()) }
func (buf *BufReader) ReadInt56BE() int64 { return int64(buf.ReadUint56BE()) }
func (buf *BufReader) ReadInt56LE() int64 { return int64(buf.ReadUint56LE()) }
func (buf *BufReader) ReadInt64BE() int64 { return int64(buf.ReadUint64BE()) }
func (buf *BufReader) ReadInt64LE() int64 { return int64(buf.ReadUint64LE()) }
func (buf *BufReader) ReadIntBE() int     { return int(buf.ReadUint64BE()) }
func (buf *BufReader) ReadIntLE() int     { return int(buf.ReadUint64LE()) }
func (buf *BufReader) ReadUintBE() uint   { return uint(buf.ReadUint64BE()) }
func (buf *BufReader) ReadUintLE() uint   { return uint(buf.ReadUint64LE()) }
