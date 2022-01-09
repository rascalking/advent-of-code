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

type Position struct {
    Horizontal int
    Depth int
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

    pos := Position{0, 0}
    for _, command := range getInput(f) {
        switch command.Direction {
        case "forward":
            pos.Horizontal += command.Distance
        case "down":
            pos.Depth += command.Distance
        case "up":
            pos.Depth -= command.Distance
            if pos.Depth < 0 {
                pos.Depth = 0
            }
        default:
            panic(command)
        }
        fmt.Println(command, pos)
    }
    fmt.Println(pos)
	fmt.Println(pos.Horizontal * pos.Depth)
}
