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

	result1 := 0
	// result2 := 0

	prev := ""
	current := ""
	next := ""

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		prev = current
		current = next
		next = scanner.Text()
		if current == "" {
			continue
		}
		result1 += stepPart1(prev, current, next)
		// result2 += stepPart2(scanner.Text(), 12)
	}

	prev = current
	current = next
	next = ""
	result1 += stepPart1(prev, current, next)

	fmt.Println("result1", result1)
	// fmt.Println("result2", result2)

	return nil
}

func stepPart1(prev string, current string, next string) int {
	result := 0
	fmt.Println("prev    ", prev)
	fmt.Println("current ", current)
	fmt.Println("next    ", next)
	currentCopy := []rune(strings.Clone(current))
	for i, c := range current {
		currentCopy[i] = ' '
		if c != '@' {
			continue
		}
		neighbors := 0
		if i > 0 && current[i-1] == '@' {
			neighbors++
		}
		if i < len(current)-1 && current[i+1] == '@' {
			neighbors++
		}
		if prev != "" {
			if i > 0 && prev[i-1] == '@' {
				neighbors++
			}
			if prev[i] == '@' {
				neighbors++
			}
			if i < len(prev)-1 && prev[i+1] == '@' {
				neighbors++
			}
		}
		if next != "" {
			if i > 0 && next[i-1] == '@' {
				neighbors++
			}
			if next[i] == '@' {
				neighbors++
			}
			if i < len(next)-1 && next[i+1] == '@' {
				neighbors++
			}
		}
		if neighbors < 4 {
			result++
			currentCopy[i] = 'x'
		}
	}
	fmt.Println("modded  ", string(currentCopy))
	fmt.Println("result", result)
	return result
}
