#! /usr/bin/env python3

with open("input.txt","r") as f:
	raw = [eval(x) for x in f.read().splitlines()]

target = 105950735
target_i = 564

#https://stackoverflow.com/questions/29914913/sum-of-two-numbers-equal-to-the-given-number
def find_wrong(numbers, target):
	d = {}
	for i in range(len(numbers)):
		if (target-numbers[i]) in d.keys():
			return [numbers[i], target-numbers[i]]
		d[numbers[i]] = i
	return []


#https://www.geeksforgeeks.org/find-subarray-with-given-sum/
def find_set(numbers, target):
	curr_sum = numbers[0]
	start = 0
	n = len(numbers)

	i = 1
	while i <= n:
		while curr_sum > target and start < i-1:
			curr_sum -= numbers[start]
			start += 1

		if curr_sum == target:
			return numbers[start:i]

		if i < n:
			curr_sum += numbers[i]
		i += 1
   

if __name__ == '__main__':
	i = 25
	#while i < len(raw):
	#	if len(find_wrong(raw[i-25:i], raw[i])) == 0: 
	#		print(raw[i], i)
	#	i += 1
	sub = find_set(raw, target)
	print(min(sub) + max(sub))
