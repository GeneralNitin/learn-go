package main

import "sync"

func main() {
	ch := make(chan int)
	//ch := make(chan int, 1) //

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {

	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {

	}(ch, wg)

	wg.Done()
}
