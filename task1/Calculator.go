package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	add = iota
	sub
	mul  = 3
	div  = 4
	left = -2
)

type Sign struct {
}

type calculator struct {
	Sign
}

var (
	ops  []int
	nums []float64
)

func (this *Sign) Add(a, b float64) float64 {
	return a + b
}

func (this *Sign) Sub(a, b float64) float64 {
	return a - b
}

func (this *Sign) Mul(a, b float64) float64 {
	return a * b
}

func (this *Sign) Div(a, b float64) float64 {
	return a / b
}

func (this *calculator) deal(str string) float64 {
	ops = []int{}
	nums = []float64{0}
	str = this.InitExp(str)
	var tmp float64
	var start int
	var flag bool = true
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '+':
			if flag {
				tmp, _ = strconv.ParseFloat(str[start:i], 64)
				nums = append(nums, tmp)

			} else {
				flag = true
			}
			for len(ops) > 0 && add <= ops[len(ops)-1] {
				this.SignalCalu()
			}
			ops = append(ops, add)
			start = i + 1
		case '-':
			if flag {
				tmp, _ = strconv.ParseFloat(str[start:i], 64)
				nums = append(nums, tmp)

			} else {
				flag = true
			}
			for len(ops) > 0 && sub <= ops[len(ops)-1]+1 {
				this.SignalCalu()
			}
			ops = append(ops, sub)
			start = i + 1
		case '*':
			if flag {
				tmp, _ = strconv.ParseFloat(str[start:i], 64)
				nums = append(nums, tmp)

			} else {
				flag = true
			}
			for len(ops) > 0 && mul <= ops[len(ops)-1] {
				this.SignalCalu()
			}
			ops = append(ops, mul)
			start = i + 1
		case '/':
			if flag {
				tmp, _ = strconv.ParseFloat(str[start:i], 64)
				nums = append(nums, tmp)

			} else {
				flag = true
			}
			for len(ops) > 0 && div <= ops[len(ops)-1]+1 {
				this.SignalCalu()
			}
			ops = append(ops, div)
			start = i + 1
		case '(':
			ops = append(ops, left)
			start = i + 1
		case ')':
			tmp, _ = strconv.ParseFloat(str[start:i], 64)
			nums = append(nums, tmp)
			for ops[len(ops)-1] != left {
				this.SignalCalu()
			}
			ops = ops[:len(ops)-1]
			flag = false
		}
	}
	tmp, _ = strconv.ParseFloat(str[start:], 64)
	fmt.Println(tmp, nums, ops)
	nums = append(nums, tmp)
	for len(ops) > 0 {
		if len(ops) > 1 && ops[len(ops)-2] >= ops[len(ops)-1]-1 {
			tmp := ops[len(ops)-1]
			ops = ops[:len(ops)-1]
			this.SignalCalu()
			ops = append(ops, tmp)
		}
		this.SignalCalu()
	}
	return nums[len(nums)-1]
}

func (this *calculator) InitExp(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "(+", "(0+")
	str = strings.ReplaceAll(str, "(-", "(0-")
	return str
}

func (this *calculator) SignalCalu() {
	if len(ops) < 1 || len(nums) < 2 {
		return
	}
	switch ops[len(ops)-1] {
	case sub:
		nums[len(nums)-2] = this.Sub(nums[len(nums)-2], nums[len(nums)-1])
		nums = nums[:len(nums)-1]
	case add:
		nums[len(nums)-2] = this.Add(nums[len(nums)-2], nums[len(nums)-1])
		nums = nums[:len(nums)-1]
	case mul:
		nums[len(nums)-2] = this.Mul(nums[len(nums)-2], nums[len(nums)-1])
		nums = nums[:len(nums)-1]
	case div:
		nums[len(nums)-2] = this.Div(nums[len(nums)-2], nums[len(nums)-1])
		nums = nums[:len(nums)-1]
	}
	fmt.Println(ops[len(ops)-1])
	ops = ops[:len(ops)-1]
}

func NewCalculator() *calculator {
	return &calculator{}
}

/*
func main() {
	str := "(1+2-(4+7*5.9)/2.7)-(4/2.8+5*3)  +1"
	ca := NewCalculator()
	fmt.Println(ca.deal(str))
}
*/
