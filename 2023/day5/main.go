package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
  var seeds []string
	scanner := bufio.NewScanner(file)
	string := ""
	for scanner.Scan() {
		input := scanner.Text()
		string += input + "\n"
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	for _, line := range strings.Split(string, "\n\n") {
		if strings.Contains(line, "seeds") {
			seeds = strings.Split(strings.Split(line, ": ")[1], " ")
		} else if strings.Contains(line,"seed-to-soil"){
      // fmt.Println(strings.Split(strings.Split(line, ":")[1], "\n"))
    }
	}

  fmt.Println(seeds[1])
}
