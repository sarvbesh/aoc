package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rows := len(grid)
	cols := len(grid[0])

	startCol := -1 // col index at start
	for c := 0; c < cols; c++ {
		if grid[0][c] == 'S' {
			startCol = c
			break
		}
	}
	if startCol == -1 {
		log.Fatal("no s found")
	}

	beams := map[int]struct{}{startCol: {}} // current col index by beams

	var splits int64 = 0 // count total

	for r := 1; r < rows; r++ {
		next := make(map[int]struct{})

		for col := range beams {
			if col < 0 || col >= cols {
				continue
			}

			ch := grid[r][col]

			switch ch {
			case '.':
				next[col] = struct{}{}

			case '^':
				splits++
				if col-1 >= 0 {
					next[col-1] = struct{}{}
				}
				if col+1 < cols {
					next[col+1] = struct{}{}
				}

			default:
				next[col] = struct{}{}
			}
		}

		beams = next
		if len(beams) == 0 {
			break
		}
	}

	fmt.Printf("splits: %d\n", splits) // 1711
}