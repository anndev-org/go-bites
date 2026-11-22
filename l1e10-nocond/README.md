# Ep 10 · A switch with no value

Part of [Ep 10](../l1e10-switch).

```go
h := 15
switch {
case h < 12:
	fmt.Println("morning")
case h < 17:
	fmt.Println("afternoon")
}
// afternoon
```

Leave out the value and it's the same as `switch true`: each case is a condition, checked from top to bottom, and the first true one runs. It's a clean way to write a long chain of `if` / `else if`.

Try it: add a `default:` case that prints `evening`, then set `h` to 20.
