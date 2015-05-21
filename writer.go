package binary

import (
	"bufio"
	"io"
)

type IWriter interface {
	io.Writer
	io.ByteWriter
	WriteRune(r rune) (n int, err error)
}

type IFlush interface {
	Flush() error
}

type Writer struct {
	w   IWriter
	wb  [MaxVarintLen64]byte
	err error
}

func NewWriter(w IWriter) *Writer {
	return &Writer{w: w}
}

func NewBufferWriter(buf []byte) *Writer {
	return &Writer{w: NewBuffer(buf)}
}

func NewBufioWriter(w io.Writer, size int) *Writer {
	return &Writer{w: bufio.NewWriterSize(w, size)}
}

func (writer *Writer) Backend() IWriter {
	return writer.w
}

func (writer *Writer) Error() error {
	return writer.err
}

func (writer *Writer) Flush() error {
	if flusher, ok := writer.w.(IFlush); ok {
		return flusher.Flush()
	}
	return nil
}

func (writer *Writer) Write(b []byte) (n int, err error) {
	n, err = writer.w.Write(b)
	writer.err = err
	return
}

func (writer *Writer) WriteByte(b byte) (err error) {
	err = writer.w.WriteByte(b)
	writer.err = err
	return
}

func (writer *Writer) WritePacket(b []byte, spliter Spliter) {
	if writer.err != nil {
		return
	}
	spliter.Write(writer, b)
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
	_, writer.err = writer.w.WriteRune(r)
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
	writer.err = writer.w.WriteByte(byte(v))
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
