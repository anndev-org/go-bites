# Ep 5 · Outside a function

Part of [Ep 5](../l1e05-short-decl), and the answer to Ep 4's puzzle. It **doesn't build**, on purpose:

```
# example/hello
./hello.go:3:1: syntax error: non-declaration statement outside function body
```

Outside a function, everything starts with a keyword (`package`, `import`, `var`, `func`, …), so `:=` isn't available there. The fix: `var x = 42`.
