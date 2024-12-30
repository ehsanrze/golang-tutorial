## Concurrency

### What is concurrency?
In Go, concurrency refers to the ability of a program to handle **multiple tasks at the same time**.

**🚨 Important note:** concurrency **is not parallelism**, although it enables parallelism. If you have **only one processor**, your program can **still be concurrent**, but it cannot be parallel. On the other hand, a well-written concurrent program might **run** efficiently **in parallel** on a **multiprocessor**. ([Google I/O 2012](https://www.youtube.com/watch?v=f6kdp27TYZs) - Rob Pike)

### Key concepts of the concurrency

- Goroutines
- Channels
- WaitGroup, Mutex, Ticker

#### Goroutines
A goroutine is **a lightweight thread** of execution in Go. It is managed by the Go runtime and allows you to **run functions or methods concurrently** with other goroutines, including the **main goroutine** (_which runs the main function_).

##### Key feature of goroutine
- Efficient and Lightweight
- Concurrency, Not Parallelism (We talked about it before.)
- Go Scheduler
  - The Go runtime uses a scheduler to manage goroutines. It multiplexes thousands (or even millions) of goroutines onto a small number of OS threads.
- Simple to Start
  - A goroutine is created by simply prefixing a function or method call with the go keyword.

##### Examples
```go
package main

import (
	"fmt"
	"time"
)

func proccess1() {
	fmt.Println("Process 1 has finished!")
}

func main() {

    for i := 1; i <= 5; i++ {
      go proccess1()
    }
	
	fmt.Println("Main process has being started.")
	//time.Sleep(1 * time.Second) // Why we need this?
	fmt.Println("Main process has finished.")
}
```

**Important Note:** By default, Go doesn't wait to all its goroutines to finish. We can do this in several ways:
- Use time package and sleep until all process to be finished. (It is an awful method, but it works =). )
- Use channels ( We'll discuss earlier.)

#### Channels

Go channels are used for communication between goroutines. Channels allow goroutines to safely **share data**.


##### Key feature of channels
- Type-Specific
  - Each channel is associated with a specific data type. For example, a channel of type `chan int` can only transfer integers.
- Bidirectional or Directional
  - Channels can be used for **sending and receiving** data. However, they **can also be restricted** to only sending or only receiving.
- Blocking Behavior
  - Send: When you send data into a channel, the operation blocks until another goroutine reads the data.
  - Receive: When you read from a channel, the operation blocks until another goroutine sends data

##### Examples

```go
// Declares a channel for send and receive integer values
var ch chan int 
ch = make(chan int) // Initializes the channel


go func() {
    ch <- 42 // Sends value 42 into the channel
}()

val := <-ch // Receives value from the channel
fmt.Println(val) // Outputs: 42


```
**Buffered Channels**
A buffered channel allows you to specify a capacity, so sending only blocks when the buffer is full.

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2

ch <- 3 // would block because the buffer is full
```

**Close a Channel**
You can close a channel to signal that no more values will be sent.
```go
close(ch)
```

**Restricting Channels**

- Send-Only Channels 
  - Declared with `chan <-` 
  - This channel can only send data
- Receive-Only Channels:
  - Declared with `<- chan` 
  - This channel can only receive data

```go
package main

import "fmt"

func sendData(ch chan<- int) {
    ch <- 42
}

func receiveData(ch <-chan int) {
    fmt.Println(<-ch)
}

func main() {
    ch := make(chan int)

    go sendData(ch)
    receiveData(ch)
}
```
