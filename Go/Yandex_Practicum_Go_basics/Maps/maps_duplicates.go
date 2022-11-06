// Дедупликация строк
// Представьте, что на входе есть слайс строк:
// input := []string{
//     "cat",
//     "dog",
//     "bird",
//     "dog",
//     "parrot",
//     "cat",
// } 
// Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса:
// func RemoveDuplicates(input []string) []string

package main

import "fmt"

func RemoveDuplicates(input []string) []string {
	var output []string
	temp_map := make(map[string]int, len(input))

	for index, val := range input {
		if _, ok := temp_map[val]; ok == false {
			output = append(output, val)
			temp_map[val] = index
		}
	}

	return output
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	} 

	fmt.Println(RemoveDuplicates(input))
}



// Ответ:
// package main

// import "fmt"

// func RemoveDuplicates(input []string) []string {
//     output := make([]string, len(input))
//     copy(output, input)

//     inputSet := make(map[string]struct{}, len(input))
//     outputIdx := 0
//     for _, v := range input {
//         if _, ok := inputSet[v]; !ok {
//             output[outputIdx] = v
//             outputIdx++
//         }
//         inputSet[v] = struct{}{}
//     }

//     return output[:outputIdx]
// }

// func main() {
//     input := []string{
//         "cat",
//         "dog",
//         "bird",
//         "dog",
//         "parrot",
//         "cat",
//     }
//     fmt.Println(input)

//     fmt.Println(RemoveDuplicates(input))
// }