# Ep 12 · Pointers, but safe

A pointer holds the memory address of a value. `&x` gives you a pointer to `x`, and `*p` is the value the pointer points to. Writing through `*p` changes the original variable.

```go
func main() {
	x := 1
	p := &x
	*p = 2
	fmt.Println(x)
}
// 2
```

```sh
go run .
```

`p := &x` makes `p` a pointer to `x` (its type is `*int`). `*p = 2` follows the pointer and sets the variable it points to, which is `x`. So `x` is now 2. (This is Ep 11's puzzle.)

Try it: add `fmt.Println(*p)` just before `*p = 2`. What does it print?

More from this episode: [l1e12-no-arith](../l1e12-no-arith) (Go has no pointer arithmetic).

Sources: [A Tour of Go: Pointers](https://go.dev/tour/moretypes/1) · Go spec: [Pointer types](https://go.dev/ref/spec#Pointer_types), [Address operators](https://go.dev/ref/spec#Address_operators).
