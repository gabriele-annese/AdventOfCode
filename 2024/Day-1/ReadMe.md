# Intro
Il codice recupera tamite http request l'input restutiuto dall'endpoit `https://adventofcode.com/2024/day/1/input`. 
Per poter ottenere correttamente l'input, è necessario fornire il `Session ID` della propria sessione. Questo perché l'input restituito dall'endpoint è personalizzato per ciascun utente loggato, quindi ogni sessione avrà un contenuto diverso.

## Recuperare il Session ID
Il `Session ID` è un identificativo unico che viene generato quando un utente accede al sito di Advent of Code. È necessario utilizzare questo ID per autenticarsi e ottenere i dati relativi alla propria sessione.


### 1. Accedere al proprio account su Advent of Code
- [Vai su Advent of Code 2024](https://adventofcode.com/2024).
- Accedi al tuo account se non lo hai già fatto.

### 2. Recuperare il Session ID dal browser
Una volta loggato, apri gli Strumenti di Sviluppo del browser (in genere si può fare clic destro sulla pagina e selezionare "Ispeziona" o usare il tasto F12).
Vai alla sezione Storage (di solito sotto la voce "Application" nei browser come Chrome).
Seleziona Cookies e cerca il cookie denominato session.
Copia il valore di questo cookie. Questo è il Session ID che dovrai utilizzare per fare richieste all'endpoint.

> [!WARNING]  
>**Non condividere mai il Session ID con altri, poiché consente di accedere ai dati del tuo account.**
Per motivi di sicurezza, è consigliato conservare il Session ID in variabili d'ambiente o in file di configurazione che non vengano mai pubblicati o condivisi (ad esempio, aggiungendo il file di configurazione a .gitignore se usi Git).


