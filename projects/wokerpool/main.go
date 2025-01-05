package main

import "fmt"

func main() {
	// using the worker pool

	// create a new task
	tasks := make([]Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = Task{ID: i + 1}
	}

	// create a worker pol with the task and set the concurrency limit

	wp := WorkerPool{
		Tasks:            tasks,
		concurrencyLimit: 5, // no. of workers that can run at a time
	}

	// start the worker pool to process the tasks
	wp.Run()

	fmt.Println("All the tasks has been processed")
}
