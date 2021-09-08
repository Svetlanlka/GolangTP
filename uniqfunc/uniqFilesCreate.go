package uniqfunc

import (
	"fmt"
	"io"
	"os"
)

func GetReader(filepath, filename string) (io.Reader, *os.File) {
	var (
		input io.Reader
		file  *os.File
	)

	if filename != "" {
		f, err := os.Open(filepath + filename)
		if err != nil {
			fmt.Printf("File %s not open\n", filename)
			panic(err)
		}

		input = f
		file = f
	} else {
		input = os.Stdin
	}

	return input, file
}

func GetWriter(filepath, filename string) (io.Writer, *os.File) {
	var (
		input io.Writer
		file  *os.File
	)

	if filename != "" {
		f, err := os.OpenFile(filepath+filename, os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("File %s not open\n", filename)
			panic(err)
		}

		input = f
		file = f
	} else {
		input = os.Stdout
	}

	return input, file
}
