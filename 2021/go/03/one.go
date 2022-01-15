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

type Accumulator struct {
	a []map[string]int
}

func newAccumulator(length int) Accumulator {
	acc := Accumulator{}
	acc.a = make([]map[string]int, length)
	for i := 0; i < length; i++ {
		acc.a[i] = map[string]int{"0": 0, "1": 0}
	}
	return acc
}

func (a *Accumulator) addEntry(entry []string) {
	for i, bit := range entry {
		a.a[i][bit] += 1
	}
}

func (a *Accumulator) gamma() int64 {
	bits := ""
	for _, bit := range a.a {
		if bit["0"] > bit["1"] {
			bits += "0"
		} else {
			bits += "1"
		}
	}

	g, err := strconv.ParseInt(bits, 2, 0)
	check(err)
	return g
}

func (a *Accumulator) epsilon() int64 {
	bits := ""
	for _, bit := range a.a {
		if bit["0"] < bit["1"] {
			bits += "0"
		} else {
			bits += "1"
		}
	}

	g, err := strconv.ParseInt(bits, 2, 0)
	check(err)
	return g
}

func accumulateBits(r io.Reader) Accumulator {
	var acc Accumulator
	accInited := false
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		bits := strings.Split(line, "")
		if accInited == false {
			acc = newAccumulator(len(bits))
			accInited = true
		}
		acc.addEntry(bits)
	}
	check(scanner.Err())
	return acc
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)
	acc := accumulateBits(f)
	fmt.Println(acc.gamma() * acc.epsilon())
}
