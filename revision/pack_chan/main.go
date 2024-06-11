package main

import (
	"fmt"

	"github.com/devj1003/myniceprogram/helpers"
)

// import (
// 	"fmt"

// 	"github.com/devj1003/myniceprogram/helpers"
// )

// func main() {

// 	fmt.Println("Hello")

// 	var myVar helpers.SomeType
// 	myVar.TypeName = "Some Name"
// 	myVar.TypeNumber = 2

// 	fmt.Println(myVar.TypeName)
// 	fmt.Println(myVar.TypeNumber)
// }

const numPool = 1000

func CalculateValue(intChan chan int) {

	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	fmt.Println(num)
}
