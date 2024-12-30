package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	user := getUserFromDB()
	wg := sync.WaitGroup{}
	ch := make(chan any, 2)

	wg.Add(2)
	go getTotalFriendsCount(&wg, ch)
	go getFullName(&wg, ch)
	wg.Wait()
	close(ch)

	endTime := time.Since(startTime)
	fmt.Println(endTime)

	fmt.Println(user)
	for msg := range ch {
		fmt.Println(msg)
	}
}

func getUserFromDB() string {
	time.Sleep(100 * time.Millisecond)
	return "Nitin"
}

func getTotalFriendsCount(wg *sync.WaitGroup, ch chan any) {
	time.Sleep(150 * time.Millisecond)
	wg.Done()
	ch <- 1000
}

func getFullName(wg *sync.WaitGroup, ch chan any) {
	time.Sleep(100 * time.Millisecond)
	wg.Done()
	ch <- "Nitin Kumar"
}
