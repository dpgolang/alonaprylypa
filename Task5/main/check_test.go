package main

import (
	"testing"
)

func TestInitVal(t *testing.T) {
	var str string
	str = InitVal(125, true, "", "", "")
	if str != "сто двадцать пять " {
		t.Error("Expected сто двадцать пять ,got ", str)
	}
}
func TestCorrectEnd(t *testing.T) {
	var str string
	str = CorrectEnd(1, "тысяча ", "тысячи ", "тысяч ")
	if str != "тысяча " {
		t.Error("Expected тысяча got: ", str)
	}
}
func TestReverse(t *testing.T) {
	var str string
	str = Reverse("one", "two ")
	if str != "two one" {
		t.Error("Expected two one, got ", str)
	}
}
func TestStringValue(t *testing.T) {
	var str string
	str = StringValue(1245)
	if str != "одна тысяча двести сорок пять " {
		t.Error("Expected одна тысяча двести сорок пять ,got ", str)
	}
}
