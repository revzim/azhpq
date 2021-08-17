// revzim <https://github.com/revzim>
// AZHPQ - HEAP PRIORITY QUEUE
package azhpq

type (
	// HeapPriorityQueue --
	// HEAP PRIORITY QUEUE (HPQ) TYPE
	HeapPriorityQueue struct {
		Compare func(a, b *QueueNode) bool `json:"compare"`
		Queue   []*QueueNode               `json:"queue"`
		Size    int                        `json:"size"`
	}

	// QueueNode --
	// ITEM TYPE TO POPULATE HPQ
	QueueNode struct {
		Value    interface{} `json:"value"`
		Priority int         `json:"priority"`
	}
)

const ()

var ()

// Compare --
// COMPARE PRIO OF A & B
func Compare(a, b *QueueNode) bool {
	return a.Priority > b.Priority
}

// New --
// NEW HEAP PRIO QUEUE STRUCTURE
func New() *HeapPriorityQueue {
	hpq := &HeapPriorityQueue{
		Compare: Compare,
		Queue:   make([]*QueueNode, 0),
		Size:    0,
	}
	return hpq
}

// Fix --
// COMPARES & FIXES BINARY HEAP, RETURNS NEW INDEX
func (hpq *HeapPriorityQueue) Fix(i int, val *QueueNode) int {
	var compareIndex int
	var compareQueueNode *QueueNode
	for i > 0 {
		compareIndex = (i - 1) >> 1
		compareQueueNode = hpq.Queue[compareIndex]
		if !hpq.Compare(val, compareQueueNode) {
			break
		}
		hpq.Queue[i] = compareQueueNode
		i = compareIndex
	}
	return i
}

// percUp --
// PERCOLATES UP & FIXES BINARY HEAP
func (hpq *HeapPriorityQueue) percUp(i int) {
	val := hpq.Queue[i]
	index := hpq.Fix(i, val)
	hpq.Queue[index] = val
}

// percDown --
// PERCOLATES DOWN & LEVELS BINARY HEAP
func (hpq *HeapPriorityQueue) percDown(i int) {
	heapSize := uint(hpq.Size) >> 1 // LOGICAL SHIFT | UINT -> ZERO FILL RIGHT SHIFT
	val := hpq.Queue[i]
	var leftIndex int
	var rightIndex int
	var heapChoice *QueueNode
	for uint(i) < heapSize {
		leftIndex = (i << 1) + 1
		rightIndex = leftIndex + 1
		heapChoice = hpq.Queue[leftIndex]
		if rightIndex < hpq.Size {
			if hpq.Compare(hpq.Queue[rightIndex], heapChoice) {
				leftIndex = rightIndex
				heapChoice = hpq.Queue[rightIndex]
			}
		}
		if !hpq.Compare(heapChoice, val) {
			break
		}
		hpq.Queue[i] = heapChoice
		i = leftIndex
	}
	hpq.Queue[i] = val
}

// RemoveAt --
// REMOVES ITEM AT SPECIFIED INDEX & RETURNS IF FOUND
func (hpq *HeapPriorityQueue) RemoveAt(i int) *QueueNode {
	if i > hpq.Size-1 || i < 0 {
		return nil
	}
	hpq.percUp(i)
	return hpq.Poll()
}

// Poll --
// RETRIEVES (& REMOVES) HIGHEST PRIO ITEM
func (hpq *HeapPriorityQueue) Poll() *QueueNode {
	if hpq.Size == 0 {
		return nil
	}
	val := hpq.Queue[0]
	if hpq.Size > 1 {
		hpq.Size--
		hpq.Queue[0] = hpq.Queue[hpq.Size]
		hpq.percDown(0)
	} else {
		hpq.Size--
	}
	return val
}

// Peek --
// VIEW THE HIGHEST PRIO ITEM & RETURN
func (hpq *HeapPriorityQueue) Peek() *QueueNode {
	if hpq.Size == 0 {
		return nil
	} else {
		return hpq.Queue[0]
	}
}

// IsEmpty --
// RETURNS TRUE IF EMPTY, FALSE OTHERWISE
func (hpq *HeapPriorityQueue) IsEmpty() bool {
	return hpq.Size == 0
}

// ForEach --
// ITERATES OVER BINARY HEAP & ALLOWS DATA MANIPULATION
func (hpq *HeapPriorityQueue) ForEach(cb func(val *QueueNode, index int)) {
	if !hpq.IsEmpty() {
		clonedHpq := hpq.Clone()
		i := 0
		for !clonedHpq.IsEmpty() {
			cb(clonedHpq.Poll(), i)
			i++
		}
	}
}

