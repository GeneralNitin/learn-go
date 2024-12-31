### Content

1. How go is different and similar from other programming languages
2. Install go Locally
3. Write our first program in GO
4. How Go lang works, it's lexer, variables, conditional, loops and functions
5. Structs, web programming, create and decode json, how to make web request, and how to accept web requests
6. How to build REST APIs, after the basic APIs, Using MongoDB
7. Concurrent programming, sync package, go routine, mutex and channels.

#### Notes
- command to run the file
  - `go run main.go`
- help
  - go help
  - go help <topic>
  - go help gopath
  - go env GOPATH

- Go Lexer
  - golang.org/doc
  - golang.org/ref/spec

- Types
  - String
  - Bool
  - Integer    | uint8, uint64, int8, int64, uintptr
  - Floating   | float32 float64
  - Complex
  - Array
  - Slices
  - Maps
  - Structs
  - Pointers
  - Functions
  - Channels

- Alias
  - byte alias for uint8
  - rune alias for int32

- Walrus Operator: `:=`, not allowed outside the method level

- Resources:
  - go.dev
  - pkg.go.dev -- has all the packages

- GO ENV
  - go evn -- command
  - GOOS="windows" go build
  - GOOS="linux" GOARCH="arm64" go build

- Memory Management
  - new()
    - Allocate memory but no INIT
    - you will get a memory address
    - zeroed storage
  - make()
    - Allocate memory and INIT
    - you will get a memory address
    - non-zeroed storage
  - pkg.go.dev/runtime