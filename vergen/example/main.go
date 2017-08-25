package main

import (
	"fmt"

	"github.com/Akagi201/utilgo/vergen/example/version"
)

// go:generate go run vergen/main.go
func main() {
	fmt.Printf("This is version: %v\n", version.VgVersion)
	fmt.Printf("Commit         : %v\n", version.VgHash)
	fmt.Printf("Clean          : %v\n", version.VgClean)
}
