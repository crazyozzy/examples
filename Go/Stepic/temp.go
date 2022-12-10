package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	const timeFormat = "2006-01-02 15:04:05"

	iByte, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	iStr := strings.TrimRight(string(iByte), "\r\n")
	iStr = strings.TrimRight(iStr, "\n")

	//2020-05-15 08:00:00
	iTime, err := time.Parse(timeFormat, iStr)
	if err != nil && err != io.EOF {
		panic(err)
	}

	if iTime.Hour() >= 13 {
		fmt.Println(iTime.AddDate(0, 0, 1).Format(timeFormat))
	} else {
		fmt.Println(iTime.Format(timeFormat))
	}
}