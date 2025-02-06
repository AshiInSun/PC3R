package main

import (
	"fmt"
	"sync"
)

// declaration varglobale
var nbWorkers = 3

// main
func main() {
	fmt.Println("Hello, World!")
	ch := make(chan string, nbWorkers)
	var wg sync.WaitGroup
	wg.Add(1)

	//reader
	go func() {
		Reader("stop_times.txt", ch)
		wg.Done() // Signale que Reader() est terminé
	}()

	//workers
	for i := 0; i < nbWorkers; i++ {
		wg.Add(1)
		go func() {
			Worker(ch)
			wg.Done() // Signale que Worker() est terminé
		}()
	}

	wg.Wait()
	println("Fin du programme")
}
