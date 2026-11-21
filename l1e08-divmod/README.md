# Ep 8 · Your own function with two results

Part of [Ep 8](../l1e08-atoi).

```go
func divmod(a, b int) (int, int) {
	return a / b, a % b
}
```

The result types go in parentheses after the parameters, and `return` lists one value for each. The caller takes them all at once:

```go
q, r := divmod(7, 2) // 3 1
```
