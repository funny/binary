package binary

import (
	"bytes"
	"github.com/funny/utest"
	"math/rand"
	"testing"
)

func ReadWriteTest(t *testing.T, n int, callback func(r *Reader, w *Writer)) {
	var buffer bytes.Buffer
	r := NewReader(&buffer)
	w := NewWriter(&buffer)
	for i := 0; i < n; i++ {
		callback(r, w)
		buffer.Reset()
		r.Reset(&buffer)
		w.Reset(&buffer)
	}
}

func Test_ReadWrite(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b := RandBytes(256)

		n, err := w.Write(b)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(b))

		c := make([]byte, len(b))
		n, err = r.Read(c)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(b))
		utest.EqualNow(t, b, c)
	})
}

func Test_ReadWrite_Bytes(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b := RandBytes(256)

		w.WriteBytes(b)
		utest.IsNilNow(t, w.Error())

		c := r.ReadBytes(len(b))
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b, c)
	})
}

func Test_ReadWrite_String(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b := string(RandBytes(256))

		w.WriteString(b)
		utest.IsNilNow(t, w.Error())

		c := r.ReadString(len(b))
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b, c)
	})
}

func Test_ReadWrite_Uint8(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint8(rand.Intn(256))
		w.WriteUint8(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint8()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint16BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint16(rand.Intn(0xFFFF))
		w.WriteUint16BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint16BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint16LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint16(rand.Intn(0xFFFF))
		w.WriteUint16LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint16LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint24BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFF))
		w.WriteUint24BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint24BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint24LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFF))
		w.WriteUint24LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint24LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint32BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFFFF))
		w.WriteUint32BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint32BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint32LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFFFF))
		w.WriteUint32LE(v1)
		utest.IsNilNow(t, w.Error())
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint32LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint40BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFF))
		w.WriteUint40BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint40BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint40LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFF))
		w.WriteUint40LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint40LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint48BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
		w.WriteUint48BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint48BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint48LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
		w.WriteUint48LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint48LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint56BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
		w.WriteUint56BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint56BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint56LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
		w.WriteUint56LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint56LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint64BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUint64BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint64BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uint64LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUint64LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUint64LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Uvarint(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUvarint(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUvarint()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Varint(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteVarint(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadVarint()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Float32BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := float32(rand.NormFloat64())
		w.WriteFloat32BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadFloat32BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Float32LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := float32(rand.NormFloat64())
		w.WriteFloat32LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadFloat32LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Float64BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := rand.NormFloat64()
		w.WriteFloat64BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadFloat64BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Float64LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := rand.NormFloat64()
		w.WriteFloat64LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadFloat64LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int8(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int8(rand.Intn(256))
		w.WriteInt8(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt8()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int16BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int16(rand.Intn(0xFFFF))
		w.WriteInt16BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt16BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int16LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int16(rand.Intn(0xFFFF))
		w.WriteInt16LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt16LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int24BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int32(rand.Intn(0xFFFFFF))
		w.WriteInt24BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt24BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int24LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int32(rand.Intn(0xFFFFFF))
		w.WriteInt24LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt24LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int32BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int32(rand.Intn(0xFFFFFFFF))
		w.WriteInt32BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt32BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int32LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int32(rand.Intn(0xFFFFFFFF))
		w.WriteInt32LE(v1)
		utest.IsNilNow(t, w.Error())
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt32LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int40BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Intn(0xFFFFFFFFFF))
		w.WriteInt40BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt40BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int40LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Intn(0xFFFFFFFFFF))
		w.WriteInt40LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt40LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int48BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
		w.WriteInt48BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt48BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int48LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
		w.WriteInt48LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt48LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int56BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
		w.WriteInt56BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt56BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int56LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
		w.WriteInt56LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt56LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int64BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteInt64BE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt64BE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_Int64LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteInt64LE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadInt64LE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_IntBE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteIntBE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadIntBE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_IntLE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteIntLE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadIntLE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_UintBE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUintBE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUintBE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}

func Test_ReadWrite_UintLE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUintLE(v1)
		utest.IsNilNow(t, w.Error())

		v2 := r.ReadUintLE()
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, v1, v2)
	})
}
