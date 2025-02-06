package main

import (
	"fmt"
	"strings"
)

func Worker(ch chan string) {
	for line := range ch {
		// Traitement de la ligne : on recupere l'heure d'arrivée
		parts := strings.Split(line, ",")
		arrivalTime := parts[1]
		fmt.Println(arrivalTime)
	}
}
