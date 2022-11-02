package main

import (
	"fmt"
	"log"
)

func add(a, b int) (int, error) {
	return a + b, nil
}

func main() {
	res, err := add(2, 3)

	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	fmt.Printf("sum: %d\n", res)

	sum := 0

	for i := 0; i < 10; i++ {
		//sum += i
		sum = sum + i
	}

	fmt.Printf("sum i: %d\n", sum)

	var balance [10]int64
	balance = [10]int64{1, 112, 34, 3452, 453, 4555, 666, 7, 888, 9653}
	
	fmt.Printf("balance: %v", balance)
}
