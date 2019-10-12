package main

import (
	"fmt"
	"math"
)

/**
采用牛顿法计算平方根
@param x 需要求的值
@param n 循环次数
*/
func Sqrt(x float64, n int) float64 {
	z := float64(1)
	for i := 1; i <= n; i++ {
		//fmt.Println("第", i, "次循环 z = ", z)
		z2 := z * z
		if z2 == x {
			return z
		}
		//（*注：* 如果你对该算法的细节感兴趣，上面的 z² − x 是 z² 到它所要到达的值（即 x）的距离，
		// 除以的 2z 为 z² 的导数，我们通过 z² 的变化速度来改变 z 的调整量。
		// 这种通用方法叫做牛顿法。 它对很多函数，特别是平方根而言非常有效。）
		z -= (z2 - x) / (2 * z)
	}
	return z
}
func main() {
	fmt.Println(Sqrt(2, 20))
	fmt.Println(math.Sqrt(2))
}
