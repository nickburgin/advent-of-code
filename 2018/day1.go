package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
)

func main() {
	args := os.Args[1:]

	filename := args[0]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	frequencyAdjustmentsStrings := strings.Split(string(data), "\n")

	foundFrequencies := map[int64]bool{}

	var calibrationResult int64
	calibrationResult = 0
	foundFrequencies[calibrationResult] = true

	foundFrequencyTwice := false

	for !foundFrequencyTwice {
		for _, frequencyAdjustmentString := range frequencyAdjustmentsStrings {
			if frequencyAdjustmentString == "" {
				continue
			}
			frequencyAdjustment, err := strconv.ParseInt(frequencyAdjustmentString, 10, 64)
			if err != nil {
				panic(err)
			}
			calibrationResult = calibrationResult + frequencyAdjustment
			if foundFrequencies[calibrationResult] {
				foundFrequencyTwice = true
				break
			}
			foundFrequencies[calibrationResult] = true
		}
	}

	fmt.Printf("Frequency Calibration: %v\n", calibrationResult)
}
