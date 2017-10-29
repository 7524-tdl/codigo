package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type providerDirectory struct {
	temperatura        string
	humedad            string
	probabilidadLluvia string
}

func main() {
	start := time.Now()

	baseURL := "http://localhost:8082/"

	directory := providerDirectory{
		temperatura:        baseURL + "temp",
		humedad:            baseURL + "hum",
		probabilidadLluvia: baseURL + "rc",
	}

	tchan := make(chan string)
	hchan := make(chan string)
	pchan := make(chan string)

	go getData(directory.temperatura, tchan)
	go getData(directory.humedad, hchan)
	go getData(directory.probabilidadLluvia, pchan)

	reporteFinal := construirReporte(tchan, hchan, pchan)
	fmt.Println(reporteFinal)
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Tiempo para elaborar el informe: %s", elapsed))
}

func construirReporte(tchan chan string, hchan chan string, pchan chan string) string {
	reporteFinal := <-tchan
	reporteFinal += <-hchan
	reporteFinal += <-pchan
	return reporteFinal
}

func getData(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		c <- link + ": Sin datos"
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	c <- bodyString
}
