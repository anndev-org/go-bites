# Ep 15 · Copy or window

Arrays are values: `b := a` copies all three elements, so changing `b` leaves `a` alone. A slice shares storage with its array: `s := a[0:2]` is a window onto `a`, so `s[1] = 8` changes `a[1]`.

```go
func main() {
	a := [3]int{1, 2, 3}
	b := a
	s := a[0:2]
	b[0] = 7
	s[1] = 8
	fmt.Println(a, b, s)
}
// [1 8 3] [7 2 3] [1 8]
```

```sh
go run .
```

Sources: [Effective Go: Arrays](https://go.dev/doc/effective_go#arrays) ("Arrays are values."), [Slices](https://go.dev/doc/effective_go#slices) · [A Tour of Go: Slices are like references to arrays](https://go.dev/tour/moretypes/8) · Go spec: [Slice types](https://go.dev/ref/spec#Slice_types)
