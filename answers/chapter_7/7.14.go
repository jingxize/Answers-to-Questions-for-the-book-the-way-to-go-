/*
编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片）。

如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。

如果您想要反转 Unicode 编码的字符串，请使用 []int32 类型的切片。
 */
package main

import (
	"fmt"
)

func main() {
	a := "aldjk"
	c := explodeStr(a)
	fmt.Printf("%s", c)
}

func explodeStr(str string) (b string) {
	a := len(str)
	str2 := []byte(str)
	for c := 0; c < a/2; c++ {
		str2[c], str2[a-c-1] = str2[a-c-1], str2[c]
	}
	b = string(str2)
	return
}
