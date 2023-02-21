package haffman

import (
	heap "github.com/pipikai/go_zip/heap"
)

type Node struct {
	key    int16
	weight int64
	left   *Node
	right  *Node
}
type Haffmancode struct {
	wmap map[byte]int64

	originData []byte
	heap       *heap.Heap[Node]
	Tree       Node
	Code       map[byte]string
}

func (h *Haffmancode) initWeights() {
	h.wmap = map[byte]int64{}
	for _, v := range h.originData {
		h.wmap[v] += 1
	}
	temp := []Node{}
	for k, v := range h.wmap {
		if v != 0 {
			temp = append(temp, Node{int16(k), v, nil, nil})
		}
	}
	h.heap = heap.NewHeap(temp, func(a, b Node) bool { return a.weight < b.weight })
}

func (h *Haffmancode) buildTree() {
	// 首先找到最小的两个节点
	for h.heap.Len() > 1 {
		t1 := h.heap.GetTop()
		h.heap.Pop()
		t2 := h.heap.GetTop()
		h.heap.Pop()
		tempNode := Node{
			key:    -1,
			weight: t1.weight + t2.weight,
			left:   &t1,
			right:  &t2,
		}
		h.heap.Push(tempNode)
	}
	// 排除当文件内容为空时
	if h.heap.Len() > 0 {
		h.Tree = h.heap.GetTop()
	}
}
func (h *Haffmancode) GetWeights() map[byte]int64 {
	return h.wmap
}
func (h *Haffmancode) GetTree() Node {
	return h.Tree
}

func (h *Haffmancode) GetCodeTable() map[byte]string {
	return h.Code
}
func (h *Haffmancode) dfs(node *Node, code string) {
	if node == nil {
		return
	}
	var lcode, rcode string
	lcode = code + "0"
	rcode = code + "1"
	if node.left != nil && node.left.key != -1 {
		h.Code[byte(node.left.key)] = lcode
	}
	if node.right != nil && node.right.key != -1 {
		h.Code[byte(node.right.key)] = rcode
	}
	h.dfs(node.left, lcode)
	h.dfs(node.right, rcode)
}
func (h *Haffmancode) buidHaffCode() {
	h.Code = make(map[byte]string)
	h.dfs(&h.Tree, "")
}

func NewHaffmanCode(data []byte) *Haffmancode {
	res := &Haffmancode{
		originData: data,
	}
	res.initWeights()
	res.buildTree()
	res.buidHaffCode()
	return res
}
