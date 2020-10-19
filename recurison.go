package main

import "fmt"

func main() {
	fmt.Println(1 * 2 * 3 * 4 * 5 * 6 * 7)
	fmt.Println(fact(7))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
