package uniqfunc_test

import (
	"GolangTP/calculator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniq(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "Calculate",
			in:   "(1+2)-3",
			out:  "0",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.out, calculator.Calculator(test.in))
		})
	}
}
