package astar

import (
	"container/list"
	"math"
)

var directionMap = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type aStar struct {
	start     *Node
	end       *Node
	searchMap [][]int
	openList  *list.List
	closeList *list.List
}

func NewAStar(start *Node, end *Node, searchMap [][]int) *aStar {
	star := &aStar{
		start:     start,
		end:       end,
		searchMap: searchMap,
		openList:  list.New(),
		closeList: list.New(),
	}
	star.openList.PushBack(star.start)
	return star
}

func (a *aStar) FindPath() *Node {
	for a.openList.Len() != 0 {
		element := a.getMinElement()
		curNode := element.Value.(*Node)
		a.openList.Remove(element)
		a.closeList.PushBack(curNode)
		if curNode.Col == a.end.Col && curNode.Row == a.end.Row {
			a.end = curNode
			break
		}
		for _, dm := range directionMap {
			a.handNode(curNode, curNode.Row+dm[0], curNode.Col+dm[1])
		}
	}
	return a.end
}

func (a *aStar) getMinElement() *list.Element {
	element := a.openList.Front()
	for e := a.openList.Front(); e != nil; e = e.Next() {
		if element.Value.(*Node).F > e.Value.(*Node).F {
			element = e
		}
	}
	return element
}

func (a *aStar) handNode(curNode *Node, row int, col int) {
	if a.isNotPass(row, col) {
		return
	}
	node := NewNode(row, col)

	if !a.isNotExistInCloseList(node) {
		return
	}
	if a.isNotExistInOpenList(node) && a.isNotExistInCloseList(node) {
		a.setF(node)
		node.father = curNode
		a.openList.PushBack(node)
	}
}

func (a *aStar) isNotExistInCloseList(node *Node) bool {
	for e := a.closeList.Front(); e != nil; e = e.Next() {
		n := e.Value.(*Node)
		if n.Row == node.Row && n.Col == node.Col {
			return false
		}
	}
	return true
}

func (a *aStar) isNotExistInOpenList(node *Node) bool {
	for e := a.openList.Front(); e != nil; e = e.Next() {
		n := e.Value.(*Node)
		if n.Row == node.Row && n.Col == node.Col {
			return false
		}
	}
	return true
}

func (a *aStar) isNotPass(row int, col int) bool {
	maxRow := len(a.searchMap)
	maxCol := len(a.searchMap[0])
	if row < 0 || row >= maxRow {
		return true
	}
	if col < 0 || col >= maxCol {
		return true
	}
	if a.searchMap[row][col] == -1 {
		return true
	}
	return false
}

func (a *aStar) setF(node *Node) {
	g := math.Abs(float64(a.start.Row-node.Row)) + math.Abs(float64(a.start.Col-node.Col))
	h := math.Abs(float64(a.end.Row-node.Row)) + math.Abs(float64(a.end.Col-node.Col))
	node.F = int(g + h)
}

func (a *aStar) PrintNode() {
	walkMap := make([][]int, len(a.searchMap))
	for k := range walkMap {
		walkMap[k] = make([]int, len(a.searchMap[0]))
	}
	n := a.end
	for n != nil {
		walkMap[n.Row][n.Col] = 1
		n = n.father
	}
	count := 0
	for _, v := range walkMap {
		for _, v2 := range v {
			if v2 == 1 {
				count++
			}
			print(v2)
		}
		println()
	}
	println("一共需要", count, "步")
}
