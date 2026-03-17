package laststone

func lastStoneWeight(stones []int) int {
	h := newHeap()
	for _, stone := range stones {
		h.insert(stone)
	}

	for h.len() > 1 {
		y, _ := h.popMax()
		x, _ := h.popMax()

		if y != x {
			h.insert(y - x)
		}

	}
	if h.len() == 1 {
		res, _ := h.popMax()
		return res
	}
	return 0
}

type heap struct {
	stones []int
}

func newHeap() *heap {
	return &heap{
		stones: []int{0},
	}
}

func (h *heap) len() int {
	return len(h.stones) - 1
}

func (h *heap) insert(v int) {
	last := len(h.stones)
	h.stones = append(h.stones, v)

	// 아래에서부터 비교하며 교환해 올라가기
	cur := last
	parent := cur / 2 // binary tree

	for parent >= 1 && h.stones[parent] < h.stones[cur] {
		h.stones[parent], h.stones[cur] = h.stones[cur], h.stones[parent]
		cur = parent
		parent = cur / 2
	}
}

func (h *heap) popMax() (int, bool) {
	if len(h.stones) == 1 {
		return 0, false
	}

	// pop root, move last to root pos, compare items starting from root
	root := h.stones[1]
	last := len(h.stones) - 1
	h.stones[1] = h.stones[last]
	h.stones = h.stones[:last]

	last = len(h.stones) - 1
	i := 1
	for i <= last {
		swap := i

		if idx := 2 * i; idx <= last && h.stones[idx] > h.stones[swap] {
			swap = idx
		}
		if idx := 2*i + 1; idx <= last && h.stones[idx] > h.stones[swap] {
			swap = idx
		}

		if swap != i {
			h.stones[swap], h.stones[i] = h.stones[i], h.stones[swap]
			i = swap
		} else {
			break
		}
	}

	return root, true
}
