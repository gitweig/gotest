package gotest

import (
	"testing"
)

func Test_Add(t *testing.T) {
	if ret := Add(2, 4); ret != 6 {
		t.Error("Add function error!")
	} else {
		t.Log("Add function success!")
	}
}

func Test_Multity(t *testing.T) {
	if ret := Multity(2, 4); ret != 8 {
		t.Error("Multity function error!")
	} else {
		t.Log("Multity function success!")
	}
}
