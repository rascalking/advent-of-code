package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

type Entry struct {
	patterns []string
	outputs  []string
}

func parseEntry(line string) Entry {
	fields := strings.Split(line, "|")
	return Entry{
		patterns: strings.Split(strings.Trim(fields[0], " "), " "),
		outputs:  strings.Split(strings.Trim(fields[1], " "), " "),
	}
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	uniqueLengths := map[int]bool{2: true, 3: true, 4: true, 7: true}
	count := 0
	for _, line := range readLines(f) {
		entry := parseEntry(line)
		for _, val := range entry.outputs {
			if uniqueLengths[len(val)] {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
