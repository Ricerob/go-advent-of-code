package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Only completed half of day one
// Learned about runes, strings, for loops, and small bits about os/ reading files
func main() {
	fmt.Println("Advent of Code 2023: Day 1")

	var runningSum int = 0
	lines := read_input()
	for i := 0; i < len(lines); i++ {
		runes := []rune(lines[i])
		start_idx := find_first_number_ascending(runes)
		end_idx := find_first_number_descending(runes)
		if start_idx == -1 || end_idx == -1 {
			fmt.Println("Found one without digits:", lines[i])
			continue
		} else {
			newDigit := string(runes[start_idx]) + string(runes[end_idx])
			digit, err := strconv.Atoi(newDigit)
			if err != nil {
				panic(err)
			}

			runningSum += digit
		}

	}
	fmt.Println(runningSum)
}

// Finds first numer, returns index location
func find_first_number_ascending(runes []rune) int {
	var index int = -1
	for i := 0; i < len(runes); i++ {
		if is_numeric(runes[i]) {
			index = i
			break
		}
	}
	return index
}

func find_first_number_descending(runes []rune) int {
	var index int = -1
	for i := len(runes) - 1; i >= 0; i-- {
		if is_numeric(runes[i]) {
			index = i
			break
		}
	}
	return index
}

func is_numeric(char rune) bool {
	return unicode.IsDigit(char)
}

func read_input() []string {
	var lines []string

	readFile, err := os.Open("input.txt")

	// Error handling
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(readFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
