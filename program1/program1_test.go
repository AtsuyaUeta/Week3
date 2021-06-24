package main

import "testing"

func TestProgram101(t *testing.T) {
	str := Program1("1")
	if str != "I" {
		t.Error("Test01 is failed")
	}
}

func TestProgram102(t *testing.T) {
	str := Program1("99")
	if str != "XCIX" {
		t.Error("Test02 is failed")
	}
}

func TestProgram103(t *testing.T) {
	str := Program1("101")
	if str != "CI" {
		t.Error("Test03 is failed")
	}
}

func TestProgram104(t *testing.T) {
	str := Program1("2008")
	if str != "MMVIII" {
		t.Error("Test04 is failed")
	}
}
