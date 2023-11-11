package vm

// Registers
const (
	PTR = 100
)

// Op
const (
	// Loaders
	LOAD_STRING   = 1
	LOAD_NUM      = 2
	LOAD_FLOAT    = 3
	LOAD_LONG_NUM = 4
	LOAD_ARRAY    = 5

	// Misc
	PROPACCESS      = 10
	FUNC_CALL       = 11
	EVAL            = 12
	CALL_BCFUNC     = 13
	RETURN_BCFUNC   = 14
	COPY            = 15
	EXIT            = 16
	COND_JUMP       = 17
	JUMP            = 18
	JUMP_COND_NEG   = 19
	BCFUNC_CALLBACK = 20
	PROPSET         = 21
	TRY             = 22
	THROW           = 23

	// Comparisons
	COMP_EQUAL               = 50
	COMP_NOT_EQUAL           = 51
	COMP_LESS_THAN           = 54
	COMP_GREATHER_THAN       = 55
	COMP_LESS_THAN_EQUAL     = 56
	COMP_GREATHER_THAN_EQUAL = 57

	// Math
	ADD = 100
	MUL = 101
	SUB = 102
	DIV = 103
)

type Vm struct {
	op        map[int]func(*Vm)
	registers map[byte]byte
	byteCode  []byte
	ram       []byte
}
