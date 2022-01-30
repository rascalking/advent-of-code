package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	basins    [][]Point
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

func (hm *HeightMap) findBasins() {
	hm.basins = make([][]Point, 0, len(hm.lowPoints))
	for _, lowPoint := range hm.lowPoints {
		toVisit := make([]Point, 0, 1)
		basin := make([]Point, 0, 1)
		visited := make(map[Point]bool, 0)
		toVisit = append(toVisit, lowPoint)
		for {
			visiting := toVisit[0]
			toVisit = toVisit[1:]
			if visited[visiting] {
				if len(toVisit) == 0 {
					break
				} else {
					continue
				}
			} else {
				visited[visiting] = true
			}
			if hm.grid[visiting.x][visiting.y] < 9 {
				basin = append(basin, visiting)
				for _, adj := range hm.findAdjacent(visiting.x, visiting.y) {
					if !visited[adj] {
						toVisit = append(toVisit, adj)
					}
				}
			}
			if len(toVisit) == 0 {
				break
			}
		}
		fmt.Println(lowPoint, len(basin), basin)
		hm.basins = append(hm.basins, basin)
	}
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	hm := parseHeightMap(readLines(f))
	hm.findLowPoints()
	fmt.Printf("%v low points\n", len(hm.lowPoints))
	hm.findBasins()
	fmt.Printf("%v basins\n", len(hm.basins))
	sort.Slice(hm.basins, func(i, j int) bool {
		return len(hm.basins[i]) > len(hm.basins[j])
	})
	product := 1
	for i := 0; i < 3; i++ {
		fmt.Printf("basin %v size %v\n", i, len(hm.basins[i]))
		product *= len(hm.basins[i])
	}
	fmt.Println(product)
}
