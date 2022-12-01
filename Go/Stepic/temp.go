package main

import (
	"fmt"
	"unicode"
)

func main() {
	var iStrText string
	var oWrongPass bool
	
	fmt.Scan(&iStrText)

	for _, v := range iStrText {
		if !unicode.Is(unicode.Latin, v) && !unicode.IsDigit(v) {
			oWrongPass = true
			break
		}
	}

	if !oWrongPass && len(iStrText) >= 5 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}