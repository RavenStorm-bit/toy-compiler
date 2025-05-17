package compiler

import (
	"github.com/RavenStorm-bit/toy-compiler/ast"
	"github.com/RavenStorm-bit/toy-compiler/vm"
)

// Compile converts an AST expression into VM instructions.
func Compile(node ast.Expression) []vm.Instruction {
	switch n := node.(type) {
	case *ast.IntegerLiteral:
		return []vm.Instruction{{Op: vm.OpPush, Arg: n.Value}}
	case *ast.InfixExpression:
		left := Compile(n.Left)
		right := Compile(n.Right)
		op := vm.OpAdd
		switch n.Operator {
		case "+":
			op = vm.OpAdd
		case "-":
			op = vm.OpSub
		case "*":
			op = vm.OpMul
		case "/":
			op = vm.OpDiv
		}
		ins := append(left, right...)
		ins = append(ins, vm.Instruction{Op: op})
		return ins
	default:
		return nil
	}
}
