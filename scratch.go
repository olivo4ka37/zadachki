package main

import (
	"container/heap"
)

// MaxHeap is a max-heap of integers.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinHeap is a min-heap of integers.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MedianFinder keeps track of the median of a stream of numbers.
type MedianFinder struct {
	low  *MaxHeap // max-heap
	high *MinHeap // min-heap
}

// Constructor initializes a new MedianFinder.
func Constructor() MedianFinder {
	low := &MaxHeap{}
	high := &MinHeap{}
	heap.Init(low)
	heap.Init(high)
	return MedianFinder{
		low:  low,
		high: high,
	}
}

// AddNum adds a number to the data structure.
func (mf *MedianFinder) AddNum(num int) {
	if mf.low.Len() == 0 || num <= (*mf.low)[0] {
		heap.Push(mf.low, num)
	} else {
		heap.Push(mf.high, num)
	}

	// Balance the heaps
	if mf.low.Len() > mf.high.Len()+1 {
		heap.Push(mf.high, heap.Pop(mf.low))
	} else if mf.high.Len() > mf.low.Len() {
		heap.Push(mf.low, heap.Pop(mf.high))
	}
}

// FindMedian returns the median of all elements so far.
func (mf *MedianFinder) FindMedian() float64 {
	if mf.low.Len() > mf.high.Len() {
		return float64((*mf.low)[0])
	} else if mf.low.Len() < mf.high.Len() {
		return float64((*mf.high)[0])
	}
	return (float64((*mf.low)[0]) + float64((*mf.high)[0])) / 2.0
}
