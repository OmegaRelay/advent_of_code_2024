// Part 1: complete
// Part 2: todo
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type rule struct {
	first int
	last  int
}
type rules []rule
type update []int

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error no input")
	}

	inputFile := os.Args[1]

	if inputFile == "" {
		log.Fatal("error no input")
	}

	rules, updates, err := getRulesAndUpdatesFromFile(inputFile)
	if err != nil {
		log.Fatalf("could not read input file: %s", err.Error())
	}

	startTime := time.Now().UnixMicro()
	ans := part1(rules, updates)
	ansTime := time.Now().UnixMicro() - startTime
	fmt.Printf("Part 1: \n\tans: %d\n\tin %dus\n\n", ans, ansTime)

}

func part1(r rules, us []update) int {
	ret := 0
	for _, update := range us {
		if r.isCompliant(update) {
			fmt.Println(update)
			ret += update[len(update)/2]
		}
	}
	return ret
}

func (rs rules) isCompliant(u update) bool {
	for _, rule := range rs {
		firstFound := false
		if !rule.doesApply(u) {
			continue
		}

		for _, update := range u {
			if update == rule.first {
				firstFound = true
			}
			if update == rule.last && !firstFound {
				return false
			}
		}
	}
	return true
}

func (r rule) doesApply(u update) bool {
	nrOccurances := 0
	for _, update := range u {
		if update == r.first || update == r.last {
			nrOccurances++
		}
	}
	return !(nrOccurances < 2)
}

func getRulesAndUpdatesFromFile(filePath string) (rules, []update, error) {
	rules := rules{}
	updates := []update{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	splitStr := strings.Split(string(data), "\n\n")

	rulesStr := splitStr[0]
	for _, ruleStr := range strings.Split(rulesStr, "\n") {
		firstLastStr := strings.Split(ruleStr, "|")

		first, err := strconv.ParseInt(firstLastStr[0], 10, 32)
		if err != nil {
			return nil, nil, err
		}

		last, err := strconv.ParseInt(firstLastStr[1], 10, 32)
		if err != nil {
			return nil, nil, err
		}

		rules = append(rules, rule{
			first: int(first),
			last:  int(last),
		})
	}

	updatesStr := splitStr[1]
	for _, updateStr := range strings.Split(updatesStr, "\n") {
		update := update{}
		for _, updateValStr := range strings.Split(updateStr, ",") {
			updateVal, err := strconv.ParseInt(updateValStr, 10, 32)
			if err != nil {
				return nil, nil, err
			}
			update = append(update, int(updateVal))
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}
