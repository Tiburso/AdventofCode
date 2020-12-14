#! /usr/bin/env python3

from switch import Switch

with open("input.txt", "r") as f:
	raw = f.read().splitlines()

d = dict(enumerate(raw))

def fix():
	i = 0
	accumulator = 0
	traversed = []

	while i < len(d):
		l = d[i].split(" ")
		command = l[0]
		value = eval(l[1])

		with Switch(command) as case:
			if case("nop"):
				traversed.append(i)
				i += 1
			if case("acc"):
				accumulator += value
				traversed.append(i)
				i += 1
			if case("jmp"):
				traversed.append(i)
				i += value

print(accumulator)	


