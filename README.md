# Toy Compiler

A simple arithmetic expression compiler written in Go. This compiler can parse and evaluate basic mathematical expressions with the following operations:
- Addition (+)
- Subtraction (-)
- Multiplication (*)
- Division (/)

## Features
- Lexical analysis (tokenization)
- Parsing with operator precedence
- Expression evaluation
- Interactive REPL

## Structure
- `token/`: Token definitions
- `lexer/`: Tokenizer that converts input into tokens
- `ast/`: Abstract Syntax Tree definitions
- `parser/`: Parser that builds the AST from tokens
- `evaluator/`: Evaluates the AST to produce results
- `repl/`: Interactive Read-Eval-Print Loop

## Usage

First, install Go if you haven't already.

Then run:
```bash
go run main.go
```

Example expressions:
```
>> 2 + 3
5
>> 10 - 4 * 2
2
>> (10 - 4) * 2
12
```

Type `exit` or `quit` to exit the REPL.

## How it works

1. **Lexer**: Takes the input string and converts it into tokens
2. **Parser**: Takes tokens and builds an AST using operator precedence
3. **Evaluator**: Traverses the AST and computes the result
4. **REPL**: Provides an interactive interface to test expressions