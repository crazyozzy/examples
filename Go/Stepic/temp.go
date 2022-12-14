package main

import (
	"fmt" // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"sync"
	// "time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	locker := new(sync.Mutex)

	go func(){
		for i := 0; i < n; i++ {
			
			// locker.Lock()
			// x1 := <- in1
			// x2 := <- in2
			// locker.Unlock()
			// go func (x1, x2, i int)  {
			// 	out <- fn(x1) + fn(x2)
			// }(x1, x2, i)
		}
	}()
}

const N = 5

func main() {

    in1, in2, out := make(chan int, 1), make(chan int, 1), make(chan int, N)

    merge2Channels(some, in1, in2, out, N)

    for i := 0; i < N; i++ {

        in1 <- i

        in2 <- i

    }

    for i := 0; i < N; i++ {

        fmt.Println(<-out)

    }

}

func some(val int) int {

    return val

}