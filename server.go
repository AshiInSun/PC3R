package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func calculate(serverch chan Paquet, ctx context.Context) {
	for {
		select {
		case paquet, ok := <-serverch:
			if !ok {
				// Si le canal est fermé, on sort de la boucle
				fmt.Println("Canal serverch fermé, arrêt de calculate.")
				return
			}

			// Convertir les heures en minutes
			tA, err := convertTimeToMinutes(paquet.heureArrivee)
			if err != nil {
				fmt.Println("Erreur convertissage heure arrivée:", err)
				continue
			}

			tD, err := convertTimeToMinutes(paquet.heureDepart)
			if err != nil {
				fmt.Println("Erreur convertissage heure départ:", err)
				continue
			}
			time.Sleep(100 * time.Millisecond)
			// Calcul de la durée
			paquet.tempArret = tA - tD

			// Envoyer le paquet mis à jour à la goroutine Worker
			select {
			case paquet.workerChan <- paquet:
			case <-ctx.Done():
				fmt.Println("Contexte annulé, arrêt de calculate.")
				return
			default:
				fmt.Println("Worker occupé, c'est fini !")
			}

		case <-ctx.Done():
			// Si le contexte est annulé, on quitte proprement
			fmt.Println("Contexte annulé, arrêt de calculate.")
			return
		}
	}
}

func convertTimeToMinutes(timeStr string) (int, error) {
	parts := strings.Split(timeStr, ":")

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("erreur conversion heure: %s", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("erreur conversion minutes: %s", err)
	}

	return hours*60 + minutes, nil
}
