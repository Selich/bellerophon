package main


var cost = map[string]int{
	"PUSH1": 3,
	"PUSH2": 3,
	"PUSH3": 3,
	"PUSH32": 3,
	"MSTORE": 3,
	"MSTORE8": 3,
	"ADD": 3,
	"MUL": 5,
	"SDIV": 5,
	"EXP": 10,
}


func calculateGas(operation string, value int) int{
	var gas int
	gas = cost[operation]
	return gas
}


