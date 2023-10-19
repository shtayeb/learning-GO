package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "done"
}

func main() {
	var channel chan string
	go printMessage("GO is great", channel)
	response := <-channel
	println(response)

	// printMessage("We miss classes !")
}
