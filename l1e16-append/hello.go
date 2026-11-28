package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	s := a[0:2]
	s = append(s, 9)
	fmt.Println(a)
}
