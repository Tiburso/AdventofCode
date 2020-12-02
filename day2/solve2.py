#! /usr/bin/env python3

nomes = [n[:-1] for n in open("input.txt", "r")]

final = 0
for nome in nomes:
	split = nome.split(" ")
	pos1, pos2 = split[0].split("-")
	target = split[1][0]
	word = split[2]

	if (word[eval(pos1)-1] == target) != (word[eval(pos2)-1] == target):
		final += 1


print(final)


