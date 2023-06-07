package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var currentElfCalories, maxElfCalories int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Blank line indicates a new elf's inventory
			if currentElfCalories > maxElfCalories {
				maxElfCalories = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error parsing input:", err)
				return
			}
			currentElfCalories += calories
		}
	}
	// Check if the last elf had the most calories
	if currentElfCalories > maxElfCalories {
		maxElfCalories = currentElfCalories
	}
	fmt.Println("Elf carrying the most calories:", maxElfCalories)
}
