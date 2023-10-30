package vm

import (
	"log"
)

type Vm struct {
	Stack     *Stack
	Registers map[string]*Register
	opMap     map[byte]string

	ByteCode []byte
}

type OperationNotFoundError struct {
	Operation string
}

type InvalidOperationError struct {
	Reason string
}

func (e *OperationNotFoundError) Error() string {
	return "Operation not found: " + e.Operation
}

func (e *InvalidOperationError) Error() string {
	return "Invalid operation: " + e.Reason
}

func NewVm() *Vm {
	V := Vm{
		Registers: make(map[string]*Register),
		Stack:     NewStack(),
		opMap: map[byte]string{
			0x01: "set",
			0x02: "print",
			0x03: "add",
			0x04: "mul",
			0x05: "push",
			0x06: "div",
			0x07: "xor",
			0x08: "equal",
			0x09: "left_shift",
			0x10: "right_shift",
			0x11: "mov",
			0x12: "load",
			0x13: "pop",
		},
	}

	return &V
}

func (vm *Vm) GetRegister(name string) *Register {
	if reg, ok := vm.Registers[name]; ok {
		return reg
	}

	return nil
}

func (vm *Vm) SetRegister(name string, value int) {
	if reg, ok := vm.Registers[name]; ok {
		reg.SetValue(value)
	} else {
		reg := NewRegister(name, value)
		vm.Registers[name] = reg
	}
}

func (vm *Vm) Run(bytecode []byte) {
	vm.ByteCode = bytecode

	opMap := map[byte]func(int) int{
		0x01: vm.Set,
		0x02: vm.Print,
		0x03: vm.Add,
		0x04: vm.Mul,
		0x05: vm.Push,
		0x06: vm.Div,
		0x07: vm.Xor,
		0x08: vm.Equal,
		0x09: vm.LeftShift,
		0x10: vm.RightShift,
		0x11: vm.Mov,
		0x12: vm.Load,
		0x13: vm.Pop,
	}

	for i := 0; i < len(vm.ByteCode); {
		op := vm.ByteCode[i]

		fn, exists := opMap[op]
		if exists {
			i += fn(i)
		} else {
			log.Printf("Unrecognized opcode: %X\n", op)
			break
		}

		//log.Println(vm.Registers)
		//log.Println(vm.Stack)
	}
}
