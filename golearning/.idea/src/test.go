package main

import "fmt"

func main() {
	boool := true
	defer func() {
		fmt.Println("这是defer语句")
	}()
	if boool {
		fmt.Println("判断内语句")
		return
	}
	return
}

