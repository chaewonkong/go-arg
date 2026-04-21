// Package coursesche2
//
//   - https://leetcode.com/problems/course-schedule-ii/description/
//   - Time: O(V + E); V = numCourses, E = len(prerequisites)
//   - Space: O(V + E); V number of Vertexes, E number of Edges
package coursesche2

import "slices"

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		graph[b] = append(graph[b], a)
	}

	visited := make([]int, numCourses)

	courses := []int{}
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if visited[node] == 1 {
			return false
		}
		if visited[node] == 2 {
			return true
		}

		visited[node] = 1
		for _, edge := range graph[node] {
			if !dfs(edge) {
				return false
			}
		}

		visited[node] = 2
		courses = append(courses, node)
		return true
	}

	for v := range numCourses {
		if !dfs(v) {
			return []int{}
		}
	}

	slices.Reverse(courses)
	return courses
}
