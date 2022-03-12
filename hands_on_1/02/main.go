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

func (p person) pSpeak() {
	fmt.Println("Hello my name is", p.first, p.last, "and my age is", p.age)

}

func (sa secretAgent) saSpeak() {
	fmt.Println("Shaken not stirred")
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

	fmt.Println("p1 first name is", p1.first)
	p1.pSpeak()
	//fmt.Println(sa1.licenseToKill)
	sa1.saSpeak()

	sa1.p.pSpeak()

}
