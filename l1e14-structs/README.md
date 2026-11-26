# Ep 14 · Structs, no classes

A struct is a collection of fields, each with a name and a type. You make one with a struct literal: list the values in field order, or name the fields you want. Fields you leave out start at their zero value. An `&` in front gives you a pointer to a new struct.

```go
type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}
	q := Point{X: 1}
	r := &Point{1, 2}
	fmt.Println(p, q, r)
}
// {1 2} {1 0} &{1 2}
```

```sh
go run .
```

`q` names only `X`, so `Y` is `0`. `r` is a `*Point`, and `fmt.Println` prints it with a leading `&`.

Go does not have classes: there's no `class` keyword and no type inheritance. Methods can be defined on types; that's Level 2.

Try it: change `q := Point{X: 1}` to `q := Point{1}`. What does `go run .` say?

More from this episode: [l1e14-pointer](../l1e14-pointer) (`p.X` through a pointer, Ep 13's puzzle).

Sources: [A Tour of Go: Structs](https://go.dev/tour/moretypes/2), [Struct Fields](https://go.dev/tour/moretypes/3), [Pointers to structs](https://go.dev/tour/moretypes/4), [Struct Literals](https://go.dev/tour/moretypes/5), [Methods](https://go.dev/tour/methods/1) ("Go does not have classes.") · Go spec: [Struct types](https://go.dev/ref/spec#Struct_types), [Composite literals](https://go.dev/ref/spec#Composite_literals) · Go FAQ: [Why is there no type inheritance?](https://go.dev/doc/faq#inheritance)
