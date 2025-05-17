# Parser Package Design

The parser consumes tokens from the lexer and builds the AST using a
Pratt-inspired approach. It keeps track of the current and next tokens and
parses expressions according to operator precedence.

Key concepts:

- `Parser` holds the lexer and an error list.
- `parseExpression` is recursive and respects precedence via `peekPrecedence` and
  `curPrecedence`.
- Only integer literals and infix operators are currently supported.

The parser emits `ast.Expression` nodes that the evaluator or future bytecode
compiler can traverse.

