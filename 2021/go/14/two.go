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

func parseRules(lines []string) map[string][]string {
	rules := make(map[string][]string)
	for _, line := range lines {
		var match, insert string
		_, err := fmt.Sscanf(line, "%2s -> %1s", &match, &insert)
		check(err)
		matchPieces := strings.SplitN(match, "", 2)
		rules[match] = []string{matchPieces[0] + insert, insert + matchPieces[1]}
	}
	return rules
}

func parsePolymer(line string) map[string]int64 {
	polymer := make(map[string]int64)
	for i := 0; i < len(line)-1; i++ {
		polymer[line[i:i+2]]++
	}
	return polymer
}

func growPolymer(in map[string]int64, rules map[string][]string) map[string]int64 {
	out := make(map[string]int64)
	for i := range in {
		for _, o := range rules[i] {
			out[o] += in[i]
		}
	}
	return out
}

func scorePolymer(polymer map[string]int64, extra rune) int64 {
	scores := make(map[rune]int64)
	for s := range polymer {
		runes := []rune(s)
		scores[runes[0]] += polymer[s]
	}
	scores[extra]++
	var min int64 = math.MaxInt64
	var max int64 = 0
	var least, most rune
	for c := range scores {
		if scores[c] > max {
			max = scores[c]
			most = c
		}
		if scores[c] < min {
			min = scores[c]
			least = c
		}
	}
	fmt.Println(string(least), min, string(most), max)
	return max - min
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	template := lines[0]
	polymer := parsePolymer(template)
	fmt.Printf("%s %+v\n", template, polymer)
	rules := parseRules(lines[1:])
	fmt.Printf("%+v\n", rules)
	for i := 0; i < 40; i++ {
		//fmt.Printf("%d %d %+v\n", i, scorePolymer(polymer), polymer)
		polymer = growPolymer(polymer, rules)
	}
	fmt.Println(scorePolymer(polymer, rune(template[len(template)-1])))
}
