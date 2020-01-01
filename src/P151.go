package main

import (
	"bytes"
	"fmt"
)

func main()  {
	s := "  hello world!  "
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	srcBytes := []byte(s)
	bytesSlice := bytes.Split(srcBytes, []byte{' '})
	var dstBytes []byte
	for i := len(bytesSlice) - 1; i >= 0; i-- {
		// 剔除[]byte{''}的元素
		if len(bytesSlice[i]) == 0 {
			continue
		}
		dstBytes = append(dstBytes, bytesSlice[i]...)
		if i > 0 {
			dstBytes = append(dstBytes, ' ')
		}
	}

	l := len(dstBytes)
	if l != 0 {
		if dstBytes[l-1] == ' ' {
			dstBytes = dstBytes[:l-1]
		}
	}
	return string(dstBytes)
}

