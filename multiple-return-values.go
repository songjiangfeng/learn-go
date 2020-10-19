package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, d := vals()
	fmt.Println(d)
}

func vals() (int, int) {
	return 3, 7
}
