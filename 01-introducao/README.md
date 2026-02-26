# introduction to go

first contact with go: variable declaration, type system, and constants.

## concepts covered

### variable declaration
```go
var name string = "explicit type"
var age = 25              // type inference
shorthand := "quick"      // short declaration
```

### type system
- `string`: text data
- `int`: integers
- `float32/float64`: floating-point numbers
- `bool`: true/false

### constants
```go
const pi = 3.14159
const appName string = "MyApp"
```

## key takeaways

- go is statically typed but supports type inference
- `:=` can only be used inside functions
- constants are immutable and evaluated at compile time
- unused variables cause compilation errors

## run

```bash
go run main.go
```
