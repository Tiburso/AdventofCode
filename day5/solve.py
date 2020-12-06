#! /usr/bin/env python3

with open("input.txt", "r") as f:
	input = f.read().splitlines()

total = []
for lugar in input:
	row = 127
	col = 7
	
	min = 0
	for rows in lugar[:7]:
		if rows == "F":
			row -= round((row-min) / 2)
		elif rows == "B":
			min += (row-min)//2
	
	if lugar[7] == "F":
		row = min
	
	min = 0
	for cols in lugar[7:]:
		if cols == "L":
			col -= round((col-min) / 2)
		elif cols == "R":
			min += round((col-min)/2)
	
	if lugar[-1] == "L":
		col = min
	
	id = row * 8 + col
	total.append(id)

org = sorted(total)
for i in range(len(org)):
	if org[i]+1 != org[i+1]:
		print(org[i]+1)
		break

	



