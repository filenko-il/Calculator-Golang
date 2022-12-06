/* Прочитав множетсво ресурсов и пройдя разные курсы, купив книжку, все что получилось написать у меня за две недели, с римскими цифрами.
 Первый файл на арабских цифрах был написан в первый же день, а вот трудности с римскими цифрами возникли на протяженности двух недель,
 с которыми так и не получилось до конца доработать, чтобы все работало гладко(
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var operator string
	var number1, number2 int
	var number3, number4 string
	fmt.Print("Введите первое число: ")
	fmt.Scan(&number3)
	number1 = romanToInt(number3)
	if number1 > 10 {
		fmt.Print("Превышено число максимума по заданию")
		return
	}

	fmt.Print("Введите второе число: ")
	fmt.Scan(&number4)
	number2 = romanToInt(number4)
	if number2 > 10 {
		fmt.Print("Превышено число максимума по заданию")
		return
	}
	{

		fmt.Print("\n Введите оператор (+ - * /): ")
		fmt.Scan(&operator)
		switch operator {
		case "+":
			fmt.Printf("%v %s %v = %v\n", (number3), operator, (number4), (number3 + number4))
		case "/":
			fmt.Printf("%v %s %v = %v\n", (number3), operator, (number4), integerToRoman(number1/number2))
		case "*":
			fmt.Printf("%v %s %v = %v\n", (number3), operator, (number4), integerToRoman(number1/number2))
		case "-":
			fmt.Printf("%v %s %v = %v\n", (number3), operator, (number4), integerToRoman(number1-number2))
		default:
			fmt.Println("Некорректный оператор")
		}
	}
}
func Roman(number int) string {
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

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

// Функция преобразования Римских чисел в арабские
func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
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
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
