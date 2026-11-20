# Ep 7 · `int + float64`? No

Go doesn't mix numeric types for you. `i + f`, with `i` an `int` and `f` a `float64`, doesn't build (see [l1e07-mismatched](../l1e07-mismatched)). You convert one side yourself, with `T(v)`:

```go
i := 1
f := 2.5
fmt.Println(float64(i) + f) // 3.5
fmt.Println(i + int(f))     // 3: int(f) discards the fraction
```

```sh
go run .
```

```
3.5
3
```

Converting a floating-point number to an integer discards the fraction (truncation towards zero), so `int(-2.5)` is `-2`.

The exception is **untyped constants**, like a plain `1`: they take the type their context needs. See [l1e07-untyped](../l1e07-untyped).

Try it: add `fmt.Println(i + 2.5)`. An untyped constant still has to fit its new type, and `2.5` can't be an `int`, so Go says `2.5 (untyped float constant) truncated to int`.

Sources: [A Tour of Go: Type conversions](https://go.dev/tour/basics/13) · [Numeric Constants](https://go.dev/tour/basics/16) · Go spec: [Operators](https://go.dev/ref/spec#Operators), [Conversions](https://go.dev/ref/spec#Conversions), [Constants](https://go.dev/ref/spec#Constants).
