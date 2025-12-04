package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func maxJoltage1(line string) int {
	maxPrevDigit := -1
	maxValue := 0

	for _, ch := range line {
		d := int(ch - '0')

		if maxPrevDigit != 1 {
			value := maxPrevDigit*10 + d

			if value > maxValue {
				maxValue = value
			}
		}

		if d > maxPrevDigit {
			maxPrevDigit = d
		}
	}

	return maxValue
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		j := maxJoltage1(line)
		total += j
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("o/p joltage: %d\n", total)
}