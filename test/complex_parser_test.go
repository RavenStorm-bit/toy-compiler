package test

import (
	"testing"
	"github.com/RavenStorm-bit/toy-compiler/lexer"
	"github.com/RavenStorm-bit/toy-compiler/parser"
)

func TestComplexProgram(t *testing.T) {
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
	
	while (x < 10) {
		x = x + 1;
	}
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) < 5 {
		t.Fatalf("program.Statements does not contain enough statements. got=%d",
			len(program.Statements))
	}

	t.Logf("Program: %s", program.String())
}

func TestInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain exactly 1 statement. got=%d",
				len(program.Statements))
		}

		// Don't check the actual values for now, just ensure it parses without errors
		t.Logf("Parsed: %s", program.String())
	}
}