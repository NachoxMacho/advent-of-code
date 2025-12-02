package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"strconv"
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
	dial1 := 50
	zeros1 := 0

	dial2 := 50
	zeros2 := 0
	z2 := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		dial1, err = step_part1(dial1, scanner.Text())
		if err != nil {
			return err
		}
		if dial1 == 0 {
			zeros1++
		}

		dial2, z2, err = step_part2(dial2, scanner.Text())
		if err != nil {
			return err
		}
		zeros2 += z2
	}

	fmt.Println("found", zeros1, "zeros in password 1.")
	fmt.Println("found", zeros2, "zeros in password 2.")
	return nil
}

func step_part1(start int, input string) (int, error) {
	direction := 1
	if input[0] == 'L' {
		direction = -1
	}

	inputNum, err := strconv.ParseInt(input[1:], 10, 64)
	if err != nil {
		return 0, err
	}

	// slog.Info("start", slog.Int("start", start), slog.String("input", input), slog.Int("move", int(inputNum)), slog.Int("end", start+int(inputNum)*direction))

	return (start + int(inputNum)*direction) % 100, nil
}

func step_part2(start int, input string) (int, int, error) {
	direction := 1
	if input[0] == 'L' {
		direction = -1
	}

	inputNum, err := strconv.ParseInt(input[1:], 10, 64)
	if err != nil {
		return 0, 0, err
	}

 	end := (start + int(inputNum % 100)*direction)
	zeros := int(math.Abs(math.Trunc(float64(inputNum) / 100.0)))
	if end >= 100 {
		zeros++
	}
	if end <= 0 && start != 0 {
		zeros++
	}

	if end < 0 {
		end += 100
	}

	// slog.Info("step", slog.Int("start", start), slog.String("input", input), slog.Int("move", int(inputNum)), slog.Int("end", end), slog.Int("zeros", zeros))

	return int(math.Abs(float64(end % 100))), zeros, nil
}
