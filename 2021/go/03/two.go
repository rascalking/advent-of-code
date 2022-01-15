package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(r io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	check(scanner.Err())
	return lines
}

func countBits(lines []string) []map[string]int {
	counts := make([]map[string]int, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		counts[i] = map[string]int{"0": 0, "1": 0}
	}
	for _, line := range lines {
		for i, bit := range strings.Split(line, "") {
			counts[i][bit] += 1
		}
	}
	return counts
}

func filter(lines []string, fn func(string, map[string]int) bool) string {
	result := "fuck"
	for i := 0; i < len(lines[0]); i++ {
		counts := countBits(lines)
		filtered := make([]string, 0)
		for _, line := range lines {
			if fn(strings.Split(line, "")[i], counts[i]) {
				filtered = append(filtered, line)
			}
		}
		if len(filtered) == 1 {
			result = filtered[0]
			break
		}
		lines = filtered
	}
	return result
}

func lineToInt(line string) int64 {
	i, err := strconv.ParseInt(line, 2, 0)
	check(err)
	return i
}

func oxygenFilter(bit string, counts map[string]int) bool {
	keep := "1"
	if counts["0"] > counts["1"] {
		keep = "0"
	}
	return bit == keep
}

func oxygen(lines []string) int64 {
	return lineToInt(filter(lines, oxygenFilter))
}

func co2Filter(bit string, counts map[string]int) bool {
	keep := "0"
	if counts["0"] > counts["1"] {
		keep = "1"
	}
	return bit == keep
}

func co2(lines []string) int64 {
	return lineToInt(filter(lines, co2Filter))
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	ox := oxygen(lines)
	co := co2(lines)

	fmt.Println(ox, co)
	fmt.Println(ox * co)
}
