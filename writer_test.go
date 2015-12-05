package binary

import (
	"bytes"
	"encoding/binary"
	"testing"
)

var bmBuf = make([]byte, MaxVarintLen64)
var bmBuffer bytes.Buffer
var bmWriter = NewWriter(&bmBuffer)
var bmBuffer2 = Buffer{Data: bmBuf}

func Benchmark_LittleEndian_PutUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint16(bmBuf, uint16(i))
	}
}

func Benchmark_PutUintLE_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintLE(bmBuf, 2, uint64(i))
	}
}

func Benchmark_PutUint16LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint16LE(bmBuf, uint16(i))
	}
}

func Benchmark_WriteUint16LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint16LE(uint16(i))
	}
}

func Benchmark_Buffer_WriteUint16LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint16LE(uint16(i))
	}
}

// =====================================================

func Benchmark_BigEndian_PutUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint16(bmBuf, uint16(i))
	}
}

func Benchmark_PutUintBE_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintBE(bmBuf, 2, uint64(i))
	}
}

func Benchmark_PutUint16BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint16BE(bmBuf, uint16(i))
	}
}

func Benchmark_WriteUint16BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint16BE(uint16(i))
	}
}

func Benchmark_Buffer_WriteUint16BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint16BE(uint16(i))
	}
}

// =====================================================

func Benchmark_PutUintLE_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintLE(bmBuf, 3, uint64(i))
	}
}

func Benchmark_PutUint24LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint24LE(bmBuf, uint32(i))
	}
}

func Benchmark_WriteUint24LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint24LE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint24LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint24LE(uint32(i))
	}
}

// =====================================================

func Benchmark_PutUintBE_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintBE(bmBuf, 3, uint64(i))
	}
}

func Benchmark_PutUint24BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint24BE(bmBuf, uint32(i))
	}
}

func Benchmark_WriteUint24BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint24BE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint24BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint24BE(uint32(i))
	}
}

// =====================================================

func Benchmark_LittleEndian_PutUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint32(bmBuf, uint32(i))
	}
}

func Benchmark_PutUintLE_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintLE(bmBuf, 4, uint64(i))
	}
}

func Benchmark_PutUint32LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint32LE(bmBuf, uint32(i))
	}
}

func Benchmark_WriteUint32LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint32LE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint32LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint32LE(uint32(i))
	}
}

// =====================================================

func Benchmark_BigEndian_PutUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(bmBuf, uint32(i))
	}
}

func Benchmark_PutUintBE_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintBE(bmBuf, 4, uint64(i))
	}
}

func Benchmark_PutUint32BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint32BE(bmBuf, uint32(i))
	}
}

func Benchmark_WriteUint32BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint32BE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint32BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint32BE(uint32(i))
	}
}

// =====================================================

func Benchmark_PutUintLE_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintLE(bmBuf, 5, uint64(i))
	}
}

func Benchmark_PutUint40LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint40LE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint40LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint40LE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint40LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint40LE(uint64(i))
	}
}

// =====================================================

func Benchmark_PutUintBE_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintBE(bmBuf, 5, uint64(i))
	}
}

func Benchmark_PutUint40BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint40BE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint40BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint40BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint40BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint40BE(uint64(i))
	}
}

// =====================================================

func Benchmark_PutUintLE_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintLE(bmBuf, 6, uint64(i))
	}
}

func Benchmark_PutUint48LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint48LE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint48LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint48LE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint48LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint48LE(uint64(i))
	}
}

// =====================================================

func Benchmark_PutUintBE_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintBE(bmBuf, 6, uint64(i))
	}
}

func Benchmark_PutUint48BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint48BE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint48BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint48BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint48BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint48BE(uint64(i))
	}
}

// =====================================================

func Benchmark_PutUintLE_7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintLE(bmBuf, 7, uint64(i))
	}
}

func Benchmark_PutUint56LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint56LE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint56LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint56LE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint56LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint56LE(uint64(i))
	}
}

// =====================================================

func Benchmark_PutUintBE_7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintBE(bmBuf, 7, uint64(i))
	}
}

func Benchmark_PutUint56BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint56BE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint56BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint56BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint56BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint56BE(uint64(i))
	}
}

// =====================================================

func Benchmark_LittleEndian_PutUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint64(bmBuf, uint64(i))
	}
}

func Benchmark_PutUintLE_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintLE(bmBuf, 8, uint64(i))
	}
}

func Benchmark_PutUint64LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint64LE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint64LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint64LE(uint64(i))
	}
}

// =====================================================

func Benchmark_BigEndian_PutUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint64(bmBuf, uint64(i))
	}
}

func Benchmark_PutUintBE_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUintBE(bmBuf, 8, uint64(i))
	}
}

func Benchmark_PutUint64BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint64BE(bmBuf, uint64(i))
	}
}

func Benchmark_WriteUint64BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.Reset()
		bmWriter.WriteUint64BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint64BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer2.WritePos = 0
		bmBuffer2.WriteUint64BE(uint64(i))
	}
}
