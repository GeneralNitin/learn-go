package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["Java"] = "Java"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	// read specific value
	fmt.Println("JS shorts for:", languages["JS"])
	fmt.Printf("JS shorts for: %v, with type: %T\n", languages["404"], languages["404"])

	// delete value
	delete(languages, "RB")

	// loops through key and value
	for key, value := range languages {
		fmt.Printf("Key: %v and Value: %v\n", key, value)
	}

	// contains?
	value, ok := languages["RB"]
	fmt.Printf("Value of non existing key is: %v and OK: %v\n", value, ok)

	_, exists := languages["RB"]
	if !exists {
		fmt.Printf("1:Value of Exists: %v\n", exists)
	} else {
		fmt.Printf("2:Value of Exists: %v\n", exists)
	}

	fmt.Printf("if (languages[\"404\"]) == \"\"; %v", (languages["404"]) == "")
}
