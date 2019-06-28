package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const half = 0.5
const sides = 3

type triangle struct {
	name                string
	side1, side2, side3 float64
}

func (tr *triangle) area() float64 {
	p := (tr.side1 + tr.side2 + tr.side3) * half
	return math.Sqrt(p * (p - tr.side1) * (p - tr.side2) * (p - tr.side3))
}
func (tr *triangle) check() bool {
	arr := []float64{tr.side1, tr.side2, tr.side3}
	sort.Float64s(arr)
	return arr[2] < (arr[0] + arr[1])
}
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
func valid(str []string) (values []float64) {
	for _, val := range str {
		n, err := strconv.ParseFloat(SpaceMap(val), 64)
		if err == nil && n > 0 {
			values = append(values, n)
		}
	}
	return
}
func add() (triangle, error) {
	var tr triangle
	fmt.Println("Enter a value. The correct input format is the name,\nthe length of the parties, separated by commas")
	scanner := bufio.NewReader(os.Stdin)
	str, _ := scanner.ReadString('\n')
	strings.TrimSuffix(str, "\n")
	arr := strings.Split(str, ",")
	if len(arr) != 4 {
		return tr, errors.New("don't enough arguments\n")
	}
	tr.name = arr[0]
	values := valid(arr[1:4])
	if len(values) != sides {
		return tr, errors.New("incorrect data\n")
	}
	tr.side1 = values[0]
	tr.side2 = values[1]
	tr.side3 = values[2]
	if !tr.check() {
		return tr, errors.New("the sum of the lengths of any two sides of a triangle must be greater than or equal to the length of the third side\n")
	}
	return tr, nil
}
func create() triangle {
	var tr triangle
	var err error
	for {
		tr, err = add()
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	return tr
}
func getAnswer() (list []triangle) {
	var answer string
	for {
		fmt.Println("Do you want to add a triangle? Answer yes or no:")
		_, err := fmt.Scanf("%v", &answer)
		if err != nil {
			break
		}
		r, err := regexp.Compile(`(?i)[y][e]?[s]?`)
		if err != nil {
			fmt.Print(err)
		}
		if r.MatchString(answer) {
			list = append(list, create())
		} else {
			break
		}
	}
	return list
}
func sortTriangle(list []triangle) (sortedList []triangle) {
	sort.Slice(list[:], func(i, j int) bool {
		return list[i].area() > list[j].area()
	})
	return list
}
func dispTriangleList(list []triangle) {
	fmt.Println("============= Triangles list: ===============")
	for _, value := range list {
		fmt.Printf("[Triangle %s]: %f cm\n", value.name, value.area())
	}
}
func isEmpty(list []triangle) bool {
	if len(list) == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	list := getAnswer()
	for isEmpty(list) {
		fmt.Println("Your list of triangles is empty")
		list = getAnswer()
	}
	dispTriangleList(sortTriangle(list))
}
