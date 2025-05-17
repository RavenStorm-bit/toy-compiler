# Token Package Design

This package defines the lexical tokens used throughout the compiler. A `Token`
encapsulates a `Type` and its literal string. Token types include integers,
operators, and parentheses. The lexer emits tokens of these types so the parser
can build an AST.

Key elements:

- `TokenType` identifies the kind of token, e.g. `INT`, `PLUS`, `MINUS`.
- `Token` holds the type and the original string literal.
- Constants group all supported token types in one place.

This package is intentionally small and does not depend on other packages.

