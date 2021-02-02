package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkURL(wg *sync.WaitGroup, link string) bool {

	defer wg.Done()

	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error: ", link, err)
	}
	fmt.Println(link, " is up")
	return true
}

func main() {
	var wg sync.WaitGroup

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

	for _, link := range website {
		wg.Add(1)
		go checkURL(&wg, link)
	}

	wg.Wait()
	fmt.Println("All done!")
}
