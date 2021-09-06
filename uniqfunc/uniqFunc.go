package uniqfunc

import (
	"bufio"
	"strconv"
	"strings"
)

func IsTrue(value bool) int8 {
	if value {
		return 1
	}
	return 0
}

func deleteFieldsInStr(str string) string {
	arr := strings.Split(str, " ")

	if NumFieldsIgnore < len(arr) {
		return strings.Join(arr[NumFieldsIgnore:], " ")
	}
	return ""
}

func deleteSymInStr(str string) string {
	arr := strings.Split(str, "")

	if NumCharsIgnore < len(arr) {
		return strings.Join(arr[NumCharsIgnore:], "")
	}
	return ""
}

func StrIsEqual(str1, str2 string) bool {
	if NumFieldsIgnore > 0 {
		str1 = deleteFieldsInStr(str1)
		str2 = deleteFieldsInStr(str2)
	}
	if NumCharsIgnore > 0 {
		str1 = deleteSymInStr(str1)
		str2 = deleteSymInStr(str2)
	}

	if !IgnoreSymCase && str1 != str2 || (IgnoreSymCase && !strings.EqualFold(str1, str2)) {
		return true
	}
	return false
}

func WriteStr(currentNumber int, currentStr string, writer *bufio.Writer) {
	if WithNumber && currentNumber != 0 ||
		RepeatedLines && currentNumber > 1 ||
		NoRepeatedLines && currentNumber == 1 ||
		!(WithNumber || RepeatedLines || NoRepeatedLines) {
		if WithNumber {
			writer.WriteString(strconv.Itoa(currentNumber) + " ")
		}
		writer.WriteString(currentStr)
		writer.WriteString("\n")
	}
}
