# Ep 9 · A loop with no condition

Part of [Ep 9](../l1e09-while).

```go
for {
	fmt.Println("again")
}
```

With no condition, `for` loops forever. This program never ends on its own: press Ctrl-C to stop it. In real code you leave the loop with `break` or `return`.

Try it: add a counter and `break` when it reaches 3.
