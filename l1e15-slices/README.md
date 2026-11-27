# Ep 15 · Arrays vs slices

`a` is an array: `[3]int`, three ints, and the length is part of its type. `s := a[0:2]` is a slice of it, covering elements 0 and 1 (the high bound is excluded). A slice doesn't store any data of its own: it describes a section of the array. So writing `s[0] = 9` changes `a`. (This is Ep 14's puzzle.)

```go
func main() {
	a := [3]int{1, 2, 3}
	s := a[0:2]
	s[0] = 9
	fmt.Println(a)
}
// [9 2 3]
```

```sh
go run .
```

Try it: change `s := a[0:2]` to `s := a`. What prints now?

More from this episode: [l1e15-copy](../l1e15-copy) (assigning an array copies it; slicing it doesn't).

Sources: [A Tour of Go: Arrays](https://go.dev/tour/moretypes/6), [Slices](https://go.dev/tour/moretypes/7), [Slices are like references to arrays](https://go.dev/tour/moretypes/8) · Go spec: [Array types](https://go.dev/ref/spec#Array_types), [Slice types](https://go.dev/ref/spec#Slice_types), [Slice expressions](https://go.dev/ref/spec#Slice_expressions)
