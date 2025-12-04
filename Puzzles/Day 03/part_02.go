package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func maxJoltage2(line string) string {
	k := 12
	n := len(line)
	result := make([]byte, 0, k)

	start := 0

	for picks := 0; picks < k; picks++ {
		remainingPicks := k - picks
		end := n - remainingPicks

		bestDigit := byte('0' - 1)
		bestIndex := start

		for i := start; i <= end; i++ {
			if line[i] > bestDigit {
				bestDigit = line[i]
				bestIndex = i
			}

			if bestDigit == '9' {
				break
			}
		}

		result = append(result, bestDigit)
		start = bestIndex + 1
	}

	return string(result)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		j := maxJoltage2(line)

		var val uint64 = 0

		for i := 0; i < len(j); i++ {
			val = val*10 + uint64(j[i]-'0')
		}

		total += val
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("o/p joltage: %d\n", total)
}