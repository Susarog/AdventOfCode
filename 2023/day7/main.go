package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	High int = iota
	One
	Two
	Three
	Full
	Four
	Five
)

type cardHand struct {
	cards []rune
	bid   int
}

type allCards []cardHand

func (s allCards) Less(i, j int) bool {
	m := make(map[rune]int)
	m['2'] = 2
	m['3'] = 3
	m['4'] = 4
	m['5'] = 5
	m['6'] = 6
	m['7'] = 7
	m['8'] = 8
	m['9'] = 9
	m['T'] = 10
	m['J'] = 1
	m['Q'] = 12
	m['K'] = 13
	m['A'] = 14
  // remove convertWildCard to get first solution
	leftCardType := getCardType(convertWildCard(s[i].cards))
	rightCardType := getCardType(convertWildCard(s[j].cards))

	if leftCardType != rightCardType {
		return leftCardType < rightCardType
	}

	for idx := 0; idx < len(s[i].cards); idx++ {
		leftCard := s[i].cards[idx]
		rightCard := s[j].cards[idx]
		if leftCard != rightCard {
			return m[leftCard] < m[rightCard]
		}
	}
	return true
}

func (s allCards) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s allCards) Len() int {
	return len(s)
}

func convertWildCard(s []rune) []rune {
	charCounter := make(map[rune]int, 0)
	for _, char := range s {
		if char != 'J' {
			charCounter[char] += 1
		}
	}
	highest := -1
	var bestCard rune
	for k, v := range charCounter {
		if highest < v {
			highest = v
			bestCard = k
		}
	}
	test := string(s)
	return []rune(strings.ReplaceAll(test, "J", string(bestCard)))
}

func getCardType(s []rune) int {
	charCounter := make(map[rune]int, 0)
	for _, char := range s {
		charCounter[char] += 1
	}
	highest := -1
	pair := 0
	for _, v := range charCounter {
		if highest < v {
			highest = v
		}
		if v == 2 {
			pair += 1
		}
	}
	switch highest {
	case 1:
		return High
	case 2:
		if pair == 1 {
			return One
		} else {
			return Two
		}
	case 3:
		if pair == 1 {
			return Full
		}
		return Three
	case 4:
		return Four
	default:
		return Five
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	allCards := make(allCards, 0)
	for scanner.Scan() {
		text := scanner.Text()
		cards := strings.Split(text, " ")[0]
		bidNumber, _ := strconv.Atoi(strings.Split(text, " ")[1])
		card := cardHand{[]rune(cards), bidNumber}
		allCards = append(allCards, card)
	}

	sort.Sort(allCards)
	total := 0
	for idx, card := range allCards {
		total += card.bid * (idx + 1)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
