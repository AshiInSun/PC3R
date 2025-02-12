package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

// Reader lit un fichier ligne par ligne et envoie les lignes sur le canal, tout en respectant le contexte
func Reader(fPath string, ch chan string, ctx context.Context) {
	fmt.Println("Lecture du fichier:", fPath)

	file, err := os.Open(fPath)
	if err != nil {
		fmt.Println("Erreur ouverture fichier:", err)
		close(ch)
		return
	}
	defer file.Close()
	defer close(ch) // Assurer la fermeture du canal à la fin

	scanner := bufio.NewScanner(file)

	// Skipper la première ligne
	if scanner.Scan() {
		fmt.Println("Première ligne ignorée.")
	}

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			fmt.Println("Contexte annulé, arrêt de la lecture du fichier.")
			close(ch)
			return
		default:
			ch <- scanner.Text()
		}
	}

	// Vérifier s'il y a une erreur de lecture
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lecture:", err)
	}
	close(ch)
}
