package main

import "fmt"

//电话号码的字母组合
func main(){
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

var phone map[string]string
var output []string

func backtrack(combination, next_digits string) {
	if len(next_digits) == 0 {
		output = append(output, combination)
	}else {
		digit := next_digits[0:1]
		letters := phone[digit]
		for i := 0; i < len(letters); i ++ {
			letter := letters[i : i + 1]
			backtrack(combination + letter, next_digits[1 : ])
		}
	}
}

func letterCombinations(digits string) []string {
	phone = make(map[string]string)
	phone["2"] = "abc"
	phone["3"] = "def"
	phone["4"] = "ghi"
	phone["5"] = "jkl"
	phone["6"] = "mno"
	phone["7"] = "pqrs"
	phone["8"] = "tuv"
	phone["9"] = "wxyz"
	if len(digits) != 0 {
		backtrack("", digits)
	}
	return output
}
