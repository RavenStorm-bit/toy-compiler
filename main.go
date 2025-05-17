package main

import (
	"fmt"
	"os"
	"github.com/RavenStorm-bit/toy-compiler/repl"
)

func main() {
	fmt.Println("Welcome to the Toy Compiler!")
	fmt.Println("Type expressions to evaluate them.")
	repl.Start(os.Stdin, os.Stdout)
}
