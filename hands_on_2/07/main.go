package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	v         vehicle
	fourWheel bool
}

type sedan struct {
	v      vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {

	return (fmt.Sprintln("I am a", s.v.color, "truck with", s.v.doors, "doors and I am used to drive people around"))

}

func (t truck) transportationDevice() string {
	return (fmt.Sprintln("I am a", t.v.color, "truck", "with", t.v.doors, "doors and I am used to haul stuff"))
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())

}

func main() {
	t1 := truck{
		vehicle{
			doors: 2,
			color: "silver",
		},
		true,
	}

	s1 := sedan{
		vehicle{
			doors: 4,
			color: "black",
		},
		false,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(s1.v.color)
	fmt.Println(t1.fourWheel)

	fmt.Println(s1.transportationDevice())
	fmt.Println(t1.transportationDevice())

	fmt.Println("==============")
	report(t1)
	fmt.Println("==============")
	report(s1)

}
