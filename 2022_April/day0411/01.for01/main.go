package main

import "fmt"

//func main() {
//	// 跳出多层循环；
//	for i := 0; i < 10; i++ {
//		var flag = false
//		for j := 'A'; j <= 'Z'; j++ {
//			if j == 'C' {
//				flag = true
//				break //跳出内循环
//			}
//			fmt.Printf("%v-%c\n", i, j)
//		}
//		if flag {
//			break //跳出外循环
//		}
//	}
//}

func main() {
	// 跳出多层循环；
	for i := 0; i < 10; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if j == 'C' {
				goto xx
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
xx: //lable 标签
	fmt.Println("over")
}
