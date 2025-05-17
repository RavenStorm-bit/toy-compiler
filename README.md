# Toy Compiler

A simple compiler implementation in Go that demonstrates the fundamental concepts of language processing including lexical analysis, parsing, and AST generation.

## Features

- **Lexer**: Tokenizes source code into meaningful symbols
- **Parser**: Converts tokens into an Abstract Syntax Tree (AST)
- **AST**: Hierarchical representation of the program structure

## Supported Language Features

- **Variables**: Declaration with `let`
- **Data Types**: Integers, strings, booleans
- **Functions**: Function declarations with parameters
- **Control Flow**: `if/else` statements, `while` loops
- **Expressions**: Arithmetic operations, comparisons
- **Assignments**: Variable reassignment
- **Return Statements**: Early returns from functions

## Example Code

```javascript
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
```

## Project Structure

```
toy-compiler/
├── ast/          # Abstract Syntax Tree definitions
├── lexer/        # Lexical analyzer
├── parser/       # Syntax parser
├── token/        # Token definitions
├── test/         # Test files
└── main.go       # Demo application
```

## Running the Project

```bash
# Run tests
go test ./...

# Run the demo
go run main.go
```

## Implementation Details

### Lexer
The lexer (`lexer/lexer.go`) reads the source code character by character and produces tokens. It handles:
- Keywords (let, if, else, while, fn, return)
- Identifiers and literals
- Operators and delimiters
- Whitespace skipping

### Parser
The parser (`parser/parser.go`) uses recursive descent parsing with operator precedence to build the AST. It supports:
- Pratt parsing for expressions
- Statement parsing (let, while, return)
- Function definitions and calls
- Infix and prefix expressions

### AST
The AST (`ast/ast.go`) defines the node types for the syntax tree:
- Program (root node)
- Statements (LetStatement, WhileStatement, ReturnStatement, etc.)
- Expressions (Identifier, IntegerLiteral, InfixExpression, etc.)

## Future Enhancements

- [ ] Type checking
- [ ] Code generation (to bytecode or machine code)
- [ ] More operators (++, --, +=, etc.)
- [ ] Arrays and objects
- [ ] Import/module system
- [ ] Error handling improvements
- [ ] Optimization passes

## License

MIT