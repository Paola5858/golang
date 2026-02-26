# setup guide

quick start guide for setting up this repository.

## prerequisites

- go 1.21 or higher
- git
- text editor or ide (vscode recommended)

## installation

### 1. install go

download from [go.dev/dl](https://go.dev/dl/)

verify installation:
```bash
go version
```

### 2. clone repository

```bash
git clone https://github.com/Paola5858/golang.git
cd golang
```

### 3. verify setup

```bash
# run tests
go test ./...

# format code
go fmt ./...

# run a module
cd 01-introducao
go run .
```

## ide setup

### vscode (recommended)

1. install [go extension](https://marketplace.visualstudio.com/items?itemName=golang.go)
2. open repository folder
3. accept recommended extensions prompt
4. settings are pre-configured in `.vscode/settings.json`

### other editors

- **goland**: native go support
- **vim/neovim**: use vim-go plugin
- **sublime text**: use gosublime package

## running examples

### individual modules

```bash
cd 03-escopo
go run .
```

### all modules at once

```bash
make run-all
```

### with tests

```bash
make test
```

## troubleshooting

### "go: command not found"
add go to your PATH:
- **linux/mac**: add to `~/.bashrc` or `~/.zshrc`
- **windows**: add to system environment variables

### import errors in game module
```bash
cd jogoasteroide
go mod download
```

### formatting issues
```bash
go fmt ./...
```

## next steps

1. start with `01-introducao/`
2. read each module's README
3. run the code and experiment
4. modify examples to test understanding
5. check tests in `03-escopo/` for testing patterns

## getting help

- check module-specific READMEs
- review [go documentation](https://go.dev/doc/)
- open an issue for bugs or questions
