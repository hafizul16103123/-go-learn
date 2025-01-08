package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// Mutax lock memory when race condition happened
	mut := &sync.Mutex{}
	score := []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("ADD Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("ADD Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("ADD Four")
		mut.Lock()
		score = append(score, 4)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
