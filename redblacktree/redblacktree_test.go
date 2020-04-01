package redblacktree

import (
	"math/rand"
	"testing"
	"time"
)

func TestRBInsert(t *testing.T) {
	root = &RBTree{
		Key:    0,
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Color:  Black,
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100; i++ {
		RBInsert(r.Intn(10000))
	}
	inOrderRecu(root)
}
