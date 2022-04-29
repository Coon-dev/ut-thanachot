package main

import (
	"fmt"
	"strconv"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

func possibleTimes(digits []int) int {
	// Your code here
	m := map[string]bool{
		fmt.Sprintf("%d%d:%d%d", digits[0], digits[1], digits[2], digits[3]): true,
		fmt.Sprintf("%d%d:%d%d", digits[0], digits[1], digits[3], digits[2]): true,
		fmt.Sprintf("%d%d:%d%d", digits[0], digits[2], digits[1], digits[3]): true,
		fmt.Sprintf("%d%d:%d%d", digits[0], digits[2], digits[3], digits[1]): true,
		fmt.Sprintf("%d%d:%d%d", digits[0], digits[3], digits[1], digits[2]): true,
		fmt.Sprintf("%d%d:%d%d", digits[0], digits[3], digits[2], digits[1]): true,
		fmt.Sprintf("%d%d:%d%d", digits[1], digits[0], digits[2], digits[3]): true,
		fmt.Sprintf("%d%d:%d%d", digits[1], digits[0], digits[3], digits[2]): true,
		fmt.Sprintf("%d%d:%d%d", digits[1], digits[2], digits[0], digits[3]): true,
		fmt.Sprintf("%d%d:%d%d", digits[1], digits[2], digits[3], digits[0]): true,
		fmt.Sprintf("%d%d:%d%d", digits[1], digits[3], digits[0], digits[2]): true,
		fmt.Sprintf("%d%d:%d%d", digits[1], digits[3], digits[2], digits[0]): true,
		fmt.Sprintf("%d%d:%d%d", digits[2], digits[1], digits[0], digits[3]): true,
		fmt.Sprintf("%d%d:%d%d", digits[2], digits[1], digits[3], digits[0]): true,
		fmt.Sprintf("%d%d:%d%d", digits[2], digits[0], digits[1], digits[3]): true,
		fmt.Sprintf("%d%d:%d%d", digits[2], digits[0], digits[3], digits[1]): true,
		fmt.Sprintf("%d%d:%d%d", digits[2], digits[3], digits[1], digits[0]): true,
		fmt.Sprintf("%d%d:%d%d", digits[2], digits[3], digits[0], digits[1]): true,
		fmt.Sprintf("%d%d:%d%d", digits[3], digits[1], digits[2], digits[0]): true,
		fmt.Sprintf("%d%d:%d%d", digits[3], digits[1], digits[0], digits[2]): true,
		fmt.Sprintf("%d%d:%d%d", digits[3], digits[2], digits[1], digits[0]): true,
		fmt.Sprintf("%d%d:%d%d", digits[3], digits[2], digits[0], digits[1]): true,
		fmt.Sprintf("%d%d:%d%d", digits[3], digits[0], digits[1], digits[2]): true,
		fmt.Sprintf("%d%d:%d%d", digits[3], digits[0], digits[2], digits[1]): true,
	}

	answer := 0
	for key := range m {
		hour, err := strconv.ParseUint(key[0:2], 10, 64)
		if err != nil {
			panic("hour error")
		}

		minute, err := strconv.ParseUint(key[3:5], 10, 64)
		if err != nil {
			panic("hour error")
		}

		if hour >= 0 && hour < 24 && minute >= 0 && minute < 60 {
			// valid time
			answer++
		}
	}
	return answer
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))

	fmt.Println(possibleTimes([]int{0, 0, 0, 0}))
	fmt.Println(possibleTimes([]int{0, 0, 1, 1}))
}
