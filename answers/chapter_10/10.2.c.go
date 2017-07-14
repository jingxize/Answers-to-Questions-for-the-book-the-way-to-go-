//假设 Integer 定义为 type Integer struct {n int}，
//完成 get() 方法的方法体：func (p Integer) get() int { ... }。
package main

type Integer struct {
	n int
}

func (p Integer) get() int {
	return p.n
}
