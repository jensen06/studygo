package main

import "fmt"

//func main() {
//	a := make([]int, 2, 10)
//	fmt.Println(a)      //[0 0]
//	fmt.Println(len(a)) //2
//	fmt.Println(cap(a)) //10
//}

//func main() {
//	var s []int
//	s = append(s, 1)       // [1]
//	s = append(s, 2, 3, 4) // [1 2 3 4]
//	s2 := []int{5, 6, 7}
//	s = append(s, s2...) // [1 2 3 4 5 6 7]！
//	// fmt.Println(s)
//	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
//	// fmt.Println(s2)
//	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
//}

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
