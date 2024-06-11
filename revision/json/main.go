package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	
	[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "brown",
		"has_dog": false
	}
	]
	`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Printf("Unmarshalled: %v", unmarshalled)

	// WRITE JSON
	var mySlice []Person

	var m1 Person

	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "grey"
	m1.HasDog = false
	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "Red"
	m2.HasDog = true
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		fmt.Println("Error marshalling", err)
	}

	fmt.Println(string(newJson))
}
