// +build go1.5

package binary

import (
	"bufio"
	"bytes"
	"math/rand"
	"testing"

	"github.com/funny/utest"
)

func Test_BufioReader_ReadWrite(t *testing.T) {
	var buf bytes.Buffer
	var optimizer = BufioOptimizer{
		R: bufio.NewReaderSize(&buf, 20),
	}

	for i := 0; i < 10000; i++ {
		b := RandBytes(256)
		buf.Write(b)

		r, err := optimizer.Next(len(b))
		utest.IsNilNow(t, err)

		c := make([]byte, len(b))
		n, err := r.Read(c)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(b))
		utest.EqualNow(t, b, c)
	}
}

func Test_BufioReader_Bytes(t *testing.T) {
	var buf bytes.Buffer
	var optimizer = BufioOptimizer{
		R: bufio.NewReaderSize(&buf, 20),
	}

	for i := 0; i < 10000; i++ {
		b := RandBytes(256)
		buf.Write(b)

		r, err := optimizer.Next(len(b))
		utest.IsNilNow(t, err)

		c := r.ReadBytes(len(b))
		utest.EqualNow(t, b, c)
	}
}

func Test_BufioReader_String(t *testing.T) {
	var buf bytes.Buffer
	var optimizer = BufioOptimizer{
		R: bufio.NewReaderSize(&buf, 20),
	}

	for i := 0; i < 10000; i++ {
		b := string(RandBytes(256))
		buf.WriteString(b)

		r, err := optimizer.Next(len(b))
		utest.IsNilNow(t, err)

		c := r.ReadString(len(b))
		utest.EqualNow(t, b, c)
	}
}

func BufioReaderTest(t *testing.T, f1 func(w *Writer), f2 func(r BinaryReader)) {
	var buf bytes.Buffer
	var optimizer = BufioOptimizer{
		R: bufio.NewReaderSize(&buf, 20),
	}

	var w Writer
	w.Reset(&buf)
	for i := 0; i < 10000; i++ {
		f1(&w)
	}

	r, err := optimizer.Next(buf.Len())
	utest.IsNilNow(t, err)

	for i := 0; i < 10000; i++ {
		f2(r)
	}
}

func Test_BufioReader_Uint8(t *testing.T) {
	c := make(chan uint8, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint8(rand.Intn(256))
			w.WriteUint8(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint8())
		},
	)
}

func Test_BufioReader_Uint16BE(t *testing.T) {
	c := make(chan uint16, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint16(rand.Intn(0xFFFF))
			w.WriteUint16BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint16BE())
		},
	)
}

func Test_BufioReader_Uint16LE(t *testing.T) {
	c := make(chan uint16, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint16(rand.Intn(0xFFFF))
			w.WriteUint16LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint16LE())
		},
	)
}

func Test_BufioReader_Uint24BE(t *testing.T) {
	c := make(chan uint32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint32(rand.Intn(0xFFFFFF))
			w.WriteUint24BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint24BE())
		},
	)
}

func Test_BufioReader_Uint24LE(t *testing.T) {
	c := make(chan uint32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint32(rand.Intn(0xFFFFFF))
			w.WriteUint24LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint24LE())
		},
	)
}

func Test_BufioReader_Uint32BE(t *testing.T) {
	c := make(chan uint32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint32(rand.Intn(0xFFFFFFFF))
			w.WriteUint32BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint32BE())
		},
	)
}

func Test_BufioReader_Uint32LE(t *testing.T) {
	c := make(chan uint32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint32(rand.Intn(0xFFFFFFFF))
			w.WriteUint32LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint32LE())
		},
	)
}

func Test_BufioReader_Uint40BE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Intn(0xFFFFFFFFFF))
			w.WriteUint40BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint40BE())
		},
	)
}

func Test_BufioReader_Uint40LE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Intn(0xFFFFFFFFFF))
			w.WriteUint40LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint40LE())
		},
	)
}

func Test_BufioReader_Uint48BE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
			w.WriteUint48BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint48BE())
		},
	)
}

func Test_BufioReader_Uint48LE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
			w.WriteUint48LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint48LE())
		},
	)
}

func Test_BufioReader_Uint56BE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
			w.WriteUint56BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint56BE())
		},
	)
}

func Test_BufioReader_Uint56LE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
			w.WriteUint56LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint56LE())
		},
	)
}

