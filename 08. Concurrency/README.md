## Concurrency

### What is concurrency?
In Go, concurrency refers to the ability of a program to handle **multiple tasks at the same time**.

**🚨 Important note:** concurrency **is not parallelism**, although it enables parallelism. If you have **only one processor**, your program can **still be concurrent**, but it cannot be parallel. On the other hand, a well-written concurrent program might **run** efficiently **in parallel** on a **multiprocessor**. ([Google I/O 2012](https://www.youtube.com/watch?v=f6kdp27TYZs) - Rob Pike)

### Key concepts of the concurrency

- Goroutines
- Channels
- WaitGroup, Mutex, Ticker
- Defer

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


#### Advanced Topics

##### WaitGroup
A WaitGroup is a synchronization primitive provided by the `sync` package that helps you **manage** the **execution of multiple goroutines** and **wait for them to complete** before proceeding further in your code.

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {

  var wg sync.WaitGroup

  for i := 1; i <= 5; i++ {
    wg.Add(1)
    go worker(i, &wg)
  }

  wg.Wait()
  fmt.Println("All workers completed")
}

func worker(id int, wg *sync.WaitGroup) {
  defer wg.Done()
  fmt.Printf("Worker %d starting\n", id)

  time.Sleep(time.Second)

  fmt.Printf("Worker %d done\n", id)
}

```

**Key Points**

- Concurrency Management
  - WaitGroup helps coordinate multiple goroutines, ensuring all of them complete before proceeding.
- Thread-Safe
  - WaitGroup is safe to use across multiple goroutines because it is designed to handle concurrent access.
- Defer Usage (**Important**)
  - It is common practice to use defer wg.Done() in a goroutine to ensure the Done() call is made regardless of how the function exits.



##### Mutex

Mutex (short for "**mutual exclusion**") is a synchronization primitive used to ensure that **only one goroutine can access a shared resource** (such as a variable, map, or structure) at a time. This helps **prevent race conditions** and ensures data consistency when multiple goroutines access or modify shared data.

**Key concepts**

- Lock()
  - Acquires the lock
  - If another goroutine has already locked the mutex, the calling goroutine will block (wait) until the lock becomes available
- UnLock()
  - Releases the lock
  - If other goroutines are waiting for the lock, one of them will acquire it

**Example**
```go

package main

import (
	"fmt"
	"sync"
)

const n = 10

var (
  counter int
  mutex   sync.Mutex
)

func main() {
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
```

**Important Notes:** 
- **Always use `Unlock()` in a defer statement right after acquiring the lock** with Lock(). This ensures the mutex is released even if the function exits due to an error or early return.
```go
mutex.Lock()
defer mutex.Unlock() // to ensure you always unlock the key
```
- **Avoid holding a lock for an extended period** to prevent blocking other goroutines
- Be **cautious of deadlocks**, which can occur if two or more goroutines wait indefinitely for each other to release locks


##### Ticker
Ticker is provided by the time package and is used to send periodic events on a channel at a specified interval. It is helpful when you want to perform an action repeatedly at fixed time intervals.

**How does it work?** 
- Ticker holds a **channel (C)** that **delivers the current time** at regular intervals defined by its duration. 
- You can **use it in a for loop** or a **select statement** to perform tasks periodically.

**Example**
```go
package main

import (
	"time"
	"fmt"
)

const seconds = 10

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

func sendData(data chan<- string) {
	data <- fmt.Sprintf("New Message - Time: %s", time.Now())
}

func receiveData(data string) {
	fmt.Println(data)
}

```

**Important Notes:** 
- **Always stop the ticker** using `Stop()` when you no longer need it to avoid resource leaks. 
- If you need to perform a one-time delayed operation instead of periodic ticks, consider using **time.After** or **time.Sleep** instead


#### Defer
`defer` is a keyword used to ensure that a function call is executed at the end of the enclosing function's execution, just before the function returns. It’s often used for tasks like cleaning up resources, closing files, unlocking mutexes, and handling other cleanup actions.

**Key points**
- Execution Order
  - Deferred calls are **executed in last in, first out (LIFO)** order. If multiple defer statements are used, the most recently deferred call is executed first
- Timing
  - Deferred calls are executed after the surrounding function finishes but before it returns to the caller.


**Examples**


```go
package main

import "fmt"

func main() {
    fmt.Println("Start")

    defer fmt.Println("Deferred: 1")
    defer fmt.Println("Deferred: 2")

    fmt.Println("End")
}

// Output
// Start
// End
// Deferred: 2 - We've talked about the orders earlier
// Deferred: 1
```

**Important note**
Arguments to the deferred function are **evaluated immediately** when the defer statement is executed, not when the deferred function is called. (**Arguments Evaluation**)

```go
package main

import "fmt"

func main() {
	x := 10
	defer fmt.Println("Deferred x:", x) // x is evaluated immediately

	x = 20
	fmt.Println("Value of x:", x)
}

// Output
// Value of x: 20
// Deferred x: 10

```