package uniq

import (
	"strconv"
	"strings"

	"github.com/svetlanlka/golangtp/readwrite/functors"
	"github.com/svetlanlka/golangtp/uniq/options"
)

func deleteFieldsInLines(str string, op options.Options) string {
	arr := strings.Split(str, " ")

	if op.NumFieldsIgnore < len(arr) {
		return strings.Join(arr[op.NumFieldsIgnore:], " ")
	}
	return ""
}

func deleteCharInLines(str string, op options.Options) string {
	arr := strings.Split(str, " ")

	for i, value := range arr {
		splitValue := strings.Split(value, "")
		if op.NumCharsIgnore < len(splitValue) {
			arr[i] = strings.Join(splitValue[op.NumCharsIgnore:], "")
			continue
		}
		arr[i] = ""
	}

	return strings.Join(arr, "")
}

func linesIsEqual(str1, str2 string, op options.Options) bool {
	str1 = strings.Trim(str1, " ")
	str2 = strings.Trim(str2, " ")

	if op.NumFieldsIgnore > 0 {
		str1 = deleteFieldsInLines(str1, op)
		str2 = deleteFieldsInLines(str2, op)
	}
	if op.NumCharsIgnore > 0 {
		str1 = deleteCharInLines(str1, op)
		str2 = deleteCharInLines(str2, op)
	}

	if !op.IgnoreSymCase && str1 != str2 || (op.IgnoreSymCase && !strings.EqualFold(str1, str2)) {
		return false
	}
	return true
}

func writeLine(repeatedStrNumber int, line string, writer *functors.WriterMock, op options.Options) {
	if strings.Trim(line, " ") != "" {
		functors.OutputWrite(writer, line)
	}
}

func lineSatisfiesConditionWithFlags(repeatedStrNumber int, previousStr string, op options.Options) bool {
	if op.WithNumber && repeatedStrNumber != 0 ||
		op.RepeatedLines && repeatedStrNumber > 1 ||
		op.NoRepeatedLines && repeatedStrNumber == 1 ||
		!(op.WithNumber || op.RepeatedLines || op.NoRepeatedLines) {
		return true
	}
	return false
}

func addNumberOfRepeatedLines(repeatedStrNumber int, previousStr string, op options.Options) string {
	if !op.WithNumber {
		return previousStr
	}

	number := strconv.Itoa(repeatedStrNumber)
	if strings.Trim(previousStr, " ") == "" {
		return number
	}

	return number + " " + previousStr
}
