package vm

import (
	"fmt"
	"github.com/gookit/color"
	"log"
)

// 2 5 8 10

func printRegister(val string) string {
	return fmt.Sprintf("<fg=b2c28c>%s</>", val)
}

func printIntegrer(val int) string {
	return fmt.Sprintf("<fg=56b0f0>%d</>", val)
}

func printDebug(val string) string {
	return fmt.Sprintf("<fg=363232>%s</>", val)
}
func (vm *Vm) Load(ptr int) int {
	val := int(vm.ByteCode[ptr+2])
	register := fmt.Sprintf("r%d", vm.ByteCode[ptr+1])

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>LOAD</> %v, %v\n", ptr, printRegister(register), printIntegrer(val))
	vm.SetRegister(register, val)
	vm.Stack.push(val)

	return 3
}

func (vm *Vm) Push(ptr int) int {
	val := int(vm.ByteCode[ptr+1])
	vm.Stack.push(val)

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>PUSH</> %v\n", ptr, printIntegrer(val))
	return 2
}

func (vm *Vm) Print(ptr int) int {
	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>PRINT</>\n", ptr)

	value := vm.Stack.top()
	if value != -1 {
		log.Printf("Value on the stack: %d\n", value)
	} else {
		log.Printf("Stack is empty\n")
	}

	return 1
}

func (vm *Vm) Add(ptr int) int {
	a := vm.Stack.pop()
	b := vm.Stack.pop()
	vm.Stack.push(a + b)

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>ADD </> %s\n", ptr, printDebug(fmt.Sprintf("[%d+%d]", a, b)))
	return 1
}

func (vm *Vm) Mul(ptr int) int {
	a := vm.Stack.pop()
	b := vm.Stack.pop()
	vm.Stack.push(a * b)

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>MUL </> %s\n", ptr, printDebug(fmt.Sprintf("[%d*%d]", a, b)))
	return 1
}

func (vm *Vm) Xor(ptr int) int {
	a := vm.Stack.pop()
	b := vm.Stack.pop()
	vm.Stack.push(a ^ b)

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>XOR </> %s\n", ptr, printDebug(fmt.Sprintf("[%d*%d]", a, b)))
	return 1
}

func (vm *Vm) Div(ptr int) int {
	a := vm.Stack.pop()
	b := vm.Stack.pop()
	vm.Stack.push(a / b)

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>DIV </> %s\n", ptr, printDebug(fmt.Sprintf("[%d/%d]", a, b)))
	return 1
}

func (vm *Vm) Equal(ptr int) int {
	a := vm.Stack.pop()
	b := vm.Stack.pop()

	if a == b {
		vm.Stack.push(1)
	} else {
		vm.Stack.push(0)
	}

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>EQL </> %s\n", ptr, printDebug(fmt.Sprintf("[%d==%d]", a, b)))
	return 1
}

func (vm *Vm) LeftShift(ptr int) int {
    a := vm.Stack.pop()
    b := vm.Stack.pop()
    vm.Stack.push(b << a)

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>LSHF</> %s\n", ptr, printDebug(fmt.Sprintf("[%d<<%d]", a, b)))
    return 1
}

func (vm *Vm) RightShift(ptr int) int {
    a := vm.Stack.pop()
    b := vm.Stack.pop()
    vm.Stack.push(b >> a)

	color.Printf("<fg=ffffff>|> </><fg=f08956>[</>%d<fg=f08956>]</> <fg=b67bed>RSHF</> %s\n", ptr, printDebug(fmt.Sprintf("[%d>>%d]", a, b)))
    return 1
}