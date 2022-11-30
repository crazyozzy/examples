package main

import (
	"fmt"
	"strings"
)

func main() {
	var iStrText string
	
	fmt.Scan(&iStrText)

	for _, v := range iStrText {
		fmt.Println(iStrText, string(v))
		if strings.Count(iStrText, string(v)) > 1 {
			iStrText = strings.Trim(iStrText, string(v))
			fmt.Println(iStrText, string(v))
		}
	}

	fmt.Println(iStrText)
}