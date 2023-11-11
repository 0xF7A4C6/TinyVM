package vm

import "fmt"

func (vm *Vm) showRam() {
	fmt.Println("\n-- Dump of ram:\n")
	for i, reg := range vm.ram {
		fmt.Printf(" [ram+%v] value: %v\n", i, reg)
	}
}

/*
Append []byte in ram and save len+ptr into registers
*/
func (vm *Vm) setRam(reg byte, value []byte) {
	ptr := len(vm.ram)
	vm.ram = append(vm.ram, value...)

	vm.registers[reg] = byte(ptr)
	vm.registers[reg+1] = byte(len(value))
}
