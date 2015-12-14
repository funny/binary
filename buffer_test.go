package binary

import (
	"math/rand"
	"testing"

	"github.com/funny/utest"
)

func (buf *Buffer) Grow(n int) {
	if n = buf.WritePos + n; n <= cap(buf.Data) {
		buf.Data = buf.Data[:n]
	} else {
		newData := make([]byte, n, n+512)
		copy(newData, buf.Data)
		buf.Data = newData
	}
}

func RandBytes(n int) []byte {
	n = rand.Intn(n) + 1
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(255))
	}
	return b
}

func Test_Buffer_ReadWrite(t *testing.T) {
	var buf Buffer
	for i := 0; i < 10000; i++ {
		b := RandBytes(256)
		buf.Grow(len(b))

		n, err := buf.Write(b)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(b))

		c := make([]byte, len(b))
		n, err = buf.Read(c)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(b))
		utest.EqualNow(t, b, c)
	}
}

func Test_Buffer_Bytes(t *testing.T) {
	var buf Buffer
	for i := 0; i < 10000; i++ {
		b := RandBytes(256)
		buf.Grow(len(b))

		buf.WriteBytes(b)

		c := buf.ReadBytes(len(b))
		utest.EqualNow(t, b, c)
	}
}

func Test_Buffer_String(t *testing.T) {
	var buf Buffer
	for i := 0; i < 10000; i++ {
		b := string(RandBytes(256))
		buf.Grow(len(b))

		buf.WriteString(b)

		c := buf.ReadString(len(b))
		utest.EqualNow(t, b, c)
	}
}

func Test_Buffer_Uint8(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint8(rand.Intn(256))
	buf.WriteUint8(v1)

	v2 := buf.ReadUint8()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint16BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint16(rand.Intn(0xFFFF))
	buf.WriteUint16BE(v1)

	v2 := buf.ReadUint16BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint16LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint16(rand.Intn(0xFFFF))
	buf.WriteUint16LE(v1)

	v2 := buf.ReadUint16LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint24BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFF))
	buf.WriteUint24BE(v1)

	v2 := buf.ReadUint24BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint24LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFF))
	buf.WriteUint24LE(v1)

	v2 := buf.ReadUint24LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint32BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFFFF))
	buf.WriteUint32BE(v1)

	v2 := buf.ReadUint32BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint32LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFFFF))
	buf.WriteUint32LE(v1)

	v2 := buf.ReadUint32LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint40BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteUint40BE(v1)

	v2 := buf.ReadUint40BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint40LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteUint40LE(v1)

	v2 := buf.ReadUint40LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint48BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteUint48BE(v1)

	v2 := buf.ReadUint48BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint48LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteUint48LE(v1)

	v2 := buf.ReadUint48LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint56BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteUint56BE(v1)

	v2 := buf.ReadUint56BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint56LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteUint56LE(v1)

	v2 := buf.ReadUint56LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint64BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUint64BE(v1)

	v2 := buf.ReadUint64BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint64LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUint64LE(v1)

	v2 := buf.ReadUint64LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uvarint(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUvarint(v1)

	v2 := buf.ReadUvarint()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Varint(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteVarint(v1)

	v2 := buf.ReadVarint()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float32BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := float32(rand.NormFloat64())
	buf.WriteFloat32BE(v1)

	v2 := buf.ReadFloat32BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float32LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := float32(rand.NormFloat64())
	buf.WriteFloat32LE(v1)

	v2 := buf.ReadFloat32LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float64BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := rand.NormFloat64()
	buf.WriteFloat64BE(v1)

	v2 := buf.ReadFloat64BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float64LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := rand.NormFloat64()
	buf.WriteFloat64LE(v1)

	v2 := buf.ReadFloat64LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int8(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int8(rand.Intn(256))
	buf.WriteInt8(v1)

	v2 := buf.ReadInt8()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int16BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int16(rand.Intn(0xFFFF))
	buf.WriteInt16BE(v1)

	v2 := buf.ReadInt16BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int16LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int16(rand.Intn(0xFFFF))
	buf.WriteInt16LE(v1)

	v2 := buf.ReadInt16LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int24BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFF))
	buf.WriteInt24BE(v1)

	v2 := buf.ReadInt24BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int24LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFF))
	buf.WriteInt24LE(v1)

	v2 := buf.ReadInt24LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int32BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFFFF))
	buf.WriteInt32BE(v1)

	v2 := buf.ReadInt32BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int32LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFFFF))
	buf.WriteInt32LE(v1)

	v2 := buf.ReadInt32LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int40BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteInt40BE(v1)

	v2 := buf.ReadInt40BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int40LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteInt40LE(v1)

	v2 := buf.ReadInt40LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int48BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteInt48BE(v1)

	v2 := buf.ReadInt48BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int48LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteInt48LE(v1)

	v2 := buf.ReadInt48LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int56BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteInt56BE(v1)

	v2 := buf.ReadInt56BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int56LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteInt56LE(v1)

	v2 := buf.ReadInt56LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int64BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteInt64BE(v1)

	v2 := buf.ReadInt64BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int64LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteInt64LE(v1)

	v2 := buf.ReadInt64LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_IntBE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteIntBE(v1)

	v2 := buf.ReadIntBE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_IntLE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteIntLE(v1)

	v2 := buf.ReadIntLE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_UintBE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUintBE(v1)

	v2 := buf.ReadUintBE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_UintLE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUintLE(v1)

	v2 := buf.ReadUintLE()
	utest.EqualNow(t, v1, v2)
}

