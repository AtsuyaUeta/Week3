package main

import "testing"

func Program101(t *testing.T) {
	str := Program1("1")
	if str != "I" {
		t.Error("Test01 is failed")
	}
}
