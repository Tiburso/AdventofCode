#! /usr/bin/env python3

nomes = [n for n in open("input.txt", "r")]

final = 0
for nome in nomes:
	split = nome.split(" ")
	inf, sup = split[0].split("-")
	target = split[1][0]
	word = split[2]

	count = 0
	for l in word:
		if l == target:
			count += 1

	if count >= eval(inf) and count <= eval(sup):
		final += 1

print(final)


