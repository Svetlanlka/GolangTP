package uniqfunc

import (
	. "GolangTP/options"
	"GolangTP/uniqfunc/functors"
	"strconv"
	"strings"
)

func IsTrue(value bool) int8 {
	if value {
		return 1
	}
	return 0
}

func deleteFieldsInStr(str string, op Options) string {
	arr := strings.Split(str, " ")

	if op.NumFieldsIgnore < len(arr) {
		return strings.Join(arr[op.NumFieldsIgnore:], " ")
	}
	return ""
}

func deleteSymInStr(str string, op Options) string {
	arr := strings.Split(str, " ")

	for i, value := range arr {
		splitValue := strings.Split(value, "")
		if op.NumCharsIgnore < len(splitValue) {
			arr[i] = strings.Join(splitValue[op.NumCharsIgnore:], "")
		} else {
			arr[i] = ""
		}
	}

	return strings.Join(arr, "")
}

func StrIsEqual(str1, str2 string, op Options) bool {
	str1 = strings.Trim(str1, " ")
	str2 = strings.Trim(str2, " ")

	if op.NumFieldsIgnore > 0 {
		str1 = deleteFieldsInStr(str1, op)
		str2 = deleteFieldsInStr(str2, op)
	}
	if op.NumCharsIgnore > 0 {
		str1 = deleteSymInStr(str1, op)
		str2 = deleteSymInStr(str2, op)
	}

	if !op.IgnoreSymCase && str1 != str2 || (op.IgnoreSymCase && !strings.EqualFold(str1, str2)) {
		return false
	}
	return true
}

func WriteStr(currentNumber int, currentStr string, writer *functors.WriterMock, op Options, eof bool) {
	if op.WithNumber && currentNumber != 0 ||
		op.RepeatedLines && currentNumber > 1 ||
		op.NoRepeatedLines && currentNumber == 1 ||
		!(op.WithNumber || op.RepeatedLines || op.NoRepeatedLines) {
		if op.WithNumber {
			number := strconv.Itoa(currentNumber)
			if strings.Trim(currentStr, " ") == "" {
				currentStr = number
			} else {
				currentStr = number + " " + currentStr
			}
		}
		if !eof {
			currentStr += "\n"
		}
		if strings.Trim(currentStr, " ") != "" {
			functors.OutputWrite(writer, currentStr)
		}
	}
}
