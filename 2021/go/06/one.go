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

func parseFish(line string) []int {
	fish := make([]int, 0)
	for _, numStr := range strings.Split(line, ",") {
		n, err := strconv.Atoi(numStr)
		check(err)
		fish = append(fish, n)
	}
	return fish
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	fishes := parseFish(lines[0])
	fmt.Println(fishes)
	for day := 0; day < 80; day++ {
		todayFishLen := len(fishes)
		for i := 0; i < todayFishLen; i++ {
			fish := fishes[i] - 1
			if fish < 0 {
				fish = 6
				fishes = append(fishes, 8)
			}
			fishes[i] = fish
		}
	}
	fmt.Println(len(fishes))
}
