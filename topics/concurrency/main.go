package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//ctxFp := context.Background()
	//context.WithTimeout(context.Background(), 100*time.Second)
	fpResultChan := make(chan int, 1000000)

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
	defer func() {
		fmt.Println("Will Close The Channel")
		close(resultChan)
	}()
	for i := 0; i < 10000; i++ {
		time.Sleep(1 * time.Second)
		//if i%2 == 0 {
		data := i
		resultChan <- data
		fmt.Println("Writing: ", data)
		//} else {
		//	//fmt.Println("Skipping: ", i)
		//}
	}
}

func processResponse(wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	for res := range resultChan {
		//select {
		//case res, ok := <-resultChan:
		//	if !ok {
		//		return
		//	}
		consumeResponse(res)
		//}
	}
}

func consumeResponse(res int) {
	//time.Sleep(2 * time.Second)
	fmt.Println("Consumed: ", res)
}
