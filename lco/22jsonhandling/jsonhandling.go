package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	encodeJson()
}

func encodeJson() {
	lcoCourse := []course{
		{"GoLang", 299, "LCO.in", "pass@123", []string{"web-dev", "js"}},
		{"Java", 299, "LCO.in", "pass@123", []string{"web-dev", "js"}},
		{"JS", 0, "LCO.in", "pass@123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}
