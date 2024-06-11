package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Daniel",
	}

	myMap["me"] = me
	fmt.Println(myMap["me"].FirstName, myMap["me"].LastName)

}
