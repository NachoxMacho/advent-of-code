package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"strings"
)

const nums = "987654321"

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

	result1 := 0.0
	result2 := 0.0

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		result1 += stepPart2(scanner.Text(), 2)
		result2 += stepPart2(scanner.Text(), 12)

	}

	fmt.Println("result1", result1)
	fmt.Println("result2", result2)

	return nil
}

// For posterity, but able to use part 2 to solve part 1
func stepPart1(input string) int {
	for i, c := range nums {
		index := strings.IndexRune(input, c)
		if index == -1 {
			continue
		}
		for j, c := range nums {
			index2 := strings.IndexRune(input[index+1:], c)
			if index2 == -1 {
				continue
			}
			return (9-i)*10 + (9 - j)
		}
	}

	return 0
}

func stepPart2(input string, digits int) float64 {
	if digits == 0 {
		return 0
	}
	for i, c := range nums {
		index := strings.IndexRune(input, c)
		if index == -1 {
			continue
		}
		if digits == 1 {
			return float64(9 - i)
		}
		res := stepPart2(input[index+1:], digits-1)
		if res != 0 {
			return float64(9-i)*math.Pow10(digits-1) + res
		}
	}

	return 0
}
