package main

import (
	"errors"
	"log"
)

func Divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by 0")

	}
	result = x / y
	return result, nil
}

func main() {

	result, err := Divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of division is:", result)
}
