package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	. "GolangTP/uniqfunc"
)

const filepath string = "testfiles/"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if IsTrue(WithNumber)+IsTrue(RepeatedLines)+IsTrue(NoRepeatedLines) > 1 {
		flag.Usage()
		return
	}

	fmt.Print("\twithNumber: ", WithNumber, "\trepeatedLines: ", RepeatedLines,
		"\tnoRepeatedLines: ", NoRepeatedLines, "\n\tnumFieldsIgnore: ", NumFieldsIgnore,
		"\n\tnumCharsIgnore: ", NumCharsIgnore, "\n\tignoreSymCase: ", IgnoreSymCase, "\n")

	input, fileIn := GetReader(filepath)
	output, fileOut := GetWriter(filepath)
	defer fileIn.Close()
	defer fileOut.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	var (
		currentStr    string = ""
		currentNumber int    = 0
	)

	fmt.Print("RESULT\n")
	for scanner.Scan() {
		if StrIsEqual(scanner.Text(), currentStr) {
			if currentNumber > 0 {
				WriteStr(currentNumber, currentStr, writer)
			}
			currentStr = scanner.Text()
			currentNumber = 0
		}
		currentNumber++
	}
	WriteStr(currentNumber, currentStr, writer)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
		panic(err)
	}

	writer.Flush()
}
