package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(name string) (letters [][]rune) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		letters = append(letters, []rune(line))
	}

	return
}

func solution1(letters [][]rune) (total int) {
	goal := "XMAS"
	reversed := "SAMX"

	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]) - 3; j++ {
			curr := string(letters[i][j:j+4])
			if curr == goal || curr == reversed {
				total++
			}
		}
	}

	// Scan vertically
	for i := 0; i < len(letters) - 3; i++ {
		for j := 0; j < len(letters[i]); j++ {
			curr := string([]rune{letters[i][j], letters[i+1][j], letters[i+2][j], letters[i+3][j]})
			if curr == goal || curr == reversed {
				total++
			}
		}
	}

	// Scan diagonally to the right
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			if i+3 < len(letters) && j+3 < len(letters[i]) {
				curr := string([]rune{letters[i][j], letters[i+1][j+1], letters[i+2][j+2], letters[i+3][j+3]})
				if curr == goal || curr == reversed {
					total++
				}
			}
		}
	}

	// Scan diagonally to the left
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			if i+3 < len(letters) && j-3 >= 0 {
				curr := string([]rune{letters[i][j], letters[i+1][j-1], letters[i+2][j-2], letters[i+3][j-3]})
				if curr == goal || curr == reversed {
					total++
				}
			}
		}
	}

	return
}

func solution2(letters [][]rune) (total int) {
	goal := "MAS"
	reversed := "SAM"

	for i := 1; i < len(letters) - 1; i++ {
		for j := 1; j < len(letters[i]) - 1; j++ {
			if letters[i][j] == 'A' {
				// Then find the diag left and right where A is the center
				diag_right := string([]rune{letters[i-1][j-1], letters[i][j], letters[i+1][j+1]})
				diag_left := string([]rune{letters[i-1][j+1], letters[i][j], letters[i+1][j-1]})

				if (diag_right == goal || diag_right == reversed) && (diag_left == goal || diag_left == reversed) {
					total++
				}
			}
		}
	}

	return
}

func main() {
	letters := readFile("input.txt")

	// total := solution1(letters)
	total := solution2(letters)

	fmt.Println(total)
}