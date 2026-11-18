# Ep 5 · `:=` vs `var`

📖 [A Tour of Go: Short variable declarations](https://go.dev/tour/basics/10) · Go spec: [Short variable declarations](https://go.dev/ref/spec#Short_variable_declarations), [Variable declarations](https://go.dev/ref/spec#Variable_declarations)

Inside a function, these three lines declare the same thing: an `int` holding 42.

- `var a int = 42`: name, type, value.
- `var b = 42`: no type, so Go infers it from the value. An untyped integer constant like `42` becomes an `int`.
- `c := 42`: the short variable declaration. It's shorthand for `var c = 42`, and it works **only inside functions**.

```sh
go run .
```

```
42 42 42
```

Three folders go with this episode:

| Folder | What happens |
| --- | --- |
| `l1e05-short-decl` (this one) | Three declarations of 42 → prints `42 42 42` |
| [`l1e05-outside`](../l1e05-outside) | Ep 4's puzzle: `x := 42` outside a function → `syntax error: non-declaration statement outside function body` |
| [`l1e05-redeclare`](../l1e05-redeclare) | `x := 2` after `x := 1` → `no new variables on left side of :=` |

## Try it

- Add `fmt.Printf("%T %T %T\n", a, b, c)`. All three are `int`.
- Change `c := 42` to `c := 4.2` and print `%T` again. Now it's `float64`.

The exact rule (Go spec): `:=` can redeclare a variable, but only in a multi-variable declaration where at least one variable on the left is new (as in `v, err := ...`). That comes up in a later episode.
