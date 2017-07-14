//假设定义： type Integer int，完成 get() 方法的方法体: func (p Integer) get() int { ... }。
package main

import "fmt"

type newInt struct {
	newInt int
}

func (i *newInt) get() int {
	return i.newInt
}
func main() {
	data := new(newInt)
	data.newInt = 10
	fmt.Printf("%v", data.get())
}
