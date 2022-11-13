// Реализуйте тип Stopwatch, который будет сохранять время каждой промежуточной фиксации секундомера и выдавать результаты относительно общего времени старта.
// Тип должен обладать следующими методами:
// Start() — запустить/сбросить секундомер;
// SaveSplit() — сохранить промежуточное время;
// GetResults() []time.Duration — вернуть текущие результаты.
// Пример работы:
// func main() {
//     sw := Stopwatch{}
//     sw.Start()
//
//     time.Sleep(1 * time.Second)
//     sw.SaveSplit()
//
//     time.Sleep(500 * time.Millisecond)
//     sw.SaveSplit()
//
//     time.Sleep(300 * time.Millisecond)
//     sw.SaveSplit()
//
//     fmt.Println(sw.GetResults())
// }
// [1.004643466s 1.505609247s 1.806269902s]

package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	startTime time.Time
	stopWatches []time.Time
}

func (s *Stopwatch) Start () {
	s.startTime = time.Now()
}

func (s *Stopwatch) SaveSplit() {
	s.stopWatches = append(s.stopWatches, time.Now())
}

func (s Stopwatch) GetResults () (res [] float64) {
	for _, v := range s.stopWatches {
		res = append(res, float64(v.Sub(s.startTime).Seconds()))
	}

	return res
}

func main() {
    sw := Stopwatch{}
    sw.Start()

    time.Sleep(1 * time.Second)
    sw.SaveSplit()

    time.Sleep(500 * time.Millisecond)
    sw.SaveSplit()

    time.Sleep(800 * time.Millisecond)
    sw.SaveSplit()

    fmt.Println(sw.GetResults())
}


// Решение:
// package main

// import (
//     "fmt"
//     "time"
// )

// type Stopwatch struct {
//     startTime time.Time
//     splits    []time.Time
// }

// func (sw *Stopwatch) Start() {
//     sw.startTime = time.Now()
//     sw.splits = nil
// }

// func (sw *Stopwatch) SaveSplit() {
//     sw.splits = append(sw.splits, time.Now())
// }

// func (sw Stopwatch) GetResults() (retResults []time.Duration) {
//     for _, splitTime := range sw.splits {
//         retResults = append(retResults, splitTime.Sub(sw.startTime))
//     }

//     return
// }

// func main() {
//     sw := Stopwatch{}
//     sw.Start()

//     time.Sleep(1 * time.Second)
//     sw.SaveSplit()

//     time.Sleep(500 * time.Millisecond)
//     sw.SaveSplit()

//     time.Sleep(300 * time.Millisecond)
//     sw.SaveSplit()

//     fmt.Println(sw.GetResults())
// }
