/*
编写一个程序，要求能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组
*/
package main

import (
	"fmt"
)

func main() {
	a := [5]string{"a", "b", "b", "d", "d"}
	b := [len(a)]string{}
	for i := 0; i < len(a)-1; i++ {
		if a[i] != a[i+1] {
			b[i] = a[i+1]
		}
	}
	for _, str := range b {
		fmt.Printf(str)
	}

}
