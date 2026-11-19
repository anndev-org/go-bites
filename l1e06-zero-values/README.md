# Ep 6 · No uninitialized variables

📖 [A Tour of Go: Zero values](https://go.dev/tour/basics/12) · Go spec: [The zero value](https://go.dev/ref/spec#The_zero_value)

In Go, a variable declared without a value isn't left uninitialized: it starts at its **zero value**.

| Type | Zero value |
| --- | --- |
| numbers (`int`, `float64`, …) | `0` |
| `string` | `""` (empty) |
| `bool` | `false` |
| pointers, slices, maps, channels, functions, interfaces | `nil` |

```sh
go run .
```

```
0 "" false <nil>
```

`%q` prints the empty string as `""`, so you can see it. `fmt` prints a nil pointer as `<nil>`.

Two folders go with this episode:

| Folder | What happens |
| --- | --- |
| `l1e06-zero-values` (this one) | Four types, no values → prints `0 "" false <nil>` |
| [`l1e06-var-n`](../l1e06-var-n) | Ep 5's puzzle: `var n int`, then print it → `0` |

## Try it

- Add `var f float64` and print it. Is it `0` or `0.0`?
- Add `var b []int` and print `b == nil`.
