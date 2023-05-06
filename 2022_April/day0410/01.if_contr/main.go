package main

import "fmt"

func main() {
	var age = 19
	if age > 18 {
		fmt.Println("成年了")
	} else if age > 7 {
		fmt.Println("上小学")
	} else {
		fmt.Println("最快乐的时光")
	}

	// for.1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for.2
	var i = 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	// for.3
	var j = 0
	for j < 10 {
		fmt.Println(i)
		j++
	}

	// for.4
	for {
		fmt.Println("无限循环")
	}

	// for.5
	s := "hello"
	fmt.Println(len(s))
	for i, v := range s {
		fmt.Println(i, v)

	}

}
