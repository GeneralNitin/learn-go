package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	ch := make(chan int)
	//ch := make(chan int, 1) // buffered channel, that 1 can be any value

	//ch <- 5
	//fmt.Println(<-ch)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	//go func(ch <-chan int, wg *sync.WaitGroup) { // <-chan makes the channel Read ONLY
	go func(ch chan int, wg *sync.WaitGroup) {
		//value, isChannelOpen := <-ch // to know if the channel is closed and the value is the zero value i.e. int 0, string "" etc
		value := <-ch

		fmt.Println("Received Value:", value)
		//fmt.Println("Received Value:", <-ch) // do this with buffered channel
		wg.Done()
	}(ch, wg)

	//go func(ch chan<- int, wg *sync.WaitGroup) { // chan<- makes the channel Write ONLY
	go func(ch chan int, wg *sync.WaitGroup) {
		// adding value to channel
		ch <- 5
		//close(ch) // closing a channel, we can still read from a closed channel
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
