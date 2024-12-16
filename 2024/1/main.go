package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

// Solution 1: Calculate the total absolute difference
func solution1(leftNumbers, rightNumbers []int) int {
	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	var total float64
	for i, left := range leftNumbers {
		total += math.Abs(float64(left - rightNumbers[i]))
	}

	fmt.Println("Solution 1 Output:", int(total))
	return int(total)
}

// Solution 2: Calculate the weighted sum
func solution2(leftNumbers, rightNumbers []int) int {
	rightCount := make(map[int]int)
	for _, right := range rightNumbers {
		rightCount[right]++
	}

	var total int
	for _, left := range leftNumbers {
		if right, ok := rightCount[left]; ok {
			total += left * right
		}
	}

	fmt.Println("Solution 2 Output:", total)
	return total
}

// Helper function to parse the input file
func parseFile(filePath string) ([]int, []int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
	}
	defer file.Close()

	var leftNumbers, rightNumbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var left, right int
		fmt.Sscanf(scanner.Text(), "%d %d", &left, &right)
		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}

	return leftNumbers, rightNumbers
}

func main() {
	// Parse the input file
	leftNumbers, rightNumbers := parseFile("input.txt")

	// Call both solutions
	_ = solution1(leftNumbers, rightNumbers)
	_ = solution2(leftNumbers, rightNumbers)
}
