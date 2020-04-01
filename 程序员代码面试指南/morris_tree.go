package main

type TreeNode struct {
	left, right *TreeNode
	val         int
}

func morris(root *TreeNode) {
	cur := root
	for cur != nil {
		mostRight := cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		}
		cur = cur.right
	}
}

func preMorris(root *TreeNode) {
	cur := root
	for cur != nil {
		mostRight := cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				println(cur.val)
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		} else {
			println(cur.val)
		}
		cur = cur.right
	}
}

func inorderMorris(root *TreeNode) {
	cur := root
	for cur != nil {
		mostRight := cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		}
		println(cur.val)
		cur = cur.right
	}
}

func postMorris(root *TreeNode) {
	cur := root
	for cur != nil {
		mostRight := cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		}
		cur = cur.right
	}
}
