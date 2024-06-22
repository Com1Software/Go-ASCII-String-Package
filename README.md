# Go-ASCII-String-Package
Go package to convert ASCII numeric code to string character and string characterto ASCII numeric code 


## Example Code

```shell
package main

import (
	"fmt"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

func main() {
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

