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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// initiate channel
	ch := make(chan string)

	// iterate over links and check status
	for _, link := range links {
		go checkLinkStatus(link, ch)
	}

	// infinite loop to keep the program running
	for l := range ch {
		// function literal to create a new goroutine
		go func(link string) {
			// sleep for 5 seconds before checking the link again
			time.Sleep(5 * time.Second)
			checkLinkStatus(link, ch)
		}(l)
	}

}

func checkLinkStatus(link string, c chan string) {
	// check if link is up
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send link to channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send link to channel
	c <- link
}
