package TwoHeaps

import (
	"container/heap"
	"fmt"
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

// ConstructorTwoHeaps initializes a new MedianFinder.
func ConstructorTwoHeaps() MedianFinder {
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
func (this *MedianFinder) AddNum(num int) {
	if this.low.Len() == 0 || num <= (*this.low)[0] {
		heap.Push(this.low, num)
	} else {
		heap.Push(this.high, num)
	}

	// Balance the heaps
	if this.low.Len() > this.high.Len()+1 {
		heap.Push(this.high, heap.Pop(this.low))
	} else if this.high.Len() > this.low.Len() {
		heap.Push(this.low, heap.Pop(this.high))
	}
}

// FindMedian returns the median of all elements so far.
func (this *MedianFinder) FindMedian() float64 {
	if this.low.Len() > this.high.Len() {
		return float64((*this.low)[0])
	} else if this.low.Len() < this.high.Len() {
		return float64((*this.high)[0])
	}
	return (float64((*this.low)[0]) + float64((*this.high)[0])) / 2.0
}

func main() {
	finder := ConstructorTwoHeaps()
	finder.AddNum(1)
	finder.AddNum(2)
	fmt.Println(finder.FindMedian()) // Output: 1.5
	finder.AddNum(3)
	fmt.Println(finder.FindMedian()) // Output: 2
}
