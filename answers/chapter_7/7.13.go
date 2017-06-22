/*
假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？
*/
package main

import (
	"fmt"
)

func main() {
	str := "lsjxiweslkdj"
	newStr := str[len(str)/2:] + str[:len(str)/2]
	fmt.Printf(newStr)
}
