# Ep 10 · `switch` without `break`

In Go, a `switch` runs only the matching case. The `break` you'd write at the end of each case in C or Java is automatic.

```go
switch n := 2; n {
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
}
// two
```

```sh
go run .
```

`n` is 2, so case 2 runs, and then the switch is done: case 3 never runs. (This is Ep 9's puzzle.)

Try it: add `fallthrough` after `fmt.Println("three")`. What does `go run .` say?

More from this episode: [l1e10-fallthrough](../l1e10-fallthrough) (run the next case too), [l1e10-nocond](../l1e10-nocond) (a switch with no value).

Sources: [A Tour of Go: Switch](https://go.dev/tour/flowcontrol/9), [Switch evaluation order](https://go.dev/tour/flowcontrol/10), [Switch with no condition](https://go.dev/tour/flowcontrol/11) · Go spec: [Expression switches](https://go.dev/ref/spec#Expression_switches), [Fallthrough statements](https://go.dev/ref/spec#Fallthrough_statements).
