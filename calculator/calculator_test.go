package calculator_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/svetlanlka/golangtp/calculator"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  string
		err  bool
	}{
		{"Basic calculate sum", "1+2", "3", false},
		{"Basic calculate difference", "1-2", "-1", false},
		{"Basic calculate multiplication", "10*2", "20", false},
		{"Basic calculate division", "4/2", "2", false},
		{"Basic float calculate sum", "1.2+2.55", "3.75", false},
		{"Basic float calculate difference", "1.1-234", "-232.9", false},
		{"Basic float calculate multiplication", "1.1*1.1", "1.21", false},
		{"Basic float calculate division", "1/8", "0.125", false},

		{"Operation with brackets: simple parse brackets", "(1)", "1", false},
		{"Operation with brackets: parse minus in brackets", "(-1)", "-1", false},
		{"Operation with brackets: parse operation sum in brackets", "(1+1)", "2", false},
		{"Operation with brackets: parse operation difference in brackets", "(1-1)", "0", false},
		{"Operation with brackets: parse operation multiplication in brackets", "(2*3)", "6", false},
		{"Operation with brackets: parse operation division in brackets", "(3/2)", "1.5", false},
		{"Operation with brackets: parse znak in brackets", "1+(-2)", "-1", false},

		{"Parse znak in number: first operand positive, second negative", "+1--2", "3", false},
		{"Parse znak in number: first operand positive, second positive", "+1.1 ++2.1", "3.2", false},
		{"Parse znak in number: first operand negative, second positive", "-1.1+20", "18.9", false},
		{"Parse znak in number: first operand negative, second negative", "-1.1--2", "0.9", false},

		{"Priority: +*", "10+12*7", "94", false},
		{"Priority: +/", "10+9/2", "14.5", false},
		{"Priority: brackets", "(10+12)*7", "154", false},
		{"Priority: * combo", "3*4*5", "60", false},
		{"Priority: / combo", "12/2/6", "1", false},
		{"Priority: + combo", "3+4+5", "12", false},
		{"Priority: - combo", "3-4-5.5", "-6.5", false},

		{"Zero: division", "12/0", "infinity", false},
		{"Zero: multiplication", "12*0", "0", false},

		{"Multiply test 1", "1 + (5 + 2 * (4 + 5)) + 1", "25", false},
		{"Multiply test 2", "10 + (3*3 + 1)", "20", false},
		{"Bracket spam", "(1)+(2)-(3)+(4)-(5)+(6)-(7)+(8)+(0)-(0)+(0)-(0)+(0)-(0)", "6", false},
		{"Inner bracket spam", "(((((((((1)))))))))", "1", false},

		{"Can not read symbols", "3re+59", "", true},
		{"Empty expression", "", "", true},
		{"Unknown operation", "10^2", "", true},

		{"Not valid syntax with point: point must go only after numbers", "3 + .3 + 1", "", true},
		{"Not valid syntax with point: after point must be number", "3. + (1 + 1)", "", true},
		{"Not valid syntax with point: point can not be in begin of str", ".3 + 1", "", true},
		{"Not valid syntax with point: point can not be in end of str", "3 + 1.", "", true},

		{"Not valid syntax with brackets: before '(' can not be digit", "3(+1)", "", true},
		{"Not valid syntax with brackets: before '(' can not be ')'", "()(+1)", "", true},
		{"Not valid syntax with brackets: ')' can not be in begin of str", ") +1", "", true},
		{"Not valid syntax with brackets: befor ')' can not digit", "1 + 5 +) 1", "", true},
		{"Not valid syntax with brackets: brackets not close", "(1 + 5 + 1", "", true},
		{"Not valid syntax with brackets: brackets not open", "1 + 5 + 1)", "", true},

		{"Not valid syntax with operands: no left operand", "*12", "", true},
		{"Not valid syntax with operands: no right operand", "175+", "", true},
		{"Not valid syntax with operands: too many operands", "(20+1.1)+----90+1", "", true},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculator.Calculator(test.in)
			if test.err {
				if err == nil {
					t.Error("Function must be finished with error")
				}
			} else {
				if err != nil {
					t.Error("Function finished with error: " + err.Error())
				} else {
					require.Equal(t, roundNumber(result, 5), roundNumber(test.out, 5))
				}
			}

		})
	}
}

func roundNumber(number string, precision float64) float64 {
	floatNumber, _ := strconv.ParseFloat(number, 64)
	prec := math.Pow(10, precision)
	return math.Round((floatNumber * prec) / prec)
}
