package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{0,0}
	fmt.Println(largestNumber(nums))
}

type bysort [][]byte

func (b bysort) Len() int {
	return len(b)
}

func (b bysort) Less(i, j int) bool {
	length := len(b[i])
	if length > len(b[j]) {
		length = len(b[j])
	}
	for m := 0; m < length; m ++ {
		if b[i][m] == b[j][m] {
			continue
		}else {
			return b[i][m] < b[j][m]
		}
	}
	if length == len(b[i]) {
		tmp1 := append(b[i], b[j][length : ]...)
		tmp2 := append(b[j][length : ], b[i]...)
		for m := 0; m < len(tmp1); m ++ {
			if tmp1[m] == tmp2[m] {
				continue
			}else {
				return tmp1[m] < tmp2[m]
			}
		}
	}else {
		tmp2 := append(b[j], b[i][length : ]...)
		tmp1 := append(b[i][length : ], b[j]...)
		for m := 0; m < len(tmp1); m ++ {
			if tmp1[m] == tmp2[m] {
				continue
			}else {
				return tmp1[m] < tmp2[m]
			}
		}
	}
	return false
}

func (b bysort) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func largestNumber(nums []int) string {
	var record bysort
	var res []byte
	for i := 0; i < len(nums); i ++ {
		record = append(record, []byte(strconv.Itoa(nums[i])))
	}
	sort.Sort(record)
	i := len(record) - 1
	for i > 0 && record[i][0] == '0' {
		i --
	}
	for ; i >= 0; i -- {
		res = append(res, record[i]...)
	}
	return string(res)
}
