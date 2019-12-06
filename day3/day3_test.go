package main

import "testing"

func TestParseWirePath(t *testing.T) {
	cases := []struct {
		in  string
		out []Turn
	}{
		{"R8,U5,L5,D3", []Turn{{DirRight, 8}, {DirUp, 5}, {DirLeft, 5}, {DirDown, 3}}},
		{"U7,R6,D4,L4", []Turn{{DirUp, 7}, {DirRight, 6}, {DirDown, 4}, {DirLeft, 4}}},
	}

	for _, c := range cases {
		t.Logf("input: %v\n", c.in)

		out, err := ParseWirePath(c.in)
		if err != nil {
			t.Errorf("errored: %v", err)
			continue
		}

		if len(c.out) != len(out) {
			t.Errorf("length mismatch: %v != %v", out, c.out)
			continue
		}

		for i := 0; i < len(out); i++ {
			if c.out[i] != out[i] {
				t.Errorf("element %v mismatch: %v != %v", i, out, c.out)
				continue
			}
		}
	}
}

func TestClosestIntersection(t *testing.T) {
	cases := []struct {
		wire1     string
		wire2     string
		magnitude int
	}{
		{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			6,
		},
		{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			159,
		},
		{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			135,
		},
	}

	for _, c := range cases {
		t.Logf("input: %v + %v\n", c.wire1, c.wire2)

		wire1, err := ParseWirePath(c.wire1)
		if err != nil {
			t.Errorf("Error while parsing wire1: %v\n", err)
			continue
		}

		wire2, err := ParseWirePath(c.wire2)
		if err != nil {
			t.Errorf("Error while parsing wire2: %v\n", err)
			continue
		}

		point, ok := ClosestIntersection(wire1, wire2)
		if !ok {
			t.Errorf("No intersections found\n")
			continue
		}

		if point.Magnitude() != c.magnitude {
			t.Errorf("mismatch: %v.Magnitude (%v) != %v\n", point, point.Magnitude(), c.magnitude)
		}
	}
}

func TestStepsToIntersection(t *testing.T) {
	cases := []struct {
		wire1 string
		wire2 string
		steps int
	}{
		{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			30,
		},
		{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			610,
		},
		{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			410,
		},
	}

	for _, c := range cases {
		t.Logf("input: %v + %v\n", c.wire1, c.wire2)

		wire1, err := ParseWirePath(c.wire1)
		if err != nil {
			t.Errorf("Error while parsing wire1: %v\n", err)
			continue
		}

		wire2, err := ParseWirePath(c.wire2)
		if err != nil {
			t.Errorf("Error while parsing wire2: %v\n", err)
			continue
		}

		steps, ok := StepsToIntersection(wire1, wire2)
		if !ok {
			t.Errorf("No intersections found\n")
			continue
		}

		if steps != c.steps {
			t.Errorf("mismatch: %v != %v\n", steps, c.steps)
		}
	}
}
