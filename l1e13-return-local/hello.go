package main

import "fmt"

func newCount() *int {
	n := 7
	return &n
}

func main() {
	fmt.Println(*newCount())
}
