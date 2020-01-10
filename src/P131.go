package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition1(s))
}

func partition1(s string) [][]string {
	var res [][]string
	basic := []byte(s)
	var tmp []string
	for i := 0; i < len(s); i ++ {
		tmp = append(tmp, string(basic[i]))
	}
	res = append(res, tmp)
	for i := 0; i < len(s) - 1; i ++ {
		for j := i + 1; j < len(s); j ++ {

		}
	}
}
