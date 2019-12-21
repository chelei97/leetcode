package main

import "fmt"

//有效的括号

func main(){
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	result := make([]byte, len(s))
	length := 0
	for i := 0; i < len(s); i ++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			result[length] = s[i]
			length ++
		}else {
			if length == 0 {
				return false
			}
			if (s[i] == ')' && result[length - 1] == '(') || (s[i] == '}' && result[length - 1] == '{') || (s[i] == ']' && result[length - 1] == '[') {
				length --
			}else {
				return false
			}
		}
	}
	return length == 0
}