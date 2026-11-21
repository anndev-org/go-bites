package main

import "fmt"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, r := divmod(7, 2)
	fmt.Println(q, r)
}
