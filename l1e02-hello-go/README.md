# Ep 2 · Hello, Go

A module with one main package, made the way the official tutorial does it:

```sh
go mod init example/hello   # writes go.mod: the module path and the minimum Go version
go run .                    # compiles the main package in this folder and runs it
```

```
Hello, World!
```

`go.mod` says `go 1.27.4` because that's the toolchain that ran `go mod init`. Yours will say whatever version you have.

## Try it

- Change `package main` to `package hello`, then run `go run .` again. You get `package example/hello is not a main package`. (`go build` still succeeds: it compiles the package as a library and throws the result away.)
- Rename `func main` to `func Main`. Now it won't build: `function main is undeclared in the main package`.

The rule behind both (Go spec, "Program execution"): the main package must have package name `main` and declare a function `main` that takes no arguments and returns no value. When `main` returns, the program exits.

Source: [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started). The code is adapted from it, so this folder keeps The Go Authors' BSD-3-Clause notice (`LICENSE-GO`).
