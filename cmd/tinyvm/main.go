package main

import (
	"log"
	"tinyvm/internal/vm"
)

// OP
const (
	// push element to a register
	LOAD = 0x01

	// pop 2 first element in stack and add them, result is pushed to the stack
	ADD = 0x03

	// print first element in stack
	PRINT = 0x02

	// pop 2 first element in stack and mult them, result is pushed to the stack
	MUL = 0x04

	// push element to the stack
	PUSH = 0x05

	DIV  = 0x06
	XOR  = 0x07
	EQL  = 0x08
	LSHF = 0x09
	RSHF = 0x10
)

// REGISTERS
const (
	REG1 = 0x01
	REG2 = 0x02
	REG3 = 0x03
	REG4 = 0x04
	REG5 = 0x05
)

// mov re0 (from stack)
func main() {
	byteCode := []byte{
		PUSH, 20,
		PUSH, 10,
		ADD,
		PRINT,

		PUSH, 20,
		MUL,
		PRINT,

		PUSH, 20,
		ADD,
		PRINT,

		PUSH, 20,
		XOR,
		PRINT,

		PUSH, 20,
		LSHF,
		PRINT,
		
		PUSH, 10,
		EQL,
		PRINT,
	}

	log.Println("Bytecode:", byteCode)

	v := vm.NewVm()
	v.Run(byteCode)

	//log.Println(v.Registers)
	//log.Println(v.Stack)
}
