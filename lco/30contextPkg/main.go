package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	ctx := context.Background()
	userId := 10

	val, err := callApi(ctx, userId)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Result:", val)
	fmt.Println("Total Time Take:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func callApi(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	respChn := make(chan Response)

	go func() {
		//val, err := mimicAPICallWithDelay(2, time.Nanosecond)
		val, err := mimicAPICallWithDelay(2, time.Second)
		respChn <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("request timed out")

		case resp := <-respChn:
			return resp.value, resp.err
		}
	}
}

func mimicAPICallWithDelay(delayValue int64, duration time.Duration) (int, error) {
	time.Sleep(time.Duration(delayValue) * duration)

	return 369, nil
}
