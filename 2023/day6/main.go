package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func countTotalWins(time int, distance int) int {
	count := 0
	for i := 0; i <= time; i++ {
		remainingTime := time - i
		travel := remainingTime * i
		if travel > distance {
			count += 1
		}
	}
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	input := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}
	timeList := input[0]
	distanceList := input[1]
	times := strings.Split(standardizeSpaces(strings.TrimSpace(strings.Split(timeList, ":")[1])), " ")
	distances := strings.Split(standardizeSpaces(strings.TrimSpace(strings.Split(distanceList, ":")[1])), " ")
	newTime := strings.ReplaceAll(strings.TrimSpace(strings.Split(timeList, ":")[1]), " ", "")
	newDistance := strings.ReplaceAll(strings.TrimSpace(strings.Split(distanceList, ":")[1]), " ", "")
	wins := 1
	for idx, time := range times {
		val, _ := strconv.Atoi(time)
		distance, _ := strconv.Atoi(distances[idx])
		count := countTotalWins(val, distance)
		wins *= count
	}
	fmt.Println(wins)
	solTwoTime, _ := strconv.Atoi(newTime)
	solTwoDistance, _ := strconv.Atoi(newDistance)

	fmt.Println(countTotalWins(solTwoTime, solTwoDistance))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
