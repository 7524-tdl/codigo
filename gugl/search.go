package main

import (
	"fmt"
	"math/rand"
	"time"
)

type searchEngine interface {
	getResults(query string) []string
}

type linksEngine struct{}

type imagesEngine struct{}

type videosEngine struct{}

func (e linksEngine) getResults(query string) []string {
	fmt.Println(time.Now(), "Fetching links for:", query)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	return []string{
		"https://es.wikipedia.org/wiki/" + query,
	}
}

func (e imagesEngine) getResults(query string) []string {
	fmt.Println(time.Now(), "Fetching images for:", query)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return []string{
		"http://www.picsearch.com/index.cgi?q=" + query,
		"https://www.flickr.com/search/?text=" + query,
	}
}

func (e videosEngine) getResults(query string) []string {
	fmt.Println(time.Now(), "Fetching videos for:", query)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	return []string{
		"https://vimeo.com/search?q=" + query,
	}
}
