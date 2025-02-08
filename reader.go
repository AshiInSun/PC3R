package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reader(fPath string, ch chan string) {
	fmt.Println("Lecture du fichier:", fPath)
	file, err := os.Open(fPath)
	if err != nil {
		fmt.Println("Erreur ouverture fichier:", err)
		close(ch)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//on veut skip la première ligne
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		ch <- line
	}

	if err := scanner.Err(); err != nil { // Vérifie les erreurs de lecture
		fmt.Println("Erreur lecture:", err)
	}
	close(ch)
}
