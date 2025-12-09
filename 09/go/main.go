package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
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

	coords := []Coord{}

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		split := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			return err
		}
		coords = append(coords, Coord{x: x, y: y})
	}

	res := stepPart1(coords)
	fmt.Println("result1", res)

	// start := 0

	// for i, cell := range runes[0] {
	// 	if cell == 'S' {
	// 		start = i
	// 		break
	// 	}
	// }

	// memo := make([][]int, len(runes))
	// for i := range memo {
	// memo[i] = make([]int, len(runes[i]))
	// }

	// startTime := time.Now()
	// res = stepPart2Alt(runes, Coord{x: 0, y: start}, memo)
	// duration := time.Since(startTime)
	// fmt.Println("result2", res, "cache hits", hits, "cache misses", misses, "duration", duration)

	return nil
}

type Coord struct {
	x int
	y int
}

func stepPart1(input []Coord) int {
	maxSize := 0.0

	for _, c1 := range input {
		for _, c2 := range input {
			area := (math.Abs(float64(c1.x-c2.x)) + 1.0) * (math.Abs(float64(c1.y-c2.y)) + 1.0)
			if area > maxSize {
				fmt.Println(area, c1, c2)
				maxSize = area
			}
		}
	}

	return int(maxSize)
}

func stepPart2Alt(input [][]rune, memo [][]int) int {
	res := 0
	return res
}
