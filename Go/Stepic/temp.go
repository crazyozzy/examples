package main

import "strings"

func removeDuplicates(inputStream, outputStream chan string) {
	dupl := make(map[string]bool)

	for v := range inputStream {
		if _, ok := dupl[v]; !ok {
			outputStream <- v
			dupl[v] = true
		}
	}

	close(outputStream)
}