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

func readLines(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func parseDrawnNumbers(line string) []int {
	var nums []int
	for _, numStr := range strings.Split(line, ",") {
		n, err := strconv.Atoi(numStr)
		check(err)
		nums = append(nums, n)
	}
	return nums
}

func parseBoards(lines []string) map[int]*Board {
	boards := make(map[int]*Board, 0)
	boardNum := 0
	for {
		board := Board{}
		for i := 0; i < 5; i++ {
			var nums [5]int
			_, err := fmt.Sscanf(lines[i], "%2d %2d %2d %2d %2d", &nums[0], &nums[1], &nums[2], &nums[3], &nums[4])
			check(err)
			for j, n := range nums {
				board.numbers[i][j] = n
			}
		}
		boards[boardNum] = &board
		boardNum += 1
		if len(lines) >= 6 {
			lines = lines[6:]
		} else {
			break
		}
	}
	return boards
}

type Board struct {
	numbers [5][5]int
	marked  [5][5]bool
}

func (b *Board) mark(num int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] == num {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *Board) isWinner() bool {
	for i := 0; i < 5; i++ {
		if b.marked[i][0] &&
			b.marked[i][1] &&
			b.marked[i][2] &&
			b.marked[i][3] &&
			b.marked[i][4] {
			return true
		}
		if b.marked[0][i] &&
			b.marked[1][i] &&
			b.marked[2][i] &&
			b.marked[3][i] &&
			b.marked[4][i] {
			return true
		}
	}
	return false
}

func (b *Board) score(called int) int {
	var n int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				n += b.numbers[i][j]
			}
		}
	}

	return n * called
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	lines := readLines(f)
	nums := parseDrawnNumbers(lines[0])
	boards := parseBoards(lines[2:])
	lastWinnerScore := -1

	for _, called := range nums {
		for i, board := range boards {
			board.mark(called)
			if board.isWinner() {
				lastWinnerScore = board.score(called)
				delete(boards, i)
			}
		}
	}
	fmt.Println(lastWinnerScore)
}
