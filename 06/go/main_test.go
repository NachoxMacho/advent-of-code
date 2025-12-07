package main

import "testing"

func TestStepPart2String(t *testing.T) {
	input := []string {
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
	}
	ops := []rune("*   +   *   +  ")

	runes := [][]rune{}
	for _, s := range input {
		runes = append(runes, []rune(s))
	}

	if res, _ := stepPart2String(runes, ops); res != 3263827 {
		t.Errorf("expected %d, got %d", 3263827, res)
	}
}
