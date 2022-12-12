package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	const timeFormat = "02.01.2006 15:04:05"
	const startTime = 1589570165
	var min, sec int

	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	iStr := fmt.Sprintf("0h%dm%ds", min, sec)

	iDur, err := time.ParseDuration(iStr)
	if err != nil && err != io.EOF {
		panic(err)
	}

	iTime := time.Unix(startTime, 0).UTC()
	iTime = iTime.Add(iDur)

	fmt.Println(iTime.Format(time.UnixDate))
}