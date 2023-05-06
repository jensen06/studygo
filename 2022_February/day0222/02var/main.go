package main

import "fmt"

var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	}
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
