package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	fmt.Printf("(c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	fmt.Println(string(34))
	for x := 1; x < 128; x++ {
		fmt.Printf(ASCIIToString(x))
	}
	fmt.Println("")
	chr := "A"
	fmt.Println(StringToASCII(chr))
	chr = "B"
	fmt.Println(StringToASCII(chr))
	chr = "C"
	fmt.Println(StringToASCII(chr))
}

func StringToASCII(val string) int {
	chr := ""
	rtn := 0
	for x := 1; x < 128; x++ {
		chr = ASCIIToString(x)
		if strings.Contains(chr, val) {
			rtn = x
			x = 128
		}
	}
	return rtn
}

func ASCIIToString(val int) string {
	rtn := string(val)
	return rtn
}
