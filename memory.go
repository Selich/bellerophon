package main

import "fmt"

type Memory struct {
	store []byte
	gas uint64
}
func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) PrintMemory(arr []string) {
	fmt.Printf("Device: %+v", arr)
}
