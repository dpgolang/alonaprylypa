package main

import (
	"fmt"
	"strconv"
)

//const maxNumber=1000000000000
var hunds = []string{
	"",
	"сто ",
	"двести ",
	"триста ",
	"четыреста ",
	"пятьсот ", "шестьсот ",
	"семьсот ",
	"восемьсот ",
	"девятьсот ",
}
var tens = []string{
	"",
	"десять ",
	"двадцать ",
	"тридцать ",
	"сорок ",
	"пятьдесят ",
	"шестьдесят ",
	"семьдесят ",
	"восемьдесят ",
	"девяносто ",
}
var part20 = []string{
	"",
	"один ",
	"два ",
	"три ",
	"четыре ",
	"пять ",
	"шесть ",
	"семь ",
	"восемь ",
	"девять ",
	"десять ",
	"одиннадцать ",
	"двенадцать ",
	"тринадцать ",
	"четырнадцать ",
	"пятнадцать ",
	"шестнадцать ",
	"семнадцать ",
	"восемьнадцать ",
	"девятнадцать ",
}
func initVal(val int, male bool, one, two, five string) (str string) {
	var num = val % 1000
	if num <= 0 {
		return ""
	}
	if !male {
		part20[1] = "одна "
		part20[2] = "две "
	}else{
		part20[1]="один "
		part20[2]="два "
	}
	str = hunds[num/100]

	if num%100 < 20 {
		str += part20[num%100]
	} else {
		str += tens[num%100/10]
		str += part20[num%10]
	}
	str += correctEnd(num, one, two, five)
	return str
}
//Choosing the right case ending noun
func correctEnd(val int, one, two, five string) string {
	t := val % 100
	if t > 20 {
		t = val % 10
	} else {
		t = val % 20
	}
	switch t {
	case 1:
		return one
	case 2, 3, 4:
		return two
	default:
		return five
	}
}
func append(s, s1 string) (result string) {
	return s1 + s
}
func stringValue(value int) (str string) {
	minus := false
	if value < 0 {
		value = -value
		minus = true
	}
	n := value
	if n == 0 {
		str += "ноль"
	}
	if n%1000 != 0 {
		str += initVal(n, true, "", "", "")
	}
	n /= 1000
	newStr := append(str, initVal(n, false, "тысяча ", "тысячи ", "тысяч "))
	n /= 1000
	newStr = append(newStr, initVal(n, true, "миллион ", "миллиона ", "миллионов "))
	n /= 1000
	newStr = append(newStr, initVal(n, true, "миллиард ", "миллиарда ", "миллиардов "))
	n /= 1000
	newStr = append(newStr, initVal(n, true, "триллион ", "триллиона ", "триллионов "))
	n /= 1000
	newStr = append(newStr, initVal(n, true, "триллиард ", "триллиарда ", "триллиардов "))
	if minus {
		newStr = append(newStr, "минус ")
	}
	return newStr
}
func getAnswer() int {
	var (
		n   int
		str string
		err error
	)
	fmt.Println("Please enter the value of n")
	for {
		fmt.Scanf("%s", &str)
		n, err = strconv.Atoi(str)
		if err != nil {
			err = fmt.Errorf("Please, write the correct value, %s\n", err)
			fmt.Print(err)
		} else {
			break
		}
	}
	return n
}
func main() {
	n := getAnswer()
	fmt.Print(stringValue(n))
}
