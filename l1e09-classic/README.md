# Ep 9 · The classic `for`

Part of [Ep 9](../l1e09-while).

```go
for i := 0; i < 3; i++ {
	fmt.Println(i) // 0, 1, 2
}
```

Three parts: the init statement (runs once), the condition (checked before each pass) and the post statement (runs after each pass). There are no parentheses around them, and the braces are always required. Each part is optional: drop the init and post, and you have Go's `while`.

Try it: wrap the three parts in parentheses, C-style. What does `go run .` say?
