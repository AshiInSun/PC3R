package main

import (
	"context"
	"fmt"
)

func Reductor(reductorch chan int, ctx context.Context) {
	var total int

	for {
		select {
		case i, ok := <-reductorch:
			if !ok {
				// Canal fermé : on sort proprement et affiche les résultats
				fmt.Println("Canal fermé, calcul du total...")
				goto CALCUL
			}
			total += i

		case <-ctx.Done():
			// Contexte annulé : on sort proprement et affiche les résultats
			fmt.Println("Contexte annulé, calcul du total...")
			goto CALCUL
		}
	}

CALCUL:
	// Calcul des jours, heures, minutes
	heures := total / 60
	jours := heures / 24
	heures = heures % 24
	minutes := total % 60
	fmt.Println("TOTAL :", jours, "jours", heures, "heures", minutes, "minutes")
}
