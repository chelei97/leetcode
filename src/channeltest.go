package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan int)
	go send(ch1)
	go recieve(ch1)
	time.Sleep(1*time.Second)
}

func send(ch1 chan int)  {
	for i := 0; i < 100; i ++ {
		ch1 <- i
	}
}

func recieve(ch1 chan int)  {
	for {
		fmt.Println(<- ch1)
	}
}
