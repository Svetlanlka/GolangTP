package main

import (
	"GolangTP/calculator"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var expression string
	in := bufio.NewReader(os.Stdin)
	expression, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}

	fmt.Println(calculator.Calculator(expression))
}
