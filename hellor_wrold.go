package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	sum := 0
	// 循环
	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}
