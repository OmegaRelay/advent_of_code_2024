package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type list []int

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error no input")
	}

	inputFile := os.Args[1]

	if inputFile == "" {
		log.Fatal("error no input")
	}

	list1, list2, err := getListsFromFile(inputFile)
	if err != nil {
		log.Fatalf("could not read input file: %s", err.Error())
	}

	startTime := time.Now().UnixMicro()
	part1Ans := part1(list1, list2)
	part1Time := time.Now().UnixMicro() - startTime

	startTime = time.Now().UnixMicro()
	part2Ans := part2(list1, list2)
	part2Time := time.Now().UnixMicro() - startTime

	fmt.Printf("Part 1: \n\tans: %d\n\tin %dus\n\n", part1Ans, part1Time)
	fmt.Printf("Part 2: \n\tans: %d\n\tin %dus\n\n", part2Ans, part2Time)
}

func getListsFromFile(filePath string) (list, list, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	values := strings.Fields(string(data))
	list1 := list{}
	list2 := list{}
	for i, v := range values {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, nil, err
		}
		if i%2 == 0 {
			list1 = append(list1, int(value))
		} else {
			list2 = append(list2, int(value))
		}
	}

	return list1, list2, nil
}

func part1(list1 list, list2 list) int {
	list1.sortSmallToLarge()
	list2.sortSmallToLarge()

	distanceList := list{}
	if len(list1) <= len(list2) {
		for i := range list1 {
			if list1[i] < list2[i] {
				distanceList = append(distanceList, (list2[i] - list1[i]))
			} else {
				distanceList = append(distanceList, (list1[i] - list2[i]))
			}
		}
	} else {
		for i := range list2 {
			if list1[i] < list2[i] {
				distanceList = append(distanceList, (list2[i] - list1[i]))
			} else {
				distanceList = append(distanceList, (list1[i] - list2[i]))
			}
		}
	}

	return distanceList.sum()
}

func part2(list1 list, list2 list) int {
	locationToCount := make(map[int]int) // maintain a map of already counted values
	scoreList := list{}

	for _, v := range list1 {
		count := 0
		count, isFound := locationToCount[v]
		if isFound {
			scoreList = append(scoreList, v*count)
			continue
		}

		for _, w := range list2 {
			if v == w {
				count++
			}
		}
		locationToCount[v] = count
		scoreList = append(scoreList, v*count)
	}

	return scoreList.sum()
}

func (l list) sortSmallToLarge() {
	len := len(l)
	for i := 1; i < len; i++ {
		if l[i-1] < l[i] {
			continue
		}

		for j := i; j > 0; j-- {
			if l[j-1] < l[j] {
				break
			}

			tmp := l[j]
			l[j] = l[j-1]
			l[j-1] = tmp
		}
	}
}

func (l list) sum() int {
	s := 0
	for _, v := range l {
		s += v
	}
	return s
}
