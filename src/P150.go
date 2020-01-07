package main

import (
	"fmt"
	"strconv"
)

func main() {
	token := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(token))
}

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	res := ""
	for i := 0; i < len(tokens); i ++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			push(&stack, tokens[i])
		}else {
			s1, _ := strconv.Atoi(pop(&stack))
			s2, _ := strconv.Atoi(pop(&stack))
			if tokens[i] == "+" {
				res = strconv.Itoa(s1 + s2)
				//fmt.Printf("%d + %d", s1, s2)
				push(&stack, res)
			}else if tokens[i] == "*" {
				res = strconv.Itoa(s1 * s2)
				//fmt.Printf("%d * %d", s1, s2)
				push(&stack, res)
			}else if tokens[i] == "/" {
				res = strconv.Itoa(s2 / s1)
				//fmt.Printf("%d / %d", s2, s1)
				push(&stack, res)
			}else if tokens[i] == "-" {
				res = strconv.Itoa(s2 - s1)
				//fmt.Printf("%d - %d", s2, s1)
				push(&stack, res)
			}
		}
	}
	s1, _ := strconv.Atoi(pop(&stack))
	return s1
}

func pop(stack *[]string) string {
	tmp := (*stack)[len(*stack) - 1]
	*stack = (*stack)[ : len(*stack) - 1]
	return tmp
}

func push(stack *[]string, s string) {
	*stack = append(*stack, s)
}
