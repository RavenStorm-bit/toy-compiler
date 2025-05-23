package ast

import (
    "bytes"
    "strings"
    "github.com/RavenStorm-bit/toy-compiler/token"
)

type Node interface {
    TokenLiteral() string
    String() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

type IntegerLiteral struct {
    Token token.Token
    Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type InfixExpression struct {
    Token    token.Token // The operator token, e.g. +
    Left     Expression
    Operator string
    Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
    var out bytes.Buffer
    out.WriteString("(")
    out.WriteString(ie.Left.String())
    out.WriteString(" " + ie.Operator + " ")
    out.WriteString(ie.Right.String())
    out.WriteString(")")
    return out.String()
}

// Program is the root node of every AST
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    }
    return ""
}

func (p *Program) String() string {
    var out bytes.Buffer
    for _, s := range p.Statements {
        out.WriteString(s.String())
    }
    return out.String()
}

// ExpressionStatement wraps expressions that can stand alone as statements
type ExpressionStatement struct {
    Token      token.Token
    Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }
    return ""
}

// Identifier represents variable names
type Identifier struct {
    Token token.Token
    Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// LetStatement represents variable declaration
type LetStatement struct {
    Token token.Token
    Name  *Identifier
    Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
    var out bytes.Buffer
    out.WriteString(ls.TokenLiteral() + " ")
    out.WriteString(ls.Name.String())
    out.WriteString(" = ")
    if ls.Value != nil {
        out.WriteString(ls.Value.String())
    }
    out.WriteString(";")
    return out.String()
}

// StringLiteral represents string values
type StringLiteral struct {
    Token token.Token
    Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// Boolean represents true/false values
type Boolean struct {
    Token token.Token
    Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// IfExpression represents if/else conditionals
type IfExpression struct {
    Token       token.Token
    Condition   Expression
    Consequence *BlockStatement
    Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
    var out bytes.Buffer
    out.WriteString("if")
    out.WriteString(ie.Condition.String())
    out.WriteString(" ")
    out.WriteString(ie.Consequence.String())
    if ie.Alternative != nil {
        out.WriteString("else ")
        out.WriteString(ie.Alternative.String())
    }
    return out.String()
}

// BlockStatement represents a block of statements
type BlockStatement struct {
    Token      token.Token
    Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
    var out bytes.Buffer
    for _, s := range bs.Statements {
        out.WriteString(s.String())
    }
    return out.String()
}

// FunctionLiteral represents function definitions
type FunctionLiteral struct {
    Token      token.Token
    Parameters []*Identifier
    Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
    var out bytes.Buffer
    params := []string{}
    for _, p := range fl.Parameters {
        params = append(params, p.String())
    }
    out.WriteString(fl.TokenLiteral())
    out.WriteString("(")
    out.WriteString(strings.Join(params, ", "))
    out.WriteString(") ")
    out.WriteString(fl.Body.String())
    return out.String()
}

// WhileStatement represents while loops
type WhileStatement struct {
    Token     token.Token
    Condition Expression
    Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
    var out bytes.Buffer
    out.WriteString("while")
    out.WriteString(ws.Condition.String())
    out.WriteString(" ")
    out.WriteString(ws.Body.String())
    return out.String()
}

// ReturnStatement represents return statements
type ReturnStatement struct {
    Token       token.Token
    ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
    var out bytes.Buffer
    out.WriteString(rs.TokenLiteral() + " ")
    if rs.ReturnValue != nil {
        out.WriteString(rs.ReturnValue.String())
    }
    out.WriteString(";")
    return out.String()
}

// CallExpression represents function calls
type CallExpression struct {
    Token     token.Token // The '(' token
    Function  Expression  // Identifier or FunctionLiteral
    Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
    var out bytes.Buffer

    args := []string{}
    for _, a := range ce.Arguments {
        args = append(args, a.String())
    }

    out.WriteString(ce.Function.String())
    out.WriteString("(")
    out.WriteString(strings.Join(args, ", "))
    out.WriteString(")")

    return out.String()
}

// AssignmentStatement represents variable assignments
type AssignmentStatement struct {
    Token token.Token // the = token
    Name  *Identifier
    Value Expression
}

func (as *AssignmentStatement) statementNode()       {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignmentStatement) String() string {
    var out bytes.Buffer

    out.WriteString(as.Name.String())
    out.WriteString(" = ")
    out.WriteString(as.Value.String())
    out.WriteString(";")

    return out.String()
}