package main

import (
	"fmt"
	"os"
)

//判断两个字符串是否相对打乱
func strJudge(sa, sb string) bool {
	if len(sa) != len(sb) {
		return false
	}
	if sa == sb {
		fmt.Println("字符串相等")
		return false
	}
	b1 := []byte(sa)
	b2 := []byte(sb)
	if string(byteSort(b1)) == string(byteSort(b2)) {
		fmt.Println("字符串相对打乱")
		return true
	} else {
		fmt.Println("字符串包含不同字符")
		return false
	}
	return false
}

//对byte类型slice进行排序，返回排序之后的切片
func byteSort(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return b
}

func main() {
	slice := make([]string, 2)
	if len(os.Args) != 3 {
		fmt.Printf("请输入正确的参数\n")
	}
	for i, args := range os.Args[1:] {
		slice[i] = args
	}
	fmt.Println(strJudge(slice[0], slice[1]))
}
