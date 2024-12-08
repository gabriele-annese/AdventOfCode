package httpInput

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	SessionID string `json:"session"`
}

// Metodo per recuperare l'input dell'advent of code tramite http
func GetInput() string {

	url := "https://adventofcode.com/2024/day/1/input"
	config := getSessionsID()

	// Creazione della reuquest
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Errore nella creazione della request: %v\n", err)
	}

	// Aggiunta del cookie di sessione
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: config.SessionID, // Inserisci qui il tuo cookie di sessione
	})

	//Client HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Errore durante l'invio della request: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Errore nella risposta: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Errore durante la lettura della response: %v", err)
	}

	return string(body)

}

func getSessionsID() Config {
	//Leggi file di configurazione
	data, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal("Errore durante la lettura del file di configurazione: %v", err)
	}

	//Parso il Json nella struttra config
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal("Errore durante il parsing del file di configurazione: %v", err)
	}

	return config

}
