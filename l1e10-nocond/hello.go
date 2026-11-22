package main

import "fmt"

func main() {
	h := 15
	switch {
	case h < 12:
		fmt.Println("morning")
	case h < 17:
		fmt.Println("afternoon")
	}
}
