package main
import(
	"reflect"
	"testing"
)
func TestGetArray(t *testing.T) {
	var numbers []string
	numbers=GetArray(10)
	var arrStr=[]string{"1","2","3"}
	if reflect.DeepEqual(numbers,arrStr){
		t.Error("Expected arrStr, got ",numbers)
	}
}
func TestNumberSquar(t *testing.T) {
	var value int
	var num Number
	num.Num = 4
	value = num.Squar()
	if value!=16{
		t.Error("Expected 16, got ",value)
	}
}
