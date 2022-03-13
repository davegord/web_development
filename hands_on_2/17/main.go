package main

import "fmt"

func main() {

	s := "I'm sorry Dave.  I can't do that"
	fmt.Println(s)
	bs := []byte(s)
	fmt.Println(bs)
	ss := string(bs)
	fmt.Println(ss)

	fmt.Println(ss[0:14])
	fmt.Println(ss[18:])

	for _, v := range ss {
		fmt.Println(v)
	}
}
