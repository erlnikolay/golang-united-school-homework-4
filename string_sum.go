package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the expression has at least one whitespace character
	errorHasSpaceChar = errors.New("input has at least one whitespace character")
	// OPerand is not number
	errorIsNotNumber = errors.New("Operand is not number")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and
// when input string contains whitespace (" 3 + 5 ")
//
// For the cases, when the input expression is not valid(contains characters,
// that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package
// wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
func countOfOperand(input string) (err error) {
	var modifyOperations string
	//fmt.Println(input)
	modifyOperations = strings.ReplaceAll(input, "+", "#")
	modifyOperations = strings.ReplaceAll(modifyOperations, "-", "#")
	//fmt.Println(modifyOperations)
	//fmt.Println(strings.Count(modifyOperations, "#"))
	if strings.Count(modifyOperations, "#") > 1 || strings.Count(modifyOperations, "#") < 1 {
		return fmt.Errorf("%v", errorNotTwoOperands)
	} else {
		return nil
	}
}

func StringSum(input string) (output string, err error) {
	var operandOne int
	var operandTwo int
	var calcResult int = 0
	var firstMinusTrigger bool // false - plus, true - minus
	var operationTrigger bool  // false - plus, true - minus
	//fmt.Println(input)
	if len(input) == 0 {
		output = ""
		//err = errorEmptyInput
		return output, fmt.Errorf("%v", errorEmptyInput)
	}
	if strings.Index(input, " ") >= 0 {
		output = ""
		//err = errorHasSpaceChar
		return output, fmt.Errorf("%v", errorHasSpaceChar)
	}
	// detele first plus from string
	firstMinusTrigger = false
	if len(input) != 0 && strings.Index(input, "+") == 0 {
		input = string([]rune(input)[1:])
	}
	if len(input) != 0 && strings.Index(input, "-") == 0 {
		input = string([]rune(input)[1:])
		firstMinusTrigger = true
	}
	rErr := countOfOperand(input)
	if rErr != nil {
		return "", rErr
	}
	// take operators
	operationTrigger = false
	var indexOfOperation int
	if strings.Index(input, "+") > 0 {
		indexOfOperation = strings.Index(input, "+")
		//fmt.Println(indexOfOperation)
		operandOne, err = strconv.Atoi(string([]rune(input)[:indexOfOperation]))
		//fmt.Println(operandOne)
		if err != nil {
			//err = errorIsNotNumber
			return "", fmt.Errorf("%v", errorIsNotNumber)
		}
		operandTwo, err = strconv.Atoi(string([]rune(input)[(indexOfOperation + 1):]))
		//fmt.Println(operandTwo)
		if err != nil {
			//err = errorIsNotNumber
			return "", fmt.Errorf("%v", errorIsNotNumber)
		}
	} else if strings.Index(input, "-") > 0 {
		indexOfOperation = strings.Index(input, "-")
		//fmt.Println(indexOfOperation)
		operandOne, err = strconv.Atoi(string([]rune(input)[:indexOfOperation]))
		if err != nil {
			//err = errorIsNotNumber
			return "", fmt.Errorf("%v", errorIsNotNumber)
		}
		operandTwo, err = strconv.Atoi(string([]rune(input)[(indexOfOperation + 1):]))
		if err != nil {
			//err = errorIsNotNumber
			return "", fmt.Errorf("%v", errorIsNotNumber)
		}
		operationTrigger = true
	}
	//fmt.Println("operations: ")
	//fmt.Printf("first symbol operation trigger: %b\n", firstMinusTrigger)
	//fmt.Printf("operation trigger: %b\n", operationTrigger)
	if firstMinusTrigger == true { // first minus
		if operationTrigger == true { // operation minus
			calcResult = -operandOne - operandTwo
		} else { // operation plus
			calcResult = -operandOne + operandTwo
		}
	} else { // first plus
		if operationTrigger == true { // operation minus
			calcResult = operandOne - operandTwo
		} else { // opeartion plus
			calcResult = operandOne + operandTwo
		}
	}
	//fmt.Printf("result: %d\n", calcResult)
	if calcResult != 0 {
		return strconv.Itoa(calcResult), nil
	} else {
		//err = errorEmptyInput
		return "", fmt.Errorf("%v", errorEmptyInput)
	}
}

//func main() {
//	var res string
//	var err error
//
//	res, err = StringSum("+10-7")
//	fmt.Printf("Common result is %s\n", res)
//	if err != nil {
//		fmt.Printf("error: %s\n", err)
//	}
//}
