package binary

import (
	"encoding/base64"
	"github.com/funny/utest"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandBytes(n int) []byte {
	n = rand.Intn(n) + 1
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(255))
	}
	return b
}

func Test_Delim_Spliter(t *testing.T) {
	ReadWriteTest(t, 1000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		b2 := make([]byte, base64.StdEncoding.EncodedLen(len(b1)))
		base64.StdEncoding.Encode(b2, b1)

		w.WritePacket(b2, SplitByLine)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b3 := r.ReadPacket(SplitByLine)
		utest.IsNilNow(t, r.Error())

		b4 := make([]byte, base64.StdEncoding.DecodedLen(len(b3)))
		n, err := base64.StdEncoding.Decode(b4, b3)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, b1, b4[:n])
	})
}

func Test_Uvarint_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUvarint)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUvarint)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint8_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(255)
		w.WritePacket(b1, SplitByUint8)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint8)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint16BE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint16BE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint16BE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint16LE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint16LE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint16LE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint24BE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint24BE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint24BE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint24LE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint24LE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint24LE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint32BE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint32BE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint32BE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint32LE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint32LE)
		utest.IsNilNow(t, w.Error())
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint32LE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint40BE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint40BE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint40BE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint40LE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint40LE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint40LE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint48BE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint48BE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint48BE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint48LE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint48LE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint48LE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint56BE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint56BE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint56BE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint56LE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint56LE)
		utest.IsNilNow(t, w.Error())
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint56LE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint64BE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint64BE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint64BE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}

func Test_Uint64LE_Spliter(t *testing.T) {
	ReadWriteTest(t, 10000, func(r *Reader, w *Writer) {
		b1 := RandBytes(1024)
		w.WritePacket(b1, SplitByUint64LE)
		w.Flush()
		utest.IsNilNow(t, w.Error())

		b2 := r.ReadPacket(SplitByUint64LE)
		utest.IsNilNow(t, r.Error())
		utest.EqualNow(t, b1, b2)
	})
}
