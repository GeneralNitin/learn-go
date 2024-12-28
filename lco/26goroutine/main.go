package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually it is a pointer

// mutex
var signals = []string{"test"}
var mut sync.Mutex

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

	wg.Wait()
	fmt.Println(signals)
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
		mut.Lock()
		signals = append(signals, site)
		mut.Unlock()

		fmt.Printf("%v status code for the site: %v\n", get.StatusCode, site)
	}
}
