package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:300/learn?coursename=react&paymentid=1baecXW=="

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println("protocol/scheme:", result.Scheme)
	fmt.Println("host:", result.Host)
	fmt.Println("path:", result.Path)
	fmt.Println("port():", result.Port())
	fmt.Println("raw query:", result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params are: %t\n", qParams)

	fmt.Println(qParams["coursename"])
	fmt.Println(qParams["paymentid"])

	for key, val := range qParams {
		fmt.Printf("Param Name and Value: %v::%v\n", key, val)
	}

	partsOfUrl := url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
