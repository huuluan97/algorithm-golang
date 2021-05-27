package some_sorts

type maxHeap struct {
	slice    []int
	heapSize int
}

func buildMaxHeap(slice []int) maxHeap {
	heap := maxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		heap.MaxHeapify(i)
	}
	return heap
}

func (h maxHeap) MaxHeapify(i int) {
	l, r := 2*i + 1, 2*i + 2
	max := i

	if l < h.size() && h.slice[l] > h.slice[max] {
		max = l
	}

	if r < h.size() && h.slice[r] > h.slice[max] {
		max = r
	}

	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}
}

func (h maxHeap) size() int { return h.heapSize } // ???

func HeapSort(slice []int) []int {
	heap := buildMaxHeap(slice)
	for i := len(heap.slice) - 1; i >= 1; i-- {
		heap.slice[0], heap.slice[i] = heap.slice[i], heap.slice[0]
		heap.heapSize--
		heap.MaxHeapify(0)
	}
	return heap.slice
}