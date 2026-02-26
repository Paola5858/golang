# structs in go

custom data types and methods for object-oriented patterns.

## what are structs?

structs are composite data types that group related fields together. they're go's way of creating custom types without traditional classes.

## defining structs

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

## field visibility

- **uppercase first letter**: exported (public)
- **lowercase first letter**: unexported (private to package)

```go
type User struct {
    Name     string  // public
    password string  // private
}
```

## creating instances

```go
// literal syntax
user := User{Name: "Alice", Age: 30}

// zero values
var user User  // Name: "", Age: 0

// partial initialization
user := User{Name: "Bob"}  // Age: 0
```

## methods

functions attached to struct types.

### value receiver
```go
func (p Person) Greet() string {
    return "Hello, " + p.Name
}
```

### pointer receiver
```go
func (p *Person) Birthday() {
    p.Age++  // modifies original struct
}
```

## when to use pointer receivers

- method needs to modify the struct
- struct is large (avoid copying)
- consistency (if one method uses pointer, use for all)

## struct embedding

go supports composition over inheritance:

```go
type Employee struct {
    Person        // embedded struct
    EmployeeID int
}
```

## run

```bash
go run main.go
```
