package binary

import (
	"github.com/funny/unitest"
	"math/rand"
	"testing"
)

func ReadWriteTest(t *testing.T, n int, callback func(r *Reader, w *Writer)) {
	buffer := NewBuffer(nil)
	r := NewReader(buffer)
	w := NewWriter(buffer)
	for i := 0; i < n; i++ {
		callback(r, w)
	}
}

func Test_Uint8(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint8(rand.Intn(256))
		w.WriteUint8(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint8()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint8(t, v1, "==", v2)
	})
}

func Test_Uint16BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint16(rand.Intn(0xFFFF))
		w.WriteUint16BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint16BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint16(t, v1, "==", v2)
	})
}

func Test_Uint16LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint16(rand.Intn(0xFFFF))
		w.WriteUint16LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint16LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint16(t, v1, "==", v2)
	})
}

func Test_Uint24BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFF))
		w.WriteUint24BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint24BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint32(t, v1, "==", v2)
	})
}

func Test_Uint24LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFF))
		w.WriteUint24LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint24LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint32(t, v1, "==", v2)
	})
}

func Test_Uint32BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFFFF))
		w.WriteUint32BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint32BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint32(t, v1, "==", v2)
	})
}

func Test_Uint32LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint32(rand.Intn(0xFFFFFFFF))
		w.WriteUint32LE(v1)
		unitest.AssertNotError(t, w.Error())
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint32LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint32(t, v1, "==", v2)
	})
}

func Test_Uint40BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFF))
		w.WriteUint64BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint64BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uint40LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFF))
		w.WriteUint40LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint40LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uint48BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
		w.WriteUint48BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint48BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uint48LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
		w.WriteUint48LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint48LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uint56BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
		w.WriteUint56BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint56BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uint56LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
		w.WriteUint56LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint56LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uint64BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUint64BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint64BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uint64LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUint64LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUint64LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Uvarint(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteUvarint(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadUvarint()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertUint64(t, v1, "==", v2)
	})
}

func Test_Varint(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
		w.WriteVarint(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadVarint()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertInt64(t, v1, "==", v2)
	})
}

func Test_Float32BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := float32(rand.NormFloat64())
		w.WriteFloat32BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadFloat32BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertFloat32(t, v1, "==", v2)
	})
}

func Test_Float32LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := float32(rand.NormFloat64())
		w.WriteFloat32LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadFloat32LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertFloat32(t, v1, "==", v2)
	})
}

func Test_Float64BE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := rand.NormFloat64()
		w.WriteFloat64BE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadFloat64BE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertFloat64(t, v1, "==", v2)
	})
}

func Test_Float64LE(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		v1 := rand.NormFloat64()
		w.WriteFloat64LE(v1)
		w.Flush()
		unitest.AssertNotError(t, w.Error())

		v2 := r.ReadFloat64LE()
		unitest.AssertNotError(t, r.Error())
		unitest.AssertFloat64(t, v1, "==", v2)
	})
}
