package vm

import (
	"fmt"
	"log"
)

func NewVm(byteCode []byte) *Vm {
	fmt.Println("Bytecode:", byteCode)

	v := Vm{
		ram:      []byte{},
		byteCode: byteCode,
		registers: map[byte]byte{
			PTR: 0,
		},
	}

	v.init()
	return &v
}

func (vm *Vm) init() {
	vm.op = map[int]func(*Vm){
		EXIT:                     _EXIT,
		LOAD_STRING:              _LOAD_STRING,
		LOAD_NUM:                 _LOAD_NUM,
		LOAD_FLOAT:               _LOAD_FLOAT,
		LOAD_LONG_NUM:            _LOAD_LONG_NUM,
		LOAD_ARRAY:               _LOAD_ARRAY,
		COMP_EQUAL:               _COMP_EQUAL,
		COMP_NOT_EQUAL:           _COMP_NOT_EQUAL,
		COMP_LESS_THAN:           _COMP_LESS_THAN,
		COMP_GREATHER_THAN:       _COMP_GREATHER_THAN,
		COMP_LESS_THAN_EQUAL:     _COMP_LESS_THAN_EQUAL,
		COMP_GREATHER_THAN_EQUAL: _COMP_GREATHER_THAN_EQUAL,
		ADD:                      _ADD,
		MUL:                      _MUL,
		DIV:                      _DIV,
		SUB:                      _SUB,
		JUMP:                     _JUMP,
		COND_JUMP:                _COND_JUMP,
		JUMP_COND_NEG:            _JUMP_COND_NEG,
		COPY:                     _COPY,
	}
}

func (vm *Vm) getDstLeftRight() (byte, byte, byte) {
	return vm.getByte(), vm.getByte(), vm.getByte()
}

func (vm *Vm) getByte() byte {
	defer func() {
		vm.registers[PTR]++
	}()

	return vm.byteCode[vm.registers[PTR]]
}

func (vm *Vm) Run() error {
	for int(vm.registers[PTR]) < len(vm.byteCode) {
		op := vm.getByte()

		if _, ok := vm.op[int(op)]; !ok {
			return fmt.Errorf("invalid op: %v", op)
		}

		vm.op[int(op)](vm)
		log.Printf("[ptr+%v] %v - %v", vm.registers[PTR], op, vm.registers)
	}

	vm.showReg()
	vm.showRam()
	return nil
}
