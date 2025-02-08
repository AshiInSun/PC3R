package main

import "fmt"

func Reductor(reductorch chan int) {
	var total int
	for i := range reductorch {
		total += i
	}
	heures := total / 60
	jours := heures / 24
	heures = heures % 24
	minutes := total % 60
	fmt.Println("TOTAL :", jours, "jours", heures, "heures", minutes, "minutes")
}
