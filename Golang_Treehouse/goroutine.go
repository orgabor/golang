/*
A "goroutine" is a simple way to make several function calls simultaneously.
The work gets split up among all your CPU cores, and they all work on it at the same time.

*/
/*
A goroutine is a simple way to make several function calls simultaneously.
The work gets split up among your CPU cores, and they all work on it at the same time.
A Go program starts with a single goroutine, which runs the main function

*/

/*
The first longTask call runs, sleeps for 3 seconds, and then returns.
Then the second longTask call runs, sleeps for 3 seconds, and so on.
Because the calls to longTask run one at a time, the program takes just over 9 seconds to complete.
Prepend the go keyword to a function call to launch another goroutine, which runs alongside the first.

*/

/*
Now calls to longTask run in separate goroutines.
After each goroutine kicks off, it goes back to the main goroutine and launches the next goroutine.
The problem is that as soon as the main goroutine finishes, the program exits.
So the other goroutines don't get a chance to do anything.
So just as a quick fix, we'll add a call to Sleep to the main function:

*/

/*
More info:
Rob Pike - 'Concurrency Is Not Parallelism'
https://www.youtube.com/watch?v=cN_DpYBzKso
Rob Pike - 'Concurrency Is Not Parallelism'

https://tour.golang.org/concurrency/1

*/

package main

import (
	"fmt"
	"time"
)

func longTask() {
	fmt.Println("Starting long task")
	time.Sleep(3 * time.Second)
	fmt.Println("Long task finished")
}

func main() {
	go longTask()
	go longTask()
	go longTask()
	time.Sleep(4 * time.Second)
}
