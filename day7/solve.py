#! /usr/bin/env python3

from re import *

with open("input.txt", "r") as f:
	raw = f.read()

def parse_raw():
	bags = {}
	bag_description = compile(r"([a-z]+ [a-z]+) bags contain (.+)")
	formula = compile(r"(\d+) ([a-z]+ [a-z]+) bag")
	for bag, contents_raw in bag_description.findall(raw):
		contents = {inner[1]: eval(inner[0]) for inner in formula.findall(contents_raw)}
		bags.setdefault(bag, contents)
	return bags

def up(name):
	for bag in bags: 
		if name in bags[bag].keys() and bag not in traversed:
			traversed.append(bag)
			count(bag)

def down(name):
	if len(bags[name]) == 0:
		return 0
	else:
		total = 0
		for inner in bags[name]:
			total += bags[name][inner] + bags[name][inner]*down(inner)
		return total  
		
traversed = []
bags = parse_raw()

if __name__ == '__main__':
	#count("shiny gold")
	#print(len(traversed))
	print(down("shiny gold"))

