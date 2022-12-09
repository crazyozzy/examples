package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("task.txt")
	if err != nil {
		panic("Cannot open input file!")
	}

	buf := bufio.NewReader(file)

	for i := 1; ; i++ {
		num, err := buf.ReadString(';')
		if err != nil && err != io.EOF {
			panic("Cannot read new number!")
		} else if err == io.EOF {
			break
		}

		if num == "0;" {
			fmt.Println("Index of 0 is: ", i)
			break
		}
	}
}