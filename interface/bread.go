package main

import "fmt"

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	value string
}

type StrawBerryJam struct {
}

type SpoonOfStrawberryJam struct {
}

type Nutella struct {
}

type SpoonOfNutella struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

func (s *SpoonOfNutella) String() string {
	return "+ nutella"
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.value += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.value
}

func (j *StrawBerryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (n *Nutella) GetOneSpoon() SpoonOfJam {
	return &SpoonOfNutella{}
}

func main() {
	bread := &Bread{}
	// jam := &StrawBerryJam{}
	jam := &Nutella{}

	bread.PutJam(jam)

	fmt.Println(bread)
}
