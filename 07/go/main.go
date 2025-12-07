package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	err := run()
	if err != nil {
		slog.Error("failed", slog.String("error", err.Error()))
	}
}

func run() error {
	f, err := os.Open("input-real.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	runes := [][]rune{}

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		runes = append(runes, []rune(scanner.Text()))
	}

	res := stepPart1(runes)
	fmt.Println("result1", res)

	start := 0

	for i, cell := range runes[0] {
		if cell == 'S' {
			start = i
			break
		}
	}

	memo := make([][]int, len(runes))
	for i := range memo {
		memo[i] = make([]int, len(runes[i]))
	}

	startTime := time.Now()
	res = stepPart2Alt(runes, Coord{x: 0, y: start}, memo)
	duration := time.Since(startTime)
	fmt.Println("result2", res, "cache hits", hits, "cache misses", misses, "duration", duration)

	return nil
}

func stepPart1(input [][]rune) int {
	res := 0
	beams := []int{}

	for j, cell := range input[0] {
		if cell == 'S' {
			beams = append(beams, j)
			break
		}
	}

	for i, row := range input {
		if i == 0 {
			continue
		}
		fmt.Println("working on row:", i)
		fmt.Println("scanning beams:", beams)
		newBeams := []int{}
		for _, index := range beams {
			if row[index] == '^' {
				fmt.Println("found a splitter", index)
				res++
				if index+1 < len(row) && !slices.Contains(newBeams, index+1) {
					fmt.Println("adding index+1", index+1)
					newBeams = append(newBeams, index+1)
				}
				if index-1 >= 0 && !slices.Contains(newBeams, index-1) {
					fmt.Println("adding index-1", index-1)
					newBeams = append(newBeams, index-1)
				}
				continue
			}
			if !slices.Contains(newBeams, index) {
				newBeams = append(newBeams, index)
			}
		}
		beams = newBeams
	}

	return res
}

type Coord struct {
	x int
	y int
}

var hits = 0
var misses = 0

func stepPart2Alt(input [][]rune, coord Coord, memo [][]int) int {
	if coord.x >= len(input) {
		return 1
	}
	if coord.y < 0 || coord.y >= len(input[coord.x]) {
		return 0
	}
	if memo[coord.x][coord.y] != 0 {
		hits++
		return memo[coord.x][coord.y]
	}
	misses++
	res := 0
	if input[coord.x][coord.y] == '^' {
		res = stepPart2Alt(input, Coord{x: coord.x + 1, y: coord.y + 1}, memo) + stepPart2Alt(input, Coord{x: coord.x + 1, y: coord.y - 1}, memo)
	} else {
		res = stepPart2Alt(input, Coord{x: coord.x + 1, y: coord.y}, memo)
	}
	memo[coord.x][coord.y] = res
	return res
}
