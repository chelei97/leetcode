package main

import (
	"fmt"
	"sort"
)

func main() {
	deck := []int{17,13,11,2,3,5,7}
	fmt.Println(deckRevealedIncreasing(deck))
}

func deckRevealedIncreasing(deck []int) []int {
	if len(deck) < 2 {
		return deck
	}
	sort.Ints(deck)
	queue := make([]int, 0)
	for i := len(deck) - 1; i >= 0 ; i -- {
		add1(&queue, deck[i])
		if i == 0 {
			break
		}
		add1(&queue, poll(&queue))
	}
	for i := len(deck) - 1; i >= 0; i-- {
		deck[i] = poll(&queue)
	}
	return deck
}

func poll(queue *[]int) int {
	tmp := (*queue)[0]
	*queue = (*queue)[1 : ]
	return tmp
}

func add1(queue *[]int, value int) {
	*queue = append(*queue, value)
}
