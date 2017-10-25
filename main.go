package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://clubnacionalquerido.com",
		"https://google.com",
		"http://guaranigrado.fi.uba.ar",
		"https://www.lingscars.com/",
		"https://golang.org/",
		"https://noexisteestesitio.co.uk",
	}

	c := make(chan string)

	for _, url := range urls {
		go getStatus(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func getStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- url + " is down!"
	} else {
		c <- url + " is working!"
	}
}
