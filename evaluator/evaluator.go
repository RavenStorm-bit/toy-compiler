package evaluator

import (
	"fmt"
	"github.com/RavenStorm-bit/toy-compiler/ast"
)

func Eval(node ast.Node) int64 {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return node.Value
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}
	return 0
}

func evalInfixExpression(operator string, left, right int64) int64 {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right == 0 {
			fmt.Println("Error: Division by zero")
			return 0
		}
		return left / right
	default:
		return 0
	}
}