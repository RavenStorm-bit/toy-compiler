package parser

import (
	"fmt"
	"strconv"
	"github.com/RavenStorm-bit/toy-compiler/ast"
	"github.com/RavenStorm-bit/toy-compiler/lexer"
	"github.com/RavenStorm-bit/toy-compiler/token"
)

const (
	_ int = iota
	LOWEST
	SUM        // +, -
	PRODUCT    // *, /
	PREFIX     // -X
	CALL       // myFunction(X)
)

var precedences = map[token.TokenType]int{
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseExpression() ast.Expression {
	return p.parseExpression(LOWEST)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	if p.curToken.Type != token.INT && p.curToken.Type != token.LPAREN {
		msg := fmt.Sprintf("unexpected token: %s", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}

	left := p.parseIntegerLiteral()

	for p.peekToken.Type != token.EOF && precedence < p.peekPrecedence() {
		if !p.isInfixOperator(p.peekToken.Type) {
			return left
		}

		p.nextToken()
		left = p.parseInfixExpression(left)
	}

	return left
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) isInfixOperator(t token.TokenType) bool {
	return t == token.PLUS || t == token.MINUS || t == token.ASTERISK || t == token.SLASH
}