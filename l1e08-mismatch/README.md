# Ep 8 · One variable, two values

Part of [Ep 8](../l1e08-atoi). This is Ep 7's puzzle, and it **doesn't build**, on purpose:

```
# example/hello
./hello.go:9:7: assignment mismatch: 1 variable but strconv.Atoi returns 2 values
```

`strconv.Atoi` returns two values, an `int` and an `error`, so you need two variables: `n, err := strconv.Atoi("42")`.
