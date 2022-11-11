// Задание
// Перед вами фрагмент исходного кода, который может быть скомпилирован, но требует улучшения.
// Сохраните его себе в IDE.
// Отформатируйте утилитой gofmt. Вот пример вызова: gofmt -w ..
// Проверьте линтером go vet и исправьте все найденные проблемы.
// Добавьте для описания типа и функций комментарии, которые обработает godoc.
// Запустите godoc и проверьте сгенерированную документацию в браузере.

package arrint

import (
	"fmt"
	"strings"
)

// Создание нового типа слайса
type ArrInt []int

// Функция добавления в массив
func Add(a, b ArrInt) ArrInt {
	length := len(a)
	if length-len(b) > 0 {
		length = len(b)
	}
	c := make(ArrInt, length)
	for i := 0; i < length; i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

// Метод String преобразует ArrInt в строку и возвращает её.
func (a ArrInt) String() string {
	out := make([]string,
		len(a))

	for i, v := range a {
		out[i] = fmt.Sprintf(`<%d>`, v)
	}
	return fmt.Sprintf(`[%s]`, strings.Join(out, ` `),)
}
