package main

import "fmt"

func main() {
	type gator int

	var g1 gator
	var x int

	g1 = 42
	x = 8

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println("------")
	fmt.Println(x)

}
