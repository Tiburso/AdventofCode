package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isLevelSafe(numbers []int) bool {
	increasing := numbers[0] < numbers[1]

	for i := 1; i < len(numbers); i++ {
		prev, curr := numbers[i-1], numbers[i]

		if curr == prev || math.Abs(float64(curr-prev)) > 3 {
			return false
		}

		if increasing && curr < prev || !increasing && curr > prev {
			return false
		}
	}	

	return true
}

func tryLevelwithOneLess(numbers []int, index int) bool {
	newNumbers := make([]int, 0, len(numbers)-1)
	newNumbers = append(newNumbers, numbers[:index]...)
	newNumbers = append(newNumbers, numbers[index+1:]...)

	return isLevelSafe(newNumbers)
}

func loopLevel(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if tryLevelwithOneLess(numbers, i) {
			return true
		}
	}

	return false
}

func convertLine(line string) []int {
	numbers := make([]int, 0, len(strings.Fields(line)))

	for _, number := range strings.Fields(line) {
		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatalf("Error converting number: %v", err)
		}
		numbers = append(numbers, n)
	}

	return numbers
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := convertLine(line)

		if loopLevel(numbers) {
			total++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Printf("Total: %d\n", total)
}
