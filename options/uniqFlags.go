package options

import "flag"

type Options struct {
	NoRepeatedLines, RepeatedLines, WithNumber, IgnoreSymCase bool
	NumFieldsIgnore, NumCharsIgnore                           int
}

var Op Options

func init() {
	flag.BoolVar(&(Op.WithNumber), "c", false, "the number of occurrences of lines in the inpu")
	flag.BoolVar(&(Op.RepeatedLines), "d", false, "print only those lines that were repeated inhe inut data")
	flag.BoolVar(&(Op.NoRepeatedLines), "u", false, "print only thos lines that have not been repead in the input data")
	flag.IntVar(&(Op.NumFieldsIgnore), "f", 0, "ignore the first num_fields fields in the line")
	flag.IntVar(&(Op.NumCharsIgnore), "s", 0, "ignore the first num_chars characters in the string")
	flag.BoolVar(&(Op.IgnoreSymCase), "i", false, "case-insensitve")
}
