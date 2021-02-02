package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkURL(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error: %s %s\n", link, err)
		c <- link
		return
	}
	fmt.Printf("%s is up\n", link)
	c <- link
}

func main() {
	website := []string{
		"http://localhost:8080",
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://twitter.com",
		"http://instagram.com",
	}

	c := make(chan string)

	for _, link := range website {
		go checkURL(c, link)
	}

	fmt.Println("All started!")
	/* Keep getting value from channel forever and for each value received, run the for loop */
	for l := range c {
		/* goroutine with a function literal */
		go func(link string) { /* function literal accepts a string */
			time.Sleep(2 * time.Second)
			checkURL(c, link)
		}(l) /* pass string link value received to function literal
		 * l is passed by value inside the function literal.
		 * Accepted as link.
		 * Reason to pass it by value, and not ptr/reference:
		 * The value of l keeps updating as messages keep coming from channel.
		 * 		Since we sleep inside the lambda, the value of `l` could have changed.
		 * 		link inside lambda could have been updated. So lambda gets executed for
		 *		incorrect input.
		 */
	}

	/*
	 * Fetch from channel the expected number of messages.
	 *
	 * time.Sleep(3 * time.Second)
	 * for i := len(website); i > 0; i-- {
	 * 	fmt.Println(<-c)
	 * }
	 */

	/*
	 * Forever keep waiting for message on channel,
	 * run loop for every value received.
	 * for l := range c {
	 *   fmt.Println(l)
	 * }
	 */
	fmt.Println("All done!")
}
