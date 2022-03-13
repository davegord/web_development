package main

import "fmt"

type gator int

type flamingo bool

func (g gator) greeting() {
	fmt.Println("Hello I am a gator")

}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g1.greeting()

	var f1 flamingo
	f1 = true

	bayou(f1)
	bayou(g1)

}
