# Ep 8 · Two return values

A Go function can return any number of results. [`strconv.Atoi`](https://pkg.go.dev/strconv#Atoi) returns two: the number, and an `error`.

```go
n, err := strconv.Atoi("42")
fmt.Println(n, err) // 42 <nil>
n, err = strconv.Atoi("abc")
fmt.Println(n, err) // 0 strconv.Atoi: parsing "abc": invalid syntax
```

```sh
go run .
```

When it works, `err` is `nil`. When the text isn't a number, `Atoi` returns `0` and an error that says why. Always check the error: Effective Go calls discarding it "terrible practice".

The second line uses `=`, not `:=`: `n` and `err` already exist (Ep 5).

Try it: change `"abc"` to `"99999999999999999999"`. That's too big for an `int`, so you get the largest `int` (`9223372036854775807` where `int` is 64 bits) and `value out of range`.

More from this episode: [l1e08-divmod](../l1e08-divmod) (your own function with two results), [l1e08-blank](../l1e08-blank) (skip one with `_`), [l1e08-mismatch](../l1e08-mismatch) (Ep 7's puzzle: one variable for two values).

Sources: [A Tour of Go: Multiple results](https://go.dev/tour/basics/6) · [Effective Go: Multiple return values](https://go.dev/doc/effective_go#multiple-returns), [The blank identifier in multiple assignment](https://go.dev/doc/effective_go#blank_assign) · Go spec: [Assignment statements](https://go.dev/ref/spec#Assignment_statements).
