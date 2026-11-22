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
| 3 | Unused variable? Won't compile | [Go FAQ: unused variables and imports](https://go.dev/doc/faq#unused_variables_and_imports) | [l1e03-unused-var](l1e03-unused-var), [l1e03-unused-import](l1e03-unused-import), [l1e03-blank](l1e03-blank) |
| 4 | Capital letter = public | [A Tour of Go: Exported names](https://go.dev/tour/basics/3) | [l1e04-exported](l1e04-exported), [l1e04-unexported](l1e04-unexported), [l1e04-println](l1e04-println) |
| 5 | `:=` vs `var` | [A Tour of Go: Short variable declarations](https://go.dev/tour/basics/10) | [l1e05-short-decl](l1e05-short-decl), [l1e05-outside](l1e05-outside), [l1e05-redeclare](l1e05-redeclare) |
| 6 | No uninitialized variables | [A Tour of Go: Zero values](https://go.dev/tour/basics/12) | [l1e06-zero-values](l1e06-zero-values), [l1e06-var-n](l1e06-var-n) |
| 7 | `int + float64`? No | [A Tour of Go: Type conversions](https://go.dev/tour/basics/13) | [l1e07-conversions](l1e07-conversions), [l1e07-mismatched](l1e07-mismatched), [l1e07-untyped](l1e07-untyped) |
| 8 | Two return values | [A Tour of Go: Multiple results](https://go.dev/tour/basics/6) | [l1e08-atoi](l1e08-atoi), [l1e08-divmod](l1e08-divmod), [l1e08-blank](l1e08-blank), [l1e08-mismatch](l1e08-mismatch) |
| 9 | Go has no `while` | [A Tour of Go: For is Go's "while"](https://go.dev/tour/flowcontrol/3) | [l1e09-while](l1e09-while), [l1e09-classic](l1e09-classic), [l1e09-forever](l1e09-forever) |

## Checking the examples

`./check.sh` checks every episode with the current Go release. Each one must be `gofmt`-clean and pass `go vet` and `go build`, print exactly its `output.txt` if it has one, or fail with the error text in its `expected.txt`. CI runs this every week, so the examples stay true on the latest Go.

## License

Our own code is under the [BSD Zero Clause License](LICENSE) (0BSD): copy it into anything, no notice needed. Examples adapted from A Tour of Go or other Go documentation stay under the Go project's BSD-3-Clause license; those folders say so at the top of the file and include `LICENSE-GO`.

Not affiliated with the Go project or Google. Go is a trademark of Google LLC.
