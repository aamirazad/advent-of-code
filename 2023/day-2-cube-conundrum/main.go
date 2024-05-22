package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	DICE_COUNT := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}

	// var games []map[string]int

	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	i := 0
	final := 0
	for scanner.Scan() {
		i++
		game_possible := true
		// Find the start of the first bag
		regex := regexp.MustCompile(`\w+.*?: `)
		remove := regex.FindStringSubmatch(scanner.Text())
		cut_line, _ := strings.CutPrefix(scanner.Text(), remove[0])

		// Take the formatted line and interpret each of the three rounds
		for _, round := range strings.Split(cut_line, "; ") {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				values := strings.Split(cube, " ")
				number, _ := strconv.Atoi(values[0])
				// if values[1] == "red" && values[0] > "12" {
				// 	game_possible = false
				// }
				for color, max := range DICE_COUNT {
					if values[1] == color && number > max {
						game_possible = false
					}
				}
			}
		}
		if game_possible {
			final += i
		}
	}
	fmt.Printf("Result: %d \n", final)

}
