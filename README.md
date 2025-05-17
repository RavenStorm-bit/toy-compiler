# Toy Compiler

A simple arithmetic expression compiler written in Go. This compiler parses expressions and executes them on a small stack-based virtual machine. Supported operations include:
- Addition (+)
- Subtraction (-)
- Multiplication (*)
- Division (/)

## Features
- Lexical analysis (tokenization)
- Parsing with operator precedence
- Bytecode generation
- Stack-based VM
- Interactive REPL

## Structure
- `token/`: Token definitions
- `lexer/`: Tokenizer that converts input into tokens
- `ast/`: Abstract Syntax Tree definitions
 - `parser/`: Parser that builds the AST from tokens
 - `compiler/`: Converts the AST into VM instructions
 - `vm/`: Executes bytecode with a stack machine
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
3. **Compiler & VM**: The AST is compiled into bytecode and executed by the stack-based VM
4. **REPL**: Provides an interactive interface to test expressions

