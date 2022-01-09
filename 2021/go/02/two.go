package main

import (
	"bufio"
	"fmt"
    "io"
	"os"
	"strconv"
    "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Command struct {
    Direction string
    Distance int
}

type Submarine struct {
    Horizontal int
    Depth int
    Aim int
}

func getInput(r io.Reader) []Command {
    var commands []Command

	scanner := bufio.NewScanner(r)
    for scanner.Scan() {
		line := scanner.Text()
        if line == "" {
            continue
        }
        fields := strings.SplitN(line, " ", 2)
		n, err := strconv.Atoi(fields[1])
        check(err)
        commands = append(commands, Command{fields[0], n})
    }
    check(scanner.Err())
    return commands
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

    sub := Submarine{}
    for _, command := range getInput(f) {
        switch command.Direction {
        case "forward":
            sub.Horizontal += command.Distance
            sub.Depth += sub.Aim * command.Distance
            if sub.Depth < 0 {
                sub.Depth = 0
            }
        case "down":
            sub.Aim += command.Distance
        case "up":
            sub.Aim -= command.Distance
        default:
            panic(command)
        }
        fmt.Println(command, sub)
    }
    fmt.Println(sub)
	fmt.Println(sub.Horizontal * sub.Depth)
}
