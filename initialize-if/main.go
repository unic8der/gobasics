package main

import "fmt"

func main() {
	if n := foo(100); n != 0 {
		fmt.Println(n)

	}
}

func foo(number int) int {
	return number + 1
}
