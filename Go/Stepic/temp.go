package main

import "fmt" // пакет используется для проверки ответа, не удаляйте его

type battery struct {
	energy int
}

func (b battery) String() string {
	var oStr string

	for i := 0; i < 10 - b.energy; i++ {
		oStr += " "
	}
	for i := 0; i < b.energy; i++ {
		oStr += "X"
	}

	return "[" + oStr + "]"
}

func main() {
    // batteryForTest - не забывайте об имени
	var iEnergy string
	fmt.Scan(&iEnergy)
	batteryForTest := battery{}

	for _, v := range iEnergy {
		if string(v) == "1" {
			batteryForTest.energy++
		}
	}
// } Скобка, закрывающая функцию main() вам не видна, но она есть
}