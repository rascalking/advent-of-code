package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func parseRules(lines []string) map[string]string {
	rules := make(map[string]string)
	for _, line := range lines {
		var match, insert string
		_, err := fmt.Sscanf(line, "%2s -> %1s", &match, &insert)
		check(err)
		matchPieces := strings.SplitN(match, "", 2)
		rules[match] = matchPieces[0] + insert
	}
	return rules
}

func growPolymer(in string, rules map[string]string) string {
	out := make([]string, len(in)-1)
	for i := 0; i < len(in)-1; i++ {
		out[i] = rules[in[i:i+2]]
	}
	return strings.Join(out, "") + in[len(in)-1:]
}

func scorePolymer(polymer string) int {
	scores := make(map[rune]int)
	for _, c := range polymer {
		scores[c]++
	}
	min := math.MaxInt
	max := 0
	for c := range scores {
		if scores[c] > max {
			max = scores[c]
		}
		if scores[c] < min {
			min = scores[c]
		}
	}
	return max - min
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	polymer := lines[0]
	rules := parseRules(lines[1:])
	//fmt.Printf("%+v\n", rules)
	for i := 0; i < 10; i++ {
		//fmt.Println(polymer, len(polymer))
		polymer = growPolymer(polymer, rules)
	}
	fmt.Println(scorePolymer(polymer))
}
