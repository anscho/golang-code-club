package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func isEqual(child, otherChild *Node) bool {
	if child == nil && otherChild == nil {
		return true
	}
	if child == nil || otherChild == nil {
		return false
	}
	return child.Equals(*otherChild)
}

func (n Node) Equals(other Node) bool {
	return n.value == other.value &&
		isEqual(n.left, other.left) &&
		isEqual(n.right, other.right)
}
