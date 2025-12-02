package main

import (
	"fmt"
	"testing"
)

func TestStepPart2(t *testing.T) {

	var tests = []struct {
		inDial    int
		inInput   string
		wantDial  int
		wantZeros int
	}{
		// sample
		{
			inDial:    50,
			inInput:   "L68",
			wantDial:  82,
			wantZeros: 1,
		},
		{
			inDial:    82,
			inInput:   "L30",
			wantDial:  52,
			wantZeros: 0,
		},
		{
			inDial:    52,
			inInput:   "R48",
			wantDial:  0,
			wantZeros: 1,
		},
		{
			inDial:    0,
			inInput:   "L5",
			wantDial:  95,
			wantZeros: 0,
		},
		{
			inDial:    95,
			inInput:   "R60",
			wantDial:  55,
			wantZeros: 1,
		},
		{
			inDial:    55,
			inInput:   "L55",
			wantDial:  0,
			wantZeros: 1,
		},
		{
			inDial:    0,
			inInput:   "L1",
			wantDial:  99,
			wantZeros: 0,
		},
		{
			inDial:    99,
			inInput:   "L99",
			wantDial:  0,
			wantZeros: 1,
		},
		{
			inDial:    0,
			inInput:   "R14",
			wantDial:  14,
			wantZeros: 0,
		},
		{
			inDial:    14,
			inInput:   "L82",
			wantDial:  32,
			wantZeros: 1,
		},

		//edge cases
		{
			inDial:    0,
			inInput:   "L42",
			wantDial:  58,
			wantZeros: 0,
		},
		{
			inDial:    10,
			inInput:   "R420",
			wantDial:  30,
			wantZeros: 4,
		},
		{
			inDial:    10,
			inInput:   "R490",
			wantDial:  0,
			wantZeros: 5,
		},
		{
			inDial:    0,
			inInput:   "L600",
			wantDial:  0,
			wantZeros: 6,
		},
		{
			inDial:    10,
			inInput:   "L600",
			wantDial:  10,
			wantZeros: 6,
		},
		{
			inDial:    20,
			inInput:   "L333",
			wantDial:  87,
			wantZeros: 4,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("dial: %d input: %s", tt.inDial, tt.inInput)
		t.Run(testname, func(t *testing.T) {
			outDial, outZeros, _ := step_part2(tt.inDial, tt.inInput)
			if outDial != tt.wantDial {
				t.Errorf("dial: got %d want %d", outDial, tt.wantDial)
			}
			if outZeros != tt.wantZeros {
				t.Errorf("zeros: got %d want %d", outZeros, tt.wantZeros)
			}
		})

	}
}
