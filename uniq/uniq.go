package uniq

import (
	"github.com/svetlanlka/golangtp/readwrite/functors"
	"github.com/svetlanlka/golangtp/uniq/options"
)

func Uniq(values []string, op options.Options) []string {
	currentStr := ""
	currentNumber := 0

	writer := functors.NewWriterMock()

	for _, value := range values {
		if !linesIsEqual(value, currentStr, op) {
			if currentNumber > 0 {
				writeStr(currentNumber, currentStr, writer, op, false)
			}

			currentNumber = 0
			currentStr = value
		}

		currentNumber++
	}
	writeStr(currentNumber, currentStr, writer, op, true)

	return writer.GetValues()
}
