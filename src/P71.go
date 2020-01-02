package main

import "fmt"

func main() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	res := []byte(path)
	for i := 0; i < len(res); {
		if res[i] == '.' && res[i + 1] == '.' {
			tmp := 0
			for j := 0; j < i - 1; j ++ {
				if res[j] == '/' {
					tmp = j
				}
			}
			for j := tmp; j < i + 2; j ++ {
				res[j] = ' '
			}
			i = i + 3
			continue
		}
		if res[i] == '.' {
			res[i] = ' '
			res[i - 1] = ' '
			i = i + 2
			continue
		}
		if i != len(res) - 1 && res[i] == '/' && res[i + 1] == '/' {
			res[i + 1] = ' '
			i = i + 2
			continue
		}
		i ++
	}
	if len(res) != 1 && res[len(res) - 1] == '/' {
		res[len(res) - 1] = ' '
	}
	return string(res)
}

