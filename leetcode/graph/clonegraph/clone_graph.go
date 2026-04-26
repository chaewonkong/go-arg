// Package clonegraph
//
//   - https://leetcode.com/problems/clone-graph/
//   - time: O(N+E); len(nodes) + len(edges)
//   - space: O(N)
package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var dfs func(node *Node, visited map[*Node]*Node) *Node
	dfs = func(cur *Node, visited map[*Node]*Node) *Node {
		if cur == nil {
			return nil
		}
		if cloned, ok := visited[cur]; ok {
			return cloned
		}
		cloned := &Node{Val: cur.Val}
		visited[cur] = cloned

		for _, neighbor := range cur.Neighbors {
			cloned.Neighbors = append(cloned.Neighbors, dfs(neighbor, visited))
		}

		return cloned
	}

	return dfs(node, make(map[*Node]*Node))
}
