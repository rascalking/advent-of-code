package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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
	mappings map[string]int
}

func normalize(raw string) string {
	chars := strings.Split(raw, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func parseEntry(line string) Entry {
	fields := strings.Split(line, "|")
	entry := Entry{mappings: make(map[string]int, 0)}
	for _, pattern := range strings.Split(strings.Trim(fields[0], " "), " ") {
		entry.patterns = append(entry.patterns, normalize(pattern))
	}
	for _, output := range strings.Split(strings.Trim(fields[1], " "), " ") {
		entry.outputs = append(entry.outputs, normalize(output))
	}
	entry.calculateMappings()
	return entry
}

func (e *Entry) addMapping(pattern string, number int) {
	e.mappings[normalize(pattern)] = number
}

func (e *Entry) calculateMappings() {
	// start with the unique ones
	tmp := make(map[int]string, 0)
	for i := 0; i < len(e.patterns); i++ {
		switch len(e.patterns[i]) {
		case 2:
			e.addMapping(e.patterns[i], 1)
			tmp[1] = e.patterns[i]
		case 3:
			e.addMapping(e.patterns[i], 7)
			tmp[7] = e.patterns[i]
		case 4:
			e.addMapping(e.patterns[i], 4)
			tmp[4] = e.patterns[i]
		case 7:
			e.addMapping(e.patterns[i], 8)
			tmp[8] = e.patterns[i]
		}
	}

	//var top, topLeft, topRight, middle, bottomLeft, bottomRight, bottom rune
	var topRight, bottomRight rune

	/*
		// the char in 7 and not in 1 is top
		for _, char := range tmp[7] {
			if !strings.ContainsRune(tmp[1], char) {
				top = char
			}
		}
	*/

	// we know 1, 4, 7, 8
	// we know top

	// two of the 6-len patterns can be identified by a number we already know
	for i := 0; i < len(e.patterns); i++ {
		switch len(e.patterns[i]) {
		case 6:
			for _, char := range tmp[8] {
				if strings.ContainsRune(tmp[1], char) && !strings.ContainsRune(e.patterns[i], char) {
					// if the missing character is in 1, the number is 6, the character is top-right
					topRight = char
					tmp[6] = e.patterns[i]
					e.addMapping(e.patterns[i], 6)
				} else if strings.ContainsRune(tmp[4], char) && !strings.ContainsRune(e.patterns[i], char) {
					// if the missing character is in 4, the number is 0, the character is middle
					//middle = char
					tmp[0] = e.patterns[i]
					e.addMapping(e.patterns[i], 0)
				}
			}
		}
	}

	// we know 0, 1, 4, 6, 7, 8
	// we know top, topRight, middle

	// process of elimination gets us the third 6-len pattern, 9, and bottom-left
	for i := 0; i < len(e.patterns); i++ {
		switch len(e.patterns[i]) {
		case 6:
			if _, ok := e.mappings[e.patterns[i]]; !ok {
				e.addMapping(e.patterns[i], 9)
				tmp[9] = e.patterns[i]
				/*
					for _, char := range e.patterns[i] {
						if !strings.ContainsRune(tmp[8], char) {
							bottomLeft = char
							break
						}
					}
				*/
				break
			}
		}
	}

	// we know 0, 1, 4, 6, 7, 8, 9
	// we know top, topRight, middle, bottomLeft

	// now that we know top-right, we can figure out bottom-right from 1
	for _, char := range tmp[1] {
		if char != topRight {
			bottomRight = char
			break
		}
	}

	// we know 0, 1, 4, 6, 7, 8, 9
	// we know top, topRight, middle, bottomLeft, bottomRight

	for i := 0; i < len(e.patterns); i++ {
		switch len(e.patterns[i]) {
		case 5:
			if !strings.ContainsRune(e.patterns[i], topRight) {
				e.addMapping(e.patterns[i], 5)
			} else if !strings.ContainsRune(e.patterns[i], bottomRight) {
				e.addMapping(e.patterns[i], 2)
			} else {
				e.addMapping(e.patterns[i], 3)
			}
		}
	}
}

func (e *Entry) value() int {
	value := 0
	for i := 0; i < len(e.outputs); i++ {
		value += e.mappings[e.outputs[i]] * int(math.Pow10(3-i))
	}
	return value
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	sum := 0
	for _, line := range readLines(f) {
		entry := parseEntry(line)
		fmt.Printf("%v\n", entry)
		fmt.Printf("%d\n\n", entry.value())
		sum += entry.value()
	}
	fmt.Println(sum)
}
