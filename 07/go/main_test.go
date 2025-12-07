package main

import "testing"

func TestStepPart1(t *testing.T) {
	input := []string {
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",

	}

	runes := [][]rune{}
	for _, s := range input {
		runes = append(runes, []rune(s))
	}

	if res := stepPart1(runes); res != 21 {
		t.Errorf("expected %d, got %d", 21, res)
	}
}

func TestStepPart2(t *testing.T) {
	input := []string {
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",

	}

	runes := [][]rune{}
	for _, s := range input {
		runes = append(runes, []rune(s))
	}

	memo := make([][]int, len(runes))
	for i := range memo {
		memo[i] = make([]int, len(runes[i]))
	}

	if res := stepPart2Alt(runes, Coord{x: 0, y: 7}, memo); res != 40 {
		t.Errorf("expected %d, got %d", 40, res)
	}
}
