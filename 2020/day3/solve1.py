#! /usr/bin/env python3

import numpy

#text = f.read().splitlines()

map = [n[:-1] for n in open("input.txt","r")]

def going_down(left, down):
	x = 0
	y = 0

	trees = 0
	while y < len(map):
		if map[y][x] == "#":
			trees += 1

		x = (x + left) % len(map[0])
		y += down

	return trees

result = [going_down(1,1), going_down(3,1), going_down(5,1), going_down(7,1), going_down(1,2)]

print(result)

print(numpy.prod(result))
