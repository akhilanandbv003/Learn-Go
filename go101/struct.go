package main

import "fmt"

type student struct {
	name string
	age  int
	class int
}

type classleader struct {
	name string
	student
	monitorstudents bool
}

func main() {
	s1 := student{
		name: "ak",
		age:  7,
	class: 10,
	}

	cl := classleader{
		name: "clname",
		student: student{
			age:  54,
			name: "bv",
		},
		monitorstudents:true,
	}

	fmt.Println(s1,cl)
	fmt.Println(s1.age,s1.name,s1.class)
	//fmt.Println(s2.age,s2.name,s2.class)// s2.class defaults to 0
	fmt.Println(cl.age,cl.name,cl.class,cl.student.name)// s2.class defaults to 0
	anonymousstruct()
}


func anonymousstruct(){
	fmt.Println("anonymousstruct")

	s1 := struct {
		name  string
		age int
	}{
		name:"akka",
		age:28,
	}
	fmt.Println(s1)
}
