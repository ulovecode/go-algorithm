package cn

//设计一个支持以下两种操作的数据结构：
//
// void addWord(word)
//bool search(word)
//
//
// search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。
//
// 示例:
//
// addWord("bad")
//addWord("dad")
//addWord("mad")
//search("pad") -> false
//search("bad") -> true
//search(".ad") -> true
//search("b..") -> true
//
//
// 说明:
//
// 你可以假设所有单词都是由小写字母 a-z 组成的。
// Related Topics 设计 字典树 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
type WordDictionary struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{
			next: new([]*Node),
		},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
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

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return dfsTree(0, len(word), this.root, word)
}
func dfsTree(position int, wLen int, node *Node, word string) bool {
	if position == wLen {
		if node.isComplete {
			return true
		}
		return false
	}
	var result = false
	w := int32(word[position])
	nodes := node.next
	for _, n := range *nodes {
		if n.val == w || w == '.' {
			result = result || dfsTree(position+1, wLen, n, word)
		}
	}
	return result
}

type Node struct {
	val        int32
	next       *[]*Node
	isComplete bool
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//leetcode submit region end(Prohibit modification and deletion)
