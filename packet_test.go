package binary

import (
	"github.com/funny/utest"
	"io"
	"testing"
)

func Test_PacketReaderWriter(t *testing.T) {
	buffer := NewBuffer(nil)
	r := NewPacketReader(SplitByUint16BE, buffer)
	w := NewPacketWriter(SplitByUint16BE, buffer)
	for i := 0; i < 10000; i++ {
		p1 := RandBytes(1024)
		p2 := make([]byte, len(p1))

		n, err := w.Write(p1)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(p1))

		err = w.Flush()
		utest.IsNilNow(t, err)

		n, err = io.ReadFull(r, p2)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(p1))
		utest.EqualNow(t, p1, p2)
	}
}
