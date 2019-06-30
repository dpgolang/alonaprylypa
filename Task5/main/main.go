package main

import (
	"fmt"
	"strconv"
)
const maxNumber=1000000000000
var hunds = [10]string{"","сто ","двести ","триста ","четыреста ","пятьсот ","шестьсот ","семьсот ","восемьсот ","девятьсот "}
var tens = [10]string{"","десять ","двадцать ","тридцать ","сорок ","пятьдесят ","шестьдесят ","семьдесят ","восемьдесят ","девяносто "}
func initVal(val int, male bool, one, two, five string)(str string){
	frac20:=[20]string{"","один ","два ","три ","четыре ","пять ","шесть ","семь ","восемь ","девять ","десять ","одиннадцать ","двенадцать ","тринадцать ","четырнадцать ","пятнадцать ","шестнадцать ","семнадцать ","восемьнадцать ","девятнадцать "}
	var num=val%1000
	if num==0{return ""}
	if num<0{return ""}
	if(!male){
		frac20[1]="одна "
		frac20[2]="две "
	}
	str=hunds[num/100]

	if num%100<20{
		str+=frac20[num%100]
	}else{
		str+=tens[num%100/10]
		str+=frac20[num%10]
	}
	str+=correctEnd(num,one,two,five)
	return str
}
func correctEnd(val int,one,two,five string)string{
	t:=val%100
	if t>20{
		t = val%10
	}else{
		t = val%20
	}
	switch t{
	case 1: return one
	case 2,3,4: return two
	default: return five
	}
	return ""
}
func reverse(s,s1 string) (result string) {
	return s1+s
}
func stringValue(value int)(str string){
	minus:=false
	if value<0{
		value = -value
		minus = true
	}
	n:=value
	if n==0{str+="ноль"}
	if n%1000!=0{str+=initVal(n,true,"","","")}
	n/=1000
	newStr:=reverse(str,initVal(n,false,"тысяча ","тысячи ","тысяч "))
	n/=1000
	newStr=reverse(newStr,initVal(n,true,"миллион ","миллиона ","миллионов "))
	n/=1000
	newStr=reverse(newStr,initVal(n,true,"миллиард ","миллиарда ","миллиардов "))
	if minus {
		newStr=reverse(newStr,"минус ")
	}
	return newStr
}
func getAnswer()int{
	var (
		n   int
		str string
		err error
	)
	fmt.Println("Please enter the value of n")
	for {
		fmt.Scan(&str)
		n, err = strconv.Atoi(str)
		if err != nil || n > maxNumber {
			fmt.Print("Please, write the correct value\n")
		} else {
			break
		}
	}
	return n
}
func main(){
	n:=getAnswer()
	fmt.Print(stringValue(n))
}
