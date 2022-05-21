package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

const START = "start"
const END = "end"

type Graph struct {
	Edges map[string][]string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func containsItem(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func containsPath(slice [][]string, item []string) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}
	return false
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

func addPath(paths [][]string, path []string) [][]string {
	if !containsPath(paths, path) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		paths = append(paths, tmp)
	}
	return paths
}

func (g *Graph) FindPaths(node string, path []string) [][]string {
	paths := make([][]string, 0)
	for _, n := range g.Edges[node] {
		n_path := append(path, n)
		if n == END {
			paths = addPath(paths, n_path)
		} else if (n == strings.ToLower(n)) && containsItem(path, n) {
			continue
		} else {
			recursed := g.FindPaths(n, n_path)
			for _, p := range recursed {
				paths = addPath(paths, p)
			}
		}
	}
	return paths
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	graph := Graph{Edges: make(map[string][]string)}
	for _, line := range lines {
		nodes := strings.SplitN(strings.TrimSpace(line), "-", 2)
		for i, node := range nodes {
			other := nodes[(i+1)%2]
			nexts, ok := graph.Edges[node]
			if !ok {
				nexts = make([]string, 0)
			}
			if !containsItem(nexts, other) {
				nexts = append(nexts, other)
				graph.Edges[node] = nexts
			}
		}
	}
	//fmt.Printf("%+v\n", graph)
	paths := graph.FindPaths(START, []string{START})
	//fmt.Printf("%+v\n", paths)
	fmt.Println(len(paths))
}
