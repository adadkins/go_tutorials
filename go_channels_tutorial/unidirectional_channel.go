package main

func UnidirectionalChannel(input int) int {
	// create unidirectional channel
	channel := make(chan int)
	go sendData(channel)
	read := <-channel
	return read
}

func sendData(sendCh chan<- int) {
	sendCh <- 10
}
