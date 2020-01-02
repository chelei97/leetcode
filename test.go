package main

import "fmt"

func main() {
	i := []byte{'a'}
	i = append(i, ' ', 'a')
	fmt.Println(string(i))
}
