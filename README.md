# Bytecode
```
Bytecode: [1 200 14 116 105 110 121 32 118 109 32 111 110 32 116 111 112 2 13 5 2 14 3 100 15 13 14 100 16 14 14 50 17 15 16 17 17 0 16]
```

```go
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
```

# Sample
```go
func add(a int, b int) int {
    return a + b 
}

func main() {
    for {
        x := "tiny vm on top"

        a := 3
        b := 5

        if add(a, b) == add(b, b) {
            continue
        }

        break
    }
}
```

Lot of things missing
Inspired by https://github.com/jwillbold/rusty-jsyc