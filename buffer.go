package binary

import (
	"errors"
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

func (b *Buffer) Reset(data []byte) {
	b.Data = data
	b.ReadPos = 0
}

func (b *Buffer) Bytes() []byte {
	return b.Data
}

func (b *Buffer) Length() int {
	return len(b.Data)
}

// io.Reader
func (b *Buffer) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if b.ReadPos >= len(b.Data) {
		return 0, io.EOF
	}
	n := copy(p, b.Data[b.ReadPos:])
	b.ReadPos += n
	return n, nil
}

// io.ByteReader
func (b *Buffer) ReadByte() (byte, error) {
	if b.ReadPos >= len(b.Data) {
		return 0, io.EOF
	}
	r := b.Data[b.ReadPos]
	b.ReadPos += 1
	return r, nil
}

// io.ReaderAt
func (b *Buffer) ReadAt(p []byte, off int64) (int, error) {
	if off < 0 {
		return 0, errors.New("binary.Buffer.ReadAt: negative offset")
	}
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
	if b.ReadPos == len(b.Data) {
		return 0, 0, io.EOF
	}
	if c := b.Data[b.ReadPos]; c < utf8.RuneSelf {
		b.ReadPos += 1
		return rune(c), 1, nil
	}
	r, n := utf8.DecodeRune(b.Data[b.ReadPos:])
	b.ReadPos += n
	return r, n, nil
}

func (b *Buffer) ReadBytes(delim byte) ([]byte, error) {
	if b.ReadPos >= len(b.Data) {
		return nil, io.EOF
	}
	s := b.ReadPos
	for i := b.ReadPos; i < len(b.Data); i++ {
		if b.Data[i] == delim {
			b.ReadPos = i + 1
			return b.Data[s:b.ReadPos], nil
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
