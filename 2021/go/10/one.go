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

type Stack struct {
	stack []string
	size  int
}

func NewStack() Stack {
	return Stack{stack: make([]string, 0)}
}

func (s *Stack) push(item string) {
	s.stack = append(s.stack, item)
	s.size += 1
}

func (s *Stack) pop() (string, error) {
	if s.size == 0 {
		return "", fmt.Errorf("empty stack")
	}
	item := s.stack[s.size-1]
	s.stack = s.stack[:s.size-1]
	s.size -= 1
	return item, nil
}

func (s *Stack) peek() (string, error) {
	if s.size == 0 {
		return "", fmt.Errorf("empty stack")
	}
	return s.stack[s.size-1], nil
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

var scorer = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func scoreLine(line string) int {
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
				return scorer[c]
			}
		} else {
			panic(fmt.Sprintf("invalid character: '%s'", c))
		}
	}
	return 0
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	score := 0
	for _, line := range readLines(f) {
		score += scoreLine(line)
	}
	fmt.Println(score)
}