// Add --
// ADD AN ITEM TO THE BINARY HEAP
func (hpq *HeapPriorityQueue) Add(val *QueueNode) {
	i := hpq.Size
	hpq.Queue = append(hpq.Queue, val)
	hpq.Size++
	index := hpq.Fix(i, val)
	hpq.Queue[index] = val
}

// AddMany --
// ADD MANY ITEMS TO THE BINARY HEAP
func (hpq *HeapPriorityQueue) AddMany(vals ...*QueueNode) {
	for index := range vals {
		hpq.Add(vals[index])
	}
}

// Remove --
// WILL ATTEMPT TO REMOVE AN ITEM FROM THE HEAP
// RETURN TRUE IF FOUND, FALSE OTHERWISE
func (hpq *HeapPriorityQueue) Remove(val *QueueNode) bool {
	didRemove := false
	for i := 0; i < hpq.Size; i++ {
		if !hpq.Compare(hpq.Queue[i], val) &&
			!hpq.Compare(val, hpq.Queue[i]) {
			hpq.RemoveAt(i)
			didRemove = true
			break
		}
	}
	return didRemove
}

// RemoveOne --
// WILL REMOVE ONE AND RETURN ONE VALUE FROM THE BINARY HEAP IF FOUND, NIL OTHERWISE
func (hpq *HeapPriorityQueue) RemoveOne(cb func(val *QueueNode) bool) *QueueNode {
	for i := 0; i < hpq.Size; i++ {
		if cb(hpq.Queue[i]) {
			return hpq.RemoveAt(i)
		}
	}
	return nil
}

// RemoveMany --
// WILL REMOVE LIMIT AMOUNT OF ITEMS FROM THE BINARY HEAP AND RETURN THEM ALL IF FOUND, OTHERWISE RETURN NIL
func (hpq *HeapPriorityQueue) RemoveMany(cb func(val *QueueNode) bool, limit int) []*QueueNode {
	if hpq.Size < 1 {
		return nil
	} else {
		size := 0
		var removedItems []*QueueNode
		tmpSize := 0
		tmpItems := make([]*QueueNode, hpq.Size)
		for size < limit && !hpq.IsEmpty() {
			val := hpq.Poll()
			if cb(val) {
				removedItems = append(removedItems, val)
				size++
			} else {
				tmpItems[tmpSize] = val
				tmpSize++
			}
		}
		i := 0
		for i < tmpSize {
			hpq.Add(tmpItems[i])
			i++
		}
		return removedItems
	}
}

// ReplaceTop --
// FORCE REPLACE THE HIGHEST MOST PRIORITY VALUE
func (hpq *HeapPriorityQueue) ReplaceTop(val *QueueNode) *QueueNode {
	if hpq.Size == 0 {
		return nil
	} else {
		replacedVal := hpq.Queue[0]
		hpq.percDown(0)
		return replacedVal
	}
}

// NSmallest --
// RETURNS THE N AMOUNT OF THE 'SMALLEST' ITEMS, NIL IF SIZE == 0
func (hpq *HeapPriorityQueue) NSmallest(n int) []*QueueNode {
	if hpq.Size == 0 {
		return nil
	}
	n = min(hpq.Size, n)
	newHPQ := New()
	var newSize int
	if n > 0 {
		newSize = (n - 1) * (n - 1)
	} else {
		newSize = 0
	}
	newSize++
	newHPQ.Size = min(newSize, hpq.Size)
	newHPQ.Queue = hpq.Queue[:newSize]
	smallest := make([]*QueueNode, n)
	for i := 0; i < n; i++ {
		smallest[i] = newHPQ.Poll()
	}
	return smallest
}

// min --
// INT MIN HELPER METHOD (MATH PKG IS FLOAT64)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Clone --
// RETURNS CLONED STRUCTURE
func (hpq *HeapPriorityQueue) Clone() *HeapPriorityQueue {
	clonedHpq := New()
	for i := range hpq.Queue {
		clonedHpq.Queue = append(clonedHpq.Queue, hpq.Queue[i])
	}
	clonedHpq.Size = hpq.Size
	clonedHpq.Compare = hpq.Compare
	return clonedHpq
}

// Trim --
// GC & TRIMMING OF BINARY HEAP
func (hpq *HeapPriorityQueue) Trim() {
	hpq.Queue = hpq.Queue[0:hpq.Size]
	hpq.Size = len(hpq.Queue)
}
