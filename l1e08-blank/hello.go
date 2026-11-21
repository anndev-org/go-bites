package main

import "fmt"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, _ := divmod(7, 2)
	fmt.Println(q)
}
