package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
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

	runes := [][]rune{}

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		runes = append(runes, []rune(scanner.Text()))
	}

	res := stepPart1(runes)
	fmt.Println("result1", res)

	// start := 0

	// for i, cell := range runes[0] {
	// 	if cell == 'S' {
	// 		start = i
	// 		break
	// 	}
	// }

	memo := make([][]int, len(runes))
	for i := range memo {
		memo[i] = make([]int, len(runes[i]))
	}

	// startTime := time.Now()
	// res = stepPart2Alt(runes, Coord{x: 0, y: start}, memo)
	// duration := time.Since(startTime)
	// fmt.Println("result2", res, "cache hits", hits, "cache misses", misses, "duration", duration)

	return nil
}

func stepPart1(input [][]rune) int {
	res := 0

	return res
}

func stepPart2Alt(input [][]rune, memo [][]int) int {
	res := 0
	return res
}
