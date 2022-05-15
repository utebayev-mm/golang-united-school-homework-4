package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

// use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	spaceCounter := 0
	runeInput := []rune(input)
	for i := 0; i < len(runeInput); i++ {
		if runeInput[i] != '-' && runeInput[i] != '+' && runeInput[i] != ' ' && (runeInput[i] > 57 || runeInput[i] < 48) {
			return "", fmt.Errorf("%s", errorNotTwoOperands)
		}
		if runeInput[i] == ' ' {
			spaceCounter++
		}

	}
	if spaceCounter == len(runeInput) || input == "" {
		return "", fmt.Errorf("%s", errorEmptyInput)
	}
	sign1 := 1
	sign2 := 1
	var numbers []int
	var indexes []int

	for i := 0; i < len(runeInput); i++ {
		if runeInput[i] <= '9' && runeInput[i] >= '0' {
			number, _ := strconv.Atoi(string(runeInput[i]))
			numbers = append(numbers, number)
			indexes = append(indexes, i)
		}
	}
	for i := 0; i < indexes[0]; i++ {
		if runeInput[i] == '-' {
			sign1 = -1
		} else {
			sign1 = 1
		}
	}

	for i := indexes[0]; i < indexes[1]; i++ {
		if runeInput[i] == '-' {
			sign2 = -1
			break
		}
	}
	number1 := numbers[0] * sign1
	number2 := numbers[1] * sign2
	return strconv.Itoa(number1 + number2), nil
}
