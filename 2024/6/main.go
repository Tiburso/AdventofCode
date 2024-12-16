package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction int
const (
	up direction = iota
	right
	down
	left
)

type point struct {
	x int
	y int
}

type Guard struct {
	x int
	y int
	dir direction 
}


var d = []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func solve(maze [][]rune, guard Guard) (int, map[point]bool) {
	var total int
	visited := make(map[point]bool)
	width, height := len(maze), len(maze[0])

	for {
		curr := point{x: guard.x, y: guard.y}
		if _, ok := visited[curr]; !ok {
			visited[curr] = true
			total++
		}

		if guard.x <= 0 || guard.x >= width-1 || guard.y <= 0 || guard.y >= height-1 {
			break
		}

		newPoint := point{x: guard.x + d[guard.dir].x, y: guard.y + d[guard.dir].y}
		if maze[newPoint.x][newPoint.y] == '#' {
			guard.dir = (guard.dir + 1) % 4
		} else {
			guard.x, guard.y = newPoint.x, newPoint.y
		}
	}

	return total, visited
}

func isLoop(maze [][]rune, guard Guard) bool {
	width, height := len(maze), len(maze[0])

	visited := make(map[Guard]bool)
	for {
		if guard.x <= 0 || guard.x >= width-1 || guard.y <= 0 || guard.y >= height-1 {
			break
		}

		newX, newY := guard.x+d[guard.dir].x, guard.y+d[guard.dir].y
		if maze[newX][newY] == '#' {
			guard.dir = (guard.dir + 1) % 4
		} else {
			guard.x, guard.y = newX, newY
		}

		if visited[guard] {
			return true
		}

		visited[guard] = true
	}

	return false
}

func copyMaze(maze [][]rune) [][]rune {
	copyMaze := make([][]rune, len(maze))
	for i, row := range maze {
		copyMaze[i] = make([]rune, len(row))
		copy(copyMaze[i], row)
	}

	return copyMaze
}

func setLoops(maze [][]rune, guard Guard, visited map[point]bool) int {
	var total int
	originalMaze := copyMaze(maze)
	for p := range visited {
		maze := copyMaze(originalMaze)
	
		if p.x == guard.x && p.y == guard.y {
			continue
		}

		maze[p.x][p.y] = '#'
		if isLoop(maze, guard) {
			total++
		}
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var maze [][]rune
	var guard Guard

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, []rune(line))

		for j, c := range line {
			if c == '^' {
				guard = Guard{x: i, y: j, dir: up}
			}
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(solve(maze, guard))
	_, visited := solve(maze, guard)
	fmt.Println(setLoops(maze, guard, visited))
}