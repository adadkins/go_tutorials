package main

import (
	"fmt"
)

func squareCalc(input int, squareChan chan int) {
	sum := 0
	for i := 0; i < input; i++ {
		sum += input
	}
	squareChan <- sum
}

func cubeCalc(input int, cubeChan chan int) {
	firstSum := 0
	for i := 0; i < input; i++ {
		firstSum += input
	}
	secondSum := 0
	for i := 0; i < input; i++ {
		secondSum += firstSum
	}
	cubeChan <- secondSum
}

func SquareCubeAddition(input int) int {
	squareChan := make(chan int)
	cubeChan := make(chan int)

	fmt.Println("Main going to call hello go goroutine")
	go squareCalc(input, squareChan)
	go cubeCalc(input, cubeChan)

	squareResult := <-squareChan
	cubeResult := <-cubeChan

	fmt.Println("Sum:", squareResult+cubeResult)
	return squareResult + cubeResult
}
