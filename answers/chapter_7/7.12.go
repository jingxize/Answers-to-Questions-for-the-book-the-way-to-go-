/*
编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
*/
package main

import (
	"fmt"
)

func main() {
	i := 3
	str := "sldjfeiow"
	a, b := explodeStr(i, str)
	fmt.Printf(a)
	fmt.Printf("\n")
	fmt.Printf(b)
}

func explodeStr(i int, str string) (a string, b string) {
	str2 := []byte(str)
	a = string(str2[:i])
	b = string(str2[i:])
	return
}
