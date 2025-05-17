package main

import (
	"fmt"
	"github.com/RavenStorm-bit/toy-compiler/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to the Toy Compiler!")
	fmt.Println("Type expressions to evaluate them.")
	repl.Start(os.Stdin, os.Stdout)
}
