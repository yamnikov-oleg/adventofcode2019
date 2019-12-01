package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func fuelRequired(mass int) int {
	return (mass / 3) - 2
}

func fuelRequiredAdditional(mass int) int {
	fuel := fuelRequired(mass)
	if fuel <= 0 {
		return 0
	}

	return fuel + fuelRequiredAdditional(fuel)
}

func main() {
	inputFileName := flag.String("i", "input.txt", "input.txt")

	flag.Parse()

	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(inputFile)
	totalFuel := 0
	totalFuelAdditional := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("Invalid mass %q: %v\n", line, err)
			os.Exit(1)
		}
		totalFuel += fuelRequired(int(mass))
		totalFuelAdditional += fuelRequiredAdditional(int(mass))
	}

	fmt.Printf("Total fuel required: %v\n", totalFuel)
	fmt.Printf("Total fuel with additional: %v\n", totalFuelAdditional)
}
