package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func main() {
	var input string
	var sum int
	fi, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = string(fi)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := string(scanner.Text())
		lenStr := len(line)

		for i := 0; i < lenStr; i++ {
			//Check if char is a number
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				var concatNumbers string
				concatNumbers = string(line[i])
				//if found the number reverse my string
				reverseString := reverse(string(line))

				for r := 0; r < lenStr; r++ {
					if _, err := strconv.Atoi(string(reverseString[r])); err == nil {

						concatNumbers += string(reverseString[r])
						fmt.Printf("%v this is number \n", concatNumbers)

						j, err := strconv.Atoi(concatNumbers)
						if err != nil {
							fmt.Printf("error occurred: %v\n", err)
						}
						sum += j
						//break if i found second number
						break
					}
				}
				//break if i found second number
				break
			}

		}
	}

	fmt.Printf("The sum is %v \n", sum)

}
