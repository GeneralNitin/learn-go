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
		log.Fatal("Error")
	}

	fmt.Println("Result:", val)
	fmt.Println("Total Time Take:", time.Since(start))
}

func callApi(ctx context.Context, userId int) (int, error) {
	val, err := mimicAPICallWithDelay(2, time.Second)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func mimicAPICallWithDelay(delayValue int64, duration time.Duration) (int, error) {
	time.Sleep(time.Duration(delayValue) * duration)

	return 369, nil
}
