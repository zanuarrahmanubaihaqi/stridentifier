package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	CountUniqueChars(str)
}

func CountUniqueChars(Str string) {
	arr := strings.Split(Str, "")

	var nSlice []map[string]int
	nArr := map[string]int{}
	var x string

	for index, value := range arr {
		if x != value {
			if x != "" {
				nSlice = append(nSlice, nArr)
				nArr = map[string]int{}
			}
			x = value
		}
		nArr[value]++
		if (len(arr) - 1) == index {
			nSlice = append(nSlice, nArr)
		}
	}
	fmt.Println(nSlice)
}

// Buat function untuk hitung jumlah masing-masing karakter unik yang berurutan dari sebuah string.Fn Name : CountUniqueChar(input)Input  : StringOutput  : []map[string]int
// Contoh : Input  : “aaaasssiia”output  : [[“a”: 4], [“s”: 3], [“i”: 2], [“a”: 1]]

func OtherFunction(str string) {
	kata := "aaabbbcccdae"
	splitKata := strings.Split(kata, "")
	sort.Strings(splitKata)
	var huruf []string
	var count []uint8

	for i := 0; i < len(splitKata); i++ {
		if !slices.Contains(huruf, splitKata[i]) {
			huruf = append(huruf, splitKata[i])
			count = append(count, 1)
		} else {
			count[len(count)-1] += 1
		}
	}

	for j := 0; j < len(huruf); j++ {
		fmt.Println(huruf[j], "=", count[j])
	}
}
// Output :
// a = 4
// b = 3
// c = 3
// d = 1
// e = 1
