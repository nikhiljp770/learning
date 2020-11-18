package main

import "fmt"

type Sofa struct {
	cofeetable string
	chairs     string
}

type Diningtable struct {
	chair string
	table string
}

type Hall struct {
	sofa        Sofa
	diningtable Diningtable
}

type Bed struct {
	cot    string
	pillow string
}

type Bathroom struct {
	soap  string
	paste string
}

type Room struct {
	bed      Bed
	bathroom Bathroom
}

type Kitchen struct {
	stove    string
	groceres string
}

type House struct {
	hall    Hall
	room    Room
	kitchen Kitchen
}

func main() {

	s := Sofa{cofeetable: "wooden", chairs: "three"}
	d := Diningtable{chair: "three", table: "wood"}
	h := Hall{sofa: s, diningtable: d}
	b := Bed{cot: "bedsheet", pillow: "two"}
	bb := Bathroom{soap: "dove", paste: "colgate"}
	r := Room{bathroom: bb, bed: b}
	k := Kitchen{stove: "prestage", groceres: "onion"}
	ho := House{hall: h, room: r, kitchen: k}

	s.print()
	d.print()
	h.print()
	b.print()
	bb.print()
	r.print()
	ho.print()
}

func (s Sofa) print() {
	fmt.Printf("the sofa has %v, %v \n", s.cofeetable, s.chairs)
}

func (d Diningtable) print() {
	fmt.Printf("diningtable has %v, %v \n", d.chair, d.table)
}

func (h Hall) print() {
	fmt.Printf("hall has %v, %v \n", h.sofa, h.diningtable)
}

func (b Bed) print() {
	fmt.Printf("bed has %v, %v \n", b.cot, b.pillow)
}

func (bb Bathroom) print() {
	fmt.Printf("bathroom has %v, %v \n", bb.soap, bb.paste)
}

func (r Room) print() {
	fmt.Printf("room has %v, %v \n", r.bed, r.bathroom)
}

func (k Kitchen) print() {
	fmt.Printf("kitchen has %v, %v \n", k.stove, k.groceres)
}

func (ho House) print() {
	fmt.Printf("house has %v, %v,%v \n", ho.hall, ho.room, ho.kitchen)
}
