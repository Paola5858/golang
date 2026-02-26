# go learning path

![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)

structured learning repository covering go fundamentals through practical examples and projects.

## why this repository?

this isn't just a collection of code snippets. each module:
- demonstrates a specific concept with working code
- includes comprehensive documentation
- provides practical examples and use cases
- follows go best practices and conventions

ideal for developers learning go or reviewing fundamentals.

## structure

```
01-introducao/          variables, types, constants
02-fundamentos/         operators, control flow, data input, collections
03-escopo/              variable and function scope
04-funcoes/             functions, multiple returns
05-ponteiros/           pointers and memory references
06-structs/             custom data structures
jogoasteroide/          practical project: asteroid game with ebiten
```

## quick start

first time? see [SETUP.md](SETUP.md) for detailed installation guide.

```bash
# navigate to any module
cd 01-introducao

# run the code
go run .
```

## core concepts

### basics
- type system (string, int, float, bool)
- variable declaration and inference
- constants and iota

### control structures
- conditionals (if/else, switch)
- loops (for, range)
- defer, panic, recover

### functions
- multiple return values
- named returns
- variadic functions
- closures

### data structures
- arrays and slices
- maps
- structs and methods
- interfaces

### memory management
- pointers and references
- value vs pointer receivers
- memory allocation

## practical projects

### banco da paola
basic banking system demonstrating:
- state management
- input validation
- error handling
- control flow

### login glamouroso
authentication system covering:
- user input processing
- credential validation
- conditional logic

### jogo asteroides
full game implementation using ebiten:
- game loop architecture
- collision detection
- sprite rendering
- score system
- entity management

## development setup

```bash
# check go installation
go version

# run tests
go test ./...

# format code
go fmt ./...

# run static analysis
go vet ./...
```

### using makefile

```bash
make test      # run all tests
make fmt       # format code
make vet       # static analysis
make run-all   # execute all modules
```

## learning resources

- [official go documentation](https://go.dev/doc/)
- [effective go](https://go.dev/doc/effective_go)
- [go by example](https://gobyexample.com/)
- [go tour](https://go.dev/tour/)

## progress tracking

- [x] basic syntax and types
- [x] control structures
- [x] functions and scope
- [x] pointers
- [x] structs
- [x] unit testing basics
- [ ] interfaces
- [ ] concurrency (goroutines, channels)
- [ ] error handling patterns
- [ ] advanced testing (mocks, benchmarks)

---

## contributing

contributions are welcome! see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## license

this project is licensed under the MIT license - see [LICENSE](LICENSE) for details.

## acknowledgments

- [go documentation](https://go.dev/doc/) for comprehensive language reference
- [effective go](https://go.dev/doc/effective_go) for best practices
- [ebiten](https://ebiten.org/) for game development framework

active learning repository. updated as concepts are mastered.
