# pointers in go

understanding memory references and pointer mechanics.

## what are pointers?

pointers store memory addresses instead of values directly. they enable:
- efficient memory usage (pass references, not copies)
- direct memory manipulation
- shared state between functions

## syntax

```go
var x int = 10
var ptr *int = &x    // ptr holds the address of x

fmt.Println(ptr)     // prints memory address
fmt.Println(*ptr)    // dereferences: prints 10
```

## operators

- `&` (address-of): gets the memory address of a variable
- `*` (dereference): accesses the value at a memory address

## pass by value vs pass by reference

### by value (default)
```go
func modify(x int) {
    x = 20  // only changes local copy
}
```

### by reference (using pointers)
```go
func modify(x *int) {
    *x = 20  // changes original variable
}
```

## when to use pointers

- modifying function parameters
- avoiding large struct copies
- implementing data structures (linked lists, trees)
- working with methods that mutate state

## common pitfalls

- nil pointer dereference (runtime panic)
- forgetting to initialize pointers
- unnecessary pointer usage for small types

## run

```bash
go run main.go
```
