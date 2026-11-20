# Ep 7 · Untyped constants adapt

Part of [Ep 7](../l1e07-conversions).

```go
f := 2.5
fmt.Println(f + 1) // 3.5
```

`1` is an **untyped constant**. When one operand is an untyped constant and the other isn't, the constant is converted to the other operand's type, here `float64`. So no `float64(1)` is needed.

It must fit that type, though: with `i := 1`, `i + 2.5` fails with `2.5 (untyped float constant) truncated to int`.
