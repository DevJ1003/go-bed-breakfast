package main

import "fmt"

func main() {

	myVar := "rhino"

	switch myVar {
	case "cat":
		fmt.Println("cat is set to cat")

	case "dog":
		fmt.Println("cat is set to dog")

	case "fish":
		fmt.Println("cat is set to fish")

	case "horse":
		fmt.Println("cat is set to horse")

	default:
		fmt.Println("cat is something else")
	}
}
