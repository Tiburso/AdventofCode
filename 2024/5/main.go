package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data	int
	children_set map[int]bool
}

func createGraph(globalMap map[int]*Node, parent, child int) {
	if _, ok := globalMap[parent]; !ok {
		globalMap[parent] = &Node{data: parent, children_set: make(map[int]bool)}
	}
	
	if _, ok := globalMap[child]; !ok {
		globalMap[child] = &Node{data: child, children_set: make(map[int]bool)}
	}

	globalMap[parent].children_set[child] = true
}

func validSequence(globalMap map[int]*Node, sequence []string) (bool, []string) {
	size := len(sequence)
	valid := true
	
	for i := 0; i < size; i++{
		for j := i+1; j < size; j++ {
			m, err := strconv.Atoi(sequence[j])
			if err != nil {
				log.Fatal(err)
			}
			
			n, err := strconv.Atoi(sequence[i])
			if err != nil {
				log.Fatal(err)
			}

			if _, ok := globalMap[n].children_set[m]; !ok {
				valid = false
				sequence[i], sequence[j] = sequence[j], sequence[i]
			}
		}
	}

	return valid, sequence
}

func main() {
	file, err := os.Open("input.txt")	
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	globalMap := make(map[int]*Node)

	scanner := bufio.NewScanner(file)
	
	// Part 1 should generate the graph
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		// Split the line into parent and child
		var parent, child int
		fmt.Sscanf(line, "%d|%d\n", &parent, &child)

		createGraph(globalMap, parent, child)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 2 should validate each sequence
	var total int
	for scanner.Scan() {
		sequence := strings.Split(scanner.Text(), ",")
	
		valid, sequence := validSequence(globalMap, sequence)
		if !valid {
			element := sequence[len(sequence) / 2]
			n, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}

			total += n
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}