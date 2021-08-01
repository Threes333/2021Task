package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unsafe"
)

type MyString string

func (this MyString) One() bool {
	for _, v := range this {
		if v < rune('0') || v > rune('9') {
			fmt.Println(this, " 不是全数字")
			return false
		}
	}
	fmt.Println(this, " 是全数字")
	return true
}

func (this MyString) Two() bool {
	for i := 0; i < len(this); i++ {
		if this[i] < 'a' || this[i] > 'z' {
			fmt.Println(this, " 不是全小写字母")
			return false
		}
	}
	fmt.Println(this, " 是全小写字母")
	return true
}

func (this MyString) Three() bool {
	for i := 0; i < len(this); i++ {
		if this[i] < 'A' || this[i] > 'Z' {
			fmt.Println(this, " 不是全大写字母")
			return false
		}
	}
	fmt.Println(this, " 是全大写字母")
	return true
}

func (this *MyString) Four() {
	for i := 0; i < len(*this); i++ {
		if (*this)[i] < 'A' || (*this)[i] > 'z' || ((*this)[i] > 'Z' && (*this)[i] < 'a') {
			fmt.Println("非全字母")
			return
		}
	}
	ptr := (*reflect.StringHeader)(unsafe.Pointer(this))
	tmp := make([]byte, ptr.Len)
	for i := 0; i < len(*this); i++ {
		if (*this)[i] >= 'a' && (*this)[i] <= 'z' {
			tmp[i] = (*this)[i] - 32
		} else {
			tmp[i] = (*this)[i]
		}
	}
	ptr.Data = uintptr(unsafe.Pointer(&tmp[0]))
	fmt.Println("将小写字母转换为大写字母后： ", *this)
}

func (this MyString) Five() {
	var arr []int
	if this.One() {
		arr = make([]int, len(this))
		for i, v := range this {
			arr[i] = int(v - '0')
		}
	} else {
		fmt.Println("非全数字")
		return
	}
	sort.Ints(arr)
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func (this MyString) Six() {
	for i := 0; i < len(this); i++ {
		if this[i] < 'A' || this[i] > 'z' || (this[i] > 'Z' && this[i] < 'a') {
			fmt.Println("非全字母")
			return
		}
	}
	arr := []byte(this)
	sort.Slice(arr, func(i, j int) bool {
		fst, sec := arr[i], arr[j]
		if arr[i] < 'a' {
			fst += byte(32)
		}
		if arr[j] < 'a' {
			sec += byte(32)
		}
		return fst < sec
	})
	for _, v := range arr {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}

func (this MyString) split() (res []MyString) {
	tmp := strings.Split(string(this), ",")
	p1 := (*reflect.SliceHeader)(unsafe.Pointer(&res))
	p2 := (*reflect.SliceHeader)(unsafe.Pointer(&tmp))
	*p1 = *p2
	return
}

func (this MyString) Start() {
	this.One()
	this.Two()
	this.Three()
	this.Four()
	this.Five()
	this.Six()
}

func main() {
	var str MyString = "acbdw,1269547,AASIDX,AIUydjs,12sjaA,3819247,ausSHSzio,IUFISsi"
	l := str.split()
	for _, v := range l {
		v.Start()
	}
}
