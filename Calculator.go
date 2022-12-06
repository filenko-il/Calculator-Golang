package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var operator string
	fmt.Print("Введите первое число: ")
	var number1 = 1
	if number1 < 11 {
		fmt.Scanln(&number1)
	}
	if number1 > 10 {
		fmt.Print("Превышено число максимума по заданию")
		return
	}
	fmt.Print("Введите второе число: ")
	var number2 = 1
	if number2 < 11 {
		fmt.Scanln(&number2)
	}
	if number2 > 10 {
		fmt.Print("Превышено число максимума по заданию")
		return
	}

	fmt.Print("\n Введите оператор (+ - * /): ")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		fmt.Printf(" %v %s %v = %v\n", number1, operator, number2, number1+number2)
	case "/":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1/number2)
	case "*":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1*number2)
	case "-":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1-number2)
	default:
		fmt.Println("Некорректный оператор")
	}
}
func integerToRoman(number3, number4 int) string {
	var operator string
	maxRomanNumber := 10
	if number3 > maxRomanNumber {
		return strconv.Itoa(number3)
	}
	if number4 > maxRomanNumber {
		return strconv.Itoa(number4)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number3 >= conversion.value {
			roman.WriteString(conversion.digit)
			number3 -= conversion.value
		}
		for number4 >= conversion.value {
			roman.WriteString(conversion.digit)
			number4 -= conversion.value
		}

	}
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		fmt.Printf("%v %s %v = %v\n", number3, operator, number4, number3+number4)
	case "/":
		fmt.Printf("%v %s %v = %v\n", number3, operator, number4, number3/number4)
	case "*":
		fmt.Printf("%v %s %v = %v\n", number3, operator, number4, number3*number4)
	case "-":
		fmt.Printf("%v %s %v = %v\n", number3, operator, number4, number3-number4)
	default:
		fmt.Println("Некорректный оператор")
	}

	return roman.String()
}
