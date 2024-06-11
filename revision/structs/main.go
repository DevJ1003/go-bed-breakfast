package main

import (
	"fmt"
)

type myStruct struct {
	FirstName string
}

func (f *myStruct) printFirstName() string {
	return f.FirstName
}

func main() {

	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	// fmt.Println("myVar is set to", myVar.FirstName)
	// fmt.Println("myVar2 is se to", myVar2.FirstName)
	fmt.Println("myVar is set to", myVar.printFirstName())
	fmt.Println("myVar2 is se to", myVar2.printFirstName())

}
