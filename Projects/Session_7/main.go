package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// #1 Goroutine and Ticker Example
	//timerProgram(10)

	// #2 Wait Group Example
	//multipleWorkerProgram(10)

	// #3 Mutex Example
	//counterProgram(1000)

}

func timerProgram(seconds int) {
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

func sendData(data chan<- string) {
	data <- fmt.Sprintf("New Message - Time: %s", time.Now())
}

func receiveData(data string) {
	fmt.Println(data)
}

func multipleWorkerProgram(numberOfWorkers int) {
	var wg sync.WaitGroup

	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers completed")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second) // Simulate work.

	fmt.Printf("Worker %d done\n", id)
}

var (
	counter int
	mutex   sync.Mutex
)

func counterProgram(n int) {
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	counter++
	fmt.Println("Counter:", counter)
	mutex.Unlock()

}
