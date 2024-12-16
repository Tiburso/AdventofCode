#! /usr/bin/env python3


#def get_answers(group):
#	count = 0
#	answered = []
#	for ans in group:
#		for x in ans:
#			if x not in answered:
#				count += 1
#				answered.append(x)
#
#	return count

def get_answers(group):
	all_answered = [m for m in group[0]]
	regular_answers = []

	for ans in group[1:]:
		i = 0
		for a in all_answered:
			if a not in ans and a not in regular_answers:
				regular_answers.append(a)

	result = list(set(all_answered) - set(regular_answers))
	return len(result)


total = 0
#get input
with open("input.txt","r") as f:
	group = []
	for x in f.read().splitlines():
		if x == "":
			total += get_answers(group)
			group = []
		else:
			group.append(x)
			

total += get_answers(group)

print(total)