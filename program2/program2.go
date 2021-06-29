package main

import (
	"fmt"
	//	"strconv"
	"strings"
)

func Program2(roma string) int {
	num := 0                       //数字を0に初期化
	i := 0                         //配列の何番目を数えているか
	ele := strings.Split(roma, "") //要素を1文字ずつに分割
	length := len(ele)             //要素の数、文字列の長さ

	for i < length-1 {

		if ele[i] == "M" {
			num = num + 1000
			i = i + 1
		} else if ele[i] == "C" && ele[i+1] == "M" {
			num = num + 900
			i = i + 2
		} else if ele[i] == "D" {
			num = num + 500
			i = i + 1
		} else if ele[i] == "C" && ele[i+1] == "D" {
			num = num + 400
			i = i + 2
		} else if ele[i] == "C" {
			num = num + 100
			i = i + 1
		} else if ele[i] == "X" && ele[i+1] == "C" {
			num = num + 90
			i = i + 2
		} else if ele[i] == "L" {
			num = num + 50
			i = i + 1
		} else if ele[i] == "X" && ele[i+1] == "L" {
			num = num + 40
			i = i + 2
		} else if ele[i] == "X" {
			num = num + 10
			i = i + 1
		} else if ele[i] == "I" && ele[i+1] == "X" {
			num = num + 9
			i = i + 2
		} else if ele[i] == "V" {
			num = num + 5
			i = i + 1
		} else if ele[i] == "I" && ele[i+1] == "V" {
			num = num + 4
			i = i + 2
		} else if ele[i] == "I" {
			num = num + 1
			i = i + 1
		}
	}

	if i == length-1 {
		if ele[i] == "M" {
			num = num + 1000
		} else if ele[i] == "D" {
			num = num + 500
		} else if ele[i] == "C" {
			num = num + 100
		} else if ele[i] == "L" {
			num = num + 50
		} else if ele[i] == "X" {
			num = num + 10
		} else if ele[i] == "V" {
			num = num + 5
		} else if ele[i] == "I" {
			num = num + 1
		}
	}
	return num
}

func main() {
	//	roma := "MCMXC"
	//	roma := "MMVIII"
	//	roma := "XCIX"
	roma := "XLVII"
	num := Program2(roma)
	fmt.Println(num)
}
