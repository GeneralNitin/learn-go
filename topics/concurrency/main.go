package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//ctxFp := context.Background()
	//context.WithTimeout(context.Background(), 100*time.Second)
	fpResultChan := make(chan int, 10)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		//defer wg.Done()
		//clientCall(fpResultChan)
		//defer wg.Done()
		go clientCall(&wg, fpResultChan)
		go processResponse(&wg, fpResultChan)
	}()

	//go func() {
	//	defer wg.Done()
	//	processResponse(fpResultChan)
	//}()

	wg.Wait()
	fmt.Println("Done!!!")
}

func clientCall(wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		resultChan <- i + 1
		fmt.Println("Writing: ", i)
	}
	close(resultChan)
}

func processResponse(wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	for {
		select {
		case res, ok := <-resultChan:
			if !ok {
				return
			}
			consumeResponse(res)
		}
	}
}

func consumeResponse(res int) {
	time.Sleep(250 * time.Millisecond)
	fmt.Println("Result: ", res)
}
