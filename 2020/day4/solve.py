#! /usr/bin/env python3
from switch import Switch

def find_right(passport):
	count = len(passport)

	for part in passport:
		tag, value = part.split(":")

		#comment the rest for 1st ex
		with Switch(tag) as case:
			if case("cid"):
				count -= 1
			if case("byr"):
				if eval(value) < 1920 or eval(value) > 2002:
					return 0
			if case("iyr"):
				if eval(value) < 2010 or eval(value) > 2020:
					return 0
			if case("eyr"):
				if eval(value) < 2020 or eval(value) > 2030:
					return 0
			if case("hgt"):
				if value[-2:] == "cm":
					if eval(value[:-2]) < 150 or eval(value[:-2]) > 193:
						return 0
				elif value[-2:] == "in":
					if eval(value[:-2]) < 59 or eval(value[:-2]) > 76:
						return 0
				else:
					return 0
			if case("hcl"):
				if len(value) != 7 or value[0] != "#" or not all(x in "abcdef" or x in "0123456789" for x in value[1:]):
					return 0
			if case("ecl"):
				if value not in ["amb","blu","brn","gry","grn","hzl","oth"]:
					return 0
			if case("pid"):
				if len(value) != 9:
					return 0

	
	if count == len(words):
		return 1

	return 0




total = 0
passport = []
words = ["byr","iyr","eyr", "hgt","hcl", "ecl", "pid"]
with open("input.txt","r") as f:
	for l in f.readlines():
		if l != "\n":
			passport += l[:-1].split(" ")
		else:
			total += find_right(passport)
			passport = []

	#last one didnt count
	total += find_right(passport)

print(total) 