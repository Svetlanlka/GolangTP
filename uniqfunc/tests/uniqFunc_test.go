package uniqfunc_test

import (
	//	. "GolangTP/options"

	"GolangTP/options"
	. "GolangTP/uniqfunc"
	"GolangTP/uniqfunc/tests/testcases"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniq(t *testing.T) {
	testCases := []struct {
		name    string
		in      string
		options options.Options
		out     string
	}{
		{
			name:    "Without params",
			in:      testcases.TestIn1,
			options: options.Options{},
			out:     testcases.TestOut1,
		},
		{
			name:    "With count of str",
			in:      testcases.TestIn2,
			options: options.Options{WithNumber: true},
			out:     testcases.TestOut2,
		},
		{
			name:    "Out repeat lines",
			in:      testcases.TestIn3,
			options: options.Options{RepeatedLines: true},
			out:     testcases.TestOut3,
		},
		{
			name:    "Out no repeat lines",
			in:      testcases.TestIn4,
			options: options.Options{NoRepeatedLines: true},
			out:     testcases.TestOut4,
		},
		{
			name:    "Case ignore",
			in:      testcases.TestIn5,
			options: options.Options{IgnoreSymCase: true},
			out:     testcases.TestOut5,
		},
		{
			name:    "Ignore fields",
			in:      testcases.TestIn6,
			options: options.Options{NumFieldsIgnore: 1},
			out:     testcases.TestOut6,
		},
		{
			name:    "Ignore chars",
			in:      testcases.TestIn7,
			options: options.Options{NumCharsIgnore: 1},
			out:     testcases.TestOut7,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, strings.SplitAfter(test.out, "\n"), Uniq(strings.Split(test.in, "\n"), test.options))
		})
	}
}
