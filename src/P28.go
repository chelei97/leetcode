package main

import "fmt"

func main()  {
	var haystack, needle string = "hello", "b"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack) - len(needle); i++ {
		if haystack[i : i + len(needle)] == needle {
			return i
		}
	}
	return -1
}
