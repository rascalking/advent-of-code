package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func parseLine(l string) Line {
	var line Line
	fmt.Sscanf(l, "%d,%d -> %d,%d", &line.p1.x, &line.p1.y, &line.p2.x, &line.p2.y)
	return line
}

type Region struct {
	max  Point
	grid [][]int
}

func newRegion(lines []Line) Region {
	region := Region{}

	for _, line := range lines {
		if line.p1.x > region.max.x {
			region.max.x = line.p1.x
		}
		if line.p2.x > region.max.x {
			region.max.x = line.p2.x
		}
		if line.p1.y > region.max.y {
			region.max.y = line.p1.y
		}
		if line.p2.y > region.max.y {
			region.max.y = line.p2.y
		}
	}
	region.grid = make([][]int, region.max.x+1)
	for x := range region.grid {
		region.grid[x] = make([]int, region.max.y+1)
	}

	for _, line := range lines {
		var step [2]int
		if line.p1.x == line.p2.x {
			if line.p1.y < line.p2.y {
				step = [2]int{0, 1}
			} else {
				step = [2]int{0, -1}
			}
		} else if line.p1.y == line.p2.y {
			if line.p1.x < line.p2.x {
				step = [2]int{1, 0}
			} else {
				step = [2]int{-1, 0}
			}
		} else if line.p1.x < line.p2.x {
			if line.p1.y < line.p2.y {
				step = [2]int{1, 1}
			} else {
				step = [2]int{1, -1}
			}
		} else {
			if line.p1.y < line.p2.y {
				step = [2]int{-1, 1}
			} else {
				step = [2]int{-1, -1}
			}
		}

		pos := line.p1
		for {
			region.grid[pos.x][pos.y] += 1
			if pos == line.p2 {
				break
			}
			pos.x += step[0]
			pos.y += step[1]
		}
	}

	return region
}

func (r *Region) countIntersections() int {
	count := 0
	for x := 0; x < r.max.x+1; x++ {
		for y := 0; y < r.max.y+1; y++ {
			if r.grid[x][y] > 1 {
				count += 1
			}
		}
	}
	return count
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	var lines []Line
	for _, line := range readLines(f) {
		lines = append(lines, parseLine(line))
	}
	region := newRegion(lines)
	fmt.Println(region.countIntersections())
}
