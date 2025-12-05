package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
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

	result1 := 0
	result2 := 0

	prev := ""
	current := ""
	next := ""

	m := []string{""}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		m = append(m, scanner.Text())
	}
	m = append(m, "")
	neighbors := [][]int{}
	neighborsLen := 0
	for i, line := range m {
		if line == "" {
			continue
		}
		prev = m[i-1]
		current = m[i]
		next = m[i+1]
		if current == "" {
			continue
		}
		result1 += stepPart1(prev, current, next)

		neighbors = append(neighbors, []int{})
		for j := range line {
			neighbors[neighborsLen] = append(neighbors[neighborsLen], calculateNeighbors(prev, current, next, j))
		}
		neighborsLen++
	}

	startAlt := time.Now()
	result2Alt := 0
	for i := range neighbors {
		for j := range neighbors[i] {
			result2Alt += stepPart2Alt(neighbors, i, j, false)
		}
	}
	durationAlt := time.Since(startAlt)

	start2 := time.Now()
	for {
		totalOut := 0
		for i, line := range m {
			if line == "" {
				continue
			}
			prev = m[i-1]
			current = m[i]
			next = m[i+1]
			if current == "" {
				continue
			}
			out, newLine := stepPart2(prev, current, next)
			m[i] = newLine
			totalOut += out
		}
		if totalOut == 0 {
			break
		}
		result2 += totalOut
	}
	duration2 := time.Since(start2)

	for i := range neighbors {
		fmt.Println(strings.ReplaceAll(m[i+1], ".", " "))
		for j := range neighbors[i] {
			if neighbors[i][j] <= 0 {
				fmt.Print(" ")
				continue
			}
			fmt.Print(neighbors[i][j])

			if m[i+1][j] != '@' {
				fmt.Println("mismatch at", i, j, "left", neighbors[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println("result1", result1)
	fmt.Println("result2", result2, "time", duration2)
	fmt.Println("result2Alt", result2Alt, "time", durationAlt)

	return nil
}

func stepPart1(prev string, current string, next string) int {
	result := 0
	for i, c := range current {
		if c != '@' {
			continue
		}
		neighbors := calculateNeighbors(prev, current, next, i)
		if neighbors < 4 {
			result++
		}
	}
	return result
}

func stepPart2(prev string, current string, next string) (int, string) {
	result := 0
	currentCopy := []rune(strings.Clone(current))
	for i, c := range current {
		if c != '@' {
			continue
		}
		neighbors := calculateNeighbors(prev, current, next, i)
		if neighbors < 4 {
			result++
			currentCopy[i] = '.'
		}
	}
	return result, string(currentCopy)
}

func stepPart2Alt(neighbors [][]int, x int, y int, decrease bool) int {

	// out of bounds or not able to remove
	if x < 0 || x >= len(neighbors) {
		return 0
	}
	if y < 0 || y >= len(neighbors[x]) {
		return 0
	}
	if neighbors[x][y] < 0 {
		return 0
	}
	if decrease {
		neighbors[x][y]--
	}
	if neighbors[x][y] >= 4 {
		return 0
	}

	neighbors[x][y] = -1
	removed := 1
	removed += stepPart2Alt(neighbors, x-1, y-1, true)
	removed += stepPart2Alt(neighbors, x-1, y, true)
	removed += stepPart2Alt(neighbors, x-1, y+1, true)
	removed += stepPart2Alt(neighbors, x, y-1, true)
	removed += stepPart2Alt(neighbors, x, y+1, true)
	removed += stepPart2Alt(neighbors, x+1, y-1, true)
	removed += stepPart2Alt(neighbors, x+1, y, true)
	removed += stepPart2Alt(neighbors, x+1, y+1, true)

	return removed
}

func calculateNeighbors(prev string, current string, next string, x int) int {
	if current[x] != '@' {
		return -1
	}
	neighbors := 0
	if x > 0 && current[x-1] == '@' {
		neighbors++
	}
	if x < len(current)-1 && current[x+1] == '@' {
		neighbors++
	}
	if prev != "" {
		if x > 0 && prev[x-1] == '@' {
			neighbors++
		}
		if prev[x] == '@' {
			neighbors++
		}
		if x < len(prev)-1 && prev[x+1] == '@' {
			neighbors++
		}
	}
	if next != "" {
		if x > 0 && next[x-1] == '@' {
			neighbors++
		}
		if next[x] == '@' {
			neighbors++
		}
		if x < len(next)-1 && next[x+1] == '@' {
			neighbors++
		}
	}
	return neighbors
}
