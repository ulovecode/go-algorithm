package redblacktree

import (
	"math/rand"
	"testing"
	"time"
)

func TestRBInsert(t *testing.T) {
	tree := NewTree()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	tree.Insert(5000, string(r.Intn(10000)))

	for i := 0; i < 100; i++ {
		n := r.Int63n(10000)
		if n != 5000 {
			tree.Insert(n, string(r.Intn(10000)))
		}
	}
	tree.Delete(5000)
	tree.Traverse()
	println(tree.size)

}
