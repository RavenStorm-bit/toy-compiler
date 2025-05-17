package lexer

import (
	"github.com/RavenStorm-bit/toy-compiler/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `5 + 10 * 2`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "5"},
		{token.PLUS, "+"},
		{token.INT, "10"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
