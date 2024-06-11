package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("The integer i says:", i)

	whatWasSaid, theOtherThing := saySomething()
	fmt.Println("The function says", whatWasSaid, "and", theOtherThing)
}

func saySomething() (string, string) {
	return "something", "else"
}
