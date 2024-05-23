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
	part2()
}

func part1() {
	DICE_COUNT := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}

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

func part2() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	final := 0
	for scanner.Scan() {
		// Find the start of the first bag
		regex := regexp.MustCompile(`\w+.*?: `)
		remove := regex.FindStringSubmatch(scanner.Text())
		cut_line, _ := strings.CutPrefix(scanner.Text(), remove[0])

		// Take the formatted line and interpret each of the three rounds
		max_cubes := map[string]int{
			"red": 0, "green": 0, "blue": 0,
		}
		for _, round := range strings.Split(cut_line, "; ") {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				values := strings.Split(cube, " ")
				number, _ := strconv.Atoi(values[0])
				if number > max_cubes[string(values[1])] {
					max_cubes[string(values[1])] = number
				}
			}
		}
		final += (max_cubes["red"] * max_cubes["green"] * max_cubes["blue"])
	}
	fmt.Printf("Output: %d \n", final)
}
