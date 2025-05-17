package vm

import (
	"fmt"
	"github.com/RavenStorm-bit/toy-compiler/bytecode"
)

const StackSize = 2048
const GlobalsSize = 65536

// VM is the virtual machine that executes bytecode
type VM struct {
	constants    []interface{}
	instructions bytecode.Instructions
	stack        []interface{}
	sp           int // stack pointer, points to next free slot
	globals      []interface{}
}

// New creates a new VM instance
func New(bytecode *bytecode.Bytecode) *VM {
	return &VM{
		constants:    bytecode.Constants,
		instructions: bytecode.Instructions,
		stack:        make([]interface{}, StackSize),
		sp:           0,
		globals:      make([]interface{}, GlobalsSize),
	}
}

// LastPoppedStackElem returns the last popped element
func (vm *VM) LastPoppedStackElem() interface{} {
	return vm.stack[vm.sp]
}

// Run executes the bytecode
func (vm *VM) Run() error {
	ip := 0 // instruction pointer

	for ip < len(vm.instructions) {
		ip++ // advance for next iteration

		op := bytecode.Opcode(vm.instructions[ip-1])

		switch op {
		case bytecode.OpConstant:
			constIndex := int(vm.instructions[ip])<<8 | int(vm.instructions[ip+1])
			ip += 2
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}

		case bytecode.OpAdd, bytecode.OpSub, bytecode.OpMul, bytecode.OpDiv:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}

		case bytecode.OpPop:
			vm.pop()

		default:
			return fmt.Errorf("unknown opcode: %d", op)
		}
	}

	return nil
}

func (vm *VM) push(o interface{}) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

func (vm *VM) pop() interface{} {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

func (vm *VM) executeBinaryOperation(op bytecode.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	leftValue, ok := left.(int64)
	if !ok {
		return fmt.Errorf("expected integer, got %T", left)
	}

	rightValue, ok := right.(int64)
	if !ok {
		return fmt.Errorf("expected integer, got %T", right)
	}

	var result int64

	switch op {
	case bytecode.OpAdd:
		result = leftValue + rightValue
	case bytecode.OpSub:
		result = leftValue - rightValue
	case bytecode.OpMul:
		result = leftValue * rightValue
	case bytecode.OpDiv:
		if rightValue == 0 {
			return fmt.Errorf("division by zero")
		}
		result = leftValue / rightValue
	default:
		return fmt.Errorf("unknown binary operator: %d", op)
	}

	return vm.push(result)
}
