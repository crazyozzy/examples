package main

import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	"os"            // пакет используется для проверки ответа, не удаляйте его
)



func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
                                            // все полученные значения имеют тип пустого интерфейса
	var iFloat1, iFloat2 float64

	switch v := value1.(type) {
	case float64:
		iFloat1 = v
	default:
		fmt.Printf("value=%v: %T", v, v)
		os.Exit(1)
	}

	switch v := value2.(type) {
	case float64:
		iFloat2 = v
	default:
		fmt.Printf("value=%v: %T", v, v)
		os.Exit(1)
	}

	switch operation {
	case "+":
		fmt.Printf("%.4f", iFloat1 + iFloat2)
	case "-":
		fmt.Printf("%.4f", iFloat1 - iFloat2)
	case "*":
		fmt.Printf("%.4f", iFloat1 * iFloat2)
	case "/":
		fmt.Printf("%.4f", iFloat1 / iFloat2)
	default:
		fmt.Print("неизвестная операция")
	}
}