package main

import "fmt"

func main() {

	var myString string
	myString = "Green"
	fmt.Println("Before func call myString is:", myString)

	changUsingPointer(&myString)
	fmt.Println("After func call myString is:", myString)

}

func changUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
