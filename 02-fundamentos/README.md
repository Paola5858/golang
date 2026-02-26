# fundamentals

core go concepts: operators, control flow, data structures, and user input.

## modules

### operators
- arithmetic: `+`, `-`, `*`, `/`, `%`
- logical: `&&`, `||`, `!`
- comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`

### control structures
- conditionals: `if/else`, `switch`
- loops: `for`, `range`
- break and continue

### data structures

#### arrays and slices
```go
// array: fixed size
var arr [5]int

// slice: dynamic size
slice := []int{1, 2, 3}
slice = append(slice, 4)
```

#### maps
```go
// key-value pairs
capitals := map[string]string{
    "SP": "São Paulo",
    "RJ": "Rio de Janeiro",
}
```

### user input
```go
var name string
fmt.Scan(&name)
```

## practical projects

### banco da paola
banking system demonstrating:
- state management with variables
- conditional logic for transactions
- input validation
- basic error handling

### login glamouroso
authentication system covering:
- user credential storage
- input processing
- validation logic
- access control

## running examples

each subdirectory contains a focused example:

```bash
cd arrays-slice
go run .
```

## key concepts

- go has no while loop, only `for`
- slices are references to arrays
- maps must be initialized before use
- `range` provides index and value
- `fmt.Scan` reads from stdin

## testing

```bash
# run all tests in module
go test ./...

# run with verbose output
go test -v ./...
```
