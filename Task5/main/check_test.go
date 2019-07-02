package main

import (
	"testing"
)

func TestInitVal(t *testing.T) {
	var str string
	str = initVal(125, true, "", "", "")
	if str != "сто двадцать пять " {
		t.Error("Expected сто двадцать пять ,got ", str)
	}
}
func TestCorrectEnd(t *testing.T) {
	var str string
	str = correctEnd(1, "тысяча ", "тысячи ", "тысяч ")
	if str != "тысяча " {
		t.Error("Expected тысяча got: ", str)
	}
}
func TestAppend(t *testing.T) {
	var str string
	str = append("one", "two ")
	if str != "two one" {
		t.Error("Expected two one, got ", str)
	}
}
func TestStringValue(t *testing.T) {
	var str string
	str = stringValue(1245)
	if str != "одна тысяча двести сорок пять " {
		t.Error("Expected одна тысяча двести сорок пять ,got ", str)
	}
}
