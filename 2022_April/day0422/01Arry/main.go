package main

import "fmt"

//func main() {
//	// 声明切片类型
//	var a []string              //声明一个字符串切片
//	var b = []int{}             //声明一个整型切片并初始化
//	var c = []bool{false, true} //声明一个布尔切片并初始化
//	// var d = []bool{false, true} //声明一个布尔切片并初始化
//	fmt.Println(a)        //[]
//	fmt.Println(b)        //[]
//	fmt.Println(c)        //[false true]
//	fmt.Println(a == nil) //true
//	fmt.Println(b == nil) //false
//	fmt.Println(c == nil) //false
//	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
//}

//func main() {
//	a := [5]int{1, 2, 3, 4, 5}
//	s := a[1:3] // s := a[low:high]
//	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
//}

//func main() {
//	a := [5]int{1, 2, 3, 4, 5}
//	s := a[1:3] // s := a[low:high]
//	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
//	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
//	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
//}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}
