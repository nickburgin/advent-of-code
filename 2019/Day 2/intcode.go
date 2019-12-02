package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	program := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		program += row
	}

	integerStrings := strings.Split(program, ",")

	initialMemory := make([]int64, 0, len(integerStrings))

	args := os.Args[1:]

	paramsProvided := false
	var noun int64 = 0
	var verb int64 = 0

	solving := false
	var target int64 = 0

	switch len(args) {
	case 0:
	case 1:
		var err error
		target, err = strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("target was not an integer %s", args[0]))
		}
		solving = true
	case 2:
		var err error
		noun, err = strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("noun was not an integer %s", args[0]))
		}
		verb, err = strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("verb was not an integer %s", args[1]))
		}
		paramsProvided = true
	default:
	}

	for _, integerString := range integerStrings {
		integer, err := strconv.ParseInt(integerString, 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse integer %s", integerString))
		}
		initialMemory = append(initialMemory, integer)
	}

	if solving {
		for noun = 0; noun <= 99; noun++ {
			for verb = 0; verb <= 99; verb++ {
				memory := make([]int64, len(initialMemory))
				copy(memory, initialMemory)
				memory[1] = noun
				memory[2] = verb
				runProgram(memory)
				if memory[0] == target {
					fmt.Printf("found noun %v and verb %v", noun, verb)
					return
				}
			}
		}
	}

	if paramsProvided {
		initialMemory[1] = noun
		initialMemory[2] = verb
	}

	err := runProgram(initialMemory)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", initialMemory)
}

func runProgram (memory []int64) error {
	for i := 0; i < len(memory); i += 4 {
		opcode := memory[i]

		switch opcode {
		case 1:
			op1 := memory[i+1]
			op2 := memory[i+2]
			op3 := memory[i+3]

			memory[op3] = memory[op1] + memory[op2]
		case 2:
			op1 := memory[i+1]
			op2 := memory[i+2]
			op3 := memory[i+3]

			memory[op3] = memory[op1] * memory[op2]
		case 99:
			return nil
		default:
			return fmt.Errorf("unknown opcode %v at position %v", opcode, i)
		}
	}
	return nil
}