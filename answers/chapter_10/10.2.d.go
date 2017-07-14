//定义： func f(i int) {}; var v Integer ，如何就 v 作为参数调用f？
package main

import "fmt"

type Integer struct {
	n int
}

func (i Integer) f() {
	fmt.Printf("use the par i :%v", i.n)
}

func main() {
	v := new(Integer)
	v.n = 10
	v.f()
}
