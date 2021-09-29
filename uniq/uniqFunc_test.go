package uniq

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/svetlanlka/golangtp/uniq/options"
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
			in:      TestInWithoutParams,
			options: options.Options{},
			out:     TestOutWithoutParams,
		},
		{
			name:    "With count of str",
			in:      TestInWithCount,
			options: options.Options{WithNumber: true},
			out:     TestOutWithCount,
		},
		{
			name:    "Out repeat lines",
			in:      TestInRepeatLines,
			options: options.Options{RepeatedLines: true},
			out:     TestOutRepeatLines,
		},
		{
			name:    "Out no repeat lines",
			in:      TestInNoRepeatLines,
			options: options.Options{NoRepeatedLines: true},
			out:     TestOutNoRepeatLines,
		},
		{
			name:    "Case ignore",
			in:      TestInCaseIgnore,
			options: options.Options{IgnoreSymCase: true},
			out:     TestOutCaseIgnore,
		},
		{
			name:    "Ignore fields",
			in:      TestInFieldsIgnore,
			options: options.Options{NumFieldsIgnore: 1},
			out:     TestOutFieldsIgnore,
		},
		{
			name:    "Ignore chars",
			in:      TestInCharsIgnore,
			options: options.Options{NumCharsIgnore: 1},
			out:     TestOutCharsIgnore,
		},
		{
			name: "Multi test",
			in:   TestInMulti,
			options: options.Options{WithNumber: true, IgnoreSymCase: true,
				NumCharsIgnore: 3, NumFieldsIgnore: 2},
			out: TestOutMulti,
		},
		{
			name:    "Whitespace lines",
			in:      TestInWhitespaceLines,
			options: options.Options{WithNumber: true, NumFieldsIgnore: 1},
			out:     TestOutWhitespaceLines,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, strings.SplitAfter(test.out, "\n"),
				Uniq(strings.Split(test.in, "\n"), test.options))
		})
	}
}

var TestInWithoutParams string = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

var TestOutWithoutParams string = `I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.`

var TestInWithCount = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
I love music of Kartik.






end`

var TestOutWithCount = `3 I love music.
1
2 I love music of Kartik.
1 Thanks.
3 I love music of Kartik.
6
1 end`

var TestInRepeatLines = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

var TestOutRepeatLines = `I love music.
I love music of Kartik.
I love music of Kartik.`

var TestInNoRepeatLines = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
Thanks.`

var TestOutNoRepeatLines = `
Thanks.
Thanks.`

var TestInCaseIgnore = `I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
I love music of kartik.
I love MuSIC of Kartik.`

var TestOutCaseIgnore = `I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.`

var TestInFieldsIgnore = `We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`

var TestOutFieldsIgnore = `We love music.

I love music of Kartik.
Thanks.`

var TestInCharsIgnore = `I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`

var TestOutCharsIgnore = `I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`

var TestInMulti = `I and ertmy 000friend hjk music.
I , ertmy 54vfriend lov music.
I , ertmy 000friend 789 music. 
I , ___my 54vfriend 6f4 music. 


dsfds fbgfb I LOVE MUSIC OF KARTIK.
_ _ o love music of kartik.
 Love music so MUCH
 Love music sO MuCh

 _Love.
 LOVE.`

var TestOutMulti = `4 I and ertmy 000friend hjk music.
2
2 dsfds fbgfb I LOVE MUSIC OF KARTIK.
2  Love music so MUCH
3`

var TestInWhitespaceLines = `                 
                  
        
 _Love.
 LOVE.`

var TestOutWhitespaceLines = `5`
