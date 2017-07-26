package main

import (
	"fmt"

	"github.com/natejenson/LearnGo/Structs/animals"
)

func main() {
	c := animals.Animal{Name: "Cow", Hp: 7}
	s := &animals.Shark{animals.Animal{"Megaladon", 100}, 50}
	fmt.Printf("%v attempting to eat %v...\n", s.Name, c.Name)
	survival(s, c)
}

func survival(e animals.Eater, b animals.Animal) {
	c := "has not"
	if e.Eat(b) {
		c = "has"
	}
	fmt.Printf("...%v eaten %v", c, b.Name)
}
