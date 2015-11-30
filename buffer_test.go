package binary

import (
	"github.com/funny/unitest"
	"math/rand"
	"testing"
	"time"
)

func BufferTest(t *testing.T, buffer *Buffer) {
	for i := 0; i < 10000; i++ {
		prevLen := len(buffer.Data)
		p1 := RandBytes(1024)
		n, err := buffer.Write(p1)
		unitest.AssertNotError(t, err)
		unitest.AssertInt(t, n, "==", len(p1))
		unitest.AssertInt(t, len(buffer.Data), "==", prevLen+len(p1))

		p2 := make([]byte, len(p1))
		n, err = buffer.Read(p2)
		unitest.AssertNotError(t, err)
		unitest.AssertInt(t, n, "==", len(p1))
		unitest.AssertBytes(t, p1, p2)
		unitest.AssertInt(t, buffer.ReadPos, "==", prevLen+len(p1))

		prevLen = len(buffer.Data)
		p3 := RandBytes(1024)
		n, err = buffer.Write(p3)
		unitest.AssertNotError(t, err)
		unitest.AssertInt(t, n, "==", len(p3))
		unitest.AssertInt(t, len(buffer.Data), "==", prevLen+len(p3))
		for j := 0; j < len(p3); j++ {
			b, err := buffer.ReadByte()
			unitest.AssertNotError(t, err)
			unitest.AssertByte(t, b, "==", p3[j])
		}

		prevLen = len(buffer.Data)
		p4 := RandBytes(1024)
		n, err = buffer.Write(p4)
		unitest.AssertNotError(t, err)
		unitest.AssertInt(t, n, "==", len(p4))
		unitest.AssertInt(t, len(buffer.Data), "==", prevLen+len(p4))
		p5 := make([]byte, len(p4))
		for j := 0; j < len(p4); j++ {
			n, err := buffer.ReadAt(p5, int64(prevLen+j))
			unitest.AssertNotError(t, err)
			unitest.AssertInt(t, n, "==", len(p4)-j)
			unitest.AssertBytes(t, p4[j:n+j], p5[:n])
		}
		n, err = buffer.Read(p5)
		unitest.AssertNotError(t, err)
		unitest.AssertInt(t, n, "==", len(p4))
		unitest.AssertBytes(t, p4, p5)
		unitest.AssertInt(t, buffer.ReadPos, "==", prevLen+len(p4))
	}
}

func Test_Buffer(t *testing.T) {
	buffer := new(Buffer)
	BufferTest(t, buffer)
}

func Test_PoolBuffer(t *testing.T) {
	pool := NewBufferPool(1, 1, 32)
	buffer := pool.NewBuffer()
	BufferTest(t, buffer)
}

func Test_MemPool(t *testing.T) {
	MemPoolTest(t, 1, 1, 10)
	MemPoolTest(t, 1, 2, 16)
}

func MemPoolTest(t *testing.T, total, min, max int) {
	pool := NewBufferPool(total, min, max)

	for i := min; i <= max; i++ {
		min1 := (i-1)*1024 + 1
		max1 := i * 1024

		b1 := pool.alloc(min1, min1)
		b2 := pool.alloc(max1, max1)

		unitest.AssertInt(t, cap(b1.Data), "==", max1)
		unitest.AssertInt(t, cap(b2.Data), "==", max1)

		b1.Free()
		b2.Free()
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100000; i++ {
		size := (rand.Intn(max-min+1) + min) * 1024
		b := pool.alloc(size, size)
		b.Free()
	}
	/*
		for _, class := range pool.classes {
			unitest.Pass(t, class.length == class.maxlen)
		}
	*/
}

func Benchmark_MemPool_AllocFree(b *testing.B) {
	pool := NewBufferPool(20, 2, 32)
	size := (rand.Intn(30) + 2) * 1024
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := pool.alloc(0, size)
		x.Free()
	}
}

func Benchmark_AllocFree(b *testing.B) {
	size := (rand.Intn(30) + 2) * 1024

	for i := 0; i < b.N; i++ {
		_ = make([]byte, size)
	}
}

func Benchmark_MemPool_Select(b *testing.B) {
	c := make(chan int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case c <- i:
		default:
		}
		select {
		case _ = <-c:
		default:
		}
	}
}

func Benchmark_MemPool_Make512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, 512)
	}
}

func Benchmark_MemPool_Make1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, 1024)
	}
}

func Benchmark_MemPool_Make4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, 4096)
	}
}

func Benchmark_MemPool_Make8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, 8192)
	}
}
