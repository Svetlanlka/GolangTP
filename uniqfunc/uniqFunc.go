package uniqfunc

import (
	. "GolangTP/options"
	"GolangTP/uniqfunc/functors"
)

func Uniq(values []string, op Options) []string {
	var (
		currentStr    string = ""
		currentNumber int    = 0
	)
	writer := functors.NewWriterMock()

	for _, value := range values {
		if StrIsEqual(value, currentStr, op) {
			if currentNumber > 0 {
				WriteStr(currentNumber, currentStr, writer, op, false)
			}

			currentNumber = 0
			currentStr = value
		}
		
		currentNumber++
	}
	//fmt.Println("number: " + strconv.Itoa(currentNumber) + " curStr: " + currentSt)
	WriteStr(currentNumber, currentStr, writer, op, true)

	return writer.GetValues()
}
