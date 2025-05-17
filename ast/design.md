# AST Package Design

The Abstract Syntax Tree (AST) package defines node types representing the
structure of parsed expressions. Each node implements the `Node` interface and
provides a `TokenLiteral()` method and a `String()` representation.

Current node types:

- `IntegerLiteral` – wraps an integer token and its value.
- `InfixExpression` – holds a left and right expression with an operator.

The AST isolates parsing details from later stages like evaluation or bytecode
generation. Additional node types can be added as the language grows.

