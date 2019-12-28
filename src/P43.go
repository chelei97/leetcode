package main

import "fmt"

func main() {
	fmt.Println(multi("408", 5, 0))
	fmt.Println(multiply("408", "5"))
}

func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	res := "0"
	for i := len(num2) - 1; i >= 0; i -- {
		tmp := multi(num1, int(num2[i] - '0'), len(num2) - 1 - i)
		res = add(res, tmp)
	}
	tmp := []byte(res)
	i := 0
	for ; i < len(res) && tmp[i] == '0'; i ++ {}
	if i == len(res) {
		return "0"
	}
	return string(tmp[i : ])
}

func add(m, n string) string {
	lenm, lenn := len(m), len(n)
	m, n = reverse(m), reverse(n)
	maxLen := max6(lenm, lenn)
	res := make([]byte, maxLen + 1)
	carry := 0
	for i := 0; i < maxLen; i ++ {
		add1 := 0
		add2 := 0
		if i < lenm {
			add1 = int(m[i] - '0')
		}
		if i < lenn {
			add2 = int(n[i] - '0')
		}
		tmp := add1 + add2 + carry
		carry = tmp / 10
		res[i] = byte('0' + (tmp % 10))
	}
	if carry != 0 {
		res[len(res) - 1] = '1'
		return reverse(string(res))
	}
	return reverse(string(res[ : maxLen]))
}

func max6(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func multi(num1 string, multiplicand, index int) string {
	carry := 0
	length := len(num1)
	res := make([]byte, length + index + 1)
	for i := 0; i < len(res); i ++ {
		res[i] = '0'
	}
	for i := length - 1; i >= 0; i -- {
		multiplier := int(num1[i] - '0')
		tmp := multiplicand * multiplier + carry
		res[length - 1 -i + index] = byte('0' + (tmp % 10))
		carry = tmp / 10
	}
	if carry != 0 {
		res[len(res) - 1] = byte('0' + carry)
		return reverse(string(res))
	}
	return reverse(string(res[ : length + index]))
}

func reverse(s string) string {
	s1 := []byte(s)
	length := len(s) - 1
	for i := 0; i < (length + 1) / 2; i ++ {
		s1[i], s1[length- i] = s1[length- i], s1[i]
	}
	return string(s1)
}