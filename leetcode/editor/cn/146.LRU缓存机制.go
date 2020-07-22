//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
// 进阶:
//
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
// 示例:
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
//

package leetcode

import "container/list"

type Entry struct {
	value int
	e     *list.Element
}

type LRUCache struct {
	cache    map[int]Entry
	preList  *list.List
	capacity int
	length   int
}

func ConstructorLRUCache(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]Entry, capacity),
		capacity: capacity,
		preList:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	entry := this.cache[key]
	this.preList.MoveToBack(entry.e)
	return entry.value
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		this.preList.Remove(v.e)
		delete(this.cache, v.e.Value.(int))
	} else {
		if len(this.cache) == this.capacity {
			removeKey := this.preList.Remove(this.preList.Front())
			delete(this.cache, removeKey.(int))
		}
	}
	e := this.preList.PushBack(key)
	this.cache[key] = Entry{value: value, e: e}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := ConstructorLRUCache(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
