package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Path[1:]
	engines := []searchEngine{
		imagesEngine{},
		videosEngine{},
		linksEngine{},
	}
	results := make([]string, 0, 50)
	c := make(chan []string)
	for _, engine := range engines {
		go func(e searchEngine) {
			c <- e.getResults(query)
		}(engine)
	}

	for i := 0; i < len(engines); i++ {
		result := <-c
		for _, r := range result {
			results = append(results, r)
		}
	}

	fmt.Fprintf(w, "Search results for %s:\n%s", query, strings.Join(results[:], "\n"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
