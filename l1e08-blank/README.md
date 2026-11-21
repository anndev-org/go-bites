# Ep 8 · Skip a result with `_`

Part of [Ep 8](../l1e08-atoi).

```go
q, _ := divmod(7, 2) // 3
```

You need one variable for each result, but the **blank identifier** `_` (Ep 3) can stand in for one you don't need. Don't use it to skip an `error`: always check those.
