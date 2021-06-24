package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Program1(number string) string {
	roma := ""
	num := strings.Split(number, "")
	length := len(num)
	for i := 0; i < length; i++ {
		roma += numtoroma(num[i], length-i)
	}
	return roma
}

func numtoroma(num string, i int) string {
	var roma string
	if i == 1 {
		switch num {
		case "1":
			roma = "I"
		case "2":
			roma = "II"
		case "3":
			roma = "III"
		case "4":
			roma = "IV"
		case "5":
			roma = "V"
		case "6":
			roma = "VI"
		case "7":
			roma = "VII"
		case "8":
			roma = "VIII"
		case "9":
			roma = "IX"

		}
	}
	return roma
}

func main() {
	var number int
	fmt.Printf("Please input number: ")
	fmt.Scan(&number)
	num := strconv.Itoa(number)
	roma := Program1(num)
	fmt.Printf("%s\n", roma)
}
