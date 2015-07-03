package binary

import (
	"bufio"
	"bytes"
	"io"
)

type RuneWriter interface {
	WriteRune(r rune) (n int, err error)
}

type FlushWriter interface {
	Flush() error
}

type Writer struct {
	w   io.Writer
	wb  [MaxVarintLen64]byte
	err error
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func NewBufferWriter(buf []byte) *Writer {
	return &Writer{w: bytes.NewBuffer(buf)}
}

func NewBufioWriter(w io.Writer, size int) *Writer {
	return &Writer{w: bufio.NewWriterSize(w, size)}
}

func (writer *Writer) Writer() io.Writer {
	return writer.w
}

func (writer *Writer) Error() error {
	return writer.err
}

func (writer *Writer) Flush() error {
	if flusher, ok := writer.w.(FlushWriter); writer.err != nil && ok {
		writer.err = flusher.Flush()
	}
	return writer.err
}

func (writer *Writer) Write(b []byte) (n int, err error) {
	n, err = writer.w.Write(b)
	writer.err = err
	return
}

func (writer *Writer) WritePacket(b []byte, spliter Spliter) {
	if writer.err != nil {
		return
	}
	spliter.Write(writer, b)
}

func (writer *Writer) WriteByte(b byte) (err error) {
	if _, ok := writer.w.(io.ByteWriter); !ok {
		writer.w = bufio.NewWriter(writer.w)
	}
	err = writer.w.(io.ByteWriter).WriteByte(b)
	writer.err = err
	return
}

func (writer *Writer) WriteBytes(b []byte) {
	if writer.err != nil {
		return
	}
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteString(s string) {
	writer.WriteBytes([]byte(s))
}

func (writer *Writer) WriteRune(r rune) {
	if _, ok := writer.w.(RuneWriter); !ok {
		writer.w = bufio.NewWriter(writer.w)
	}
	_, writer.err = writer.w.(RuneWriter).WriteRune(r)
}

func (writer *Writer) WriteUvarint(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:]
	n := PutUvarint(b, v)
	b = b[:n]
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteVarint(v int64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:]
	n := PutVarint(b, v)
	b = b[:n]
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint8(v uint8) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:1]
	b[0] = v
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint16BE(v uint16) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:2]
	PutUint16BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint16LE(v uint16) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:2]
	PutUint16LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint24BE(v uint32) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:3]
	PutUint24BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint24LE(v uint32) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:3]
	PutUint24LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint32BE(v uint32) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:4]
	PutUint32BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint32LE(v uint32) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:4]
	PutUint32LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint40BE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:5]
	PutUint40BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint40LE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:5]
	PutUint40LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint48BE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:6]
	PutUint48BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint48LE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:6]
	PutUint48LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint56BE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:7]
	PutUint56BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint56LE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:7]
	PutUint56LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint64BE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:8]
	PutUint64BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteUint64LE(v uint64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:8]
	PutUint64LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteFloat32BE(v float32) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:4]
	PutFloat32BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteFloat32LE(v float32) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:4]
	PutFloat32LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteFloat64BE(v float64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:8]
	PutFloat64BE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteFloat64LE(v float64) {
	if writer.err != nil {
		return
	}
	b := writer.wb[:8]
	PutFloat64LE(b, v)
	_, writer.err = writer.w.Write(b)
}

func (writer *Writer) WriteInt16BE(v int16) { writer.WriteUint16BE(uint16(v)) }
func (writer *Writer) WriteInt16LE(v int16) { writer.WriteUint16LE(uint16(v)) }
func (writer *Writer) WriteInt32BE(v int32) { writer.WriteUint32BE(uint32(v)) }
func (writer *Writer) WriteInt32LE(v int32) { writer.WriteUint32LE(uint32(v)) }
func (writer *Writer) WriteInt40BE(v int64) { writer.WriteUint40BE(uint64(v)) }
func (writer *Writer) WriteInt40LE(v int64) { writer.WriteUint40LE(uint64(v)) }
func (writer *Writer) WriteInt48BE(v int64) { writer.WriteUint48BE(uint64(v)) }
func (writer *Writer) WriteInt48LE(v int64) { writer.WriteUint48LE(uint64(v)) }
func (writer *Writer) WriteInt56BE(v int64) { writer.WriteUint56BE(uint64(v)) }
func (writer *Writer) WriteInt56LE(v int64) { writer.WriteUint56LE(uint64(v)) }
func (writer *Writer) WriteInt64BE(v int64) { writer.WriteUint64BE(uint64(v)) }
func (writer *Writer) WriteInt64LE(v int64) { writer.WriteUint64LE(uint64(v)) }
