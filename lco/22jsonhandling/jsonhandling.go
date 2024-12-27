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
	//encodeJson()
	decodeJson()
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

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "GoLang",
		"Price": 299,
		"website": "LCO.in",
		"tags": ["web-dev","js"]
	}`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was invalid")
	}

	// other way of unMarshaling json in a map way
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("Output2: %#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key: %v :: Value: %v :: Type: %T\n", k, v, v)
	}

}
