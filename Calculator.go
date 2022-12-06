package main

import (
   "fmt"
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
      fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1+number2)
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

