package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	p             person
	licenseToKill bool
}

func (p person) Speak() {
	fmt.Println("Hello my name is", p.first, p.last, "and my age is", p.age)

}

func (sa secretAgent) Speak() {
	fmt.Println("Shaken not stirred")
}

type superHuman interface {
	Speak()
}

func testsuperHuman(sh superHuman) {
	sh.Speak()
}

func main() {
	p1 := person{
		"Dave",
		"Gordon",
		42,
	}

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
			45,
		},
		true,
	}

	testsuperHuman(p1)
	testsuperHuman(sa1)

}
