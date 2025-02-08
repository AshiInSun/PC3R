package main

import (
	"fmt"
	"strings"
)

func Worker(readerch chan string, serverch chan Paquet, reductorch chan int) {
	thisch := make(chan Paquet)
	for line := range readerch {
		parts := strings.Split(line, ",")
		// Traitement de la ligne : on creer des paquets avec les données
		p := Paquet{
			heureDepart:  parts[1],
			heureArrivee: parts[2],
			tempArret:    0,
			workerChan:   thisch,
		}
		//On envoie le paquet au serveur
		serverch <- p
		//On attend la réponse du serveur
		paquet := <-thisch
		reductorch <- paquet.tempArret
		fmt.Println(paquet.tempArret)
	}
	close(thisch)
}
