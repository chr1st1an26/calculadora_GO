package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("Calculadora perrona")
		fmt.Println("S-Suma")
		fmt.Println("R-Resta")
		fmt.Println("M-Multiplicación")
		fmt.Println("D-División")
		fmt.Println("Q-Salir")

		var choice string
		fmt.Print("Elije una opción S,R,M,D,E: ")
		fmt.Scanln(&choice)
		choice = strings.ToUpper(choice)

		switch choice {
		case "S":
			performOperation('+')
		case "R":
			performOperation('-')
		case "M":
			performOperation('*')
		case "D":
			performOperation('/')
		case "E":
			fmt.Println("adios guapo")
			return
		default:
			fmt.Println("eso no se puede bb")
		}
	}
}

func performOperation(operator rune) {
	var num1, num2 float64
	fmt.Print("pon el primer numero inutil: ")
	fmt.Scanln(&num1)

	fmt.Print("pon el segundo numero inutil: ")
	fmt.Scanln(&num2)

	var result float64

	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("no puedes dividir numeros sin valor.")
			return
		}
	}

	resultStr := strconv.FormatFloat(result, 'f', -1, 64)
	if strings.Contains(resultStr, ".") {
		fmt.Printf("total bb: %s\n", resultStr)
	} else {
		intResult, _ := strconv.Atoi(resultStr)
		fmt.Printf("resultado woooa: %d\n", intResult)
	}
}
