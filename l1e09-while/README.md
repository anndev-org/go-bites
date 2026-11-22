# Ep 9 · Go has no `while`

Go has only one loop keyword: `for`. With just a condition, it's Go's `while`.

```go
n := 1
for n < 100 {
	n *= 2
}
fmt.Println(n) // 128
```

```sh
go run .
```

The condition is checked before each pass: 1, 2, 4, … 64 is still under 100, so `n` doubles once more to 128, and then the loop stops. (This is Ep 8's puzzle.)

Try it: change `100` to `1000`. What does it print now?

More from this episode: [l1e09-classic](../l1e09-classic) (init; condition; post), [l1e09-forever](../l1e09-forever) (no condition at all). `for` also has a `range` form; it comes with slices, maps and strings later in the series.

Sources: [A Tour of Go: For](https://go.dev/tour/flowcontrol/1), [For continued](https://go.dev/tour/flowcontrol/2), [For is Go's "while"](https://go.dev/tour/flowcontrol/3), [Forever](https://go.dev/tour/flowcontrol/4) · Go spec: [For statements](https://go.dev/ref/spec#For_statements), [Keywords](https://go.dev/ref/spec#Keywords).
