// +build go1.5

package binary

import (
	"bufio"
	"io"
)

var _ BinaryReader = (*bufioReader)(nil)

type BufioOptimizer struct {
	R      *bufio.Reader
	reader bufioReader
	buffer Buffer
}

func (bo *BufioOptimizer) Next(n int) (BinaryReader, error) {
	if buffered := bo.R.Buffered(); n > buffered {
		bo.reader.remind = n - buffered
		n = buffered
	} else {
		bo.reader.remind = 0
	}

	data, err := bo.R.Peek(n)
	if err != nil {
		return nil, err
	}
	_, err = bo.R.Discard(n)
	if err != nil {
		return nil, err
	}

	if bo.reader.remind == 0 {
		bo.buffer.ReadPos = 0
		bo.buffer.Data = data
		return &bo.buffer, nil
	} else {
		bo.reader.r = bo.R
		bo.reader.readPos = 0
		bo.reader.data = data
		bo.reader.err = nil
		return &bo.reader, nil
	}
}

type bufioReader struct {
	r       *bufio.Reader
	data    []byte
	maked   []byte
	readPos int
	remind  int
	err     error
}

func (br *bufioReader) readForward(n int) (b []byte) {
	for {
		if br.err != nil {
			return zero[:n]
		}

		if br.readPos+n <= len(br.data) {
			b = br.data[br.readPos:]
			br.readPos += n
			return
		}

		dataRemind := len(br.data) - br.readPos
		remind := br.remind + dataRemind
		if remind > cap(br.maked) {
			br.maked = make([]byte, remind, remind+512)
		}
		copy(br.maked, br.data[br.readPos:])
		br.data = br.maked[:remind]
		br.remind = 0
		br.readPos = 0

		_, br.err = io.ReadFull(br.r, br.data[dataRemind:])
	}
}

func (br *bufioReader) Error() error {
	return br.err
}

func (br *bufioReader) Read(b []byte) (int, error) {
	bb := br.readForward(len(b))
	if br.err == nil {
		return copy(b, bb), nil
	}
	return 0, br.err
}

func (br *bufioReader) ReadByte() (byte, error) {
	return br.ReadUint8(), br.err
}

func (br *bufioReader) ReadBytes(n int) (b []byte) {
	b = make([]byte, n)
	bb := br.readForward(n)
	if br.err == nil {
		copy(b, bb)
	}
	return
}

func (br *bufioReader) ReadString(n int) string {
	bb := br.readForward(n)
	if br.err == nil {
		return string(bb)
	}
	return ""
}

func (br *bufioReader) ReadUvarint() uint64 {
	v, err := ReadUvarint(br)
	if err != nil {
		panic(err)
	}
	return v
}

func (br *bufioReader) ReadVarint() int64 {
	v, err := ReadVarint(br)
	if err != nil {
		panic(err)
	}
	return v
}

func (br *bufioReader) ReadUint8() uint8 {
	return uint8(br.readForward(1)[0])
}

func (br *bufioReader) ReadUint16BE() uint16 {
	return GetUint16BE(br.readForward(2))
}

func (br *bufioReader) ReadUint16LE() uint16 {
	return GetUint16LE(br.readForward(2))
}

func (br *bufioReader) ReadUint24BE() uint32 {
	return GetUint24BE(br.readForward(3))
}

func (br *bufioReader) ReadUint24LE() uint32 {
	return GetUint24LE(br.readForward(3))
}

func (br *bufioReader) ReadUint32BE() uint32 {
	return GetUint32BE(br.readForward(4))
}

func (br *bufioReader) ReadUint32LE() uint32 {
	return GetUint32LE(br.readForward(4))
}

func (br *bufioReader) ReadUint40BE() uint64 {
	return GetUint40BE(br.readForward(5))
}

func (br *bufioReader) ReadUint40LE() uint64 {
	return GetUint40LE(br.readForward(5))
}

func (br *bufioReader) ReadUint48BE() uint64 {
	return GetUint48BE(br.readForward(6))
}

func (br *bufioReader) ReadUint48LE() uint64 {
	return GetUint48LE(br.readForward(6))
}

func (br *bufioReader) ReadUint56BE() uint64 {
	return GetUint56BE(br.readForward(7))
}

func (br *bufioReader) ReadUint56LE() uint64 {
	return GetUint56LE(br.readForward(7))
}

func (br *bufioReader) ReadUint64BE() uint64 {
	return GetUint64BE(br.readForward(8))
}

func (br *bufioReader) ReadUint64LE() uint64 {
	return GetUint64LE(br.readForward(8))
}

func (br *bufioReader) ReadFloat32BE() float32 {
	return GetFloat32BE(br.readForward(4))
}

func (br *bufioReader) ReadFloat32LE() float32 {
	return GetFloat32LE(br.readForward(4))
}

func (br *bufioReader) ReadFloat64BE() float64 {
	return GetFloat64BE(br.readForward(8))
}

func (br *bufioReader) ReadFloat64LE() float64 {
	return GetFloat64LE(br.readForward(8))
}

func (br *bufioReader) ReadInt8() int8     { return int8(br.ReadUint8()) }
func (br *bufioReader) ReadInt16BE() int16 { return int16(br.ReadUint16BE()) }
func (br *bufioReader) ReadInt16LE() int16 { return int16(br.ReadUint16LE()) }
func (br *bufioReader) ReadInt24BE() int32 { return int32(br.ReadUint24BE()) }
func (br *bufioReader) ReadInt24LE() int32 { return int32(br.ReadUint24LE()) }
func (br *bufioReader) ReadInt32BE() int32 { return int32(br.ReadUint32BE()) }
func (br *bufioReader) ReadInt32LE() int32 { return int32(br.ReadUint32LE()) }
func (br *bufioReader) ReadInt40BE() int64 { return int64(br.ReadUint40BE()) }
func (br *bufioReader) ReadInt40LE() int64 { return int64(br.ReadUint40LE()) }
func (br *bufioReader) ReadInt48BE() int64 { return int64(br.ReadUint48BE()) }
func (br *bufioReader) ReadInt48LE() int64 { return int64(br.ReadUint48LE()) }
func (br *bufioReader) ReadInt56BE() int64 { return int64(br.ReadUint56BE()) }
func (br *bufioReader) ReadInt56LE() int64 { return int64(br.ReadUint56LE()) }
func (br *bufioReader) ReadInt64BE() int64 { return int64(br.ReadUint64BE()) }
func (br *bufioReader) ReadInt64LE() int64 { return int64(br.ReadUint64LE()) }
func (br *bufioReader) ReadIntBE() int     { return int(br.ReadUint64BE()) }
func (br *bufioReader) ReadIntLE() int     { return int(br.ReadUint64LE()) }
func (br *bufioReader) ReadUintBE() uint   { return uint(br.ReadUint64BE()) }
func (br *bufioReader) ReadUintLE() uint   { return uint(br.ReadUint64LE()) }
