// Задание 1 из 4
// Попробуйте реализовать обобщение операции умножения для чисел и строк.
// Если первый аргумент функции — строка, то повторить её b раз, а если число, то вернуть a*b.
//
// func Mul(a interface{}, b int) interface{} {
//
// }

package main

import "fmt"

func Mul(a interface{}, b int) interface{} {
	switch a.(type) {
	case string:
		res := ""
		for i := 1; i <= b; i++ {
			res += a.(string)
		}
		return res
	case int:
		return a.(int) * b
	default:
		return nil
	}
}

func main() {
	s := "This is string"
	n := 10

	fmt.Println(Mul(s, 3))
	fmt.Println(Mul(n, 5))
}


// Ответ:
// func Mul(a interface{}, b int) interface{} {
//     switch va := a.(type) {
//     case int:
//         return va * b
//     case string:
//         return strings.Repeat(va, b)
//     case fmt.Stringer:
//         return strings.Repeat(va.String(), b)
//     default :
//         return nil
//     }
// }
