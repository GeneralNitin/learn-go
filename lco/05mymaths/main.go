package main

import (
	cr "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to maths in go lang")

	var num1 = 2
	var num2 = 4.5
	fmt.Println("The sum is:", num1+int(num2))
	fmt.Println("The sum is:", float64(num1)+num2)

	// random number
	//rand.Seed(time.Now().UnixNano())
	//rand.New(rand.NewSource(4))
	fmt.Println(rand.Intn(5), ": Random num from rand lib") // random from 0 to 4

	randNumber, _ := cr.Int(cr.Reader, big.NewInt(5))
	fmt.Println(randNumber, ": Random num from crypto lin") // random from 0 to 4
}
