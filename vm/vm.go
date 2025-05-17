package vm

import "fmt"

// OpCode represents the type of instruction.
type OpCode byte

const (
	OpPush OpCode = iota
	OpAdd
	OpSub
	OpMul
	OpDiv
)

// Instruction is a single VM instruction.
type Instruction struct {
	Op  OpCode
	Arg int64 // used by OpPush
}

// VM is a simple stack-based virtual machine.
type VM struct {
	instructions []Instruction
	stack        []int64
	ip           int
}

// New creates a new VM with instructions.
func New(instructions []Instruction) *VM {
	return &VM{instructions: instructions}
}

// Run executes the instructions and returns the final stack top.
func (vm *VM) Run() (int64, error) {
	for vm.ip < len(vm.instructions) {
		ins := vm.instructions[vm.ip]
		switch ins.Op {
		case OpPush:
			vm.stack = append(vm.stack, ins.Arg)
		case OpAdd, OpSub, OpMul, OpDiv:
			if len(vm.stack) < 2 {
				return 0, fmt.Errorf("stack underflow")
			}
			b := vm.pop()
			a := vm.pop()
			var res int64
			switch ins.Op {
			case OpAdd:
				res = a + b
			case OpSub:
				res = a - b
			case OpMul:
				res = a * b
			case OpDiv:
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				res = a / b
			}
			vm.stack = append(vm.stack, res)
		default:
			return 0, fmt.Errorf("unknown opcode %d", ins.Op)
		}
		vm.ip++
	}
	if len(vm.stack) != 1 {
		return 0, fmt.Errorf("invalid program state")
	}
	return vm.stack[0], nil
}

func (vm *VM) pop() int64 {
	val := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return val
}
