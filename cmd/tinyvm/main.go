package main

import (
	"log"
	"tinyvm/internal/lexer"
	"tinyvm/internal/vm"
)

// mov re0 (from stack)
func main() {
	byteCode, err := lexer.ProccessFile("./bytecode.asm")
	if err != nil {
		panic(err)
	}

	log.Println("Bytecode:", byteCode)

	v := vm.NewVm()
	v.Run(byteCode)

	log.Println("Registers:", v.Registers)
	log.Println("Stack:", v.Stack)
}
