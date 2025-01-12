package string_sum

import (
	"errors"
	"fmt"
	"regexp"
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
			_, err := strconv.Atoi(input)
			if err != nil {
				return "", fmt.Errorf("%w", err)
			}
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
		if runeInput[i] == ' ' {
			spaceCounter++
		}
	}

	if spaceCounter == len(runeInput) || input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	regexp := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	numbers := regexp.FindAllString(input, -1)
	result := 0
	var arr []int
	for i := 0; i < len(numbers); i++ {
		number, err := strconv.Atoi(numbers[i])
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}
		arr = append(arr, number)
	}
	if len(arr) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	var sign1, sign2 int
	lastIndex := 0
	for i := 0; i < len(runeInput); i++ {
		if runeInput[i] == '-' {
			sign1 = -1
		}
		if runeInput[i] <= 57 && runeInput[i] >= 48 {
			lastIndex = i
			break
		}
	}

	for i := lastIndex; i < len(runeInput); i++ {
		if runeInput[i] > 57 || runeInput[i] < 48 {
			lastIndex = i
			break
		}
	}
	for i := lastIndex; i < len(runeInput); i++ {
		if runeInput[i] == '-' {
			sign2 = -1
		}
		if runeInput[i] <= 57 && runeInput[i] >= 48 {
			lastIndex = i
			break
		}
	}
	if sign1 == 0 {
		sign1 = 1
	}
	if sign2 == 0 {
		sign2 = 1
	}
	if arr[0] < 0 {
		arr[0] = arr[0] * -1
	}
	if arr[1] < 0 {
		arr[1] = arr[1] * -1
	}
	arr[0] = arr[0] * sign1
	arr[1] = arr[1] * sign2
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	return strconv.Itoa(result), nil
}
