package main

import "fmt"

func main() {
//These are called slices in GO
	var x[5] int
	fmt.Println(x)
	x[4] = 99
	fmt.Println(x)




}


/*
composite litral
x := type{value}
x := []int{9,22,99,32,76}

type slice of int
e*/