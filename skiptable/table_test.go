package skiptable

import (
	"math/rand"
	"testing"
	"time"
)

func TestJumpTable_Insert(t *testing.T) {
	for i:=0 ;i <1;i++ {
		table := New()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		table.Insert(13, nil)
		for i := 0; i < 1000; i++ {
			table.Insert(r.Intn(100000), nil)
		}

		node := table.Search(13)
		if node.key != 13{
			panic("")
		}
		println(node.key)
	}
}