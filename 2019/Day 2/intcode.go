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

	integers := make([]int64, 0, len(integerStrings))

	args := os.Args[1:]

	paramsProvided := false
	var noun int64 = 0
	var verb int64 = 0

	switch len(args) {
	case 0:
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
		integers = append(integers, integer)
	}

	if paramsProvided {
		integers[1] = noun
		integers[2] = verb
	}

	for i := 0; i < len(integers); i += 4 {
		opcode := integers[i]
		end := false

		switch opcode {
		case 1:
			op1 := integers[i+1]
			op2 := integers[i+2]
			op3 := integers[i+3]

			integers[op3] = integers[op1] + integers[op2]
		case 2:
			op1 := integers[i+1]
			op2 := integers[i+2]
			op3 := integers[i+3]

			integers[op3] = integers[op1] * integers[op2]
		case 99:
			end = true
		default:
			fmt.Println(fmt.Errorf("unknown opcode %v at position %v", opcode, i))
			end = true
		}

		if end {
			break
		}
	}

	fmt.Printf("%v", integers)
}
