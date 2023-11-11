package vm

/*
	- Memory
*/

func _COPY(v *Vm) {
	dst := v.getByte()
	src := v.getByte()
	v.setReg(dst, v.getReg(src))
}

/*
	- Branch
*/

func _JUMP(v *Vm) {
	offset := v.getByte()
	v.setReg(PTR, byte(offset))
}

// jmp cond, where
func _COND_JUMP(v *Vm) {
	cond := v.getByte()
	offset := v.getByte()

	if v.getReg(cond) == 1 {
		v.setReg(PTR, byte(offset))
	}
}

func _JUMP_COND_NEG(v *Vm) {
	cond := v.getByte()
	offset := v.getByte()

	if v.getReg(cond) != 1 {
		v.setReg(PTR, byte(offset))
	}
}

func _EXIT(v *Vm) {
	v.setReg(PTR, byte(len(v.byteCode)))
}

/*
	- Loaders
*/

func _LOAD_STRING(v *Vm) {
	dst := v.getByte()
	str := v.loadString()

	v.setRam(dst, str)
}

func _LOAD_NUM(v *Vm) {
	dst := v.getByte()
	val := v.getByte()

	v.setReg(dst, val)
}

func _LOAD_FLOAT(v *Vm) {
	v.setReg(PTR, 0)
}

func _LOAD_LONG_NUM(v *Vm) {
	v.setReg(PTR, 0)
}

func _LOAD_ARRAY(v *Vm) {
	dst := v.getByte()
	v.setReg(dst, 0)
}

/*
	- Conditions
*/

// eql, reg, src0, src1
func _COMP_EQUAL(v *Vm) {
	dst, l, r := v.getDstLeftRight()
	left := v.getReg(l)
	right := v.getReg(r)

	result := byte(0)
	if left == right {
		result = 1
	}

	v.setReg(dst, result)
}

func _COMP_NOT_EQUAL(v *Vm) {
	dst, l, r := v.getDstLeftRight()
	left := v.getReg(l)
	right := v.getReg(r)

	result := byte(0)
	if left != right {
		result = 1
	}

	v.setReg(dst, result)
}

func _COMP_LESS_THAN(v *Vm) {
	dst, l, r := v.getDstLeftRight()
	left := v.getReg(l)
	right := v.getReg(r)

	result := byte(0)
	if left < right {
		result = 1
	}

	v.setReg(dst, result)
}

func _COMP_GREATHER_THAN(v *Vm) {
	dst, l, r := v.getDstLeftRight()
	left := v.getReg(l)
	right := v.getReg(r)

	result := byte(0)
	if left > right {
		result = 1
	}

	v.setReg(dst, result)
}

func _COMP_LESS_THAN_EQUAL(v *Vm) {
	dst, l, r := v.getDstLeftRight()
	left := v.getReg(l)
	right := v.getReg(r)

	result := byte(0)
	if left <= right {
		result = 1
	}

	v.setReg(dst, result)
}

func _COMP_GREATHER_THAN_EQUAL(v *Vm) {
	dst, l, r := v.getDstLeftRight()
	left := v.getReg(l)
	right := v.getReg(r)

	result := byte(0)
	if left >= right {
		result = 1
	}

	v.setReg(dst, result)
}

/*
	- Math
*/

func _ADD(v *Vm) {
	dst, l, r := v.getDstLeftRight()
	left := v.getReg(l)
	right := v.getReg(r)

	v.setReg(dst, left+right)
}

func _MUL(v *Vm) {
	dst, src0, src1 := v.getDstLeftRight()
	v.setReg(dst, v.registers[src0]*v.registers[src1])
}

func _DIV(v *Vm) {
	dst, src0, src1 := v.getDstLeftRight()
	v.setReg(dst, v.registers[src0]/v.registers[src1])
}

func _SUB(v *Vm) {
	dst, src0, src1 := v.getDstLeftRight()
	v.setReg(dst, v.registers[src0]-v.registers[src1])
}
