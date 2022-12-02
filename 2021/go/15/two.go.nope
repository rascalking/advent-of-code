package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/beefsack/go-astar"
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

type Cave struct {
	Grid    [][]*Position
	Average float64
}

func (c *Cave) PrintGrid() {
	for y := 0; y < len(c.Grid); y++ {
		for x := 0; x < len(c.Grid[0]); x++ {
			fmt.Printf("%d", int(c.Grid[y][x].Risk))
		}
		fmt.Printf("\n")
	}
}

type Position struct {
	Cave      *Cave
	X         int
	Y         int
	Risk      float64
	Neighbors []astar.Pather
}

func parseCave(lines []string) *Cave {
	cave := &Cave{}

	// fill in the initial grid
	grid := make([][]*Position, len(lines))
	for y, line := range lines {
		positions := make([]*Position, len(line))
		chars := strings.Split(line, "")
		for x, char := range chars {
			val, err := strconv.Atoi(char)
			check(err)
			pos := &Position{
				Cave: cave,
				X:    x,
				Y:    y,
				Risk: float64(val),
			}
			positions[x] = pos
		}
		grid[y] = positions
	}

	// generate the full grid
	ySize := len(grid)
	xSize := len(grid[0])
	fullGrid := make([][]*Position, ySize*5)
	for y := 0; y < len(fullGrid); y++ {
		fullGrid[y] = make([]*Position, xSize*5)
	}
	for gridY := 0; gridY < 5; gridY++ {
		for gridX := 0; gridX < 5; gridX++ {
			if (gridX == 0) && (gridY == 0) {
				// just copy over the original grid
				for y := 0; y < ySize; y++ {
					for x := 0; x < xSize; x++ {
						fullGrid[y][x] = grid[y][x]
					}
				}
			} else {
				// we need to generate new Positions for the other grids
				inc := gridX + gridY
				for y := 0; y < ySize; y++ {
					for x := 0; x < xSize; x++ {
						orig := grid[y][x]
						newRisk := float64((int(orig.Risk) + inc) % 10)
						pos := &Position{
							Cave: cave,
							X:    (gridX * xSize) + x,
							Y:    (gridY * ySize) + y,
							Risk: newRisk,
						}
						fullGrid[pos.Y][pos.X] = pos
					}
				}
			}
		}
	}
	cave.Grid = fullGrid

	// link up neighbors after the grid is full
	for y := 0; y < len(cave.Grid); y++ {
		for x := 0; x < len(cave.Grid[0]); x++ {
			pos := cave.Grid[y][x]
			neighbors := make([]astar.Pather, 0, 4)
			steps := [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
			for _, step := range steps {
				xN := pos.X + step[0]
				yN := pos.Y + step[1]
				if (xN >= 0) && (xN < len(cave.Grid[0])) &&
					(yN >= 0) && (yN < len(cave.Grid)) {
					neighbor := cave.Grid[yN][xN]
					neighbors = append(neighbors, neighbor)
				}
			}
			pos.Neighbors = neighbors
		}
	}

	// calculate the average risk for use in the estimated cost
	var sum float64
	for y := 0; y < len(cave.Grid); y++ {
		for x := 0; x < len(cave.Grid[y]); x++ {
			sum += cave.Grid[y][x].Risk
		}
	}
	cave.Average = sum / float64(len(cave.Grid)*len(cave.Grid[0]))

	return cave
}

func (p *Position) PathNeighbors() []astar.Pather {
	return p.Neighbors
}

func (p *Position) PathNeighborCost(to astar.Pather) float64 {
	toPos := to.(*Position)
	return toPos.Risk
}

func (p *Position) PathEstimatedCost(to astar.Pather) float64 {
	toPos := to.(*Position)
	distance := math.Abs(float64(toPos.X-p.X)) + math.Abs(float64(toPos.Y-p.Y))
	return p.Cave.Average * distance
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	cave := parseCave(lines)
	//cave.PrintGrid()
	//fmt.Printf("%+v\n", cave)
	start := cave.Grid[0][0]
	end := cave.Grid[len(cave.Grid)-1][len(cave.Grid[0])-1]
	//fmt.Printf("%+v, %+v\n", start, end)
	_, distance, found := astar.Path(start, end)
	if !found {
		panic("no path found")
	}
	fmt.Println(distance)
}
