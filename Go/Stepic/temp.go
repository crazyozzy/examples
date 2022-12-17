package main

import (
	"fmt" // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"sync"
	// "time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(){
        x1, x2 := make([]int, n), make([]int, n)
		wg := new(sync.WaitGroup)

		for i := 0; i < n; i++ {
			x1[i] = <- in1
			x2[i] = <- in2
		}

		for i := 0; i < n; i++ {
			go func (i int, x1 []int, wg *sync.WaitGroup)  {
				wg.Add(1)
				x1[i] = fn(x1[i])
				wg.Done()
			}(i, x1, wg)
			go func (i int, x2 []int, wg *sync.WaitGroup)  {
				wg.Add(1)
				x2[i] = fn(x2[i])
				wg.Done()
			}(i, x2, wg)
		}
		wg.Wait()

		for i := 0; i < n; i++ {
			out <- x1[i] + x2[i]
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