package uniq

import (
	"github.com/svetlanlka/golangtp/readwrite/functors"
	"github.com/svetlanlka/golangtp/uniq/options"
)

func Uniq(data []string, op options.Options) []string {
	previousStr := ""
	repeatedStrNumber := -1

	writer := functors.NewWriterMock()

	for _, currentStr := range data {
		repeatedStrNumber++

		if linesIsEqual(currentStr, previousStr, op) {
			continue
		}

		if repeatedStrNumber <= 0 || !lineSatisfiesConditionWithFlags(repeatedStrNumber, previousStr, op) {
			repeatedStrNumber = 0
			previousStr = currentStr
			continue
		}

		lineToWrite := addNumberOfRepeatedLines(repeatedStrNumber, previousStr, op)
		lineToWrite += "\n"
		writeLine(repeatedStrNumber, lineToWrite, writer, op)
		repeatedStrNumber = 0
		previousStr = currentStr
	}

	repeatedStrNumber++
	if lineSatisfiesConditionWithFlags(repeatedStrNumber, previousStr, op) {
		lineToWrite := addNumberOfRepeatedLines(repeatedStrNumber, previousStr, op)
		writeLine(repeatedStrNumber, lineToWrite, writer, op)
	}

	return writer.GetValues()
}
