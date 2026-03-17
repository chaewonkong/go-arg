package topnfreq

func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		cntMap[num]++
	}

	h := newHeap()
	for n, cnt := range cntMap {
		h.insert(newElement(n, cnt))

		if h.len() > k {
			h.popMin() // remove
		}
	}

	results := []int{}
	for _, el := range h.elements() {
		results = append(results, el.value)
	}

	return results
}

// heap min heap
type heap struct {
	arr []element
}

type element struct {
	value int
	cnt   int
}

func newElement(v, cnt int) element {
	return element{value: v, cnt: cnt}
}

func newHeap() *heap {
	return &heap{
		arr: []element{newElement(0, 0)}, // padding
	}
}

func (h *heap) insert(v element) {
	last := len(h.arr)
	h.arr = append(h.arr, v)

	cur := last
	parent := cur / 2
	for parent >= 1 && h.arr[parent].cnt > h.arr[cur].cnt {
		h.arr[parent], h.arr[cur] = h.arr[cur], h.arr[parent]
		cur = parent
		parent = parent / 2 // 올라가며 탐색
	}
}

func (h *heap) len() int {
	return len(h.arr) - 1
}

func (h *heap) elements() []element {
	if len(h.arr) == 1 {
		return []element{}
	}
	return h.arr[1:]
}

func (h *heap) popMin() (element, bool) {
	if len(h.arr) == 1 {
		return element{}, false
	}

	root := h.arr[1]
	last := len(h.arr) - 1
	h.arr[1] = h.arr[last]
	h.arr = h.arr[:last]

	// 1 ~ last까지 위에서부터 비교하며 swap
	last = last - 1 // 하나 pop했으므로 length - 1
	i := 1
	for i <= last {
		swap := i
		if idx := 2 * i; idx <= last && (h.arr[swap].cnt > h.arr[idx].cnt) {
			swap = idx
		}

		if idx := 2*i + 1; idx <= last && (h.arr[swap].cnt > h.arr[idx].cnt) {
			swap = idx
		}

		if swap != i {
			h.arr[swap], h.arr[i] = h.arr[i], h.arr[swap]
			i = swap
		} else {
			break
		}
	}

	return root, true
}

func topNFreqWithBucket(nums []int, k int) []int {
	cntMap := make(map[int]int)

	// check frequency and make deduplication
	for _, num := range nums {
		cntMap[num]++
	}

	bucket := make([][]int, len(nums)+1) // cnt를 index로
	for num, cnt := range cntMap {
		bucket[cnt] = append(bucket[cnt], num)
	}

	res := make([]int, 0, k)

	// starting from last item; greatest cnt
	for i := len(bucket) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, bucket[i]...)
	}

	return res
}
