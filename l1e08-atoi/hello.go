package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("42")
	fmt.Println(n, err)
	n, err = strconv.Atoi("abc")
	fmt.Println(n, err)
}
