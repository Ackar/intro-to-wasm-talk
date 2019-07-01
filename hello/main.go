package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello WebAssembly!")
}

// GOOS=js GOARCH=wasm go build -o main.wasm
