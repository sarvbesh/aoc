/*
	total number of integers covered by a list of input ranges after those ranges have been sorted and merged

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval2 struct {
	L, R int64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := []Interval2{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		a, _ := strconv.ParseInt(parts[0], 10, 64)
		b, _ := strconv.ParseInt(parts[1], 10, 64)
		ranges = append(ranges, Interval2{a, b})
	}

	sort.Slice(ranges, func(i, j int) bool { return ranges[i].L < ranges[j].L })

	merged := []Interval2{}
	for _, x := range ranges {
		if len(merged) == 0 || x.L > merged[len(merged)-1].R+1 {
			merged = append(merged, x)
		} else {
			if x.R > merged[len(merged)-1].R {
				merged[len(merged)-1].R = x.R
			}
		}
	}

	var total int64 = 0
	for _, iv := range merged {
		total += (iv.R - iv.L + 1)
	}

	fmt.Printf("fresh ingredient IDs: %d\n", total) // 338258295736104
}