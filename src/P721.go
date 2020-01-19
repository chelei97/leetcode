package main

import "fmt"

func main() {
	account := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	fmt.Println(accountsMerge(account))
}

func accountsMerge(accounts [][]string) [][]string {
	union := make([]int, len(accounts))
	//并查集
	for i := 0; i < len(union); i ++ {
		union[i] = i
	}
	
}

func unionSer() {

}


