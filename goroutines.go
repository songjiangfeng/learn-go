package main

import (
	"fmt"
	"time"
)

func main() {
	go f("goroutine")

	f("direct")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)

	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}
