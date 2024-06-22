# Go-ASCII-String-Package
Go package to convert ASCII numeric code to string character and string character to ASCII numeric code 

## Installation Build and Run

```shell
git clone https://github.com/Com1Software/Test-GoWebServer.git gowebserver
cd gowebserver
go mod init testserver
go mod tidy
go build
testserver


```


## Example Code

```shell
package main

import (
	"fmt"
	"runtime"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

func main() {

	fmt.Printf("(c) Copyright Com1Software 1992-2024\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
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

```


## Useful Links



ASCII

[ American Standard Code for Information Interchange](https://en.wikipedia.org/wiki/ASCII)

How do I convert a sting character into ASCII in Golang? 

[ASCII doesn't exactly exist in Go.](https://groups.google.com/g/golang-nuts/c/Hvm9Nq3dF2M)

[Conversions Spec ](https://go.dev/ref/spec#Conversions)

