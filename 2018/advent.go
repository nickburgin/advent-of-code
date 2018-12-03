package main

import (
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	args := os.Args[1:]

	function := args[0]

	filename := args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	switch function {
	case "calibrate":
		fmt.Printf("Frequency Calibration: %v\n", calibrate(lines))
	case "checksum":
		fmt.Printf("Checksum: %v\n", inventoryChecksum(lines))
	case "findboxes":
		fmt.Printf("Common Letters: %v\n", findBoxes(lines))
	}
}
func findBoxes(boxIDs []string) string {
	for index, boxID := range boxIDs {
		if boxID == "" {
			continue
		}
		for i := index; i < len(boxIDs); i++ {
			otherBoxID := boxIDs[i]
			if otherBoxID == "" {
				continue
			}

			commonLetters := findCommonLetters(boxID, otherBoxID)

			if len(commonLetters) == len(boxID) - 1 {
				return commonLetters
			}
		}
	}
	return ""
}

func findCommonLetters(s string, s2 string) string {
	var resultArray []rune
	otherRunes := []rune(s2)
	for i, character := range []rune(s) {
		if character == otherRunes[i] {
			resultArray = append(resultArray, character)
		}
	}

	return string(resultArray)
}

func inventoryChecksum(boxIDs []string) int64 {
	counts := map[int64]int64{}

	for _, boxID := range boxIDs {
		characterCounts := map[rune]int64{}
		twoFound := false
		threeFound := false

		for _, character := range []rune(boxID) {
			characterCounts[character] = characterCounts[character] + 1
		}

		for _, count := range characterCounts {
			if count == 2 && !twoFound {
				twoFound = true
				counts[2] = counts[2] + 1
			}
			if count == 3 && !threeFound {
				threeFound = true
				counts[3] = counts[3] + 1
			}
		}
	}

	return counts[2] * counts[3]
}

func calibrate(frequencyAdjustmentsStrings []string) int64 {
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

	return calibrationResult
}