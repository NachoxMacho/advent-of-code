package main

import "testing"

func TestStepPart1(t *testing.T) {

	if res := stepPart1("","..@@.@@@@.", "@@@.@.@.@@"); res != 5 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 5)
	}
	if res := stepPart1("..@@.@@@@.", "@@@.@.@.@@", "@@@@@.@.@@"); res != 1 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 1)
	}
	if res := stepPart1("@@@.@.@.@@", "@@@@@.@.@@", "@.@@@@..@."); res != 1 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 1)
	}
	if res := stepPart1("@@@@@.@.@@", "@.@@@@..@.", "@@.@@@@.@@"); res != 0 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 0)
	}
	if res := stepPart1("@.@@@@..@.", "@@.@@@@.@@", ".@@@@@@@.@"); res != 2 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 2)
	}
	if res := stepPart1("@@.@@@@.@@", ".@@@@@@@.@",".@.@.@.@@@"); res != 0 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 0)
	}
	if res := stepPart1(".@@@@@@@.@",".@.@.@.@@@","@.@@@.@@@@"); res != 0 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 0)
	}
	if res := stepPart1(".@.@.@.@@@","@.@@@.@@@@",".@@@@@@@@."); res != 1 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 1)
	}
	if res := stepPart1("@.@@@.@@@@",".@@@@@@@@.","@.@.@@@.@."); res != 0 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 0)
	}
	if res := stepPart1(".@@@@@@@@.","@.@.@@@.@.",""); res != 3 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 3)
	}

	if res := stepPart1("...",".@.","..."); res != 1 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 1)
	}
	if res := stepPart1("...","@@@","..."); res != 3 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 3)
	}
	if res := stepPart1("@.@","@@@","..."); res != 2 {
		t.Errorf("stepPart1 Test: got %d want %d", res, 3)
	}
}

// func TestStepPart2(t *testing.T) {
//
// 	if res := stepPart2("91", 1); res != 9.0 {
// 		t.Errorf("step_part2Test: got %f want %d", res, 9)
// 	}
// 	if res := stepPart2("912", 2); res != 92.0 {
// 		t.Errorf("step_part2Test: got %f want %d", res, 92)
// 	}
// 	if res := stepPart2("987654321111111", 12); res != 987654321111.0 {
// 		t.Errorf("step_part2Test: got %f want %d", res, 987654321111)
// 	}
// 	if res := stepPart2("811111111111119", 12); res != 811111111119.0 {
// 		t.Errorf("step_part2Test: got %f want %d", res, 811111111119)
// 	}
// 	if res := stepPart2("234234234234278", 12); res != 434234234278.0 {
// 		t.Errorf("step_part2Test: got %f want %d", res, 434234234278)
// 	}
// 	if res := stepPart2("818181911112111", 12); res != 888911112111.0 {
// 		t.Errorf("step_part2Test: got %f want %d", res, 888911112111)
// 	}
//
// }
