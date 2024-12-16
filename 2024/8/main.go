package main

import (
	"bufio"
	"os"
)

type Position struct {
	X int
	Y int
}

func isAntinodeValid(antinode Position, width, height int) bool {
	return antinode.X >= 0 && antinode.X < width && antinode.Y >= 0 && antinode.Y < height
}

// func debugAntinodes(antinodes map[Position]bool, width, height int) {
// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			pos := Position{x, y}
// 			if antinodes[pos] {
// 				print("#")
// 			} else {
// 				print(".")
// 			}
// 		}
// 		println()
// 	}
// }

func findResonant(antinodes map[Position]bool, position, other Position, width, height int) {
	// Find resonant works like the first but works for every dx, dy combination
	dx := other.X - position.X
	dy := other.Y - position.Y

	// Find the resonant position
	for {
		position.X += dx
		position.Y += dy

		if !isAntinodeValid(position, width, height) {
			break
		}

		antinodes[position] = true
	}

	for {
		other.X -= dx
		other.Y -= dy

		if !isAntinodeValid(other, width, height) {
			break
		}

		antinodes[other] = true
	}
}

func findAntiNodes(antennas map[rune][]Position, width, height int, resonant bool) int {
	antinodes := make(map[Position]bool)

	for _, positions := range antennas {
		size := len(positions)
		for i, pos := range positions {
			for j := i + 1; j < size; j++ {
				other := positions[j]
				dx := other.X - pos.X
				dy := other.Y - pos.Y

				antinode := Position{other.X + dx, other.Y + dy}
				opposite := Position{pos.X - dx, pos.Y - dy}

				if !resonant {
					if isAntinodeValid(antinode, width, height) {
						antinodes[antinode] = true
					}
	
					if isAntinodeValid(opposite, width, height) {
						antinodes[opposite] = true
					}
				} else {
					findResonant(antinodes, pos, other, width, height)
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	antennas := make(map[rune][]Position)

	scanner := bufio.NewScanner(file)
	var width, height int
	for scanner.Scan() {
		line := scanner.Text()
		
		width = len(line)
		for x, c := range line {
			if c == '.' {
				continue
			}
			pos := Position{x, height}
			antennas[c] = append(antennas[c], pos)
		}
		height++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	total := findAntiNodes(antennas, width, height, true)
	println(total)
}