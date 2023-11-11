package vm

import (
	"encoding/binary"
	"fmt"
)

func (vm *Vm) showReg() {
	fmt.Println("\n-- Dump of registers:\n")
	for i, reg := range vm.registers {
		fmt.Printf(" [reg+%v] value: %v\n", i, reg)
	}
}

func (vm *Vm) setReg(reg byte, value byte) {
	vm.registers[reg] = value
}

func (vm *Vm) getReg(reg byte) byte {
	// Todo: check if reg exist.
	return vm.registers[reg]
}

func (vm *Vm) loadArrayFromRegister() (out []byte) {
	arrLength := vm.getByte()

	for i := 0; i < int(arrLength); i++ {
		out = append(out, vm.getReg(vm.getByte()))
	}

	return out
}

func (vm *Vm) loadString() (out []byte) {
	stringLength := vm.getByte() // 256 len || (vm.getByte() << 8)

	for i := 0; i < int(stringLength); i++ {
		out = append(out, vm.getByte())
	}

	return out
}

func (vm *Vm) loadLongNum() int32 {
	numBytes := make([]byte, 4)

	for i := 0; i < 4; i++ {
		numBytes[i] = vm.getByte()
	}

	num := int32(binary.BigEndian.Uint32(numBytes))
	return num
}