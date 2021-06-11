package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, res float32
	var resInt int32
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, %):")
	fmt.Scanln(&op)

	if op == "%" {
		if b != 0 {
			resInt = int32(a) % int32(b)
			fmt.Printf("%d\n", resInt)
		} else {
			fmt.Println("Нельзя делить на 0")
			os.Exit(1)
		}
	} else {
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b != 0 {
				res = a / b
			} else {
				fmt.Println("Нельзя делить на 0")
				os.Exit(1)
			}
		default:
			fmt.Println("Операция выбрана неверно")
			os.Exit(1)
		}

		fmt.Printf("%f\n", res)
	}
}
