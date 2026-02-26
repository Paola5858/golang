# functions in go

function declaration, multiple returns, and scope management.

## basic syntax

```go
func functionName(param1 type1, param2 type2) returnType {
    // function body
    return value
}
```

## multiple parameters of same type

```go
func add(x, y int) int {
    return x + y
}
```

## multiple return values

go functions can return multiple values, commonly used for error handling:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## named return values

```go
func calculate(x, y int) (sum, product int) {
    sum = x + y
    product = x * y
    return  // naked return
}
```

## variadic functions

accept variable number of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// usage
sum(1, 2, 3, 4, 5)
```

## function as values

functions are first-class citizens:

```go
func apply(fn func(int) int, value int) int {
    return fn(value)
}

double := func(x int) int { return x * 2 }
result := apply(double, 5)  // 10
```

## closures

functions can capture variables from outer scope:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
c()  // 1
c()  // 2
```

## defer

executes function after surrounding function returns:

```go
func example() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// prints: hello world
```

## best practices

- keep functions small and focused
- use multiple returns for error handling
- prefer named returns for complex functions
- document exported functions

## run

```bash
go run .
```
