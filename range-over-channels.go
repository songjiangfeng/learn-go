package main

import "fmt"

func main() {
	quenue := make(chan string, 2)
	quenue <- "one"
	quenue <- "two"
	close(quenue)
	for elem := range quenue {
		fmt.Println(elem)
	}
}
