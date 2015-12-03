package binary

import (
	"io"
	"math"
)

type Reader struct {
	R   io.Reader
	buf [MaxVarintLen64]byte
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{R: r}
}

func (reader *Reader) Reset(r io.Reader) {
	reader.R = r
	reader.err = nil
}

func (reader *Reader) Error() error {
	return reader.err
}

func (reader *Reader) Read(b []byte) (n int, err error) {
	if reader.err != nil {
		return 0, reader.err
	}
	n, err = reader.R.Read(b)
	reader.err = err
	return
}

func (reader *Reader) ReadBytes(n int) (b []byte) {
	var nn int
	b = make([]byte, n)
	nn, reader.err = io.ReadFull(reader.R, b)
	return b[:nn]
}

func (reader *Reader) ReadString(n int) string {
	return string(reader.ReadBytes(n))
}

func (reader *Reader) ReadUvarint() (v uint64) {
	if reader.err != nil {
		return
	}
	v, reader.err = ReadUvarint(reader.R.(io.ByteReader))
	return
}

func (reader *Reader) ReadVarint() (v int64) {
	if reader.err != nil {
		return
	}
	v, reader.err = ReadVarint(reader.R.(io.ByteReader))
	return
}

func (reader *Reader) seek(n int) []byte {
	if reader.err != nil {
		return nil
	}
	b := reader.buf[:n]
	_, reader.err = io.ReadFull(reader.R, b)
	return b
}

func (reader *Reader) ReadUint8() (v uint8) {
	b := reader.seek(1)
	if reader.err != nil {
		return 0
	}
	return uint8(b[0])
}

func (reader *Reader) ReadUint16BE() uint16 {
	b := reader.seek(2)
	if reader.err != nil {
		return 0
	}
	return GetUint16BE(b)
}

func (reader *Reader) ReadUint16LE() uint16 {
	b := reader.seek(2)
	if reader.err != nil {
		return 0
	}
	return GetUint16LE(b)
}

func (reader *Reader) ReadUint24BE() uint32 {
	b := reader.seek(3)
	if reader.err != nil {
		return 0
	}
	return GetUint24BE(b)
}

func (reader *Reader) ReadUint24LE() uint32 {
	b := reader.seek(3)
	if reader.err != nil {
		return 0
	}
	return GetUint24LE(b)
}

func (reader *Reader) ReadUint32BE() uint32 {
	b := reader.seek(4)
	if reader.err != nil {
		return 0
	}
	return GetUint32BE(b)
}

func (reader *Reader) ReadUint32LE() uint32 {
	b := reader.seek(4)
	if reader.err != nil {
		return 0
	}
	return GetUint32LE(b)
}

func (reader *Reader) ReadUint40BE() uint64 {
	b := reader.seek(5)
	if reader.err != nil {
		return 0
	}
	return GetUint40BE(b)
}

func (reader *Reader) ReadUint40LE() uint64 {
	b := reader.seek(5)
	if reader.err != nil {
		return 0
	}
	return GetUint40LE(b)
}

func (reader *Reader) ReadUint48BE() uint64 {
	b := reader.seek(6)
	if reader.err != nil {
		return 0
	}
	return GetUint48BE(b)
}

func (reader *Reader) ReadUint48LE() uint64 {
	b := reader.seek(6)
	if reader.err != nil {
		return 0
	}
	return GetUint48LE(b)
}

func (reader *Reader) ReadUint56BE() uint64 {
	b := reader.seek(7)
	if reader.err != nil {
		return 0
	}
	return GetUint56BE(b)
}

func (reader *Reader) ReadUint56LE() uint64 {
	b := reader.seek(7)
	if reader.err != nil {
		return 0
	}
	return GetUint56LE(b)
}

func (reader *Reader) ReadUint64BE() uint64 {
	b := reader.seek(8)
	if reader.err != nil {
		return 0
	}
	return GetUint64BE(b)
}

func (reader *Reader) ReadUint64LE() uint64 {
	b := reader.seek(8)
	if reader.err != nil {
		return 0
	}
	return GetUint64LE(b)
}

func (reader *Reader) ReadFloat32BE() float32 {
	b := reader.seek(4)
	if reader.err != nil {
		return float32(math.NaN())
	}
	return GetFloat32BE(b)
}

func (reader *Reader) ReadFloat32LE() float32 {
	b := reader.seek(4)
	if reader.err != nil {
		return float32(math.NaN())
	}
	return GetFloat32LE(b)
}

func (reader *Reader) ReadFloat64BE() float64 {
	b := reader.seek(8)
	if reader.err != nil {
		return math.NaN()
	}
	return GetFloat64BE(b)
}

func (reader *Reader) ReadFloat64LE() float64 {
	b := reader.seek(8)
	if reader.err != nil {
		return math.NaN()
	}
	return GetFloat64LE(b)
}

func (reader *Reader) ReadInt16BE() int16 { return int16(reader.ReadUint16BE()) }
func (reader *Reader) ReadInt16LE() int16 { return int16(reader.ReadUint16LE()) }
func (reader *Reader) ReadInt32BE() int32 { return int32(reader.ReadUint32BE()) }
func (reader *Reader) ReadInt32LE() int32 { return int32(reader.ReadUint32LE()) }
func (reader *Reader) ReadInt40BE() int64 { return int64(reader.ReadUint40BE()) }
func (reader *Reader) ReadInt40LE() int64 { return int64(reader.ReadUint40LE()) }
func (reader *Reader) ReadInt48BE() int64 { return int64(reader.ReadUint48BE()) }
func (reader *Reader) ReadInt48LE() int64 { return int64(reader.ReadUint48LE()) }
func (reader *Reader) ReadInt56BE() int64 { return int64(reader.ReadUint56BE()) }
func (reader *Reader) ReadInt56LE() int64 { return int64(reader.ReadUint56LE()) }
func (reader *Reader) ReadInt64BE() int64 { return int64(reader.ReadUint64BE()) }
func (reader *Reader) ReadInt64LE() int64 { return int64(reader.ReadUint64LE()) }
