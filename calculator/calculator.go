package calculator

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Calculator(line string) (string, error) {
	prior := make(map[string]int)
	priorityInit(prior)

	var (
		numbers, signs  []string
		current         string = ""
		isNumber        bool
		isCalculated    bool  = false
		isBracketClosed bool  = true
		err             error = nil
	)

	for i := 0; i < utf8.RuneCountInString(line); {
		current, isNumber, err = getNextElem(line, &i, current)
		fmt.Println("get line ", i, ":", current)
		if err != nil {
			return "", err
		}
		if current == "" {
			break
		}

		if isNumber == false || current == "" {
			if current == "(" {
				signs = append(signs, current)
				continue
			}
			numbers, signs, err = putSignAndTryDoOperation(current, signs, numbers, prior, &isBracketClosed, &isCalculated)
			if err != nil {
				return "", err
			}
		} else {
			numbers = append(numbers, current)
		}
	}
	isCalculated = true
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
	prior map[string]int, isBracketClosed, isCalculated *bool) ([]string, []string, error) {
	var err error
	if current == ")" {
		*isBracketClosed = false
	}

	for !canToPutSign(current, signs, prior) {
		*isCalculated = true
		signs, *isCalculated = deleteBrackets(current, signs)
		if *isCalculated {
			*isBracketClosed = true
			break
		}
		numbers, signs, *isCalculated, err = calcStackChange(numbers, signs)
		if err != nil {
			return numbers, signs, err
		}
		if !(*isCalculated) {
			break
		}
	}
	if !(*isBracketClosed) {
		return numbers, signs, errors.New("no open brackets")
	}
	if IsSign(current) {
		signs = append(signs, current)
	}
	return numbers, signs, nil
}
