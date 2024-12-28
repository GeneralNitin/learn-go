package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually it is a pointer

func main() {
	//go greeter("Hello")
	//greeter("World")

	sites := []string{
		"https://www.google.com",
		"https://www.x.com",
		"https://www.facebook.com",
		"https://www.github.com",
	}

	for _, site := range sites {
		go getStatusCode(site)
		wg.Add(1)
	}

	defer wg.Wait()
}

//func greeter(s string) {
//	for i := 0; i <= 5; i++ {
//		fmt.Println(s)
//		time.Sleep(3 * time.Millisecond)
//	}
//}

func getStatusCode(site string) {
	defer wg.Done()

	get, err := http.Get(site)
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
	} else {
		fmt.Printf("%v status code for the site: %v\n", get.StatusCode, site)
	}
}
