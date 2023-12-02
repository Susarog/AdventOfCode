package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	RED_MAX   = 12
	GREEN_MAX = 13
	BLUE_MAX  = 14
)

func solOne(input string) int {
	total := 0
	splitColon := strings.Split(strings.TrimSpace(input), ":")
	gameNumber, _ := strconv.Atoi(strings.Split(splitColon[0], " ")[1])

	gameSets := strings.Split(strings.TrimSpace(splitColon[1]), ";")

	m := make(map[string]int)
	isNotValidSet := false
	for _, set := range gameSets {
		if isNotValidSet {
			break
		}
		m["red"] = 0
		m["blue"] = 0
		m["green"] = 0

		cubesInfo := strings.Split(strings.TrimSpace(set), ", ")
		for _, cubeInfo := range cubesInfo {
			cube := strings.Split(cubeInfo, " ")
			color := string(cube[1])
			amount := string(cube[0])
			val, _ := strconv.Atoi(amount)
			m[color] += val
		}
		for k, v := range m {
			isNotValidSet = ((k == "red" && v > RED_MAX) || (k == "blue" && v > BLUE_MAX) || (k == "green" && v > GREEN_MAX))
			if isNotValidSet {
				total -= gameNumber
				break
			}
		}
	}
	total += gameNumber
	return total
}

func solTwo(input string) int {
	total := 1
	splitColon := strings.Split(strings.TrimSpace(input), ":")
	gameSets := strings.Split(strings.TrimSpace(splitColon[1]), ";")

	m := make(map[string]int)
	m["red"] = 0
	m["blue"] = 0
	m["green"] = 0
	isNotValidSet := false
	for _, set := range gameSets {
		tempMap := make(map[string]int)
		tempMap["red"] = 0
		tempMap["blue"] = 0
		tempMap["green"] = 0

		if isNotValidSet {
			break
		}
		cubesInfo := strings.Split(strings.TrimSpace(set), ", ")
		for _, cubeInfo := range cubesInfo {
			cube := strings.Split(cubeInfo, " ")
			color := string(cube[1])
			amount := string(cube[0])
			val, _ := strconv.Atoi(amount)
			tempMap[color] += val
		}
		for k, v := range m {
			if v < tempMap[k] {
				m[k] = tempMap[k]
			}
		}
	}
	for _, v := range m {
		total *= v
	}
	return total
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		input := scanner.Text()
		total += solTwo(input)
		// total += solOne(input)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
