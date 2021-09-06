package uniqfunc

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func GetReader(filepath string) (io.Reader, *os.File) {
	var (
		input io.Reader
		file  *os.File
	)

	if filename := filepath + flag.Arg(0); filename != filepath {
		f, err := os.Open(filename)
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

func GetWriter(filepath string) (io.Writer, *os.File) {
	var (
		input io.Writer
		file  *os.File
	)

	if filename := filepath + flag.Arg(1); filename != filepath {
		f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
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
