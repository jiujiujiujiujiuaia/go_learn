package main

import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
type LRUCache struct {
	Space      int
	hashMap    map[int]*list.Element
	LinkedList *list.List
}

//注意，需要一个node放在双端队列中
//因为当容量超出需要删除的时候，node中的key可以帮忙也删除
//map中的值
type Node struct {
	Key int
	Val int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{}
	lruCache.Space = capacity
	lruCache.hashMap = make(map[int]*list.Element)
	lruCache.LinkedList = list.New()
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.hashMap[key]; ok {
		this.Put(key, val.Value.(Node).Val)
		return val.Value.(Node).Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.hashMap[key]; ok {
		//1.删除原有节点
		this.LinkedList.Remove(val)
		//2.赋值一个新的节点
		node := Node{Val: value, Key: key}
		//3.插入到队首，并且更新map
		this.LinkedList.PushFront(node)
		newVal := this.LinkedList.Front()
		this.hashMap[key] = newVal
	} else {
		//1.判断是否超出缓存，如果超出，则删除队尾。无论如何直接插到队首
		if len(this.hashMap) == this.Space {
			tailNode := this.LinkedList.Back().Value.(Node)
			this.LinkedList.Remove(this.LinkedList.Back())
			delete(this.hashMap, tailNode.Key)
		}
		//同上
		node := Node{Val: value, Key: key}
		this.LinkedList.PushFront(node)
		newVal := this.LinkedList.Front()
		this.hashMap[key] = newVal
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
