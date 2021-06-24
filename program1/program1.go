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
		default:
			roma = ""
		}
	} else if i == 2 {
		switch num {
		case "1":
			roma = "X"
		case "2":
			roma = "XX"
		case "3":
			roma = "XXX"
		case "4":
			roma = "XL"
		case "5":
			roma = "L"
		case "6":
			roma = "LX"
		case "7":
			roma = "LXX"
		case "8":
			roma = "LXXX"
		case "9":
			roma = "XC"
		default:
			roma = ""
		}
	} else if i == 3 {
		switch num {
		case "1":
			roma = "C"
		case "2":
			roma = "CC"
		case "3":
			roma = "CCC"
		case "4":
			roma = "CD"
		case "5":
			roma = "D"
		case "6":
			roma = "DC"
		case "7":
			roma = "DCC"
		case "8":
			roma = "DCCC"
		case "9":
			roma = "CM"
		default:
			roma = ""
		}
	} else if i == 4 {
		switch num {
		case "1":
			roma = "M"
		case "2":
			roma = "MM"
		case "3":
			roma = "MMM"
		case "4":
			roma = "MMMM"
		default:
			roma = ""
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
