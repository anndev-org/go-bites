# Ep 13 · Return a pointer to a local? Fine

A Go function can return the address of its own local variable. The variable lives on after the function returns, so the pointer stays valid. (This is Ep 12's puzzle.)

```go
func newCount() *int {
	n := 7
	return &n
}

func main() {
	fmt.Println(*newCount())
}
// 7
```

```sh
go run .
```

When it can, the compiler keeps a function's local variables in that function's stack frame. Here the address of `n` leaves the function, so the compiler can't prove `n` is unused after `newCount` returns, and it puts `n` on the garbage-collected heap instead. This check is called escape analysis.

You can see the decision with the compiler's `-m` flag ("Print optimization decisions"):

```sh
go build -gcflags=-m
```

Among its lines (the others are about inlining and `fmt.Println`'s arguments):

```
./hello.go:6:2: moved to heap: n
```

Where a variable lives doesn't change what the program does. The Go FAQ: "From a correctness standpoint, you don't need to know."

Try it: change `newCount` to return `n` instead of `&n` (and its result type to `int`), then run `go build -gcflags=-m` again. Does `moved to heap: n` still appear?

Sources: [Effective Go: Composite literals](https://go.dev/doc/effective_go#composite_literals) ("unlike in C, it's perfectly OK to return the address of a local variable") · [Go FAQ: How do I know whether a variable is allocated on the heap or the stack?](https://go.dev/doc/faq#stack_or_heap) · [cmd/compile](https://pkg.go.dev/cmd/compile) (`-m`).
