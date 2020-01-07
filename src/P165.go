package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	version1, version2 := "1.0", "1.0.0"
	fmt.Println(compareVersion(version1, version2))
}

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	length := max8(len(v1), len(v2))
	for i := 0; i < length; i ++ {
		t1, t2 := 0, 0
		if i < len(v1) {
			t1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			t2, _ = strconv.Atoi(v2[i])
		}
		if t1 > t2 {
			return 1
		}
		if t1 < t2 {
			return -1
		}
	}
	return 0
}

func max8(a, b int) int {
	if a < b {
		return b
	}
	return a
}
