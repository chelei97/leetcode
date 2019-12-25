package main

import (
	"container/list"
	"fmt"
)

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}

func dailyTemperatures(T []int) []int {
	l := list.New()
	for i := 0; i < len(T); i ++ {
		if l.Len() == 0 || T[l.Back().Value.(int)] >= T[i] {
			l.PushBack(i)
			continue
		}
		for l.Len() != 0 && T[l.Back().Value.(int)] < T[i] {
			T[l.Back().Value.(int)] = i - l.Back().Value.(int)
			l.Remove(l.Back())
		}
		l.PushBack(i)
	}
	for l.Len() != 0 {
		T[l.Back().Value.(int)] = 0
		l.Remove(l.Back())
	}
	return T
}
