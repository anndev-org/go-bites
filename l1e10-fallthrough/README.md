# Ep 10 · `fallthrough`

Part of [Ep 10](../l1e10-switch).

```go
switch n := 2; n {
case 2:
	fmt.Println("two")
	fallthrough
case 3:
	fmt.Println("three")
}
// two
// three
```

`fallthrough` is how you opt in to running the next case. It jumps straight into the next case's statements without checking that case's value: `three` prints even though `n` is 2. It must be the last statement in its case, and it can't be used in the last case of a switch.
