# Ep 3 · Unused variable? Won't compile

📖 [Go FAQ: unused variables and imports](https://go.dev/doc/faq#unused_variables_and_imports) · Go spec: [Variable declarations](https://go.dev/ref/spec#Variable_declarations), [Import declarations](https://go.dev/ref/spec#Import_declarations)

This is Ep 2's puzzle. It **doesn't build**, on purpose:

```sh
go run .
```

```
# example/hello
./hello.go:4:6: declared and not used: x
```

`4:6` is line 4, column 6: the `x` in `var x = 42`.

Three folders go with this episode:

| Folder | What happens |
| --- | --- |
| `l1e03-unused-var` (this one) | a local variable that's never used → `declared and not used: x` |
| [`l1e03-unused-import`](../l1e03-unused-import) | an import that's never used → `"os" imported and not used` |
| [`l1e03-blank`](../l1e03-blank) | `_ = x` uses `x`, so it builds |

## Why so strict?

The Go FAQ gives two reasons: an unused variable may indicate a bug, and unused imports slow down compilation, an effect that grows as a program accumulates code and programmers. And the Go compiler doesn't report warnings, only errors that stop the build.

## Try it

- Move `var x = 42` above `func main`, so it's a package-level variable. Now it builds: the rule is about variables declared inside a function body.
- Use it: `fmt.Println(x)` (and `import "fmt"`). It builds and prints `42`.
- Write `x := 1` then `x = 2`. Assigning isn't using: it's still `declared and not used: x`.
