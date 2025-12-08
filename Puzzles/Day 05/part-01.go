/*
reads two distinct sets of data from input.txt

1. list of inclusive integer ranges (intervals)
2. list of individual integer IDs

goal is to merge the overlapping or contiguous ranges

then count how many of the individual IDs fall inside these merged intervals
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

type Interval1 struct {
	L, R int64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := []Interval1{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		a, _ := strconv.ParseInt(parts[0], 10, 64)
		b, _ := strconv.ParseInt(parts[1], 10, 64)
		ranges = append(ranges, Interval1{a, b})
	}

	sort.Slice(ranges, func(i, j int) bool { return ranges[i].L < ranges[j].L })

	merged := []Interval1{}
	for _, x := range ranges {
		if len(merged) == 0 || x.L > merged[len(merged)-1].R+1 {
			merged = append(merged, x)
		} else {
			if x.R > merged[len(merged)-1].R {
				merged[len(merged)-1].R = x.R
			}
		}
	}

	freshCount := 0

	for scanner.Scan() {
		idStr := strings.TrimSpace(scanner.Text())
		if idStr == "" {
			continue
		}
		id, _ := strconv.ParseInt(idStr, 10, 64)

		i := sort.Search(len(merged), func(i int) bool { return merged[i].L > id })
		if i > 0 && merged[i-1].L <= id && id <= merged[i-1].R {
			freshCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fresh ingredient IDs: %d\n", freshCount) // 617
}