package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"example.com/httpInput"
)

/*
Due condizioni da rispettare
- che i livelli siano crescenti o deceresenti
- se lo sono i numeri tra di loro posso differire di un minimo di 1 ad un massimo di 3 numeri
*/

func Check_adjacent(crescente bool, level_list []string) bool {
	var valid_report bool = false
	var temp_list []int
	for _, s := range level_list {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Errorf("Errore durante il parsing: %v \n", err)
		}
		temp_list = append(temp_list, number)
	}

	if crescente {
		fmt.Printf("Crescente\n")
		for i := 0; i < len(temp_list)-1; i++ {
			n := temp_list[i]
			difference := temp_list[i+1] - n
			//Calcolo la differenza tra i numeri, se il numero e' negativo in una lista crescente scarto il report. la differnza deve essere compresa tra 1 e 3
			if !(difference >= 1 && difference <= 3) || difference < 0 {
				valid_report = false
				break
			}
			valid_report = true
		}
	} else {
		fmt.Printf("Decrescente\n")
		for i := 0; i < len(temp_list)-1; i++ {
			n := temp_list[i]
			difference := temp_list[i+1] - n
			//Calcolo la differenza tra i numeri, se il numero e' positivo in una lista crescente scarto il report. la differnza deve essere compresa tra 1 e 3
			if !(math.Abs(float64(difference)) >= 1 && math.Abs(float64(difference)) <= 3) || difference > 0 {
				valid_report = false
				break
			}
			valid_report = true
		}
	}

	return valid_report

}

// Metodo che controlla se la lista e' crescente o decrescente
func Check_Decreasing_Increasing(report string, safe_report *int32) {
	level_list := strings.Split(report, " ")

	//Controllo se la lista deve essere decrescente o crescente
	//Se primo e secondo numero sono uguali il livello non e' valido
	first_numbers, err := strconv.Atoi(level_list[0])
	if err != nil {
		fmt.Errorf("Errore durante il parsing: %v \n", err)
	}
	second_number, err := strconv.Atoi(level_list[1])
	if err != nil {
		fmt.Errorf("Errore durante il parsing: %v \n", err)
	}

	if first_numbers < second_number {
		// La lista deve essere crescente
		if Check_adjacent(true, level_list) {
			*safe_report++
		}
	} else if first_numbers > second_number {
		// La lista deve essere decrescente
		if Check_adjacent(false, level_list) {
			*safe_report++
		}
	} else {
		fmt.Printf("Non è possibile stabilire se la lista è crescente o decrescente. Report Non valido \n")
	}
}

func fist_part(body string, safe_report *int32) {

	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		report := scanner.Text()
		fmt.Printf("The report: %s \n", report)
		Check_Decreasing_Increasing(report, *&safe_report)
	}
}

/*
	Colore	Codice ANSI
	Nero	\033[30m
	Rosso	\033[31m
	Verde	\033[32m
	Giallo	\033[33m
	Blu		\033[34m
	Magenta	\033[35m
	Ciano	\033[36m
	Bianco	\033[37m
*/

func main() {
	var safe_report int32 = 0
	httpInput.GetAscii()
	body := httpInput.GetInput("2024", "2")

	fmt.Println("\n\033[35m************PRIMA PARTE************\033[0m")
	fist_part(body, &safe_report)

	fmt.Printf("\033[32mReport validi: %d\033[0m", safe_report)

	fmt.Println("\n\n\033[35m************SECONDA PARTE************\033[0m")
}
