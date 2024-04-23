package main

import (
	"fmt"
	"strings"
)

type Responses struct {
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

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
