# Runner Package

The runner package provides the execution environment for toy language programs.

## Purpose

The runner handles:
- File execution
- Source code processing
- REPL enhancement
- Error reporting
- Program lifecycle

## Architecture

### Execution Pipeline

```
Source Code -> Lexer -> Parser -> Compiler -> VM -> Result
     |           |         |          |        |        |
     +--------Error----Handling----Chain-------+--------+
```

### Components

1. **File Runner**: Executes source files
2. **Source Runner**: Executes string input
3. **REPL Runner**: Interactive execution
4. **Error Handler**: Unified error reporting

## Implementation

### File Execution

```go
func RunFile(filename string) error {
    // Read file
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("cannot read file: %w", err)
    }
    
    // Execute source
    return RunSource(string(source))
}
```

### Source Execution

```go
func RunSource(source string) error {
    // Tokenize
    lexer := lexer.New(source)
    
    // Parse
    parser := parser.New(lexer)
    program := parser.ParseProgram()
    if len(parser.Errors()) > 0 {
        return handleParseErrors(parser.Errors())
    }
    
    // Compile
    compiler := compiler.New()
    err := compiler.Compile(program)
    if err != nil {
        return handleCompileError(err)
    }
    
    // Execute
    vm := vm.New(compiler.Bytecode())
    err = vm.Run()
    if err != nil {
        return handleRuntimeError(err)
    }
    
    return nil
}
```

### REPL Enhancement

```go
type REPLRunner struct {
    compiler *compiler.Compiler
    vm       *vm.VM
    globals  map[string]interface{}
}

func (r *REPLRunner) Run(input string) (interface{}, error) {
    // Parse single statement or expression
    // Compile incrementally
    // Execute in persistent VM
    // Return result or error
}
```

## Error Handling

### Error Types

1. **File Errors**: File not found, permission denied
2. **Syntax Errors**: Invalid tokens, parse failures
3. **Compile Errors**: Undefined variables, type errors
4. **Runtime Errors**: Division by zero, stack overflow

### Error Reporting

```go
type Error struct {
    Type     ErrorType
    Message  string
    Line     int
    Column   int
    Filename string
}

func (e Error) String() string {
    return fmt.Sprintf("%s:%d:%d: %s: %s",
        e.Filename, e.Line, e.Column, e.Type, e.Message)
}
```

### Example Output

```
test.toy:5:10: SyntaxError: unexpected token ')'
    print("Hello" ))
                  ^
```

## Features

### Command Line Interface

```bash
# Run a file
toy run script.toy

# Execute expression
toy eval "2 + 3 * 4"

# Start REPL
toy repl

# Check syntax
toy check script.toy
```

### Options

```go
type Options struct {
    Debug       bool   // Enable debug output
    Optimize    bool   // Enable optimizations
    StackSize   int    // VM stack size
    MemoryLimit int64  // Memory limit in bytes
}
```

### Performance Monitoring

```go
type Stats struct {
    ParseTime    time.Duration
    CompileTime  time.Duration
    ExecuteTime  time.Duration
    Instructions int
    MaxStackSize int
}
```

## Program Types

### Script Mode
- Executes entire file
- Returns exit code
- Handles imports

### Expression Mode
- Evaluates single expression
- Returns result value
- Quick calculations

### Interactive Mode
- Persistent state
- Command history
- Tab completion

## Testing

### Test Categories

1. **File execution tests**
2. **Error handling tests**
3. **Integration tests**
4. **Performance tests**

### Test Utilities

```go
func TestRunFile(t *testing.T) {
    tests := []struct {
        name     string
        file     string
        expected interface{}
        wantErr  bool
    }{
        // Test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Run test
        })
    }
}
```

## Future Enhancements

1. **Module System**
   - Import/export
   - Package management
   - Dependency resolution

2. **Debugging**
   - Breakpoints
   - Step execution
   - Variable inspection

3. **Optimization**
   - AST optimization
   - Bytecode optimization
   - Caching

4. **Tooling**
   - Formatter
   - Linter
   - Documentation generator
