package leetcode

import (
	"testing"
)

func TestConstructorLRUCache(t *testing.T) {
	lruCache := ConstructorLRUCache(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	println(lruCache.Get(3))
	lruCache.Put(5, 5)
	lruCache.Put(6, 6)
	lruCache.Put(7, 7)
	println(lruCache.Get(1))
	println(lruCache.Get(2))
	println(lruCache.Get(3))
	println(lruCache.Get(4))
	println(lruCache.Get(5))
}
