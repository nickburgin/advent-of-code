package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	input := append([]string{}, "abcdef")
	input = append(input, "bababc")
	input = append(input, "abbcde")
	input = append(input, "abcccd")
	input = append(input, "aabcdd")
	input = append(input, "abcdee")
	input = append(input, "ababab")

	var expectedResult int64 = 12

	assert.Equal(t, expectedResult, inventoryChecksum(input))
}

func TestFindBoxes(t *testing.T) {
	input := append([]string{}, "abcde")
	input = append(input, "fghij")
	input = append(input, "klmno")
	input = append(input, "pqrst")
	input = append(input, "fguij")
	input = append(input, "axcye")
	input = append(input, "wvxyz")

	expectedResult := "fgij"

	assert.Equal(t, expectedResult, findBoxes(input))
}

func TestFindCommonLetters(t *testing.T) {
	expectedResult := "fgij"
	assert.Equal(t, expectedResult, findCommonLetters("fghij", "fguij"))
}