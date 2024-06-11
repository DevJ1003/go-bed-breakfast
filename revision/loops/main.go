package main

import "fmt"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Haan", "mary@haan.com", 20})
	users = append(users, User{"Lia", "Waah", "lia@waah.com", 40})
	users = append(users, User{"Taro", "Hola", "taro@hola.com", 50})

	for _, l := range users {
		fmt.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
