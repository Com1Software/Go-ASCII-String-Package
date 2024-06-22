package asciistring

import (
	"strings"
)
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
