// Adapted from A Tour of Go, "Goroutines" (https://go.dev/tour/concurrency/1).
// Copyright 2009 The Go Authors. BSD-3-Clause, see LICENSE-GO in this folder.
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
