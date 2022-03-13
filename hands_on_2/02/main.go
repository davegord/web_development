package main

import "fmt"

func main() {
	m1 := map[string]int{
		"Dave":   42,
		"Joe":    37,
		"Jess":   41,
		"Rachel": 9,
		"Blake":  5,
	}
	fmt.Println(m1)
	fmt.Println("=================")

	for i, _ := range m1 {
		fmt.Println(i)
	}

	fmt.Println("=================")

	for i, v := range m1 {
		fmt.Printf("%v \t\t%v\n", i, v)
	}

}
