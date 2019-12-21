package main

import "fmt"

//电话号码的字母组合
func main(){
	digits := "23"
	fmt.Println(letterCombinations1(digits))
}

var phone1 map[string]string
var output1 []string

func backtrack1(combination, next_digits string) {
	if len(next_digits) == 0 {
		output1 = append(output1, combination)
	}else {
		letters := phone1[next_digits[0:1]]
		for i := 0; i < len(letters); i ++ {
			backtrack1(combination + string(letters[i]), next_digits[1 : ])
		}
	}
}

func letterCombinations1(digits string) []string {
	phone1 = make(map[string]string)
	phone1["2"] = "abc"
	phone1["3"] = "def"
	phone1["4"] = "ghi"
	phone1["5"] = "jkl"
	phone1["6"] = "mno"
	phone1["7"] = "pqrs"
	phone1["8"] = "tuv"
	phone1["9"] = "wxyz"
	if len(digits) != 0 {
		backtrack1("", digits)
	}
	return output1
}

