package binary

import (
	"github.com/funny/unitest"
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
		unitest.AssertNotError(t, err)
		unitest.AssertInt(t, n, "==", len(p1))

		err = w.Flush()
		unitest.AssertNotError(t, err)

		n, err = io.ReadFull(r, p2)
		unitest.AssertNotError(t, err)
		unitest.AssertInt(t, n, "==", len(p1))
		unitest.AssertBytes(t, p1, p2)
	}
}
