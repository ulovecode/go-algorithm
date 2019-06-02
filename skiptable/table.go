package skiptable

import (
	"math/rand"
	"time"
)

type skipTable struct {
	header *Node
	r      *rand.Rand
}

func New() *skipTable {
	return &skipTable{
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
		header: newNode(-int(^uint(0) >> 1), nil),
	}
}

func (jt *skipTable) Search(k int) *Node {
	// println("walkPreviousNode")
	node, _ := walkPreviousNode(jt.header, k)
	// println("共需要", count, "步")
	return node
}

func (jt *skipTable) Insert(k int, v interface{}) {
	sn := jt.header
	newN := newNode(k, v)
	node, _ := walkPreviousNode(sn, k)
	if node.key == k {
		node.value = v
		return
	}
	currentLevel := 0
	newN.level = currentLevel
	jt.setNode(newN, node)
	lowNode := newN
	for jt.isPromotion() {
		// println("isPromotion")
		currentLevel++
		upNewN := setUpNewDownNode(currentLevel, k, v, newN)
		if jt.header.level < currentLevel {
			updateHeader(jt, upNewN)
		}
		leftUpNode := findLeftUpNode(lowNode, upNewN)
		if leftUpNode != nil {
			jt.setNode(upNewN, leftUpNode)
		}
		newN = upNewN
	}
}

func updateHeader(jt *skipTable, upNew *Node) {
	jt.header = upNew
}

func setUpNewDownNode(level int, k int, v interface{}, newN *Node) *Node {
	upNew := newNode(k, v)
	upNew.level = level
	newN.up = upNew
	upNew.down = newN
	return upNew
}

func findLeftUpNode(newN *Node, upNew *Node) *Node {
	var leftUpNode *Node
	leftNewNode := newN.left
leftBreak:
	for leftNewNode != nil {
		leftNewUpNode := leftNewNode.up
		for leftNewUpNode != nil {
			if leftNewUpNode.level == upNew.level {
				leftUpNode = leftNewUpNode
				break leftBreak
			}
			leftNewUpNode = leftNewUpNode.up
		}
		leftNewNode = leftNewNode.left
	}
	return leftUpNode
}

// new 后面插入cur
func (jt *skipTable) setNode(q *Node, p *Node) {
	q.left = p
	if p.right != nil {
		q.right = p.right
		p.right.left = q
	}
	p.right = q

}

func walkPreviousNode(curNode *Node, k int) (*Node, int) {
	// println("k:", k, " level", curNode.level)
	count := 0
	for curNode.key < k {
		count++
		if curNode.right == nil {
			break
		}
		// println("curNode.key < k ", curNode.key, "-->", curNode.right.key)
		curNode = curNode.right
	}
	for curNode.key > k {
		count++
		if curNode.left == nil {
			break
		}
		// println("curNode.key > k ", curNode.key, "-->", curNode.left.key)
		curNode = curNode.left
	}
	if curNode.down != nil {
		var newCount int
		// println("curNode.down ", curNode.key, "-->", curNode.down.key)
		curNode, newCount = walkPreviousNode(curNode.down, k)
		count += newCount
	}
	return curNode, count
}

func (jt *skipTable) isPromotion() bool {
	n := jt.r.Intn(2)
	if n == 0 {
		return false
	}
	return true
}
