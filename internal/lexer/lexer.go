package lexer

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// OPCODE
const (
	SET   = 0x01
	PRINT = 0x02
	ADD   = 0x03
	MUL   = 0x04
	PUSH  = 0x05
	DIV   = 0x06
	XOR   = 0x07
	EQL   = 0x08
	LSHF  = 0x09
	RSHF  = 0x10
	MOV   = 0x11
	LOAD  = 0x12
	POP   = 0x13
)

var (
	opMap = map[string]byte{
		"SET":   0x01,
		"PRINT": 0x02,
		"ADD":   0x03,
		"MUL":   0x04,
		"PUSH":  0x05,
		"DIV":   0x06,
		"XOR":   0x07,
		"EQL":   0x08,
		"LSHF":  0x09,
		"RSHF":  0x10,
		"MOV":   0x11,
		"LOAD":  0x12,
		"POP":   0x13,
	}
)

// REGISTERS
const (
	REG1 = 0x01
	REG2 = 0x02
	REG3 = 0x03
	REG4 = 0x04
	REG5 = 0x05
)

func intToBytes(num int) []byte {
	bytes := make([]byte, 8)

	binary.LittleEndian.PutUint64(bytes, uint64(num))

	return bytes
}

func ProccessFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	out := []byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		tokens := strings.Split(strings.ReplaceAll(line, " ", ""), ",")
		log.Println(tokens)

		op_hex, ok := opMap[tokens[0]]
		if !ok {
			panic(fmt.Sprintf("invalid op: %v", tokens))
		}

		out = append(out, op_hex)
		log.Println(op_hex, out)

		if len(tokens) == 1 {
			continue
		}

		for _, arg := range tokens[1:] {
			op_hex, ok := opMap[arg]

			if !ok {
				if strings.HasPrefix(arg, "REG") {
					a, err := strconv.Atoi(strings.Split(arg, "REG")[1])
					if err != nil {
						panic(err)
					}
					
					out = append(out, intToBytes(a)[0])
					continue
				}

				a, err := strconv.Atoi(arg)
				if err != nil {
					panic(err)
				}
				
				out = append(out, intToBytes(a)[0])
				continue
			}

			out = append(out, op_hex)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
