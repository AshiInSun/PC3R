package main

import (
	"context"
	"fmt"
	"strings"
)

func Worker(readerch chan string, serverch chan Paquet, reductorch chan int, ctx context.Context) {
	thisch := make(chan Paquet)
	defer close(thisch) // Fermer le canal à la fin

	for {
		select {
		case line, ok := <-readerch:
			if !ok {
				// Si readerch est fermé, on sort proprement
				fmt.Println("Canal readerch fermé, arrêt du Worker.")
				return
			}

			// Traitement de la ligne
			parts := strings.Split(line, ",")
			p := Paquet{
				heureDepart:  parts[1],
				heureArrivee: parts[2],
				tempArret:    0,
				workerChan:   thisch,
			}

			// Envoi au serveur
			serverch <- p

			// Attente de la réponse du serveur
			select {
			case paquet := <-thisch:
				reductorch <- paquet.tempArret
				fmt.Println(paquet.tempArret)
			case <-ctx.Done():
				// Si le contexte est annulé, on sort de la boucle
				fmt.Println("Contexte annulé, arrêt du Worker.")
				return
			}

		case <-ctx.Done():
			// Si le contexte est annulé, on arrête le Worker
			fmt.Println("Contexte annulé, arrêt du Worker.")
			return
		}
	}
}
