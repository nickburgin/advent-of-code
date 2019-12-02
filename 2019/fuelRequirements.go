package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var requiredFuelMass int64 = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		payloadMass, err := strconv.ParseInt(row, 10, 64)
		if err != nil {
			continue
		}
		requiredFuelMass += (payloadMass/3)-2
	}

	fmt.Print(requiredFuelMass)
}
