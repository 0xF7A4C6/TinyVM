package main

import (
	_ "tinyvm/internal/parser"
	"tinyvm/internal/vm"
)

func main() {
	/*
		vm := vm.NewVm([]byte{
			// Load string into RAM
			vm.LOAD_STRING, 200, 14, 116, 105, 110, 121, 32, 118, 109, 32, 111, 110, 32, 116, 111, 112, // tiny vm on top

			vm.LOAD_NUM, 13, 5, // mov 5 into reg+13
			vm.LOAD_NUM, 14, 3, // mov 3 into reg+14

			vm.ADD, 15, 13, 14, // 5 + 3
			vm.ADD, 16, 14, 14, // 5 + 5

			vm.COMP_EQUAL, 17, 15, 16, //  if (5+3) == (5+5) {
			vm.COND_JUMP, 17, 0,       //  restart }
			vm.EXIT,				   //  else { exit }
		})
	*/

	vm := vm.NewVm([]byte{
		vm.LOAD_STRING, 200, 18, 104, 116, 116, 112, 115, 58, 47, 47, 103, 111, 111, 103, 108, 101, 46, 99, 111, 109,
		vm.PUSH, 19, // ptr of resp body
		vm.PUSH, 18,
		vm.PUSH, 0,
		vm.CALL_BIND, 12, vm.BIND_HTTP_REQ,
		vm.EXIT,
	})

	if err := vm.Run(); err != nil {
		panic(err)
	}
}
