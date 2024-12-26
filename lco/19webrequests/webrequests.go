package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response: %T\n", response)

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Println(content)
}
