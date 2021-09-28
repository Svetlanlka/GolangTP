package readwrite

import (
	"bufio"
	"fmt"
	"os"

	"github.com/svetlanlka/golangtp/readwrite/functors"
)

func Read(filepath, filename string, reader *functors.ReaderMock) bool {
	input := os.Stdin
	if filename != "" {
		f, err := os.Open(filepath + filename)
		if err != nil {
			fmt.Printf("File %s not open\n", filename)
			return false
		}
		defer f.Close()
		input = f
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		functors.OutputRead(reader, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
		return false
	}

	return true
}

func Write(filepath, filename string, writer *functors.WriterMock, values []string) bool {
	output := os.Stdout

	if filename != "" {
		f, err := os.OpenFile(filepath+filename, os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("File %s not open\n", filename)
			return false
		}
		defer f.Close()

		output = f
	}

	bufioWriter := bufio.NewWriter(output)
	for _, value := range values {
		bufioWriter.WriteString(value)
	}
	bufioWriter.Flush()

	return true
}
