package vm

type Stack struct {
	memory []int
}

func NewStack() *Stack {
	return &Stack{
		memory: make([]int, 0),
	}
}

func (S *Stack) pop() int {
	if len(S.memory) == 0 {
		return -1
	}

	val := S.memory[len(S.memory)-1]
	S.memory = S.memory[:len(S.memory)-1]

	return val
}

func (S *Stack) push(value int) {
	S.memory = append(S.memory, value)
}

func (S *Stack) len() int {
	return len(S.memory)
}

func (S *Stack) top() int {
	if len(S.memory) == 0 {
		return -1
	}
	return S.memory[len(S.memory)-1]
}
