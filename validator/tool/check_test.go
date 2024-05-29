package tool

import (
	"testing"
)

func TestCheckParam(t *testing.T) {
	var person = Person{
		Name: "1252074183@qq.com",
		Age:  1,
	}
	err := CheckParams(person)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("check success")
	}
}
