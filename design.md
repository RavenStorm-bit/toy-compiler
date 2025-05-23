# Project Design

This document describes the planned evolution of the toy compiler into a small but Turing-complete virtual machine (VM) implemented in Go.

## Goals
- Provide a minimal yet practical language capable of general computation.
- Keep the implementation simple and well structured.
- Maintain a REPL for experimentation while supporting script execution.

## Architecture Overview
1. **Frontend**
   - **Lexer**: Converts input source code into tokens.
   - **Parser**: Builds an Abstract Syntax Tree (AST) from the tokens.
   - **AST**: Represents expressions, statements, and program structure.
2. **Bytecode Generation**
   - Traverse the AST and emit bytecode instructions.
   - Instructions will include arithmetic, variable access, conditional jumps, loops, and function calls.
3. **Virtual Machine**
   - A stack-based VM executes the bytecode.
   - Supports a call stack, global and local variables, and basic heap allocation.
4. **Standard Library**
   - Provide small built-in functions (e.g., print) to facilitate programming.
5. **REPL and CLI**
   - Continue to offer an interactive REPL.
   - Add the ability to run source files directly.

## Directory Structure

The project is organized into the following packages:

```
toy-compiler/
├── token/        # Token type definitions
├── lexer/        # Lexical analysis (tokenization)
├── ast/          # Abstract Syntax Tree definitions
├── parser/       # Parser implementation
├── compiler/     # Bytecode generation from AST
├── bytecode/     # Bytecode instruction definitions
├── vm/           # Virtual machine implementation
├── evaluator/    # Direct AST evaluation (legacy)
├── stdlib/       # Standard library functions
├── repl/         # Interactive REPL
├── runner/       # File execution support
├── main.go       # CLI entry point
├── go.mod        # Go module definition
├── README.md     # Project documentation
└── design.md     # This design document
```

### Package Descriptions

- **token/**: Defines token types for all language constructs (numbers, operators, keywords, etc.)
- **lexer/**: Converts source code into a stream of tokens
- **ast/**: Defines node types for the Abstract Syntax Tree
- **parser/**: Builds AST from tokens using recursive descent parsing
- **compiler/**: Traverses AST and generates bytecode instructions
- **bytecode/**: Defines bytecode instruction format and constants
- **vm/**: Stack-based virtual machine that executes bytecode
- **evaluator/**: Current tree-walking interpreter (will be phased out)
- **stdlib/**: Built-in functions like print, len, etc.
- **repl/**: Interactive Read-Eval-Print Loop
- **runner/**: Executes source files from the command line

This structure supports incremental development while maintaining clean separation of concerns.

## Planned Language Features
- Integer and string literals.
- Variables with lexical scope.
- Arithmetic and comparison operators (`+`, `-`, `*`, `/`, `==`, `!=`, `<`, `>`).
- `if`/`else` expressions.
- `while` and `for` loops.
- Function definitions and calls (allowing recursion).
- Simple arrays and maps for composite data.

These features are sufficient for Turing completeness while keeping the language approachable.

## Bytecode Format (draft)
- `CONST <value>` – push a constant onto the stack
- `LOAD <name>` / `STORE <name>` – variable access
- `ADD`, `SUB`, `MUL`, `DIV` – arithmetic
- `CMP_EQ`, `CMP_NEQ`, `CMP_LT`, `CMP_GT` – comparisons
- `JMP <addr>` – unconditional jump
- `JMP_IF_FALSE <addr>` – conditional jump based on stack top
- `CALL <addr>` / `RET` – function invocation
- `POP` – discard stack top

This instruction set will expand as needed for arrays, strings, and built-ins.

## Example
```txt
let fact = fn(n) {
    if (n <= 1) { 1 } else { n * fact(n - 1) }
}
print(fact(5))  // outputs 120
```
The above program demonstrates recursion, conditionals, and function calls running on the planned VM.

## Milestones
1. Extend the parser to support statements, variables, and functions.
2. Implement bytecode generation for the extended AST.
3. Build the VM with a stack and call frames.
4. Add standard library functions and a file-based runner.
5. Expand the REPL to execute statements and load source files.

By following this design, the project will transition from a simple evaluator into a small yet fully capable VM suitable for experimenting with language features.