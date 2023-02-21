package heap

type comparefunc[T comparable] func(a, b T) bool
type Heap[T comparable] struct {
	data    []T
	compare comparefunc[T]
	length  int
}

func parentIdx(i int) int {
	return (i - 1) >> 1
}
func leftIdx(i int) int {
	return (i << 1) + 1
}
func rightIdx(i int) int {
	return (i << 1) + 2
}

// 将当前最小子节点移到根, 若找到最小则随着最小子节点向下递归
func (h *Heap[T]) percolateDown(rootIdx int) {
	selectedIdx := rootIdx
	if left := leftIdx(rootIdx); left < h.length && h.compare(h.data[left], h.data[selectedIdx]) {
		selectedIdx = left
	}
	if right := rightIdx(rootIdx); right < h.length && h.compare(h.data[right], h.data[selectedIdx]) {
		selectedIdx = right
	}
	if selectedIdx != rootIdx {
		h.data[rootIdx], h.data[selectedIdx] = h.data[selectedIdx], h.data[rootIdx]
		h.percolateDown(selectedIdx)
	}
}

// 将当前节点和父节点做比较，若小于则交换位置
func (h *Heap[T]) percolateUp(childIdx int) {
	parenIdx := parentIdx(childIdx)
	if parenIdx >= 0 && h.compare(h.data[childIdx], h.data[parenIdx]) {
		h.data[childIdx], h.data[parenIdx] = h.data[parenIdx], h.data[childIdx]
		h.percolateUp(parenIdx)
	}
}
func (h *Heap[T]) heapCreat() {
	// 从最后一个非叶子节点开始，进行 down操作，
	for rootIdx := (h.length >> 1) - 1; rootIdx >= 0; rootIdx-- {
		h.percolateDown(rootIdx)
	}
}

func (h *Heap[T]) GetTop() T {
	return h.data[0]
}

func (h *Heap[T]) Len() int {
	return h.length
}
func (h *Heap[T]) IsEmpty() bool {
	return h.length == 0
}

// 放入最后一个元素 然后执行up操作
func (h *Heap[T]) Push(element T) {
	h.data = append(h.data, element)
	h.length++
	h.percolateUp(h.length - 1)
}

// 将顶部弹出，最后一个元素放入顶部执行 down 操作
func (h *Heap[T]) Pop() T {
	heapValue := h.data[0]
	h.data[0] = h.data[h.length-1]
	h.data = h.data[:h.length-1]
	h.length--
	h.percolateDown(0)
	return heapValue
}
func NewHeap[T comparable](arr []T, comparef comparefunc[T]) *Heap[T] {
	heap := &Heap[T]{
		data:    arr,
		compare: comparef,
		length:  len(arr),
	}
	heap.heapCreat()
	return heap
}
