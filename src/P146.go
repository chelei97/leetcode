package main

import "container/list"

func main() {

}

type LRUCache struct {
	record map[int]list.List
}

func Constructor2(capacity int) LRUCache {
	return LRUCache{record:make(map[int]list.List, capacity)}
}


func (this *LRUCache) Get(key int) int {

}


func (this *LRUCache) Put(key int, value int)  {

}
