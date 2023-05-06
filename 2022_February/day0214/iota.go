package main

import "fmt"

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	fmt.Println("a:", a, "b:", b)
	fmt.Println("c:", c, "d:", d)
	fmt.Println("e:", e, "f:", f)
}
