package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/svetlanlka/golangtp/calculator"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	expression, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		return
	}

	answer, err := calculator.Calculator(expression)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(answer)
}
