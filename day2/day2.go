package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Program is the number array of the program.
type Program []int

// ParseProgram parses a program from a string of comma-separated numbers.
func ParseProgram(s string) (Program, error) {
	positionsRaw := strings.Split(strings.TrimSpace(s), ",")
	positions := make([]int, 0, len(positionsRaw))
	for _, posRaw := range positionsRaw {
		pos, err := strconv.ParseInt(posRaw, 10, 64)
		if err != nil {
			return nil, err
		}

		positions = append(positions, int(pos))
	}
	return positions, nil
}

// CopyProgram creates a copy of the program. You can modify the copy without
// affecting its original.
func CopyProgram(p Program) Program {
	programCopy := make(Program, len(p))
	copy(programCopy, p)
	return programCopy
}

// Computer contains current state of program execution.
type Computer struct {
	Program Program
	// Position of currently executed operation.
	Position int
	// Has the program halted (reached opcode 99)?
	Halted bool
}

// NewComputer makes a new computer with the given program loaded in.
func NewComputer(program Program) *Computer {
	return &Computer{
		Program:  CopyProgram(program),
		Position: 0,
		Halted:   false,
	}
}

// RunOne runs the current operation in the program.
func (c *Computer) RunOne() {
	if c.Halted {
		return
	}

	opcode := c.Program[c.Position]
	switch opcode {
	case 1:
		operand1 := c.Program[c.Position+1]
		operand2 := c.Program[c.Position+2]
		output := c.Program[c.Position+3]
		c.Program[output] = c.Program[operand1] + c.Program[operand2]
		c.Position += 4
		break
	case 2:
		operand1 := c.Program[c.Position+1]
		operand2 := c.Program[c.Position+2]
		output := c.Program[c.Position+3]
		c.Program[output] = c.Program[operand1] * c.Program[operand2]
		c.Position += 4
		break
	case 99:
		c.Halted = true
		break
	default:
		panic(fmt.Sprintf("Invalid opcode: %v", opcode))
	}
}

// Run runs all the operations in the program until its halts or until
// `maxOperations` number of operations is executed.
func (c *Computer) Run(maxOperations int) {
	for operationCount := 0; operationCount < maxOperations; operationCount++ {
		c.RunOne()
		if c.Halted {
			return
		}
	}
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	programRaw, err := ioutil.ReadAll(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	program, err := ParseProgram(string(programRaw))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	program[1] = 12
	program[2] = 2

	const maxOperations = 1000 * 1000

	computer := NewComputer(program)
	computer.Run(maxOperations)

	if !computer.Halted {
		fmt.Printf("Did not halt after %v operations, exiting\n", maxOperations)
	}

	fmt.Printf("Value at position 0: %v\n", computer.Program[0])

	const targetOutput = 19690720
	for noun := 0; noun <= 99; noun++ {
		program[1] = noun
		for verb := 0; verb <= 99; verb++ {
			program[2] = verb

			computer := NewComputer(program)
			computer.Run(maxOperations)

			if !computer.Halted {
				fmt.Printf("Did not halt after %v operations, exiting\n", maxOperations)
			}

			output := computer.Program[0]
			if output == targetOutput {
				fmt.Printf("Noun %v and verb %v produce output %v\n", noun, verb, output)
				break
			}
		}
	}
}
