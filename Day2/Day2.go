package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Determine which games would have been possible if the bag had been loaded with only
					 12 red cubes, 13 green cubes, and 14 blue cubes.
What is the sum of the IDs of those games?
*/

type cube struct {
	color    string
	quantity int
}

type Games struct {
	ID    int
	cubes []*cube
}

// Take ID splitting where line = Game
func GetGamesID(spliline string) int {

	game := strings.Split(spliline, ":")

	i, err := strconv.Atoi(strings.TrimSpace(strings.Split(game[0], "Game")[1]))
	if err != nil {
		panic(err)
	}
	return i

}

func main() {

	// Variables
	/*
		SumID := 0
		ListElfCubes := []cubes{
			{
				color:    "red",
				quantity: 12,
			},
			{
				color:    "green",
				quantity: 13,
			},
			{
				color:    "blue",
				quantity: 14,
			},
		}

	*/
	//ListGames := []games{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ";")

		// for each split create a new game
		game := new(Games)
		for _, spliline := range splits {

			if strings.Contains(spliline, "Game") {
				game.ID = GetGamesID(spliline)
			}
			// TODO add property in game structured
		}
		fmt.Println(game.ID)
		break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
