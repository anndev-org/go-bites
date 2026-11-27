package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a
	s := a[0:2]
	b[0] = 7
	s[1] = 8
	fmt.Println(a, b, s)
}
