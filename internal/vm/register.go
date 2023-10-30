package vm

type Register struct {
	Name  string
	Value int
}

func NewRegister(Name string, Value int) *Register {
	return &Register{
		Name:  Name,
		Value: Value,
	}
}

func (R *Register) GetName() string {
	return R.Name
}

func (R *Register) GetValue() int {
	return R.Value
}

func (R *Register) SetName(name string) {
	R.Name = name
}

func (R *Register) SetValue(Value int) {
	R.Value = Value
}
