# scope in go

understanding variable and function visibility across packages and blocks.

## scope levels

### package scope
variables declared outside functions are accessible throughout the package.

```go
var globalVar = "accessible everywhere in package"

func main() {
    fmt.Println(globalVar)
}
```

### function scope
variables declared inside functions are local to that function.

```go
func example() {
    localVar := "only here"
    // localVar dies when function returns
}
```

### block scope
variables in `if`, `for`, `switch` blocks are limited to that block.

```go
if x := 10; x > 5 {
    // x only exists here
}
// x is undefined here
```

## exported vs unexported

### exported (public)
starts with uppercase letter, accessible from other packages.

```go
func PublicFunction() {}
var PublicVar = 10
```

### unexported (private)
starts with lowercase letter, only accessible within the same package.

```go
func privateFunction() {}
var privateVar = 10
```

## shadowing

inner scope can shadow outer scope variables:

```go
x := 10
if true {
    x := 20  // different variable
    fmt.Println(x)  // prints 20
}
fmt.Println(x)  // prints 10
```

## best practices

- minimize global variables
- keep scope as narrow as possible
- avoid shadowing unless intentional
- use meaningful names to prevent confusion

## run

```bash
go run .
```

note: this module requires multiple files to demonstrate package-level scope.
