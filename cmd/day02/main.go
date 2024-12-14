// Part A: Success
// Part B: Fail
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type level int
type report []level

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error no input")
	}

	inputFile := os.Args[1]

	if inputFile == "" {
		log.Fatal("error no input")
	}

	reports, err := getReportsFromFile(inputFile)
	if err != nil {
		log.Fatalf("could not read input file: %s", err.Error())
	}

	startTime := time.Now().UnixMicro()
	reportsCpy := make([]report, len(reports))
	copy(reportsCpy, reports)
	ans := part1(reportsCpy)
	ansTime := time.Now().UnixMicro() - startTime
	fmt.Printf("Part 1: \n\tans: %d\n\tin %dus\n\n", ans, ansTime)

	startTime = time.Now().UnixMicro()
	ans = part2(reports)
	ansTime = time.Now().UnixMicro() - startTime
	fmt.Printf("Part 2: \n\tans: %d\n\tin %dus\n\n", ans, ansTime)
}

func part1(reports []report) int {
	nrSafeReports := 0
	for _, report := range reports {
		if report.isSafe(false) {
			nrSafeReports++
		}
	}
	return nrSafeReports
}

func part2(reports []report) int {
	nrSafeReports := 0
	for _, report := range reports {
		if report.isSafe(true) {
			nrSafeReports++
		}
	}
	return nrSafeReports
}

func getReportsFromFile(filePath string) ([]report, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	reports := []report{}

	resportStrings := strings.Split(string(data), "\n")
	for _, reportString := range resportStrings {
		report := report{}
		levels := strings.Split(reportString, " ")
		if reportString == "" {
			continue
		}
		for _, levelString := range levels {
			lvl, err := strconv.ParseInt(levelString, 10, 32)
			if err != nil {
				return nil, err
			}
			report = append(report, level(lvl))
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func (r report) isSafe(hasDampener bool) bool {
	nrBadLevels := 0
	maxBadLevels := 0
	if hasDampener {
		maxBadLevels = 1
	}

	i := 0
	isIncrementing, err := r[i].isIncrementing(r[i+1])
	for err != nil {
		r = r.removeLevel(i)
		nrBadLevels++
		if nrBadLevels > maxBadLevels {
			return false
		}
		isIncrementing, err = r[i].isIncrementing(r[i+1])
	}

	for i := 0; i < len(r)-1; i++ {
		ret := r.checkLevel(isIncrementing, i)
		if !ret {
			r = r.removeLevel(i)
			i--
			nrBadLevels++
		}

		if nrBadLevels > maxBadLevels {
			return false
		}
	}

	return true
}

func (r report) checkLevel(isIncrementing bool, i int) bool {
	diff := r[i+1] - r[i]

	if diff > 0 {
		if !isIncrementing {
			return false
		}
	} else if diff < 0 {
		if isIncrementing {
			return false
		}
	} else {
		return false
	}

	if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
		return false
	}

	return true
}

func (r report) removeLevel(i int) report {
	slice := r
	for ; i < len(slice)-1; i++ {
		slice[i] = slice[i+1]
	}
	return slice[:(len(slice) - 1)]
}

func (l level) isIncrementing(comp level) (bool, error) {
	diff := comp - l
	if diff > 0 {
		return true, nil
	} else if diff < 0 {
		return false, nil
	} else {
		return false, errors.New("neither")
	}
}
