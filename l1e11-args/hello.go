package main

import "fmt"

func main() {
	count := 1
	defer fmt.Println("deferred:", count)
	count = 2
	fmt.Println("now:", count)
}
