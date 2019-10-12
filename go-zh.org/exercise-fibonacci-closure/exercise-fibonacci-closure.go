package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	index, before, cur := 1, 0, 1
	return func() int {
		if index == 1 {
			index++
			return before
		} else if index == 2 {
			index++
			return cur
		} else {
			temp := before
			before = cur
			cur += temp
			return cur
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
