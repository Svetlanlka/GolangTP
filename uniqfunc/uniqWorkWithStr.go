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
	arr := strings.Split(str, "")

	if op.NumCharsIgnore < len(arr) {
		return strings.Join(arr[op.NumCharsIgnore:], "")
	}
	return ""
}

func StrIsEqual(str1, str2 string, op Options) bool {
	if op.NumFieldsIgnore > 0 {
		str1 = deleteFieldsInStr(str1, op)
		str2 = deleteFieldsInStr(str2, op)
	}
	if op.NumCharsIgnore > 0 {
		str1 = deleteSymInStr(str1, op)
		str2 = deleteSymInStr(str2, op)
	}

	if !op.IgnoreSymCase && str1 != str2 || (op.IgnoreSymCase && !strings.EqualFold(str1, str2)) {
		return true
	}
	return false
}

func WriteStr(currentNumber int, currentStr string, writer *functors.WriterMock, op Options, eof bool) {
	if op.WithNumber && currentNumber != 0 ||
		op.RepeatedLines && currentNumber > 1 ||
		op.NoRepeatedLines && currentNumber == 1 ||
		!(op.WithNumber || op.RepeatedLines || op.NoRepeatedLines) {
		if op.WithNumber {
			currentStr = strconv.Itoa(currentNumber) + " " + currentStr
		}
		if !eof {
			currentStr += "\n"
		}
		functors.OutputWrite(writer, currentStr)
	}
}
