# Go Bites

The code from **Go Bites**, a short-video Go course for developers on TikTok [@anndev.org](https://www.tiktok.com/@anndev.org). It follows [A Tour of Go](https://go.dev/tour) and the official docs at [go.dev/doc](https://go.dev/doc/), using **Go 1.27**.

There's one folder per episode, and each one is a standalone module:

```sh
cd l1e01-why-go
go run .
```

Some episodes show code that **intentionally doesn't build**, so you can see Go's error for yourself. Their README says so, and `expected.txt` holds the error text you should get.

## Episodes

### Level 1: Go Basics

| Ep | Title | Source | Code |
| --- | --- | --- | --- |
| 1 | Why does Go exist? | [Go FAQ: Origins](https://go.dev/doc/faq#Origins) | [l1e01-why-go](l1e01-why-go) |
| 2 | Hello, Go | [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started) | [l1e02-hello-go](l1e02-hello-go) |

## Checking the examples

`./check.sh` checks every episode with the current Go release. Each one must be `gofmt`-clean and pass `go vet` and `go build`, print exactly its `output.txt` if it has one, or fail with the error text in its `expected.txt`. CI runs this every week, so the examples stay true on the latest Go.

## License

Our own code is under the [BSD Zero Clause License](LICENSE) (0BSD): copy it into anything, no notice needed. Examples adapted from A Tour of Go or other Go documentation stay under the Go project's BSD-3-Clause license; those folders say so at the top of the file and include `LICENSE-GO`.

Not affiliated with the Go project or Google. Go is a trademark of Google LLC.
