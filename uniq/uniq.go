package uniq

import (
	"github.com/svetlanlka/golangtp/readwrite/functors"
	"github.com/svetlanlka/golangtp/uniq/options"
)

func Uniq(data []string, op options.Options) []string {
	previousStr := ""
	repeatedStrNumber := 0

	writer := functors.NewWriterMock()

	for _, currentStr := range data {
		if !linesIsEqual(currentStr, previousStr, op) {
			if repeatedStrNumber > 0 {
				if lineSatisfiesConditionWithFlags(repeatedStrNumber, previousStr, op) {
					lineToWrite := addNumberOfRepeatedLines(repeatedStrNumber, previousStr, op)
					lineToWrite += "\n"
					writeLine(repeatedStrNumber, lineToWrite, writer, op)
				}
			}

			repeatedStrNumber = 0
			previousStr = currentStr
		}

		repeatedStrNumber++
	}
	if lineSatisfiesConditionWithFlags(repeatedStrNumber, previousStr, op) {
		lineToWrite := addNumberOfRepeatedLines(repeatedStrNumber, previousStr, op)
		writeLine(repeatedStrNumber, lineToWrite, writer, op)
	}

	return writer.GetValues()
}
