package binary

import (
	"bytes"
	"github.com/funny/unitest"
	"testing"
)

func Test_Buffer(t *testing.T) {
	buffer := new(Buffer)

	for i := 0; i < 10000; i++ {
		prevLen := len(buffer.Data)
		p1 := RandBytes(1024)
		n, err := buffer.Write(p1)
		unitest.NotError(t, err)
		unitest.Pass(t, n == len(p1))
		unitest.Pass(t, len(buffer.Data) == prevLen+len(p1))

		p2 := make([]byte, len(p1))
		n, err = buffer.Read(p2)
		unitest.NotError(t, err)
		unitest.Pass(t, n == len(p1))
		unitest.Pass(t, bytes.Equal(p1, p2))
		unitest.Pass(t, buffer.ReadPos == prevLen+len(p1))

		prevLen = len(buffer.Data)
		p3 := RandBytes(1024)
		n, err = buffer.Write(p3)
		unitest.NotError(t, err)
		unitest.Pass(t, n == len(p3))
		unitest.Pass(t, len(buffer.Data) == prevLen+len(p3))
		for j := 0; j < len(p3); j++ {
			b, err := buffer.ReadByte()
			unitest.NotError(t, err)
			unitest.Pass(t, b == p3[j])
		}

		prevLen = len(buffer.Data)
		p4 := RandBytes(1024)
		n, err = buffer.Write(p4)
		unitest.NotError(t, err)
		unitest.Pass(t, n == len(p4))
		unitest.Pass(t, len(buffer.Data) == prevLen+len(p4))
		p5 := make([]byte, len(p4))
		for j := 0; j < len(p4); j++ {
			n, err := buffer.ReadAt(p5, int64(prevLen+j))
			unitest.NotError(t, err)
			unitest.Pass(t, n == len(p4)-j)
			unitest.Pass(t, bytes.Equal(p4[j:n+j], p5[:n]))
		}
		n, err = buffer.Read(p5)
		unitest.NotError(t, err)
		unitest.Pass(t, n == len(p4))
		unitest.Pass(t, bytes.Equal(p4, p5))
		unitest.Pass(t, buffer.ReadPos == prevLen+len(p4))
	}
}
