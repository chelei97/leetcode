package main

import (
	"container/heap"
	"fmt"
)

func main() {
	result := make([]int, 5, 10)
	heap.Init()
	fmt.Println(result)
}

