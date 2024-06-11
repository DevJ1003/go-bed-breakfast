package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := divide(100.0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result of division:", result)
}

func divide(x, y float32) (float32, error) {

	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by 0")
	}

	result = x / y
	return result, nil
}
