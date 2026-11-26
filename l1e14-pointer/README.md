# Ep 14 · `p.X` through a pointer

`p` is a pointer to a struct (`*Point`). To reach its field `X` you could write `(*p).X`, but Go lets you write just `p.X`: the spec says `p.X` is shorthand for `(*p).X`. (This is Ep 13's puzzle: it builds and prints 9.)

```go
type Point struct{ X, Y int }

func main() {
	p := &Point{1, 2}
	p.X = 9
	fmt.Println(p.X)
}
// 9
```

```sh
go run .
```

Sources: [A Tour of Go: Pointers to structs](https://go.dev/tour/moretypes/4) · Go spec: [Selectors](https://go.dev/ref/spec#Selectors).
