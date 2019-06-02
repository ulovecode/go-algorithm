package astar

type Node struct {
	Row int
	Col int
	F   int
	father    *Node
}

func NewNode(row, col int) *Node {
	return &Node{
		Col: col,
		Row: row,
	}
}
