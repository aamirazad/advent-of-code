package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var final int

	for scanner.Scan() {
		// Extract the numbers into a list
		numbers := []int{}
		for _, letter := range scanner.Text() {
			if letter >= '0' && letter <= '9' {
				numbers = append(numbers, int(letter-'0'))
			}
		}
		var str1 string = fmt.Sprint(numbers[0])
		var str2 string = fmt.Sprint(numbers[len(numbers)-1])
		num, _ := strconv.Atoi(str1 + str2)
		final += num

	}

	fmt.Printf("Output: %d \n", final)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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

	var final int

	wordToNum := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	for scanner.Scan() {
		// Extract the numbers into a list
		numbers := []int{}
		var word string
		for _, letter := range scanner.Text() {
			if letter >= '0' && letter <= '9' {
				numbers = append(numbers, int(letter-'0'))
				word = ""
			} else {
				word += string(letter)
				for i, j := range wordToNum {
					if strings.Contains(word, i) {
						numbers = append(numbers, j)
						word = string(word[len(word)-1])
					}
				}
			}
		}
		var str1 string = fmt.Sprint(numbers[0])
		var str2 string = fmt.Sprint(numbers[len(numbers)-1])
		num, _ := strconv.Atoi(str1 + str2)
		final += num
	}
	fmt.Printf("Output: %d \n", final)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