var bmBuffer = Buffer{Data: bmBuf}

func Benchmark_Buffer_ReadUint16LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint16LE()
	}
}

func Benchmark_Buffer_ReadUint16BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint16BE()
	}
}

func Benchmark_Buffer_ReadUint24LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint24LE()
	}
}

func Benchmark_Buffer_ReadUint24BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint24BE()
	}
}

func Benchmark_Buffer_ReadUint32LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint32LE()
	}
}

func Benchmark_Buffer_ReadUint32BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint32BE()
	}
}

func Benchmark_Buffer_ReadUint40LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint40LE()
	}
}

func Benchmark_Buffer_ReadUint40BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint40BE()
	}
}

func Benchmark_Buffer_ReadUint48LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint48LE()
	}
}

func Benchmark_Buffer_ReadUint48BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint48BE()
	}
}

func Benchmark_Buffer_ReadUint56LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint56LE()
	}
}

func Benchmark_Buffer_ReadUint56BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint56BE()
	}
}

func Benchmark_Buffer_ReadUint64LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint64LE()
	}
}

func Benchmark_Buffer_ReadUint64BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUint64BE()
	}
}

func Benchmark_Buffer_ReadUvarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadUvarint()
	}
}

func Benchmark_Buffer_ReadVarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.ReadPos = 0
		bmBuffer.ReadVarint()
	}
}

func Benchmark_Buffer_WriteUint16BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint16BE(uint16(i))
	}
}
func Benchmark_Buffer_WriteUint16LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint16LE(uint16(i))
	}
}

func Benchmark_Buffer_WriteUint24BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint24BE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint24LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint24LE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint32BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint32BE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint32LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint32LE(uint32(i))
	}
}

func Benchmark_Buffer_WriteUint40BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint40BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint40LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint40LE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint48BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint48BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint48LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint48LE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint56BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint56BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint56LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint56LE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint64BE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint64BE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUint64LE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUint64LE(uint64(i))
	}
}

func Benchmark_Buffer_WriteUvarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteUvarint(uint64(i))
	}
}

func Benchmark_Buffer_WriteVarint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmBuffer.WritePos = 0
		bmBuffer.WriteVarint(int64(i))
	}
}
