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

	result1 := 0
	ranges := []Range{}

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		line := scanner.Text()
		split := strings.Split(line, "-")
		l, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return err
		}
		h, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return err
		}
		ranges = append(ranges, Range{int(l), int(h)})
	}
	slices.SortStableFunc(ranges, func(a Range, b Range) int {
		if a.Low < b.Low {
			return -1
		}
		if a.Low > b.Low {
			return 1
		}
		return 0
	})
	// fmt.Println(ranges)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return err
		}
		res := stepPart1(ranges, int(num))
		// fmt.Println(num, ",", res)

		if res {
			result1++
		}
	}
	fmt.Println("result1", result1)

	start := time.Now()
	res := stepPart2Alt(ranges)
	duration := time.Since(start)
	fmt.Println("result2", res, "time", duration)
	return nil
}

func stepPart1(ranges []Range, n int) bool {
	for _, r := range ranges {
		// fmt.Println("comparing", n, "with range", r)
		if n >= r.Low && n <= r.High {
			// fmt.Println("found", n)
			return true
		}
	}
	return false
}

func stepPart2Alt(ranges []Range) int {
	rangesCopy := slices.Clone(ranges)
	res := 0
	comparisons := 0
	for i, r1 := range rangesCopy {
		for j, r2 := range rangesCopy {
			if i == j {
				continue
			}
			if r2.High < r1.Low {
				continue
			}
			if r2.Low > r1.High {
				break
			}
			comparisons++
			if r1.High >= r2.High && r1.Low <= r2.Low {
				rangesCopy[j].Low = 0
				rangesCopy[j].High = 0
				continue
			}
			if r1.High >= r2.Low && r1.Low <= r2.Low {
				rangesCopy[j].Low = r1.High + 1
				continue
			}
			if r1.Low <= r2.High && r1.High >= r2.High {
				rangesCopy[j].High = r1.Low - 1
				continue
			}
		}
	}
	fmt.Println("comparisons", comparisons, "size", len(ranges))
	for _, r := range rangesCopy {
		if r.Low == 0 && r.High == 0 {
			continue
		}
		res += r.High - r.Low + 1
	}
	return res
}
