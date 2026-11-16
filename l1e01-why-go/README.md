# Level 1 · Ep 1 — Why does Go exist?

📖 [Go FAQ: Origins](https://go.dev/doc/faq#Origins) · [A Tour of Go: Goroutines](https://go.dev/tour/concurrency/1)

The episode shows one line of this program: `go say("world")`. Putting `go` in front of a function call runs it as a **goroutine**, concurrently with the rest of `main`.

```sh
go run .
```

You'll see `hello` and `world` interleaved, and the order can change from run to run.

**Look closely:** sometimes there are only four `world`s. When `main` returns, the program ends, even if a goroutine hasn't finished. Level 3 covers how to wait for goroutines.

The example is from A Tour of Go, © The Go Authors, BSD-3-Clause (see `LICENSE-GO`).
