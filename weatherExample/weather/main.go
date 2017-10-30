package main

import (
	"fmt"
	"net/http"
)

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	tp := temperatureProvider{}
	temp := tp.getTemperature()
	fmt.Fprintf(w, "Temperatura: %d ºC\n", temp)
}

func humidityHandler(w http.ResponseWriter, r *http.Request) {
	hp := humidityProvider{}
	hum := hp.getHumidity()
	fmt.Fprintf(w, "Humedad: %d%% \n", hum)
}

func rainingChanceHandler(w http.ResponseWriter, r *http.Request) {
	rp := rainingChanceProvider{}
	rc := rp.getRainingChance()
	fmt.Fprintf(w, "Probabilidad de lluvia: %d%% \n", rc)
}

func main() {
	http.HandleFunc("/temp", temperatureHandler)
	http.HandleFunc("/hum", humidityHandler)
	http.HandleFunc("/rc", rainingChanceHandler)
	http.ListenAndServe(":8082", nil)
}
