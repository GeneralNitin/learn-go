package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests")

	performGetRequest()
	performPostRequest()
	performPostFormRequest()
}

func performGetRequest() {
	const myUrl = "http://localhost:8000/get"

	resp, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Content length:", resp.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Response 1:", responseString.String())

	//fmt.Println(string(content))
}

func performPostRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"courseName": "Go Land",
			"price": 0,
			"platform": "lco.dev"
		}`)

	resp, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)

	fmt.Println("Response 2:", string(content))
}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "Nitin")
	data.Add("lastName", "Kumar")
	data.Add("email", "nitin@nitin.com")

	resp, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)
	fmt.Println("Response 3:", string(content))
}
