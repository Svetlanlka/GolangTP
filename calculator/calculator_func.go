package calculator

import (
	"errors"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func IsSign(sym string) bool {
	if sym == "/" || sym == "*" || sym == "+" || sym == "-" {
		return true
	}
	return false
}

func calculateExpr(value1, value2, sign string) (string, error) {
	var result float64
	a, err1 := strconv.ParseFloat(value1, 64)
	b, err2 := strconv.ParseFloat(value2, 64)
	if err1 != nil || err2 != nil {
		return "", errors.New("not parse float")
	}

	switch sign {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return "", errors.New("wrong operation")
	}
	return strconv.FormatFloat(result, 'f', -1, 64), nil
}

func priorityInit(priority map[string]int) {
	priority["+"] = 1
	priority["-"] = 1
	priority["*"] = 2
	priority["/"] = 2
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

// return true, if we can put into stack new sign
func canToPutSign(cur string, signs []string, prior map[string]int) bool {
	if cur == ")" {
		return false
	}
	if len(signs) < 1 {
		return true
	}
	if signs[len(signs)-1] == "(" {
		return true
	}

	return prior[cur] > prior[signs[len(signs)-1]]
}

func deleteBrackets(current string, signs []string) ([]string, bool) {
	if current == ")" && (len(signs) > 0 && signs[len(signs)-1] == "(") {
		signs = signs[:(len(signs) - 1)]
		return signs, true
	}
	return signs, false
}

func getNextElem(line string, i *int, last string) (string, bool, error) {
	var result string = ""
	var len int = utf8.RuneCountInString(line)

	for ; *i < len; (*i)++ {
		var sym = rune(line[*i])

		if unicode.IsSpace(sym) {
			continue
		}
		if string(sym) == "." {
			if result == "" || *i+1 >= len || !unicode.IsDigit(rune(line[(*i)+1])) {
				return "", false, errors.New("point must be after or before digit")
			}
			result += string(sym)
			continue
		}
		if IsSign(string(sym)) {
			if (last == "" || IsSign(last) || last == "(") && (string(sym) == "-" || string(sym) == "+") && result == "" {
				result += string(sym)
				continue
			}
			(*i)++
			return string(sym), false, nil
		}
		if string(sym) == "(" {
			if last == "" || IsSign(last) || last == "(" {
				(*i)++
				return string(sym), false, nil
			}
			return "", false, errors.New("before '(' not correct symbol")
		}
		if string(sym) == ")" {
			if !(last == "" || IsSign(last)) {
				(*i)++
				return string(sym), false, nil
			}
			return "", false, errors.New("before ')' not correct symbol")
		}
		if unicode.IsDigit(sym) {
			if !(last == "" || IsSign(last) || last == "(") {
				return "", false, errors.New("before digit not correct symbol")
			}
			result += string(sym)
			if *i+1 < len && (!unicode.IsDigit(rune(line[(*i)+1])) && string(line[(*i)+1]) != ".") {
				(*i)++
				break
			}
			continue
		}
		return "", false, errors.New("error, wrong symbol in str")
	}

	return result, true, nil
}
