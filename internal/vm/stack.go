package vm

import "fmt"

type Stack struct {
	memory []byte
}

func NewStack() *Stack {
	return &Stack{
		memory: make([]byte, 0),
	}
}

func (S *Stack) showStack() {
	fmt.Println("\n-- Dump of stack:\n")
	for i, mem := range S.memory {
		fmt.Printf(" [mem+%v] value: %v\n", i, mem)
	}
}

func (S *Stack) pop() byte {
	if len(S.memory) == 0 {
		return 0
	}

	val := S.memory[len(S.memory)-1]
	S.memory = S.memory[:len(S.memory)-1]

	return val
}

func (S *Stack) push(value byte) {
	S.memory = append(S.memory, value)
}

func (S *Stack) len() int {
	return len(S.memory)
}

func (S *Stack) top() byte {
	if len(S.memory) == 0 {
		return 1
	}
	return S.memory[len(S.memory)-1]
}