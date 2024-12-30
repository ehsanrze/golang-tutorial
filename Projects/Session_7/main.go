package main

import (
	"fmt"
	"time"
)

func sendData(data chan<- string) {
	data <- fmt.Sprintf("New Message - Time: %s", time.Now())
}

func receiveData(data string) {
	fmt.Println(data)
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	stop := time.After(10 * time.Second)

	ch := make(chan string)

	for {
		select {
		case <-ticker.C:
			go sendData(ch)
		case msg := <-ch:
			receiveData(msg)
		case <-stop:
			ticker.Stop()
			fmt.Println("The counter has been stopped.")
			return
		}
	}
}
