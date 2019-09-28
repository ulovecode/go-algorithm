package cn

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
// 示例:
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//
// 说明:
//
//
// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。
//
// Related Topics 设计 字典树
//leetcode submit region begin(Prohibit modification and deletion)
//
type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &Node{
			next: new([]*Node),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var node = *this.root
	wLen := len(word)
	for i, w := range word {
		var childrenRoot *Node
		nodes := node.next
		for _, n := range *nodes {
			if n.val == w {
				childrenRoot = n
				break
			}
		}
		if childrenRoot == nil {
			childrenRoot = &Node{val: w, next: new([]*Node)}
			*node.next = append(*node.next, childrenRoot)
		}
		if i == wLen-1 {
			childrenRoot.isComplete = true
		}
		node = *childrenRoot
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var node = *this.root
	wLen := len(word)
	for i, w := range word {
		var childrenRoot *Node
		nodes := node.next
		for i, n := range *nodes {
			if n.val == w {
				childrenRoot = (*nodes)[i]
				break
			}
		}
		if childrenRoot == nil {
			return false
		}
		if i == wLen-1 && childrenRoot.isComplete {
			return true
		}
		node = *childrenRoot
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var node = *this.root
	wLen := len(prefix)
	for i, w := range prefix {
		var childrenRoot *Node
		nodes := node.next
		for i, n := range *nodes {
			if n.val == w {
				childrenRoot = (*nodes)[i]
				break
			}
		}
		if childrenRoot == nil {
			return false
		}
		if i == wLen-1 {
			return true
		}
		node = *childrenRoot
	}
	return false
}

type Node struct {
	val        int32
	next       *[]*Node
	isComplete bool
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)

//
//type Trie struct {
//	root *Node
//}
//
///** Initialize your data structure here. */
//func Constructor() Trie {
//	return Trie{
//		root: &Node{
//			next: new([]*Node),
//		},
//	}
//}
//
///** Inserts a word into the trie. */
//func (this *Trie) Insert(word string) {
//	if word == "" {
//		return
//	}
//	var node = *this.root
//	wLen := len(word)
//	for i, w := range word {
//		var childrenRoot *Node
//		nodes := node.next
//		for _, n := range *nodes {
//			if n.val == w {
//				childrenRoot = n
//				break
//			}
//		}
//		if childrenRoot == nil {
//			childrenRoot = &Node{val: w, next: new([]*Node)}
//			*node.next = append(*node.next, childrenRoot)
//		}
//		if i == wLen-1 {
//			childrenRoot.isComplete = true
//		}
//		node = *childrenRoot
//	}
//}
//
///** Returns if the word is in the trie. */
//func (this *Trie) Search(word string) bool {
//	var node = *this.root
//	wLen := len(word)
//	for i, w := range word {
//		var childrenRoot *Node
//		nodes := node.next
//		for i, n := range *nodes {
//			if n.val == w {
//				childrenRoot = (*nodes)[i]
//				break
//			}
//		}
//		if childrenRoot == nil {
//			return false
//		}
//		if i == wLen-1 && childrenRoot.isComplete {
//			return true
//		}
//		node = *childrenRoot
//	}
//	return false
//}
//
///** Returns if there is any word in the trie that starts with the given prefix. */
//func (this *Trie) StartsWith(prefix string) bool {
//	var node = *this.root
//	wLen := len(prefix)
//	for i, w := range prefix {
//		var childrenRoot *Node
//		nodes := node.next
//		for i, n := range *nodes {
//			if n.val == w {
//				childrenRoot = (*nodes)[i]
//				break
//			}
//		}
//		if childrenRoot == nil {
//			return false
//		}
//		if i == wLen-1 {
//			return true
//		}
//		node = *childrenRoot
//	}
//	return false
//}
//
//type Node struct {
//	val        int32
//	next       *[]*Node
//	isComplete bool
//}
//
