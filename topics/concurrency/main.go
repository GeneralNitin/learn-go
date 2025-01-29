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

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Add(2)
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
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Nanosecond)
		if i%2 == 0 {
			data := i + 1
			resultChan <- data
			fmt.Println("Writing: ", data)
		} else {
			//fmt.Println("Skipping: ", i)
		}
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
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Consumed: ", res)
}
