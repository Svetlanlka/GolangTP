package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func isSign(sym string) bool {
	if sym == plus || sym == minus || sym == multiplication || sym == division {
		return true
	}

	return false
}

func calculateExpr(first, second, sign string) (string, error) {
	result := 0.0

	a, err1 := strconv.ParseFloat(first, 64)
	if err1 != nil {
		return "", errors.New("cannot convert first operand to float")
	}

	b, err2 := strconv.ParseFloat(second, 64)
	if err2 != nil {
		return "", errors.New("cannot convert second operand to float")
	}

	switch sign {
	case plus:
		result = a + b
	case minus:
		result = a - b
	case multiplication:
		result = a * b
	case division:
		result = a / b
	default:
		return "", errors.New("wrong operation")
	}

	return strconv.FormatFloat(result, 'f', -1, 64), nil
}

func calcStackChange(numbers, signs []string) ([]string, []string, bool, error) {
	lastSignIdx := len(signs) - 1
	lastNumberIdx := len(numbers) - 1

	if lastNumberIdx < 0 || lastNumberIdx < 1 {
		return numbers, signs, false, nil
	}

	result, err := calculateExpr(numbers[lastNumberIdx-1], numbers[lastNumberIdx], signs[lastSignIdx])
	numbers = numbers[:(lastNumberIdx - 1)]
	signs = signs[:(lastSignIdx)]
	numbers = append(numbers, result)

	return numbers, signs, true, err
}

func checkSign(cur string, signs []string) bool {
	if cur == rightBracket {
		return false
	}
	if len(signs) < 1 {
		return true
	}
	if signs[len(signs)-1] == leftBracket {
		return true
	}

	return prior[cur] > prior[signs[len(signs)-1]]
}

func deleteBrackets(current string, signs []string) ([]string, bool) {
	if current == rightBracket && (len(signs) > 0 && signs[len(signs)-1] == leftBracket) {
		signs = signs[:(len(signs) - 1)]
		return signs, true
	}

	return signs, false
}

func checkAndGetNextElem(line []string, i int, last string) (string, int, bool, error) {
	result := ""
	len := len(line)

	for ; i < len; i++ {
		sym := line[i]
		symRune := []rune(sym)[0]

		if unicode.IsSpace(symRune) {
			continue
		}
		if sym == "." {
			if result == "" || i+1 >= len || !unicode.IsDigit([]rune(line[i+1])[0]) {
				return "", i, false, errors.New("point must be after or before digit")
			}
			result += sym
			continue
		}
		if isSign(sym) {
			if (last == "" || isSign(last) || last == leftBracket) && (sym == minus || sym == plus) && result == "" {
				result += sym
				continue
			}
			i++
			return sym, i, false, nil
		}
		if sym == leftBracket {
			if last == "" || isSign(last) || last == leftBracket {
				i++
				return sym, i, false, nil
			}
			return "", i, false, errors.New("before '(' not correct symbol")
		}
		if sym == rightBracket {
			if !(last == "" || isSign(last)) {
				i++
				return sym, i, false, nil
			}
			return "", i, false, errors.New("before ')' not correct symbol")
		}
		if unicode.IsDigit(symRune) {
			fmt.Println("last ", last)
			if !(last == "" || isSign(last) || last == leftBracket) {
				return "", i, false, errors.New("before digit not correct symbol")
			}
			result += sym
			if i+1 < len && (!unicode.IsDigit([]rune(line[i+1])[0]) && line[i+1] != ".") {
				i++
				break
			}
			continue
		}

		return "", i, false, errors.New("error, wrong symbol in str")
	}

	return result, i, true, nil
}
