package main

import (
	"bufio"
	"fmt"
	"os"
)

const INVALID = -1
type Count struct {
	StartIndex int
	Occurences int
	Value int 
}

func (c Count) String() string {
	return fmt.Sprintf("StartIndex: %d, Occurences: %d, Value: %d;", c.StartIndex, c.Occurences, c.Value)
}

func decode(input string) []int {
	free := false
	var index int
	var decoded []int

	for _, r := range input {
		// Convert rune to int
		count := int(r - '0')
		
		if free {
			for i := 0; i < count; i++ {
				decoded = append(decoded, -1)
			}
			free = false
		} else {
			for i := 0; i < count; i++ {
				decoded = append(decoded, index)
			}
			index++
			free = true
		}
	}

	return decoded
}

func organize(input []int) []int {
	for start, end := 0, len(input); start < end; {
		// while start is not "."
		for start < end && input[start] != -1 {
			start++
		}

		// while end is "."
		for start < end && input[end-1] == -1 {
			end--
		}

		// swap between start and end
		if start < end {
			input[start], input[end-1] = input[end-1], input[start]
			start++
			end--
		}
	}

	return input
}

func calculate(input []int) int {
	var total int
	for i, v := range input {
		if v == -1 {
			continue
		}

		total += v * i
	}
	return total
}

func decodeCount(input string) []Count {
	var counts []Count
	free := false
	value := 0

	for i := 0; i < len(input); i++{
		count := int(input[i] - '0')
		if count == 0 {
			free = !free
			continue
		}

		if free {
			counts = append(counts, Count{StartIndex: i, Occurences: count, Value: INVALID})
		} else {
			counts = append(counts, Count{StartIndex: i, Occurences: count, Value: value})
			value++
		}

		free = !free
	}

	return counts
}

func organizeCounts(input []Count) []Count {
	var organized []Count
	
	for start, end := 0, len(input); start < end; {
		for start < end && input[start].Value != INVALID {
			organized = append(organized, input[start])
			start++
		}

		for start < end && input[end-1].Value == INVALID {
			end--
		}

		// Now starting on the end look at the count which has capacity to be moved to the start
		for i := end - 1; i >= start; i-- {
			if input[i].Value != INVALID && input[i].Occurences <= input[start].Occurences {
				organized = append(organized, input[i])
				input[i].Value = INVALID
				input[start].Occurences -= input[i].Occurences
			}
		}

		if input[start].Occurences != 0 {
			organized = append(organized, input[start])
		}
		start++
	}

	return organized
}

func convertToIntArray(input []Count) []int {
	var output []int
	for _, c := range input {
		for i := 0; i < c.Occurences; i++ {
			output = append(output, c.Value)
		}
	}
	return output
}

func part1(input string) int {
	output := decode(input)
	organized := organize(output)
	return calculate(organized)
}

func part2(input string) int {
	output := decodeCount(input)
	organized := organizeCounts(output)
	intArr := convertToIntArray(organized)

	return calculate(intArr)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	total := part1(input)
	fmt.Println(total)

	total = part2(input)
	fmt.Println(total)
}
