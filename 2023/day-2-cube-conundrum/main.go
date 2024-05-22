package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1()
}

func part1() {
	games := []map[string]int{}
	
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		prefix := "Game "
		pattern := fmt.Sprintf("%s(\\d+)", prefix)
		regex := regexp.MustCompile(pattern)
		match := regex.FindStringSubmatch(scanner.Text())
		games = append(games, make(map[string]int), make(map[string]int))
		num, _ := strconv.Atoi(match[1])
		games[1]["blue"] = num
	}


}
