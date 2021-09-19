package main

import (
	. "github.com/svetlanlka/golangtp/uniq"
	// . "GolangTP/uniqfunc"
	"bufio"
	"flag"
	"fmt"
	"os"
)

const filepath string = "testfiles/"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if IsTrue(Op.WithNumber)+IsTrue(Op.RepeatedLines)+IsTrue(Op.NoRepeatedLines) > 1 {
		flag.Usage()
		return
	}

	input, fileIn := GetReader(filepath, flag.Arg(0))
	output, fileOut := GetWriter(filepath, flag.Arg(1))
	defer fileIn.Close()
	defer fileOut.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	reader := functors.NewReaderMock()

	for scanner.Scan() {
		functors.OutputRead(reader, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
		panic(err)
	}

	var values []string = Uniq(reader.GetValues(), Op)

	for _, value := range values {
		writer.WriteString(value)
	}
	writer.Flush()
}
