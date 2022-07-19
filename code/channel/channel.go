package main

import "fmt"

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := numPool
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	fmt.Print(num)
}
