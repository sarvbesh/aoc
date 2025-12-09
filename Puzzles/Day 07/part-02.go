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
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(grid) == 0 {
		log.Fatal("empty grid")
	}

	rows := len(grid)
	cols := len(grid[0])

	startRow, startCol := -1, -1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				startRow, startCol = r, c
				break
			}
		}
		if startRow != -1 {
			break
		}
	}
	if startRow == -1 {
		log.Fatal("no S found in grid")
	}

	beams := map[int]int64{
		startCol: 1,
	}

	var totalTimelines int64 = 0

	for r := startRow + 1; r < rows; r++ {
		next := make(map[int]int64)

		for col, cnt := range beams {
			if cnt == 0 || col < 0 || col >= cols {
				continue
			}

			ch := grid[r][col]

			switch ch {
			case '.', 'S':
				if r == rows-1 {
					totalTimelines += cnt
				} else {
					next[col] += cnt
				}

			case '^':
				left := col - 1
				if left < 0 || left >= cols {
					totalTimelines += cnt
				} else {
					if r == rows-1 {
						totalTimelines += cnt
					} else {
						next[left] += cnt
					}
				}

				right := col + 1
				if right < 0 || right >= cols {
					totalTimelines += cnt
				} else {
					if r == rows-1 {
						totalTimelines += cnt
					} else {
						next[right] += cnt
					}
				}

			default:
				if r == rows-1 {
					totalTimelines += cnt
				} else {
					next[col] += cnt
				}
			}
		}

		beams = next
		if len(beams) == 0 {
			break
		}
	}

	fmt.Printf("timelines: %d\n", totalTimelines) // 36706966158365
}