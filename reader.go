package binary

import (
	"bufio"
	"bytes"
	"io"
	"math"
)

type BufferedReader interface {
	io.Reader
	io.ByteReader
	io.RuneReader
	ReadBytes(delim byte) (line []byte, err error)
}

type Reader struct {
	r   BufferedReader
	buf [MaxVarintLen64]byte
	err error
}

func NewReader(r BufferedReader) *Reader {
	return &Reader{r: r}
}

func NewBufferReader(buf []byte) *Reader {
	return &Reader{r: bytes.NewBuffer(buf)}
}

func NewBufioReader(r io.Reader, size int) *Reader {
	return &Reader{r: bufio.NewReaderSize(r, size)}
}

func (reader *Reader) Reader() BufferedReader {
	return reader.r
}

func (reader *Reader) Error() error {
	return reader.err
}

func (reader *Reader) Read(b []byte) (n int, err error) {
	if reader.err != nil {
		return 0, reader.err
	}
	n, err = reader.r.Read(b)
	reader.err = err
	return
}

func (reader *Reader) ReadByte() (b byte, err error) {
	if reader.err != nil {
		return 0, reader.err
	}
	b, err = reader.r.ReadByte()
	reader.err = err
	return
}

func (reader *Reader) ReadRune() (r rune, n int, err error) {
	if reader.err != nil {
		return 0, 0, reader.err
	}
	r, n, err = reader.r.ReadRune()
	reader.err = err
	return
}

func (reader *Reader) Delimit(delim byte) (b []byte) {
	if reader.err != nil {
		return nil
	}
	b, reader.err = reader.r.ReadBytes(delim)
	return
}

func (reader *Reader) ReadPacket(spliter Spliter) (b []byte) {
	if reader.err != nil {
		return nil
	}
	b = spliter.Read(reader)
	return
}

func (reader *Reader) ReadFull(b []byte) (n int, err error) {
	if reader.err != nil {
		return 0, reader.err
	}
	n, err = io.ReadFull(reader.r, b)
	reader.err = err
	return
}

func (reader *Reader) ReadBytes(n int) (b []byte) {
	b = make([]byte, n)
	nn, _ := reader.ReadFull(b)
	return b[:nn]
}

func (reader *Reader) ReadString(n int) string {
	return string(reader.ReadBytes(n))
}

func (reader *Reader) ReadUvarint() (v uint64) {
	if reader.err != nil {
		return
	}
	v, reader.err = ReadUvarint(reader.r)
	return
}

func (reader *Reader) ReadVarint() (v int64) {
	if reader.err != nil {
		return
	}
	v, reader.err = ReadVarint(reader.r)
	return
}

func (reader *Reader) ReadUint8() (v uint8) {
	if reader.err != nil {
		return
	}
	v, reader.err = reader.r.ReadByte()
	return
}

func (reader *Reader) seek(n int) []byte {
	if reader.err != nil {
		return nil
	}
	b := reader.buf[:n]
	_, reader.err = io.ReadFull(reader.r, b)
	return b
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
