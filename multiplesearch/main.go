package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urls := []string{
		"http://localhost:8081/miguel",
		"http://localhost:8081/fideos",
		"http://localhost:8081/gatos",
		"http://localhost:8081/agua",
		"http://localhost:8081/pascal",
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
	res, err := http.Get(url)
	if err != nil {
		c <- url + " is down!"
	} else {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		c <- bodyString
	}
}
