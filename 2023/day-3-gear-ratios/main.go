package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1()
}

func part1() {
	schematic := [][]rune{}

	// open file
	f, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for j, letter := range line {
			row[j] = letter
		}
		schematic = append(schematic, row)
	}
	for _, line := range schematic {
		num_parts := make([]string, 0)
		num_parts = nil
		for _, character := range line {
			if character >= '1' && character <= '9' {
				num_parts = append(num_parts, string(character))
			} else if character == '.' && num_parts != nil{
				// Once the full number has been found, figure out if it should be counted
				fmt.Println(num_parts)
				num_parts = nil
			}
		}
		num_parts = nil
	}
}
