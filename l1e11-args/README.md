# Ep 11 · deferred arguments are evaluated right away

Part of [Ep 11](../l1e11-defer).

```go
count := 1
defer fmt.Println("deferred:", count)
count = 2
fmt.Println("now:", count)
// now: 2
// deferred: 1
```

The call is deferred, but its arguments are evaluated when the `defer` statement runs. `count` was 1 at that moment, so the deferred call prints 1, even though `count` is 2 by the time it runs.

Sources: [A Tour of Go: Defer](https://go.dev/tour/flowcontrol/12) ("The deferred call's arguments are evaluated immediately") · Go spec: [Defer statements](https://go.dev/ref/spec#Defer_statements).
