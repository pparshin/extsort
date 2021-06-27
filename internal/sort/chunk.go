package sort

import (
	"container/heap"
)

type StringHeap []string

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h StringHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StringHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type Chunk struct {
	h *StringHeap
}

func NewChunk(n int) *Chunk {
	h := make(StringHeap, 0, n)
	heap.Init(&h)

	return &Chunk{
		h: &h,
	}
}

func (ch *Chunk) Add(s string) {
	heap.Push(ch.h, s)
}

func (ch *Chunk) Len() int {
	return ch.h.Len()
}

func (ch *Chunk) ToArray() []string {
	res := make([]string, 0, ch.h.Len())
	for ch.Len() > 0 {
		v := heap.Pop(ch.h).(string)
		res = append(res, v)
	}

	return res
}
