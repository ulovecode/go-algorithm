package interview

import (
	"github.com/ulovecode/go-algorithm/leetcode/editor/cn"
)

type Node struct {
	next *Node
	val  uint8
}

func add(a, b string) string {
	if a == "" || b == "" {
		if a == "" {
			return b
		}
		if b == "" {
			return a
		}
	}
	if len(a) < len(b) {
		return add(b, a)
	}
	a = cn.Reverse(a)
	b = cn.Reverse(b)
	nodeA := numToNode(a)
	nodeB := numToNode(b)
	head := nodeA
	var preNodeA *Node
	var c uint8 = 0
	for nodeB != nil {
		nodeA.val = nodeA.val - '0' + nodeB.val - '0' + c
		c = 0
		if nodeA.val >= 10 {
			nodeA.val %= 10
			c = 1
		}
		nodeB = nodeB.next
		preNodeA = nodeA
		nodeA = nodeA.next
	}
	if c == 1 {
		if nodeA != nil {
			nodeA.next = &Node{
				next: nil,
				val:  c,
			}
		} else {
			preNodeA.next = &Node{
				next: nil,
				val:  c,
			}
		}

	}
	res := ""
	for head != nil {
		res += string(head.val + '0')
		head = head.next
	}
	return cn.Reverse(res)
}

func numToNode(n string) *Node {
	var head = &Node{
		next: nil,
		val:  n[0],
	}
	cur := head
	for i := 1; i < len(n); i++ {
		cur.next = &Node{val: n[i]}
		cur = cur.next
	}
	return head
}
