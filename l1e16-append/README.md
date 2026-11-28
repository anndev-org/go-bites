# Ep 16 · The append surprise

`s := a[0:2]` has length 2 but capacity 3: the array behind it has room for one more element. So `append(s, 9)` doesn't need a new array. It writes the 9 into the array, where the 3 was.

```go
func main() {
	a := [3]int{1, 2, 3}
	s := a[0:2]
	s = append(s, 9)
	fmt.Println(a)
}
// [1 2 9]
```

```sh
go run .
```

Try it: print `len(s)` and `cap(s)` after the append.

Sources: [A Tour of Go: Slice length and capacity](https://go.dev/tour/moretypes/11), [Appending to a slice](https://go.dev/tour/moretypes/15) · Go spec: [Appending to and copying slices](https://go.dev/ref/spec#Appending_and_copying_slices) ("Otherwise, append re-uses the underlying array.")
