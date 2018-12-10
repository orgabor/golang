package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"https://orgabor.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) //2nd argument channel
	}

	for l := range c { //infinite loop l is for link
		//	fmt.Println(<-c) #original code
		// t.Sleep(3 * time.Second) //pauses the routine for 3 seconds in the main routine -bad idea
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l) //call anonymus function - function literal
		//passing l as an argument to function literal
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {

		fmt.Println(link + "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link

}
