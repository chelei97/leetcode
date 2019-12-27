package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string, 0)
	for _, v := range strs {
		s := (by)(v)
		sort.Sort(s)
		result[string(s)] = append(result[string(s)], v)
	}
	res := make([][]string, 0)
	for _, v := range result {
		res = append(res, v)
	}
	return res
}

type by []byte

func (a by) Len() int {    	 // 重写 Len() 方法
	return len(a)
}
func (a by) Swap(i, j int){     // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a by) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[j] < a[i]
}

