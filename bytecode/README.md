# Bytecode Package

The bytecode package defines the instruction set and format for the virtual machine.

## Purpose

This package provides:
- Opcode definitions
- Instruction encoding/decoding
- Bytecode data structures
- Instruction manipulation utilities

## Instruction Set Architecture

### Design Principles

1. **Stack-based**: Operations work on stack values
2. **Fixed-width opcodes**: 1 byte per opcode
3. **Variable operands**: 0-2 operands per instruction
4. **Big-endian encoding**: Multi-byte values use big-endian

### Opcode Categories

#### Constants and Literals
- `OpConstant`: Load constant from pool
- `OpTrue`: Push boolean true
- `OpFalse`: Push boolean false
- `OpNull`: Push null value

#### Arithmetic Operations
- `OpAdd`: Addition
- `OpSub`: Subtraction
- `OpMul`: Multiplication
- `OpDiv`: Division
- `OpMod`: Modulo

#### Comparison Operations
- `OpEqual`: Equality check
- `OpNotEqual`: Inequality check
- `OpGreaterThan`: Greater than
- `OpLessThan`: Less than

#### Control Flow
- `OpJump`: Unconditional jump
- `OpJumpNotTrue`: Conditional jump
- `OpLoop`: Loop constructs

#### Stack Operations
- `OpPop`: Remove top element
- `OpDup`: Duplicate top element
- `OpSwap`: Swap top two elements

#### Variables
- `OpSetGlobal`: Store global variable
- `OpGetGlobal`: Load global variable
- `OpSetLocal`: Store local variable
- `OpGetLocal`: Load local variable

#### Functions
- `OpCall`: Function call
- `OpReturn`: Return from function
- `OpReturnValue`: Return with value

#### Data Structures
- `OpArray`: Create array
- `OpHash`: Create hash/map
- `OpIndex`: Array/map access

## Instruction Format

```
[opcode: 1 byte][operand1: 0-2 bytes][operand2: 0-2 bytes]
```

### Examples

1. **OpConstant**
   ```
   [0x00][index_high][index_low]
   ```
   Loads constant at index from constant pool

2. **OpAdd**
   ```
   [0x01]
   ```
   No operands, works on stack

3. **OpJump**
   ```
   [0x0A][addr_high][addr_low]
   ```
   Jumps to address

## Data Structures

### Instructions
```go
type Instructions []byte
```

### Bytecode
```go
type Bytecode struct {
    Instructions Instructions
    Constants    []interface{}
}
```

### Definition
```go
type Definition struct {
    Name          string
    OperandWidths []int
}
```

## Utility Functions

- `Make(op Opcode, operands ...int) []byte`: Create instruction
- `Lookup(op byte) (*Definition, error)`: Get opcode definition
- `ReadOperands(def *Definition, ins Instructions) ([]int, int)`: Decode operands
- `String(ins Instructions) string`: Human-readable format

## Future Extensions

1. **Optimization Opcodes**
   - Combined operations (e.g., increment)
   - Specialized comparisons

2. **Advanced Features**
   - Closures
   - Exception handling
   - Coroutines

3. **Performance**
   - Superinstructions
   - Inline caching
