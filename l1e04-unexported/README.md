# Ep 4 · Not exported

Part of [Ep 4](../l1e04-exported). It **doesn't build**, on purpose: `main` calls `greet.secret()`, and `secret` starts with a lower-case letter.

```
# example/hello
./hello.go:10:20: name secret not exported by package greet
```
