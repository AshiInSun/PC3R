package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(serverch chan Paquet) {
	for paquet := range serverch {
		//convertissage des heures en int
		parts := strings.Split(paquet.heureArrivee, ":")
		if len(parts) == 1 {
			println("Erreur")
			continue
		}
		hA, err := strconv.Atoi(parts[0])
		mA, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Erreur convertissage heure arrivee", err)
		}
		tA := hA*60 + mA
		parts = strings.Split(paquet.heureDepart, ":")
		hD, err := strconv.Atoi(parts[0])
		mD, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Erreur convertissage heure depart", err)
		}
		tD := hD*60 + mD

		paquet.tempArret = tA - tD
		paquet.workerChan <- paquet
	}
}
