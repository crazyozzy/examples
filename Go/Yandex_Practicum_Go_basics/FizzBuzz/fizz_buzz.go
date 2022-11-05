// Перед вами классическое задание для собеседований на любом языке программирования — FizzBuzz.
// Напишите программу, которая выводит на экран числа от 1 до 100.
// При этом вместо чисел, кратных трём, программа должна выводить слово Fizz, а вместо кратных пяти — Buzz.
// Если число кратно и трём, и пяти, программа должна выводить слово FizzBuzz.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case ( i % 3 == 0 && i % 5 == 0 ):
			fmt.Println("FizzBuzz")
		case ( i % 3 == 0 ):
			fmt.Println("Fizz")
		case ( i % 5 == 0 ):
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}



// Ответ:
// package main
//
// import "fmt"
//
// func main() {
//     for i := 1; i <= 100; i++ {
//         found := false
//
//         if i%3 == 0 {
//             fmt.Printf("Fizz")
//             found = true
//         }
//         if i%5 == 0 {
//             fmt.Printf("Buzz")
//             found = true
//         }
//
//         if !found {
//             fmt.Println(i)
//             continue
//         }
//
//         fmt.Println()
//     }
// }
