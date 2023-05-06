package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOk bool

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)

}
