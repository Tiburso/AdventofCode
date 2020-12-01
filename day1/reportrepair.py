#! /usr/bin/env python3

import os 

f = open("report.txt", "r")


report = []
while(True):
	line = f.readline()
	if len(line) == 0:
		break
	report.append(eval(line))

f.close()

result = []
for i in range(len(report)):
	for l in range(i+1, len(report)):
		for t in range(l+1, len(report)):
			if report[i]+report[l]+report[t] == 2020:
				print(report[i]*report[l]*report[t])


