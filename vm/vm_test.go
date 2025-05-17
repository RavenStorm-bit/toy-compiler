package vm_test

import (
	"fmt"
	"testing"

	"github.com/RavenStorm-bit/toy-compiler/compiler"
	"github.com/RavenStorm-bit/toy-compiler/lexer"
	"github.com/RavenStorm-bit/toy-compiler/parser"
	"github.com/RavenStorm-bit/toy-compiler/vm"
)

func run(input string) (int64, error) {
	l := lexer.New(input)
	p := parser.New(l)
	expr := p.ParseExpression()
	if len(p.Errors()) != 0 || expr == nil {
		return 0, fmt.Errorf("parser errors: %v", p.Errors())
	}
	ins := compiler.Compile(expr)
	machine := vm.New(ins)
	return machine.Run()
}

func TestSimpleArithmetic(t *testing.T) {
	result, err := run("2 + 3 * 4")
	if err != nil {
		t.Fatalf("run error: %v", err)
	}
	if result != 14 {
		t.Fatalf("expected 14, got %d", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := run("20 / 5 - 1")
	if err != nil {
		t.Fatalf("run error: %v", err)
	}
	if result != 3 {
		t.Fatalf("expected 3, got %d", result)
	}
}
