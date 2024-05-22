package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
				numbers = append(numbers, int(letter - '0'))
			}
		}
		var str1 string = fmt.Sprint(numbers[0])
		var str2 string = fmt.Sprint(numbers[len(numbers) - 1])
		num, _ := strconv.Atoi(str1 + str2)
		final += num
		
	}

	fmt.Printf("Output: %d \n", final)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
