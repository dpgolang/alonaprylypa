package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)
var file,separator,algorithm string
const amountTicketNumb=6
func init(){
	flag.StringVar(&file,"file","/home/alonaprylypa/go/ElementaryTasks/Task6/main/text.txt","address of file")
	flag.StringVar(&separator,"sep",",","text separator")
	flag.StringVar(&algorithm,"alg","Moskow","choose the algorithm, Moskow or Piter")
	flag.Parse()
}
func readFile()string{
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
func isTicket(ticket string)(int,error){
	num,err:=strconv.Atoi(ticket)
	return num,err
}
func createArr(str string)[]int{
	var (
		tickArr[]int
		err error
		ticket int
	)
	strArr:=strings.Split(str,separator)
	for _,value:=range strArr{
		ticket,err=isTicket(value)
		if err==nil&&len(value)==amountTicketNumb{
			tickArr=append(tickArr,ticket)
		}
	}
	return tickArr
}
func moskowTick(tickets[]int)int{
	var count int
	for _,tick:=range tickets{
		if tick/100000+tick%100000/10000+tick%10000/1000==tick%1000/100+tick%100/10+tick%10{
			count++
		}
	}
	return count
}
func piterTick(tickets[]int)int{
	var count,sum1,sum2 int
	var numbArr[]int
	for _,tick:=range tickets{
		sum1=0
		sum2=0
		numbArr=[]int{tick/100000,tick%100000/10000,tick%10000/1000,tick%1000/100,tick%100/10,tick%10}
		for i:=0;i<len(numbArr);i++{
			if numbArr[i]%2==0{
				sum1+=numbArr[i]
			} else{
				sum2+=numbArr[i]
			}
		}
		if sum1==sum2{
			count++
		}
	}
	return count
}
func countLuckyTicket(tickets[]int)int{
if algorithm=="Piter"{
	return piterTick(tickets)
}
return moskowTick(tickets)
}
func main() {
	content:=readFile()
	fmt.Print(content)
	tickArr:=createArr(content)
	fmt.Print("The number of lucky ticket: ",countLuckyTicket(tickArr))
}
