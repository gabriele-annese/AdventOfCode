package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"example.com/httpInput"
)

func CalculateSimilarity(leftNumbers []int, rightNumbers []int) int32 {
	var similarity int32
	var founded int

	for _, lefet_num := range leftNumbers {
		founded = 0 // inizilizzo a 0 la variabile foudnded

		for _, right_num := range rightNumbers {
			if lefet_num == right_num {
				founded += 1
			}
		}
		similarity += int32(lefet_num * founded)
	}

	return similarity
}

// Metodo per calcolare la differenza tra input1 e input 2
func CalculateDistance(leftNumbers []int, rightNumbers []int) int32 {
	var distance int32

	for i, lefet_num := range leftNumbers {
		//Dato che hanno la stessa lunghezza posso utilizzare lo stesso indice
		right_num := rightNumbers[i]
		distance += int32(math.Abs(float64(lefet_num - right_num)))
	}

	return distance
}

// Metodo per mettere in ordine crescente una lista di interi
func Ascending(list_number *[]int) {
	sort.Slice(*list_number, func(i, j int) bool {
		return (*list_number)[i] < (*list_number)[j]
	})
}

// Metodo per creare due liste separate di numeri
func CreateInput(body string, leftNumbers *[]int, rightNumbers *[]int) {
	//Leggo riga per riga tramite lo scanner
	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		line := scanner.Text()
		//Separo le stringhe in due liste e le converto in int
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			leftNum, err := strconv.Atoi(parts[0])
			if err == nil {
				*leftNumbers = append(*leftNumbers, leftNum)
			} else {
				log.Fatal("Error converting string to int: %v\n", err)
			}
			rightNum, err := strconv.Atoi(parts[1])
			if err == nil {
				*rightNumbers = append(*rightNumbers, rightNum)
			} else {
				log.Fatal("Error converting string to int: %v\n", err)
			}
		}

	}
	// Controlla se c'è stato un errore nello scanner
	if err := scanner.Err(); err != nil {
		log.Fatal("Errore nella lettura dell'input:", err)
	}

}

func main() {
	fmt.Println(`
       *     *
      ***   ***
     ***** *****
    **************
   ***************
      |||  |||
   MERRY CHRISTMAS
    FROM GABRIELE 🎅
    `)
	fmt.Println("\n************PRIMA PARTE************")
	var leftNumbers, rightNumbers []int
	//Recupero l'input tramite http e lo converto in due liste separate di numeri
	body := httpInput.GetInput("2024", "1")
	CreateInput(body, &leftNumbers, &rightNumbers)

	//Metto in ordine crescente le due liste tramite slice
	Ascending(&leftNumbers)
	Ascending(&rightNumbers)

	//Calcolo la distanza
	distance := CalculateDistance(leftNumbers, rightNumbers)
	fmt.Println("Result: ", distance)

	fmt.Println("\n************SECONDA PARTE************")
	//Calcolo la simiralita
	similarity := CalculateSimilarity(leftNumbers, rightNumbers)
	fmt.Println("Result: ", similarity)

}
