package main

import (
	"fmt"
	"strconv"
	"strings"
)

const maxNumber = 10000000000

type Number struct {
	Num int
}

func (numb *Number) Squar() int {
	return numb.Num * numb.Num
}
func GetArray(n int) (arr []string) {
	var numb Number
	var s string
	for i := 0; ; i++ {
		numb.Num = i
		if numb.Squar() < n {
			s = strconv.Itoa(numb.Num)
			arr = append(arr, s)
		} else {
			break
		}
	}
	return
}
func printArr(arr []string) {
	fmt.Print(strings.Join(arr, ","))
}
func start() {
	var (
		n   int
		str string
		err error
	)
	fmt.Println("The program displays a series of natural numbers separated by commas, the square of which is less than the given n")
	fmt.Println("Please enter the value of n")
	for {
		fmt.Scan(&str)
		n, err = strconv.Atoi(str)
		if err != nil || n < 0 || n > maxNumber {
			fmt.Print("Please, write the correct value\n")
		} else {
			break
		}
	}
	printArr(GetArray(n))
}
func main() {
	start()
}
