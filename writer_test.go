package binary

import (
	"testing"
)

var bmReader = NewReader(&bmBuffer)
var bmWriter = NewWriter(&bmBuffer)

func Benchmark_Reader_ReadUint16LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint16LE()
	}
}

func Benchmark_Reader_ReadUint16BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint16BE()
	}
}

func Benchmark_Reader_ReadUint24LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint24LE()
	}
}

func Benchmark_Reader_ReadUint24BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint24BE()
	}
}

func Benchmark_Reader_ReadUint32LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint32LE()
	}
}

func Benchmark_Reader_ReadUint32BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint32BE()
	}
}

func Benchmark_Reader_ReadUint40LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint40LE()
	}
}

func Benchmark_Reader_ReadUint40BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint40BE()
	}
}

func Benchmark_Reader_ReadUint48LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint48LE()
	}
}

func Benchmark_Reader_ReadUint48BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint48BE()
	}
}

func Benchmark_Reader_ReadUint56LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint56LE()
	}
}

func Benchmark_Reader_ReadUint56BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint56BE()
	}
}

func Benchmark_Reader_ReadUint64LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint64LE()
	}
}

func Benchmark_Reader_ReadUint64BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUint64BE()
	}
}

func Benchmark_Reader_ReadUvarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadUvarint()
	}
}

func Benchmark_Reader_ReadVarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmReader.ReadVarint()
	}
}

func Benchmark_Writer_WriteUint16LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint16LE(uint16(i))
	}
}

func Benchmark_Writer_WriteUint16BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint16BE(uint16(i))
	}
}

func Benchmark_Writer_WriteUint24LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint24LE(uint32(i))
	}
}

func Benchmark_Writer_WriteUint24BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint24BE(uint32(i))
	}
}

func Benchmark_Writer_WriteUint32LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint32LE(uint32(i))
	}
}

func Benchmark_Writer_WriteUint32BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint32BE(uint32(i))
	}
}

func Benchmark_Writer_WriteUint40LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint40LE(uint64(i))
	}
}

func Benchmark_Writer_WriteUint40BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint40BE(uint64(i))
	}
}

func Benchmark_Writer_WriteUint48LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint48LE(uint64(i))
	}
}

func Benchmark_Writer_WriteUint48BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint48BE(uint64(i))
	}
}

func Benchmark_Writer_WriteUint56LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint56LE(uint64(i))
	}
}

func Benchmark_Writer_WriteUint56BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint56BE(uint64(i))
	}
}

func Benchmark_Writer_WriteUint64LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint64LE(uint64(i))
	}
}

func Benchmark_Writer_WriteUint64BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUint64BE(uint64(i))
	}
}

func Benchmark_Writer_WriteUvarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteUvarint(uint64(i))
	}
}

func Benchmark_Writer_WriteVarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmWriter.WriteVarint(int64(i))
	}
}
