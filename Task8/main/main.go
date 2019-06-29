package main

import (
	"fmt"
	"strconv"
)

func initValue() int {
	fmt.Print("It must be a positive number:\n")
	var (
		str   string
		value int
		err   error
	)
	for {
		fmt.Scan(&str)
		value, err = strconv.Atoi(str)
		if err != nil || value < 0 {
			fmt.Print("Please, write the correct value\n")
		} else {
			break
		}
	}
	return value
}
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func fib(b int) []int {
	var fibonacci []int
	for i := 0; ; i++ {
		fn := Fibonacci(i)
		if fn > b {
			break
		}
		fibonacci = append(fibonacci, fn)
	}
	return fibonacci
}
func sortFib(begin, end int, fibonacci []int) (numbArr []int) {
	for _, value := range fibonacci {
		if value >= begin && value <= end {
			numbArr = append(numbArr, value)
		}
	}
	return
}
func writeArr(arrOfNumbers []int) {
	fmt.Println("Fibonacci series: ", arrOfNumbers)
}
func isEndBiggerThanBegin(a, b int) bool {
	return b > a
}
func getAnswer() {
	fmt.Println("Enter the value of the beginning of the series")
	begin := initValue()
	fmt.Println("Enter the value of the ending of the series")
	end := initValue()
	if isEndBiggerThanBegin(begin, end) {
		fibArr := fib(end)
		numbArr := sortFib(begin, end, fibArr)
		writeArr(numbArr)
	} else {
		fmt.Print("The first value must be less than the second\n")
		getAnswer()
	}
}
func main() {
	getAnswer()
}
