package main

import (
	"time"
)

type clock struct {
	year, month, day int
	stop             chan bool
}

var service chan *clock

func Newc(year, month, day int) *clock {
	tmp := &clock{year: year, month: month, day: day, stop: make(chan bool)}
	service <- tmp
	return tmp
}

func NewcDefault() *clock {
	tmp := &clock{year: time.Now().Year(), month: int(time.Now().Month()), day: time.Now().Day(), stop: make(chan bool)}
	service <- tmp
	return tmp
}

func (this *clock) Update(year, month, day int) {
	this.stop <- true
	this.year, this.month, this.day = year, month, day
	service <- this
}

func (this *clock) UpdateDefault() {
	this.stop <- true
	tmp := time.Now()
	this.year, this.month, this.day = tmp.Year(), int(tmp.Month()), tmp.Day()
	service <- this
}

func (this *clock) Add() {
	month := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if this.year%4 == 0 || (this.year%400 == 0 && this.year%100 != 0) {
		month[1] = 29
	}
	if this.day == month[this.month-1] {
		this.day = 1
		if this.month == 12 {
			this.year, this.month = this.year+1, 1
		} else {
			this.month += 1
		}
	} else {
		this.day += 1
	}
}

/*func main() {
	service = make(chan *clock,3)
	go func() {
		for {
			select {
			case ck := <-service:
				go func() {
					for {
						select {
						case <- time.After(time.Second):
							fmt.Println(ck.year, ck.month, ck.day)
							//ck.Add()
						case <- ck.stop:
							return
						}
					}
				}()
			}
		}
	}()
	_ = NewcDefault()
	for {

	}
}*/
