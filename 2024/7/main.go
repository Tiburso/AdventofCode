package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concatNumbers(a, b int) int {
	strA, strB := strconv.Itoa(a), strconv.Itoa(b)
	str := strA + strB
	
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return result
}

func isSolvableAux(objective int64, items []int, index int, sum int64) bool {
	if index == len(items) {
		return sum == objective
	}

	return isSolvableAux(objective, items, index+1, sum+int64(items[index])) || isSolvableAux(objective, items, index+1, sum*int64(items[index]))
}

func isSolvable(objective int64, items []int) bool {
	return isSolvableAux(objective, items, 1, int64(items[0]))
}

func isSolvable2(objective int64, items []int) bool {
	if isSolvableAux(objective, items, 1, int64(items[0])) {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var total int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var objective int64
		var items []int
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			panic("invalid input")
		}

		objective, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}

		for _, item := range strings.Fields(parts[1]) {
			value, err := strconv.Atoi(strings.TrimSpace(item))
			if err != nil {
				panic(err)
			}
			items = append(items, value)
		}

		if isSolvable(objective, items) {
			total += objective
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}