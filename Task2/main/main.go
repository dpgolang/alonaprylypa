package main

import (
	"fmt"
	"regexp"
	"strconv"
)
type envelope struct{
	length,width float64
}
func isValidVal(value *float64)bool{
	var str string
	fmt.Scanf("%s", &str)
	n, err:= strconv.ParseFloat(str,64)
	if err!=nil||n<=0{
		fmt.Println("Please, enter a correct value. It must be an unsigned number")
		return false
	}else{
		*value=n
		return true
	}
}
func (env *envelope) checkParams(){
	if env.width>env.length{
		env.width, env.length = env.length, env.width
	}
}
func initEnvelopsSize()(firstEnv, secondEnv envelope){
	fmt.Print("Please, enter a length of the first envelope:")
	for {
		if isValidVal(&firstEnv.length){break}
	}
	fmt.Print("Please, enter a width of the first envelope:")
	for {
		if isValidVal(&firstEnv.width){break}
	}
	fmt.Print("Please, enter a length of the second envelope:")
	for {
		if isValidVal(&secondEnv.length){break}
	}
	fmt.Print("Please, enter a width of the second envelope:")
	for {
		if isValidVal(&secondEnv.width){break}
	}
	return
}
func returnProgram(){
	var answer string
	fmt.Println("Are you want to continue?")
	fmt.Scanf("%v",&answer)
	r, err := regexp.Compile(`(?i)[y][e]?[s]?`)
	if err!=nil{
		fmt.Println("check correctness of regexp")
	}
	for r.MatchString(answer){
		programStart()
		fmt.Println("Thank you! Goodbye!")
		break
	}

}
func putOneIntoAnother(env1, env2 envelope){
	p:=env1.length
	a:=env2.length
	q:=env1.width
	b:=env2.width
	if (p <= a && q <= b) || ((p > a &&q<=b)&&(((a+b)/(p+q))*((a+b)/(p+q))+((a-b)/(p-q))*((a-b)/(p-q)))>=2){
		fmt.Println("We can put the first envelope in the second")
	}else if (a <= p && b <= q) || ((a > p && b <= q )&&(((p+q)/(a+b))*((p+q)/(a+b))+((p-q)/(a-b))*((p-q)/(a-b)))>=2){
		fmt.Println("We can put the second envelope in the first")
	}else{
		fmt.Println("We can't put one envelope in another")
	}
}
func programStart(){
	firstEnv,secondEnv:=initEnvelopsSize()
	firstEnv.checkParams()
	secondEnv.checkParams()
	putOneIntoAnother(firstEnv,secondEnv)
	returnProgram()
}
func main(){
	programStart()
}
