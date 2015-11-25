package binary

import (
	"sync/atomic"
	"unsafe"
)

type BufferPool struct {
	classes []slabClass
	min     int
	max     int
	off     int
}

func NewBufferPool(total /* M */, min /* K */, max /* K */ int) *BufferPool {
	total *= 1024 * 1024
	num := max - min + 1
	each := total / num

	pool := &BufferPool{make([]slabClass, num), min, max, 0}

	for i, j := min, 0; i <= max; i, j = i+1, j+1 {
		size := i * 1024
		items := each / size

		class := &pool.classes[j]
		class.chunks = make([]chunk, items)
		class.page = make([]byte, items*size)

		for k := 0; k < items; k++ {
			ck := &class.chunks[k]
			ck.pool = pool
			ck.Data = class.page[k*size : k*size+size : k*size+size]
			class.Push(ck)
		}
	}

	pool.min *= 1024
	pool.max *= 1024
	pool.off = pool.min - 1023
	return pool
}

func (pool *BufferPool) NewBuffer() *Buffer {
	buf := &Buffer{
		pool:  pool,
		chunk: pool.alloc(0, pool.min),
	}
	buf.Data = buf.chunk.Data
	return buf
}

func (pool *BufferPool) Manage(buf *Buffer) {
	buf.pool = pool
	buf.chunk = pool.alloc(0, pool.min)
	buf.Data = buf.chunk.Data
}

func (pool *BufferPool) alloc(size, capacity int) *chunk {
	if capacity <= pool.max {
		if capacity < pool.min {
			capacity = pool.min
		}
		m := pool.classes[(capacity-pool.off)/1024].Pop()
		if m != nil {
			// reset its length to given size
			m.Data = m.Data[:size]
			return m
		}
	}
	return &chunk{Data: make([]byte, size, capacity)}
}

func (pool *BufferPool) free(c *chunk) {
	if cap(c.Data) >= pool.min && cap(c.Data) <= pool.max {
		pool.classes[(cap(c.Data)-pool.off)/1024].Push(c)
	}
}

type slabClass struct {
	chunks []chunk
	page   []byte
	head   unsafe.Pointer
}

func (class *slabClass) Push(item *chunk) {
	for {
		item.next = atomic.LoadPointer(&class.head)
		if atomic.CompareAndSwapPointer(&class.head, item.next, unsafe.Pointer(item)) {
			break
		}
	}
}

func (class *slabClass) Pop() *chunk {
	var ptr unsafe.Pointer
	for {
		ptr = atomic.LoadPointer(&class.head)
		if ptr == nil {
			break
		}
		if atomic.CompareAndSwapPointer(&class.head, ptr, ((*chunk)(ptr)).next) {
			break
		}
	}
	return (*chunk)(ptr)
}

type chunk struct {
	Data []byte
	pool *BufferPool
	next unsafe.Pointer
}

func (c *chunk) Free() {
	if c.pool == nil {
		return
	}
	c.pool.free(c)
}
