package main

import "fmt"


type shape interface {
	getAria() float64
}

type square struct {
	sideLengt float64
}

type treangle struct {
	height float64
	base float64
}

func main () {
	sh := square{}
	sh.sideLengt = 10

	tr := treangle{}
	tr.base = 10
	tr.height = 10

	printAria(sh)
	printAria(tr)
}

func printAria(s shape) {
	fmt.Println(s.getAria())
}

func (sq square) getAria() float64 {
	return sq.sideLengt * sq.sideLengt
}

func (tr treangle) getAria() float64 {
	return 0.5 * tr.base * tr.height
} 