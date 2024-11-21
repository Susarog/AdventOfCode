package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type coordinate struct {
	x, y int
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PrintSlice[T any](slice []T) {
	fmt.Println(slice)
}

func transpose[T any](matrix [][]T) [][]T {
	row := len(matrix)
	col := len(matrix[0])

	m := make([][]T, col)
	for i := range m {
		m[i] = make([]T, row)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[j][i] = matrix[i][j]
		}
	}
	return m
}

func getNonGalaxyRow(matrix [][]string) []int {
	m := make([]int, 0)
	for i, row := range matrix {
		if !slices.Contains(row, "#") {
			m = append(m, i)
		}
	}
	return m
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	image := make([][]string, 0)

	// image contains the image input from the problem
	i := 0
	for scanner.Scan() {
		input := scanner.Text()
		row := strings.Split(input, "")
		image = append(image, row)
		i += 1
	}
	nonGalaxyRowsIndex := getNonGalaxyRow(image)
	nonGalaxyColumnIndex := getNonGalaxyRow(transpose(image))

	coordinates := make([]coordinate, 0)
	for y, row := range image {
		for x, val := range row {
			if val == "#" {
				coordinates = append(coordinates, coordinate{x: x, y: y})
			}
		}
	}

	expandedCoordinates := make([]coordinate, len(coordinates))
	copy(expandedCoordinates, coordinates)

	for _, idx := range nonGalaxyRowsIndex {
		for i := range coordinates {
			if coordinates[i].y > idx {
				expandedCoordinates[i].y += 999999
			}
		}
	}

	for _, idx := range nonGalaxyColumnIndex {
		for i := range coordinates {
			if coordinates[i].x > idx {
				expandedCoordinates[i].x += 999999
			}
		}
	}

	shortestPathAmt := 0
	for i := 0; i < len(expandedCoordinates); i++ {
		fromCoord := expandedCoordinates[i]
		for j := i + 1; j < len(expandedCoordinates); j++ {
			currCoord := expandedCoordinates[j]
			xDist := absInt(fromCoord.x - currCoord.x)
			yDist := absInt(fromCoord.y - currCoord.y)
			shortestPathAmt += xDist + yDist
		}
	}
	fmt.Println(shortestPathAmt)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
