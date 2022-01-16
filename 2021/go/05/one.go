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
		fmt.Printf("Considering (%d, %d) -> (%d, %d)\n", line.p1.x, line.p1.y, line.p2.x, line.p2.y)
		if line.p1.x == line.p2.x {
			if line.p1.y < line.p2.y {
				for y := line.p1.y; y < line.p2.y+1; y++ {
					region.mark(line.p1.x, y)
				}
			} else {
				for y := line.p2.y; y < line.p1.y+1; y++ {
					region.mark(line.p1.x, y)
				}
			}
		} else {
			if line.p1.x < line.p2.x {
				for x := line.p1.x; x < line.p2.x+1; x++ {
					region.mark(x, line.p1.y)
				}
			} else {
				for x := line.p2.x; x < line.p1.x+1; x++ {
					region.mark(x, line.p1.y)
				}
			}
		}
	}

	return region
}

func (r *Region) mark(x int, y int) {
	r.grid[x][y] += 1
	fmt.Printf("Marked (%d, %d), new value %d\n", x, y, r.grid[x][y])
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
		l := parseLine(line)
		if (l.p1.x == l.p2.x) || (l.p1.y == l.p2.y) {
			lines = append(lines, l)
		}
	}
	region := newRegion(lines)
	fmt.Println(region.countIntersections())
}
