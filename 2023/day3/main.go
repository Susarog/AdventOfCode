package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type location struct {
	x int
	y int
}

type enginePart struct {
	partNumber string
	coordinate location
}

func atRange(symbol location, part enginePart) int {
	x := part.coordinate.x
	y := part.coordinate.y
	xMax := part.coordinate.x + len(part.partNumber) - 1
	symbolStartX := symbol.x - 1
	symbolStartY := symbol.y - 1

	for i := 0; i < 3; i += 1 {
		for j := 0; j < 3; j += 1 {
			if symbolStartY+i == y && symbolStartX+j >= x && symbolStartX+j <= xMax {
				val, _ := strconv.Atoi(part.partNumber)
				return val
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	y := -1
	symbols := make([]location, 0)
	parts := make([]enginePart, 0)
	for scanner.Scan() {
		input := scanner.Text()
		value := ""
		x := 0
		y += 1
		isFirstDigit := true
		for idx, v := range input {
			if unicode.IsDigit(rune(v)) {
				if isFirstDigit {
					x = idx
					isFirstDigit = false
				}
				value += string(v)
				if idx == len(input)-1 {
					parts = append(parts, enginePart{value, location{x, y}})
					isFirstDigit = true
					value = ""
				}
			} else {
				if value != "" {
					parts = append(parts, enginePart{value, location{x, y}})
					isFirstDigit = true
					value = ""
				}
				if string(v) != "." {
					symbols = append(symbols, location{x: idx, y: y})
				}
			}
		}
	}

	// solution 1
	total := 0
	for _, symbolCoordinate := range symbols {
		for _, enginePart := range parts {
			total += atRange(symbolCoordinate, enginePart)
		}
	}
	fmt.Println(total)
	// solution 2
	total = 0
	for _, symbolCoordinate := range symbols {
		gearParts := make([]int, 0)
		for _, enginePart := range parts {
			if val := atRange(symbolCoordinate, enginePart); val != 0 {
				gearParts = append(gearParts, val)
			}
		}
		if len(gearParts) == 2 {
			total += gearParts[0] * gearParts[1]
		}
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
