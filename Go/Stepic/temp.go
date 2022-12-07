package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(iN uint) uint {
		iStr := strconv.Itoa(int(iN))
		oStr := ""

		for _, v := range iStr {
			switch string(v) {
			case "2", "4", "6", "8":
				oStr += string(v)
			}
		}
		
		temp, _ := strconv.ParseUint(oStr, 10, 64)
		iN = uint(temp)

		if iN == 0 {
			return 100
		} else {
			return iN
		}
	}

	fmt.Print(fn(727178))
}