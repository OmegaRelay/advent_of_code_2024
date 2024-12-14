package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

const regex = "(do\\(\\))|(don't\\(\\))|(mul\\([0-9]{1,3},[0-9]{1,3}\\))"
const regexMul = "(mul\\([0-9]{1,3},[0-9]{1,3}\\))"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error no input")
	}

	inputFile := os.Args[1]

	if inputFile == "" {
		log.Fatal("error no input")
	}

	instruction, err := getInstructionFromFile(inputFile)
	if err != nil {
		log.Fatalf("could not read input file: %s", err.Error())
	}

	startTime := time.Now().UnixMicro()
	ans := part1(instruction)
	ansTime := time.Now().UnixMicro() - startTime
	fmt.Printf("Part 1: \n\tans: %d\n\tin %dus\n\n", ans, ansTime)

	startTime = time.Now().UnixMicro()
	ans = part2(instruction)
	ansTime = time.Now().UnixMicro() - startTime
	fmt.Printf("Part 2: \n\tans: %d\n\tin %dus\n\n", ans, ansTime)
}

func part1(instruction string) int {
	r, _ := regexp.Compile(regexMul)
	mulInstructions := r.FindAllString(instruction, -1)

	ans := 0
	for _, mul := range mulInstructions {
		var a, b int
		fmt.Sscanf(mul, "mul(%d,%d)", &a, &b)
		ans += a * b
	}
	return ans
}

func part2(instruction string) int {
	r, _ := regexp.Compile(regex)
	instructions := r.FindAllString(instruction, -1)
	mulEnabled := true

	ans := 0
	for _, instruction := range instructions {
		switch instruction {
		case "do()":
			mulEnabled = true
		case "don't()":
			mulEnabled = false
		default:
			// assume mul instruction?
			if !mulEnabled {
				continue
			}
			var a, b int
			fmt.Sscanf(instruction, "mul(%d,%d)", &a, &b)
			ans += a * b
		}
	}
	return ans
}

func getInstructionFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
