package main

import "fmt"

func main() {
	println("Hello, Mom!")

	var sum int = 0
	for i := 0; i < 90000; i++ {
		sum += i
	}

	fmt.Printf("the sum is %d", sum);
}