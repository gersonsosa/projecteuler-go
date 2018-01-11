package datastructures

import "errors"

type MaxHeap interface {
    Add(i int)
    Peek() (int, error)
    Poll() (int, error)
}

type maxHeap struct {
    Items []int
}

func NewMaxHeap() MaxHeap {
	return &maxHeap{make([]int, 0)}
}

func getLeftChildIdx(parentIdx int) int {
    return 2 * parentIdx + 1
}

func getRightChildIdx(parentIdx int) int {
    return 2 * parentIdx + 2
}

func getParentIdx(childIdx int) int {
    return (childIdx - 1) / 2
}

func (m maxHeap) hasLeftChild(idx int) bool {
    return getLeftChildIdx(idx) < len(m.Items)
}

func (m maxHeap) hasRightChild(idx int) bool {
    return getRightChildIdx(idx) < len(m.Items)
}

func  (m maxHeap) hasParent(idx int) bool {
    return getParentIdx(idx) >= 0
}

func (m maxHeap) leftChild(idx int) int {
    return m.Items[getLeftChildIdx(idx)]
}

func (m maxHeap) rightChild(idx int) int {
    return m.Items[getRightChildIdx(idx)]
}

func (m maxHeap) parent(idx int) int {
    return m.Items[getParentIdx(idx)]
}

func (m *maxHeap) swap(indexOne, indexTwo int) {
    m.Items[indexOne], m.Items[indexTwo] = m.Items[indexTwo], m.Items[indexOne]
}

func (m maxHeap) Peek() (int, error) {
    if len(m.Items) == 0 {
        return 0, errors.New("empty heap")
    }
    return m.Items[0], nil
}

func (m *maxHeap) Poll() (int, error) {
    size := len(m.Items)
    if size == 0 {
        return 0, errors.New("empty heap")
    }
    item := m.Items[0]
    m.Items[0] = m.Items[size - 1]
    m.Items = m.Items[0 : size - 1]
    m.heapifyDown()
    return item, nil
}

func (m *maxHeap) heapifyDown() {
    idx := 0
    for m.hasLeftChild(idx) {
        largestChildIdx := getLeftChildIdx(idx)
        if m.hasRightChild(idx) && m.rightChild(idx) > m.leftChild(idx) {
            largestChildIdx = getRightChildIdx(idx)
        }

        if m.Items[idx] < m.Items[largestChildIdx] {
            m.swap(idx, largestChildIdx)
        } else {
            break
        }
        idx = largestChildIdx
    }
}

func (m *maxHeap) Add(i int) {
    m.Items = append(m.Items, i)
    m.heapifyUp()
}

func (m *maxHeap) heapifyUp() {
    idx := len(m.Items) - 1
    for m.hasParent(idx) && m.parent(idx) < m.Items[idx] {
        m.swap(idx, getParentIdx(idx))
        idx = getParentIdx(idx)
    }
}
