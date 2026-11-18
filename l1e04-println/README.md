# Ep 4 · One letter

Part of [Ep 4](../l1e04-exported), and the answer to Ep 3's puzzle. It **doesn't build**, on purpose:

```
# example/hello
./hello.go:6:6: undefined: fmt.println (but have Println)
```

The fix is one letter: `fmt.Println`. Package `fmt` exports `Println`, with a capital P.
