# Ep 7 · `int + float64`

Part of [Ep 7](../l1e07-conversions). It **doesn't build**, on purpose:

```
# example/hello
./hello.go:8:14: invalid operation: i + f (mismatched types int and float64)
```

`i := 1` makes an `int` and `f := 2.5` makes a `float64` (Ep 5: the type is inferred). For `+`, both operands must have the same type, and Go won't convert one for you. Write `float64(i) + f` or `i + int(f)` instead.
