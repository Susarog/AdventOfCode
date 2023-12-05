package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalPoints := 0
	totalScratchCards := 0

	m2 := map[int]int{}
	for scanner.Scan() {
		input := scanner.Text()
		splitColon := strings.Split(strings.TrimSpace(input), ":")
    fmt.Println()
		cardNumber, _ := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(splitColon[0], "Card", "")))
		splitCard := strings.Split(strings.TrimSpace(splitColon[1]), " | ")
		winningNumbers := strings.Split(strings.TrimSpace(splitCard[0]), " ")
		userNumbers := strings.Split(strings.TrimSpace(splitCard[1]), " ")

		m := map[string]bool{}
		points := 0
		for _, number := range winningNumbers {
			if number != "" {
				m[string(number)] = true
			}
		}

		for _, number := range userNumbers {
			number = strings.TrimSpace(number)
			if m[string(number)] {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points

		// sol 2
		copyAmount := 0
		for _, number := range winningNumbers {
			if number != "" {
				m[string(number)] = true
			}
		}

		for _, number := range userNumbers {
			number = strings.TrimSpace(number)
			if m[string(number)] {
				copyAmount += 1
			}
		}
		m2[cardNumber] += 1

		for i := 0; i < m2[cardNumber]; i += 1 {
			for num := cardNumber + 1; num <= cardNumber+copyAmount; num += 1 {
				m2[num] += 1
			}
		}
	}
	fmt.Println(m2)
	for _, value := range m2 {
		totalScratchCards += value
	}
	// fmt.Println(totalPoints)
	fmt.Println(totalScratchCards)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
