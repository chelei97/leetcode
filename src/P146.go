package main

import (
	"container/list"
)

func main() {
	obj := Constructor(2)
	obj.Put(2,1)
	obj.Put(1,1)
	obj.Put(2,3)
	obj.Put(4,1)
	obj.Get(1)
	obj.Get(2)
}

type LRUCache struct {
	value list.List
	cap int
	record map[int]*list.Element
}

type Pair struct {
	key int
	value int
}

func Constructor(capacity int) LRUCache {
	value := list.List{}
	cap := capacity
	record := make(map[int]*list.Element)
	return LRUCache{record:record, cap:cap, value:value}
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.record[key]; !ok {
		return -1
	}
	this.value.MoveToFront(this.record[key])
	return this.record[key].Value.(Pair).value
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.record[key]; ok {
		this.value.MoveToFront(this.record[key])
		this.record[key].Value = Pair{value:value,key:key,}
		return
	}
	this.value.PushFront(Pair{key : key, value : value,})
	this.record[key] = this.value.Front()
	if this.value.Len() <= this.cap {
		return
	}
	back := this.value.Back()
	this.value.Remove(back)
	delete(this.record, back.Value.(Pair).key)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
