package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var iStr string
	var iStrSlice []string
	var iFloatSlice [2]float64

	iStdin := bufio.NewScanner(os.Stdin)
	if iStdin.Scan() {
		iStr = iStdin.Text()
	}

	iStrSlice = strings.Split(iStr, ";")

	iStrSlice[0] = strings.Replace(iStrSlice[0], " ", "", -1)
	iStrSlice[1] = strings.Replace(iStrSlice[1], " ", "", -1)
	iStrSlice[0] = strings.Replace(iStrSlice[0], ",", ".", -1)
	iStrSlice[1] = strings.Replace(iStrSlice[1], ",", ".", -1)
	fmt.Println(iStrSlice)

	iFloatSlice[0], _ = strconv.ParseFloat(iStrSlice[0], 64)
	iFloatSlice[1], _ = strconv.ParseFloat(iStrSlice[1], 64)

	fmt.Printf("%.4f", iFloatSlice[0] / iFloatSlice[1])
}