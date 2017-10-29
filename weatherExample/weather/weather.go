package main

import (
	"fmt"
	"math/rand"
	"time"
)

type searchEngine interface {
	getResults(query string) []string
}

type temperatureProvider struct{}

type humidityProvider struct{}

type rainingChanceProvider struct{}

/* Devuelve un porcentaje de humedad random entre 1 y 100, tarda entre 0 y 4 segundos */
func (h humidityProvider) getHumidity() int {
	// Inicializo el random con una semilla basada en el tiempo
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	start := time.Now()

	time.Sleep(time.Duration(r.Intn(4)) * time.Second)
	hum := r.Intn(99) + 1

	// Se imprime el tiempo de demora
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Humedad - tiempo de respuesta: %s", elapsed))

	return hum
}

/* Devuelve un porcentaje de probabilidad de lluvia random entre 1 y 100, tarda entre 0 y 4 segundos */
func (rcp rainingChanceProvider) getRainingChance() int {
	// Inicializo el random con una semilla basada en el tiempo
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	start := time.Now()

	time.Sleep(time.Duration(r.Intn(4)) * time.Second)
	rc := r.Intn(99) + 1

	// Se imprime el tiempo de demora
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Probabilidad de lluvia - tiempo de respuesta: %s", elapsed))

	return rc
}

/* Devuelve una temperatura random entre 10 y 20 grados, tarda entre 0 y 4 segundos */
func (t temperatureProvider) getTemperature() int {
	// Inicializo el random con una semilla basada en el tiempo
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	start := time.Now()

	time.Sleep(time.Duration(r.Intn(4)) * time.Second)
	temp := r.Intn(10) + 10

	// Se imprime el tiempo de demora
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Temperatura - tiempo de respuesta: %s", elapsed))
	return temp
}
