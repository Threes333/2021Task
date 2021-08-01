package main

import (
	"fmt"
)

type NodeType int

var signs chan bool

type Node struct {
	Val        NodeType
	Next, Last *Node
}

type Linklist struct {
	begin, end *Node
}

func (this *Linklist) Get() {
	if this.begin == nil {
		fmt.Println("Empty")
		return
	}
	var res string
	for node := range this.iter() {
		res += fmt.Sprintf("%v ⇆ ", node.Val)
	}
	res += fmt.Sprintf("(%v)", this.begin.Val)
	fmt.Println(res)
}

func (this *Linklist) Update(OldVal, NewVal NodeType) {
	for node := range this.iter() {
		if node.Val == OldVal {
			node.Val = NewVal
			return
		}
	}
	fmt.Println("Not found")
}

func (this *Linklist) Delete(val NodeType) {
	if this.end == nil {
		fmt.Println("Empty")
		return
	}
	var idx *Node
	for node := range this.iter() {
		if node.Val == val {
			idx = node
			signs <- true
		}
	}
	idx.Last.Next = idx.Next
	idx.Next.Last = idx.Last
}

func (this *Linklist) TailAdd(val NodeType) {
	if this.end == nil {
		tmp := &Node{Val: val}
		tmp.Next, tmp.Last = tmp, tmp
		this.begin, this.end = tmp, tmp
	} else {
		tmp := &Node{val, this.begin, this.end}
		this.end.Next, this.begin.Last = tmp, tmp
		this.end = tmp
	}
}

func (this *Linklist) iterator() []*Node {
	res := make([]*Node, 0)
	if this.begin != nil {
		res = append(res, this.begin)
		for tmp := this.begin.Next; tmp != this.begin; tmp = tmp.Next {
			res = append(res, tmp)
		}
	}
	return res
}

func (this *Linklist) iter() chan *Node {
	msg := make(chan *Node, 1)
	go func() {
		if this.begin != nil {
			msg <- this.begin
			tmp := this.begin.Next
			for {
				select {
				case <-signs:
					close(msg)
					return
				default:
					if tmp == this.begin {
						signs <- true
					} else {
						msg <- tmp
						tmp = tmp.Next
					}
				}
			}
		}
	}()
	return msg
}

func NewLinklist() *Linklist {
	return &Linklist{}
}

func init() {
	signs = make(chan bool, 1)
}

/*
func main(){
	l := NewLinklist()
	l.Get()
	l.TailAdd(1)
	l.TailAdd(2)
	l.TailAdd(3)
	l.Get()
	l.Update(1,4)
	l.Get()
	l.Delete(3)
	l.Get()
}
*/
