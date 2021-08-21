package main

import (
	"fmt"
	"strconv"
	"strings"
)


func getOpcode(memory Memory, stack Stack, code string) int {
	words := strings.Fields(code)
	value, _ := strconv.ParseInt(words[1], 16, 64)

	switch words[0] {
	case "PUSH1":
		opPush1(&memory, &stack, uint8(value))
	case "PUSH2":
		return 0x61
	case "PUSH3":
		return 0x62
	case "PUSH32":
		return 0x7F
	case "MSTORE":
		value, _ := stack.Pop()
		offset,_  := stack.Pop()
		memory.store[offset] = byte(value)
		return 0x52
	case "MSTORE8":
		return 0x53
	case "ADD":
		return 0x01
	case "MUL":
		return 0x02
	case "SDIV":
		return 0x05
	case "EXP":
		return 0x0A
	default:
		return 0x00
	}

}

func genereateBytecode(memory Memory, stack Stack, code string) {
	opcode := getOpcode(memory, stack, code)
	h := fmt.Sprintf("%x", opcode)
	fmt.Println(h)
}

