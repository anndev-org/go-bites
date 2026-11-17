# Ep 3 · Unused import

Part of [Ep 3](../l1e03-unused-var). It **doesn't build**, on purpose:

```
# example/hello
./hello.go:3:8: "os" imported and not used
```

To fix it, delete the import, or let a tool do it: [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) and gopls add missing imports and remove unused ones. To keep an import only for its side effects, the spec's way is a blank import: `import _ "some/package"`.
