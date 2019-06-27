package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)
type Triangle struct{
	name string
	side1,side2,side3 float64
}
func (tr *Triangle) area()float64{
	p:= (tr.side1 + tr.side2 + tr.side3) * 0.5;
	return math.Sqrt(p * (p - tr.side1) * (p - tr.side2) * (p - tr.side3))
}
//func addTriangle()[]Triangle{
//	fmt.Print("Enter a value.The correct input format is the name,\nthe length of the parties, separated by commas\n")
//	var lsOfTriangle [] Triangle
//	var triangle =Triangle{}
//	scanner:=bufio.NewReader(os.Stdin)
//	str,_:=scanner.ReadString('\n')
//	strings.TrimSuffix(str,"\n")
//	arr:=strings.Split(str,",")
//	//var err error
//	triangle.name=arr[0]
//	for i:=1;i<len(arr);i++{
//		_,err:=strconv.ParseFloat(strings.TrimSpace(arr[i]),64)
//		if err!=nil{
//			fmt.Print("Enter a correct value!")
//			addTriangle()
//			break
//		}}
//	triangle.side1=arr[1]
//
//	//triangle.side1,err=strconv.ParseFloat(strings.TrimSpace(arr[1]),64)
//	//if err!=nil{
//	//  fmt.Print("Enter a correct value!")
//	//  addTriangle()
//	//}
//	//triangle.side2,err=strconv.ParseFloat(strings.TrimSpace(arr[2]),64)
//	//if err!=nil{
//	//  fmt.Print("Enter a correct value!")
//	//  addTriangle()
//	//}
//	//triangle.side3,err=strconv.ParseFloat(strings.TrimSpace(arr[3]),64)
//	//if err!=nil{
//	//  fmt.Print("Enter a correct value!")
//	//  addTriangle()
//	//}
//	if checkASizeOfTr(triangle){
//		lsOfTriangle=append(lsOfTriangle,triangle)
//	}  else{addTriangle()}
//}
//func sortByAreaOfTriangle(){
//	sort.Slice(listOfTriangle[:],func(i,j int)bool{
//		return listOfTriangle[i].area()>listOfTriangle[j].area()
//	})
//}
//func displASortTriangles(){
//	fmt.Print("============= Triangles list: ===============\n")
//	for _,value:=range listOfTriangle{
//		fmt.Printf("[Triangle %s]: %f cm\n",value.name,value.area())
//	}
//}
//func checkASizeOfTr(tr Triangle)bool{
//	arr:=[]float64{tr.side1,tr.side2,tr.side3}
//	sort.Float64s(arr)
//	return arr[2]<(arr[0]+arr[1])
//}
func isCorrectValue(arr[]string){

}
func addTriangle()(triangle Triangle){
	fmt.Print("Enter a value.The correct input format is the name,\nthe length of the parties, separated by commas\n")
	scan:=bufio.NewReader(os.Stdin)
	str,_:=scan.ReadString('\n')
	strings.TrimSuffix(str,"\n")
	arr:=strings.Split(str,",")
	triangle.name=arr[0]
	for i:=1;i<len(arr);i++{
		_,err:=strconv.ParseFloat(strings.TrimSpace(arr[i]),64)
		if err!=nil{
			fmt.Print("Enter a correct value!")
			addTriangle()
			break
		}}
		triangle.side1=arr[1]
	return
}
func getAnAnswer()bool{
	var answer string
		fmt.Scanf("%v",&answer)
		r, _ := regexp.Compile(`(?i)[y][e]?[s]?`)
		return r.MatchString(answer)
}
func main(){
	var listOfTriangl []Triangle
	fmt.Print("Do you want to add a triangle? Answer yes or no:")
	for getAnAnswer(){
		listOfTriangl=append(listOfTriangl,addTriangle())
	}
}