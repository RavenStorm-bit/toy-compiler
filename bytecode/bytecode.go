package bytecode

import (
	"fmt"
)

// Instructions is a sequence of bytecode instructions
type Instructions []byte

// Bytecode represents compiled bytecode
type Bytecode struct {
	Instructions Instructions
	Constants    []interface{}
}

// Opcode represents a single bytecode instruction
type Opcode byte

const (
	// OpConstant loads a constant onto the stack
	OpConstant Opcode = iota
	// OpAdd pops two values and pushes their sum
	OpAdd
	// OpSub pops two values and pushes their difference
	OpSub
	// OpMul pops two values and pushes their product
	OpMul
	// OpDiv pops two values and pushes their quotient
	OpDiv
	// OpPop removes the top stack element
	OpPop
	// OpTrue pushes true onto the stack
	OpTrue
	// OpFalse pushes false onto the stack
	OpFalse
	// OpEqual compares two values for equality
	OpEqual
	// OpNotEqual compares two values for inequality
	OpNotEqual
	// OpGreaterThan compares if left > right
	OpGreaterThan
	// OpJumpNotTrue jumps if top of stack is not true
	OpJumpNotTrue
	// OpJump unconditional jump
	OpJump
)

// Definition describes an opcode's structure
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant:    {"OpConstant", []int{2}},
	OpAdd:         {"OpAdd", []int{}},
	OpSub:         {"OpSub", []int{}},
	OpMul:         {"OpMul", []int{}},
	OpDiv:         {"OpDiv", []int{}},
	OpPop:         {"OpPop", []int{}},
	OpTrue:        {"OpTrue", []int{}},
	OpFalse:       {"OpFalse", []int{}},
	OpEqual:       {"OpEqual", []int{}},
	OpNotEqual:    {"OpNotEqual", []int{}},
	OpGreaterThan: {"OpGreaterThan", []int{}},
	OpJumpNotTrue: {"OpJumpNotTrue", []int{2}},
	OpJump:        {"OpJump", []int{2}},
}

// Lookup returns the definition for an opcode
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

// Make creates a bytecode instruction
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, operand := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			instruction[offset] = byte(operand >> 8)
			instruction[offset+1] = byte(operand)
		}
		offset += width
	}

	return instruction
}
