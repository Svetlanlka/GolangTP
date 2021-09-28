package calculator

import (
	"errors"
	"fmt"
	"strings"
)

func Calculator(inputLine string) (string, error) {
	prior := map[string]int{
		plus:           1,
		minus:          1,
		multiplication: 2,
		division:       2,
	}

	var (
		numbers, signs []string
		current        string
		isNumber       bool
		err            error
	)

	line := strings.Split(inputLine, "")
	fmt.Println(line, len(line))
	for i := 0; i < len(line); {
		current, isNumber, err = checkAndGetNextElem(line, &i, current)
		fmt.Println("get line ", i, ":", current)
		if err != nil {
			return "", err
		}
		if current == "" {
			break
		}

		if isNumber == false || current == "" {
			if current == leftBracket {
				signs = append(signs, current)
				continue
			}
			numbers, signs, err = putSignAndTryDoOperation(current, signs, numbers, prior)
			if err != nil {
				return "", err
			}
		} else {
			numbers = append(numbers, current)
		}
	}
	isCalculated := true
	for isCalculated {
		numbers, signs, isCalculated, err = calcStackChange(numbers, signs)
	}

	if len(numbers) < 1 {
		return "", errors.New("not enter expression")
	}
	if len(signs) > 0 {
		return "", errors.New("no close brackets or not complete operation")
	}
	return numbers[0], err
}

func putSignAndTryDoOperation(current string, signs, numbers []string,
	prior map[string]int) ([]string, []string, error) {
	var err error
	isCalculated := false
	isBracketClosed := true

	if current == rightBracket {
		isBracketClosed = false
	}

	for !checkSign(current, signs, prior) {
		isCalculated = true
		signs, isCalculated = deleteBrackets(current, signs)
		if isCalculated {
			isBracketClosed = true
			break
		}
		numbers, signs, isCalculated, err = calcStackChange(numbers, signs)
		if err != nil {
			return numbers, signs, err
		}
		if !(isCalculated) {
			break
		}
	}
	if !(isBracketClosed) {
		return numbers, signs, errors.New("no open brackets")
	}
	if isSign(current) {
		signs = append(signs, current)
	}
	return numbers, signs, nil
}
