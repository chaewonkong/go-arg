package bfsdfs

import "fmt"

func BFS(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node)

		for _, next := range graph[node] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
}

func DFS(graph map[int][]int, node int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true
	fmt.Println(node)

	for _, next := range graph[node] {
		DFS(graph, next, visited)
	}
}
