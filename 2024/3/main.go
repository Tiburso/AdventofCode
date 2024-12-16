package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func validMults(line string) int {
	var total int
	
	re, err := regexp.Compile(`(do(n't)?\(\))|(mul\((\d+),(\d+)\))`)
	if err != nil {
		log.Fatal(err)
	}

	shouldSkip := false
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		action := match[0]

		if action == "do()" {
			shouldSkip = false
			continue
		} else if action == "don't()" {
			shouldSkip = true
			continue
		}

		if shouldSkip {
			continue
		}

		a, err := strconv.Atoi(match[4])
		if err != nil {
			log.Fatal(err)
		}
		
		b, err := strconv.Atoi(match[5])
		if err != nil {
			log.Fatal(err)
		}

		total += a * b
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineTotal string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineTotal += line
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := validMults(lineTotal)

	fmt.Println(total)
}