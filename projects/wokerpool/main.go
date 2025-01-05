package main

import "fmt"

func main() {
	// using the worker pool

	// create a new task
	//tasks := make([]Task, 20)
	//for i := 0; i < 20; i++ {
	//	tasks[i] = Task{ID: i + 1}
	//}

	tasks := []Task{
		&EmailTask{Email: "n1@n.com", Subject: "Hello", MessageBody: "Worl"},
		&EmailTask{Email: "n2@n.com", Subject: "Hello", MessageBody: "Worl"},
		&EmailTask{Email: "n3@n.com", Subject: "Hello", MessageBody: "Worl"},
		&EmailTask{Email: "n4@n.com", Subject: "Hello", MessageBody: "Worl"},
		&EmailTask{Email: "n5@n.com", Subject: "Hello", MessageBody: "Worl"},
		&ImageProcessingTask{ImageUrl: "nitin1.png"},
		&ImageProcessingTask{ImageUrl: "nitin2.png"},
		&ImageProcessingTask{ImageUrl: "nitin3.png"},
		&ImageProcessingTask{ImageUrl: "nitin4.png"},
		&ImageProcessingTask{ImageUrl: "nitin5.png"},
	}

	// create a worker pol with the task and set the concurrency limit

	pool := WorkerPool{
		Tasks:            tasks,
		concurrencyLimit: 3, // no. of workers that can run at a time
	}

	// start the worker pool to process the tasks
	pool.Run()

	fmt.Println("All the tasks has been processed")
}
