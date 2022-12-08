package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// i-процентная ставка по займу в месяц
	// S-сумма кредита
	// n -количество месяцев, в которые платится ипотека
	// m - количество месяцев в году
	//var i float64
	//var S float64 = 1300000
	//var n float64 = 168
	//var n float64 = 84
	var m float64 = 12
	//Возвращает новое значение bufio.Scanner
	//Scanner читает данные из стандартного ввода (с клавиатуры)
	scanner := bufio.NewScanner(os.Stdin)

	//Запрашиваем данные у пользователя
	fmt.Print("Введите сумму кредита:")
	scanner.Scan()
	Sum := scanner.Text()
	S, err := strconv.ParseFloat(Sum, 64)
	if err != nil {
		fmt.Println("Пожалуйста введите число с плавающей точкой")

	}

	// Запрашиваем данные у пользователя
	fmt.Print("Количество месяцев, в которые платится ипотека: Пр.84\n")
	scanner.Scan()
	number := scanner.Text()
	n, err := strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Println("Пожалуйста введите число с плавающей точкой")

	}

	//Запрашиваем данные у пользователя
	fmt.Print("Введите процентную ставку по займу в месяц Пр:(12,9 % процент по кредиту : 100 ) т.е.=0.129\n")
	scanner.Scan()
	input := scanner.Text()
	in, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Пожалуйста введите число с плавающей точкой")

	}

	i := in / m // (12,9 % процент по кредиту : 100 ) / m = 12 месяцев
	test := math.Pow((1 + i), n)

	resalt := (S * (i * math.Pow((1+i), n))) / (math.Pow((1+i), n) - 1)

	fmt.Printf("Сумма ежемесячного платежа составляет: %0.2f, %0.2f \n", resalt, test)

}
