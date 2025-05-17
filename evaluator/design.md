# Evaluator Package Design

The evaluator walks the AST and computes the result of expressions. It returns
integer values and reports errors such as division by zero.

Design points:

- `Eval` is a recursive function that pattern matches on node types.
- `evalInfixExpression` handles arithmetic operators; it checks for division by
  zero and prints an error message.
- Additional node types will be supported as the language evolves.

The evaluator serves as a reference implementation of expression semantics and
will later inform the bytecode generator and VM.

