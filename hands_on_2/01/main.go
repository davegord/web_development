package main

import "fmt"

func main() {
	s1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22}
	fmt.Println(s1)

	for i, _ := range s1 {
		fmt.Println(i)
	}

	fmt.Println("--------")

	for i, v := range s1 {
		fmt.Println(i, v)
	}
}
