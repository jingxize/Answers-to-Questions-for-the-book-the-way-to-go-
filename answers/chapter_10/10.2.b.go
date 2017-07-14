//定义： func f(i int) {}; var v Integer ，如何就 v 作为参数调用f？
package main

import "fmt"

type Integer int

func (Integer) f(i int) {
	fmt.Printf("use the par i :%v", i)
}

func main() {
	var v Integer
	v.f(10)
}
