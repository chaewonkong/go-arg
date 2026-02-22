package trie

type Trie struct {
	Root Node
}

func New() *Trie {
	return &Trie{
		Root: *newNode(),
	}
}

func (t *Trie) Search(target string) *Node {
	return t.Root.search(target, 0)
}

func (t *Trie) Insert(val string) {
	t.Root.insert(val, 0)
}

func (t *Trie) Delete(target string) {
	t.Root.delete(target, 0)
}

type Node struct {
	IsEntry  bool
	Children []*Node // len=26, nil..
}

func newNode() *Node {
	return &Node{
		Children: make([]*Node, 26),
	}
}

func (n *Node) search(target string, index int) *Node {
	if index == len(target) {
		if n.IsEntry {
			return n
		}
		return nil
	}

	nextChar := target[index]
	nextIdx := letterToIdx(nextChar)
	nextNode := n.Children[nextIdx]
	if nextNode == nil {
		return nil
	}

	return nextNode.search(target, index+1)
}

func (n *Node) insert(val string, index int) {
	if index == len(val) {
		n.IsEntry = true
		return
	}

	nextChar := val[index]
	nextIdx := letterToIdx(nextChar)
	nextNode := n.Children[nextIdx]
	if nextNode == nil {
		n.Children[nextIdx] = newNode()
		n.Children[nextIdx].insert(val, index+1)
		return
	}

	nextNode.insert(val, index+1)
}

func (n *Node) delete(target string, index int) bool {
	if index == len(target) {
		if n.IsEntry {
			n.IsEntry = false
		}
	} else {
		nextChar := target[index]
		nextIdx := letterToIdx(nextChar)
		nextNode := n.Children[nextIdx]
		if nextNode != nil {
			// 제거해도 안전하다면 제거한다
			if nextNode.delete(target, index+1) {
				n.Children[nextIdx] = nil
			}
		}
	}

	if n.IsEntry {
		return false
	}
	for _, child := range n.Children {
		if child != nil {
			return false
		}
	}

	return true
}

func letterToIdx(c byte) int {
	return int(c) - 97
}
