package main

import (
	"fmt"
	"sync"
	"time"
)

// Single unit of work - Task Definition
type Task interface {
	Process()
}

// Type 1: Email Task
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// function to process the Task - Way to process the task
func (t *EmailTask) Process() {
	fmt.Printf("Sending Email to %v\n", t.Email)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Type 2: Image Processing Task
type ImageProcessingTask struct {
	ImageUrl string
}

// function to process the Task - Way to process the task
func (t *ImageProcessingTask) Process() {
	fmt.Printf("Processing the image %v\n", t.ImageUrl)
	// Simulate a time consuming process
	time.Sleep(5 * time.Second)
}

// Worker Pool Definition
type WorkerPool struct {
	Tasks            []Task         // Tasks List
	concurrencyLimit int            // No of concurrent process that can be handled together
	tasksChan        chan Task      // defines a channel to send task to workers
	wg               sync.WaitGroup // Synchronising the completion of task. This will we used to wait for the task to be completed
}

// Worker Pool function, to execute the task in the pool

// method that receives tasks from the task channels and processes them
// defining a method on WorkerPool Struct

func (pool *WorkerPool) worker() {
	for task := range pool.tasksChan {
		task.Process()
		pool.wg.Done()
	}
}

// The run method: This method initialises the channel, sets the concurrency, creates the go routines ans sends tasks
// over the channel
func (pool *WorkerPool) Run() {
	// Initialize the tasks channel with the capacity equals the number of tasks
	pool.tasksChan = make(chan Task, len(pool.Tasks))

	// Start the worker go routines, number of workers equals to the concurrency variable
	for i := 0; i < pool.concurrencyLimit; i++ {
		go pool.worker()
	}

	// add the number of tasks in the wait group counter
	pool.wg.Add(len(pool.Tasks))

	// send tasks to the task channel
	for _, task := range pool.Tasks {
		pool.tasksChan <- task
	}

	// close the task chan after sending all tasks to Signal no more tasks will be sent
	close(pool.tasksChan)

	// wait for all the tasks to be processed using the wait method of the wait group
	pool.wg.Wait()
}
