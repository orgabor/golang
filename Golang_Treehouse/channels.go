/*
Channels solve two problems simultaneously:
they let your goroutines communicate with each other,
and they let you synchronize your goroutines' operations.

Here's an updated version of our longTask function that takes a random number of seconds to complete.
We're using the rand package to generate a random number.
We've also updated longTask to return the number of seconds it took to its caller.


Suppose we want to call longTask as a goroutine.
We can't just say time := go longTask(). We'll get a syntax error if we try to run this.
And even if that were allowed, what would we print in the like following the call to longTask?
The goroutine doesn't return a value right away!
It's just not possible to communicate between goroutines using simple return values.
So instead, we're going to create a channel, and use that to communicate with our goroutine.
The channel will allow longTask to pass a value back to the main goroutine.
And when the main goroutine encounters the instruction to read from the channel,
it will wait until it has a value to proceed.

We call the built-in make method to create a channel.
The type for a channel is chan.
And then we need to specify what type of values the channel will accept; we'll use int.
We're going to change longTask's declaration so it accepts a channel as an argument,
and get rid of the return value. So we'll pass our new channel to longTask, get rid of the time variable,
 and call it as a goroutine.
Then, instead of printing the time variable, we'll read a value from the channel.
Then we'll change the longTask declaration to match the call in main. We'll accept a chan int parameter,
 and get rid of the return value.
Instead of returning the time we waited at the end of longTask,
 we're going to write that value to the channel. That value will then be received and printed down in main.

We can start multiple goroutines if we want,
and they can all write to the same channel.
Then as long as we read from the channel that same number of times,
the program will wait until all goroutines have finished before it exits.

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTask(channel chan int) {
	delay := rand.Intn(5)
	fmt.Println("Starting long task")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Long task finished")
	channel <- delay
}

func main() {
	rand.Seed(time.Now().Unix())
	channel := make(chan int)
	for i := 1; i <= 3; i++ {

		go longTask(channel)
	}
	for i := 1; i <= 3; i++ {
		fmt.Println("Took", <-channel, "seconds")
	}
}

/*
Feladat

In another file in the channels package, we define a writeToChannel function:

func writeToChannel(channel chan string) {
  channel <- "hello"
}
Update the readFromChannel function to do the following:

Create a new channel that takes string values
Call writeToChannel as a goroutine, and pass your new channel as an argument
Read a string from the channel and return that string


+++++++++Megoldas+++++++++++++++++++++

package channels

func readFromChannel() string {
  // CREATE A CHANNEL FOR string VALUES HERE
  channel := make (chan string)
  // HERE, CALL writeToChannel AS A GOROUTINE, AND PASS IT YOUR CHANNEL
  go writeToChannel(channel)
  // HERE, READ A STRING FROM YOUR CHANNEL AND RETURN IT
  return <-channel
}


*/
