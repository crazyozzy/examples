// В последнем рассмотренном примере реализация функции Read не очень эффективна — генератор случайных чисел возвращает 64-битное число, то есть 8 байт.
// Из них используем только 1.

// Попробуйте реализовать более эффективное решение.
// Для упрощения примера считайте, что функция будет принимать только слайсы, длина которых кратна 8.
// Для преобразования числа в слайс байт можно использовать функцию из стандартной библиотеки binary.LittleEndian.PutUint64([ ]byte, uint64).

package randbyte

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
)

type generator struct {
    rnd rand.Source // Генератор случайных чисел. Вообще rand.Rand уже реализует интерфейс io.Reader, но для примера мы реализуем его самостоятельно.
}

// New — обратите внимание, что мы возвращаем generator, присвоенный интерфейсу io.Reader, сама структура generator неэкспортируемая.
// Мы скрыли внутри пакета все детали.
func New(seed int64) io.Reader {
    return &generator{
        rnd: rand.NewSource(seed),
    }
}

// Read — реализация io.Reader
func (g *generator) Read(bytes []byte) (n int, err error) { // error — это тип ошибки, подробнее мы рассмотрим его в следующем разделе.
    if len(bytes) % 8 != 0 {
        return 0, err
    }

    for i := 0; i < len(bytes); i += 8 {
        randInt := g.rnd.Int63()  // функция возвращает положительное число в пределах от 0 до 2^63
        fmt.Println(randInt)
        // randByte := byte(randInt) // приводим к типу byte
        // bytes[i] = randByte
        binary.LittleEndian.PutUint64(bytes[i:], uint64(randInt))
    }
    return len(bytes), nil
}


// Ответ:
// // Read — реализация io.Reader
// func (g *generator) Read(bytes []byte) (n int, err error) { // error это тип ошибки, подробнее мы рассмотрим его в следующем разделе.
//     for i := 0; i+8 < len(bytes); i += 8 {
//         binary.LittleEndian.PutUint64(bytes[i:i+8], uint64(g.rnd.Int63()))
//     }
//     return len(bytes), nil
// }