package binary

import (
	"io"
	"unicode/utf8"
)

type Buffer struct {
	Data    []byte
	ReadPos int
}

func NewBuffer(data []byte) *Buffer {
	return &Buffer{Data: data}
}

func MakeBuffer(size, capacity int) *Buffer {
	return NewBuffer(make([]byte, size, capacity))
}

func (b *Buffer) Grows(n int) (i int) {
	i = len(b.Data)

	newLen := len(b.Data) + n
	if cap(b.Data) >= newLen {
		b.Data = b.Data[:newLen]
		return
	}

	data := make([]byte, newLen, cap(b.Data)/4+newLen)
	copy(data, b.Data)
	b.Data = data
	return
}

func (b *Buffer) Reset(size, capacity int) {
	b.ReadPos = 0
	if capacity > cap(b.Data) {
		b.Grows(capacity - len(b.Data))
	}
	b.Data = b.Data[:size]
}

func (b *Buffer) Bytes() []byte {
	return b.Data
}

func (b *Buffer) Length() int {
	return len(b.Data)
}

// io.Reader
func (b *Buffer) Read(p []byte) (int, error) {
	n, err := b.ReadAt(p, int64(b.ReadPos))
	b.ReadPos += n
	return n, err
}

// io.ByteReader
func (b *Buffer) ReadByte() (byte, error) {
	if b.ReadPos == len(b.Data) {
		return 0, io.EOF
	}
	r := b.Data[b.ReadPos]
	b.ReadPos += 1
	return r, nil
}

// io.ReaderAt
func (b *Buffer) ReadAt(p []byte, off int64) (int, error) {
	if int(off) >= len(b.Data) {
		return 0, io.EOF
	}
	n := len(p)
	if n+int(off) > len(b.Data) {
		n = len(b.Data) - int(off)
	}
	copy(p, b.Data[off:])
	return n, nil
}

// io.RuneReader
func (b *Buffer) ReadRune() (rune, int, error) {
	r, n := utf8.DecodeRune(b.Data[b.ReadPos:])
	b.ReadPos += n
	return r, n, nil
}

func (b *Buffer) ReadBytes(delim byte) ([]byte, error) {
	for i := 0; i < len(b.Data); i++ {
		if b.Data[i] == delim {
			return b.Data[:i+1], nil
		}
	}
	return nil, io.EOF
}

// io.Writer
func (b *Buffer) Write(p []byte) (int, error) {
	i := b.Grows(len(p))
	copy(b.Data[i:], p)
	return len(p), nil
}

// io.ByteWriter
func (b *Buffer) WriteByte(c byte) error {
	i := b.Grows(1)
	b.Data[i] = c
	return nil
}

func (b *Buffer) WriteRune(r rune) (int, error) {
	i := b.Grows(utf8.UTFMax)
	s := utf8.EncodeRune(b.Data[i:], r)
	n := utf8.UTFMax - s
	b.Data = b.Data[:len(b.Data)-n]
	return s, nil
}
