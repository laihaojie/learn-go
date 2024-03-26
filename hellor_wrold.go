package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	sum := 0
	// 循环
	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	fmt.Println(quote.Go())

	fmt.Println(Hello("world"))

	var message string = "Hello, World!"

	fmt.Println(message)

	message1 := "Hello, World!"

	fmt.Println(message1)
}
