package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
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
	// Vérifier si un argument est fourni
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <nombre>")
		return
	}

	// Récupérer le premier argument
	arg, _ := strconv.Atoi(os.Args[1])
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(arg)*time.Second)
	defer cancel() // Annule le contexte à la fin

	fmt.Println("Hello, World!")
	ch := make(chan string, nbWorkers)
	serverch := make(chan Paquet, nbWorkers)
	reductorch := make(chan int, nbWorkers)
	var wg sync.WaitGroup
	var workerGroup sync.WaitGroup
	wg.Add(3) // 3 pour les 3 goroutines Reader, calculate et Reductor

	//reader
	go func() {
		defer wg.Done() // Signale que Reader() est terminé
		Reader("stop_times.txt", ch, ctx)
	}()
	//server
	go func() {
		defer wg.Done()
		calculate(serverch, ctx)
	}()
	//reductor
	go func() {
		defer wg.Done()
		Reductor(reductorch, ctx)
	}()
	//workers
	for i := 0; i < nbWorkers; i++ {
		workerGroup.Add(1)
		go func() {
			defer workerGroup.Done()
			Worker(ch, serverch, reductorch, ctx)
		}()
	}
	//attendre la fin des workers
	workerGroup.Wait()
	close(reductorch)
	close(serverch)
	wg.Wait()
	println("Fin du programme")
}
