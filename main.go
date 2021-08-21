package main

import (
	"bufio"
	"log"
	"os"
)

//fmt.Println(hex.EncodeToString(buf))
//expected := "EEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"
//  var bytecode := "60016020526002606452600361ff0052600362ffffff526005601053";
// keccak256 ??
// keccak256 = "ab2744998886b708acadc0a32428d0aa1953e83924383d21c6de5dac852ccbcc"

func main() {
	var stack Stack
	var memory Memory
	file, err := os.Open("code.asm")

	if err != nil { log.Fatal(err) }
	defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		genereateBytecode(memory, stack, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
