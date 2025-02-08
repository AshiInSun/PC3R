package main

import (
	"fmt"
	"sync"
)

// declaration varglobale
var nbWorkers = 3

type Paquet struct {
	heureDepart  string
	heureArrivee string
	tempArret    int
	workerChan   chan Paquet
}

// main
func main() {
	fmt.Println("Hello, World!")
	ch := make(chan string, nbWorkers)
	serverch := make(chan Paquet, nbWorkers)
	reductorch := make(chan int, nbWorkers)
	var wg sync.WaitGroup
	var workerGroup sync.WaitGroup
	wg.Add(3) // 3 pour les 3 goroutines Reader, calculate et Reductor

	//reader
	go func() {
		Reader("stop_times.txt", ch)
		wg.Done() // Signale que Reader() est terminé
	}()
	//server
	go func() {
		calculate(serverch)
		wg.Done() // Signale que calculate() est terminé
	}()
	//reductor
	go func() {
		Reductor(reductorch)
		wg.Done() // Signale que Reductor() est terminé
	}()
	//workers
	for i := 0; i < nbWorkers; i++ {
		workerGroup.Add(1)
		go func() {
			Worker(ch, serverch, reductorch)
			workerGroup.Done() // Signale que Worker() est terminé
		}()
	}
	//attendre la fin des workers
	workerGroup.Wait()
	close(reductorch)
	close(serverch)
	wg.Wait()
	println("Fin du programme")
}
