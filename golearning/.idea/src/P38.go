package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//var n int = 2
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {

	if n <= 0 {
		return ""
	} else if n == 1 {
		return "1"
	}else if n == 2 {
		return "11"
	}else {
		var result string  = "11"//11
		for i := 0; i < n - 2; i++ {//n-2
			var temp string = ""
			var count int = 1
			for j := 0; j < len(result); {
				if j <= len(result) - 2 && result[j] == result[j + 1] {
					for (j + count) <= len(result) - 1 && result[j] == result[j + count] {
						count ++
					}
					temp = temp + strconv.Itoa(count) + strconv.Itoa(int(result[j] - '0'))
					j = j + count
					count = 1
				}else {
					temp = temp + "1" + strconv.Itoa(int(result[j] - '0'))
					j ++
				}
			}
			result = temp
		}
		return result
	}
	return ""
}

