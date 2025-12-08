/*
	calculates the number of cells marked with '@' that have fewer than 4 adjacent neighbors also marked with '@'

*/

package main

import (
	"bufio" // buffer io
	"fmt"
	"log"
	"os"
)

// represent 8 possible directions

var dir = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rows := len(grid)
	cols := len(grid[0])

	accessCount := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}

			adj := 0
			for _, d := range dir {
				nr, nc := r+d[0], c+d[1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					if grid[nr][nc] == '@' {
						adj++
					}
				}
			}

			if adj < 4 {
				accessCount++
			}
		}
	}

	fmt.Printf("accessible rolls of paper: %d\n", accessCount) // 1527
}