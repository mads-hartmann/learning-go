```sh
# Running
go run hello.go

# Compiling
go build -o hello_world hello.go 

# Formatting
go fmt hello.go

# -l to print files that are incorrectly formatted
# -w to modify files in-place
goimports -l -w . 

# Linting
golint ./...

# Vetting
go vet hello.go
go vet ./...

# Staticcheck (Used by the golang.go vscode extension)
staticcheck ./...
```
