package main

import (
	"fmt"
	"math"
)

//func main() {
//	// 当i=5时跳出for循环
//	for i := 0; i < 10; i++ {
//		if i == 5 {
//			break // 跳出for循环
//		}
//		fmt.Println(i)
//	}
//	fmt.Println("over")
//}

func main() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换！
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
