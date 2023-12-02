package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func getCalibrationValues(input string) int{
	m := make(map[string]string)
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"
	m["four"] = "4"
	m["five"] = "5"
	m["six"] = "6"
	m["seven"] = "7"
	m["eight"] = "8"
	m["nine"] = "9"
	var strings []string
	for i := 0; i < len(input); i += 1 {
		if unicode.IsDigit(rune(input[i])) {
			strings = append(strings, string(input[i]))
		} else {
			for k, v := range m {
				if i+len(k) > len(input) {
					continue
				}
				if input[i:i+len(k)] == k {
					strings = append(strings, v)
					break
				}
			}
		}
	}

	convertedInput := strings[0] + strings[len(strings)-1]
	val, _ := strconv.Atoi(convertedInput)
	return val
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
		total += getCalibrationValues(input)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
