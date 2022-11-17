// Задание 4
// Напишите функцию MyCheck(input string) error, которая должна проверять строку на соответствие следующим параметрам:
// Строка не должна содержать цифр (found numbers).
// Длина должна быть меньше 20 символов (line is too long).
// Строка должна иметь два пробела (no two spaces).
// В скобках приведены тексты возвращаемых ошибок.
// Функция должна находить все возможные ошибки за один вызов. Результат должен содержать слайс ошибок, по которым строка не прошла проверку, или быть nil. Подсказка: определите свой тип для слайса ошибок.
// package main

// import (
//     "bufio"
//     "errors"
//     "fmt"
//     "os"
//     "strings"
// )

// // 1) вставьте определение типа для []error
// // 2) определите метод Error для вашего типа, который будет выводить
// //    все ошибки слайса
// // 3) реализуйте функцию MyCheck
// //
// // ...

// func main() {
//     for {
//         fmt.Printf("Укажите строку (q для выхода): ")
//         reader := bufio.NewReader(os.Stdin)
//         ret, err := reader.ReadString('\n')
//         if err != nil {
//             fmt.Println(err)
//             continue
//         }
//         ret = strings.TrimRight(ret, "\n")
//         if ret == `q` {
//             break
//         }
//         if err = MyCheck(ret); err != nil {
//             fmt.Println(err)
//         } else {
//             fmt.Println(`Строка прошла проверку`)
//         }
//     }
// }

package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

// 1) вставьте определение типа для []error
type Errors []error

// 2) определите метод Error для вашего типа, который будет выводить все ошибки слайса
func (err Errors) Error() string {
	var output string
	for _, v := range err {
		output += v.Error() + "; "
	}

	return output
}

// 3) реализуйте функцию MyCheck
func MyCheck(input string) error {
	var (
		err Errors
		hasDigits bool
		spaces uint8
	)

	if len([]rune(input)) >= 20 {
		err = append(err, errors.New("Line is too long (>=20 characters)!"))
	}

	for _, v := range input {
		switch {
		case v >= '0' && v <= '9':
			hasDigits = true
		case v == ' ':
			spaces++
		}
	}

	if hasDigits {
		err = append(err, errors.New("Line has digits!"))
	}

	if spaces < 2 {
		err = append(err, errors.New("Line must have at least two spaces!"))
	}

	if len(err) == 0 {
		return nil
	}

	return err
}

//
// ...

func main() {
    for {
        fmt.Printf("Укажите строку (q для выхода): ")
        reader := bufio.NewReader(os.Stdin)
        ret, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(err)
            continue
        }
        ret = strings.TrimRight(ret, "\n")
        if ret == `q` {
            break
        }
        if err = MyCheck(ret); err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(`Строка прошла проверку`)
        }
    }
} 



// Ответ:
// type SliceError []error

// func (errs SliceError) Error() string {
//     var out string
//     for _, err := range errs {
//         out += err.Error() + ";"
//     }
//     return strings.TrimRight(out, `;`)
// }

// func MyCheck(input string) error {
//     var (
//         err      SliceError
//         spaces   int
//         hasDigit bool
//     )
//     if len([]rune(input)) >= 20 {
//         err = SliceError{errors.New(`line is too long`)}
//     }
//     for _, ch := range input {
//         if ch == ' ' {
//             spaces++
//         } else if ch >= '0' && ch <= '9' {
//             hasDigit = true
//         }
//     }
//     if hasDigit {
//         err = append(err, errors.New(`found numbers`))
//     }
//     if spaces != 2 {
//         err = append(err, errors.New(`no two spaces`))
//     }
//     if len(err) == 0 {
//         return nil
//     }
//     return err
// }