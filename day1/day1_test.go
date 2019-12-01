package main

import "testing"

func TestFuelRequied(t *testing.T) {
	cases := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, c := range cases {
		fuel := fuelRequired(c.mass)
		if fuel != c.fuel {
			t.Errorf("Mass %v: expected %v got %v", c.mass, c.fuel, fuel)
		}
	}
}

func TestFuelRequiedAdditional(t *testing.T) {
	cases := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, c := range cases {
		fuel := fuelRequiredAdditional(c.mass)
		if fuel != c.fuel {
			t.Errorf("Mass %v: expected %v got %v", c.mass, c.fuel, fuel)
		}
	}
}
