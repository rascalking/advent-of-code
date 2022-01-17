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

func emptyFish() map[int]int {
	return map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
}

func parseFish(line string) map[int]int {
	fish := emptyFish()
	for _, numStr := range strings.Split(line, ",") {
		n, err := strconv.Atoi(numStr)
		check(err)
		fish[n] += 1
	}
	return fish
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	fish := parseFish(lines[0])
	for day := 0; day < 256; day++ {
		newFish := emptyFish()
		for i := 1; i <= 8; i++ {
			newFish[i-1] = fish[i]
		}
		newFish[6] += fish[0]
		newFish[8] = fish[0]
		fish = newFish
	}
	totalFish := 0
	for _, count := range fish {
		totalFish += count
	}
	fmt.Println(totalFish)
}
