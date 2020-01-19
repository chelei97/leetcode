package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(52))
}

func convertToTitle(n int) string {
	var record [26]byte
	var res []byte
	record[0] = 'Z'
	for i := 1; i < 26; i ++ {
		record[i] = byte(i + int('A') - 1)
	}
	for n > 0 {
		tmp := n % 26
		res = append(res, record[tmp])
		if n == 26 {
			break
		}
		n = (n - tmp) / 26
	}
	for i := 0; i < len(res) / 2; i ++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
	}
	return string(res)
}
