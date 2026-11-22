package main

import "fmt"

func main() {
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n)
}
