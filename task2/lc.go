package main

import "strconv"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, v := range nums {
		if x, ok := hash[target-v]; ok {
			return []int{x, i}
		}
		hash[v] = i
	}
	return []int{}
}

func isValid(s string) bool {
	const (
		xiao  byte = 1
		zhong byte = 2
		da    byte = 3
	)
	stack := make([]byte, 0)
	for _, v := range s {
		length := len(stack)
		switch v {
		case '(':
			stack = append(stack, xiao)
		case '[':
			stack = append(stack, zhong)
		case '{':
			stack = append(stack, da)
		case ')':
			if length < 1 || stack[length-1] != xiao {
				return false
			}
			stack = stack[:length-1]
		case ']':
			if length < 1 || stack[length-1] != zhong {
				return false
			}
			stack = stack[:length-1]
		case '}':
			if length < 1 || stack[length-1] != da {
				return false
			}
			stack = stack[:length-1]
		}
	}
	return len(stack) == 0
}

func evalRPN(tokens []string) int {
	arr := make([]int, 0)
	for _, v := range tokens {
		length := len(arr) - 1
		switch v {
		case "+":
			arr[length-1] += arr[length]
			arr = arr[:length]
		case "-":
			arr[length-1] -= arr[length]
			arr = arr[:length]
		case "*":
			arr[length-1] *= arr[length]
			arr = arr[:length]
		case "/":
			arr[length-1] /= arr[length]
			arr = arr[:length]
		default:
			num, _ := strconv.Atoi(v)
			arr = append(arr, num)
		}
	}
	return arr[0]
}
