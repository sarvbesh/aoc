package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dirs = [][2]int{
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

	grid := [][]byte{}
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rows := len(grid)
	cols := len(grid[0])

	adj := make([][]int, rows)
	for i := range adj {
		adj[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}

			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
					adj[r][c]++
				}
			}
		}
	}

	type cell struct{ r, c int }
	queue := []cell{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '@' && adj[r][c] < 4 {
				queue = append(queue, cell{r, c})
			}
		}
	}

	removed := 0

	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		r, c := cur.r, cur.c

		if grid[r][c] != '@' {
			continue
		}

		grid[r][c] = '.'
		removed++

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
				adj[nr][nc]--
				if adj[nr][nc] < 4 {
					queue = append(queue, cell{nr, nc})
				}
			}
		}
	}

	fmt.Printf("rolls of paper removed: %d\n", removed)
}