package main

import "fmt"

func main() {
	m := &maxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}

// MaxHeap structhas a slice that holds the array.

type maxHeap struct {
	slice []int
}

// Insert adds an element to the heap.

func (h *maxHeap) insert(key int) {
	h.slice = append(h.slice, key)
	h.maxHeapifyUp(len(h.slice) - 1)
}

// Extract returns the largest key, and removes it from the heap.

func (h *maxHeap) Extract() int {

	extracted := h.slice[0]
	l := len(h.slice) - 1

	// when slice is empty returns -1
	if len(h.slice) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	// take out the last index and put it in the root
	h.slice[0] = h.slice[l]
	h.slice = h.slice[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom top

func (h *maxHeap) maxHeapifyUp(index int) {
	for h.slice[parent(index)] < h.slice[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from bottom top

func (h *maxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.slice) - 1
	l, r := leftchild(index), rightchild(index)
	childToCompare := 0

	//loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.slice[l] >= h.slice[r] { // when left child is larger or equal
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.slice[index] < h.slice[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftchild(index), rightchild(index)
		} else {
			return
		}
	}
}

// get the parent index

func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index

func leftchild(i int) int {
	return (i * 2) + 1
}

// get the right child index

func rightchild(i int) int {
	return (i * 2) + 2
}

//swap keys in the slice

func (h *maxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]
}
