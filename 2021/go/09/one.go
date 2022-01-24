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

type Point struct {
	x int
	y int
}

type HeightMap struct {
	grid      [][]int
	lowPoints []Point
}

func parseHeightMap(lines []string) HeightMap {
	hm := HeightMap{
		grid: make([][]int, len(lines[0])),
	}
	for x := 0; x < len(lines[0]); x++ {
		hm.grid[x] = make([]int, len(lines))
	}
	for y := 0; y < len(lines); y++ {
		for x, nStr := range strings.Split(lines[y], "") {
			n, err := strconv.Atoi(nStr)
			check(err)
			hm.grid[x][y] = n
		}
	}
	return hm
}

func (hm *HeightMap) print() {
	for y := 0; y < len(hm.grid[0]); y++ {
		for x := 0; x < len(hm.grid); x++ {
		}
	}
}

func (hm *HeightMap) findLowPoints() {
	hm.lowPoints = make([]Point, 0)
	for x := 0; x < len(hm.grid); x++ {
		for y := 0; y < len(hm.grid[x]); y++ {
			isLowPoint := true
			for _, p := range hm.findAdjacent(x, y) {
				if hm.grid[p.x][p.y] <= hm.grid[x][y] {
					isLowPoint = false
					break
				}
			}
			if isLowPoint {
				hm.lowPoints = append(hm.lowPoints, Point{x, y})
			}
		}
	}
}

func (hm *HeightMap) findAdjacent(x int, y int) []Point {
	adj := make([]Point, 0, 4)
	if x > 0 {
		adj = append(adj, Point{x - 1, y})
	}
	if x < len(hm.grid)-1 {
		adj = append(adj, Point{x + 1, y})
	}
	if y > 0 {
		adj = append(adj, Point{x, y - 1})
	}
	if y < len(hm.grid[0])-1 {
		adj = append(adj, Point{x, y + 1})
	}
	return adj
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	hm := parseHeightMap(readLines(f))
	hm.findLowPoints()
	riskLevelSum := 0
	for _, lp := range hm.lowPoints {
		riskLevelSum += hm.grid[lp.x][lp.y] + 1
	}
	fmt.Println(riskLevelSum)
}
