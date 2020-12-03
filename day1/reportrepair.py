#! /usr/bin/env python3

report = [eval(n) for n in open("input.txt", "r")]

f.close()

for i in range(len(report)):
	for l in range(i+1, len(report)):
		for t in range(l+1, len(report)):
			if report[i]+report[l]+report[t] == 2020: 
				print(report[i]*report[l]*report[t])


