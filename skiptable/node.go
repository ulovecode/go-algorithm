package skiptable

type Node struct {
	key                   int
	value                 interface{}
	up, down, left, right *Node
	level                 int
}

func newNode(k int, v interface{}) *Node {
	return &Node{
		key:   k,
		value: v,
	}
}

func (n *Node) equals(node *Node) bool {
	if node == nil {
		return false
	}
	if n.key != node.key {
		return false
	}
	if n.value != node.value {
		return false
	}
	if n.level != node.level {
		return false
	}
	return true
}
