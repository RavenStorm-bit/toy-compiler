package main

import (
	"fmt"
	"github.com/RavenStorm-bit/toy-compiler/lexer"
	"github.com/RavenStorm-bit/toy-compiler/parser"
)

func main() {
	fmt.Println("Toy Compiler Demo")
	fmt.Println("=================")
	
	// Sample program
	input := `
	let x = 5;
	let y = 10;
	let add = fn(a, b) {
		return a + b;
	};
	
	let result = add(x, y);
	
	if (result > 10) {
		return "result is greater than 10";
	} else {
		return "result is 10 or less";
	}
	`
	
	// Create lexer and parser
	l := lexer.New(input)
	p := parser.New(l)
	
	// Parse the program
	program := p.ParseProgram()
	
	// Check for errors
	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		return
	}
	
	// Print the abstract syntax tree
	fmt.Println("\nAbstract Syntax Tree:")
	fmt.Println(program.String())
	
	// Print all the tokens
	fmt.Println("\nTokens:")
	l = lexer.New(input)
	for {
		tok := l.NextToken()
		fmt.Printf("Type: %-10s Literal: %s\n", tok.Type, tok.Literal)
		if tok.Type == "EOF" {
			break
		}
	}
}