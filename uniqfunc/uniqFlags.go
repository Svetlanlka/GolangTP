package uniqfunc

import "flag"

var (
	NoRepeatedLines, RepeatedLines, WithNumber, IgnoreSymCase bool
	NumFieldsIgnore, NumCharsIgnore                           int
)

func init() {
	flag.BoolVar(&WithNumber, "c", false, "the number of occurrences of lines in the input")
	flag.BoolVar(&RepeatedLines, "d", false, "print only those lines that were repeated in the input data")
	flag.BoolVar(&NoRepeatedLines, "u", false, "print only those lines that have not been repeated in the input data")
	flag.IntVar(&NumFieldsIgnore, "f", 0, "ignore the first num_fields fields in the line")
	flag.IntVar(&NumCharsIgnore, "s", 0, "ignore the first num_chars characters in the string")
	flag.BoolVar(&IgnoreSymCase, "i", false, "case-insensitive")
}
