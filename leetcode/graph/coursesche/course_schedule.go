// Package coursesche course schedule
//
//   - https://leetcode.com/problems/course-schedule/description/
//   - Time: O(V + E); V = numCourses, E = len(prerequisites)
//   - Space: O(V + E); V number of Vertexes, E number of Edges
package coursesche

func canFinish(numCourses int, prerequisites [][]int) bool {
	// make graph
	graph := make([][]int, numCourses)
	for _, course := range prerequisites {
		cur, pre := course[0], course[1]
		graph[pre] = append(graph[pre], cur)
	}

	visited := make([]int, numCourses) // 0: unvisited, 1: visiting in current iteration, 2: visited before

	var dfs func(node int) bool

	dfs = func(node int) bool {
		if visited[node] == 1 {
			return false
		}

		if visited[node] == 2 {
			return true
		}

		visited[node] = 1
		for _, next := range graph[node] {
			if !dfs(next) {
				return false
			}
		}

		// now safe
		visited[node] = 2
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
