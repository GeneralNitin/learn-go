package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go Lang")

	fileName := "./myfile.txt"
	content := "This needs to go in a file"

	file, err := os.Create(fileName)
	//if err != nil {
	//	panic(err)
	//}
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)

	defer file.Close()

	readFile(fileName)
}

func readFile(fileName string) {
	bytes, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Data Bytes:", bytes)
	fmt.Println("Data Bytes:", string(bytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
