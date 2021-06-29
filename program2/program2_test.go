package main

import "testing"

func TestProgram201(t *testing.T) {
	num := Program2("MCMXC")
	if num != 1990 {
		t.Error("Test01 is failed")
	}
}

func TestProgram202(t *testing.T) {
	num := Program2("MMVIII")
	if num != 2008 {
		t.Error("Test02 is failed")
	}
}

func TestProgram203(t *testing.T) {
	num := Program2("XCIX")
	if num != 99 {
		t.Error("Test03 is failed")
	}
}

func TestProgram204(t *testing.T) {
	num := Program2("XLVII")
	if num != 47 {
		t.Error("Test04 is failed")
	}
}
