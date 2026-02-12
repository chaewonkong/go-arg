package design

type Node struct {
	Val  int
	Prev *Node
}

type MinStack struct {
	MinTail   *Node
	StackTail *Node
}

func Constructor() MinStack {
	return MinStack{
		MinTail:   nil,
		StackTail: nil,
	}
}

func (this *MinStack) Push(val int) {
	node := &Node{
		Val: val,
	}

	if this.StackTail == nil {
		this.StackTail = node
	} else {
		tail := this.StackTail
		this.StackTail = node
		node.Prev = tail
	}

	if this.MinTail == nil {
		mn := &Node{
			Val: val,
		}
		this.MinTail = mn
		return
	}

	if this.MinTail.Val > val {
		// add to min
		mn := &Node{
			Val: val,
		}
		min := this.MinTail
		mn.Prev = min
		this.MinTail = mn

		return
	}

	mn := &Node{
		Val: this.MinTail.Val,
	}
	min := this.MinTail
	mn.Prev = min
	this.MinTail = mn
}

func (this *MinStack) Pop() {
	stackPrev := this.StackTail.Prev
	minPrev := this.MinTail.Prev

	this.StackTail = stackPrev
	this.MinTail = minPrev
}

func (this *MinStack) Top() int {
	return this.StackTail.Val
}

func (this *MinStack) GetMin() int {
	return this.MinTail.Val
}
