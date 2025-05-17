package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

// Token types
const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers + literals
    IDENT  = "IDENT"  // add, foobar, x, y
    INT    = "INT"    // 1234
    STRING = "STRING" // "hello world"

    // Operators
    PLUS     = "+"
    MINUS    = "-"
    ASTERISK = "*"
    SLASH    = "/"

    // Assignment
    ASSIGN = "="

    // Comparison
    EQ     = "=="
    NOT_EQ = "!="
    LT     = "<"
    GT     = ">"

    // Delimiters
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"
    SEMICOLON = ";"
    COMMA     = ","

    // Keywords
    LET      = "LET"
    FUNCTION = "FUNCTION"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    WHILE    = "WHILE"
    FOR      = "FOR"
    RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
    "let":    LET,
    "fn":     FUNCTION,
    "true":   TRUE,
    "false":  FALSE,
    "if":     IF,
    "else":   ELSE,
    "while":  WHILE,
    "for":    FOR,
    "return": RETURN,
}

// LookupIdent checks if an identifier is a keyword
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}