func Test_BufioReader_Uint64BE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteUint64BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint64BE())
		},
	)
}

func Test_BufioReader_Uint64LE(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteUint64LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUint64LE())
		},
	)
}

func Test_BufioReader_UintBE(t *testing.T) {
	c := make(chan uint, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteUintBE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUintBE())
		},
	)
}

func Test_BufioReader_UintLE(t *testing.T) {
	c := make(chan uint, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteUintLE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUintLE())
		},
	)
}

func Test_BufioReader_Uvarint(t *testing.T) {
	c := make(chan uint64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteUvarint(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadUvarint())
		},
	)
}

func Test_BufioReader_Varint(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteVarint(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadVarint())
		},
	)
}

func Test_BufioReader_Float32BE(t *testing.T) {
	c := make(chan float32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := float32(rand.NormFloat64())
			w.WriteFloat32BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadFloat32BE())
		},
	)
}

func Test_BufioReader_Float32LE(t *testing.T) {
	c := make(chan float32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := float32(rand.NormFloat64())
			w.WriteFloat32LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadFloat32LE())
		},
	)
}

func Test_BufioReader_Float64BE(t *testing.T) {
	c := make(chan float64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := rand.NormFloat64()
			w.WriteFloat64BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadFloat64BE())
		},
	)
}

func Test_BufioReader_Float64LE(t *testing.T) {
	c := make(chan float64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := rand.NormFloat64()
			w.WriteFloat64BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadFloat64BE())
		},
	)
}

func Test_BufioReader_Int8(t *testing.T) {
	c := make(chan int8, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int8(rand.Intn(256))
			w.WriteInt8(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt8())
		},
	)
}

func Test_BufioReader_Int16BE(t *testing.T) {
	c := make(chan int16, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int16(rand.Intn(0xFFFF))
			w.WriteInt16BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt16BE())
		},
	)
}

func Test_BufioReader_Int16LE(t *testing.T) {
	c := make(chan int16, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int16(rand.Intn(0xFFFF))
			w.WriteInt16LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt16LE())
		},
	)
}

func Test_BufioReader_Int24BE(t *testing.T) {
	c := make(chan int32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int32(rand.Intn(0xFFFFFF))
			w.WriteInt24BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt24BE())
		},
	)
}

func Test_BufioReader_Int24LE(t *testing.T) {
	c := make(chan int32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int32(rand.Intn(0xFFFFFF))
			w.WriteInt24LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt24LE())
		},
	)
}

func Test_BufioReader_Int32BE(t *testing.T) {
	c := make(chan int32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int32(rand.Intn(0xFFFFFFFF))
			w.WriteInt32BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt32BE())
		},
	)
}

func Test_BufioReader_Int32LE(t *testing.T) {
	c := make(chan int32, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int32(rand.Intn(0xFFFFFFFF))
			w.WriteInt32LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt32LE())
		},
	)
}

func Test_BufioReader_Int40BE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Intn(0xFFFFFFFFFF))
			w.WriteInt40BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt40BE())
		},
	)
}

func Test_BufioReader_Int40LE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Intn(0xFFFFFFFFFF))
			w.WriteInt40LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt40LE())
		},
	)
}

func Test_BufioReader_Int48BE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
			w.WriteInt48BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt48BE())
		},
	)
}

func Test_BufioReader_Int48LE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
			w.WriteInt48LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt48LE())
		},
	)
}

func Test_BufioReader_Int56BE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
			w.WriteInt56BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt56BE())
		},
	)
}

func Test_BufioReader_Int56LE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
			w.WriteInt56LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt56LE())
		},
	)
}

func Test_BufioReader_Int64BE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteInt64BE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt64BE())
		},
	)
}

func Test_BufioReader_Int64LE(t *testing.T) {
	c := make(chan int64, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteInt64LE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadInt64LE())
		},
	)
}

func Test_BufioReader_IntBE(t *testing.T) {
	c := make(chan int, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteIntBE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadIntBE())
		},
	)
}

func Test_BufioReader_IntLE(t *testing.T) {
	c := make(chan int, 10000)
	BufioReaderTest(t,
		func(w *Writer) {
			v1 := int(rand.Int63n(0x7FFFFFFFFFFFFFFF))
			w.WriteIntLE(v1)
			c <- v1
		},
		func(r BinaryReader) {
			utest.EqualNow(t, <-c, r.ReadIntLE())
		},
	)
}
