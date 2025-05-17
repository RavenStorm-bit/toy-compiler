package runner

import (
	"fmt"
	"io/ioutil"
	"github.com/RavenStorm-bit/toy-compiler/compiler"
	"github.com/RavenStorm-bit/toy-compiler/lexer"
	"github.com/RavenStorm-bit/toy-compiler/parser"
	"github.com/RavenStorm-bit/toy-compiler/vm"
)

// RunFile executes a source file
func RunFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", filename, err)
	}

	return RunSource(string(data))
}

// RunSource executes source code
func RunSource(source string) error {
	l := lexer.New(source)
	p := parser.New(l)

	// TODO: Parse the full program, not just an expression
	expression := p.ParseExpression()
	if len(p.Errors()) != 0 {
		return fmt.Errorf("parser errors: %v", p.Errors())
	}

	comp := compiler.New()
	err := comp.Compile(expression)
	if err != nil {
		return fmt.Errorf("compiler error: %w", err)
	}

	machine := vm.New(comp.Bytecode())
	err = machine.Run()
	if err != nil {
		return fmt.Errorf("vm error: %w", err)
	}

	// TODO: Get the result from the VM
	return nil
}
