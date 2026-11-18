# Ep 5 · `:=` twice

Part of [Ep 5](../l1e05-short-decl). It **doesn't build**, on purpose:

```
# example/hello
./hello.go:5:4: no new variables on left side of :=
```

`:=` declares, so it needs at least one new variable on its left. To change `x`, use plain assignment: `x = 2`.

(`_ = x` is the blank identifier from Ep 3: it uses `x`, so the only error left is the one about `:=`.)
