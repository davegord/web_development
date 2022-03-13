package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	//str := p.fName + " is walking"
	//return str
	return (fmt.Sprintln(p.fName, "is walking"))
}

func main() {

	p1 := person{
		fName:   "Dave",
		lName:   "Gordon",
		favFood: []string{"pizza", "wings", "sushi", "steak"},
	}

	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for _, v := range p1.favFood {
		fmt.Println(v)
	}

	str := p1.walk()
	fmt.Printf("%v", str)

}
