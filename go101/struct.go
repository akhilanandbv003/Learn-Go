package main

import "fmt"

type student struct {
	name string
	age  int
	class int
}

func main() {
	s1 := student{
		name: "ak",
		age:  7,
	class: 10,
	}

	s2 := student{
		name: "bv",
		age:  54,
	}
	fmt.Println(s1,s2)
	fmt.Println(s1.age,s1.name,s1.class)
	fmt.Println(s2.age,s2.name,s2.class)// s2.class defaults to 0
}
