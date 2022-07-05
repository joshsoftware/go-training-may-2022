package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops in endpoint")
	} else {
		fmt.Printf("%d code for %s\n", res.StatusCode, endpoint)
	}
}

var wg sync.WaitGroup

func main() {

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://gooogle.dev", // it will print oops in endpoint
		"https://github.dev",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}
