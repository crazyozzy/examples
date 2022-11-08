// Есть фигуры:
// type figures int

// const(
//     square figures = iota // квадрат
//     circle // круг
//     triangle // равносторонний треугольник
// )
// Напишите функцию с такой сигнатурой:
// func area(figures)(func(float64) float64, bool)
// Функция должна возвращать:
// функцию для вычисления площади фигуры;
// true, если фигура известна;
// false, если фигура неизвестна.
// Нужно, чтобы её можно было применять так:
// ar, ok := area(myFigure)
// if !ok {
//     fmt.Println("Ошибка")
//     return
// }
// myArea := ar(x)

package main

import (
	"fmt"
	"math"
)

type figures int

const(
    square figures = iota // квадрат
    circle // круг
    triangle // равносторонний треугольник
)

const Pi = 3.1415926

func area(figure figures)(func(float64) float64, bool) {
	switch figure {
	case square:
		return func(f float64) float64 {return f* f}, true
	case circle:
		return func(f float64) float64 {return Pi *f *f}, true
	case triangle:
		return func(f float64) float64 {return f * f * math.Sqrt(3) / 4}, true
	default:
		return nil, false
	}
}

func main(){
	var myFigure figures
	myFigure = triangle

	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}

	myArea := ar(3)

	fmt.Printf("Area is %.2f", myArea)
}


// Ответ:
// func area(f figures) (func(float64) float64, bool) {
//     switch f {
//     case square:
//         return func(x float64) float64 { return x * x }, true
//     case circle:
//         return func(x float64) float64 { return 3.142 * x * x }, true
//     case triangle:
//         return func(x float64) float64 { return 0.433 * x * x }, true
//     default:
//         return nil, false
//     }
// }