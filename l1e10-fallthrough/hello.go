package main

import "fmt"

func main() {
	switch n := 2; n {
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	}
}
