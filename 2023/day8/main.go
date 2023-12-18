package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type element struct {
	left  string
	right string
}

func checkAtEndPath(node string) bool {
	return node[len(node)-1] != 'Z'
}

func gcd(a int, b int) int {
	// Euclidean algorithm
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func lcm(a int, b int) int {
	return (a * b / gcd(a, b))
}

func lcmm(args []int) int {
	// Recursively iterate through pairs of arguments
	// i.e. lcm(args[0], lcm(args[1], lcm(args[2], args[3])))

	if len(args) == 2 {
		return lcm(args[0], args[1])
	} else {
		arg0 := args[0]
		shift := args[1:]
		return lcm(arg0, lcmm(shift))
	}
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
		// remove empty line
		if len(text) != 0 {
			input = append(input, text)
		}
	}
	instructions := input[0]
	startingNodes := make([]string, 0)
	network := make(map[string]element)
	for i := 1; i < len(input); i++ {
		nodes := strings.Split(input[i], " = ")
		currNode := nodes[0]
		test := strings.NewReplacer("(", "", ")", "")
		elements := strings.Split(test.Replace(nodes[1]), ", ")
		leftElement := elements[0]
		rightElement := elements[1]
		network[currNode] = element{leftElement, rightElement}

		if currNode[len(currNode)-1] == 'A' {
			startingNodes = append(startingNodes, currNode)
		}
	}

	// first solution
	// currNode := "AAA"
	// steps := 0
	// for currNode != "ZZZ" {
	// 	for _, instruction := range instructions {
	// 		steps += 1
	// 		if instruction == 'R' {
	// 			currNode = network[currNode].right
	// 		} else {
	// 			currNode = network[currNode].left
	// 		}
	// 	}
	// }
	// fmt.Println(steps)

	// fmt.Println(startingNodes)
	// brute forcey. it will work eventually but i will need to find a faster solution
	nodeSteps := make([]int, 0)
	for _, node := range startingNodes {
		tempNode := node
		steps := 0
		for checkAtEndPath(tempNode) {
			for _, instruction := range instructions {
				steps += 1
				if instruction == 'R' {
					tempNode = network[tempNode].right
				} else {
					tempNode = network[tempNode].left
				}
			}
		}
		// check if it reaches Z
		nodeSteps = append(nodeSteps, steps)
	}

	fmt.Println(lcmm(nodeSteps))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
