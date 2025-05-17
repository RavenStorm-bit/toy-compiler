package compiler

import (
	"github.com/RavenStorm-bit/toy-compiler/ast"
	"github.com/RavenStorm-bit/toy-compiler/bytecode"
)

// Compiler traverses the AST and generates bytecode
type Compiler struct {
	instructions bytecode.Instructions
	constants    []interface{}
}

// New creates a new Compiler instance
func New() *Compiler {
	return &Compiler{
		instructions: bytecode.Instructions{},
		constants:    []interface{}{},
	}
}

// Compile generates bytecode from an AST node
func (c *Compiler) Compile(node ast.Node) error {
	// TODO: Implement AST traversal and bytecode generation
	return nil
}

// Bytecode returns the compiled bytecode
func (c *Compiler) Bytecode() *bytecode.Bytecode {
	return &bytecode.Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}
