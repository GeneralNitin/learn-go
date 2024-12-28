package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition - LCO")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	//mut2 := &sync.RWMutex{}

	var scores = []int{0}

	wg.Add(3)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One Routine")
		mut.Lock()
		//mut2.RLock()
		scores = append(scores, 1)
		//mut2.RUnlock()
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three Routine")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(scores)
}
