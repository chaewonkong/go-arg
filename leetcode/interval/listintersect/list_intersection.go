package listintersect

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0

	results := [][]int{}

	for i < len(firstList) && j < len(secondList) {
		first := firstList[i]
		second := secondList[j]

		start := max(first[0], second[0])
		end := min(first[1], second[1])

		if start <= end { // effective intersection
			results = append(results, []int{start, end})
		}

		if first[1] < second[1] {
			i++
		} else {
			j++
		}
	}

	return results
}
