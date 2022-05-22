package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Point struct {
	X int
	Y int
}

type Paper struct {
	Marked map[Point]bool
}

type Axis int

const (
	X Axis = iota
	Y
)

type Fold struct {
	Along Axis
	Value int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(r io.Reader) ([]Point, []Fold) {
	points := make([]Point, 0)
	folds := make([]Fold, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		check(err)
		points = append(points, Point{X: x, Y: y})
	}
	check(scanner.Err())
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var axis string
		var value int
		_, err := fmt.Sscanf(line, "fold along %1s=%d", &axis, &value)
		check(err)
		switch axis {
		case "x":
			folds = append(folds, Fold{Along: X, Value: value})
		case "y":
			folds = append(folds, Fold{Along: Y, Value: value})
		}
	}
	check(scanner.Err())
	return points, folds
}

func newPaper(points []Point) Paper {
	paper := Paper{
		Marked: make(map[Point]bool),
	}
	for _, p := range points {
		paper.Marked[p] = true
	}
	return paper
}

func foldPaper(orig Paper, fold Fold) Paper {
	folded := Paper{Marked: make(map[Point]bool)}
	switch fold.Along {
	case X:
		for p := range orig.Marked {
			if !orig.Marked[p] {
				continue
			}
			if p.X > fold.Value {
				foldedPoint := Point{
					X: fold.Value - (p.X - fold.Value),
					Y: p.Y,
				}
				folded.Marked[foldedPoint] = true
			} else {
				folded.Marked[p] = true
			}
		}
	case Y:
		for p := range orig.Marked {
			if !orig.Marked[p] {
				continue
			}
			if p.Y > fold.Value {
				foldedPoint := Point{
					X: p.X,
					Y: fold.Value - (p.Y - fold.Value),
				}
				folded.Marked[foldedPoint] = true
			} else {
				folded.Marked[p] = true
			}
		}
	}
	return folded
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	points, folds := readInput(f)
	//fmt.Printf("%+v\n%+v\n", points, folds)
	paper := newPaper(points)
	//fmt.Printf("%+v\n\n%+v\n", paper, folds)
	folded := foldPaper(paper, folds[0])
	//fmt.Printf("%+v\n", folded)
	fmt.Println(len(folded.Marked))
}
