package main

import (
	"bufio"
	"fmt"
	"io"
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

type Stack struct {
	stack []string
	Size  int
}

func NewStack() Stack {
	return Stack{stack: make([]string, 0)}
}

func (s *Stack) push(item string) {
	s.stack = append(s.stack, item)
	s.Size += 1
}

func (s *Stack) pop() (string, error) {
	if s.Size == 0 {
		return "", fmt.Errorf("empty stack")
	}
	item := s.stack[s.Size-1]
	s.stack = s.stack[:s.Size-1]
	s.Size -= 1
	return item, nil
}

func (s *Stack) peek() (string, error) {
	if s.Size == 0 {
		return "", fmt.Errorf("empty stack")
	}
	return s.stack[s.Size-1], nil
}

var opener = map[string]bool{
	"(": true,
	"[": true,
	"{": true,
	"<": true,
}

var closer = map[string]bool{
	")": true,
	"]": true,
	"}": true,
	">": true,
}

var matcher = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var reverseMatcher = make(map[string]string, 0)

func init() {
	for k := range matcher {
		reverseMatcher[matcher[k]] = k
	}
}

func scoreBrokenLine(line string) int {
	brokenScore := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	stack := NewStack()
	for _, c := range strings.Split(line, "") {
		if opener[c] {
			stack.push(c)
		} else if closer[c] {
			top, err := stack.peek()
			if err != nil {
				panic(err)
			}
			if matcher[c] == top {
				stack.pop()
			} else {
				return brokenScore[c]
			}
		} else {
			panic(fmt.Sprintf("invalid character: '%s'", c))
		}
	}
	return 0
}

func completeLine(line string) string {
	stack := NewStack()
	for _, c := range strings.Split(line, "") {
		if opener[c] {
			stack.push(c)
		} else if closer[c] {
			top, err := stack.peek()
			check(err)
			if matcher[c] == top {
				stack.pop()
			} else {
				return ""
			}
		} else {
			panic(fmt.Sprintf("invalid character: '%s'", c))
		}
	}

	completion := ""
	for {
		if stack.Size == 0 {
			break
		}
		c, err := stack.pop()
		check(err)
		completion += reverseMatcher[c]
	}
	return completion
}

func scoreCompletion(completion string) int {
	completionScore := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	score := 0
	for _, c := range strings.Split(completion, "") {
		score = score*5 + completionScore[c]
	}
	return score
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	scores := make([]int, 0)
	for _, line := range readLines(f) {
		completion := completeLine(line)
		if len(completion) > 0 {
			scores = append(scores, scoreCompletion(completion))
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
