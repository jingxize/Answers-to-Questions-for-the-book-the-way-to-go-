/*
编写一个程序，使用冒泡排序的方法排序一个包含整数的切片
*/
package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 7, 9, 5, 2}
	bubbleSort(arr[:])
}

func bubbleSort(a []int) {
	var i int = 0
	for j := 0; j < len(a)-1; j++ {
		for i = 0; i < len(a)-j-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	for _, x := range a {
		fmt.Print(x)
	}
}
