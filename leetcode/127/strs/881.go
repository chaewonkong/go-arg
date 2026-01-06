package strs

func firstUniqChar(s string) int {
	const MAX = 100001
	countMap := make(map[rune]int)
	indexMap := make(map[rune]int)
	minIdx := MAX

	for i, r := range s {
		countMap[r]++

		if _, ok := indexMap[r]; !ok {
			indexMap[r] = i
		}
	}

	for k, v := range countMap {
		if v == 1 {
			offset := indexMap[k]
			minIdx = min(offset, minIdx)
		}
	}

	// no uniq char
	if minIdx == MAX {
		return -1
	}

	return minIdx
}
