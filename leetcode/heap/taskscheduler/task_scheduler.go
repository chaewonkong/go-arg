package taskscheduler

func leastInterval(tasks []byte, n int) int {
	time := 0

	freq := make(map[byte]int)

	for _, task := range tasks {
		freq[task]++
	}

	heap := newMaxHeap()
	for _, cnt := range freq {
		heap.push(cnt)
	}

	// cool down queue
	queue := []task{}

	for heap.len() > 0 || len(queue) > 0 {
		// task 진행
		time++

		// heap 에서 실행
		if heap.len() > 0 {
			curCnt := heap.pop()
			curCnt--

			if curCnt > 0 {
				// cool down queue에 추가
				queue = append(queue, task{cnt: curCnt, available: time + n}) // n 만큼 cool down
			}
		}

		// queue에 있고 cool down 완료
		if len(queue) > 0 && queue[0].available == time {
			heap.push(queue[0].cnt)
			queue = queue[1:]
		}
	}

	return time
}

// task
type task struct {
	cnt int

	// 다시 실행 가능한 시간
	available int
}

// max heap
type maxHeap struct {
	arr []int // count for each task
}

func newMaxHeap() *maxHeap {
	return &maxHeap{
		arr: []int{0},
	}
}

func (h *maxHeap) len() int {
	return len(h.arr) - 1
}

func (h *maxHeap) push(v int) {
	h.arr = append(h.arr, v)
	last := h.len()

	cur := last
	parent := cur / 2

	for parent >= 1 && h.arr[parent] < h.arr[cur] {
		// swap
		h.arr[parent], h.arr[cur] = h.arr[cur], h.arr[parent]
		cur = parent
		parent = cur / 2
	}
}

func (h *maxHeap) pop() int {
	if h.len() == 0 {
		return 0 // not ok
	}

	root := h.arr[1]
	last := h.len()
	h.arr[1] = h.arr[last]
	h.arr = h.arr[:last]

	last = h.len()
	i := 1

	for i <= last {
		swap := i

		if idx := 2 * i; idx <= last && h.arr[swap] < h.arr[idx] {
			swap = idx
		}
		if idx := 2*i + 1; idx <= last && h.arr[swap] < h.arr[idx] {
			swap = idx
		}

		if swap != i {
			// swap
			h.arr[i], h.arr[swap] = h.arr[swap], h.arr[i]
			i = swap
		} else {
			break
		}
	}

	return root
}
