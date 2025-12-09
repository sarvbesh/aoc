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

	isBlankCol := func(c int) bool {
		for r := 0; r < rows; r++ {
			if grid[r][c] != ' ' {
				return false
			}
		}
		return true
	}

	var total int64 = 0
	c := 0

	for c < cols {
		if isBlankCol(c) {
			c++
			continue
		}

		start := c
		for c < cols && !isBlankCol(c) {
			c++
		}
		end := c

		rowsText := make([]string, rows)
		for r := 0; r < rows; r++ {
			rowsText[r] = strings.TrimSpace(string(grid[r][start:end]))
		}

		op := rowsText[rows-1]
		if op != "+" && op != "*" {
			log.Fatalf("invalid operator: %q", op)
		}

		nums := []int64{}
		for r := 0; r < rows-1; r++ {
			if rowsText[r] == "" {
				continue
			}
			n, err := strconv.ParseInt(rowsText[r], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}

		if len(nums) == 0 {
			continue
		}

		res := nums[0]
		for i := 1; i < len(nums); i++ {
			if op == "+" {
				res += nums[i]
			} else {
				res *= nums[i]
			}
		}

		total += res
	}

	fmt.Printf("total: %d\n", total)
}