package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

type Turn struct {
	Dir      Dir
	Distance int
}

func ParseWirePath(path string) ([]Turn, error) {
	turns := []Turn{}

	for _, turnRaw := range strings.Split(path, ",") {
		var dir Dir
		dirRune := turnRaw[0]
		switch dirRune {
		case 'U':
			dir = DirUp
			break
		case 'R':
			dir = DirRight
			break
		case 'D':
			dir = DirDown
			break
		case 'L':
			dir = DirLeft
			break
		default:
			return nil, fmt.Errorf("unknown direction %v", dirRune)
		}

		distanceRaw := turnRaw[1:]
		distance, err := strconv.ParseInt(distanceRaw, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse distance: %w", err)
		}

		turns = append(turns, Turn{dir, int(distance)})
	}

	return turns, nil
}

type Point struct {
	X, Y int
}

func (p Point) Magnitude() int {
	x := p.X
	if x < 0 {
		x = -x
	}

	y := p.Y
	if y < 0 {
		y = -y
	}

	return x + y
}

func WireToPoints(wire []Turn) []Point {
	cursor := Point{0, 0}
	points := []Point{cursor}

	for _, turn := range wire {
		for i := 0; i < turn.Distance; i++ {
			switch turn.Dir {
			case DirUp:
				cursor.Y++
				break
			case DirRight:
				cursor.X++
				break
			case DirDown:
				cursor.Y--
				break
			case DirLeft:
				cursor.X--
				break
			}
			points = append(points, cursor)
		}
	}

	return points
}

func PointSet(points []Point) map[Point]struct{} {
	set := make(map[Point]struct{}, len(points))
	for _, point := range points {
		set[point] = struct{}{}
	}
	return set
}

func ClosestIntersection(wire1 []Turn, wire2 []Turn) (Point, bool) {
	// Count number of times each point on the field was crossed (visited) by a wire.
	visited := map[Point]int{}
	for point := range PointSet(WireToPoints(wire1)) {
		if visited[point] < 1 {
			visited[point]++
		}
	}
	for point := range PointSet(WireToPoints(wire2)) {
		if visited[point] < 2 {
			visited[point]++
		}
	}

	// Get intersection points by filtering out points with more than 1 visit.
	intersections := []Point{}
	for point, visits := range visited {
		// Ignore 0,0 point
		isOrigin := point.X == 0 && point.Y == 0
		if visits > 1 && !isOrigin {
			intersections = append(intersections, point)
		}
	}

	// No intersections
	if len(intersections) == 0 {
		return Point{}, false
	}

	// Find intersection closest to the central point (0, 0)
	closest := intersections[0]
	for _, point := range intersections {
		if point.Magnitude() < closest.Magnitude() {
			closest = point
		}
	}

	return closest, true
}

func CountSteps(wire []Turn) map[Point]int {
	steps := map[Point]int{}
	for i, point := range WireToPoints(wire) {
		if steps[point] == 0 {
			steps[point] = i
		}
	}
	return steps
}

func StepsToIntersection(wire1 []Turn, wire2 []Turn) (int, bool) {
	stepsMap1 := CountSteps(wire1)
	stepsMap2 := CountSteps(wire2)

	stepsTo := -1
	for point, steps1 := range stepsMap1 {
		if point.X == 0 && point.Y == 0 {
			continue
		}

		steps2, ok := stepsMap2[point]
		if !ok {
			continue
		}

		if stepsTo == -1 || (steps1+steps2) < stepsTo {
			stepsTo = steps1 + steps2
		}
	}

	if stepsTo == -1 {
		return 0, false
	}

	return stepsTo, true
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	input, err := ioutil.ReadAll(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wirePathsRaw := strings.Split(strings.TrimSpace(string(input)), "\n")

	wirePath1, err := ParseWirePath(wirePathsRaw[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wirePath2, err := ParseWirePath(wirePathsRaw[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	intersection, ok := ClosestIntersection(wirePath1, wirePath2)
	if !ok {
		fmt.Println("No intersection found")
		os.Exit(1)
	}

	fmt.Printf("Intersection closest to the center: %v\n", intersection)
	fmt.Printf("Manhattan distance: %v\n", intersection.Magnitude())

	stepsToFirst, ok := StepsToIntersection(wirePath1, wirePath2)
	if !ok {
		fmt.Println("No intersection found")
		os.Exit(1)
	}

	fmt.Printf("Soonest intersection is reached after combined steps of: %v\n", stepsToFirst)
}
