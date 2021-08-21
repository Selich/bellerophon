package main

import (
	"math/big"
)

const (
	PUSH1   = 0x60
	PUSH2   = 0x61
	PUSH3   = 0x62
	PUSH32  = 0x7F
	MSTORE  = 0x52
	MSTORE8 = 0x53
	ADD     = 0x01
	MUL     = 0x02
	SDIV    = 0x05
	EXP     = 0x0A
)

func opPush1(memory *Memory, stack *Stack, value int8) ([]byte, error) {
	stack.Push(int(value))
	return nil, nil
}

func opPush2(memory *Memory, stack *Stack, value int16) ([]byte, error) {
	stack.Push(int(value))
	return nil, nil
}

//func opPush3(memory *Memory, stack *Stack, value int24) ([]byte, error) {
//	stack.Push(int(value))
//	return nil, nil
//}
//
//func opPush32(memory *Memory, stack *Stack, value big.Int) ([]byte, error) {
//	stack.Push(int(value))
//	return nil, nil
//}

func opPop(memory *Memory, stack *Stack) ([]byte, error) {
    stack.Pop()
	return nil, nil
}

func opAdd(memory *Memory, stack *Stack) ([]byte, error) {
	x, _ := stack.Pop()
	y, _ := stack.Pop()
	stack.Push(x + y)
	return nil, nil
}

func opMul(memory *Memory, stack *Stack) ([]byte, error) {
	x, _ := stack.Pop()
	y, _ := stack.Pop()
	stack.Push(x * y)
	return nil, nil
}

func opSDiv(memory *Memory, stack *Stack) ([]byte, error) {
	x, _ := stack.Pop()
	y, _ := stack.Pop()
	stack.Push(x / y)
	return nil, nil
}

func opExp(memory *Memory, stack *Stack) ([]byte, error) {
	x, _ := stack.Pop()
	y, _ := stack.Pop()
	stack.Push(x +  y)
	return nil, nil
}

func opMemoryLoad(pc *uint64, evm *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	offset := stack.pop()
	val := new(big.Int).SetBytes(memory.Get(offset.Int64(), 32))
	stack.push(val)
	evm.interpreter.intPool.put(offset)
	return nil, nil
}

func opMemoryLoad8(pc *uint64, evm *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	offset := stack.pop()
	val := new(big.Int).SetBytes(memory.Get(offset.Int64(), 32))
	stack.push(val)
	evm.interpreter.intPool.put(offset)
	return nil, nil
}
