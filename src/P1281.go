package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(4421))
}

func subtractProductAndSum(n int) int {
	ji, he := n % 10, n % 10
	n = (n - he) / 10
	for n > 0 {
		tmp := n % 10
		ji = ji * tmp
		he = he + tmp
		n = (n - tmp) / 10
	}
	return ji - he
}