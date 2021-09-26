package main

import (
	"github.com/svetlanlka/golangtp/readwrite"
	"github.com/svetlanlka/golangtp/readwrite/functors"
	. "github.com/svetlanlka/golangtp/uniq"
	"github.com/svetlanlka/golangtp/uniq/options"

	"flag"
)

const filepath string = "testfiles/"

func main() {
	flag.Parse()
	if !options.CheckFlags(options.Op) {
		return
	}
	reader := functors.NewReaderMock()
	if !readwrite.Read(filepath, flag.Arg(0), reader) {
		return
	}

	var values []string = Uniq(reader.GetValues(), options.Op)

	writer := functors.NewWriterMock()
	if !readwrite.Write(filepath, flag.Arg(1), writer, values) {
		return
	}
}
