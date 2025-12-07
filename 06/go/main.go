package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := run()
	if err != nil {
		slog.Error("failed", slog.String("error", err.Error()))
	}
}

type Range struct {
	Low  int
	High int
}

func run() error {
	f, err := os.Open("input-real.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	m := []string{}
	runes := [][]rune{}

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		m = append(m, scanner.Text())
		runes = append(runes, []rune(scanner.Text()))
	}

	numbers := [][]int{}
	for _, s := range m {
		if strings.Contains(s, "+") {
			continue
		}
		split := strings.Split(s, " ")
		split = slices.DeleteFunc(split, func(a string) bool { return a == "" })
		n := []int{}
		for _, c := range split {
			i, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				return err
			}
			n = append(n, int(i))
		}
		numbers = append(numbers, n)
	}
	splitOps := strings.Split(m[len(m)-1], " ")
	splitOps = slices.DeleteFunc(splitOps, func(a string) bool { return a == "" })

	res, err := stepPart1(numbers, splitOps)
	if err != nil {
		return err
	}
	fmt.Println("result1", res)

	ops := runes[len(runes)-1]
	runes = slices.Delete(runes, len(runes)-1, len(runes))

	start := time.Now()
	res, _ = stepPart2String(runes, ops)
	duration := time.Since(start)
	fmt.Println("result2", res, "duration", duration)
	return nil
}

func stepPart1(input [][]int, ops []string) (int, error) {
	res := 0

	for i, op := range ops {
		if op == "" {
			continue
		}
		switch op {
		case "*":
			temp := 1
			for j := range input {
				temp *= input[j][i]
			}
			res += temp
		case "+":
			temp := 0
			for j := range input {
				temp += input[j][i]
			}
			res += temp
		}
	}

	return res, nil
}

func stepPart2String(input [][]rune, opString []rune) (int, error) {

	result := 0
	var op rune
	temp := 0

	opStringLength := len(opString)

	for i := 0; i < len(input[0]); i++ {
		if i < opStringLength && opString[i] != ' ' {
			result += temp
			switch opString[i] {
			case '*':
				temp = 1
			case '+':
				temp = 0
			}
		}
		n := []rune{}
		for j := range input {
			if i >= len(input[j]) {
				continue
			}
			if input[j][i] == ' ' {
				continue
			}
			n = append(n, input[j][i])
		}
		if len(n) == 0 {
			continue
		}
		number, err := strconv.ParseInt(string(n), 10, 64)
		if err != nil {
			return 0, err
		}
		switch op {
		case '*':
			temp *= int(number)
		case '+':
			temp += int(number)
		}
	}

	return result + temp, nil
}
