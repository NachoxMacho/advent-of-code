package main

import "testing"

func TestStepPart1(t *testing.T) {
	ranges := []Range{
		{
			Low: 3,
			High: 5,
		},
		{
			Low: 10,
			High: 14,
		},
		{
			Low: 16,
			High:20,
		},
		{
			Low: 12,
			High: 18,
		},
	}

	if res := stepPart1(ranges, 1); res != false {
		t.Errorf("expected false, got %v", res)
	}
	if res := stepPart1(ranges, 3); res != true {
		t.Errorf("expected true, got %v", res)
	}
	if res := stepPart1(ranges, 5); res != true {
		t.Errorf("expected true, got %v", res)
	}
	if res := stepPart1(ranges, 8); res != false {
		t.Errorf("expected false, got %v", res)
	}
	if res := stepPart1(ranges, 10); res != true {
		t.Errorf("expected true, got %v", res)
	}
	if res := stepPart1(ranges, 11); res != true {
		t.Errorf("expected true, got %v", res)
	}
	if res := stepPart1(ranges, 17); res != true {
		t.Errorf("expected true, got %v", res)
	}
	if res := stepPart1(ranges, 32); res != false {
		t.Errorf("expected false, got %v", res)
	}

}

func TestStepPart2Alt(t *testing.T) {
	ranges := []Range{
		{
			Low: 3,
			High: 5,
		},
		{
			Low: 10,
			High: 14,
		},
		{
			Low: 16,
			High:20,
		},
		{
			Low: 12,
			High: 18,
		},
	}
	want := 14

	if res := stepPart2Alt(ranges); res != want {
		t.Errorf("expected %d, got %d", want, res)
	}

	ranges = []Range{
		{
			Low: 534042781792,
			High: 2379343072995,
		},
	}
	want = 1845300291204
	if res := stepPart2Alt(ranges); res != want {
		t.Errorf("expected %d, got %d", want, res)
	}

	ranges = []Range{
		{
			Low: 534042781792,
			High: 2379343072995,
		},
		{
			Low: 534042781792,
			High: 2379343072994,
		},
	}
	want = 1845300291204
	if res := stepPart2Alt(ranges); res != want {
		t.Errorf("expected %d, got %d", want, res)
	}

	ranges = []Range{
		{
			Low: 534042781792,
			High: 2379343072995,
		},
		{
			Low: 534042781792,
			High: 2379343072997,
		},
	}
	want = 1845300291206
	if res := stepPart2Alt(ranges); res != want {
		t.Errorf("expected %d, got %d", want, res)
	}

	ranges = []Range{
		{
			Low: 534042781792,
			High: 2379343072995,
		},
		{
			Low: 534042781792,
			High: 2379343072997,
		},
		{
			Low: 534042781795,
			High: 2379343072999,
		},
	}
	want = 1845300291208
	if res := stepPart2Alt(ranges); res != want {
		t.Errorf("expected %d, got %d", want, res)
	}

	ranges = []Range{
		{
			Low: 534042781792,
			High: 2379343072995,
		},
		{
			Low: 534042781791,
			High: 2379343072997,
		},
		{
			Low: 534042781790,
			High: 2379343072999,
		},
	}
	want = 1845300291210
	if res := stepPart2Alt(ranges); res != want {
		t.Errorf("expected %d, got %d", want, res)
	}
}
