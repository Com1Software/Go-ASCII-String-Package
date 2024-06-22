package main

import (
	"fmt"
	"runtime"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

func main() {

	fmt.Printf("(c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	fmt.Println(string(34))
	for x := 1; x < 128; x++ {
		fmt.Printf(asciistring.ASCIIToString(x))
	}
	fmt.Println("")
	chr := "A"
	fmt.Println(asciistring.StringToASCII(chr))
	chr = "B"
	fmt.Println(asciistring.StringToASCII(chr))
	chr = "C"
	fmt.Println(asciistring.StringToASCII(chr))
}
