# Lexer Package Design

The lexer scans source code and produces a stream of tokens for the parser. It
maintains an input string and reads characters sequentially. Each call to
`NextToken` returns the next `token.Token` found.

Highlights:

- Supports integer literals and the `+`, `-`, `*`, `/` operators.
- Skips whitespace between tokens.
- Tracks the current and next read positions to allow lookahead.

The lexer is a self-contained component that depends only on the `token`
package. It forms the first stage of the compiler pipeline.

