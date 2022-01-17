package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

type Crabs struct {
	min   int
	max   int
	crabs map[int]int
}

func parseCrabs(line string) Crabs {
	tmp := make([]int, 0)
	min := math.MaxInt
	max := math.MinInt
	for _, numStr := range strings.Split(line, ",") {
		n, err := strconv.Atoi(numStr)
		check(err)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		tmp = append(tmp, n)
	}
	crabs := Crabs{min: min, max: max, crabs: make(map[int]int, 0)}
	for _, n := range tmp {
		crabs.crabs[n] += 1
	}
	return crabs
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	crabs := parseCrabs(lines[0])

	minFuel := math.MaxInt
	for i := crabs.min; i < crabs.max+1; i++ {
		fuel := 0
		for pos, count := range crabs.crabs {
			cost := pos - i
			if cost < 0 {
				cost = -cost
			}
			fuel += count * cost
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Println(minFuel)
}
