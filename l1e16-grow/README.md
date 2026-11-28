# Ep 16 · A new array

The first `append` fits in the array's spare room. The second doesn't, so `append` allocates a new, bigger array, copies the elements across, and returns a slice of it. After that, `s` and `a` no longer share storage: `s[0] = 7` doesn't change `a`. That's why you always assign the result back: `s = append(s, x)`.

```go
func main() {
	a := [3]int{1, 2, 3}
	s := a[0:2]
	s = append(s, 9)
	s = append(s, 9)
	s[0] = 7
	fmt.Println(a, s)
}
// [1 2 9] [7 2 9 9]
```

```sh
go run .
```

How much bigger the new array is isn't in the spec; it's up to the runtime.

Sources: Go spec: [Appending to and copying slices](https://go.dev/ref/spec#Appending_and_copying_slices) ("append allocates a new, sufficiently large underlying array") · [A Tour of Go: Appending to a slice](https://go.dev/tour/moretypes/15) · [Effective Go: Append](https://go.dev/doc/effective_go#append) ("The result needs to be returned because … the underlying array may change.")
