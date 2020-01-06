package main

import "fmt"

func main() {
	fmt.Println(numTrees(18))
}

type region struct {
	begin int
	end int
}

var record map[region]int
func numTrees(n int) int {
	record = make(map[region]int, 0)
	r := region{
		begin : 1,
		end : n,
	}
	return helper7(r)
}

func helper7(r region) int {
	begin, end := r.begin, r.end
	if begin >= end {
		return 1
	}
	tmp := 0
	for i := begin; i <= end; i ++ {
		a, b := 0, 0
		t1 := region{
			begin: begin,
			end:   i - 1,
		}
		t2 := region{
			begin: i + 1,
			end:   end,
		}
		if k, ok := record[t1]; ok {
			a = k
		}else {
			a = helper7(t1)
			record[t1] = a
		}
		if k, ok := record[t2]; ok {
			b = k
		}else {
			b = helper7(t2)
			record[t2] = b
		}
		tmp = a * b + tmp
	}
	return tmp
}
