package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	str := make([]string, 0)
	for i := 0; i < len(split); i ++ {
		if split[i] == "" {
			continue
		}else if split[i] == "." {
			continue
		}else if split[i] == ".." {
			if len(str) == 0 {
				continue
			}else {
				str = str[ : len(str) - 1]
			}
		}else {
			str = append(str, split[i])
		}
	}
	res := strings.Join(str, "/")
	res = "/" + res
	return res
}

