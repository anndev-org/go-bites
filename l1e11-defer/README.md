# Ep 11 · `defer`: last in, first out

A `defer` statement schedules a function call to run just before the surrounding function returns. If there are several, they run in reverse order: the last one deferred runs first.

```go
func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("three")
}
// three
// two
// one
```

```sh
go run .
```

`three` prints right away. The two deferred calls wait until `main` returns, then run newest first: `two`, then `one`. (This is Ep 10's puzzle.)

The everyday use is cleanup: `defer f.Close()` right after opening a file, or `defer mu.Unlock()` right after locking, so it runs on every return path.

Try it: add `defer` in front of `fmt.Println("three")` too. Does the output change?

More from this episode: [l1e11-args](../l1e11-args) (arguments are evaluated when `defer` runs).

Sources: [A Tour of Go: Defer](https://go.dev/tour/flowcontrol/12), [Stacking defers](https://go.dev/tour/flowcontrol/13) · Go spec: [Defer statements](https://go.dev/ref/spec#Defer_statements) · [Effective Go: Defer](https://go.dev/doc/effective_go#defer).
