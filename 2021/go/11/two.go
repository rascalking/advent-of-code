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

type Grid struct {
	levels     [10][10]int
	flashed    [10][10]bool
	flashCount int
}

func (g *Grid) Print() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Printf("%d", g.levels[x][y])
		}
		fmt.Printf("\n")
	}
}

func (g *Grid) Step() int {
	flashes := make([][]int, 0)

	// start with the basic increment, note points that should flash
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			g.levels[x][y]++
			if g.levels[x][y] > 9 {
				flashes = append(flashes, []int{x, y})
			}
		}
	}

	for {
		// stop if there's no more flash candidates
		if len(flashes) == 0 {
			break
		}

		// pop a candidate
		point := flashes[len(flashes)-1]
		flashes = flashes[:len(flashes)-1]
		x := point[0]
		y := point[1]

		// skip if we've already flashed it
		if g.flashed[x][y] {
			continue
		}

		// mark it as flashed
		g.flashed[x][y] = true

		// increment the neighbors
		for x1 := x - 1; x1 <= x+1; x1++ {
			if x1 < 0 || x1 > 9 {
				continue
			}
			for y1 := y - 1; y1 <= y+1; y1++ {
				if y1 < 0 || y1 > 9 {
					continue
				}
				g.levels[x1][y1]++
				if g.levels[x1][y1] > 9 {
					flashes = append(flashes, []int{x1, y1})
				}
			}
		}
	}

	// reset the flashed octopi
	stepFlashes := 0
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if g.flashed[x][y] {
				g.flashCount++
				stepFlashes++
			}
			g.flashed[x][y] = false
			if g.levels[x][y] > 9 {
				g.levels[x][y] = 0
			}
		}
	}
	return stepFlashes
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	grid := Grid{}
	for y, line := range readLines(f) {
		for x, c := range strings.Split(line, "") {
			n, err := strconv.Atoi(c)
			check(err)
			grid.levels[x][y] = n
		}
	}
	for i := 0; i < 10000; i++ {
		numFlashes := grid.Step()
		if numFlashes == 100 {
			fmt.Println(i + 1)
			break
		}
	}
}
