# Ep 12 · No pointer arithmetic

Part of [Ep 12](../l1e12-pointers). This one **doesn't build**, on purpose:

```
# example/hello
./hello.go:8:2: invalid operation: p++ (non-numeric type *int)
```

`p++` means `p += 1`, and arithmetic only works on numbers. Unlike C, Go has no pointer arithmetic, so a pointer can't be moved to some other address. You can change what it points to only by giving it another pointer, like `p = &y`.

Sources: [A Tour of Go: Pointers](https://go.dev/tour/moretypes/1) ("Unlike C, Go has no pointer arithmetic.") · [Go FAQ: Why is there no pointer arithmetic?](https://go.dev/doc/faq#no_pointer_arithmetic) · Go spec: [IncDec statements](https://go.dev/ref/spec#IncDec_statements), [Arithmetic operators](https://go.dev/ref/spec#Arithmetic_operators).
