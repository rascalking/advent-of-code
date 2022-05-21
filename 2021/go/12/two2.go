package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const START = "start"
const END = "end"
const PATH_SEPARATOR = ","

type Graph struct {
	Edges map[string][]string
}

func AddPath(paths map[string]bool, path []string) {
	pathStr := strings.Join(path, PATH_SEPARATOR)
	paths[pathStr] = true
}

func AddPaths(paths map[string]bool, toAdd map[string]bool) {
	for p := range toAdd {
		if toAdd[p] {
			paths[p] = true
		}
	}
}

func ContainsPath(paths map[string]bool, path []string) bool {
	pathStr := strings.Join(path, PATH_SEPARATOR)
	return paths[pathStr]
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

func hasDuplicateSmall(path []string) bool {
	seen := make(map[string]bool)
	for _, node := range path {
		if node == strings.ToLower(node) {
			if seen[node] {
				return true
			} else {
				seen[node] = true
			}
		}
	}
	return false
}

func (g *Graph) FindPaths(node string, path []string) map[string]bool {
	fmt.Printf("FindPaths(%v, %v))\n", node, path)
	paths := make(map[string]bool)
	for _, n := range g.Edges[node] {
		n_path := append(path, n)
		if n == END {
			AddPath(paths, n_path)
		} else if n == START {
			continue
		} else if (n == strings.ToLower(n)) && containsItem(path, n) {
			if hasDuplicateSmall(path) {
				continue
			} else {
				AddPaths(paths, g.FindPaths(n, n_path))
			}
		} else {
			AddPaths(paths, g.FindPaths(n, n_path))
		}
	}
	fmt.Printf("\tFindPaths(%v, %v)) -> %v\n", node, path, paths)
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
	allPaths := make([]string, 0, len(paths))
	for p := range paths {
		if paths[p] {
			allPaths = append(allPaths, p)
		}
	}
	fmt.Println(len(allPaths))
}
