package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	fmt.Println(getPermutation(4, 9))
}

func getPermutation(n int, k int) string {
	res := make([]byte, n)
	size := factorial(n - 1)
	l := list.New()
	for i := 0; i < n; i ++ {
		l.PushBack(i + 1)
	}
	for i := 0; i < n; i ++ {
		index := int(math.Ceil(float64(k) / float64(size)))
		e := l.Front()
		for j := 1; j < index; j ++ {
			e = e.Next()
		}
		res[i] = byte(e.Value.(int) + '0')
		l.Remove(e)
		k = k - (index - 1) * size
		if i != n - 1 {
			size = size / (n - 1 - i)
		}
	}
	return string(res)
}

func factorial(n int) int {
	res := n
	for i := n - 1; i > 0; i -- {
		res = res * i
	}
	return res
}
