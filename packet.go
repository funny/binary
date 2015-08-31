package binary

import (
	"io"
)

type PacketReader struct {
	Spliter Spliter
	Reader  *Reader
	Buffer  Buffer
}

func NewPacketReader(spliter Spliter, reader io.Reader) *PacketReader {
	if _, ok := reader.(*Reader); !ok {
		reader = NewReader(reader)
	}
	return &PacketReader{Spliter: spliter, Reader: reader.(*Reader)}
}

func (r *PacketReader) Read(p []byte) (int, error) {
	for {
		n, err := r.Buffer.Read(p)
		if err == nil {
			return n, nil
		}

		if err != io.EOF {
			return 0, err
		}

		r.Buffer.Reset(r.Reader.ReadPacket(r.Spliter))
		if r.Reader.Error() != nil {
			return 0, r.Reader.Error()
		}
	}
}

func (r *PacketReader) ReadPacket() ([]byte, error) {
	return r.Reader.ReadPacket(r.Spliter), r.Reader.Error()
}

type PacketWriter struct {
	Spliter Spliter
	Writer  *Writer
	Buffer  Buffer
}

func NewPacketWriter(spliter Spliter, writer io.Writer) *PacketWriter {
	if _, ok := writer.(*Writer); !ok {
		writer = NewWriter(writer)
	}
	return &PacketWriter{Spliter: spliter, Writer: writer.(*Writer)}
}

func (w *PacketWriter) Write(p []byte) (int, error) {
	return w.Buffer.Write(p)
}

func (w *PacketWriter) Flush() error {
	w.Writer.WritePacket(w.Buffer.Data, w.Spliter)
	w.Buffer.Reset(w.Buffer.Data[0:0])
	return w.Writer.Flush()
}
