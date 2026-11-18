# Ep 4 · Capital letter = public

📖 [A Tour of Go: Exported names](https://go.dev/tour/basics/3) · Go spec: [Exported identifiers](https://go.dev/ref/spec#Exported_identifiers)

In Go, a name is **exported** if it starts with a capital letter: other packages can use it. A name that starts with a lower-case letter is **not exported**: only its own package can use it. There's no `public` or `private` keyword.

This module has two packages. `greet` exports `Hello` and keeps `secret` to itself; `Hello` calls `secret` from inside the package, which is fine.

```sh
go run .
```

```
Hello, Go
```

Three folders go with this episode:

| Folder | What happens |
| --- | --- |
| `l1e04-exported` (this one) | `main` calls `greet.Hello()` → prints `Hello, Go` |
| [`l1e04-unexported`](../l1e04-unexported) | `main` calls `greet.secret()` → `name secret not exported by package greet` |
| [`l1e04-println`](../l1e04-println) | Ep 3's puzzle: `fmt.println` → `undefined: fmt.println (but have Println)` |

## Try it

- The Tour's own example: print `math.pi`, then `math.Pi`. The first one fails with `undefined: math.pi (but have Pi)`.
- Rename `secret` to `Secret` (in both places in `greet.go`). Now `main` can call `greet.Secret()`.

The exact rule (Go spec): an identifier is exported if its first character is a Unicode upper-case letter **and** it's declared at package level, or it's a field or method name. A variable named `X` inside a function isn't exported: it's local.
