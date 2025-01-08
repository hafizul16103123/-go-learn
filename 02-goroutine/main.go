package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var rwMut sync.RWMutex
var signals = []string{"Test"}

// func main() {
// 	bigComputation(1)
// 	bigComputation(2)
// 	bigComputation(3)
// 	bigComputation(4)
// 	bigComputation(5)
// }
// func bigComputation(i int) {
// 	count := 0
// 	for i := 0; i < 10000000000; i++ {
// 		count++
// 	}
// 	fmt.Printf("bigComputation========> %d \n", i)
// }

// with goroutine and wait group
func main() {
	fmt.Println("main")
	go bigComputation(1)
	wg.Add(1)
	go bigComputation(2)
	wg.Add(1)
	go bigComputation(3)
	wg.Add(1)
	go bigComputation(4)
	wg.Add(1)
	go bigComputation(5)
	wg.Add(1)

	fmt.Println(signals)

	wg.Wait()
}

func bigComputation(i int) {
	fmt.Printf("bigComputation start========> %d \n", i)
	defer wg.Done()
	count := 0
	for i := 0; i < 10000000000; i++ {
		count++
	}
	mut.Lock()
	signals = append(signals, "Testing")
	mut.Unlock()

	fmt.Printf("bigComputation end ========> %d \n", i)

	rwMut.RLock()
	fmt.Println(signals)
	rwMut.RUnlock()

}
